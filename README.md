Learn rust
====

## Ownership

- At any given time, you can have either one mutable reference or any number of immutable references.
- References must always be valid.


### No dangling

```rust
fn main() {
    let reference_to_nothing = dangle();
}

// Error: s goes out of scope, and is dropped. Its memory goes away. Danger!
fn dangle() -> &String {
    let s = String::from("hello");

    &s
}

// OK
fn no_dangle() -> String {
    let s = String::from("hello");

    s
}
```
