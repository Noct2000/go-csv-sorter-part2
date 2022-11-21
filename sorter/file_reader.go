package sorter

import (
	"fmt"
	"sync"
)

func (p *Pipeline) fileReadingStage(fnames chan string, n int) (allLines chan string) {
	lines := make([]chan string, n)
	allLines = make(chan string)
	for i := 0; i < n; i++ {
		lines[i] = make(chan string)
		readFiles(fnames, lines[i])
	}
	wg := &sync.WaitGroup{}
	for i := range lines {
		wg.Add(1)
		go func(ch chan string) {
			defer wg.Done()
			for line := range ch {
				allLines <- line
			}
		}(lines[i])
	}
	go func() {
		wg.Wait()
		close(allLines)
	}()
	return allLines
}

func readFiles(fnames, lines chan string) {
	go func() {
		defer close(lines)
		for fname := range fnames {
			for i := 0; i < 10; i++ {
				lines <- fmt.Sprintf("%q content# %d\n", fname, i+1)
			}
		}
	}()
}