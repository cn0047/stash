//![allow(dead_code)]

struct Point<T> {
  y: T
}

fn one() {
  let p1:Point<i32> = Point{y:1};
  let p2:Point<u64> = Point{y:2};
  println!("[1] p1 = {}; p2 = {}", p1.y, p2.y);
}

fn main() {
  one()
}
