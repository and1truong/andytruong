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
let mut s = String::new();

```
