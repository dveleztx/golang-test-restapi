package main

import (
	"log"
	"testrest/datastore"
	"testrest/handler"
	"testrest/tempengine"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
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

	// Template
	templates := tempengine.GetTemplates()
	t := &tempengine.TemplateRegistry {
		Templates: templates,
	}
	e.Renderer = t

	/* Handlers */
	// Home
	e.GET("/", handler.Welcome())

	// Users
	e.GET("/users", handler.GetAllUsers(db))
	e.GET("/users/name/:name", handler.GetUser(db))

	/* Render API FrontEnd */
	// GET Requests
	e.GET("/index.html", handler.Index())
	e.GET("/tables", handler.UsersTable_GET(db))
	e.GET("/jsonload.html", handler.JSONLoad_GET(db))
	e.GET("/csvload.html", handler.CSVLoad_GET(db))

	// POST Requests
	//e.POST("/jsonload.html", handler.JSONLoad_POST(db))
	e.POST("/csvload.html", handler.CSVLoad_POST(db))

	// Static
	e.Static("/static", "assets")

	/* Start Web Microservice */
	e.Logger.Fatal(e.Start(":1324"))
}

func logFatal(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
