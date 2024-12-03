package cui

import (
	"fmt"
	"src/controllers"
	"src/models"
)

func Change() {
	fmt.Println("请选择想要使用的功能: ")
	fmt.Println("1. 添加功能")
	fmt.Println("2. 删除功能")
	fmt.Println("3. 更新功能")
	fmt.Println("4. 返回")
	fmt.Print("请选择功能: ")
	var id int
	fmt.Scanln(&id)
	switch id {
	case 1:
		fmt.Println("添加功能")
		Add()
	case 2:
		fmt.Println("删除功能")
		Delete()
	case 3:
		fmt.Println("更新功能")
		Update()
	case 4:
		fmt.Println("返回")
		return
	default:
		fmt.Println("无效的选择")
		return
	}
}

func Add() {
	fmt.Println("请选择要添加的类型: ")
	fmt.Println("1. 飞机")
	fmt.Println("2. 酒店")
	fmt.Println("3. 巴士")
	fmt.Print("请选择类型: ")
	var id int
	fmt.Scanln(&id)
	switch id {
	case 1:
		var flight models.Flight
		fmt.Println("添加飞机")
		fmt.Print("请输入航班号: ")
		fmt.Scanln(&flight.FlightNum)
		fmt.Print("请输入起点: ")
		fmt.Scanln(&flight.FromCity)
		fmt.Print("请输入终点: ")
		fmt.Scanln(&flight.ArivCity)
		fmt.Print("请输入价格: ")
		fmt.Scanln(&flight.Price)
		fmt.Print("请输入座位数: ")
		fmt.Scanln(&flight.NumSeats)
		fmt.Print("请输入剩余座位数: ")
		fmt.Scanln(&flight.NumAvail)
		if err := controllers.AddFlight(flight); err != nil {
			fmt.Printf("添加失败: %v\n", err)
			return
		}
		fmt.Println("添加成功")
	case 2:
		var hotel models.Hotel
		fmt.Println("添加酒店")
		fmt.Print("请输入位置: ")
		fmt.Scanln(&hotel.Location)
		hotel.Location += "Hotel"
		fmt.Print("请输入价格: ")
		fmt.Scanln(&hotel.Price)
		fmt.Print("请输入总房间数: ")
		fmt.Scanln(&hotel.NumRooms)
		fmt.Print("请输入剩余房间数: ")
		fmt.Scanln(&hotel.NumAvail)
		if err := controllers.AddHotel(hotel); err != nil {
			fmt.Printf("添加失败: %v\n", err)
			return
		}
		fmt.Println("添加成功")
	case 3:
		var bus models.Bus
		fmt.Println("添加巴士")
		fmt.Print("请输入位置: ")
		fmt.Scanln(&bus.Location)
		bus.Location += "Bus"
		fmt.Print("请输入价格: ")
		fmt.Scanln(&bus.Price)
		fmt.Print("请输入总座位数: ")
		fmt.Scanln(&bus.NumBus)
		fmt.Print("请输入剩余座位数: ")
		fmt.Scanln(&bus.NumAvail)
		if err := controllers.AddBus(bus); err != nil {
			fmt.Printf("添加失败: %v\n", err)
			return
		}
		fmt.Println("添加成功")
	default:
		fmt.Println("无效的选择")
		return
	}
}

func Delete() {
	fmt.Println("请选择要删除的类型: ")
	fmt.Println("1. 飞机")
	fmt.Println("2. 酒店")
	fmt.Println("3. 巴士")
	fmt.Print("请选择类型: ")
	var id int
	fmt.Scanln(&id)
	switch id {
	case 1:
		fmt.Println("删除飞机")
		fmt.Print("请输入航班号: ")
		var flightNum string
		fmt.Scanln(&flightNum)
		if err := controllers.DeleteFlight(flightNum); err != nil {
			fmt.Printf("删除失败: %v\n", err)
			return
		}
		fmt.Println("删除成功")
	case 2:
		fmt.Println("删除酒店")
		fmt.Print("请输入位置: ")
		var location string
		fmt.Scanln(&location)
		location += "Hotel"
		if err := controllers.DeleteHotel(location); err != nil {
			fmt.Printf("删除失败: %v\n", err)
			return
		}
		fmt.Println("删除成功")
	case 3:
		fmt.Println("删除巴士")
		fmt.Print("请输入位置: ")
		var location string
		fmt.Scanln(&location)
		location += "Bus"
		if err := controllers.DeleteBus(location); err != nil {
			fmt.Printf("删除失败: %v\n", err)
			return
		}
		fmt.Println("删除成功")
	default:
		fmt.Println("无效的选择")
		return
	}
}

func Update() {
	fmt.Println("请选择要更新的类型: ")
	fmt.Println("1. 飞机")
	fmt.Println("2. 酒店")
	fmt.Println("3. 巴士")
	fmt.Print("请选择类型: ")
	var id int
	fmt.Scanln(&id)
	switch id {
	case 1:
		fmt.Println("更新飞机")
		fmt.Print("请输入航班号: ")
		var flightNum string
		fmt.Scanln(&flightNum)
		var flight models.Flight
		flight.FlightNum = flightNum
		fmt.Print("请输入新的起点: ")
		fmt.Scanln(&flight.FromCity)
		fmt.Print("请输入新的终点: ")
		fmt.Scanln(&flight.ArivCity)
		fmt.Print("请输入新的价格: ")
		fmt.Scanln(&flight.Price)
		fmt.Print("请输入新的座位数: ")
		fmt.Scanln(&flight.NumSeats)
		fmt.Print("请输入新的剩余座位数: ")
		fmt.Scanln(&flight.NumAvail)
		if err := controllers.UpdateFlight(flightNum, flight); err != nil {
			fmt.Printf("更新失败: %v\n", err)
			return
		}
		fmt.Println("更新成功")
	case 2:
		fmt.Println("更新酒店")
		fmt.Print("请输入位置: ")
		var location string
		fmt.Scanln(&location)
		location += "Hotel"
		var hotel models.Hotel
		hotel.Location = location
		fmt.Print("请输入新的价格: ")
		fmt.Scanln(&hotel.Price)
		fmt.Print("请输入新的总房间数: ")
		fmt.Scanln(&hotel.NumRooms)
		fmt.Print("请输入新的剩余房间数: ")
		fmt.Scanln(&hotel.NumAvail)
		if err := controllers.UpdateHotel(location, hotel); err != nil {
			fmt.Printf("更新失败: %v\n", err)
			return
		}
		fmt.Println("更新成功")
	case 3:
		fmt.Println("更新巴士")
		fmt.Print("请输入位置: ")
		var location string
		fmt.Scanln(&location)
		location += "Bus"
		var bus models.Bus
		bus.Location = location
		fmt.Print("请输入新的价格: ")
		fmt.Scanln(&bus.Price)
		fmt.Print("请输入新的总座位数: ")
		fmt.Scanln(&bus.NumBus)
		fmt.Print("请输入新的剩余座位数: ")
		fmt.Scanln(&bus.NumAvail)
		if err := controllers.UpdateBus(location, bus); err != nil {
			fmt.Printf("更新失败: %v\n", err)
			return
		}
		fmt.Println("更新成功")
	default:
		fmt.Println("无效的选择")
		return
	}
}
