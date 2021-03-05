String
====

- `String` is not `str`, usually seen `&str`
  - String -> `String`
  - String slice -> `str`
  - both `String` and `str` are UTF-8 encoded
- Other string types: `OsString`, `OsStr`, `CString`, `CStr`.
  - `xxxString` -> owned
  - `xxxStr` -> borrowed.


```rust
// create new string
// ---------------------
let mut s = String::new();

// from string slice
let data = "initial contents";
let s = data.to_string();
let s = "initial contents".to_string(); // short form

// static method
let s = String::from("initial contents");

// UTF-8 encoded
let hello = String::from("السلام عليكم");
let hello = String::from("Dobrý den");
let hello = String::from("Hello");
let hello = String::from("שָׁלוֹם");
let hello = String::from("नमस्ते");
let hello = String::from("こんにちは");
let hello = String::from("안녕하세요");
let hello = String::from("你好");
let hello = String::from("Olá");
let hello = String::from("Здравствуйте");
let hello = String::from("Hola");
```
