use reqwest::Error;
use serde::Deserialize;
use tokio::time::sleep;
use std::time::Duration;
use std::time::Instant;
use serde_json;

#[derive(Deserialize, Debug)]
struct Response {
    url: String,
    args: serde_json::Value,
}

async fn fetch_data(seconds: u64) -> Result<Response, Error> {
    let request_url = format!("https://httpbin.org/delay/{}", seconds);
    let response  = reqwest::get(&request_url).await?;
    let delayed_response:Response = response.json().await?;
    Ok(delayed_response)
}

async fn calculate_lst_login(){
    sleep(Duration::from_secs(1)).await;
    println!("Logged in 2 days ago");
}

#[tokio::main]
async fn main() -> Result<(), Error> {
    println!("Hello, world!");
    let start = Instant::now();
    let data = fetch_data(5);
    let time_since = calculate_lst_login();
    let(posts,_)= tokio::join!(data, time_since);
    let duration = start.elapsed();
    println!("Fetched {:?}", posts);
    println!("Time elapsed {:?}", duration);
    Ok(())
}
