package show

import "fmt"

func ShowFinalResp(tag int, tmpRet map[string]interface{}) {
	switch tag {
	case 1:
		fmt.Printf("++++robot：%s", tmpRet["text"])
	case 2:
		fmt.Printf("++++robot：%s", tmpRet["text"])
		fmt.Printf("\n%s", tmpRet["url"])
	case 3:
		fmt.Printf("++++robot：%s", tmpRet["text"])
		newsList := tmpRet["list"].([]interface{})
		for idx, item := range newsList {
			fmt.Printf("\n -------------第%d条新闻-------------", idx+1)
			itemMap := item.(map[string]interface{})
			fmt.Printf("\n 新闻标题：%s", itemMap["article"])
			fmt.Printf("\n 来源：%s", itemMap["source"])
			fmt.Printf("\n 缩略图：%s", itemMap["icon"])
			fmt.Printf("\n 详细见：%s", itemMap["detailurl"])
		}
	case 4:
		fmt.Printf("++++robot：%s", tmpRet["text"])
		newsList := tmpRet["list"].([]interface{})
		for idx, item := range newsList {
			fmt.Printf("\n -------------第%d条做法-------------", idx+1)
			itemMap := item.(map[string]interface{})
			fmt.Printf("\n 菜名：%s", itemMap["name"])
			fmt.Printf("\n 缩略图：%s", itemMap["icon"])
			fmt.Printf("\n 评价以及热度：%s", itemMap["info"])
			fmt.Printf("\n 详细做法见：%s", itemMap["detailurl"])
		}
	}

}
