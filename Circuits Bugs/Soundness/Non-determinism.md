# Non-determinism

- [Introduction](#introduction)
- [Case](#case)
  - [1. Scroll-zkevm-circuit: is\_zero circuit](#1-scroll-zkevm-circuit-is_zero-circuit)

## Introduction

Non-determinism means there are multiple ways to create a valid proof for a certain outcome. This can be very bad in certain cases, especially nullifiers. On the contrary, in scenarios where the requirement for uniqueness of proof is not high (i.e. zkVM), some constraints can be omitted.

## Case 

### 1. Scroll-zkevm-circuit: is_zero circuit
