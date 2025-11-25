use crate::solution::model::dimension::Dimension;

pub fn part_1(contents: &Vec<String>) -> Result<usize, String> {
  let dimensions = parse_dimensions(contents);
  if dimensions.len() == 0 {
    return Ok(0);
  }
  let mut total_surface_area = 0;
  for dim in dimensions {
    total_surface_area += dim.total_wrapping_paper()
  }
  Ok(total_surface_area)
}

pub fn part_2(contents: &Vec<String>) -> Result<usize, String> {
  let dimensions = parse_dimensions(contents);
  if dimensions.len() == 0 {
    return Ok(0);
  }
  let mut total_ribbon = 0;
  for dim in dimensions {
    total_ribbon += dim.total_ribbon();
  }
  Ok(total_ribbon)
}

fn parse_dimensions(contents: &Vec<String>) -> Vec<Dimension> {
  let mut dimensions = Vec::new();
  for dim in contents {
    dimensions.push(Dimension::from_string(&dim))
  }
  dimensions
}
