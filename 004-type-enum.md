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

// Enum can store different types
// ---------------------
enum Message {
    Quit,
    Move { x: i32, y: i32 },
    Write(String),
    ChangeColor(i32, i32, i32),
}

// Method 💪
// ---------------------
impl Message {
    fn call(&self) {
        // method body would be defined here
    }
}

let m = Message::Write(String::from("hello"));
m.call();
```
