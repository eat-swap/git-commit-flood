// this file cannot be executed directly.

fn main() {
    let cnt = %d;
    let message = String::from("%s");
    for x in (0..cnt).rev() {
        println!("{}", message);
    }
}
