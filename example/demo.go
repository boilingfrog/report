package main

import (
	"fmt"
	"io"

	"github.com/report"
)

func main() {
	doc := report.NewDoc()

	err := doc.NewBuffer()
	if err != nil {
		panic(err)
	}

	if err := doc.WriteHead(); err != nil {
		fmt.Println(err)
	}

	if err := doc.WriteTitle(report.NewText("测试文档")); err != nil {
		fmt.Println(err)
	}

	if err := doc.WriteTitle3(report.NewText("                                  ———Web应用扫描")); err != nil {
		fmt.Println(err)
	}
	tableHead := [][]interface{}{
		{report.NewText("部门或型号")},
		{report.NewText("部门:研发;型号:martin;")},
		{report.NewText("监督时间")},
		{report.NewText("2020-06-04")},
	}
	table := [][]*report.TableTD{
		{
			report.NewTableTD([]interface{}{report.NewText("监督内容")}),
			report.NewTableTD([]interface{}{report.NewText("你好吗")}),
		},
		{
			report.NewTableTD([]interface{}{report.NewText("主要问题描述")}),
			report.NewTableTD([]interface{}{report.NewText("哈哈哈，我不好")}),
		},
		{
			report.NewTableTD([]interface{}{report.NewText("监督意见或建议")}),
			report.NewTableTD([]interface{}{report.NewText("今天天气好吗")}),
		},
		{
			report.NewTableTD([]interface{}{report.NewText("监督人员：yuelei")}),
		},
		{
			report.NewTableTD([]interface{}{report.NewText("部门领导：")}),
		},
	}
	// 合并单元格操作
	trSpan := [][]int{
		{0, 3},
		{0, 3},
		{0, 3},
		{4},
		{4},
	}
	// 头表格宽度
	tdw := []int{1687, 2993, 5687, 1693}
	// 单元格宽度
	thw := []int{1687, 2993, 5687, 1693}
	// 单元格高度
	tdh := []int{3, 5, 5, 2, 2}

	tableObj := report.NewTable("", true, table, tableHead, thw, trSpan, tdw, tdh)
	if err := doc.WriteTable(tableObj); err != nil {
		fmt.Println(err)
	}
	// 这一行要加上，结束word
	if err := doc.WriteEndHead(); err != nil {
		fmt.Println(err)
	}

	fmt.Println(doc.Buffer.String())

	var i io.Writer

	doc.SaveTo(i)

	fmt.Println(i)

}
