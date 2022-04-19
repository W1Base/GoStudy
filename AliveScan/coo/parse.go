package coo

import (
	"fmt"
	"strconv"
	"strings"
)

var target []string

func ParseIP(host string) []string {
	hosts := strings.Split(host, ".")
	_, err := strconv.Atoi(hosts[3]) // 将最后转化为int类型，如果转化成果则是一个IP，如果是1/24则失败，进入err
	if err != nil {
		if strings.HasSuffix(host, "/8") {

		} else if strings.HasSuffix(host, "/16") {
			addIPList(hosts[0]+"."+hosts[1], 16)
		} else if strings.HasSuffix(host, "/24") {
			addIPList(hosts[0]+"."+hosts[1]+"."+hosts[2], 24)
		} else if strings.Contains(host, "-") {

		}
	} else {
		target = append(target, host)
		fmt.Println("正在扫描单独IP地址", host)
	}
	return target
}

func addIPList(host string, mark int) {
	if mark == 8 {
		//目前的思路：扫1和255网关
		for i := 1; i <= 255; i++ {
			for j := 1; j <= 255; j++ {
				target = append(target, fmt.Sprintf("10.%v.%v.1", i, j))
			}
		}
		for i := 1; i <= 255; i++ {
			for j := 1; j <= 255; j++ {
				target = append(target, fmt.Sprintf("10.%v.%v.255", i, j))
			}
		}
	} else if mark == 16 {
		for i := 0; i <= 255; i++ {
			for j := 1; j <= 255; j++ {
				target = append(target, fmt.Sprintf("%s.%d.%d", host, i, j))
			}
		}
	} else if mark == 24 {
		for j := 1; j <= 255; j++ {
			target = append(target, fmt.Sprintf("%s.%d", host, j))
		}
	}
}
