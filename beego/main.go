package main

import (
    "github.com/astaxie/beego"
)

type MainController struct {
    beego.Controller
}

type HomeController struct {
    beego.Controller
}

func (this *MainController) Get() {
    this.Ctx.WriteString("hello world")
}

func (this *HomeController) Get() {
    this.Ctx.WriteString("Welcome to Go Beego Home Page")
}

func main() {
    beego.Router("/", &MainController{})
    beego.Router("/home", &HomeController{})
    beego.Run()
}