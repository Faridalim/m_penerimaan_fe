package main

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/labstack/echo/v4"
)

func main() {
	fmt.Println("OK")
	e := echo.New()
	e.GET("/:namefile", pageHandler)
	e.Logger.Fatal(e.Start("127.0.0.1:3001"))
}

func pageHandler(c echo.Context) error {
	namefile := c.Param("namefile")
	namefile = strings.TrimSuffix(namefile, "/")
	cookie := new(http.Cookie)
	if strings.Contains(namefile, "index.html") {
		cookie.Name = "last_path"
		cookie.Value = namefile
		cookie.Expires = time.Now().Add(5 * time.Second)
		c.SetCookie(cookie)
		c.Response().Header().Set("Cache-Control", "max-age:2000, public")
		return c.File("pages/public/index.html")
	} else {
		c.Response().Header().Set("Cache-Control", "max-age:2000, public")
		return c.File("pages/public/" + namefile)
	}
}

// func checkFile(filename string) bool {
// 	if strings.Contains(namefile, ".js") || strings.Contains(namefile, ".js")
// }
