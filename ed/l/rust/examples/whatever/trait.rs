trait Animal {
  fn create(name: &'static str) -> Self;

  fn name(&self) -> &'static str;

  fn talk(&self) {
    println!("{} cannot talk", self.name());
  }
}

struct Human {
  name: &'static str
}

impl Animal for Human {
  fn create(name: &'static str) -> Human {
    Human{name: name}
  }

  fn name(&self) -> &'static str {
    self.name
  }

  fn talk(&self) {
    println!("{} says hello", self.name());
  }
}

struct Cat {
  name: &'static str
}

impl Animal for Cat {
  fn create(name: &'static str) -> Cat {
    Cat{name: name}
  }

  fn name(&self) -> &'static str {
    self.name
  }
}

fn one() {
  let h = Human{name: "John"};
  h.talk();

  let c = Cat::create("Tom");
  c.talk();

  let h2:Human = Animal::create("Jerry");
  h2.talk();
}

trait Summable<T> {
  fn sum(&self) -> T;
}

impl Summable<i32> for Vec<i32> {
  fn sum(&self) -> i32 {
    let mut res:i32 = 0;
    for x in self  { res += x; }
    return res;
  }
}

fn two() {
  let v = vec![1,2,3];
  println!("[2] sum = {}", v.sum());
}

fn main() {
  one();
  two();
}
