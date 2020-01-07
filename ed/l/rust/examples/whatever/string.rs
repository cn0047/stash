pub fn one() {
  let s = "hello";
  println!("[1] {}", s);

  let s2:&'static str = "world";
  println!("[1] {}", s2);
  for c in s2.chars().rev() {
    println!("[1] {}", c);
  }

  if let Some(fc) = s2.chars().nth(0) {
    println!("[1] first char: {:?}", fc);
  }
}

pub fn two() {
  let mut letters = String::new();
  let mut a = 'a' as u8;
  while a <= ('z' as u8) {
    letters.push(a as char);
    letters.push_str(",");
    a += 1;
  }
  println!("[2] {}", letters);
}

pub fn three() {
  let mut s = "hey".to_string();
  s.remove(0);
  s.push_str("!");
  println!("[3] {}", s);
}

fn main() {
  one();
  two();
  three();
}
