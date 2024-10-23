# Circuit Bugs

## Introduction

Circuits play a significant role as namely **constraints system** in a ZKP scheme. It represents the constraints relationship in the execution trace of computation.

There are currently many circuit DSLs, such as [Circom](https://github.com/iden3/circom), [Cairo](https://github.com/starkware-libs/cairo), [Noir](https://github.com/noir-lang/noir), [Leo](https://github.com/AleoHQ/leo), [Zokrate](https://github.com/Zokrates/ZoKrates), [Lurk](https://github.com/lurk-lab/lurk-rs), [Chiquito](https://github.com/privacy-scaling-explorations/chiquito/) etc. Ideally, provable programs written in these languages should be **well constrained**. 

The actual situation is that implementation may be **over-constrained** or **under-constrained**, and even lead to **privacy leakage**. The above respectively undermines the completeness, soundness and zero knowledge property of ZKP. 