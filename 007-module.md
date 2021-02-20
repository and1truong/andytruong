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
        fn add_to_waitlist() {}
        fn seat_at_table() {}
    }

    mod serving {
        fn take_order() {}
        fn serve_order() {}
        fn take_payment() {}
    }
}
```
