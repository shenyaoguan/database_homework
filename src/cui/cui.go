package cui

import (
	"fmt"
)

func StartCUI() {
	for {
		fmt.Println("欢迎使用旅行预定系统")
		fmt.Println("1. 入库功能")
		fmt.Println("2. 预定功能")
		fmt.Println("3. 查询功能")
		fmt.Println("4. 检查功能")
		fmt.Println("5. 退出")
		fmt.Print("请选择功能: ")
		var id int
		fmt.Scanln(&id)
		switch id {
		case 1:
			fmt.Println("入库功能")
			Change()
		case 2:
			fmt.Println("预定功能")
			Reverse()
		case 3:
			fmt.Println("查询功能")
			Query()
		case 4:
			fmt.Println("检查功能")
			Check()
		case 5:
			fmt.Println("退出")
			return
		default:
			fmt.Println("无效的选择")
		}
	}
}
