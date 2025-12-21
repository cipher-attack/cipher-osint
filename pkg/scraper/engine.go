package scraper

import (
	"fmt"
	"sync"
	"time"
)

// Task - find work
type Task struct {
	Source string
	Target string
}

// Result 
type Result struct {
	Data string
	Err  error
}

func RunScraper(target string) string {
	sources := []string{"Google_Dork", "Tor_Crawler", "LinkedIn_DNA", "Pastebin_Leaked"}
	tasks := make(chan Task, len(sources))
	results := make(chan string, len(sources))
	var wg sync.WaitGroup

	// Worker Pool
	for w := 1; w <= 3; w++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for task := range tasks {
				// real dorks
				time.Sleep(500 * time.Millisecond) // Simulated speed
				results <- fmt.Sprintf("[+] %s: Found data for %s", task.Source, task.Target)
			}
		}(w)
	}

	// send to channel works
	for _, s := range sources {
		tasks <- Task{Source: s, Target: target}
	}
	close(tasks)

	// finshed work and report the result
	go func() {
		wg.Wait()
		close(results)
	}()

	var aggregate string
	for res := range results {
		aggregate += res + "\n"
	}
	return aggregate
}