extern crate sieve_of_eratosthenes;

fn main() {
    let primes = sieve_of_eratosthenes::sieve(5);
    print!("{:?}", primes);
}