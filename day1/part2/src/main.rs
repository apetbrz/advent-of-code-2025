use std::fs::File;
use std::io::{BufReader, prelude::*};

fn main() {
    let file = File::open("sequence.txt").expect("please input a 'sequence.txt'");
    let steps = BufReader::new(file).lines();

    let mut dial: u16 = 50;
    let mut zeroes: u16 = 0;

    steps.for_each(|line| {
        let line = line.expect("broken sequence file");
        let (direction, amount) = line.split_at(1);
        let amount: u16 = amount.parse().expect("invalid amount");
        let clicks: u16;

        (dial, clicks) = match direction {
            "L" => left,
            "R" => right,
            _ => panic!("invalid direction: {direction}"),
        }(dial, amount);

        zeroes += clicks;
    });

    println!("amount of zeroes: {zeroes}");
}

fn left(mut dial: u16, amount: u16) -> (u16, u16) {
    let mut clicks = 0;
    println!("dial at {dial}, {amount} down");
    if dial == 0 {
        dial += 100
    }
    while dial < amount {
        dial += 100;
        clicks += 1;
        println!("    click");
    }
    dial -= amount;
    if dial == 0 {
        clicks += 1
    }
    println!(" -> dial to {dial}");
    (dial, clicks)
}

fn right(mut dial: u16, amount: u16) -> (u16, u16) {
    let mut clicks = 0;
    println!("dial at {dial}, {amount} up");
    dial += amount;
    while dial >= 100 {
        dial -= 100;
        clicks += 1;
        println!("    click");
    }
    println!(" -> dial to {dial}");
    (dial, clicks)
}
