/*
   Created by jinhan on 17-8-24.
   Tip:
   Update:
*/
package main

import (
	"fmt"
	//"github.com/hunterhug/marmot/miner"
	"github.com/hunterhug/GoTaoBao/src"
	"github.com/hunterhug/parrot/util"
)

func main() {
	//miner.SetLogLevel("debug")
	fmt.Println(`
	---------------------------------------------
	|	亲爱的朋友，你好！
	|	欢迎使用SillyChen制作的小工具
	|	友好超乎你想象！
	|	如果觉得好，给我一个star！
	|	https://github.com/hunterhug/GoTaoBao
	|	QQ：459527502 版本： v1.0
	|	更新于： 20171029
	---------------------------------------------
	`)

	for {
		fmt.Println(`
	-------温柔的提示框---------
	|天猫淘宝搜索框小工具: 请按 1 |
	|天猫淘宝啥图片小工具: 请按 2 |
	|更多待续更多待续更多: 请按 x |
	--------------------------
		`)
		choice := util.Input("* 请你输入你要使用的功能:", "0")
		switch choice {
		case "1":
			src.SearchMain()
		case "2":
			src.DownloadPicMain()
		case "0":
			hello()
		default:
			hello()

		}
	}
}

func hello() {
	fmt.Println(`
	--
	- - -
	-
	--- -- - -------------
	---
	----------输入错误----- - - -- - -
	-----  -
	-
	-  --   - -- - --
	-  - -囧 -------
	   - - --    ---
	`)
}
