package main

import "github.com/report"

func main() {
	doc := report.NewDoc()
	err := doc.InitDoc("example/report.doc")
	if err != nil {
		panic(err)
	}
	defer doc.CloseReport()
	doc.WriteHead()
	doc.WriteTitle(report.NewText("测试文档"))

	doc.WriteTitle3(report.NewText("                               ———Web应用扫描"))
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
	doc.WriteTable(tableObj)
	// 这一行要加上，结束word
	doc.WriteEndHead(false, "text", "", "")

}
