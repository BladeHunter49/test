package main

import (
	"fmt"
	"mahonia"
	"net/http"
	"os"
	"time"

	//"strconv"
	//"strings"
	"regexp"
)

var str = make(chan string)
var filestr string

func ConvertToString(src string, srcCode string, tagCode string) string {

	srcCoder := mahonia.NewDecoder(srcCode)

	srcResult := srcCoder.ConvertString(src)

	tagCoder := mahonia.NewDecoder(tagCode)

	_, cdata, _ := tagCoder.Translate([]byte(srcResult), true)

	result := string(cdata)

	return result

}

func Getftp(data string) {
	resp, err1 := http.Get(data)
	if err1 != nil {
		fmt.Println("http get err:", err1)
		return
	}
	defer resp.Body.Close()

	var result string
	buf := make([]byte, 4*1024)
	for {
		n, _ := resp.Body.Read(buf)
		if n == 0 {
			break
		}
		result += string(buf[:n])
	}

	re := regexp.MustCompile(`<a href="(.*?)">ftp`)
	if re == nil {
		fmt.Println("regexp MustCompile err")
		return
	}
	downloadUrls := re.FindAllStringSubmatch(result, -1)
	for _, data := range downloadUrls {
		data[1] = ConvertToString(data[1], "gbk", "utf-8") + "\r\n"
		// fmt.Println(data[1])
		str <- data[1]
	}
	// ch <- 100
	return
}

func main() {
	url := "https://www.dy2018.com/i/100827.html"
	// str := make([]string, 0)

	resp, err1 := http.Get(url)
	if err1 != nil {
		fmt.Println("http get err:", err1)
		return
	}
	defer resp.Body.Close()

	var result string
	buf := make([]byte, 4*1024)
	for {
		n, _ := resp.Body.Read(buf)
		if n == 0 {
			break
		}
		result += string(buf[:n])
	}

	re := regexp.MustCompile(`<a href="(.*?)">https://www`)
	if re == nil {
		fmt.Println("regexp MustCompile err")
		return
	}
	downloadUrls := re.FindAllStringSubmatch(result, -1)
	for _, data := range downloadUrls {
		go Getftp(data[1])
	}

	f, err := os.Create("爬取下载链接.txt")
	if err != nil {
		fmt.Println("Create file err", err)
		return
	}
	defer f.Close()

	for {
		select {
		case <-time.After(10 * time.Second):
			fmt.Println("time out")
			break
		case <-str:
			filestr = <-str
			_, err := f.WriteString(filestr)
			if err != nil {
				fmt.Println("write file err")
				return
			}
		}
	}

	fmt.Println("program end")
}
