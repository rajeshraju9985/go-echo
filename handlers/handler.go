package handler

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"gopkg.in/mgo.v2"
	"strings"
)

type (
	Handler struct {
		DB *mgo.Session
		TeamClient string
	}
)

func userFromToken(c echo.Context) string {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	return claims["sub"].(string)
}

func teamsFromToken(c echo.Context) []string {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	teams:= strings.Split(strings.Replace(strings.Replace(strings.Replace(claims["teams"].(string), "[", "", -1),"]","",-1)," ","",-1),",")
	return teams
}