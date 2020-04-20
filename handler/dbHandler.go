package handler

import (
	"fmt"
	"strings"
	"testrest/model"

	"github.com/jinzhu/gorm"
)

//func TableHandler(fields ...interface{}) model.User {
//func TableHandler(table, data) model.User {
func TableHandler(db *gorm.DB, table string, data []string) {

	// Debugging
	fmt.Println(table)

	/* Handler Logic */
	if table == "users" {

		for index, line := range data {
			if strings.Count(line, ",") != 1 {
				fmt.Printf("\n*******************************\n" +
				"Not correct amount of entries!\nLine: %v\nEntry: %v\n" +
				"*******************************\n", index, line)
			} else {
				field := strings.Split(line, ",")
				user := model.User {
					Name: field[0],
					Age: field[1],
				}
				createEntry(db, user)
			}
		}
	}
}

func createEntry (db *gorm.DB, dbObject model.User) error {
	return db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&dbObject).Error; err != nil {
			return err
		}
		return nil
	})
}
