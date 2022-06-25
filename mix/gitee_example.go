package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
)

/**
 * @author       weimenghua
 * @time         2023-07-12 14:00
 * @description
 */

func get_index() {
	resp, err := http.Get("https://gitee.com")

	if err != nil {
		return
	}

	if resp.StatusCode != http.StatusOK {
		return
	}

	io.Copy(os.Stdout, resp.Body)

	fmt.Println(os.Stdout)
}

func create_issure() {
	url1 := "https://api.gitee.com/enterprises/2095951/issues"
	data := []byte(`{
"31": "7,6",
"37": "4",
"63": "2",
"3469": "88999111",
"title": "需求0712GO",
"priority": 0,
"description_type": "json",
"description": "这是需求的内容模板",
"issue_type_id": 150199,
"project_id": 28888726,
"plan_started_at": "2023-07-12",
"deadline": "2023-08-12",
"security_hole": 0,
"page_source": "项目",
"description_json": "{\"type\":\"doc\",\"content\":[{\"type\":\"paragraph\",\"attrs\":{\"indent\":0,\"textAlign\":\"left\"},\"content\":[{\"type\":\"text\",\"text\":\"这是需求的内容模板\"}]}]}",
"assignee_id": 0,
"extra_fields": [
{
"issue_field_id": 33,
"value": ""
},
{
"issue_field_id": 3025,
"value": ""
},
{
"issue_field_id": 3024,
"value": ""
},
{
"issue_field_id": 63,
"value": "2"
},
{
"issue_field_id": 37,
"value": "4"
},
{
"issue_field_id": 3469,
"value": "88999111"
},
{
"issue_field_id": 31,
"value": "7,6"
}
]
}`)

	req, err := http.NewRequest("POST", url1, bytes.NewBuffer(data))

	client := &http.Client{}
	cookieValue := url.QueryEscape("sensorsdata2015jssdkchannel=%7B%22prop%22%3A%7B%22_sa_channel_landing_url%22%3A%22%22%7D%7D; alpha-enable=5; close_wechat_tour=true; user_locale=zh-CN; Hm_lvt_24f17767262929947cc3631f99bfd274=1688537940,1688625642,1689067323; sensorsdata2015jssdkcross=%7B%22distinct_id%22%3A%222317018%22%2C%22first_id%22%3A%22188f238026516e8-0c5d408b57b8a48-1b525634-1484784-188f23802662c30%22%2C%22props%22%3A%7B%22%24latest_traffic_source_type%22%3A%22%E7%9B%B4%E6%8E%A5%E6%B5%81%E9%87%8F%22%2C%22%24latest_search_keyword%22%3A%22%E6%9C%AA%E5%8F%96%E5%88%B0%E5%80%BC_%E7%9B%B4%E6%8E%A5%E6%89%93%E5%BC%80%22%2C%22%24latest_referrer%22%3A%22%22%7D%2C%22identities%22%3A%22eyIkaWRlbnRpdHlfY29va2llX2lkIjoiMTg3ZWYyZGVkYzUzMDYtMDFmNTY1MTBjOTI1ZGIxLTFkNTI1NjM0LTE0ODQ3ODQtMTg3ZWYyZGVkYzYxZmI5IiwiJGlkZW50aXR5X2xvZ2luX2lkIjoiMjMxNzAxOCJ9%22%2C%22history_login_id%22%3A%7B%22name%22%3A%22%24identity_login_id%22%2C%22value%22%3A%222317018%22%7D%2C%22%24device_id%22%3A%22187ef2dedc5306-01f56510c925db1-1d525634-1484784-187ef2dedc61fb9%22%7D; go_canary=hightest; Hm_lpvt_24f17767262929947cc3631f99bfd274=1689141880; csrf_token=2jiD9xmLny7JIvMGGdX%2BOsBzofQlQcbL5PVK3k0NrCMoNR14LUne6Br57HFSjUGxupPFMKJUnhA1D0Tf1QRoWw%3D%3D; gitee-session-n=UElpSTYyRlc4ckQ5NmtYY2NKUjhoNUFkTDJYNUlqTlc0WlZRNmJwdGV4bzdRSXRYMjZPT1Jud1o2MFgzNGlIbWNqaFNDYms3NHZzZEFHSlkyTVFibzNLdUMxRWJBckxNamFmdTlsenh3WUJPanlma1ZrVmY1WWpBcUpLTDBta3hxaitHdkxuL0tUdUtWYkg4YU1sM2VydVJPQzUwN0dWaTd2eFRBZTF3NjRJdnJDVWhURFBkTndPLzNXc0ZGbDNSRDhCdDk1MHg3aUFiM2xlM1pnbmRKZW04STNXeE00aDJlVTNRMGVrSmtPUmVBTTZ1YkNiOEIxeXpSekw2TVFxWkkyWmlqV3hteXp1Vlp6UDdQcmJQdzNEOURPZkRPVVFnWmZucU5heVN1VXNMcldvZHV5WklxMFdoZ3o2MStuYkdwY244ejBKWUJzaWw3NnJZaTNqNHRCWlB5UjdISXBBUGNDVlk3RkJwcERieGxMd1pzb1Q1R09ldmQwSElsVFRRNWJleTNGdHA2R1lqNVNwRlpxOGVpUmdaajd5UWRzbHRQV1hadDVrOE52bFdpaXlwb0xxVkxWM1IwL3U1KzU3ak1YRnlaTmlWRlNZQUZKcHV5bWI5Nk9MVXlxaVl2SUEzZjJKMmFBMDVKekNPU2ZWK0tLYUFQbktaVXBPN1R1M2Z3WHc5dFRrbCtEVStoVG1JQ3M3eGppVFhVekF5THJ4bnRWbTFPOGNEUnUzMUVyNElEd2ROdWh2VXJIVUFyYnArUFRHTk1JOUtDb3JZaGo0K3FrMFRWVkxQM2g2eVlOU2dZbVNxUkYwQWx1bz0tLVo5cGNtek5CTFY3U20vWVZ5NWlBbFE9PQ%3D%3D--6f5bcd5b545ba7af3bc3df6fa9613c839fb60128")
	cookie := http.Cookie{Name: "Cookie", Value: cookieValue}
	req.AddCookie(&cookie)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	resp, err := client.Do(req)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)

	fmt.Println(os.Stdout)
}
