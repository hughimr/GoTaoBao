/*
   Created by jinhan on 17-8-24.
   Tip:
   Update:
*/
package src

import (
	"encoding/json"
	"fmt"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/hunterhug/marmot/miner"
	"github.com/hunterhug/parrot/util"
	"github.com/hunterhug/parrot/util/open"
	"github.com/hunterhug/parrot/util/xid"
	"os"
	"strconv"
	"time"
	"github.com/PuerkitoBio/goquery"
	"io"
)

// 每页11列 44个商品 // 不用 ajax方式
type SearchQuery struct {
	KsTS        string `json:"-"`           // 1503560947454_856
	Ajax        bool   `json:"ajax"`        // 不要改  true
	Bcoffset    int    `json:"bcoffset"`    // 不要改 4
	Callback    string `json:"callback"`    // jsonp857
	DataKey     string `json:"data-key"`    // s
	Ntoffset    int    `json:"ntoffset"`    // 0
	P4ppushleft string `json:"p4ppushleft"` // 1,48

	// 重要
	Page      int    `json:"-"`          // 第5页
	DataValue string `json:"data-value"` // 2156 ---> （50-1）*44=2156
	KeyWord   string `json:"q"`          // 搜索关键字
}

var (
	爬虫         *miner.Worker
	搜索链接       = "https://s.taobao.com/search?q=%s&s=%d&sort=%s&cd=false"
	有价格区间的搜索连接 = "https://s.taobao.com/search?q=%s&s=%d&sort=%s&filter=reserve_%s%.2f%s%.2f%s&cd=false"
	搜索排序       = map[int]string{
		1: "综合排序(MayBe千人千面)",
		2: "人气从高到低",
		3: "销量从高到低",
		4: "信用从高到低",
		5: "价格 低-高",
		6: "价格 高-低",
		7: "总价 低-高",
		8: "总价 高-低",
	}
	OrderMap = map[int]string{
		1: "default",     // 综合排序
		2: "renqi-desc",  // 人气从高到低
		3: "sale-desc",   // 销量从高到低
		4: "credit-desc", // 信用从高到低
		5: "price-asc",   // 价格 低-高
		6: "price-desc",  // 价格 高-低
		7: "total-asc",   //总价 低-高
		8: "total-desc",  //总价 高-低
	}
)

func init() {
	爬虫, _ = miner.New(nil)
	爬虫.SetUa(miner.RandomUa())
	爬虫.SetHeaderParm("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8")
	//爬虫.SetHeaderParm("Accept-Encoding", "gzip, deflate, br")
	爬虫.SetHeaderParm("Accept-Language", "en-US,en;q=0.5")
}

func 请问搜索如何排序() int {
	fmt.Println("我想问你,想如何排序：")
	fmt.Println("----------------")
	for k := 1; k <= len(OrderMap); k++ {
		fmt.Printf("%-20s 请选择:%d\n", 搜索排序[k], k)
	}
	fmt.Println("----------------")
	choice := util.Input("请选择：", "1")
	fmt.Println("选择完毕:" + choice)
	if i, e := util.SI(choice); e != nil {
		fmt.Println("请认真选择！")
		return 请问搜索如何排序()
	} else {
		return i
	}
}

// 搜索全部类型商品
func SearchPrepare(keyword string, page int, order int) string {
	orderstring, ok := OrderMap[order]
	if !ok {
		orderstring = "default"
		fmt.Println("排序条件出错，采用默认")
	}
	url := fmt.Sprintf(搜索链接, util.UrlE(keyword), (page-1)*44, orderstring)
	return url
}

//搜索带价格区间带全部类型商品
func SearchPrepareWithSection(keyword string, page int, order int, sectStart float64, sectEnd float64) string {
	orderstring, ok := OrderMap[order]
	if !ok {
		orderstring = "default"
		fmt.Println("排序条件出错，采用默认")
	}
	url := fmt.Sprintf(有价格区间的搜索连接, util.UrlE(keyword), (page-1)*44, orderstring, util.UrlE("price["), sectStart, util.UrlE(","), sectEnd, util.UrlE("]"))
	return url
}

// 只搜索天猫
func SearchPrepareTmall(keyword string, page int, order int) string {
	url := SearchPrepare(keyword, page, order)
	return url + "&filter_tianmao=tmall&tab=mall"
}

func Search(url string) ([]byte, error) {
	爬虫.SetUrl(url)
	return 爬虫.Get()
}

type Mods struct {
	ModData Items `json:"mods"`
	//PageName string `json:"pageName"`
}
type Items struct {
	Items   ItemList `json:"itemlist"`
	Sortbar BarList  `json:"sortbar"`
}

type BarList struct {
	Data DataList `json:"data"`
}

type DataList struct {
	Price PriceData `json:"price"`
	Pager PriceObject `json:"pager"`
}

