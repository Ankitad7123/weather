package main

import (
	"encoding/json"
	"fmt"
	"net/http"


	"github.com/gin-gonic/gin"
)

type apiConfig12 struct {
  OpenWeatherApiKey string `json:"OpenWeatherApiKey"`
 }

type weatherD struct {
 Name string `json:"name"`
 Main struct{
      Kelvin float64 `json:"temp"`
  }`json:"main"`

}



func query(c *gin.Context) {
 city := c.Param("city")

  res , err := http.Get("https://api.openweathermap.org/data/2.5/weather?APPID=" + "278395a12c454e144877bfa0196c83be"+ "&q=" + city)
  fmt.Print(res.Body)
    if err != nil {
    c.JSON(400 , gin.H{"err":err.Error()})
    return
  }
  defer res.Body.Close()

  var weather weatherD
  if err := json.NewDecoder(res.Body).Decode(&weather);err != nil {
    c.JSON(400 , gin.H{"err":err.Error()})
    return
  }



  c.JSON(200 , gin.H{"err":weather})



}


func main(){
  r := gin.Default()
  r.GET("/:city"  , query)
  r.Run()
}

