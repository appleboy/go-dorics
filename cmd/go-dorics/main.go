package main

import (
	"flag"
	"log"
	"os"
	"path"
	"strings"

	"github.com/appleboy/com/file"
	"github.com/appleboy/go-dorics/pkg/analytic"
)

var (
	output string
	board  string
)

func main() {
	flag.StringVar(&output, "output", "output", "help message for flagname")
	flag.StringVar(&board, "board", "254", "board array")
	flag.Parse()

	if file.IsFile(output) {
		log.Fatal("output path is a file")
	}

	if !file.IsDir(output) {
		if err := os.MkdirAll(output, os.ModePerm); err != nil {
			log.Fatal(err)
		}
	}

	if board == "" {
		log.Fatal("missing board value")
	}

	boards := strings.Split(board, ",")
	for _, b := range boards {
		out := path.Join(output, b+".xlsx")
		if file.IsFile(out) {
			continue
		}

		board := analytic.Spider(b)
		if err := board.Excel(out); err != nil {
			log.Fatal(err)
		}
	}
}
