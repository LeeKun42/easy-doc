package request

type RegisterUser struct {
	Account  string `json:"account"`
	NickName string `json:"nick_name"`
	Passwd   string `json:"passwd"`
}

type UserLogin struct {
	Account string `url:"account"`
	Passwd  string `url:"passwd"`
}

type UserSearch struct {
	PageIndex int `url:"page_index"`
	PageSize  int `url:"page_size"`
}

type UserUpdate struct {
	ID       int    `json:"id"`
	NickName string `json:"nick_name"`
	Account  string `json:"account"`
}

type UserUpdatePassword struct {
	ID          int    `json:"id"`
	OldPassword string `json:"old_password"`
	NewPassword string `json:"new_password"`
}
