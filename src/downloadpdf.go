/*
   Created by jinhan on 17-8-25.
   Tip:
   Update:
*/
package src

import (
	"fmt"
	"path/filepath"
	//"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/hunterhug/marmot/expert"
	"github.com/hunterhug/parrot/util"
)

// 详情页主图
func DownloadPdfMain() {
	for {
		fmt.Println(`
	-------------------------------
	欢迎使用强大的图片下载小工具
	你只需按照提示进行即可！
	联系QQ：459527502
	----------------------------------
	`)
		fmt.Println("请输入天猫淘宝链接*保存目录")
		fmt.Println("如：https://item.taobao.com/item.htm?id=40066362090*taobao")
		fmt.Println("------------以上详情页会保存在“图片/taobao”文件夹下--------------")
		url := util.Input("请输入：", "")
		downlodpdf(TripAll(url))
		if cancle() == "y" {
			break
		}
	}
}

func downlodpdf(url string) {


	//filename := util.TodayString(3)
	//filename := "默认保存"
	//if len(temp) >= 2 && temp[1] != "" {
	//	filename = util.ValidFileName(temp[1])
	//}
	//dir := filepath.Join(".", "图片", filename)
	dir:=filepath.Join(".","pdf","中怡康智库")
	util.MakeDir(dir)
	爬虫.SetUrl(url)
	//urlhost := strings.Split(url, "//")
	//if len(urlhost) != 2 {
	//	fmt.Println("网站错误：开头必须为http://或https://")
	//	return
	//}
	content, err := 爬虫.Get()
	if err != nil {
		fmt.Println(err.Error())
		return
	} else {
		docm, err := expert.QueryBytes(content)
		if err != nil {
			fmt.Println(err.Error())
			return
		} else {
			docm.Find("div.zhuanye_bao").Find("li").Each(func(num int, node *goquery.Selection) {
				pdf, e1 := node.Find("a").Eq(0).Attr("href")
				pdfname :=node.Find("a").Eq(0).Text()
				//if e == false {
				//	img, e = node.Attr("data-src")
				//}
				if e1 && pdfname != "" && pdf !="" {
					if !strings.Contains(pdf, ".pdf") {
						return
					}
					filename := pdfname
					if util.FileExist(dir + "/" + filename + ".pdf") {
						fmt.Println("文件存在：" + dir + "/" + filename)
					} else {
						fmt.Println("下载:" + pdf)
						爬虫.SetUrl(pdf)
						//if strings.HasPrefix(pdf, "//") {
						//	爬虫.SetUrl("http:" + )
						//}
						pdfpage, e := 爬虫.Get()
						if e != nil {
							fmt.Println("下载出错" + pdf + ":" + e.Error())
							return
						}
						e = util.SaveToFile(dir+"/"+filename+".pdf", pdfpage)
						if e == nil {
							fmt.Println("成功保存在" + dir + "/" + filename)
						}
						//util.Sleep(1)
						//fmt.Println("暂停1秒")
					}
				}
			})

		}

	}
}
