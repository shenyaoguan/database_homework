package cui

import (
	"fmt"
	"src/controllers"
	"src/models"
)

func Query() {
	fmt.Println("1. 查询飞机")
	fmt.Println("2. 查询酒店")
	fmt.Println("3. 查询巴士")
	fmt.Println("4. 查询客户的旅行线路")
	fmt.Print("请选择查询类型: ")
	var queryType int
	fmt.Scanln(&queryType)
	switch queryType {
	case 1:
		fmt.Println("请输入要查询的航班号: ")
		var flightNum string
		fmt.Scanln(&flightNum)
		flight, err := controllers.GetFlightByNum(flightNum)
		if err != nil {
			fmt.Printf("查询失败: %v\n", err)
			return
		}
		fmt.Println("航班号: ", flight.FlightNum)
		fmt.Println("航班座位数: ", flight.NumAvail)
		fmt.Println("航班起点: ", flight.FromCity)
		fmt.Println("航班终点: ", flight.ArivCity)
		fmt.Println("航班价格: ", flight.Price)
		fmt.Println("航班剩余座位: ", flight.NumAvail)
	case 2:
		fmt.Println("请输入要查询的酒店位置: ")
		var location string
		fmt.Scanln(&location)
		location += "Hotel"
		hotel, err := controllers.GetHotelByLocation(location)
		if err != nil {
			fmt.Printf("查询失败: %v\n", err)
			return
		}
		fmt.Println("酒店位置: ", hotel.Location)
		fmt.Println("酒店总房间数: ", hotel.NumRooms)
		fmt.Println("酒店剩余房间数: ", hotel.NumAvail)
		fmt.Println("酒店价格: ", hotel.Price)

	case 3:
		fmt.Println("请输入要查询的巴士位置: ")
		var location string
		fmt.Scanln(&location)
		location += "Bus"
		bus, err := controllers.GetBusByLocation(location)
		if err != nil {
			fmt.Printf("查询失败: %v\n", err)
			return
		}
		fmt.Println("巴士位置: ", bus.Location)
		fmt.Println("巴士总座位数: ", bus.NumBus)
		fmt.Println("巴士剩余座位数: ", bus.NumAvail)
		fmt.Println("巴士价格: ", bus.Price)
	case 4:
		fmt.Println("请输入要查询的客户姓名: ")
		var name string
		_, err := fmt.Scanln(&name)
		if err != nil {
			fmt.Println("输入错误")
			return
		}
		reservations, err := controllers.GetCustomerItinerary(name)
		if err != nil {
			fmt.Printf("查询失败: %v\n", err)
			return
		}
		for i, reservation := range reservations {
			fmt.Println("预定", i+1)
			switch res := reservation.(type) {
			case models.Flight:
				fmt.Println("航班号: ", res.FlightNum)
				fmt.Println("航班起点: ", res.FromCity)
				fmt.Println("航班终点: ", res.ArivCity)
			case models.Hotel:
				fmt.Println("酒店位置: ", res.Location)
			case models.Bus:
				fmt.Println("巴士位置: ", res.Location)
			default:
				fmt.Println("未知的预定类型")
			}
			fmt.Println("====================================")
		}

	default:
		fmt.Println("无效的选择")
	}
}
