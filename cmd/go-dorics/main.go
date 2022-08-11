package main

import (
	"context"
	"flag"
	"log"
	"os"
	"path"
	"strings"

	"github.com/appleboy/go-dorics/pkg/analytic"

	"github.com/appleboy/com/file"
	"github.com/golang-queue/queue"
)

var (
	output    string
	board     string
	queueSize int
)

func main() {
	flag.StringVar(&output, "output", "output", "help message for flagname")
	flag.StringVar(&board, "board", "254", "board array")
	flag.IntVar(&queueSize, "size", 2, "concurrent queue size")
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

	// initial queue pool
	q := queue.NewPool(queueSize)
	// shutdown the service and notify all the worker
	// wait all jobs are complete.
	defer q.Release()

	boards := strings.Split(board, ",")
	for _, b := range boards {
		go func(b string) {
			if err := q.QueueTask(func(ctx context.Context) error {
				out := path.Join(output, b+".xlsx")
				if file.IsFile(out) {
					return nil
				}

				board := analytic.Spider(b)
				return board.Excel(out)
			}); err != nil {
				panic(err)
			}
		}(b)
	}
	q.Wait()
}
