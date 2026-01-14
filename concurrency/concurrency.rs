use std::thread;
use std::time::Instant;

fn main() {
    let start = Instant::now();
    // Rustの標準スレッドで10万はOS設定(ulimit)でクラッシュするため
    // 比較用に1/10の「1万」で実行します
    let count = 10000; 
    let mut handles = vec![];

    println!("Starting to spawn {} threads...", count);

    for _ in 0..count {
        let handle = thread::spawn(|| {
            // 何もしない
        });
        handles.push(handle);
    }

    for handle in handles {
        handle.join().unwrap();
    }

    println!("Rust: Spawning {} threads took {:?}", count, start.elapsed());
}
