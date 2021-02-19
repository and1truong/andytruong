Struct
====

## Overview

```rust
// define struct
// ---------------------
struct User {
    // String is intented, we can't use `name: &string` because ownership, 
    // object must own all ownership of its fields
    name: String,
    email: String,
    sign_in_count: u64,
    active: bool,
}

// if mutable, whole object is mutable -- not just certain fields
// ---------------------
let mut some_user = User {
    email: String::from("someone@example.com"),
    name: String::from("someusername123"),
    active: true,
    sign_in_count: 1,
};

some_user.email = String::from("anotheremail@example.com");

// create from other instance
let other_user = User {
  email: String::from("other_user@example.com"),
  name: String::from("other_user"),
  ..some_user
}

// Builder
// ---------------------
fn build_user(name: String, email: String) -> User {
    return User {
        name, // shorthande for `name: name`
        email,
        active: true,
        sign_in_count: 1,
    }
}
```


## Tuple struct

```rust
struct Color(i32, i32, i32);
struct Point(i32, i32, i32);

let black = Color(0, 0, 0);
let origin = Point(0, 0, 0);
```

## Unit-like struct

```
// struct with no field!
struct nothing ();
```
