Struct
====

```rust
// define struct
// ---------------------
struct User {
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
