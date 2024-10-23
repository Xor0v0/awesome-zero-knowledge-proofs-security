# Unsecure Hash Function
- [Introduction](#introduction)
- [Case](#case)
  - [1. Secure-starknet: Hash function is not second image resistant](#1-secure-starknet-hash-function-is-not-second-image-resistant)
    - [Description](#description)
    - [The Fix](#the-fix)
  - [2. Gnark: MiMC Implementation Is Vulnerable To Length-Extension Attacks](#2-gnark-mimc-implementation-is-vulnerable-to-length-extension-attacks)
    - [Description](#description-1)
    - [Recommendation](#recommendation)



## Introduction

Hash function is used to map a message of arbitrary length to a fixed-length message digest. and is sort of one way function where it is easy to compute the output from a given input, but computationally infeasible to reverse the process. 

- Collision-Resistent: Attackers can not find 2 different inputs $x_1, x_2$ that $h(x_1) = h(x_2)$ holds,
- Hiding: The output $y$ does not leak any information about the input $x$.
- Puzzle-friendly: Given a random number $r$ and target value $y$, it is no efficient way to find the input $x$ such that $h(x||r)=y$, where `||` is the concatenation operator. One can only choose input randomly.

Unsecure hash function vulnerability refers to security weaknesses arising from the use of weak or improperly designed cryptographic hash functions. Common attack methods include:

1. Pre-image attack: given the hash output $y$, find the input $x$.
2. Second-image attack: given the input $x_1$, find another input $x_2$ such that $h(x_1) = h(x_2)$.
3. Collision attack: find two inputs $x_1, x_2$ such that $h(x_1) = h(x_2)$. [Birthday attack](https://en.wikipedia.org/wiki/Birthday_attack) is a bruteforce collision attack that exploits the mathematics behind the birthday problem in probability theory
4. [Length extension attack](https://en.wikipedia.org/wiki/Length_extension_attack): is a type of attack where an attacker can use $h(message_1)$ and the length of $message_1$ to calculate $h(message_1 || message_2)$ for an attacker-controlled $message_2$, without needing to know the content of $message_1$.

**The security requirements for hash functions depend on application scenario**. Generally require the use of cryptographic secure hash functions (CHF) $h(x)=y$ that is expected to have following properties:

- Pre-image resistance: Given a hash output $y$, it is hard to find the corresponding input $x$ .
- Second pre-image resistance: given the input $x_1$, it is hard to find another input $x_2$ such that $h(x_1) = h(x_2)$.
- Collision resistance: it should be difficult to find find two inputs $x_1, x_2$ such that $h(x_1) = h(x_2)$. Such a pair is called a cryptographic hash collision. This property is sometimes referred to as strong collision resistance. It requires a hash value at least twice as long as that required for pre-image resistance; otherwise, collisions may be found by a birthday attack.

Non-cryptographic hash functions are used in hash tables and to detect accidental errors; their constructions frequently provide no resistance to a deliberate attack. For example, a denial-of-service attack on hash tables is possible if the collisions are easy to find.


## Case

### 1. Secure-starknet: Hash function is not second image resistant

|Identifier|Severity|Location|Status|
|:-:|:-:|:-:|:-:|
|Kudelski Security|high | [secure-starknet/index.ts](https://github.com/paulmillr/scure-starknet/blob/07b25e9997b45a0c0d83ced2c0272306143f0660/index.ts#L216C1-L223C2) |[remediated](https://github.com/paulmillr/scure-starknet/tree/7ab944ae967efe19d1009764dce85ea9941fb7ca)| 

#### Description

The function `hashChain` is built upon the Pedersen hash function and used to hash an array of values. the hashchain is used to compute the contract storage address of a variable. However, the hashChain function does not
include the length of the data neither the starting value in the hash computation. Thus, is prone to second pre-image attack. Here is a proof of concept:

```ts
import { hashChain } from "micro-starknet"

var h1 = hashChain([1, 2, 3])
console.log(h1)

var h2 = hashChain([2, 3])
console.log(h2)

var h3 = hashChain([1, h2])
console.log(h3)

h1 === h3
```

#### The Fix

Implement the Array hashing method as define in the Hash functions [documentation](https://docs.starknet.io/architecture-and-concepts/cryptography/hash-functions/) and how it is done with computeHashOnElements function. The previous example can also be added to the tests to avoid regressions。

### 2. Gnark: MiMC Implementation Is Vulnerable To Length-Extension Attacks

| Identifier | Severity | Location | Status |
| :--------: | :------: | :------: | :----: |
| zkSecurity | Medium | [std/hash](https://github.com/Consensys/gnark/blob/f7b61b73a80ef3ce4cb0112eb272a4f16af172fb/std/hash/mimc/mimc.go#L28C1-L37C2) | [Fixed](https://github.com/Consensys/gnark/pull/1198) |

#### Description

Gnark's standard library implements the [MiMC hash function](https://eprint.iacr.org/2016/492) using the Miyaguchi–Preneel construction. The Miyaguchi–Preneel construction is a known way to turn a block cipher into a compression function, which can then be used to build a hash function.

Such an instantiation is well-studied and known to be secure against most attacks, except length-extension attacks. Length-extension attacks occur when a secret is used in the hash function to produce a keyed hash, for example, if one produces a digest as `hash(secret || public_data)` where `||` is a concatenation, someone else could pick up where the hashing was left off and produce a new digest `hash(secret || public_data || more_data)` without knowing the secret. Fundamentally, this is because the digest produced by the algorithm is the internal state of the hash function, which can be reused without issue to continue hashing.

As such, one could imagine an innocent developer producing a circuit where `data` and the digest are made public (through public inputs, for example) and then used in another protocol and circuit to produce a different keyed-hash on some related data (as explained above). For example:

```go
type Circuit1 struct {
    Key frontend.Variable `gnark:",secret"`
    Data [1]frontend.Variable `gnark:",public"`
    Expected frontend.Variable `gnark:",public"`
}

type Circuit2 struct {
    Key frontend.Variable `gnark:",secret"`
    Data [2]frontend.Variable `gnark:",public"`
    Expected frontend.Variable `gnark:",public"`
}

func (c *CircuitN) Define(api frontend.API) error {
    h, err := mimc.New(api)
    if err != nil { return err }
    h.Write(c.Key)
    h.Write(c.Data[:]...)
    res := h.Sum()
    api.AssertIsEqual(res, c.Expected)
}
```

In addition, the [MiMC website's FAQ](https://mimc.iaik.tugraz.at/pages/faq.php) also states:

> “In our original paper we propose to instantiate a hash function via plugging the MiMC permutation into a Sponge construction. The reason back then was security analysis. A mode like Miyaguchi-Preneel makes it harder, and in 2016 we did not feel confi dent in proposing this. In the meanwhile we did more security analysis, improving our understanding and confi dence, but our recommendation remains unchanged.”

Note that we could not fi nd usages of this API being used with secret data, but the API itself remains exposed to end users and does not forbid the caller to use it with secret data.

#### Recommendation

Document the API to make it clear that it should not be used with secret data.