use std::env;
use std::fs::File;
use std::io::{BufRead, BufReader};

fn main() {
    if let Err(err) = run() {
        eprintln!("{}", err);
    }
}

fn run() -> Result<(), Box<dyn std::error::Error>> {
    let args: Vec<_> = env::args().collect();
    let filename = &args[1];

    let file = File::open(filename)?;
    let reader = BufReader::new(file);
    let first_line = reader.lines().next().unwrap()?;

    let floor: i32 = first_line.chars()
        .map(|c| {
            match c {
                '(' => 1,
                ')' => -1,
                _ => 0,
            }
        })
        .sum();

    println!("day1, part1: {}", floor);


    let mut floor2 = 0;
    for (position, char) in first_line.chars().enumerate() {
        match char {
            '(' => floor2 += 1,
            ')' => floor2 -= 1,
            _ => (),
        }
        if floor2 == -1 {
            println!("day1, part2: {}", position + 1);
            break;
        }
    }
    Ok(())
}
