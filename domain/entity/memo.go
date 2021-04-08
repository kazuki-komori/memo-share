package entity

type Memo struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	UserID      string `json:"user_id"`
	SubjectID   string `json:"subject_id"`
}
