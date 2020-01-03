fn one(code: i8) {
    let c = match code {
        1 => "foo",
        2 => "bar",
        3|4 => "baz",
        5..=7 => "zoo",
        _ => "ooo"
    };
    println!("code = {}, c = {}", code, c);
}

fn main() {
    one(2);
    one(3);
    one(7);
    one(11);
}
