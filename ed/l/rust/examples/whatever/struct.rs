use std::mem;

struct Point {
    x: f64,
    y: f64
}

impl Point {
  fn len(&self) -> f64 {
    self.y - self.x
  }
}

fn get_point() -> Point {
    Point{x: 5.0, y: 7.0}
}

fn main() {
    let p1:Point = get_point();
    println!("p1: x = {}; y = {}; takes {} bytes", p1.x, p1.y, mem::size_of_val(&p1));

    let p2 = Box::new(get_point());
    println!("p2: x = {}; y = {}; takes {} bytes", p2.x, p2.y, mem::size_of_val(&p2));
    println!("p2: len = {}", p2.len());
}
