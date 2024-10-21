# Arithmetic Over/Under Flow

Arithmetic Overflow/Underflow vulnerabilities can occur due to improper handling of arithmetic boundaries and properties in finite fields. Let’s break down this concept and understand why these vulnerabilities happen in zero-knowledge systems.

In zero-knowledge proof systems, calculations are often carried out in finite fields. A finite field has a limited number of elements, typically defined by a large prime number  `p` . All arithmetic operations (addition, multiplication, subtraction, etc.) within these systems are performed modulo  `p`.

In normal operations, if a result exceeds the finite field’s maximum value (i.e., is greater than or equal to p), the operation result is reduced by p, effectively wrapping around. This ensures that results always remain within the range [0, p-1]. This behavior avoids traditional overflow (in the classical computing sense), as the result is always within the bounds of the field.

Even though finite fields seem to avoid typical integer overflows, **Arithmetic Overflow/Underflow** can **still occur without panic** in ZK systems:

- Input Alias Attack
  
  In ZK systems, constraints are used to define valid inputs and operations, which is on a finite field by default. But if inputs are not constrained correctly or if operations are allowed to exceed the modulus without adequate checks, overflow vulnerabilities can arise. This is known as the input alias attack.
  

## Case 1: 