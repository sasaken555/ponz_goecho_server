package main

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/sasaken555/ponz_goecho_server/routes"
)

func main() {
	e := echo.New()

	// Root Level Middleware
	// ログの設定は Apache Common Log Format っぽくすると読みやすい
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "${host} [${time_rfc3339_nano}] \"${method} ${uri}\" ${status} ${bytes_in} ${bytes_out}\n",
	}))
	e.Use(middleware.Recover())

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.GET("/users/:id", routes.GetUser)
	e.GET("/users/json", routes.GetJSONUser)
	e.GET("/teams", routes.GetTeam)
	e.POST("/save", routes.Save)
	e.POST("/users", routes.SaveUser)

	// Group Level Middleware
	g := e.Group("/admin")
	// Basic認証の設定
	g.Use(middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
		if username == "joe" && password == "secret" {
			return true, nil
		}
		return false, nil
	}))

	// Route Level Middleware you implemented...
	track := func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			println("request to /users")
			return next(c)
		}
	}
	e.GET("/users", func(c echo.Context) error {
		return c.String(http.StatusOK, "/users")
	}, track)

	e.Logger.Fatal(e.Start(":1323"))

}
