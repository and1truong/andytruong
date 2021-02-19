Module
====

- Package: A package is one or more crates that provide a set of functionality. A package contains a `Cargo.toml` file that describes how to build those crates.
- Crate: A crate is a binary or library.

### Create a package

```
$ cargo new my-project

Created binary (application) `my-project` package

$ ls my-project

Cargo.toml
src

$ ls my-project/src
main.rs # Cargo follows a convention that src/main.rs is the crate root of a binary crate with the same name as the package.
```
