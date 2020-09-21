package request

type UserLoginReq struct {
	Email       string `json:"email" validate:"email"`
	Password    string `json:"password" validate:"min=6"`
	LastLoginIP string `json:"lastLoginIP" validate:"max=15"`
}

type UserLoginResp struct {
	Nickname string `json:"nickname"`
	Token    string `json:"token"`
}

type UserRegisterReq struct {
	UserLoginReq
	Nickname string `json:"nickname" validate:"min=0,max=100"`
	Code     string `json:"code" validate:"min=4,max=8"`
}

type UserRegisterResp struct {
}
