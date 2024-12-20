# Bad Polynomial Implementation

## Introduction

Incorrect polynomial representation resulting from arithmetic operations may break assumptions and lead to erroneous computations or may result in denial of service attacks via Rust panics.

## Case
### 1. Zendoo: Missing Polynomial Normalization after Arithmetic Operations
| Identifier | Severity | Location | Status |
| :--------: | :------: | :------: | :----: |
| NCC Group | High | [algebra/src/fft/polynomial/dense.rs](https://github.com/HorizenOfficial/ginger-lib/blob/35cd278f24f70b498095190e360754b2c13cc4b8/algebra/src/fft/polynomial/dense.rs#L182C1-L208C2) | [Fixed](https://github.com/HorizenOfficial/ginger-lib/pull/112/commits/8e377aa3ba7e383681a5a3421b7bce67c201f8f7)|

- [Zendoo: Missing Polynomial Normalization after Arithmetic Operations](https://research.nccgroup.com/2021/11/30/public-report-zendoo-proof-verifier-cryptography-review/)

#### Description

The file fft/polynomial/dense.rs provides an implementation of dense polynomials to be used for FFTs. These polynomials are represented by vectors in which each entry corresponds to a coefficient. These coefficients are elements of a finite field, and as such, the sum of two coefficients may take any value in the range 0, . . . , p − 1, where p is the order of the prime field.

When adding two polynomials of the same degree using the function add(), trailing coefficients that sum to zero are not trimmed. This contradicts an underlying assumption on the shape of polynomial representations, namely that the coefficient of the leading term is non-zero.

As an example, summing the polynomials $3+2x+x^2$ and $1+(p-1)x^2$ (using the function add() provided below for reference) represented by the vectors [3, 2, 1] and [1, 0, p - 1] will result in the vector [4, 2, 0], namely the trailing position is equal to zero.

```Rust
fn add(self, other: &'a DensePolynomial<F>) -> DensePolynomial<F> {
    if self.is_zero() {
        other.clone()
    } else if other.is_zero() {
        self.clone()
    } else {
        if self.degree() >= other.degree() {
            let mut result = self.clone();
            for (a, b) in result.coeffs.iter_mut().zip(&other.coeffs) {
                *a += b
            }
            result
        } else {
            let mut result = other.clone();
            for (a, b) in result.coeffs.iter_mut().zip(&self.coeffs) {
                *a += b
            }
            // If the leading coefficient ends up being zero, pop it off.
            while result.coeffs.last().unwrap().is_zero() {
                result.coeffs.pop();
            }
            result
        }
    }
}
```

Interestingly, note that the else-clause in the add() function above does perform this trimming.

While this failure to trim leading zero coefficients is technically not inconsistent with the current polynomial representation (and should not lead to incorrect results), the implementation assumes that all trailing zeros have been trimmed from polynomials.

As a result, functions like degree() (provided below) will panic on unexpected inputs.

```Rust
/// Returns the degree of the polynomial.
pub fn degree(&self) -> usize {
    if self.is_zero() {
        0
    } else {
        assert!(self.coeffs.last().map_or(false, |coeff| !coeff.is_zero()));
        self.coeffs.len() - 1
    }
}
```

This oversight with regards to the trimming of zero coefficients applies to function add_assign(), sub() and sub_assign().

#### The Fix

Consider performing the “trimming” step of removing trailing zero coefficients from polynomials in all cases after arithmetic operations. Additionally, consider writing unit tests to catch such potential edge cases.

Zendoo Developers introduced a function named truncate_leading_zeros() which removes the leading zero coefficients of a polynomial. This function is now called prior to returning the result of the arithmetic operations add(), add_assign(), sub(), and sub_assign(). As such, this finding has been marked as “Fixed”.
