use crate::solution::model::dimension::Dimension;

pub fn part_1(contents: &Vec<String>) -> Result<usize, String> {
  let mut dimensions = Vec::new();
  for dim in contents {
    dimensions.push(Dimension::from_string(&dim))
  }
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
  Ok(0)
}
