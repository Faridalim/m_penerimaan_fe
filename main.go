package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/go-git/go-git/v5"
	"github.com/labstack/echo/v4"
)

func main() {
	fmt.Println("OK")
	e := echo.New()
	e.GET("/clone", gitClone)
	e.GET("/:namefile", pageHandler)
	e.Logger.Fatal(e.Start("127.0.0.1:3002"))
}

func pageHandler(c echo.Context) error {
	namefile := c.Param("namefile")
	namefile = strings.TrimSuffix(namefile, "/")

	// membuat check permission utk path

	cookie := new(http.Cookie)
	if strings.Contains(namefile, "app") {
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

// Bisa digunakan untuk membuat path permission dan file permission

func gitClone(c echo.Context) error {
	_, err := git.PlainClone("/result", false, &git.CloneOptions{
		URL:      "https://github.com/go-git/go-git",
		Progress: os.Stdout,
	})

	if err != nil {
		return c.JSON(403, map[string]string{
			"message": err.Error(),
		})
	}

	return c.JSON(200, map[string]string{
		"message": "sukses",
	})
}
