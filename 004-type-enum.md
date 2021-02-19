Enum
====

## Basic

```rust
enum IpAddrKind { V4, V6 }

let four = IpAddrKind::V4;
let six = IpAddrKind::V6;
```

## enum can also store value with type

```rust
enum IpAddr {
    V4(String),
    V6(String),
}

let home = IpAddr::V4(String::from("127.0.0.1"));
let loopback = IpAddr::V6(String::from("::1"));
```

## Generic

```rust
enum Option<T> {
    Some(T),
    None,
}
```

## Enum can store different types

```rust
enum Message {
    Quit,
    Move { x: i32, y: i32 },
    Write(String),
    ChangeColor(i32, i32, i32),
}
```

## Method

```rust
impl Message {
    fn call(&self) {
        // method body would be defined here
    }
}

let m = Message::Write(String::from("hello"));
m.call();
```
