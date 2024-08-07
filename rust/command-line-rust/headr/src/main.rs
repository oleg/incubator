use std::fs::File;
use std::io::{BufRead, BufReader, Read};

use anyhow::Result;
use clap::Parser;

/// Rust version of `head`
#[derive(Parser, Debug)]
#[command(author, version, about)]
struct Args {
    /// Input file(s)
    #[arg(
        value_name = "FILE",
        default_value = "-"
    )]
    files: Vec<String>,

    /// Number of lines
    #[arg(
        short('n'),
        long,
        default_value = "10",
        value_name = "LINES",
        value_parser = clap::value_parser ! (u64).range(1..)
    )]
    lines: u64,

    /// Number of bytes
    #[arg(
        short('c'),
        long,
        conflicts_with("lines"),
        value_name = "BYTES",
        value_parser = clap::value_parser ! (u64).range(1..)
    )]
    bytes: Option<u64>,
}

fn main() {
    if let Err(e) = run(Args::parse()) {
        eprintln!("{e}");
        std::process::exit(1);
    }
}

fn run(args: Args) -> Result<()> {
    let multi_file = args.files.len() > 1;
    let mut first = true;
    for filename in args.files {
        if multi_file {
            if first {
                first = false
            } else {
                println!();
            }
            println!("==> {filename} <==");
        }
        match open(&filename) {
            Ok(mut r) => {
                if let Some(num_bytes) = args.bytes {
                    let num_bytes = num_bytes as usize;
                    let mut buffer = vec![0; num_bytes];
                    let bytes_read = r.read(&mut buffer)?;
                    print!("{}", String::from_utf8_lossy(&buffer[..bytes_read]));
                } else {
                    let mut line = String::new();
                    for _ in 0..args.lines {
                        let len = r.read_line(&mut line)?;
                        if len == 0 {
                            break;
                        }
                        print!("{line}");
                        line.clear();
                    }
                }
            }
            Err(e) => eprintln!("Failed to open file: {filename}: {e}"),
        }
    }
    Ok(())
}

fn open(filename: &str) -> Result<Box<dyn BufRead>> {
    match filename {
        "-" => Ok(Box::new(BufReader::new(std::io::stdin()))),
        _ => Ok(Box::new(BufReader::new(File::open(filename)?))),
    }
}


