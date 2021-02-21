Common Collections
====

Rust’s standard library includes a number of very useful data structures called collections. Most other data types represent one specific value, but collections can contain multiple values. Unlike the built-in array and tuple types, the data these collections point to is stored on the heap, which means the amount of data does not need to be known at compile time and can grow or shrink as the program runs. Each kind of collection has different capabilities and costs, and choosing an appropriate one for your current situation is a skill you’ll develop over time.

## Vector

```rust
fn main() {
    // define vector
    // explicit: let my_vector: Vec<i32> = Vec::new();
    let mut my_vector = vec![1, 2, 3, 4];

    // simple reading
    println!("{}", my_vector[2]); // 3

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

## String

…

## Hash Map

…
