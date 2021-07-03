package main

import (
	"fmt"
)

func main() {
	var instancia string
	var ipaddr string

	fmt.Scan(&instancia)
	fmt.Scan(&ipaddr)
	fmt.Println("The instance ", instancia, " containing the ", ipaddr, "ip address.")

}
