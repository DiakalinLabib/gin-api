package main

import (
	"github.com/gin-gonic/gin"
	"html/template"
)

func main() {
	r := gin.Default()

	funcMap := template.FuncMap{
		// The name "inc" is what the function will be called in the template text.
		"add": func(x int, y int) int {
			return x + y
		},
	}

	r.SetFuncMap(funcMap)
	r.StaticFile("/assets", "./assets")
	r.LoadHTMLGlob("views/**/*.html")

	Routes(r)

	r.Run()
}
