package utils

import (
	"bizd/metion/global"
	"bizd/metion/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/xuri/excelize/v2"
	"net/url"
	"strconv"
)

func WriteToExcel(c *gin.Context, data []model.PointPosition) {
	tmpSheetName := global.SheetName
	file := excelize.NewFile()
	sheetIndex, _ := file.NewSheet(tmpSheetName)
	file.SetActiveSheet(sheetIndex) // 默认sheet
	_ = file.DeleteSheet("Sheet1")  // 删除默认创建的sheet页
	rowsCount := len(data)

	for i := -1; i < rowsCount; i++ {
		var err error
		if i == -1 {
			err = file.SetSheetRow(tmpSheetName, fmt.Sprintf("A%d", i+2), &[]interface{}{
				"客户名称", "单位名称", "地址", "人数", "实施人员", "设备别名", "状态", "负责人", "资料链接", "预计实施时间", "实施开始时间", "实施结束时间", "备注"})
		} else {
			var tmp string
			if data[i].PeopleNumbers == nil {
				tmp = ""
			} else {
				tmp = strconv.Itoa(*data[i].PeopleNumbers)
			}
			err = file.SetSheetRow(tmpSheetName, fmt.Sprintf("A%d", i+2), &[]interface{}{data[i].ClientAbbreviation, data[i].PointPositionName, data[i].Address,
				tmp, data[i].ImplementerName, data[i].CpeName, data[i].StatusName, data[i].UserName, data[i].DataLink,
				data[i].ScheduledTime, data[i].StartTime, data[i].EndTime, data[i].Remark})
		}
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
