package main

import (
	"fmt"
	"time"

	"github.com/kataras/iris"
)

var addr = iris.Addr("localhost:8080")

func serveHello(ctx iris.Context) {
	ctx.ViewData("Message", "Hello World!")
	ctx.View("index.html")
	fmt.Printf("%v : Served Page /", time.Now())
}
func main() {
	app := iris.New()

	app.RegisterView(iris.Django("./templates", ".html").Reload(true))
	app.Get("/", serveHello)
	app.Run(addr)
	fmt.Printf("%v : Server started on %v", time.Now(), addr)
	fmt.Println("Exited, see you soon!")
}
