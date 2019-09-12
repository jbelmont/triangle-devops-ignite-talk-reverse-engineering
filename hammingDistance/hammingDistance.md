# Hamming Distance

According to Wikipedia:

> In information theory, the Hamming distance between two strings of equal length is the number of positions at which the corresponding symbols are different. In other words, it measures the minimum number of substitutions required to change one string into the other, or the minimum number of errors that could have transformed one string into the other. In a more general context, the Hamming distance is one of several string metrics for measuring the edit distance between two sequences. It is named after the American mathematician Richard Hamming.

#### Hamming Examples

[Hamming Examples in Wikipedia](https://en.wikipedia.org/wiki/Hamming_distance#Examples)

The Hamming distance between:

"karolin" and "kathrin" is 3.
"karolin" and "kerstin" is 3.
1011101 and 1001001 is 2.
2173896 and 2233796 is 3.

#### Wikipedia Hamming solution written in C language

```c
int hamming_distance(unsigned x, unsigned y)
{
    int dist = 0;
    
    // Count the number of bits set
    for (unsigned val = x ^ y; val > 0; val /= 2)
    {
        // If A bit is set, so increment the count
        if (val & 1)
            dist++;
        // Clear (delete) val's lowest-order bit
    }

    // Return the number of differing bits
    return dist;
}
```