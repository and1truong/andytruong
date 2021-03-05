Collection › Vector
====

Learn more:

- https://doc.rust-lang.org/std/vec/struct.Vec.html

```rust
fn main() {
    // define vector
    // explicit: let my_vector: Vec<i32> = Vec::new();
    let mut my_vector = vec![1, 2, 3, 4];

    // access elements
    let third = &my_vector[2]; // should make reference instead of getting ownership
    println!("{} {}", third, my_vector[2]); // 3 3

    // index my not point to valid result
    match my_vector.get(490) {
        Some(number) => println("{}", number),
        None => println("Element not found"),
    }

    my_vector.push(490);
    my_vector.remove(2); // remove element 3

    for number in my_vector.iter() {
        println("{}", number);
    }
}
```

Store multiple type in Vector? We can: Enum. Using an enum plus a match expression means that Rust will ensure at compile time that every possible case is handled,

```rust
enum SpreadsheetCell {
    Int(i32),
    Float(f64),
    Text(String),
}

let row = vec![
    SpreadsheetCell::Int(3),
    SpreadsheetCell::Text(String::from("blue")),
    SpreadsheetCell::Float(10.12),
];
```
