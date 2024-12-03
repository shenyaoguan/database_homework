package cui

import (
	"fmt"
	"src/controllers"
	"src/models"
)

func Reverse() {
	fmt.Println("1. 飞机")
	fmt.Println("2. 酒店")
	fmt.Println("3. 巴士")
	fmt.Print("请选择预定类型: ")
	var resvType int
	fmt.Scanln(&resvType)
	switch resvType {
	case 1:
		var flight models.Reservation
		fmt.Println("请输入要预定的航班号: ")
		fmt.Scanln(&flight.ResvKey)
		fmt.Println("请输入预定者的姓名: ")
		fmt.Scanln(&flight.CustName)
		flight.ResvType = 1
		if err := controllers.AddReservation(flight); err != nil {
			fmt.Printf("预定失败: %v", err)
			return
		}
	case 2:
		var hotel models.Reservation
		fmt.Println("请输入要预定的酒店位置: ")
		fmt.Scanln(&hotel.ResvKey)
		fmt.Println("请输入预定者的姓名: ")
		fmt.Scanln(&hotel.CustName)
		hotel.ResvType = 2
		if err := controllers.AddReservation(hotel); err != nil {
			fmt.Printf("预定失败: %v\n", err)
			return
		}
	case 3:
		var bus models.Reservation
		fmt.Println("请输入要预定的巴士位置: ")
		fmt.Scanln(&bus.ResvKey)
		fmt.Println("请输入预定者的姓名: ")
		fmt.Scanln(&bus.CustName)
		bus.ResvType = 3
	default:
		fmt.Println("无效的选择")
	}
}
