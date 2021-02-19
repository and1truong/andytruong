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

```rust
// struct with no field!
struct nothing ();
```

## Examples

```rust
#[derive(Debug)]
struct Rectangle {
    width: u32,
    height: u32,
}

fn area(rectangle: &Rectangle) -> u32 {
    rectangle.width * rectangle.height
}

fn main() {
    let rect1 = Rectangle {
        width: 30,
        height: 50,
    };

    // can print per we used directive `derive(Debug)` for the struct
    println!("rect1 is {:?}", rect1);
    println!("pretty print {:#}", rect1);
}
```

## Methods

```rust
struct Rectangle {
    width: u32,
    height: u32,
}

impl Rectangle {
    fn area(&self) -> u32 {
        self.width * self.height
    }

    fn can_hold(&self, other: &Rectangle) -> bool {
        self.width > other.width && self.height > other.height
    }
    
    // Associated Functions
    // usage: let my_square = Rectangle::square(3);
    fn square(size: u32) -> Rectangle {
        Rectangle {
            width: size,
            height: size,
        }
    }
}
```
