package model

type User struct {
	Name		string		`json:"name"`
	Age		string		`json:"age"`
}

type UsersByID struct {
	ID		uint64		`gorm:"primary_key" json:"id"`
	Name		string		`json:"name"`
	Age		string		`json:"age"`
	CreatedAt	string		`json:"created_at"`
	ModifiedAt	string		`json:"modified_at"`
}
