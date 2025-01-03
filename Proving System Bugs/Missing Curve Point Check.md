# Missing Curve Point Check

## Introduction
Missing curve point check is one of the most common cryptographic bugs where an elliptic curve point used in a protocol is invalid (not belong to the correct curve or subgroup).

## Cases

There are a lot of cases we can study from.

### 1. Golang: Elliptic curve crash with invalid point

| Identifier | Severity | Location | Status |
| :--------: | :------: | :------: | :----: |
|Guido Vranken| Critical|[crypto/elliptic](https://github.com/golang/go/blob/master/src/crypto/elliptic/elliptic.go)|[Remidated](https://github.com/golang/go/issues/50974)|

`crypto/elliptic` implements elliptic curves over prime-order finite fields, meaning coordinates are elements of the field composed of the integers from zero to P-1 for some prime P. **Integers below zero (negative) or equal to or higher than P (overflowing) are as meaningless as coordinates as float64(1.5) or string("hi") would be**.

Regrettably, the `crypto/elliptic` API allows passing negative or overflowing big.Ints as coordinates.

Similarly, an arbitrary pair of field elements (x, y) is not a valid point if they don't satisfy the curve equation, and doing group operations (like a scalar multiplication) on them is meaningless. (Worse, it's called an invalid curve attack and can leak bits of the scalar that an attacker can plug into the evergreen Chinese Remainder Theorem.)

The (implicit in Go 1.17, made explicit in Go 1.18 by [refactor](https://github.com/golang/go/commit/30b5d6385e91ab557978c0024a9eb90e656623b7#diff-7d57224906f14d89dd65c1deadadd84620eb396477346633ed81a57c100a20e2)) contract of crypto/elliptic is: you shall pass input points though IsOnCurve (Unmarshal will do it for you); if it returns true, then all group operations will work and return valid points. If IsOnCurve returns false, behavior is undefined.

> Footnote: (0, 0) is not a valid point, Unmarshal will not return it, IsOnCurve will return false, and Marshal will behave incorrectly, but it's what group operations return if the output is the point at infinity, which can't be represented with two abelian coordinates. Again, exposing the coordinates in the API is a mistake. See this [Go 1.15 change](https://github.com/golang/go/commit/320e4adc4bd153cb0cb7e31e186fb3b4564fd0a7#diff-7d57224906f14d89dd65c1deadadd84620eb396477346633ed81a57c100a20e2L310) and the linked issue for more info.

#### The Fix

- Fix IsOnCurve as a security issue (CVE-2022-23806) to always **return false** for negative and overflowing values, because that's the function behaving inconsistently and the only one for which these inputs shouldn't be undefined behavior.
- Marshal and the group operations should **panic** for invalid field elements, and more broadly for invalid points. Hopefully there will be a better and secure API in the furure. 

### 2. 0 Bug in pairing-based BLS signature

| Identifier | Severity | Location | Status |
| :--------: | :------: | :------: | :----: |
| Nguyen Thoi Minh Quan | High | |

#### Description

Let's introduce aggregate signatures first. The basic goal of signature aggregation is the following: let's assume we have $n$ users and each have secret key $x_i$ and public key $X_i$. Each user signs its message $m_i$ as $\sigma_i=\text{Sign}(x_i, m_i)$. Now in verification, instead of verify $n$ signatures $\sigma_1, \dots, \sigma_n$ indivially, we want to verify a single aggregate signature $\sigma$ which somehow combines all the signature together. This not only reduces CPU cycles but also saves bandwidth in transferring signatures over the network.

The attacks are against non-repudiation security property, which isn’t captured in the standard "existential unforgeability" definition. As we’ll show below, non-repudiation property is far more important for aggregate signatures than for single signatures.

##### BLS protocol

BLS signature has a attractive security perporty, which is signature aggregation.

- [0 Bug](https://arxiv.org/pdf/2104.12255.pdf)
- [00 Bug](https://github.com/cryptosubtlety/00/blob/main/00.pdf)