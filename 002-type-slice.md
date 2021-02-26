Type slice
====

Learn more:

- https://doc.rust-lang.org/std/primitive.slice.html

## Overview

```
// define
let a = [1, 2, 3, 4, 5];

// access part
let slice = &a[1..3];

```

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

// Return first word string literal
fn first_word(s: &String) -> &str {
    let bytes = s.as_bytes();

    for (i, &item) in bytes.iter().enumerate() {
        if item == b' ' {
            return &s[0..i];
        }
    }

    &s[..]
}

// String can be used as &str
// str is slice
// Usage: 
//   let s = String::from("hello world");
//   let _ = first_word(&s[..]); // OK: call with String
//   let ls = "hello world";     // string literal
//   let _ = first_word(ls);     // OK: string
//   let _ = first_word(ls[..]); // OK: &str is slice
fn first_word(s: &str) -> &str {
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
