pub fn one(slice: &[i32]) {
  println!("[1] slice = {:?}", slice);
}

pub fn two(slice: &mut[i32]) {
  slice[0] = 9;
  println!("[2] slice = {:?}", slice);
}

fn main() {
  let data1 = [1,2,3,4,5];
  one(&data1[1..4]);
  let mut data2 = [1,2,3,4,5];
  two(&mut data2);
}
