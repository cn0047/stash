fn one(x:i32) {
  println!("[1] {:?}", x);
}

fn mul(a:i32, b:i32) -> i32 {
  a*b
}

fn inc(x: &mut i32) {
  *x += 1;
}

fn hello() {
  println!("hello");
}

fn two() {
  let c = hello;
  c();

  let closur = |x:i32| -> i32 { x+1 };
  let a = 6;
  println!("[2] {:?}", closur(a));
}

pub fn is_even(x: u32) -> bool {
  x%2 == 0
}

fn three() {
  let l = 500;
  let mut s = 0;
  for i in 0.. {
    let isq = i*i;
    if isq > l {
      break;
    } else if is_even(isq) {
      s += isq;
    }
  }
  println!("[3] sum = {}", s);
}

fn four() {
  let l = 500;
  let s = (0..)
    .map(|x| x*x)
    .take_while(|&x| x < l)
    .filter(|x| is_even(*x))
    .fold(0, |sum,x| sum+x);
  println!("[4] sum = {}", s);
}

fn main() {
  one(1);

  let mut z = 1;
  inc(&mut z);
  println!("z = {}", z);

  let p = mul(2, 5);
  println!("p = {}", p);

  two();
  three();
  four();
}
