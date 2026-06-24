// This example is intended to be used inside a Cargo project with reqwest installed.
// Cargo.toml dependency example:
// reqwest = { version = "0.12", features = ["blocking"] }

fn main() -> Result<(), Box<dyn std::error::Error>> {
    let response = reqwest::blocking::get("https://api.github.com")?;
    println!("Status code: {}", response.status());
    println!("Content type: {:?}", response.headers().get("content-type"));
    Ok(())
}
