use reqwest;
use tokio;

#[tokio::main]
async fn main() -> Result<(), reqwest::Error> {
    let response = reqwest::get("https://api.github.com").await?;

    println!("Status Code: {}", response.status());

    Ok(())
}
