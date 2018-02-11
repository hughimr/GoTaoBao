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
	//"encoding/json"
	//"regexp"
	"github.com/hunterhug/parrot/util"
	//"flag"
	//"runtime/pprof"
	//"os"


	//"regexp"
	"time"

	"sort"
)
//
//var cpuprofile = flag.String("cpuprofile","", "write cpu profile `file`")
//var memprofile = flag.String("memprofile","", "write memory profile to `file`")
//
//var memFile *os.File
//var cpuFile *os.File

func RemoveDuplicatesAndEmpty(a []string) (ret []string){
	a_len := len(a)
	for i:=0; i < a_len; i++{
		if (i > 0 && a[i-1] == a[i]) || len(a[i])==0{
			continue;
		}
		ret = append(ret, a[i])
	}
	return
}

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



	//／*
	//	//自定义关键字
	//	keyWordList := []string{"苹果手机壳"}
	//	//拿淘宝分类关键字
	//	for _, v := range src.GetKeywords() {
	//
	//		var m src.KeyItem
	//		err := json.Unmarshal(v, &m)
	//		if err == nil {
	//			//fmt.Println(m.LevelThree)
	//			pattern, _ := regexp.Compile(`.*家电|路由|电器|存储|耳机|3C|智能|壳|灯具.*`)
	//			if pattern.Match([]byte(m.LevelOne)) || pattern.Match([]byte(m.LevelTwo)) || pattern.Match([]byte(m.LevelThree)) {
	//			//去掉不相干的关键字
	//				pattern2, _ := regexp.Compile(".*服务|元器件|商用家具|五金工具|个性定制|文化|iPhone6|包装|A4纸.*")
	//				if !pattern2.Match([]byte(m.LevelOne)) && !pattern2.Match([]byte(m.LevelTwo)) && !pattern2.Match([]byte(m.LevelThree)) && len(m.LevelThree) != 0 {
	//
	//					//fmt.Println(m.LevelOne,m.LevelTwo,m.LevelThree)
	//					keyWordList = append(keyWordList, string(m.LevelThree))
	//
	//				}
	//
	//			}
	//		}
	//	}
	//*／

	keyWordListRaw:=[]string{"苹果手机壳","吸顶灯","吊灯","筒灯","射灯","台灯","落地灯","室外灯","壁灯","小夜灯","智能马桶","智能马桶盖","智能车机","智能后视镜","充电器","路由器","充电宝","智能穿戴","蓝牙耳机","手机壳套","保护壳套","儿童手表","智能手表","智能手环","智能配饰","智能排插","智能眼镜","路由器","网络存储设备","路由器","耳机","U盘","闪存卡","记忆棒","移动硬盘","电磁炉","电水壶","料理机","电饭煲","榨汁机","净水器","豆浆机","烤箱","电风扇","空调扇","挂烫机","扫地机","吸尘器","加湿器","除湿机","对讲机","空气净化","理发器","电子称","美容仪","按摩椅","按摩披肩","血压计","足浴器","电动牙刷","剃须刀","耳机","音响","麦克风","扩音器","低音炮","打印机","投影仪","保险柜","学习机","冰箱","空调","平板电视","油烟机","燃气灶","消毒柜","热水器","洗衣机","智能钢琴","跑步机","早教机","游戏手柄","蓝牙耳机","电压力锅","微波炉","破壁机","原汁机","空气炸锅","电蒸锅","电炖盅","电饼档","电烤箱","面包机","咖啡机","厨师机","扫地机器人","净化器","挂烫机","取暖器","干衣机","剃须刀","美容仪","洁面仪","成人牙刷","儿童牙刷","冲牙器","电子秤","足浴盆","按摩椅","按摩器","理发器","电吹风"}

	sort.Strings(keyWordListRaw)

	keyWordList:=RemoveDuplicatesAndEmpty(keyWordListRaw)


	t1:=time.Now()
	for _, v := range keyWordList {
		fmt.Printf("开始抓关键字%s\n",v)
		src.MySearchMain(v)

		//fmt.Println(v)

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
