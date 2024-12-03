package cui

import (
	"fmt"
	"src/controllers"
)

func Check() {
	fmt.Println("请输入要检查人的姓名: ")
	var name string
	_, err := fmt.Scanln(&name)
	if err != nil {
		fmt.Println("输入错误")
		return
	}
	fmt.Println("请输入要检查的路线终点: ")
	err = controllers.CheckReservationIntegrity(name)
	if err != nil {
		fmt.Printf("检查失败: %v\n", err)
		return
	}
	fmt.Println("检查成功")
}