type PriceData struct {
	Rank []RankObject `json:"rank"`

}

//解析总页面
type PriceObject struct {
	PageSize int `json:"pageSize"`
	TotalPage int `json:"totalPage"`
}

//解析页面上最多价格区间喜欢的
type RankObject struct {
	Percent int    `json:"percent"`
	Start   string `json:"start"`
	End     string `json:"end"`
}

type ItemList struct {
	Data ItemData `json:"data"`
}

// core
type ItemData struct {
	Auctions []ItemObject `json:"auctions"`
}

// 我的商品(不区分广告，某些商品做了广告会被置顶！)
type ItemObject struct {
	IsTmallObject IsTmall `json:"shopcard"`      // 是否天猫
	CommentCount  string  `json:"comment_count"` // 评论数
	Nid           string  `json:"nid"`           // 商品ID
	//CommentUrl   string `json:"comment_url"`
	//DetailUrl    string `json:"detail_url"`
	ItemLoc  string `json:"item_loc"`  // 发货地
	Nick     string `json:"nick"`      //  店铺名字
	PicUrl   string `json:"pic_url"`   // 商品图片
	RawTitle string `json:"raw_title"` // 商品标题
	//ShopLink     string `json:"shopLink"`   // 店铺URL
	UserId    string `json:"user_id"`    // 卖家ID
	ViewFee   string `json:"view_fee"`   // 小费？
	ViewPrice string `json:"view_price"` // 价格
	ViewSales string `json:"view_sales"` // 付款人数

	SectPercent int     `json:"sect_percent"` //价格区间占比
	SectStart   float64 `json:"sect_start"`   //区间起始值
	SectEnd     float64 `json:"sect_end"`     //区间结束值
}

type IsTmall struct {
	Yes bool `json:"isTmall"`
}

type KeyItem struct {
	LevelOne   string `json:"level_one"`
	LevelTwo   string `json:"level_two"`
	LevelThree string `json:"level_three,omitempty"`
}

func ParseSearchPrepare(data []byte) []byte {
	parsereg := "g_page_config = ({.*})"
	r, err := regexp.Compile(parsereg)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		bb := r.FindAllSubmatch(data, -1)
		if len(bb) > 0 {
			return bb[0][1]
		}
	}
	return []byte("")
}

// 解析到结构体
func ParseSearch(data []byte) Mods {
	items := Mods{}
	err := json.Unmarshal(data, &items)
	if err != nil {
		fmt.Println(err.Error())
	}
	return items
}

func SearchMain() {
	for {
		csv := []ItemObject{}

		fmt.Println(`
	-------------------------------
	欢迎使用强大的搜索框小工具
	你只需安装提示进行即可！
	联系QQ：459527502
	----------------------------------
	`)
		keyword := util.Input("请输入关键字(请使用+代替空格！):", "")
		keyword = strings.Replace(keyword, " ", "+", -1)
		types := 请问搜索如何排序()
		tmall := false
		if strings.Contains(strings.ToLower(util.Input("是否只搜索天猫商品(Y/y),默认N", "n")), "y") {
			tmall = true
		}

		pagestemp := util.Input("你要抓几页(1-100):", "1")
		pages, err := util.SI(pagestemp)
		if err != nil {
			fmt.Println("输入页数有问题")
			break
		}
		if pages > 100 || pages < 1 {
			fmt.Printf("你选择的页数有问题：%d\n", pages)
			break
		}
		url := ""
		for page := 1; page <= pages; page++ {
			if tmall {
				url = SearchPrepareTmall(keyword, page, types)
			} else {
				url = SearchPrepare(keyword, page, types)
			}
			fmt.Println("搜索:" + url)
			data, err := Search(url)
			if err != nil {
				fmt.Printf("抓取第%d页 失败：%s\n", page, err.Error())
			} else {
				fmt.Printf("抓取第%d页\n", page)
				/*filename := filepath.Join(".", "原始数据", util.ValidFileName(keyword), "search"+util.IS(page)+".html")
				util.MakeDirByFile(filename)
				e := util.SaveToFile(filename, data)
				if e != nil {
					fmt.Printf("保存数据在:%s 失败:%s\n", filename, e.Error())
					continue
				}
				fmt.Printf("保存数据在:%s 成功\n", filename)*/
				xx := ParseSearchPrepare(data)
				if string(xx) == "" {
					fmt.Println("这页数据为空...")
					continue
				}
				a := ParseSearch(xx)
				if len(a.ModData.Items.Data.Auctions) > 0 {
					for _, v := range a.ModData.Items.Data.Auctions {
						csv = append(csv, v)
						//fmt.Printf("%#v\n", v)
					}
				}
			}
		}

		if len(csv) == 0 {
			fmt.Println("啥都没抓到")
			continue
		}
		/**************************/
		id := xid.New().String()
		fileonly := util.TodayString(5) + "-" + id
		rootdir := filepath.Join(".", "搜索结果", util.ValidFileName(keyword))
		util.MakeDir(rootdir)
		tempdata := "排序,商品标题,店铺名,发货地址,评论数,是否天猫,小费,价格,销量,用户ID,店铺URL,商品ID,商品详情URL,商品评论URL图片地址\n"

		for k, v := range csv {
			tempdata = tempdata + fmt.Sprintf("%v,%s,%s,%s,", k+1, CD(v.RawTitle), v.Nick, v.ItemLoc)
			tempdata = tempdata + fmt.Sprintf("%s,%v,%s,%s,%s,", v.CommentCount, v.IsTmallObject.Yes, v.ViewFee, v.ViewPrice, v.ViewSales)
			s1 := "http://store.taobao.com/shop/view_shop.htm?user_number_id=" + v.UserId
			s2 := "http://detail.tmall.com/item.htm?id=" + v.Nid
			s3 := s2 + "&on_comment=1"
			tempdata = tempdata + fmt.Sprintf("%s,%s,%s,%s,%s,%s\n", v.UserId, s1, v.Nid, s2, s3, "http:"+v.PicUrl)
		}

		filekeep := rootdir + "/" + fileonly + ".csv"
		util.SaveToFile(filekeep, []byte(tempdata))
		fmt.Println("保存成功，请打开:" + filekeep)
		if strings.Contains(strings.ToLower(util.Input("是否打开文件(Y/y)", "n")), "y") {
			open.Start(filekeep)
		}
		/*************************/
		if cancle() == "y" {
			break
		}
	}

}

