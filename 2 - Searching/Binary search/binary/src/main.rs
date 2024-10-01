use lazy_static::lazy_static;
use std::io::{self, Write};


lazy_static! {
    static ref NUMBERS: Vec<u32> = (0..10000).collect();
}


fn main() {
    ui();
}

fn ui() {
    println!("Welcome to the guessing game!");
    print!("Guess a number between 0-10000: ");
    io::stdout().flush().expect("Failed to flush stdout");

    let mut input = String::new();

    match io::stdin().read_line(&mut input) {
        Ok(_) => {
            let number: u32 = input.trim().parse().expect("Please enter a valid number...");
            println!();
            guess_number(number);
        }
        Err(err) => {
            println!("Invalid number, err:\n{}", err);
        }
    }
    ui();
}

// Guess number using binary search
fn guess_number(number: u32) {
    let mut low: u32 = 0;
    let mut high: u32 = NUMBERS.len() as u32;
    let mut guesses: u32 = 0;

    binary_search(&mut low, &mut high, number, &mut guesses);
}

fn binary_search(low: &mut u32, high: &mut u32, number: u32, guesses: &mut u32) {
    let guess: u32 = (*high + *low) / 2;
    *guesses += 1;

    println!("Guess: {0}, range: {1}-{2}", guess, low, high);

    if guess == number {
        println!("\nFound {0} in {1} guesses!", number, guesses);
        return;
    }

    if guess < number {
        *low = guess + 1;
    }
    else {
        *high = guess - 1;
    }

    binary_search(low, high, number, guesses);
}
