/*
   Created by jinhan on 17-8-24.
   Tip:
   Update:
*/
package main

import (

	"github.com/hughimr/GoTaoBao/src"
	//"net/url"
	//"fmt"
	"fmt"
	//"os"
	//"github.com/PuerkitoBio/goquery"
	//"strings"
	"encoding/json"
)



func main() {
	//miner.SetLogLevel("debug")
	/*
		执行爬取天猫淘宝商品信息
	*/
	//src.MySearchMain("加湿器")


	for _,v  :=range src.GetKeywords(){

		var m src.KeyItem
		err:=json.Unmarshal(v,&m)
		if err==nil{fmt.Println(m.LevelOne)}
	}




	}
