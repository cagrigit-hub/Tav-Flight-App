
package handler

import (
	"github.com/cagrigit-hub/tav-app/model"
	"github.com/cagrigit-hub/tav-app/view/example"
	"github.com/labstack/echo/v4"
)

type ExampleHandler struct{}

func (h *ExampleHandler) HandleExampleShow(c echo.Context) error {
	u := model.Example{
		Text: "example-text",
	}
	return render(c, example.Show(u))
}

func (h *ExampleHandler) HandlePost(c echo.Context) error {
	c.Request().ParseForm()
	if c.Request().Form.Has("example") {
		u := model.Example{
			Text: c.Request().Form.Get("example"),
		}
		return render(c, example.EcOne(u))
	}
	return c.String(400, "Bad Request")
}

	