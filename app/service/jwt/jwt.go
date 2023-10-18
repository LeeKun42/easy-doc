package jwt

import (
	"context"
	"easy-doc/app/lib/redis"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/viper"
	"strconv"
	"time"
)

type Service struct {
}

func NewService() *Service {
	return &Service{}
}

type CustomClaims struct {
	UserId     int   `json:"user_id"`
	RefreshTtl int64 `json:"refresh_ttl"`
	jwt.RegisteredClaims
}

func (js *Service) Create(userId int, refreshTtl int64) string {
	//生成token
	now := time.Now()              //当前时间
	ttl := viper.GetInt("jwt.ttl") //token有效期（分钟）
	if refreshTtl == 0 {
		refreshTtl = now.Add(time.Minute * time.Duration(viper.GetInt64("jwt.refresh_ttl"))).Unix() //token刷新有效期（分钟）
	}

	exp := now.Add(time.Minute * time.Duration(ttl)) //过期时间
	//自定义jwt body内容
	claims := CustomClaims{
		userId,
		refreshTtl,
		jwt.RegisteredClaims{
			Issuer:    "easy-doc",
			Subject:   "client",
			Audience:  nil,
			ExpiresAt: jwt.NewNumericDate(exp),
			NotBefore: jwt.NewNumericDate(now),
			IssuedAt:  jwt.NewNumericDate(now),
			ID:        strconv.Itoa(userId),
		},
	}
	//生成jwt签名
	tk := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, _ := tk.SignedString([]byte(viper.GetString("jwt.secret")))

	//写入redis 用以实现刷新token，手动设置token过期功能
	ctx, _ := context.WithTimeout(context.Background(), time.Millisecond*500)
	key := "jwt:token:" + token
	redis.Cache().Set(ctx, key, exp.Unix(), time.Minute*time.Duration(ttl))

	return token
}

func (js *Service) Check(tokenString string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(viper.GetString("jwt.secret")), nil
	})
	if err != nil {
		return &CustomClaims{}, errors.New("token无效")
	}
	claims, ok := token.Claims.(*CustomClaims)
	if ok && token.Valid {
		//检查是否在redis中存在
		ctx, _ := context.WithTimeout(context.Background(), time.Millisecond*500)
		key := "jwt:token:" + tokenString
		_, err := redis.Cache().Get(ctx, key).Result()
		if err != nil { //
			fmt.Println("jwt Check redis get err：", err.Error())
			return &CustomClaims{}, errors.New("token无效")
		} else {
			//todo 检查密码是否已修改
			return claims, nil
		}
	} else {
		return &CustomClaims{}, errors.New("token无效")
	}
}

func (js *Service) Refresh(oldToken string) (string, error) {
	//检查旧token
	var claims *CustomClaims
	var err error
	if claims, err = js.Check(oldToken); err != nil {
		return "", errors.New("token无效")
	}

	if claims.RefreshTtl < time.Now().Unix() {
		return "", errors.New("token无效")
	}

	//创建新token
	token := js.Create(claims.UserId, claims.RefreshTtl)

	//删除旧token
	js.Invalidate(oldToken)

	return token, nil
}

func (js *Service) Invalidate(token string) error {
	ctx, _ := context.WithTimeout(context.Background(), time.Millisecond*500)
	key := "jwt:token:" + token
	return redis.Cache().Del(ctx, key).Err()
}
