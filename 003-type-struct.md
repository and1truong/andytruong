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
let mut someOtherUser = User {
    email: String::from("someone@example.com"),
    name: String::from("someusername123"),
    active: true,
    sign_in_count: 1,
};

someOtherUser.email = String::from("anotheremail@example.com");

// Builder
// ---------------------
fn build_user(name: String, email: String) -> User {
    return User {
        name,
        email,
        active: true,
        sign_in_count: 1,
    }
}
```
