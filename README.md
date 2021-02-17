Learn rust
====

## Hello world with number guess game

```rust
use rand::Rng;
use std::cmp::Ordering;
use std::io;

fn main() {
    println!("Guess the number!");

    let secret_number = rand::thread_rng().gen_range(1, 101);

    loop {
        println!("Please input your guess.");

        let mut guess = String::new();

        io::stdin()
            .read_line(&mut guess)
            .expect("Failed to read line");

        let guess: u32 = match guess.trim().parse() {
            Ok(num) => num,
            Err(_) => continue,
        };

        println!("You guessed: {}", guess);

        match guess.cmp(&secret_number) {
            Ordering::Less => println!("Too small!"),
            Ordering::Greater => println!("Too big!"),
            Ordering::Equal => {
                println!("You win!");
                break;
            }
        }
    }
}
```

## Ownership

- At any given time, you can have either one mutable reference or any number of immutable references.
- References must always be valid.

### Mut reference


```rust
// Error: Users of an immutable reference don’t expect the values to suddenly change out from under them!
{
    let mut s = String::from("hello");
    let r1 = &s; // no problem
    let r2 = &s; // no problem
    let r3 = &mut s; // BIG PROBLEM

    println!("{}, {}, and {}", r1, r2, r3);
}

// OK
{
    let mut s = String::from("hello");

    let r1 = &s; // no problem
    let r2 = &s; // no problem
    println!("{} and {}", r1, r2);
    // r1 and r2 are no longer used after this point

    let r3 = &mut s; // no problem
    println!("{}", r3);
}
```

### No dangling

```rust
fn main() {
    let reference_to_nothing = dangle();
}

// Error: s goes out of scope, and is dropped. Its memory goes away. Danger!
fn dangle() -> &String {
    let s = String::from("hello");

    &s
}

// OK
fn no_dangle() -> String {
    let s = String::from("hello");

    s
}
```
