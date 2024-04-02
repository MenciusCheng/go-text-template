package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

// 循环调接口
func main() {
	//arr := generateDateRange("2023-01-01", "2023-03-31")
	//fmt.Println(arr)
	//for _, item := range arr {
	//	request(item[0], item[1])
	//}

	arr := generateDateRange("2019-05-01", "2024-03-31")
	fmt.Println(arr)

	//for _, item := range arr {
	//	request_sv_usr_communityLike(item[0], item[1])
	//	time.Sleep(6 * time.Second)
	//}

	//for i, item := range arr {
	//	// sv_usr_communityTrends
	//	requestById(item[0], item[1], 47)
	//	if i != len(arr)-1 {
	//		time.Sleep(6 * time.Second)
	//	}
	//}

	//for i, item := range arr {
	//	// sv_usr_logOnlinePlayerNew
	//	requestById(item[0], item[1], 51)
	//	if i != len(arr)-1 {
	//		time.Sleep(6 * time.Second)
	//	}
	//}

	//for i, item := range arr {
	//	// sv_usr_logGuardTeamBuyRecord
	//	requestById(item[0], item[1], 92)
	//	if i != len(arr)-1 {
	//		time.Sleep(6 * time.Second)
	//	}
	//}

	//for i, item := range arr {
	//	// sv_usr_turnAwardJoinLog
	//	requestById(item[0], item[1], 39)
	//	if i != len(arr)-1 {
	//		time.Sleep(6 * time.Second)
	//	}
	//}

	//for i, item := range arr {
	//	// sv_usr_logUfoBuyTicket
	//	requestById(item[0], item[1], 44)
	//	if i != len(arr)-1 {
	//		time.Sleep(6 * time.Second)
	//	}
	//}

	//for i, item := range arr {
	//	// sv_usr_logRrackEgg
	//	requestById(item[0], item[1], 34)
	//	if i != len(arr)-1 {
	//		time.Sleep(6 * time.Second)
	//	}
	//}

	//for i, item := range arr {
	//	// sv_usr_logGiftSend
	//	requestById(item[0], item[1], 153)
	//	if i != len(arr)-1 {
	//		time.Sleep(6 * time.Second)
	//	}
	//}

	//for i, item := range arr {
	//	// sv_usr_logStore
	//	requestById(item[0], item[1], 74)
	//	if i != len(arr)-1 {
	//		time.Sleep(6 * time.Second)
	//	}
	//}

}

func generateDateRange(startDate, endDate string) [][]string {
	layout := "2006-01-02"
	start, err := time.Parse(layout, startDate)
	if err != nil {
		fmt.Println("无效的起始日期格式")
		return nil
	}

	end, err := time.Parse(layout, endDate)
	if err != nil {
		fmt.Println("无效的结束日期格式")
		return nil
	}

	var dateRange [][]string

	// 循环生成日期范围
	for start.Before(end) || start.Equal(end) {
		monthEndDate := start.AddDate(0, 1, -1)
		if monthEndDate.After(end) { // 如果结束日期在当月之前，则取结束日期作为月末
			monthEndDate = end
		}

		dateRange = append(dateRange, []string{start.Format(layout), monthEndDate.Format(layout)})

		start = monthEndDate.AddDate(0, 0, 1)
	}

	return dateRange
}

func request(dateParam, endDateParam string) {
	url := "https://staging-api.miyachat.com/staging/1111/platform_stat_admin_sg/platform/"
	method := "POST"

	text := fmt.Sprintf(`{"Func":"ReRunDataSourceSync","Obj":"dataSourceSync","Param":{"accessPath":"/vue_data_source_sync","pageId":100345049,"accountId":713,"accountName":"程孟威","applicationId":39,"projectId":1003,"areaType":"sg","id":34,"dateParam":"%s","endDateParam":"%s"}}`, dateParam, endDateParam)

	payload := strings.NewReader(text)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Cookie", "xymh-username-ip=%E7%A8%8B%E5%AD%9F%E5%A8%81%2F183.6.114.210; xymh-cookie=079cb33eb9acfa300720357333aaa2294d656831d7d4b5a46051cf5d0d2a1774; table=cf8035a8a3a3c411e430e7da5db6f74964dbe3edbf781359a9af95d010cac938")
	req.Header.Add("User-Agent", "Apifox/1.0.0 (https://apifox.com)")
	req.Header.Add("content-type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("req: %s, body:%s\n", text, string(body))
}

func request_sv_usr_communityLike(dateParam, endDateParam string) {
	url := "https://staging-api.miyachat.com/staging/1111/platform_stat_admin_sg/platform/"
	method := "POST"

	text := fmt.Sprintf(`{"Func":"ReRunDataSourceSync","Obj":"dataSourceSync","Param":{"accessPath":"/vue_data_source_sync","pageId":100345049,"accountId":713,"accountName":"程孟威","applicationId":39,"projectId":1003,"areaType":"sg","id":48,"dateParam":"%s","endDateParam":"%s"}}`, dateParam, endDateParam)

	payload := strings.NewReader(text)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Cookie", "xymh-username-ip=%E7%A8%8B%E5%AD%9F%E5%A8%81%2F183.6.114.210; xymh-cookie=079cb33eb9acfa300720357333aaa2294d656831d7d4b5a46051cf5d0d2a1774; table=cf8035a8a3a3c411e430e7da5db6f74964dbe3edbf781359a9af95d010cac938")
	req.Header.Add("User-Agent", "Apifox/1.0.0 (https://apifox.com)")
	req.Header.Add("content-type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("req: %s, body:%s\n", text, string(body))
}

func requestById(dateParam, endDateParam string, id int) {
	url := "https://staging-api.miyachat.com/staging/1111/platform_stat_admin_sg/platform/"
	method := "POST"

	text := fmt.Sprintf(`{"Func":"ReRunDataSourceSync","Obj":"dataSourceSync","Param":{"accessPath":"/vue_data_source_sync","pageId":100345049,"accountId":713,"accountName":"程孟威","applicationId":39,"projectId":1003,"areaType":"sg","id":%d,"dateParam":"%s","endDateParam":"%s"}}`, id, dateParam, endDateParam)

	payload := strings.NewReader(text)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Cookie", "xymh-username-ip=%E7%A8%8B%E5%AD%9F%E5%A8%81%2F183.6.114.210; xymh-cookie=079cb33eb9acfa300720357333aaa2294d656831d7d4b5a46051cf5d0d2a1774; table=cf8035a8a3a3c411e430e7da5db6f74964dbe3edbf781359a9af95d010cac938")
	req.Header.Add("User-Agent", "Apifox/1.0.0 (https://apifox.com)")
	req.Header.Add("content-type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("req: %s, body:%s\n", text, string(body))
}
