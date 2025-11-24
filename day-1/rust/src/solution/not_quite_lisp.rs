const UP_FLOOR: char = '(';
const DOWN_FLOOR: char = ')';

pub fn part_1(directions: &str) -> Result<isize, String> {
  let mut floor = 0;

  for direction in directions.trim().chars() {
    match direction {
      UP_FLOOR => floor += 1,
      DOWN_FLOOR => floor -= 1,
      _ => return Err(format!("Invalid direction: {}", direction)),
    }
  }

  Ok(floor)
}

pub fn part_2(directions: &str) -> Result<usize, String> {
  let mut floor = 0;

  for (idx, direction) in directions.trim().chars().enumerate() {
    match direction {
      UP_FLOOR => floor += 1,
      DOWN_FLOOR => {
        floor -= 1;
        if floor < 0 {
          return Ok(idx + 1)
        }
      },
      _ => return Err(format!("Invalid direction: {}", direction)),
    }
  }

  Ok(0)
}
