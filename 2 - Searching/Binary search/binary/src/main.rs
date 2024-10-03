use lazy_static::lazy_static;
use std::io::{self, Write};


lazy_static! {
    static ref NUMBERS: Vec<u32> = (0..10000).collect();
}


fn main() {
    ui();
}

fn ui() {
    println!("Please enter a 1 or a 2:");
    println!("1: Play the guessing game");
    println!("2: Approximate a square root");

    let mut choice = String::new();
    io::stdin().read_line(&mut choice).expect("Failed to read line");

    match choice.trim() {
        "1" => {
            guessing_game_ui();
        }
        "2" => {
            square_root_approximation_ui();
        }
        _ => {
            println!("Invalid input, please enter 1 or 2.");
        }
    }
    ui();
}

fn guessing_game_ui() {
    print!("Guess a number between 0-10000: ");
    io::stdout().flush().expect("Failed to flush stdout");

    let mut input = String::new();
    io::stdin().read_line(&mut input).expect("Failed to read line");

    match input.trim().parse::<u32>() {
        Ok(number) => guess_number(number),
        Err(_) => println!("That's not a valid whole number."),
    }
}

fn square_root_approximation_ui() {
    println!("Enter a number to approximate root:");
    let mut number = String::new();
    io::stdin().read_line(&mut number).expect("Failed to read line");

    println!("Please enter a precision (ex. 0.001):");
    let mut precision = String::new();
    io::stdin().read_line(&mut precision).expect("Failed to read line");

    match (number.trim().parse::<f64>(), precision.trim().parse::<f64>()) {
        (Ok(n), Ok(p)) => square_root(n, p),
        _ => println!("Invalid input, please enter valid floats."),
    }

}

// Guess number using binary search
fn guess_number(number: u32) {
    let mut low: u32 = 0;
    let mut high: u32 = NUMBERS.len() as u32;
    let mut guesses: u32 = 0;

    binary_search_guess(&mut low, &mut high, number, &mut guesses);
}

fn binary_search_guess(low: &mut u32, high: &mut u32, number: u32, guesses: &mut u32) {
    let guess: u32 = (*high + *low) / 2;
    *guesses += 1;

    println!("Guess: {0}, range: {1}-{2}", guess, low, high);

    if guess == number {
        println!("\nFound {0} in {1} guesses!\n", number, guesses);
        return;
    }

    if guess < number {
        *low = guess + 1;
    }
    else {
        *high = guess - 1;
    }

    binary_search_guess(low, high, number, guesses);
}


// Approximate square root to specified precision
fn square_root(x: f64, precision: f64) {
    let mut low: f64 = 0.0;
    let mut high: f64 = x / 2.0; // Optimization: No square root of x can be greater than x / 2 if x > 1
    let mut guesses: u32 = 0;

    binary_search_root(&mut low, &mut high, x, precision, &mut guesses);
}

fn binary_search_root(low: &mut f64, high: &mut f64, x: f64, precision: f64, guesses: &mut u32) {
    let mid: f64 = (*high + *low) / 2.0;
    *guesses += 1;

    println!("Guess: {0}, range: {{{1}, {2}}}", mid, low, high);

    let guess = truncate_precise(mid * mid, precision);
    let x_trunc = truncate_precise(x, precision);

    if guess == x_trunc {
        println!("\nFound square root {0} of {1} in {2} guesses!\n", mid, x, guesses);
        return;
    }

    if guess < x_trunc {
        *low = mid;
    }
    if guess > x_trunc {
        *high = mid;
    }

    binary_search_root(low, high, x, precision, guesses);
}

fn truncate_precise(x: f64, precision: f64) -> f64 {
    let first_truncate: f64 = (x / precision).trunc();
    let second_truncate: f64 = ((first_truncate * precision) * (1.0 / precision)).round() / (1.0 / precision);
    return second_truncate;
}