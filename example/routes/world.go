package main

import "github.com/labstack/echo/v4"

func Get_world(c echo.Context) error {
	return c.String(200, "World")
}