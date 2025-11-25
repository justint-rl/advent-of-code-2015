use std::arch::aarch64::vorn_u8;
use std::cmp;

pub struct Dimension {
  pub length: usize,
  pub width: usize,
  pub height: usize,
}

const DIMENSION_DELIMITER: char = 'x';

impl Dimension {
  pub fn new (length: usize, width: usize, height: usize) -> Dimension {
    Dimension { length, width, height }
  }

  pub fn from_string(s: &str) -> Dimension {
    let dims = s.split(DIMENSION_DELIMITER).collect::<Vec<&str>>();

    let l = *dims.get(0).expect(&format!("Dimension {} missing length", s));
    let w = *dims.get(1).expect(&format!("Dimension {} missing width", s));
    let h = *dims.get(2).expect(&format!("Dimension {} missing height", s));

    let l_num = l.parse::<usize>().expect(&format!("Invalid length {}", l));
    let w_num = w.parse::<usize>().expect(&format!("Invalid width {}", w));
    let h_num = h.parse::<usize>().expect(&format!("Invalid height {}", h));

    Dimension::new(l_num, w_num, h_num)
  }

  pub fn total_wrapping_paper(self) -> usize {
    let mut total_surface_area = 0;
    let smallest_surface_area = cmp::min(
      self.length * self.width,
      cmp::min(
        self.width * self.height,
        self.height * self.length,
      ),
    );
    total_surface_area += (2 * self.length * self.width);
    total_surface_area += (2 * self.height * self.width);
    total_surface_area += (2 * self.length * self.height);
    total_surface_area += smallest_surface_area;
    total_surface_area
  }

  pub fn total_ribbon(self) -> usize {
    let smallest_perimeter = cmp::min(
      2 * self.length + 2 * self.width,
      cmp::min(
        2 * self.width + 2 * self.height,
        2 * self.height + 2 * self.length,
      )
    );
    let volume = self.width * self.height * self.length;
    let total_ribbon = volume + smallest_perimeter;
    total_ribbon
  }
}
