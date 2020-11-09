package main

import (
	"net/http"
	"os"
	"strconv"

	"github.com/labstack/echo/v4"
)

func main() {
	// port is set to "MY_APP_PORT" value
	port := os.Getenv("MY_APP_PORT")
	// if there is no value assigned, it is  port 8080
	if port == "" {
		port = "8080"
	}

	//new connection setup
	e := echo.New()
	products := []map[int]string{{1: "mobiles"}, {2: "tv"}, {3: "laptops"}}

	//set GET method (end point, GET method expects Context and error)
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Well, hello there !!")
	})
	e.GET("/products/:id", func(c echo.Context) error {
		//create a map
		var product map[int]string
		// we get one of the products (key/value pair)
		for _, p := range products {
			// for all key/value pairs
			for k := range p {
				pID, err := strconv.Atoi(c.Param("id"))
				if err != nil {
					return err
				}
				//if id equals to k
				if pID == k {
					//assign
					product = p
				}
			}
		}
		if product == nil {
			return c.JSON(http.StatusNotFound, "product not found")
		}
		return c.JSON(http.StatusOK, product)
	})
	// start
	e.Logger.Print("Listening on port 8080S oordword")
	e.Logger.Fatal(e.Start(":8080"))
}
