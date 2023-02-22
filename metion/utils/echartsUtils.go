package utils

import (
	"bizd/metion/model"
)

func InitAxis() model.EchartsLine {
	var line model.EchartsLine
	var xAxis = model.Axis{
		Type: "category",
		Name: "日期",
		NameTextStyle: model.NameTextStyle{
			FontWeight: 600,
			FontSize:   18,
		},
	}
	var days []string
	for i := 3; i >= 0; i-- {
		if i == 0 {
			days = append(days, GetCurWeekStartAndEnd())
		} else {
			days = append(days, GetPrevWeeksStartAndEnd(i))
		}
	}
	xAxis.Data = days
	var yAxis = model.Axis{
		Type: "value",
		Name: "平均实施时间/min",
		NameTextStyle: model.NameTextStyle{
			FontWeight: 600,
			FontSize:   18,
		},
	}
	line.XAxis = xAxis
	line.YAxis = yAxis
	line.Tooltip.Trigger = "axis"
	line.Title = model.EchartsTitle{
		Text: "实施效率",
		Left: "center",
	}
	line.Legend.Left = "right"
	return line
}

func GetNullLine() model.NullEcharts {
	var line model.NullEcharts
	line.Title = model.EchartsTitle{
		Text: "暂无数据",
		X:    "center",
		Y:    "center",
	}
	return line
}

func InitPie() model.EchartsPie {
	var pie model.EchartsPie
	pie.Legend = model.EchartsLegend{
		Orient: "vertical",
		Left:   "left",
	}
	pie.Title = model.EchartsTitle{
		Text: "实施数据",
		Left: "center",
	}
	pie.ToolTip = model.EchartsToolTip{
		Trigger: "item",
	}
	pie.Series = []model.EchartsPieSeries{
		{
			Type: "pie",
			Label: model.EchartsLabel{
				Show:      true,
				Formatter: "{b}\n{c}个\n{d}%",
			},
		},
	}
	return pie
}
