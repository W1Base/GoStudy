package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	const url = "https://forum.butian.net/community/all/hottest?page="
	for i := 1; i <= 5; i++ {
		url := url + strconv.Itoa(i)
		// fmt.Println("***********************************************")
		// fmt.Println("正在搜集的URL", url)
		// fmt.Println("***********************************************")
		items := getTitle(url)
		for _, item := range items {
			fmt.Println("标题：", item[3], "         URL: ", item[1], "\n---------------------------------------------------------------")
		}
	}
}

func getBody(url string) string {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/99.0.4844.84 Safari/537.36")
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("http get err", err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	return string(body)
}

func getTitle(url string) [][]string {
	body := getBody(url)
	body = strings.Replace(body, "\n", "", -1)
	//fmt.Println(body)
	title := regexp.MustCompile(`                            href="(.*?)"                             target="_blank"                             rel="noopenner noreferrer"                             data-source-list="community_all_hottest"                            data-source-id="(.*?)"                        >(.*?)</a></h2>`)
	items := title.FindAllStringSubmatch(body, -1)
	return items
}
