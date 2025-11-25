mod solution; // Load module

use std::env;
use crate::solution::wrapping_paper::{part_1, part_2}; // reference for calling

/**
* https://adventofcode.com/2015/day/2
*
* Part 1:
* Iterate through each string and parse into Dimension
* Dimension will split by x and cast char to integer
* For each Dimension compute the surface area for each surface and keep track of the smallest surface area
* Sum up the surface areas for each side and add the smallest surface area
* Keep a running total surface area and add the sum from current dimension to running total surface
* Return total surface area
* Complexity:
*   - Runtime: O(n) w. n is the number of dimensions in input file
*   - Space: O(n) w. n is the number of dimensions to track. The algo is O(1) as we are tracking sum of surface area
* Part 2:
* Question:
*   - Can any of the dimension be negative?
*   - Will the dimension always be in format of lxwxh?
*/

fn main() {
  let args: Vec<String> = env::args().collect();
  println!("All args: {:?}", args);

  if args.len() < 2 {
    panic!("Missing file path: cargo run <file_path>");
  }

  let input_file_path = &args[1];
  println!("Processing file: {:?}", args);

  let multiline_content = file_processor::load_multiline(input_file_path)
    .expect(&format!("Failed to load multiline file {}", input_file_path));

  // Part 1
  let part_1_result = part_1(&multiline_content);
  println!("Part 1 result: {:?}", part_1_result);

  // Part 2
  let part_2_result = part_2(&multiline_content);
  println!("Part 2 result: {:?}", part_2_result);
}
