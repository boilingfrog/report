package main

import "github.com/report"

func main() {
	doc := report.NewDoc()
	err := doc.InitDoc("report1.doc")
	if err != nil {
		panic(err)
	}
	defer doc.CloseReport()
	doc.WriteHead()
	image1 := report.NewImage("1.gif", "../images/titlepic.gif", 50.00, 50.00, "")
	doc.WriteImage(true, "六壬网安'风险评估'安全评估报告", image1)
	doc.WriteTitle(report.NewText("测试文档"))

	doc.WriteTitle3(report.NewText("                               ———Web应用扫描"))
	tableHead := [][]interface{}{
		{report.NewText("URL")},
		{report.NewText("漏洞数量")},
		{report.NewText("URL")},
		{report.NewText("漏洞数量")},
	}
	table := [][]*report.TableTD{
		{
			report.NewTableTD([]interface{}{report.NewText("http://www.xjbtw.com/Article_Class.asp")}),
			report.NewTableTD([]interface{}{report.NewText("13")}),
			report.NewTableTD([]interface{}{report.NewText("13")}),
			report.NewTableTD([]interface{}{report.NewText("13")}),
		},
		{
			report.NewTableTD([]interface{}{report.NewText("http://www.xjbtw.com/Article_ClassC.asp")}),
			report.NewTableTD([]interface{}{report.NewText("8")}),
		},
		{
			report.NewTableTD([]interface{}{report.NewText("http://www.xjbtw.com/video_xwlb.asp")}),
			report.NewTableTD([]interface{}{report.NewText("8")}),
		},
		{
			report.NewTableTD([]interface{}{report.NewText("http://www.xjbtw.com/video_ttbb.asp")}),
			report.NewTableTD([]interface{}{report.NewText("7")}),
		},
		{
			report.NewTableTD([]interface{}{report.NewText("http://www.xjbtw.com/video_ttbb.asp")}),
		},
		{
			report.NewTableTD([]interface{}{report.NewText("http://www.xjbtw.com/video_ttbb.asp")}),
		},
	}
	trSpan := [][]int{
		{0, 0, 0, 0},
		{0, 3},
		{0, 3},
		{0, 3},
		{4},
		{4},
	}
	tdw := []int{1687, 1693, 6687, 1693}
	thw := []int{1687, 1693, 6687, 1693}
	tdh := []int{2, 4, 4, 4, 2, 2}

	tableObj := report.NewTable("", true, table, tableHead, thw, trSpan, tdw, tdh)
	doc.WriteTable(tableObj)

}
