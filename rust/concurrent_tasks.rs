use std::thread;
use std::time::{Duration, Instant};

fn simulated_task(task_id: u32) {
    thread::sleep(Duration::from_secs(1));
    println!("Task {} finished", task_id);
}

fn main() {
    let start = Instant::now();
    let mut handles = Vec::new();

    for task_id in 1..=5 {
        handles.push(thread::spawn(move || simulated_task(task_id)));
    }

    for handle in handles {
        handle.join().unwrap();
    }

    println!("Elapsed: {:?}", start.elapsed());
}
