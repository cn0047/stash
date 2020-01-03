//![allow(dead_code)]

use std::mem;

const MY_X:u8 = 7;

static X:u8 = 9;

fn main() {
    // one();
    // two();
    // three();
    four();
}

fn four() {
    let x:f64 = 2.0;
    let r:Option<f64> = if x != 0.0 { Some(10.0/x) } else { None };
    println!("{:?}", r); // Some(5.0)

    if let Some(z) = r { println!("result = {}", z); } // result = 5
}

pub fn three() {
    let y = Box::new(5); // allocates var in heap
    println!("[3] {}; {}", y, *y);

    let z= vec![1,2,3];
    println!("[3] {:?}", z);
}

pub fn two() {
    let c:char = 'x';
    println!("[2] {}", c);
}

pub fn one() {
    let a:u8 = 1;
    let mut b:u8 = 0; // mutable
    println!("[1] a = {}; b = {}; const = {}; x = {}", a, b, MY_X, X);
    b = 2;
    println!("[1] b = {}; size = {} bytes", b, mem::size_of_val(&b));
}
