package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)
func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello")
}

//getAnimal display paramerter : http://127.0.0.1:8888/animal?name=tipu&type=dog
func getAnimal(c echo.Context) error{
	animalName := c.QueryParam("name")
	animalType := c.QueryParam("type")
	return c.String(http.StatusOK, fmt.Sprintf("Name: %s , Type: %s", animalName, animalType))
}
//getAnimalWithType display output in different format
// animal/string?name=tipu&type=dog  ==> Name: tipu , Type: dog
// animal/json?name=tipu&type=dog  ==> {"name":"tipu","type":"dog"}
func getAnimalWithType(c echo.Context) error {
	
	// get value of /:data
	outputType := c.Param("data")

	animalName := c.QueryParam("name")
	animalType := c.QueryParam("type")

	if outputType == "string"{
		return c.String(http.StatusOK, fmt.Sprintf("Name: %s , Type: %s", animalName, animalType))
	}
	if outputType == "json"{
		return c.JSON(http.StatusBadRequest, map[string]string{"name": animalName, "type":animalType})
	}
   return c.JSON(http.StatusBadRequest, map[string]string{"error":"type does not match"})
}

func main() {
	fmt.Println("hello")

	e := echo.New()
	e.GET("/", hello)

	//  /animal?name=tipu&type=dog
	e.GET("/animal",getAnimal)

	//  /animal/string?name=tipu&type=dog
    //  /animal/json?name=tipu&type=dog  
	e.GET("/animal/:data",getAnimalWithType)

	e.Logger.Fatal(e.Start(":8888"))
}