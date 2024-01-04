package main

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/duke-git/lancet/v2/random"
	"github.com/imroc/req/v3"
	"github.com/samber/lo"
)

const (
	url       = "https://sc.122.gov.cn/m/mvehxh/getTfhdList"
	cookie    = "_uab_collina=169448401324509842614666; userpub=1; JSESSIONID-L=4af75bae-f20e-4e6a-a90c-c66a341e8b39; accessToken=TSwxFlfrNdFAvckVYCu7aVFsVHrR/7gYzl6aT6Z74NJL4WBaw5AEHxlYol8EqainQ0AKeCTcTVLOYOO3Eg19ba+ig1zYn0PY3UmYFxOoDnn5XqQu3cj3FtZbohHUNAYwMmYcl/WdGZXbW4Cn0WQk+vCYRCV7oj3RTOcZN3KOUtVEmTczDOcIGyQ6La2NrvNJ; _122_gt=WGT-249727-UCusbmgvXb7Wg3qEMC7mvEGbeKclVJWGtvy; _122_gt_tag=1; tmri_csfr_token=CF2D7CE9AECAE0086A1D0CA5D565C12F"
	pageTotal = 7
)

var (
	now                = time.Now().Local()
	startDate, endDate = now.AddDate(0, -2, 0).Format("2006-01-02"), now.Format("2006-01-02")
	randomStr          = random.RandNumeralOrLetter(8)
	param              = fmt.Sprintf("glbm=510100000400&hpzl=52&type=0&startTime=%s&endTime=%s&random=%s", startDate, endDate, randomStr)
	cli                = req.C().SetCommonHeaders(map[string]string{
		"Cookie":       cookie,
		"Content-Type": "application/x-www-form-urlencoded; charset=UTF-8",
	})
)

func main() {
	lastDate := ""
	fmt.Println("放号时间   川A混动|号段|随机段  nice程度")
	for page := 1; page <= pageTotal; page++ {
		items, err := post(page)
		if err != nil {
			panic(err)
		}

		for _, item := range items {
			numberRange := item.NumberRange
			if lo.Contains([]byte{'6', '8'}, numberRange[5]) {
				if numberRange[4] == numberRange[5] {
					numberRange = fmt.Sprintf("%s     !!!!", numberRange)
				} else if lo.Contains([]byte{'6', '8'}, numberRange[4]) {
					numberRange = fmt.Sprintf("%s     !!", numberRange)
				} else {
					numberRange = fmt.Sprintf("%s     !", numberRange)
				}
			} else if numberRange[4] == numberRange[5] {
				numberRange = fmt.Sprintf("%s     !!!", numberRange)
			}

			if item.Time[:6] != lastDate {
				fmt.Printf("%s     %s\n", item.Time, numberRange)
			} else {
				fmt.Printf("      %s     %s\n", item.Time[6:], numberRange)
			}
			lastDate = item.Time[:6]
		}
		time.Sleep(time.Second)
	}
}

func post(page int) (items []*Item, err error) {
	items = make([]*Item, 0)
	body, rsp := fmt.Sprintf("page=%d&%s", page, param), new(Rsp)

	ret, err := cli.R().SetBody(body).SetSuccessResult(rsp).Post(url)
	if err != nil {
		return
	}
	if ret.IsErrorState() {
		return nil, errors.New("请求失败")
	}

	if rsp.Code != 200 {
		return nil, errors.New(rsp.Msg)
	}
	if len(rsp.Data.List.Items) == 0 {
		return nil, errors.New("没有数据返回")
	}
	for _, item := range rsp.Data.List.Items {
		if item.Location != "成都车辆管理所" || item.Vehicle != "小型新能源汽车" || item.NumberRange[0] != 'A' {
			return nil, errors.New("过滤条件错误")
		}
		if item.NumberRange[1] != 'F' || strings.Contains(item.NumberRange, "4") {
			continue
		}
		item.Time = item.Time[5:]
		item.NumberRange = fmt.Sprintf("%s %s xx", item.NumberRange[:2], item.NumberRange[2:5])
		items = append(items, item)
	}

	return
}

type Rsp struct {
	Code int    `json:"code"`
	Data Data   `json:"data"`
	Msg  string `json:"message"`
}

type Data struct {
	List List `json:"list"`
}

type List struct {
	Items []*Item `json:"content"`
}

type Item struct {
	Location    string `json:"bmmc"`
	Vehicle     string `json:"hpzlStr"`
	Time        string `json:"tfrq"`
	NumberRange string `json:"subhd"`
}
