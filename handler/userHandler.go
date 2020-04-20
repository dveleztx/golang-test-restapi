package handler

import (
	"net/http"
	"strings"
	"testrest/model"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
)

func Index() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.Render(http.StatusOK, "index.html", map[string]interface{} {
			"title": "Home",
			"bodyheader": "Dashboard",
		})
	}
}
/******************************************************************************
 * users_table.html
 *****************************************************************************/
// GET
func UsersTable_GET(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {

		// Setting to singular, otherwise will search "users_by_ids"
		db.SingularTable(true)

		var users []*model.UsersByID
                if err := db.Find(&users).Error; err != nil {
                        return err
                }

		return c.Render(http.StatusOK, "users_table.html", map[string]interface{}{
			"title": "Users Table",
			"bodyheader": "Users Table",
			"users": users,
		})
	}
}

/******************************************************************************
 * jsonload.html
 *****************************************************************************/
// GET
func JSONLoad_GET(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {

		var tables []string
		if err := db.Table("information_schema.tables").Where("table_type = ? AND table_schema = ?", "BASE TABLE", "public").Pluck("table_name", &tables).Error; err != nil {
			panic(err)
		}

		return c.Render(http.StatusOK, "jsonload.html", map[string]interface{}{
			"title": "JSON Loader",
			"bodyheader": "JSON Loader",
			"database": tables,
		})
	}
}
/******************************************************************************
 * csvload.html
 *****************************************************************************/
// GET
func CSVLoad_GET(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {

		var tables []string
		if err := db.Table("information_schema.tables").Where("table_type = ? AND table_schema = ?", "BASE TABLE", "public").Pluck("table_name", &tables).Error; err != nil {
			panic(err)
		}

		return c.Render(http.StatusOK, "csvload.html", map[string]interface{}{
			"title": "CSV Loader",
			"bodyheader": "CSV Loader",
			"table": tables,
		})
	}
}

// POST
func CSVLoad_POST(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {

		// Setup Dropdown
		var tables []string
		if err := db.Table("information_schema.tables").Where("table_type = ? AND table_schema = ?", "BASE TABLE", "public").Pluck("table_name", &tables).Error; err != nil {
			panic(err)
		}

		// Retrieve Values
		var entries []string
		table := c.FormValue("table")
		fields := strings.Split(c.FormValue("csv"), "\n")
		for _, field := range fields {
			entries = append(entries, field)
		}

		// Database Handler
		TableHandler(db, table, entries)

		// TODO: Create Table Struct with table name, fields, and number of insertable columns

		return c.Render(http.StatusOK, "csvload.html", map[string]interface{}{
			"title": "CSV Loader",
			"bodyheader": "CSV Loader",
			"table": tables,
		})
	}
}

func GetAllUsers(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {

		var user []*model.User
		if err := db.Find(&user).Error; err != nil {
			return err
		}
		return c.JSON(http.StatusOK, user)
	}
}

func GetUser(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {

		//var user []*model.User
		var user model.User
		db.Where("name = ?", c.Param("name")).Find(&user)

		return c.JSON(http.StatusOK, user)
	}
}
