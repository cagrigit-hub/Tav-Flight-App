package handler

import (
	"fmt"
	"io"
	"os"

	"github.com/cagrigit-hub/tav-app/model"
	"github.com/cagrigit-hub/tav-app/view/home"
	"github.com/labstack/echo/v4"
	"github.com/xuri/excelize/v2"
)

type HomeHandler struct{}

func (h *HomeHandler) HandleHomeShow(c echo.Context) error {
	ho := model.Home{}
	return render(c, home.Show(ho))
}

func (h *HomeHandler) HandleExcelPost(c echo.Context) error {
	c.Request().ParseForm()

	file, err := c.FormFile("file")
	if err != nil {
		return err
	}
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()
	dst, err := os.Create(file.Filename)
	if err != nil {
		return err
	}
	defer dst.Close()
	if _, err = io.Copy(dst, src); err != nil {
		return err
	}

	excel, err := excelize.OpenFile(file.Filename)
	if err != nil {
		return err
	}
	defer func() {
		// Close the spreadsheet.
		if err := excel.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	rows, err := excel.GetRows("INITIAL SOLUTION")
	if err != nil {
		return err
	}
	for _, row := range rows {
		for _, colCell := range row {
			fmt.Print(colCell, "\t")
		}
		fmt.Println()
	}

	return c.String(400, "Bad Request")
}
