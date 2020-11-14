# rust

cr() {
    rustc -o /tmp/x $1 && /tmp/x
}

cr ed/l/rust/examples/whatever/hw.rs
