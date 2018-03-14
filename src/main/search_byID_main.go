package main

import (
	"github.com/hughimr/GoTaoBao/src"

	"fmt"
	//"time"
	"os"
	"bufio"
	"io"
	"strconv"
	"bytes"
	"path/filepath"
	//"time"
	"strings"
	//"time"
	"time"
	"math/rand"
	"sort"
)

//func agent(jobs <-chan int, results chan<- string) {
//	for j := range jobs {
//		fmt.Println(j)
//		//fmt.Println("worker", id, "processing job", j)
//		res:=src.SearchByID(j)
//		time.Sleep(5*time.Second)
//		results <- res
//	}
//}

func readToList(filename string) []int {
	idList:=[]int{}
	fi, err := os.Open(filename)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
	}
	defer fi.Close()

	br := bufio.NewReader(fi)
	for {
		a, _, c := br.ReadLine()
		if c == io.EOF {
			break
		}
		if string(a)!=""{
			ai,_:=strconv.Atoi(strings.Split(string(a),",")[17])
			idList=append(idList, ai)
		}

	}
	return idList
}

func writeToFile(buffer bytes.Buffer){

	filekeep :=  filepath.Base("")+ "/" + "brand.txt"
	//util.SaveToFile(filekeep, []byte(tempdata))
	os.Remove(filekeep)
	filename,_:=os.Create(filekeep)
	buffer.WriteTo(filename)
	filename.Close()
}

func RemoveDuplicatesAndEmpty2(a []int) (ret []int){
	a_len := len(a)
	for i:=0; i < a_len; i++{
		if (i > 0 && a[i-1] == a[i]) || len(strconv.Itoa(a[i]))==0{
			continue
		}
		ret = append(ret, a[i])
	}
	return
}

func main() {
	var buffer bytes.Buffer
	buffer.Reset()
	// 创建channel
	//jobs := make(chan int, 10000)
	//results := make(chan string, 10000)

	//
	//go agent(jobs, results)

	//results:=[]string{}

	idListRaw:=readToList("/Users/chenminyu/IdeaProjects/GoTaoBao/src/main/201802250307-ba8rfl06v8t098derlr0.csv")
	//// 发送9个jobs，然后关闭
	//for j := 0; j < len(idList); j++ {
	//	jobs <- idList[j]
	//}
	//close(jobs)

	sort.Ints(idListRaw)


	idList:=RemoveDuplicatesAndEmpty2(idListRaw)
	listLength:=len(idList)
	il:=0

	fmt.Println(len(idListRaw),len(idList))
	for _,j:=range idList{

		rand.Seed(time.Now().UnixNano())
		t:=rand.Intn(3)+1
		time.Sleep(time.Duration(t)*time.Second)
		res,err:=src.SearchByID(j)
		if err==nil{
			buffer.WriteString(fmt.Sprintf("%s\n",res))
			fmt.Println(res)
		}else{
			fmt.Println(j,err)
			time.Sleep(10*time.Second)
		}

		il++
		fmt.Println("剩余",listLength-il,"个")


	}

	// 最后收集结果
	//for a := 0; a < len(idList); a++ {
	//	res:=<-results
	//	buffer.WriteString(fmt.Sprintf("%s\n",res))
	//}

	//for _,r:=range results{
	//	buffer.WriteString(fmt.Sprintf("%s\n",r))
	//}

	defer writeToFile(buffer)

}


//func main() {
//
//
//
//
//	fmt.Println(src.SearchByID(532774062476))
//}

