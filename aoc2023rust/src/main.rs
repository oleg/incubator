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


    let mut sum = 0;
    for line in reader.lines() {
        let mut nums: Vec<(Option<usize>, usize)> = Vec::new();
        let line = line.unwrap();

        for (i, v) in ["1", "2", "3", "4", "5", "6", "7", "8", "9"].iter().enumerate() {
            nums.push((line.find(v), i + 1));
            nums.push((line.rfind(v), i + 1));
        }
        for (i, v) in ["one", "two", "three", "four", "five", "six", "seven", "eight", "nine"].iter().enumerate() {
            nums.push((line.find(v), i + 1));
            nums.push((line.rfind(v), i + 1));
        }

        let mut pos2val = nums.iter()
            .filter(|(pos, _)| pos.is_some())
            .map(|(pos, val)| (pos.unwrap(), val))
            .collect::<Vec<_>>();
        pos2val.sort_by(|(a, _), (b, _)| a.cmp(&b));

        let d1 = pos2val.first().unwrap().1;
        let d2 = pos2val.last().unwrap().1;
        sum += d1 * 10 + d2;
    }

    println!("{}", sum);

    Ok(())
}
