package main

import (
	"flag"
	"fmt"
	"time"

	"training.go/imgproc/filter"
	"training.go/imgproc/task"
)

func main() {
	srcDir := flag.String("src", "", "Input directory")
	dstDir := flag.String("dst", "", "Output directory")
	filterType := flag.String("filter", "grayscale", "grayscale/blur")
	taskType := flag.String("task", "waitgrp", "waitgrp/channel")
	sizePoolsize := flag.Int("poolsize", 4, "number of goroutines on channel")

	flag.Parse()
	var f filter.Filter
	switch *filterType {
	case "grayscale":
		f = &filter.Grayscale{}
	case "blur":
		f = &filter.Blur{}
	}

	var t task.Tasker
	switch *taskType {
	case "waitgrp":
		t = task.NewWaitGrpTask(*srcDir, *dstDir, f)
	case "channel":
		t = task.NewChanTask(*srcDir, *dstDir, f, *sizePoolsize)
	}

	start := time.Now()

	t.Process()

	elapsed := time.Since(start)
	fmt.Printf("Image processing took %s\n", elapsed)
}
