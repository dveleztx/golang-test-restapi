package handler

import (
	"error"
	"fmt"
	"testrest/model"
)

//func TableHandler(fields ...interface{}) model.User {
func TableHandler(table, data) model.User {

	// Debugging
	fmt.Println(table)

	/* Handler Logic */
	if table == "users" {

		fields := strings.Split(c.FormValue(data), ",")
		fmt.Println(fields)

                if len(csv) != 2 {
                        panic(err)
                }

                user := model.User{
                        Name: csv[0],
                        Age: csv[1],
                }

	}

	return User
}
