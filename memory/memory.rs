use std::time::Instant;

#[allow(dead_code)] 
struct Payload {
    data: [u8; 1024],
}

impl Payload {
    fn new() -> Self {
        Self { data: [0; 1024] }
    }
}

fn main() {
    let count = 10_000_000;
    let start = Instant::now();

    for _ in 0..count {
        let _p = Box::new(Payload::new()); 
        // Scope end: Immediate drop (free)
    }

    println!("Rust: Allocations took {:?}", start.elapsed());
}
