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
	"github.com/hunterhug/marmot/miner"
	"strconv"
)



func init() {
	爬虫, _ = miner.New(nil)
	爬虫.SetUa(miner.RandomUa())
	爬虫.SetHeaderParm("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8")
	//爬虫.SetHeaderParm("Accept-Encoding", "gzip, deflate, br")
	爬虫.SetHeaderParm("Accept-Language", "en-US,en;q=0.5")
}

// 详情页主图
func DownloadPdfMain() {
	for i:=1;i<=16;i++ {

		fmt.Println(i)
		url := "http://www.monitor.com.cn/report.aspx?id=433&page="+strconv.Itoa(i)
		fmt.Println(url)
		downlodpdf(TripAll(url))
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
	urlhost := strings.Split(url, "//")
	if len(urlhost) != 2 {
		fmt.Println("网站错误：开头必须为http://或https://")
		return
	}
	content, err := 爬虫.Get()

	//fmt.Println(url,string(content))
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
				fmt.Println(pdf,pdfname)
				if e1 && pdfname != "" && pdf !="" {
					if !strings.Contains(pdf, ".pdf") {
						return
					}
					filename := pdfname
					if util.FileExist(dir + "/" + filename + ".pdf") {
						fmt.Println("文件存在：" + dir + "/" + filename)
					} else {
						fmt.Println("下载:" + pdf)
						爬虫.SetUrl("http://www.monitor.com.cn"+pdf)
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
