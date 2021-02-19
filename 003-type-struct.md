Struct
====

```rust
// define struct
struct User {
    username: String,
    email: String,
    sign_in_count: u64,
    active: bool,
}

// create new struct instance
let someUser = User {
    email: String::from("someone@example.com"),
    username: String::from("someusername123"),
    active: true,
    sign_in_count: 1,
};

// if mutable, whole object is mutable -- not just certain fields
let mut someOtherUser = User {
    email: String::from("someone@example.com"),
    username: String::from("someusername123"),
    active: true,
    sign_in_count: 1,
};

someOtherUser.email = String::from("anotheremail@example.com");
```
