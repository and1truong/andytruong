Module system
====

- Package: A package is one or more crates that provide a set of functionality. A package contains a `Cargo.toml` file that describes how to build those crates.
  - If a package contains `src/main.rs` and `src/lib.rs`, it has two crates: a library and a binary, both with the same name as the package.
  - A package can have multiple binary crates by placing files in the src/bin directory: each file will be a separate binary crate.
- Crate: A crate is a binary or library.

### Create a package

```
$ cargo new my-project

Created binary (application) `my-project` package

$ ls my-project

Cargo.toml
src

$ ls my-project/src
main.rs

# Cargo follows a convention that src/main.rs is the crate root of a binary crate with the same name as the package.
```


### Module

```rust

/*********************
crate structure
 └── front_of_house
     ├── hosting
     │   ├── add_to_waitlist
     │   └── seat_at_table
     └── serving
         ├── take_order
         ├── serve_order
         └── take_payment
*********************/
mod front_of_house {
    mod hosting {
        pub fn add_to_waitlist() {}
        fn seat_at_table() {}
    }

    mod serving {
        fn take_order() {}
        fn serve_order() {}
        fn take_payment() {}
    }
}

// bringing Paths into Scope with the use Keyword
use crate::front_of_house::hosting;
use crate::front_of_house::hosting::add_to_waitlist;

// Calling logic inside modules using path
// Error on compiling if `pub` access is not defined.
pub fn eat_at_restaurant() {
    crate::front_of_house::hosting::add_to_waitlist();  // absolute path
    front_of_house::hosting::add_to_waitlist();         // relative path
    hosting::add_to_waitlist();                         // module was imported by use keyword
    add_to_waitlist();                                  // function was imported by use keyword
}
```

Access parent module using `super` path

```rust
fn serve_order() {}

mod back_of_house {
    fn fix_incorrect_order() {
        cook_order();
        super::serve_order();
    }

    fn cook_order() {}
}
```

Define struct/enum in module

```rust
mod back_of_house {
    // struct inside module
    pub struct Breakfast {
        pub toast: String,
        seasonal_fruit: String,
    }
    
    // Enum inside module
    pub enum Appetizer { Soup, Salad }
}
```

## Using external package

Declare it to `cargo`:

```toml
[dependencies]
rand = "0.5.5"
```

Use it

```rust
use std::io;
use rand::Rng;

fn main() {
    println!("Guess the number!");

    let secret_number = rand::thread_rng().gen_range(1, 101);

    println!("The secret number is: {}", secret_number);

    println!("Please input your guess.");

    let mut guess = String::new();

    io::stdin()
        .read_line(&mut guess)
        .expect("Failed to read line");

    println!("You guessed: {}", guess);
}
```


## Nested path

```rust
// Instead long list
use std::cmp::Ordering;
use std::io;

// Use nested path
use std::{cmp::Ordering, io};
```
