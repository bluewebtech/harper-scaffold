package main

import (
	"app/bootstrap"
)

func main() {
	Boot := bootstrap.Run()
	Boot.Run(":8080")
}
