package model

type User struct {
	ID		uint64		`gorm:"primary_key" json:"id"`
	Name		string		`json:"name"`
	Age		string		`json:"age"`
	CreatedAt	string		`json:"created_at"`
	UpdatedAt	string		`json:"updated_at"`
}
