package main

import (
	"flag"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/KelpGF/Go-Stress-Test/internal/stress"
)

func main() {
	go func() {
		log.Println("Starting server on :8080")

		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			statusCodeList := []int{200, 201, 202, 400, 401, 403, 404, 500}

			randomStatusCode := statusCodeList[rand.Intn(len(statusCodeList))]

			w.WriteHeader(randomStatusCode)
		})
		http.ListenAndServe(":8080", nil)
	}()

	time.Sleep(1 * time.Second)

	log.Println("-------------------")
	log.Println("Starting stress test")
	log.Println("-------------------")

	u := flag.String("url", "http://localhost:8080", "URL to stress test")
	r := flag.Int("requests", 100, "Number of requests to make")
	c := flag.Int("concurrency", 10, "Number of concurrent requests to make")

	result := stress.Stress(*u, *r, *c)

	log.Printf("Stress test results for %s\n", *u)
	log.Println("Total time:", result.TotalTime)
	log.Println("Total requests:", result.TotalRequests)
	for status, count := range result.TotalRequestsByStatus {
		log.Printf("Status %d: %d\n", status, count)
	}
}
