package main

import (
	ip "forum/Handler"
)

func main() {
	ip.Run()
	ip.DB.Close()
}