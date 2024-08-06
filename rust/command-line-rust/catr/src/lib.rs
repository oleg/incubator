// use clap::Parser;
// use anyhow::Result;

// /// Rust version of `cat`
// #[derive(Debug, Parser)]
// #[command(author, version, about)]
// pub struct Args {
//     /// Input file(s)
//     #[arg(value_name = "FILE", default_value = "-")]
//     pub files: Vec<String>,
//     /// Number lines
//     #[arg(short('n'), long("number"), conflicts_with("number_nonblank_lines"))]
//     pub number_lines: bool,
//     /// Number non-blank lines
//     #[arg(short('b'), long("number-nonblank"))]
//     pub number_nonblank_lines: bool,
// }


// use std::error::Error;
// use std::fs::File;
// use std::io::{self, BufRead, BufReader};

// use clap::{Arg};

// type MyResult<T> = Result<T, Box<dyn Error>>;

// pub fn get_args() -> MyResult<Args> {
//     let matches = App::new("catr")
//         .version("0.1.0")
//         .author("Oleg Prozorov <oleg.entw@gmail.com>")
//         .about("Rust cat")
//         .arg(
//             Arg::with_name("files")
//                 .value_name("FILE")
//                 .help("Input file(s)")
//                 .multiple(true)
//                 .default_value("-"),
//         )
//         .arg(
//             Arg::with_name("number")
//                 .short("n")
//                 .long("number")
//                 .help("Number lines")
//                 .takes_value(false)
//                 .conflicts_with("number_nonblank"),
//         )
//         .arg(
//             Arg::with_name("number_nonblank")
//                 .short("b")
//                 .long("number-nonblank")
//                 .help("Number non-blank lines")
//                 .takes_value(false),
//         )
//         .get_matches();
//
//     Ok(Args {
//         files: matches.values_of_lossy("files").unwrap(),
//         number_lines: matches.is_present("number"),
//         number_nonblank_lines: matches.is_present("number_nonblank"),
//     })
// }
//
// pub fn run(config: Args) -> MyResult<()> {
//     for filename in config.files {
//         match open(&filename) {
//             Err(err) => eprintln!("{}: {}", filename, err),
//             Ok(f) => {
//                 let mut line_num = 0;
//                 for line_result in f.lines() {
//                     let line = line_result?;
//                     let number_non_blank = config.number_nonblank_lines && !line.is_empty();
//                     let prefix = if (config.number_lines) || (number_non_blank) {
//                         line_num += 1;
//                         format!("{:>6}\t", line_num)
//                     } else {
//                         String::new()
//                     };
//                     println!("{}{}", prefix, line)
//                 }
//             }
//         }
//     }
//     Ok(())
// }
//
// fn open(filename: &str) -> MyResult<Box<dyn BufRead>> {
//     match filename {
//         "-" => Ok(Box::new(BufReader::new(io::stdin()))),
//         _ => Ok(Box::new(BufReader::new(File::open(filename)?))),
//     }
// }
