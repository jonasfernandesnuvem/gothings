package main

import (
	"fmt"
	"time"
)

var diadasemana string

func init() {
	diadasemana = time.Now().Weekday().String()
}

func main() {
	fmt.Println("Hoje e dia ", diadasemana)
}
