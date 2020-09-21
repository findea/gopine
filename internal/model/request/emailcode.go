package request

type SendEmailCodeReq struct {
	Email string `json:"email" validate:"email"`
}

type SendEmailCodeResp struct {
}
