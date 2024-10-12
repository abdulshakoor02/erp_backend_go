package user

type User struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserData struct {
	Id       string   `json:"id"`
	Role     string   `json:"role"`
	Fullname string   `json:"fullName"`
	Username string   `json:"username"`
	Email    string   `json:"email"`
	Features []string `json:"features"`
	TenantId string   `json:"tenant_id"`
}
