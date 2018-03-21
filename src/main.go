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
func serveData(ctx iris.Context) {
	var name = ctx.Params().Get("name")
	ctx.ViewData("Name", name)
	ctx.View("name.html")
	fmt.Printf("%v : Served Page /%v", time.Now(), name)
}
func main() {
	app := iris.New()

	app.RegisterView(iris.Django("./templates", ".html").Reload(true))
	app.Get("/", serveHello)
	app.Get("/hi/{name:string}", serveData)
	app.Run(addr)
	fmt.Printf("%v : Server started on %v", time.Now(), addr)
	fmt.Println("Exited, see you soon!")
}
