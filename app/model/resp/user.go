package resp

type UserInfo struct {
	ID        int    `json:"id"`
	Account   string `json:"account"`
	NickName  string `json:"nick_name"`
	Role      int    `json:"role"`
	Status    int    `json:"status"`
	CreatedAt string `json:"created_at"`
}
