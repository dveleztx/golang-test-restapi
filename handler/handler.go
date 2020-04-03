package handler

import (
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"net/http"
	"strings"
	"strconv"
	"testrest/model"
)

func Welcome() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.String(http.StatusOK, "Welcome to Wigel Mapping API!")
	}
}

func Index() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.Render(http.StatusOK, "index.html", map[string]interface{} {
			"title": "Home",
			"bodyheader": "Dashboard",
		})
	}
}

func JSONLoad_GET(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {

		var tables []string
		if err := db.Table("information_schema.tables").Where("table_schema = ?", "public").Pluck("table_name", &tables).Error; err != nil {
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
		if err := db.Table("information_schema.tables").Where("table_schema = ?", "public").Pluck("table_name", &tables).Error; err != nil {
			panic(err)
		}

		return c.Render(http.StatusOK, "csvload.html", map[string]interface{}{
			"title": "CSV Loader",
			"bodyheader": "CSV Loader",
			"database": tables,
		})
	}
}

// POST
func CSVLoad_POST(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {

		csv := strings.Split(c.FormValue("csv"), ",")
		if len(csv) != 5 {
			panic("Error")
		}

		id, _ := strconv.ParseUint(csv[0], 10, 64)

		user := model.User{
			ID: id,
			Name: csv[1],
			Age: csv[2],
			CreatedAt: csv[3],
			UpdatedAt: csv[4],
		}

		db.Create(&user)

		// List of Tables in Database
		tables := []string {
			"users",
		}

		return c.Render(http.StatusOK, "csvload.html", map[string]interface{}{
			"title": "CSV Loader",
			"bodyheader": "CSV Loader",
			"database": tables,
		})
	}
}

func CreateUser(db *gorm.DB, user model.User) error {
	return db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&user).Error; err != nil {
			return err
		}
		return nil
	})
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
