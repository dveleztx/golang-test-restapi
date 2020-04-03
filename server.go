package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"log"
	"testrest/datastore"
	"testrest/handler"
	"testrest/tempengine"
)

func main() {

	/* GORM */
	db, err := datastore.NewDB()
	logFatal(err)

	// For debugging GORM
	db.LogMode(true)
	defer db.Close()

	/* Echo Framework */
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// CORS
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))

	// Handler
	e.GET("/", handler.Welcome())
	e.GET("/users", handler.GetAllUsers(db))
	e.GET("/users/:name", handler.GetUser(db))

	// Template
	templates := tempengine.GetTemplates()
	t := &tempengine.TemplateRegistry {
		Templates: templates,
	}
	e.Renderer = t

	// Render Template HTMLs
	e.GET("/index.html", handler.Index())
	e.GET("/jsonload.html", handler.JSONLoad_GET(db))
	e.GET("/csvload.html", handler.CSVLoad_GET(db))

	//e.POST("/jsonload.html", handler.JSONLoad_POST(db))
	e.POST("/csvload.html", handler.CSVLoad_POST(db))

	// Static
	e.Static("/static", "assets")

	// Create user test
	/*
	var user = model.User {
		ID: 6,
		Name: "Cody",
		Age: "43",
		CreatedAt: "2018-09-18",
		UpdatedAt: "2011-07-17",
	}
	handler.CreateUser(db, user)
	*/

	e.Logger.Fatal(e.Start(":1324"))
}

func logFatal(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