func CD(a string) string {
	return TripAll(strings.Replace(a, ",", "*", -1))
}

//copy from SearchMain()
func MySearchMain(keyWord string) {

	csv := []ItemObject{}

	//	fmt.Println(`
	//-------------------------------
	//欢迎使用强大的搜索框小工具
	//你只需安装提示进行即可！
	//联系QQ：459527502
	//----------------------------------
	//`)
	//	keyword := util.Input("请输入关键字(请使用+代替空格！):", "")
	keyword := strings.Replace(keyWord, " ", "+", -1)
	//types := 请问搜索如何排序() 直接按照销量降序排序
	types := 3
	//默认搜索的商品包含天猫和淘宝
	//tmall := false
	//if strings.Contains(strings.ToLower(util.Input("是否只搜索天猫商品(Y/y),默认N", "n")), "y") {
	//	tmall = true
	//}

	//pagestemp := util.Input("你要抓几页(1-100):", "1")
	//抓取100页
	pages := 100
	//pages, err := util.SI(pagestemp)
	//if err != nil {
	//	fmt.Println("输入页数有问题")
	//	break
	//}
	//if pages > 100 || pages < 1 {
	//	fmt.Printf("你选择的页数有问题：%d\n", pages)
	//	break
	//}
	url := ""
	//解析得到价格区间
	url0 := SearchPrepare(keyword, 1, types)
	data0, err0 := Search(url0)
	if err0 != nil {
		fmt.Printf("抓取第%d页 失败：%s\n", 1, err0.Error())
	} else {
		x0 := ParseSearchPrepare(data0)
		if string(x0) == "" {
			fmt.Println("抓取起始页数据为空")
			os.Exit(1)
		}
		a0 := ParseSearch(x0)
		rankList := a0.ModData.Sortbar.Data.Price.Rank
		if len(rankList) > 0 {
			for _, v0 := range rankList {
				percent0 := v0.Percent
				start0, _ := strconv.ParseFloat(v0.Start, 64)
				end0, _ := strconv.ParseFloat(v0.End, 64)

				//解析分类第一个页面拿到总页面数
				urlP:=SearchPrepareWithSection(keyword, 1, types, start0, end0)
				dataP,errP:=Search(urlP)
				if errP!=nil{
					fmt.Println("获取总页面数失败")
				}else{
					xp:=ParseSearchPrepare(dataP)
					if string(xp)==""{
						fmt.Println("这页数据为空。")
						continue
					}else{
						ap:=ParseSearch(xp)
						pages=ap.ModData.Sortbar.Data.Pager.TotalPage
						if string(pages)==""{
							fmt.Println("总页面数没解析到")
							continue
						}else{
							fmt.Printf("总页面数是%d",pages)
						}
					}
				}




				for page := 1; page <= pages; page++ {
					url = SearchPrepareWithSection(keyword, page, types, start0, end0)

					//fmt.Println("搜索:" + url)
					data, err := Search(url)
					if err != nil {
						fmt.Printf("抓取区间[%.2f,%.2f]第%d页 失败：%s\n", start0, end0, page, err.Error())
					} else {
						fmt.Printf("抓取区间[%.2f,%.2f]第%d页\n", start0, end0, page)

						xx := ParseSearchPrepare(data)
						if string(xx) == "" {
							fmt.Println("这页数据为空...")
							continue
						}
						a := ParseSearch(xx)
						if len(a.ModData.Items.Data.Auctions) > 0 {
							for _, v := range a.ModData.Items.Data.Auctions {
								v.SectPercent = percent0
								v.SectStart = start0
								v.SectEnd = end0
								csv = append(csv, v)
								//fmt.Printf("%#v\n", v)
							}
						}
					}
				}
			}
		} else {
			url := ""
			for page := 1; page <= pages; page++ {

				url = SearchPrepare(keyword, page, types)

				//fmt.Println("搜索:" + url)
				data, err := Search(url)
				if err != nil {
					fmt.Printf("抓取第%d页 失败：%s\n", page, err.Error())
				} else {

					xx := ParseSearchPrepare(data)
					if string(xx) == "" {
						fmt.Println("这页数据为空...")
						continue
					}
					a := ParseSearch(xx)
					if len(a.ModData.Items.Data.Auctions) > 0 {
						for _, v := range a.ModData.Items.Data.Auctions {
							v.SectPercent = 0
							v.SectStart = 0.0
							v.SectEnd = 0.0
							csv = append(csv, v)
							//fmt.Printf("%#v\n", v)
						}
					}
				}
			}
		}

		if len(csv) == 0 {
			fmt.Println("啥都没抓到")
			//os.Exit(1)
		}
		/**************************/
		id := xid.New().String()
		fileonly := util.TodayString(5) + "-" + id
		nowDay := time.Now().Format("2006-01-02")
		rootdir := filepath.Join(".", "搜索结果", time.Now().Format("2006/01/02"))
		util.MakeDir(rootdir)
		tempdata := "排序,日期，关键字，所在区间占比，区间起始值，区间结束值，商品标题,店铺名,发货地址,评论数,是否天猫,小费,价格,销量,用户ID,店铺URL,商品ID,商品详情URL,商品评论URL图片地址\n"

		for k, v := range csv {
			tempdata = tempdata + fmt.Sprintf("%v,%s,%s,%d,%.2f,%.2f,%s,%s,%s,", k+1, nowDay, keyword, v.SectPercent, v.SectStart, v.SectEnd, CD(v.RawTitle), v.Nick, v.ItemLoc)
			tempdata = tempdata + fmt.Sprintf("%s,%v,%s,%s,%s,", v.CommentCount, v.IsTmallObject.Yes, v.ViewFee, v.ViewPrice, v.ViewSales)
			s1 := "http://store.taobao.com/shop/view_shop.htm?user_number_id=" + v.UserId
			s2 := "http://detail.tmall.com/item.htm?id=" + v.Nid
			s3 := s2 + "&on_comment=1"
			tempdata = tempdata + fmt.Sprintf("%s,%s,%s,%s,%s,%s\n", v.UserId, s1, v.Nid, s2, s3, "http:"+v.PicUrl)
		}

		filekeep := rootdir + "/" + fileonly + ".csv"
		//util.SaveToFile(filekeep, []byte(tempdata))
		WriteStringToFile(filekeep,string(tempdata))
	}
}

