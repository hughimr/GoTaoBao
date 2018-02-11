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
	//"regexp"
	"github.com/hunterhug/parrot/util"
	//"flag"
	//"runtime/pprof"
	//"os"


	"regexp"
	"time"
)
//
//var cpuprofile = flag.String("cpuprofile","", "write cpu profile `file`")
//var memprofile = flag.String("memprofile","", "write memory profile to `file`")
//
//var memFile *os.File
//var cpuFile *os.File

func main() {
	//
	//flag.Parse()
	//if *cpuprofile != "" {
	//	cpuFile,err := os.Create(*cpuprofile)
	//	if err != nil {
	//		log.Fatal("could not create CPU profile: ", err)
	//	}
	//	if err := pprof.StartCPUProfile(cpuFile); err != nil {
	//		log.Fatal("could not start CPU profile: ", err)
	//	}
	//	defer pprof.StopCPUProfile()
	//}
	//
	//
	//
	//	memFile,err := os.Create(*memprofile)
	//	if err != nil {
	//		log.Fatal("could not create memory profile: ", err)
	//	}



	//runtime.GC() // get up-to-date statistics
	//miner.SetLogLevel("debug")
	/*
		执行爬取天猫淘宝商品信息
	*/
	//src.MySearchMain("加湿器")

	//自定义关键字
	keyWordList := []string{"苹果手机壳"}
	//拿淘宝分类关键字
	for _, v := range src.GetKeywords() {

		var m src.KeyItem
		err := json.Unmarshal(v, &m)
		if err == nil {
			//fmt.Println(m.LevelThree)
			pattern, _ := regexp.Compile(`.*家电|路由|电器|存储|耳机|3C|智能|壳|灯具.*`)
			if pattern.Match([]byte(m.LevelOne)) || pattern.Match([]byte(m.LevelTwo)) || pattern.Match([]byte(m.LevelThree)) {
				//去掉不相干的关键字
				pattern2, _ := regexp.Compile(".*服务|元器件|商用家具|五金工具|个性定制|文化|iPhone6|包装|A4纸.*")
				if !pattern2.Match([]byte(m.LevelOne)) && !pattern2.Match([]byte(m.LevelTwo)) && !pattern2.Match([]byte(m.LevelThree)) && len(m.LevelThree) != 0 {

					//fmt.Println(m.LevelOne,m.LevelTwo,m.LevelThree)
					keyWordList = append(keyWordList, string(m.LevelThree))

				}

			}
		}
	}

	keyWordList=[]string{"电饭锅"}
	t1:=time.Now()
	for _, v := range keyWordList {
		fmt.Printf("开始抓关键字%s\n",v)
		src.MySearchMain(v)


		//if err := pprof.WriteHeapProfile(memFile); err != nil {
		//	log.Fatal("could not write memory profile: ", err)
		//}


		util.Sleep(5)
	}

	elapsed:=time.Since(t1)

	fmt.Println("程序运行时间:",elapsed)

	//defer memFile.Close()
	//pprof.StopCPUProfile()
	//cpuFile.Close()
	//memFile.Close()


}
