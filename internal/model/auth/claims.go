package auth

type UserClaims struct {
	UserID int64  `json:"userID"`
	Email  string `json:"email"`
}
