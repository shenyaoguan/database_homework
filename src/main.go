package main

import (
	"src/config"
	"src/cui"
)

func main() {
	config.ConnectDatabase()
	cui.StartCUI()
}
