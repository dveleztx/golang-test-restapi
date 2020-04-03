package handler

import (
	"fmt"
	"testrest/model"
)

func UserHandler(fields ...interface{}) model.User {

	fmt.Println(fields)

	var User = model.User {
		ID: 6,
		Name: "Cody",
		Age: "43",
		CreatedAt: "2018-09-18",
		UpdatedAt: "2011-07-17",
	}

	return User
}
