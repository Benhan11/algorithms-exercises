use lazy_static::lazy_static;


lazy_static! {
    static ref NUMBERS: Vec<u32> = (0..10000).collect();
}


fn main() {
    guess_number(7834);
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
