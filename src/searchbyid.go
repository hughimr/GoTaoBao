/*
   Created by jinhan on 17-8-24.
   Tip:
   Update:
*/
package src

import (
	//"encoding/json"
	"fmt"
	//"path/filepath"
	//"regexp"
	//"strings"
	//
	//"github.com/hunterhug/marmot/miner"
	//"github.com/hunterhug/parrot/util"
	//"github.com/hunterhug/parrot/util/open"
	//"github.com/hunterhug/parrot/util/xid"
	//"os"
	"strconv"
	//"time"
	//"github.com/PuerkitoBio/goquery"
	//"io"
	//"bytes"
	//"math/big"
)

//copy from SearchMain()
func searchByID(id int) {
	//根据id拼接商品详情url

	url := "https://trade-acs.m.taobao.com/gw/mtop.taobao.detail.getdetail/6.0/?data=%7B%22detail_v%22%3A%223.1.0%22%2C%22exParams%22%3A%22%7B%5C%22action%5C%22%3A%5C%22ipv%5C%22%2C%5C%22countryCode%5C%22%3A%5C%22CN%5C%22%2C%5C%22cpuCore%5C%22%3A%5C%224%5C%22%2C%5C%22cpuMaxHz%5C%22%3A%5C%221209600%5C%22%2C%5C%22from%5C%22%3A%5C%22search%5C%22%2C%5C%22id%5C%22%3A%5C%22" + strconv.Itoa(id) + "%5C%22%2C%5C%22item_id%5C%22%3A%5C%22" + strconv.Itoa(id) + "%5C%22%2C%5C%22latitude%5C%22%3A%5C%2223.12577%5C%22%2C%5C%22list_type%5C%22%3A%5C%22search%5C%22%2C%5C%22longitude%5C%22%3A%5C%22113.372117%5C%22%2C%5C%22osVersion%5C%22%3A%5C%2219%5C%22%2C%5C%22phoneType%5C%22%3A%5C%22Che1-CL10%5C%22%2C%5C%22search_action%5C%22%3A%5C%22initiative%5C%22%2C%5C%22soVersion%5C%22%3A%5C%222.0%5C%22%2C%5C%22utdid%5C%22%3A%5C%22VQF1POae6O4DABLWrwO0STXN%5C%22%7D%22%2C%22itemNumId%22%3A%22" + strconv.Itoa(id) + "%22%7D"
	fmt.Println(url)
	//解析得到价格区间
	//url0 := SearchPrepare(keyword, 1, types)
	data0, err0 := Search(url)
	if err0 != nil {
		fmt.Printf("抓取详情页失败：%s\n", err0.Error())
	} else {
		fmt.Println(data0)
		//x0 := ParseSearchPrepare(data0)
		//if string(x0) == "" {
		//	fmt.Println("抓取起始页数据为空")
		//	os.Exit(1)
		//}
		//a0 := ParseSearch(x0)
		//b0 := ParseSearchKeys(x0)
		//rankList := a0.ModData.Sortbar.Data.Price.Rank
		//keyWordList := b0.MainInfoData.TraceData.TraceDataObject.RsKeywords
		//if len(rankList) > 0 {
		//	for _, v0 := range rankList {
		//		percent0 := v0.Percent
		//		start0, _ := strconv.ParseFloat(v0.Start, 64)
		//		end0, _ := strconv.ParseFloat(v0.End, 64)
		//
		//		//解析分类第一个页面拿到总页面数
		//		urlP := SearchPrepareWithSection(keyword, 1, types, start0, end0)
		//		dataP, errP := Search(urlP)
		//		if errP != nil {
		//			fmt.Println("获取总页面数失败")
		//		} else {
		//			xp := ParseSearchPrepare(dataP)
		//			if string(xp) == "" {
		//				fmt.Println("这页数据为空。")
		//				continue
		//			} else {
		//				ap := ParseSearch(xp)
		//				pages = ap.ModData.Sortbar.Data.Pager.TotalPage
		//				if string(pages) == "" {
		//					fmt.Println("总页面数没解析到")
		//					continue
		//				} else {
		//					fmt.Printf("总页面数是%d", pages)
		//				}
		//			}
		//		}
		//
		//		for page := 1; page <= pages; page++ {
		//			trytimes := 0
		//			url = SearchPrepareWithSection(keyword, page, types, start0, end0)
		//
		//			//fmt.Println("搜索:" + url)
		//		AGAIN:
		//			data, err := Search(url)
		//			trytimes++
		//			if err != nil {
		//				fmt.Printf("抓取区间[%.2f,%.2f]第%d页 失败：%s\n", start0, end0, page, err.Error())
		//			} else {
		//				fmt.Printf("抓取区间[%.2f,%.2f]第%d页\n", start0, end0, page)
		//
		//				xx := ParseSearchPrepare(data)
		//				if string(xx) == "" {
		//					fmt.Println("这页数据为空...")
		//					if trytimes < 3 {
		//						time.Sleep(5 * time.Second)
		//						fmt.Println("重试第", trytimes, "次")
		//						goto AGAIN
		//					} else {
		//						continue
		//					}
		//				}
		//				trytimes = 0
		//				a := ParseSearch(xx)
		//				if len(a.ModData.Items.Data.Auctions) > 0 {
		//					for _, v := range a.ModData.Items.Data.Auctions {
		//						v.SectPercent = percent0
		//						v.SectStart = start0
		//						v.SectEnd = end0
		//						csv = append(csv, v)
		//						//fmt.Printf("%#v\n", v)
		//					}
		//				}
		//			}
		//		}
		//	}
		//} else {
		//	url := ""
		//
		//	for page := 1; page <= pages; page++ {
		//
		//		trytimes := 0
		//		url = SearchPrepare(keyword, page, types)
		//
		//		//fmt.Println("搜索:" + url)
		//	AGAIN2:
		//		data, err := Search(url)
		//		trytimes++
		//		if err != nil {
		//			fmt.Printf("抓取第%d页 失败：%s\n", page, err.Error())
		//		} else {
		//
		//			xx := ParseSearchPrepare(data)
		//			if string(xx) == "" {
		//				fmt.Println("这页数据为空...")
		//				if trytimes < 3 {
		//					time.Sleep(5 * time.Second)
		//					fmt.Println("重试第", trytimes, "次")
		//					goto AGAIN2
		//				} else {
		//					continue
		//				}
		//			}
		//			trytimes = 0
		//			a := ParseSearch(xx)
		//			if len(a.ModData.Items.Data.Auctions) > 0 {
		//				for _, v := range a.ModData.Items.Data.Auctions {
		//					v.SectPercent = 0
		//					v.SectStart = 0.0
		//					v.SectEnd = 0.0
		//					csv = append(csv, v)
		//					//fmt.Printf("%#v\n", v)
		//				}
		//			}
		//		}
		//	}
		//}
		//
		//if len(csv) == 0 {
		//	fmt.Println("啥都没抓到")
		//	//os.Exit(1)
		//}
		///**************************/
		//id := xid.New().String()
		//fileonly := util.TodayString(5) + "-" + id
		//nowDay := time.Now().Format("2006-01-02")
		//rootdir := filepath.Join(".", "搜索结果", time.Now().Format("2006/01/02"))
		//util.MakeDir(rootdir)
		//var buffer bytes.Buffer
		//buffer.Reset()
		//buffer.WriteString("排序,日期,关键字,关联关键字,所在区间占比,区间起始值,区间结束值,商品标题,店铺名,发货地址,评论数,是否天猫,小费,价格,销量,用户ID,店铺URL,商品ID,商品详情URL,商品评论URL图片地址\n")
		//
		//for k, v := range csv {
		//	buffer.WriteString(fmt.Sprintf("%v,%s,%s,%s,%d,%.2f,%.2f,%s,%s,%s,", k+1, nowDay, keyword, strings.Join(keyWordList, ";"), v.SectPercent, v.SectStart, v.SectEnd, CD(v.RawTitle), v.Nick, v.ItemLoc))
		//	buffer.WriteString(fmt.Sprintf("%s,%v,%s,%s,%s,", v.CommentCount, v.IsTmallObject.Yes, v.ViewFee, v.ViewPrice, v.ViewSales))
		//	s1 := "http://store.taobao.com/shop/view_shop.htm?user_number_id=" + v.UserId
		//	s2 := "http://detail.tmall.com/item.htm?id=" + v.Nid
		//	s3 := s2 + "&on_comment=1"
		//	buffer.WriteString(fmt.Sprintf("%s,%s,%s,%s,%s,%s\n", v.UserId, s1, v.Nid, s2, s3, "http:"+v.PicUrl))
		//}
		//
		//filekeep := rootdir + "/" + fileonly + ".csv"
		////util.SaveToFile(filekeep, []byte(tempdata))
		//filename, _ := os.Create(filekeep)
		//buffer.WriteTo(filename)
		//filename.Close()
		////WriteStringToFile(filekeep,buffer.String())
	}
}
