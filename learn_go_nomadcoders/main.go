package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/JeonJe/learngo/scrapper"
	"github.com/labstack/echo"
)

const fileName string = "jobs.csv"
// Handler
func hello(c echo.Context) error {
	return c.File("home.html")
  }
  
func handleScrape(c echo.Context) error {
	defer os.Remove(fileName)
	term := strings.ToLower( scrapper.CleanString( c.FormValue("term")))
	fmt.Println(term)
	scrapper.Scrape(term)

	return c.Attachment("jobs.csv", "job.csv")
  }
  
func main(){
	// Echo instance
	e := echo.New()
  	// Routes
  	e.GET("/", hello)
  	e.POST("/scrape", handleScrape)
	  // Start server
	e.Logger.Fatal(e.Start(":1323"))
}