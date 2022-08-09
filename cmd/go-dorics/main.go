package main

import (
	"log"

	"github.com/appleboy/go-dorics/pkg/analytic"
)

func main() {
	board := analytic.Spider("254")
	if err := board.Excel("output"); err != nil {
		log.Fatal(err)
	}
}
