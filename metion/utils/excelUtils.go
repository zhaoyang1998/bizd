package utils

import (
	"bizd/metion/global"
	"bizd/metion/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/xuri/excelize/v2"
	"net/url"
)

func WriteToExcel(c *gin.Context, data [][]string) {
	tmpSheetName := global.SheetName
	file := excelize.NewFile()
	sheetIndex, _ := file.NewSheet(tmpSheetName)
	style, _ := file.NewStyle(&excelize.Style{
		Alignment: &excelize.Alignment{
			Horizontal:      "center",
			Indent:          1,
			JustifyLastLine: true,
			ReadingOrder:    2,
			RelativeIndent:  1,
			ShrinkToFit:     true,
			Vertical:        "top",
			WrapText:        true,
		},
	})
	file.SetActiveSheet(sheetIndex) // 默认sheet
	_ = file.SetColWidth(global.SheetName, "N", "N", 150)
	_ = file.DeleteSheet("Sheet1") // 删除默认创建的sheet页
	_ = file.SetCellStyle(global.SheetName, "A1", "A1", style)
	for i, item := range data {
		var err error
		err = file.SetSheetRow(tmpSheetName, fmt.Sprintf("A%d", i+1), &item)
		if err != nil {
			panic(model.MyError{Code: 400, Message: err.Error()})
			return
		}
	}
	_ = file.Close()
	fileName := url.QueryEscape(global.ExcelName + ".xls")
	Write(c, fileName, file)
}

func Write(ctx *gin.Context, fileName string, file *excelize.File) {
	ctx.Writer.Header().Add("Access-Control-Expose-Headers", "Content-Disposition")
	ctx.Writer.Header().Add("Content-Disposition", fmt.Sprintf("attachment; filename*=utf-8''%s", fileName))
	ctx.Writer.Header().Add("Content-Type", "application/octet-stream;charset=UTF-8")
	ctx.Writer.Header().Add("Content-Transfer-Encoding", "binary")
	_ = file.Write(ctx.Writer)
}
