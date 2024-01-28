use std::env;
use std::fs::File;
use std::io::{BufRead, BufReader};

#[derive(Debug)]
struct Cubes {
    red: i32,
    green: i32,
    blue: i32,
}

#[derive(Debug)]
struct Game {
    id: i32,
    sets: Vec<Cubes>,
}
/*
Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green
 */

//12 red
//13 green
//14 blue

pub fn run() -> Result<(), Box<dyn std::error::Error>> {
    let args: Vec<_> = env::args().collect();
    let file = File::open(&args[1])?;
    let reader = BufReader::new(file);

    for line in reader.lines() {
        let game_and_sets = line?.split(":");
        let game = (&game_and_sets).next().or_else(|| None).unwrap();
        let sets = (&game_and_sets).next().or_else(|| None).unwrap();

        println!(">{}", game);
        println!(">>{}", sets);
    }

    Ok(())
}