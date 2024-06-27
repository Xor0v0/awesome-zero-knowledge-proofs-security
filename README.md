# Awesome zero knowledge proofs security
[![Awesome](https://camo.githubusercontent.com/715ee701c8a9a0dbe30aac69ed79f5712a6542f5a482a3940084ce76d494a779/68747470733a2f2f617765736f6d652e72652f62616467652e737667)](https://awesome.re/) A curated list of awesome things related to learning zero knowledge proofs security

<div align=center><img src="assets/ZK Security.jpg" style="zoom:50%;"></div>

## Table of Content
- [Awesome zero knowledge proofs security](#awesome-zero-knowledge-proofs-security)
  - [Table of Content](#table-of-content)
  - [1. Introduction](#1-introduction)
  - [2. Vulnerability Classification](#2-vulnerability-classification)
    - [FrontEnd: Circuits](#frontend-circuits)
      - [Circuit Domain Specific Bugs](#circuit-domain-specific-bugs)
        - [Soundness Error](#soundness-error)
        - [Completeness Error](#completeness-error)
        - [Zero Knowledge Error](#zero-knowledge-error)
      - [Common Bugs](#common-bugs)
    - [FrontEnd: zkVM programs](#frontend-zkvm-programs)
      - [Smart Contract](#smart-contract)
    - [Back End: Proving system](#back-end-proving-system)
      - [Unstandardized Cryptographic Implementation](#unstandardized-cryptographic-implementation)
        - [Frozen Heart](#frozen-heart)
        - [Lack of Domain Seperation](#lack-of-domain-seperation)
        - [Bad Polynomial Implementation](#bad-polynomial-implementation)
        - [Missing Curve Point check](#missing-curve-point-check)
        - [Unsecure Elliptic Curve](#unsecure-elliptic-curve)
        - [Unseure Hash Function](#unseure-hash-function)
  - [3. Security Consideration](#3-security-consideration)
    - [circom](#circom)
    - [cairo](#cairo)
  - [4. Learning Resources](#4-learning-resources)
    - [Papers](#papers)
    - [Audit Reports](#audit-reports)
    - [Blogs](#blogs)
    - [zkHACK/CTF/Puzzles](#zkhackctfpuzzles)
    - [Tools](#tools)
    - [Videos](#videos)
    - [Miscellaneous](#miscellaneous)



## 1. Introduction

[Zero Knowledge Proof (ZKP)](https://github.com/matter-labs/awesome-zero-knowledge-proofs) technology is considered as a very promising infrastructure in blockchain field, even not limited to the Web3 world.

In concept, proving system (or proof system in some context) are indeed advanced cryptographic techniques as you can see in various papers. But when it comes to a ZK application, from a development perspective, it is usually divided into two parts: front-end and back-end.

In general, ZKP is a technique for proving the correct execution of programs, which has completeness, soundness, and zero knowledge property. Specifically, the front-end is these programs that can be proven, namely circuits or zkVM programs that implement business logic, while the back-end is a proving system used to generate proof for the execution of these business logic. 

As with other programming field, the primary technical risk faced by both is code bugs.

To be more precise, circuits or zkVM programs implementation comes with its own set of vulnerability classification, disjoint from the low-level cryptography bugs that may be found in the proving system.

## 2. Vulnerability Classification

The biggest difference between circuits and zkVM programs is that circuit languages are usually domain specific (DSL), and their mental models (writing constraints) are very different from traditional programming, while the programming approach of zkVM programs is more similar to traditional programming (but not exactly the same because the underlying VM is implemented as circuits, so only some circuit friendly operations can be implemented, such as hash functions [pedersen](https://iden3-docs.readthedocs.io/en/latest/iden3_repos/research/publications/zkproof-standards-workshop-2/pedersen-hash/pedersen.html#pdf-link), [poseidon](https://eprint.iacr.org/2019/458.pdf), and MiMC[https://eprint.iacr.org/2016/492.pdf]), so the learning threshold and cost are lower.

### FrontEnd: Circuits

Circuits act a very important role as namely **arithmetization** in a ZKP scheme.

There are currently many circuit DSLs, such as [Circom](https://github.com/iden3/circom), [Cairo](https://github.com/starkware-libs/cairo), [Noir](https://github.com/noir-lang/noir), [Leo](https://github.com/AleoHQ/leo), [Zokrate](https://github.com/Zokrates/ZoKrates), [Lurk](https://github.com/lurk-lab/lurk-rs), [Chiquito](https://github.com/privacy-scaling-explorations/chiquito/) etc. Ideally, provable programs written in these languages should be **well constrained**. 

The actual situation is that implementation may be **over-constrained** or **under-constrained**, even if the protocol design and implementation are improper, it may lead to **privacy leakage**. The above respectively undermines the completeness, soundness and zero knowledge property of ZKP. 

#### Circuit Domain Specific Bugs

##### Soundness Error

under-constrained

- Nondeterministic Circuits

  - (Circom) [Circom-Pairing: Missing Output Check Constraint](https://medium.com/veridise/circom-pairing-a-million-dollar-zk-bug-caught-early-c5624b278f25)

  - (Halo2-Rust) [Scroll wave1: ModGadget is underconstrained and allows incorrect MULMOD operations to be proven](https://github.com/nullity00/zk-security-reviews/blob/main/Scroll/2023-04-scroll-zkEVM-wave1-securityreview.pdf)

  - (Noir) [DoS: Recusion / AVM trace is unlimited](https://github.com/noir-lang/noir/issues/5026)


- Range Check / Mismatching Bit Lengths

    - (Circom) [Dark Forest Missing bit Length Check](https://blog.zkga.me/df-init-circuit#:~:text=Bonus%201%3A%20Range%20Proofs)
    - (Circom) [BigInt: Missing Bit Length Check](https://github.com/0xPARC/circom-ecdsa/pull/10)
    - (Halo2-Rust) [Scroll wave1: N_BYTES parameters are not checked to prevent overflow](https://github.com/nullity00/zk-security-reviews/blob/main/Scroll/2023-04-scroll-zkEVM-wave1-securityreview.pdf)

- Arithmetic operation issue

  - (Halo2-Rust) [Scroll wave1: Zero modulus will cause a panic](https://github.com/nullity00/zk-security-reviews/blob/main/Scroll/2023-04-scroll-zkEVM-wave1-securityreview.pdf)

##### Completeness Error

over-constrained
  
##### Zero Knowledge Error

- Trusted Setup Leak
  - [ZCash counterfeiting vulnerability](https://electriccoin.co/blog/zcash-counterfeiting-vulnerability-successfully-remediated/)
  - [Vitalik: How do trusted setups work?](https://vitalik.eth.limo/general/2022/03/14/trustedsetup.html)
  - [setup-ceremony](https://zkproof.org/2021/06/30/setup-ceremonies/)

- Bad Protocol Design/Implementation

    - [Dusk-Network Plonk](https://github.com/dusk-network/plonk/issues/650)

Reference

- [Consensys: Endeavors into the zero-knowledge Halo2 proving system](https://consensys.io/diligence/blog/2023/07/endeavors-into-the-zero-knowledge-halo2-proving-system/#:~:text=How%20can%20bugs%20happen%20in%20Halo2%20circuits%3F)
- [Automated Analysis of Halo2 Circuits](https://ceur-ws.org/Vol-3429/paper3.pdf)

#### Common Bugs

- Front Running

    - [RLN Front Running Issue](https://github.com/nullity00/zk-security-reviews/blob/main/RLN/VAR-RLN.pdf)

- Replay

- Double Spending


### FrontEnd: zkVM programs

The emergence of zkVM (including zkEVM) has greatly enriched the application of zk technology, and people can prove more diverse programs, such as smart contracts (starknet based on [cairo VM](https://github.com/lambdaclass/cairo-vm), blockchain based on various EVMs such as [Polygon](https://docs.polygon.technology/zkEVM/), [Scroll](https://scroll.io/blog/zkevm), [zksync](https://github.com/matter-labs/zksync-era), etc.) and general programs ([RISC Zero](https://dev.risczero.com/api/zkvm/), [SP1](https://github.com/succinctlabs/sp1), etc.) . 

Meanwhile, it also aligns with many traditional programming fields, such as reverse engineering (A CTF [puzzle](https://github.com/weikengchen/zkctf-r0-season1) by [weikeng chen](https://github.com/weikengchen/))。

The security issues in these fields are still blank and worth further exploring in the future.

#### Smart Contract

- Solidity

  Here are some good repo for its security:
  - [Solidity Security Blog](https://github.com/sigp/solidity-security-blog)
  - [not-so-smart-contract](https://github.com/crytic/not-so-smart-contracts)
  - [List of Security Vunerabilities](https://github.com/runtimeverification/verified-smart-contracts/wiki/List-of-Security-Vulnerabilities)

- Cairo

  One thing to note: when using components from the third party repo, pay attention to some default configuration that it makes. Open Zeppelin ERC20 components, for example, **the default token decimal is fixed value: 18**. Developers need to consider if this configurations are applicable in their project scenarios.

  Here are some existing audit reports for reference:
  - [Opus-2024_01-c4](https://code4rena.com/reports/2024-01-opus#h-01-neglect-of-exceptional-redistribution-amounts-in-withdraw_helper-function)
  - [lindy-labs-aura-2023_11-tob](https://solodit.xyz/issues/healthy-loans-can-be-liquidated-trailofbits-none-lindy-labs-aura-pdf)
  - [Argent-Account-2023_6-consensys](https://consensys.io/diligence/audits/2023/06/argent-account-multisig-for-starknet/)

### Back End: Proving system

The backend is a proving system that leans towards the cryptographic part, so this part involves more secure applications of cryptographic primitives. One must note: even secure primitives may introduce vulnerabilities if used incorrectly in the larger protocol or configured in an insecure manner.

#### Unstandardized Cryptographic Implementation

##### Frozen Heart

- [TrailOfBit Blog](https://blog.trailofbits.com/2022/04/13/part-1-coordinated-disclosure-of-vulnerabilities-affecting-girault-bulletproofs-and-plonk/)

##### Lack of Domain Seperation

- [Scroll zkTier: Lack of domain separation allows proof forgery](https://github.com/nullity00/zk-security-reviews/blob/main/Scroll/2023-07-scroll-zktrie-securityreview.pdf)

##### Bad Polynomial Implementation

- [Zendoo: Missing Polynomial Normalization after Arithmetic Operations](https://research.nccgroup.com/2021/11/30/public-report-zendoo-proof-verifier-cryptography-review/)
  
##### Missing Curve Point check

- [0 Bug](https://arxiv.org/pdf/2104.12255.pdf)
- [00 Bug](https://github.com/cryptosubtlety/00/blob/main/00.pdf)

##### Unsecure Elliptic Curve


##### Unseure Hash Function

- [Micro-starknet: Hash function is not second image resistant](https://github.com/paulmillr/scure-starknet/blob/main/audit/2023-09-kudelski-audit-starknet.pdf)
- [Scroll zkTier: Unchecked usize to c_int casts allow hash collisions by length misinterpretation](https://github.com/nullity00/zk-security-reviews/blob/main/Scroll/2023-04-scroll-zkEVM-wave1-securityreview.pdf)

## 3. Security Consideration

### circom

- [blockdev's slides](https://hackmd.io/@blockdev/Bk_-jRkXa#/)

### cairo

1. No payable functions
2. Name hashed storage slots
3. Upgradeability built-in
4. Separated internal/external functions
5. Cheap execution means readable algorithms
6. Immutable variables by default
7. Safe type conversions
8. Option and Result traits

**Reference**
- [starknet book](https://book.starknet.io/ch02-14-security-considerations.html)
- [cairo-the-starknet-way-to-writing-safe-code by Nethermind Security](https://medium.com/nethermind-eth/cairo-the-starknet-way-to-writing-safe-code-8169486c7132)

## 4. Learning Resources

### Papers

- [Weak Fiat-Shamir Attacks on Modern Proof Systems](https://eprint.iacr.org/2023/691.pdf)
- [On the practical CPAD security of “exact” and threshold FHE schemes and libraries](https://eprint.iacr.org/2024/116)
- [Attacks Against the INDCPA-D Security of Exact FHE Schemes](https://eprint.iacr.org/2024/127)

### Audit Reports

- [Security Reviews](https://github.com/nullity00/zk-security-reviews) of ZK Protocols by [nullity](https://github.com/nullity00). Consists of Security Reports of 50+ ZK Protocols.
- [code4rena Report](https://code4rena.com/reports)

You can directly visit the [solodit](https://solodit.xyz/) website to get some off-the-shelf audit reports.

### Blogs

- [0xPARC Blog](https://0xparc.org/blog)
- [zkHACK Blog](https://zkhack.dev/blog/)
- [NCC Group Research Blog](https://research.nccgroup.com/)
- [Zellic Blog](https://www.zellic.io/blog/)
- [zkSecurity Blog](https://www.zksecurity.xyz/blog/)
- [David Wong Blog](https://www.cryptologie.net/)
- [LambdaClass Blog](https://blog.lambdaclass.com/)
- [Nethermind Blog](https://www.nethermind.io/blogs)
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
- [cairo-damn-vulnerable](https://github.com/credence0x/cairo-damn-vulnerable-defi)
- [starknet-security-challenges.app](https://starknet-security-challenges.app/)
- [StarknetCC-CTF](https://github.com/pscott/StarknetCC-CTF)

writeups

- [2023 Ingonyama CTF WP by shuklaayush](https://hackmd.io/@shuklaayush/SkWizdyBh)
- [2023 Ingonyama CTF Official WP](https://github.com/ingonyama-zk/zkctf-2023-writeups)
- 

### Tools
| Tool | Technique | UC	| OC | CE |
| - | - | - | - | - | 
| Circomspect | SA | ✓ | ✗ | ✗ |
| ZKAP | SA	| ✓	| ✗	| ✗ |
| halo2-analyzer | SA | ✓	| ✓ |	✗ |
| Coda | FV	| ✓	| ✓	| ✓ |
| Ecne | FV | ✓ |	✗ | ✗ |
| Picus | FV | ✓ | ✗ | ✗ |
| Aleo | FV | ✓ | ✓ | ✓ |
| SnarkProbe | DA | ✓ |	✓	| ✗ |
| CIVER|FV|✓|✗|✗ |
| GNARK/Lean | FV | ✓ | ✓	| ✓ |

### Videos

### Miscellaneous

- ["Security of ZKP projects: same but different"](https://www.aumasson.jp/data/talks/zksec_zk7.pdf) by JP Aumasson @ [Taurus](https://www.taurushq.com/). Great slides outlining the different types of zk security vulnerabilities along with examples.
- [0xPARC zk-bug-tracker](https://github.com/0xPARC/zk-bug-tracker) by [0xPARC](https://0xparc.org/) and [PSE](https://pse.dev/).
- BUG Bounty platform: [code4rena](https://code4rena.com/), [Immunefi](https://immunefi.com/).
- [l2-security-framework by QuantStamp](https://github.com/quantstamp/l2-security-framework)