package main

import (
	"html/template"
	"io"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/yizenghui/WechatQrcodeServe/orm"
	"github.com/yizenghui/WechatQrcodeServe/repository"
)

// 接入微信接口服务
func echoWxCallbackHandler(c echo.Context) error {
	repository.WechatServe(c.Response().Writer, c.Request())
	var err error
	return err
}

// Template ..
type Template struct {
	templates *template.Template
}

// Render ..
func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

//Home 主页
func Home(c echo.Context) error {
	return c.Render(http.StatusOK, "home", "")
}

//CreateTask 创建二维码任务
func CreateTask(c echo.Context) error {

	task, err := repository.NewQrcodeTask()
	if err != nil {

	}
	return c.JSON(http.StatusOK, task)
}

//CheckTask 创建二维码任务
func CheckTask(c echo.Context) error {
	token := c.Param("token")

	task, err := repository.CheckQrcodeTask(token)
	if err != nil {

	}
	return c.JSON(http.StatusOK, task)
}

func main() {
	orm.DB().AutoMigrate(&orm.Task{})

	t := &Template{
		templates: template.Must(template.ParseGlob("views/*.html")),
	}

	e := echo.New()
	e.Renderer = t
	// e.Pre(middleware.HTTPSRedirect())
	// e.Pre(middleware.HTTPSNonWWWRedirect())
	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/new", CreateTask)
	e.GET("/check/:token", CheckTask)

	// Route => handler
	e.GET("/", Home)

	e.File("/favicon.ico", "images/favicon.ico")

	e.Any("/wx_callback", echoWxCallbackHandler)

	e.Logger.Fatal(e.Start(":8061"))

}
