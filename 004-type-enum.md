Enum
====

```rust
// define and use
// ---------------------
enum IpAddrKind { V4, V6 }

let four = IpAddrKind::V4;
let six = IpAddrKind::V6;

// enum can also store
// ---------------------
enum IpAddr {
    V4(String),
    V6(String),
}

let home = IpAddr::V4(String::from("127.0.0.1"));
let loopback = IpAddr::V6(String::from("::1"));
```
