pub fn one() {
  let x = 3;
  let y = 4;
  let s = sum(x, y);
  println!("[1] s = {:?}; v1 = {}, v2 = {}", s, s.0, s.1);

  let (a,b) = s;
  println!("[1] a = {}, b = {}", a, b);
  let c = (a, b);
  println!("[1] c = {:?}", c);
}

fn sum(x:i32, y:i32) -> (i32, i32) {
  (x+y, x*y)
}

fn main() {
  one();
}
