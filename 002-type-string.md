Type slice
====

## Methods

```rust
let s = String::from("hello");
let length = s.len();

// empty String
s.clear();
```

## Loop

```rust
// Return index of first space
fn first_word(s: &String) -> usize {
    let bytes = s.as_bytes();

    for (i, &item) in bytes.iter().enumerate() {
        if item == b' ' {
            return i;
        }
    }

    s.len()
}

// Return first word literal string
fn first_word(s: &String) -> &str {
    let bytes = s.as_bytes();

    for (i, &item) in bytes.iter().enumerate() {
        if item == b' ' {
            return &s[0..i];
        }
    }

    &s[..]
}
```

## Sub-string

```rust
let s = String::from("hello");

let slice = &s[0..2]; // &str hel 
let slice = &s[..2];  // &str hel 
let slice = &s[..];   // &str hello
```
