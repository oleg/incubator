use std::error::Error;

type MyResult<T> = Result<T, Box<dyn Error>>;

pub struct Config {
    files: Vec<String>,
    lines: usize,
    bytes: Option<usize>,
}