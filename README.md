# Awesome zero knowledge proofs security
[![Awesome](https://camo.githubusercontent.com/715ee701c8a9a0dbe30aac69ed79f5712a6542f5a482a3940084ce76d494a779/68747470733a2f2f617765736f6d652e72652f62616467652e737667)](https://awesome.re/) A curated list of awesome things related to learning zero knowledge proofs security

## Table of Content
- [Awesome zero knowledge proofs security](#awesome-zero-knowledge-proofs-security)
  - [Table of Content](#table-of-content)
  - [Introduction](#introduction)
  - [Vulnerability Classification](#vulnerability-classification)
    - [Front End](#front-end)
      - [circuits](#circuits)
        - [Domain Specific Bugs](#domain-specific-bugs)
          - [Circom](#circom)
          - [Rust(Halo2)](#rusthalo2)
          - [Cairo](#cairo)
          - [Noir](#noir)
          - [Leo](#leo)
          - [Zokrates](#zokrates)
        - [Common Bugs](#common-bugs)
          - [Architetural Design Flaw](#architetural-design-flaw)
          - [Business Logic Error](#business-logic-error)
      - [zkVM programs](#zkvm-programs)
        - [Cairo Starknet Contract](#cairo-starknet-contract)
    - [Back End](#back-end)
      - [Unstandardized Cryptographic Implementation](#unstandardized-cryptographic-implementation)
        - [Frozen Heart](#frozen-heart)
        - [Bad Polynomial Implementation](#bad-polynomial-implementation)
        - [Missing Curve Point check](#missing-curve-point-check)
        - [Unsecure Elliptic Curve](#unsecure-elliptic-curve)
        - [Unseure Hash Function](#unseure-hash-function)
  - [Learning Resources](#learning-resources)
    - [Papers](#papers)
    - [Audit Reports](#audit-reports)
    - [Blogs](#blogs)
    - [zkHACK/CTF/Puzzles](#zkhackctfpuzzles)
    - [Vedios](#vedios)
    - [Miscellaneous](#miscellaneous)



## Introduction

[Zero Knowledge Proof (ZKP)](https://github.com/matter-labs/awesome-zero-knowledge-proofs) technology is considered as a very promising infrastructure in blockchain field, even not limited to the Web3 world.

In concept, proving system (or proof system in some context) are indeed advanced cryptographic techniques as you can see in various papers. But when it comes to a ZK application, from a development perspective, it is usually divided into two parts: front-end and back-end.

In general, ZKP is a technique for proving the correct execution of programs, which has completeness, soundness, and zero knowledge property. Specifically, the front-end is these programs that can be proven, namely circuits or zkVM programs that implement business logic, while the back-end is a proving system used to generate proof for the execution of these business logic. 

As with other programming field, the primary technical risk faced by both is code bugs.

To be more precise, circuits or zkVM programs implementation comes with its own set of vulnerability classification, disjoint from the low-level cryptography bugs that may be found in the proving system.

## Vulnerability Classification

### Front End

The biggest difference between circuits and zkVM programs is that circuit languages are usually domain specific (DSL), and their mental models (writing constraints) are very different from traditional programming, while the programming approach of zkVM programs is more similar to traditional programming (but not exactly the same because the underlying VM is implemented as circuits, so only some circuit friendly operations can be implemented, such as hash functions that only support [pedersen](https://iden3-docs.readthedocs.io/en/latest/iden3_repos/research/publications/zkproof-standards-workshop-2/pedersen-hash/pedersen.html#pdf-link) and [poseidon](https://eprint.iacr.org/2019/458.pdf)), so the learning threshold and cost are lower.

The following will introduce domain specific bugs and common bugs separately.

#### circuits

There are currently many circuit DSLs, such as Circom, Cairo, Noir, Leo, Zokrate, Lurk, etc. Ideally, provable programs written in these languages should be well constrained. The actual situation is that implementation may be **over-constrained** or **under-constrained**, even if the protocol design and implementation are improper, it may lead to **privacy leakage**. The above respectively undermines the completeness, soundness and zero knowledge property of zkp. 

##### Domain Specific Bugs

###### Circom

- under-constrained

    - Nondeterministic Circuits

        Case:

        - [Circom-Pairing: Missing Output Check Constraint](https://medium.com/veridise/circom-pairing-a-million-dollar-zk-bug-caught-early-c5624b278f25)

    - Mismatching Bit Lengths

        Case:

        - [Dark Forest Missing bit Length Check](https://blog.zkga.me/df-init-circuit#:~:text=Bonus%201%3A%20Range%20Proofs)
        - [BigInt: Missing Bit Length Check](https://github.com/0xPARC/circom-ecdsa/pull/10)

    - Unused Public Inputs Optimized Out

- over-constrained
    
    case 1: 

- privacy leakage

    - Trusted Setup Leak
  
        case 1: 

    - Bad Protocol Design/Implementation

        [Dusk-Network Plonk](https://github.com/dusk-network/plonk/issues/650)

###### Rust(Halo2)

- WIP

###### Cairo

- WIP

###### Noir

- WIP

###### Leo

- WIP

###### Zokrates

- WIP

##### Common Bugs

###### Architetural Design Flaw

- Front Running

    Case:
    - [RLN Front Running Problem](https://github.com/nullity00/zk-security-reviews/blob/main/RLN/VAR-RLN.pdf)

- Replay

- Double Spending

- Privacy Leakage

###### Business Logic Error

- Access Control

- Data Validation
  
- Denial of Service

Arithmetic Over/Under Flows


#### zkVM programs

The emergence of zkVM (including zkEVM) has greatly enriched the application of zk technology, and people can prove more diverse programs, such as smart contracts (starknet based on [cairo VM](https://github.com/lambdaclass/cairo-vm), blockchain based on various EVMs such as [Polygon](https://docs.polygon.technology/zkEVM/), [Scroll](https://scroll.io/blog/zkevm), [zksync](https://github.com/matter-labs/zksync-era), etc.) and general programs ([RISC Zero](https://dev.risczero.com/api/zkvm/), [SP1](https://github.com/succinctlabs/sp1), etc.) . Meanwhile, it also aligns with many traditional programming fields, such as reverse engineering (A ctf [puzzle](https://github.com/weikengchen/zkctf-r0-season1) by [weikeng chen](https://github.com/weikengchen/))。

The security issues in these fields are still blank and worth further exploring in the future.

##### Cairo Starknet Contract

**Security consideration**

**Tools**: [Cairo Fuzzer](https://github.com/FuzzingLabs/cairo-fuzzer), [Caracal](https://github.com/crytic/caracal), [Thoth](https://github.com/FuzzingLabs/thoth).

### Back End

The backend is a proving system that leans towards the cryptographic part, so this part involves more secure applications of cryptographic primitives. One must note: even secure primitives may introduce vulnerabilities if used incorrectly in the larger protocol or configured in an insecure manner.

#### Unstandardized Cryptographic Implementation

##### Frozen Heart
  
    [TrailOfBit Blog](https://blog.trailofbits.com/2022/04/13/part-1-coordinated-disclosure-of-vulnerabilities-affecting-girault-bulletproofs-and-plonk/)

##### Bad Polynomial Implementation

    [Zendoo: Missing Polynomial Normalization after Arithmetic Operations](https://research.nccgroup.com/2021/11/30/public-report-zendoo-proof-verifier-cryptography-review/)
    
##### Missing Curve Point check

    Case:
    - [0 Bug](https://arxiv.org/pdf/2104.12255.pdf)
    - [00 Bug](https://github.com/cryptosubtlety/00/blob/main/00.pdf)

##### Unsecure Elliptic Curve


##### Unseure Hash Function

## Learning Resources

### Papers

- [Algebraic Cryptanalysis of the HADES Design
Strategy: Application to Poseidon and Poseidon2](https://eprint.iacr.org/2023/537.pdf)

### Audit Reports

- [Security Reviews](https://github.com/nullity00/zk-security-reviews) of ZK Protocols by [nullity](https://github.com/nullity00). Consists of Security Reports of 20+ ZK Protocols.

### Blogs

- [0xPARC Blog](https://0xparc.org/blog)
- [zkHACK Blog](https://zkhack.dev/blog/)
- [LambdaClass Blog](https://blog.lambdaclass.com/)
- [NCC Group Research Blog](https://research.nccgroup.com/)
- [Nethermind Blog](https://www.nethermind.io/blogs)
- [zkSecurity Blog](https://www.zksecurity.xyz/blog/)
- [Ingonyama Blog](https://www.ingonyama.com/blog)
- [Open Zeppelin Blog](https://blog.openzeppelin.com/)
- [samczsun' Blog](https://samczsun.com/)

### zkHACK/CTF/Puzzles

- [zkHACKs](https://zkhack.dev/)
- [Paradigm CTF](https://ctf.paradigm.xyz/)
- [Paradigm CTF Infrastructure](https://github.com/paradigmxyz/paradigm-ctf-infrastructure)
- [Open Zeppelin CTF](https://ctf.openzeppelin.com/)
- [Ingonyama CTF](https://ctf.ingonyama.com/)
- [RareSkill ZK Puzzles](https://github.com/RareSkills/zero-knowledge-puzzles/tree/main)

writeups

- [2023 Ingonyama CTF WP by shuklaayush](https://hackmd.io/@shuklaayush/SkWizdyBh)
- [2023 Ingonyama CTF Official WP](https://github.com/ingonyama-zk/zkctf-2023-writeups)
- 

### Vedios

### Miscellaneous

- ["Security of ZKP projects: same but different"](https://www.aumasson.jp/data/talks/zksec_zk7.pdf) by JP Aumasson @ [Taurus](https://www.taurushq.com/). Great slides outlining the different types of zk security vulnerabilities along with examples.
- [0xPARC zk-bug-tracker](https://github.com/0xPARC/zk-bug-tracker) by [0xPARC](https://0xparc.org/) and [PSE](https://pse.dev/).
  