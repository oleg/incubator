use io::Error;
use io::ErrorKind::Other;
use std::{env, io};
use std::fs::File;
use std::io::{BufRead, BufReader};

fn main() {
    if let Err(err) = run() {
        eprintln!("{}", err);
    }
}

#[derive(Debug)]
struct Matcher {
    s: String,
}

impl Matcher {
    fn new() -> Matcher {
        Matcher { s: String::new() }
    }
    fn push(&mut self, c: char) {
        self.s.push(c);
    }
    fn digit(&self) -> Option<i32> {
        match self.s.as_str() {
            "one" | "1" => Some(1),
            "two" | "2" => Some(2),
            "three" | "3" => Some(3),
            "four" | "4" => Some(4),
            "five" | "5" => Some(5),
            "six" | "6" => Some(6),
            "seven" | "7" => Some(7),
            "eight" | "8" => Some(8),
            "nine" | "9" => Some(9),
            _ => None,
        }
    }
    fn push_and_match_digit(&mut self, c: char) -> Option<i32> {
        self.push(c);
        self.digit()
    }
}

fn run() -> Result<(), Box<dyn std::error::Error>> {
    let args: Vec<_> = env::args().collect();
    let file = File::open(&args[1])?;
    let reader = BufReader::new(file);

    let mut sum = 0;
    for line in reader.lines() {
        let mut matchers: Vec<Matcher> = Vec::new();
        let mut first: Option<i32> = None;
        let mut last: Option<i32> = None;
        for c in line?.chars() {
            matchers.push(Matcher::new());
            for m in &mut matchers {
                if let Some(d) = m.push_and_match_digit(c) {
                    if let None = first {
                        first = Some(d);
                    }
                    last = Some(d);
                }
            }
        }
        match (first, last) {
            (Some(f), Some(l)) => {
                sum += f * 10 + l;
            }
            _ => Err(Box::new(Error::new(Other, "failed to get first and last digit")))?,
        }
    }

    println!("{}", sum);
    Ok(())
}
