package stress

import (
	"net/http"
	"sync"
	"time"
)

type stressResult struct {
	TotalTime             time.Duration
	TotalRequests         int
	TotalRequestsByStatus map[int]int
}

func Stress(url string, requests, concurrency int) stressResult {
	result := stressResult{
		TotalRequests:         0,
		TotalRequestsByStatus: map[int]int{},
	}

	timeStart := time.Now()

	for result.TotalRequests < requests {
		wg := sync.WaitGroup{}

		simultaneousRequests := concurrency

		if requests-result.TotalRequests < concurrency {
			simultaneousRequests = requests - result.TotalRequests
		}

		for i := 0; i < simultaneousRequests; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()

				statusCode := makeRequest(url)

				result.TotalRequestsByStatus[statusCode]++
			}()
		}
		wg.Wait()

		result.TotalRequests += simultaneousRequests
	}

	result.TotalTime = time.Since(timeStart)

	return result
}

func makeRequest(url string) int {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	return res.StatusCode
}
