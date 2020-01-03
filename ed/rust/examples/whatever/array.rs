//![allow(dead_code)]

pub fn one() {
    let a:[i32;5] = [1,2,3,4,5];
    println!("array length = {}", a.len());
    for i in 0..a.len() {
        print!("{}, ", a[i]);
    }
}

fn two() {
    let mut a = Vec::new();
    a.push(1);
    a.push(22);
    a.push(3);
    println!("{:?}", a);

    match a.get(2) {
        Some(x) => println!("got {:?} at position 2", x),
        None => println!("has nothing at position 2"),
    }

    for x in &a { println!("{}", x); }

    a.pop(); // delete last element
}

fn main() {
   // one();
   two();
}
