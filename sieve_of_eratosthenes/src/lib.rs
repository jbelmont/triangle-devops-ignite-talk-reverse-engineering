pub fn sieve(max_number: usize) -> Vec<usize> {
    let mut numbers = vec![true; max_number+1];
    numbers[0] = false;
    numbers[1] = false;

    let mut primes = Vec::new();
    // iterate from 2 <= max_number
    for num in 2..max_number+1 {
        if numbers[num] {
            primes.push(num);
        }

        let mut next_number = num * num;
        while next_number <= max_number {
            numbers[next_number] = false;
            next_number += num;
        }
    }
    
    return primes
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_should_find_primes_less_than_or_equal_to_n() {
        let expected = vec![2, 3, 5];
        let actual = sieve(5);
        assert_eq!(actual, expected);

        assert_eq!(
            sieve(10),
            vec![2, 3, 5, 7]
        );
    }
}
