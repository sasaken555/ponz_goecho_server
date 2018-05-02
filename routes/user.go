package routes

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"github.com/sasaken555/ponz_goecho_server/util"
)

// GetUser ... ユーザーを返す
func GetUser(c echo.Context) error {
	// User ID from Path Parameter `users/:id`
	id := c.Param("id")
	return c.String(http.StatusOK, id)
}

// GetTeam ... QueryStringからチームを返す
func GetTeam(c echo.Context) error {
	// Get Team and Member from query string
	team := c.QueryParam("team")
	member := c.QueryParam("member")
	return c.String(http.StatusOK, "team: "+team+", member: "+member)
}

// Save ... フォームデータから値を取り出して返す
func Save(c echo.Context) error {
	// Get name and email
	name := c.FormValue("name")
	email := c.FormValue("email")
	return c.String(http.StatusOK, "name: "+name+", email: "+email)
}

// User ...ユーザー情報を格納する構造体
type User struct {
	Name  string `json:"name" xml:"name" form:"name" query:"name"`
	Email string `json:"email" xml:"email" form:"email" query:"email"`
}

// SaveUser ... 名前とEmailを持つデータを取得してJSONを返す
func SaveUser(c echo.Context) error {
	u := new(User)
	if err := c.Bind(u); err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, u)
}

// Customer ... 顧客情報
type Customer struct {
	ID        int64  `json:"id" xml:"id"`
	Name      string `json:"name" xml:"name"`
	OrderNum  int    `json:"ordernum" xml:"ordernum"`
	OrderProd string `json:"orderprod" xml:"orderprod"`
}

// GetJSONUser ... JSONを返す
func GetJSONUser(c echo.Context) error {
	userID, err := strconv.ParseInt(c.QueryParam("userId"), 10, 0)
	userName := c.QueryParam("userName")
	orderNum := util.GetRand(100)

	// userIDが整数でなければ500エラーを返す
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "You Should Provide userId as Integer!")
	}

	u := &Customer{
		ID:        userID,
		Name:      userName,
		OrderNum:  orderNum,
		OrderProd: "Blend Coffee",
	}

	return c.JSON(http.StatusOK, u)
}
