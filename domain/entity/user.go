package entity

type User struct {
	ID       string `json:"id"`
	UserName string `json:"user_name"`
	FullName string `json:"full_name"`
	Mail     string `json:"mail"`
}
