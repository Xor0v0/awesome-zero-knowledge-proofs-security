# Awesome zero knowledge proofs security
[![Awesome](https://camo.githubusercontent.com/715ee701c8a9a0dbe30aac69ed79f5712a6542f5a482a3940084ce76d494a779/68747470733a2f2f617765736f6d652e72652f62616467652e737667)](https://awesome.re/) A curated list of awesome things related to learning zero knowledge proofs security

## Introduction

Zero Knowledge Proof (ZKP) technology is considered as a very promising infrastructure in blockchain field, even not limited to the Web3 world.

In concept, proving system (or proof system in some context) are indeed advanced cryptographic techniques as you can see in various papers. But when it comes to a ZK application, from a development perspective, it is usually divided into two parts: front-end and back-end.

In general, ZKP is a technique for proving the correct execution of programs, which has integrity, soundness, and zero knowledge property. Specifically, the front-end is these programs that can be proven, namely circuits or zkVM programs that implement business logic, while the back-end is a proving system used to generate proof for the execution of these business logic. 

As with other programming field, the primary technical risk faced by both is code bugs.

To be more precise, circuits or zkVM programs implementation comes with its own set of vulnerability classification, disjoint from the low-level cryptography bugs that may be found in the proving system.

## Vulnerability Classification

### Front End

The biggest difference between circuits and zkVM programs is that circuit languages are usually domain specific (DSL), and their mental models (writing constraints) are very different from traditional programming, while the programming approach of zkVN programs is more similar to traditional programming (but not exactly the same because the underlying VM is implemented as circuits, so only some circuit friendly operations can be implemented, such as hash functions that only support pedersen and poseidon), so the learning threshold and cost are lower.

The following will introduce domain specific bugs and common bugs separately.

#### circuits

There are currently many circuit DSLs, such as Circom, Cairo, Noir, Leo, Zokrate, Lurk, etc. Ideally, provable programs written in these languages should be well constrained. The actual situation is that implementation may be **over-constrained** or **under-constrained**, even if the protocol design and implementation are improper, it may lead to **privacy leakage**. The above respectively undermines the completeness, soundness and zero knowledge property of zkp. 

##### Domain Specific Bugs

**Circom**

under-constrained

- Nondeterministic Circuits

    case 1: 

    case 2: 

- Mismatching Bit Lengths

- Unused Public Inputs Optimized Out

over-constrained
    
- over

    case 1: 

privacy leakage

- Trusted Setup Leak

    case 1:

**Rust(Halo2)**

WIP

**Cairo**

WIP

**Noir**

WIP

**Leo**

WIP

**Zokrates**

WIP

##### Common Bugs

- Arithmetic Over/Under Flows

- Bad Data Validation

- Access Control

#### zkVM programs

The emergence of zkVM (including zkEVM) has greatly enriched the application of zk technology, and people can prove more diverse programs, such as smart contracts (starknet based on cairo VM, blockchain based on various EVMs such as Polygon, Scroll, zksync, etc.) and general programs (RISC Zero, SP1, etc.) . Meanwhile, it also aligns with many traditional programming fields, such as reverse engineering。

The security issues in these fields are still blank and worth further exploring in the future.

##### Cairo Bugs

More involving smart contracts and Defi security.

### Back End

The backend is a proving system that leans towards the cryptographic part, so this part involves more secure applications of cryptographic primitives. One must note: even secure primitives may introduce vulnerabilities if used incorrectly in the larger protocol or configured in an insecure manner.

#### Unstandardized Cryptographic Implementation

##### Frozen Heart

##### Bad Polynomial Implementation

##### Unseure Hash Function

##### Unsecure Elliptic Curve

## Learning Resources

**Papers**

**Audit Reports**

**Blogs**

**Vedios**