//![allow(dead_code)]

enum Color {
    Red,
    // Green,
    // Blue
    RGBColor(u8,u8,u8)
}

fn one() {
    let mut c:Color = Color::Red;
    match c {
        Color::Red => println!("red"),
        // Color::Green => println!("green"),
        // Color::Blue => println!("blue")
        Color::RGBColor(0,0,0) => println!("black"),
        Color::RGBColor(r,g,b) => println!("rgb({},{},{})", r, g, b)
    };

    c = Color::RGBColor(2,0,9);
    match c {
        Color::Red => println!("red"),
        // Color::Green => println!("green"),
        // Color::Blue => println!("blue")
        Color::RGBColor(0,0,0) => println!("black"),
        Color::RGBColor(r,g,b) => println!("rgb({},{},{})", r, g, b)
    };
}

fn main() {
    one();
}
