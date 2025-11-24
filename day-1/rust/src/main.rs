mod solution;
// Load module

use std::env;
use crate::solution::not_quite_lisp::{part_1, part_2}; // reference for calling

/**
* https://adventofcode.com/2015/day/1
*
* Part 1:
* Use counter.
* Iterate char in string
* If see '(' increment, see ')' decrement
* return counter
* Part 2:
* Extend part 1
* In cases see ')', also check floor. If negative return index of current char
* Question:
*  - What is the floor never reaches negative?
*/

fn main() {
  let args: Vec<String> = env::args().collect();
  println!("All args: {:?}", args);

  let input_file_path = &args[1];
  println!("Processing file: {:?}", args);

  let raw_content = file_processor::load_raw(input_file_path);

  // Part 1
  let part_1_result = part_1(&raw_content);
  println!("Part 1 result: {:?}", part_1_result);

  // Part 2
  let part_2_result = part_2(&raw_content);
  println!("Part 2 result: {:?}", part_2_result);
}
