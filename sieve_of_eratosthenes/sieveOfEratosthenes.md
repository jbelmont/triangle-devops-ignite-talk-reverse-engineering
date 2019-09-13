# Sieve of Eratosthenes

## Description of Sieve of Eratosthenes

According to Wikipedia:

> In mathematics, the Sieve of Eratosthenes is a simple, ancient algorithm for finding all prime numbers up to any given limit.

#### Overview of the Sieve of Eratosthenes

[Overview of Sieve in wikipedia](https://en.wikipedia.org/wiki/Sieve_of_Eratosthenes)

To find all the prime numbers less than or equal to a given integer n by Eratosthenes' method:

1. Create a list of consecutive integers from 2 through n: (2, 3, 4, ..., n).
    1. Initially, let p equal 2, the smallest prime number.

2. Enumerate the multiples of p by counting in increments of p from 2p to n, and mark them in the list (these will be 2p, 3p, 4p, ...; the p itself should not be marked).

3. Find the first number greater than p in the list that is not marked. If there was no such number, stop. Otherwise, let p now equal this new number (which is the next prime), and repeat from step 3.

4. When the algorithm terminates, the numbers remaining not marked in the list are all the primes below n.

#### Pseudocode of Seive of Eratosthenes

```
Input: an integer n > 1.
 
 Let A be an array of Boolean values, indexed by integers 2 to n,
 initially all set to true.
 
 for i = 2, 3, 4, ..., not exceeding √n:
   if A[i] is true:
     for j = i2, i2+i, i2+2i, i2+3i, ..., not exceeding n:
       A[j] := false.
 
 Output: all i such that A[i] is true.
 ```