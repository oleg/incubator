// use clap::Parser;
// use catr::Args;
// use clap_builder::Parser;

use std::fs::File;
use std::io::{self, BufRead, BufReader};

use anyhow::Result;
use clap::Parser;

fn main() {
    if let Err(e) = run(Args::parse()) {
        eprintln!("{e}");
        std::process::exit(1);
    }
}

/// Rust version of `cat`
#[derive(Debug, Parser)]
#[command(author, version, about)]
pub struct Args {
    /// Input file(s)
    #[arg(value_name = "FILE", default_value = "-")]
    pub files: Vec<String>,
    /// Number lines
    #[arg(short('n'), long("number"), conflicts_with("number_nonblank_lines"))]
    pub number_lines: bool,
    /// Number non-blank lines
    #[arg(short('b'), long("number-nonblank"))]
    pub number_nonblank_lines: bool,
}

fn run(args: Args) -> Result<()> {
    for filename in args.files {
        match open(&filename) {
            Err(err) => eprintln!("Failed to open file: {filename}: {err}"),
            Ok(r) => {
                let mut line_number = 0;
                for line in r.lines() {
                    let line = line?;
                    if args.number_lines || args.number_nonblank_lines && !line.is_empty() {
                        line_number += 1;
                        println!("{line_number:>6}\t{line}");
                    } else {
                        println!("{line}");
                    }
                }
            }
        }
    }
    Ok(())
}

fn open(filename: &str) -> Result<Box<dyn BufRead>> {
    match filename {
        "-" => Ok(Box::new(BufReader::new(io::stdin()))),
        _ => Ok(Box::new(BufReader::new(File::open(filename)?))),
    }
}
