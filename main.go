package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"net/http"
	"strconv"
	"zwlinc.com/quickshare/ent"
)

func init() {
	db, err := gorm.Open(sqlite.Open("database.db"), &gorm.Config{})
	if err != nil {
		panic("fail to connect database")
	}
	db.AutoMigrate(&ent.Gist{})
}

func main() {
	e := echo.New()
	e.Static("/","dist")
	e.Use(middleware.CORSWithConfig(middleware.DefaultCORSConfig))
	e.GET("/gists",getGists)
	e.POST("/gists",saveGist)
	e.DELETE("gists/:id",delGist)
	e.Logger.Fatal(e.Start(":80"))
}

func getGists(c echo.Context) error{
	db, _ := gorm.Open(sqlite.Open("database.db"), &gorm.Config{})
	var gists []ent.Gist
	db.Find(&gists)
	return c.JSON(http.StatusOK,gists)
}

func delGist(c echo.Context) error{
	id,_ := strconv.Atoi(c.Param("id"))

	db, _ := gorm.Open(sqlite.Open("database.db"), &gorm.Config{})

	db.Delete(&ent.Gist{},id)

	return c.String(http.StatusOK,"Deleted")
}


func saveGist(c echo.Context) error{
	gist := struct {
		Title string `json:"title"`
		Text string `json:"text"`
	}{}

	c.Bind(&gist)
	g := ent.Gist{Title: gist.Title,Text: gist.Text}
	db, _ := gorm.Open(sqlite.Open("database.db"), &gorm.Config{})

	db.Create(&g)

	return c.JSON(200,g)
}
