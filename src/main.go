package main

import (
	"fmt"
	"src/config"
	"src/controllers"
	"src/models"
)

func main() {
	config.ConnectDatabase()

	for {
		fmt.Println("1. Add Flight")
		fmt.Println("2. View Flights")
		fmt.Println("3. Exit")
		fmt.Print("Enter your choice: ")

		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			var flight models.Flight
			fmt.Print("Enter FlightNum: ")
			fmt.Scan(&flight.FlightNum)
			fmt.Print("Enter Price: ")
			fmt.Scan(&flight.Price)
			fmt.Print("Enter NumSeats: ")
			fmt.Scan(&flight.NumSeats)
			flight.NumAvail = flight.NumSeats
			fmt.Print("Enter FromCity: ")
			fmt.Scan(&flight.FromCity)
			fmt.Print("Enter ArivCity: ")
			fmt.Scan(&flight.ArivCity)
			if err := controllers.AddFlight(flight); err != nil {
				fmt.Println("Error:", err)
			}
		case 2:
			flights, err := controllers.GetAllFlights()
			if err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("Available Flights:")
				for _, flight := range flights {
					fmt.Printf("FlightNum: %s, From: %s, To: %s, Price: %d, Seats Available: %d\n",
						flight.FlightNum, flight.FromCity, flight.ArivCity, flight.Price, flight.NumAvail)
				}
			}
		case 3:
			return
		default:
			fmt.Println("Invalid choice")
		}
	}
}
