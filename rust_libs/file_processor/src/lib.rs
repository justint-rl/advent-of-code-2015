use std::fs;
use std::fs::File;
use std::io::{self, BufRead, BufReader};

pub fn add(left: u64, right: u64) -> u64 {
    left + right
}

/**
* Load content of file as raw string
*/
pub fn load_raw(file_path: &str) -> String {
    let content = fs::read_to_string(file_path)
      .expect(&format!("Could not load file {}", file_path));
    content
}

pub fn load_multiline(file_path: &str) -> io::Result<Vec<String>> {
    let file = File::open(file_path)?;
    let reader = BufReader::new(file);
    let mut lines = Vec::new();

    for line in reader.lines() {
        lines.push(line?);
    }

    Ok(lines)
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_works() {
        let result = add(2, 2);
        assert_eq!(result, 4);
    }
}
