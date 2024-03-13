package handler

import (
	"fmt"
	"io"
	"os"
	"strconv"

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
	fmt.Println("hits here")
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

	gateToCarouselExcel, err := excelize.OpenFile("gate_karosel.xlsx")
	if err != nil {
		return err
	}

	// gate -> karosel -> distance
	distances := make(map[int]map[int]int)

	colStart := 4
	rowStart, rowEnd := 2, 8
	gate := 103
	carousel := 21

	rows, err := gateToCarouselExcel.GetRows("Sheet1")
	if err != nil {
		return err
	}
	for j := colStart; j < 18; j++ {
		row := rows[j]
		if len(row) > rowStart {
			// get each row from idxStart to idxEnd (included)
			for i := rowStart; i <= rowEnd; i++ {
				if row[i] != "" {
					if distances[gate] == nil {
						distances[gate] = make(map[int]int)
					}
					distances[gate][carousel], err = strconv.Atoi(row[i])
					if err != nil {
						fmt.Println(err)
						return err
					}
					carousel++
				}
			}
			carousel = 21
			gate++
		}
	}
	fmt.Println(distances)
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

	rows, err = excel.GetRows("içhatlar")
	if err != nil {
		return err
	}
	standIndex := 4
	acTypeIndex := 5

	stands := []string{}
	acTypes := []string{}

	for _, row := range rows {
		if len(row) > standIndex {
			stands = append(stands, row[standIndex])
		}
		if len(row) > acTypeIndex {
			acTypes = append(acTypes, row[acTypeIndex])
		}

	}
	fmt.Println(stands)
	fmt.Println(acTypes)

	return c.String(400, "Bad Request")
}
