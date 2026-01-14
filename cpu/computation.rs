use std::time::Instant;

fn main() {
    let size = 100000000;
    let mut data = vec![0; size];
    for i in 0..size { data[i] = i; }

    let start = Instant::now();
    let sum: usize = data.iter().sum();

    println!("Rust: Sum {:?} in {:?}", sum, start.elapsed());
}