func WriteStringToFile(filepath, s string) error {
	fo, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer fo.Close()

	_, err = io.Copy(fo, strings.NewReader(s))
	if err != nil {
		return err
	}

	return nil
}

func GetKeywords() [][]byte {

	keys := [][]byte{}
	key := make(map[string]interface{})

	urlKey := "https://www.taobao.com/markets/tbhome/market-list"

	dataK, errK := Search(urlKey)

	//解析页面拿到关键字
	if errK != nil || string(dataK) == "" {
		fmt.Printf("获取关键字失败.error:%s", errK.Error())
		os.Exit(1)
	}

	doc, _ := goquery.NewDocumentFromReader(strings.NewReader(string(dataK)))

	doc.Find("div.layout.layout-grid-0").Each(func(index int, sel *goquery.Selection) {
		div1 := sel.Find("div.grid-0").Find("div.col.col-main").Find("div.main-wrap.J_Region").Find("div.home-category-list.J_Module").Find("div.module-wrap")
		a1 := div1.Find("a.category-name.category-name-level1.J_category_hash").Text()
		div1.Find("ul.category-list").Find("li").Each(func(index2 int, sel2 *goquery.Selection) {
			a2 := sel2.ChildrenFiltered("a.category-name").Text()

			sel2.Find("div.category-items").Find("a").Each(func(index3 int, sel3 *goquery.Selection) {
				a3 := sel3.Text()

				key["level_one"] = a1
				key["level_two"] = a2
				key["level_three"] = a3
				keyB, _ := json.Marshal(key)
				keys = append(keys, keyB)

			})
		})
	})

	return keys

}
