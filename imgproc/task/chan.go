package task

import (
	"fmt"
	"path"
	"path/filepath"

	"training.go/imgproc/filter"
)

type ChanTask struct {
	dirCtx
	Filter   filter.Filter
	Poolsize int
}

func NewChanTask(srcDir, dstDir string, filter filter.Filter, poolsize int) Tasker {
	return &ChanTask{
		Filter: filter,
		dirCtx: dirCtx{
			SrcDir: srcDir,
			DstDir: dstDir,
			files:  buildFileList(srcDir),
		},
		Poolsize: poolsize,
	}
}

type jobReq struct {
	src string
	dst string
}

func (c *ChanTask) Process() error {
	size := len(c.files)
	jobs := make(chan jobReq, size)
	results := make(chan string, size)

	// init workers
	for w := 1; w <= c.Poolsize; w++ {
		go worker(w, c, jobs, results)
	}

	// start jobs
	for _, f := range c.files {
		filename := filepath.Base(f)
		dst := path.Join(c.DstDir, filename)
		jobs <- jobReq{
			src: f,
			dst: dst,
		}
	}
	close(jobs)

	for range c.files {
		fmt.Println(<-results)
	}
	return nil
}

func worker(id int, chanTask *ChanTask, jobs <-chan jobReq, results chan<- string) {
	workers := make(map[int]bool)
	done := make(map[int]int)

	workers[id] = true
	for j := range jobs {
		fmt.Printf("worker %d, started job %v\n", id, j)
		err := chanTask.Filter.Process(j.src, j.dst)
		fmt.Printf("worker %d, finished job %v\n", id, j)
		if _, ok := done[id]; ok {
			done[id]++
		} else {
			done[id] = 1
		}
		if err != nil {
			fmt.Printf("erreur id=%d de traitement du fichier %v\n", id, j)
		}
		results <- j.dst
	}
	fmt.Printf("\n--worker n°%d has done %d/%d works--\n\n", id, done[id], len(chanTask.files))
}
