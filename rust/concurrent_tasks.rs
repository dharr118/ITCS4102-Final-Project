use std::thread;
use std::time::{Duration, Instant};

fn task(number: i32) {
    println!("Task {} started", number);

    thread::sleep(Duration::from_secs(1));

    println!("Task {} finished", number);
}

fn main() {
    let start_time = Instant::now();

    let mut handles = vec![];

    for i in 1..=5 {
        handles.push(thread::spawn(move || {
            task(i);
        }));
    }

    for handle in handles {
        handle.join().unwrap();
    }

    let elapsed = start_time.elapsed();

    println!("Elapsed time: {:.2?}", elapsed);
}
