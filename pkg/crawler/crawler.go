package crawler

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"sort"
	"strconv"
	"strings"

	"github.com/lichmaker/boring-lottery/pkg/config"
)

const DOMAIN string = "www.cwl.gov.cn"
const URL string = "http://www.cwl.gov.cn/ygkj/wqkjgg/"

type LotteryResult struct {
	Period string
	Blue   uint64
	Red1   uint64
	Red2   uint64
	Red3   uint64
	Red4   uint64
	Red5   uint64
	Red6   uint64
}

func Get() (*LotteryResult, error) {
	dataString, _ := crawl()
	if len(dataString) <= 0 {
		return nil, fmt.Errorf("爬虫数据读取失败！")
	}
	return parse(dataString)
}

func parse(sourceData string) (*LotteryResult, error) {
	preiodPreComplie, _ := regexp.Compile(`<div class="ssqQh-dom">(\d+)</div>`)
	preiodPre := preiodPreComplie.FindString(sourceData)
	if len(preiodPre) <= 0 {
		return nil, fmt.Errorf("爬虫请求获取数据失败！请检查UA、COOKIE等配置！")
	}

	bluePreComplie, _ := regexp.Compile(`<div class="ssqBlue-dom">\[\d+\]</div>`)
	bluePre := bluePreComplie.FindString(sourceData)

	redPreComplie, _ := regexp.Compile(`<div class="ssqRed-dom">\[\d{2},\d{2},\d{2},\d{2},\d{2},\d{2}\]</div>`)
	redPre := redPreComplie.FindString(sourceData)

	numberComplie, _ := regexp.Compile(`(\d+)`)
	preiod := numberComplie.FindString(preiodPre)
	blue := numberComplie.FindString(bluePre)

	redComplie, _ := regexp.Compile(`\d{2},\d{2},\d{2},\d{2},\d{2},\d{2}`)
	red := strings.Split(redComplie.FindString(redPre), ",")
	sort.StringSlice(red).Sort()

	data := new(LotteryResult)
	data.Period = preiod
	data.Blue, _ = strconv.ParseUint(blue, 10, 64)
	data.Red1, _ = strconv.ParseUint(red[0], 10, 64)
	data.Red2, _ = strconv.ParseUint(red[1], 10, 64)
	data.Red3, _ = strconv.ParseUint(red[2], 10, 64)
	data.Red4, _ = strconv.ParseUint(red[3], 10, 64)
	data.Red5, _ = strconv.ParseUint(red[4], 10, 64)
	data.Red6, _ = strconv.ParseUint(red[5], 10, 64)
	return data, nil
}

func crawl() (string, error) {
	request := buildRequest()
	client := new(http.Client)
	respond, err := client.Do(request)
	if err != nil {
		fmt.Println("爬虫请求执行失败:", err)
		return "", err
	}
	defer respond.Body.Close()
	dataByte, _ := ioutil.ReadAll(respond.Body)
	return string(dataByte), nil
}

func buildRequest() *http.Request {
	headerUA := config.Viper.GetString("crawler.ua")
	cookieName := config.Viper.GetString("crawler.cookiename")
	cookieValue := config.Viper.GetString("crawler.cookievalue")

	request, _ := http.NewRequest("GET", URL, nil)
	request.Header.Set("User-Agent", headerUA)

	cookie := new(http.Cookie)
	cookie.Name = cookieName
	cookie.Value = cookieValue
	cookie.Domain = DOMAIN
	cookie.Path = "/"
	request.AddCookie(cookie)

	return request
}
