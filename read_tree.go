package main

import (
	"encoding/json"
	"fmt"
)

var names []string

func foTree(i interface{}) {
	arr, ok := i.(map[string]interface{})
	if ok {
		var str string
		for key, value := range arr {
			if key == "name" {
				str = value.(string)
				names = append(names, str)
			}
			if key == "children" {
				m, ook := value.([]interface{})
				if ook {
					for _, v := range m {
						foTree(v)
					}
				}
			}
		}
	}
}

func main() {
	j := `{
		"Status": 200,
		"Message": "",
		"Result": {
			"Node": {
				"id": "1",
				"name": "成龙",
				"type": "listed",
				"Category": 0,
				"Level": 1,
				"children": [
					{
						"id": "",
						"name": "基本信息",
						"Category": 1,
						"Level": 2,
						"children": [
							{
								"id": "",
								"name": "中文名：成龙",
								"type": "baseifno",
								"Category": 1,
								"Level": 3,
								"children": null
							},
							{
								"id": "",
								"name": "外文名：Jackie Chan",
								"type": "baseifno",
								"Category": 1,
								"Level": 3,
								"children": null
							},
							{
								"id": "",
								"name": "别名：陈港生（原名）、房仕龙、元楼、陈元龙",
								"type": "baseifno",
								"Category": 1,
								"Level": 3,
								"children": null
							},
							{
								"id": "",
								"name": "国籍:中国",
								"type": "baseifno",
								"Category": 1,
								"Level": 3,
								"children": null
							},{
								"id": "",
								"name": "民族:汉族",
								"type": "baseifno",
								"Category": 1,
								"Level": 3,
								"children": null
							},{
								"id": "",
								"name": "星座:白羊座",
								"type": "baseifno",
								"Category": 1,
								"Level": 3,
								"children": null
							},{
								"id": "",
								"name": "……",
								"type": "baseifno",
								"Category": 1,
								"Level": 3,
								"children": null
							}
						]
					},
					{
						"id": "",
						"name": "主要成就",
						"Category": 2,
						"Level": 2,
						"children": [
							{
								"id": "",
								"name": "第25届中国电影金鸡奖最佳男主角奖",
								"type": "achievement",
								"Category": 2,
								"Level": 3,
								"children": null
							},
							{
								"id": "",
								"name": "两届台湾电影金马奖最佳男主角奖",
								"type": "achievement",
								"Category": 2,
								"Level": 3,
								"children": null
							},{
								"id": "",
								"name": "三届香港电影金像奖最佳动作设计奖",
								"type": "achievement",
								"Category": 2,
								"Level": 3,
								"children": null
							},{
								"id": "",
								"name": "两届台湾电影金马奖最佳动作设计奖",
								"type": "achievement",
								"Category": 2,
								"Level": 3,
								"children": null
							},{
								"id": "",
								"name": "第89届奥斯卡金像奖终身成就奖",
								"type": "achievement",
								"Category": 2,
								"Level": 3,
								"children": null
							}
						]
					},
					{
						"id": "",
						"name": "代表作品",
						"Category": 3,
						"Level": 2,
						"children": [
							{
								"id": "",
								"name": "醉拳",
								"type": "achievement",
								"Category": 3,
								"Level": 3,
								"children": null
							},
							{
								"id": "",
								"name": "A计划",
								"type": "achievement",
								"Category": 3,
								"Level": 3,
								"children": [
									{
										"id": "",
										"name": "主演",
										"type": "achievement",
										"Category": 3,
										"Level": 4,
										"children": [
											{
												"id": "",
												"name": "成龙",
												"type": "achievement",
												"Category": 3,
												"Level": 5,
												"children": null
											},
											{
												"id": "",
												"name": "洪金宝",
												"type": "achievement",
												"Category": 3,
												"Level": 5,
												"children": null
											},
											{
												"id": "",
												"name": "元彪",
												"type": "achievement",
												"Category": 3,
												"Level": 5,
												"children": null
											}
										]
									}
								]
							},{
								"id": "",
								"name": "警察故事",
								"type": "achievement",
								"Category": 3,
								"Level": 3,
								"children": null
							},{
								"id": "",
								"name": "尖峰时刻",
								"type": "achievement",
								"Category": 3,
								"Level": 3,
								"children": null
							},{
								"id": "",
								"name": "十二生肖",
								"type": "achievement",
								"Category": 3,
								"Level": 3,
								"children": null
							},{
								"id": "",
								"name": "红番区",
								"type": "achievement",
								"Category": 3,
								"Level": 3,
								"children": null
							},{
								"id": "",
								"name": "神话",
								"type": "achievement",
								"Category": 3,
								"Level": 3,
								"children": null
							},{
								"id": "",
								"name": "……",
								"type": "achievement",
								"Category": 3,
								"Level": 3,
								"children": null
							}
						]
					}
				]
			}
		}
	}`
	m2 := make(map[string]interface{})
	err2 := json.Unmarshal([]byte(j), &m2)
	if err2 != nil {
		fmt.Println("err = ", err2)
		return
	}
	m3 := m2["Result"].(map[string]interface{})
	m4 := m3["Node"].(map[string]interface{})
	foTree(m4)
	fmt.Println("result=", names)
}
