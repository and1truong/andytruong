Enum
====

## Basic

```rust
enum IpAddrKind { V4, V6 }

let four = IpAddrKind::V4;
let six = IpAddrKind::V6;
```

## Enum can also store value with type

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

let some_number = Some(5);
let some_string = Some("a string");
let absent_number: Option<i32> = None;
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
## Matching

```rust
enum Coin { Penny, Nickel, Dime, Quarter(String) }

fn value_in_cents(coin: Coin) -> u8 {
    match coin {
        Coin::Penny => {
            println!("Lucky penny!");
            1
        },
        Coin::Nickel => 5,
        Coin::Dime => 10,
        Coin::Quarter(serial) => {
            println!("Serial nunber {:?}!", serial);
            25
        },
    }
}
```

Generic

```rust
fn plus_one(x: Option<i32>) -> Option<i32> {
    match x {
        Some(i) => Some(i + 1),
        // compile error if we don't handle this case
        None => None,
    }
}

let five = Some(5);
let six = plus_one(five);
let none = plus_one(None);
```

Fallback

```rust
let some_u8_value = 0u8;
match some_u8_value {
    1 => println!("one"),
    3 => println!("three"),
    5 => println!("five"),
    7 => println!("seven"),
    _ => (),
}
```
