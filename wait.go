package main

import (
	"fmt"
	"net"
	"net/url"
	"sync"
	"time"
)

// waitForServices tests and waits on the availability of a TCP host and port
func waitForServices(services []string, timeOut time.Duration) error {
	var depChan = make(chan struct{})
	var wg sync.WaitGroup
	wg.Add(len(services))
	go func() {
		for _, s := range services {
			go func(s string) {
				defer wg.Done()
				waitForService(s)
			}(s)
		}
		wg.Wait()
		close(depChan)
	}()

	select {
	case <-depChan: // services are ready
		return nil
	case <-time.After(timeOut):
		return fmt.Errorf("services aren't ready in %s", timeOut)
	}
}

func waitForService(s string) error {
	dialName := getDialName(s)

	for {
		_, err := net.Dial("tcp", dialName)
		if err == nil {
			break
		}
		time.Sleep(1 * time.Second)
	}
	return nil
}

func getDialName(s string) string {
	parsedURL, err := url.Parse(s)
	if err == nil && parsedURL.Host != "" {
		port := parsedURL.Port()
		if port == "" {
			port = "80"
		}
		return fmt.Sprintf("%s:%s", parsedURL.Hostname(), port)
	}

	return s
}
