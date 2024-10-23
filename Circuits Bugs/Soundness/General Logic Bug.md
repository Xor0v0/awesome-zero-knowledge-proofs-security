# General Logic Bug

- [Introduction](#introduction)
- [Case](#case)
  - [1: Circom-Pairing: Missing Output Check Constraint](#1-circom-pairing-missing-output-check-constraint)
    - [Background](#background)
    - [Description](#description)
    - [The Fix](#the-fix)
  - [2. No distuiguishing leaf nodes and branch nodes](#2-no-distuiguishing-leaf-nodes-and-branch-nodes)


## Introduction

 This type of problem is usually caused by a lack of necessary conditions or limitations in the system's business logic, resulting in system behavior deviating from exxample, no distinguishing leaf nodes and branch nodes in Merkle trees.

## Case

### 1: Circom-Pairing: Missing Output Check Constraint

| Identifier | Severity | Location | Status |
| :--------: | :------: | :------: | :----: |
| Veridise | High | [bls_signature.circom](https://github.com/yi-sun/circom-pairing/blob/380e5430fe3e4effbd62fdb5abb7ea93af686f97/circuits/bls_signature.circom#L76C1-L94C7) |[Fixed](https://github.com/yi-sun/circom-pairing/pull/21/commits/c686f0011f8d18e0c11bd87e0a109e9478eb9e61)|

#### Background

The circom-pairing library is a Circom implementation of elliptic curve pairing for the widely-adopted BLS12-381 curve. The library enables zkp systems to verify signature over BLS12-381 curve.

In the context of circom-pairing, for instance, we can prove the validity of a BLS12–381 signature for a given public key. The second artifact, called the verifier, can verify the validity of a proof without running the actual computation. Because proof generation time depends on the number of constraints, ZK systems aim to reduce the constraints the verifier needs to check.

The first challenge is how represent the number of bits needed to describe pair of integers (x, y) that lie on the curve. However the current zkp system can only support 254-bit integers. To bypass this limitation, the authors of circom-pairing had to create a library for integers of arbitrary bit length (big integers for short). Each library function operates over a big integer that consists of k registers containing an n-bit integer. 

The other challenge is the heavy constraints of verifing signatures over BLS12–381. A crucial design decision in circom-pairing that drastically reduced the number of generated constraints was to omit data validation in several core templates of the library. For instance, several core circuits in circom-pairing assume their inputs are properly formatted big integers that fall in a certain range. Hence, users of such circuits must also use the BigLessThan template to perform appropriate data validation for the core circuits. As we shall see next, this decision can, unfortunately, backfire.

#### Description

The bug originates in the circuits called `CoreVerifyPubkeyG1`, which, as the name suggests, has the crucial task of verifying a signature for a pub key. Before performing the signature verification, `CoreVerifyPubkeyG1` must ensure that all of its inputs conform to the assumptions made by the core circom-pairing templates. The most important assumption is that all curve coordinates belong to the finite field defined by the curve's prime number `q`.

These checks happen in the following code snippet, which uses ten instances of BigLessThan.

```circom
// Inputs:
//   - pubkey as element of E(Fq)
//   - hash represents two field elements in Fp2, in practice hash = hash_to_field(msg,2).
//   - signature, as element of E2(Fq2)
// Assume signature is not point at infinity
template CoreVerifyPubkeyG1(n, k){
  ...
  var q[50] = get_BLS12_381_prime(n, k);
  
  component lt[10];
  // check all len k input arrays are correctly formatted bigints < q (BigLessThan calls Num2Bits)
  for(var i=0; i<10; i++){
    lt[i] = BigLessThan(n, k);
    for(var idx=0; idx<k; idx++)
      lt[i].b[idx] <== q[idx];
  }
  for(var idx=0; idx<k; idx++){
    lt[0].a[idx] <== pubkey[0][idx];
    lt[1].a[idx] <== pubkey[1][idx];
    ... // Initializing parameters for rest of the inputs
  }
  ...
}
```

Let's dive into `BigLessThan` template:

```circom
/*
Inputs:
  - BigInts a, b
Output:
  - out = (a < b) ? 1 : 0
*/
template BigLessThan(n, k){
  signal input a[k];
  signal input b[k];
  signal output out;
  ...
}
```

The `out` signal is set to one if the integer represented by array a is smaller than the one represented by b and zero otherwise. But the `CoreVerifyPubkeyG1` does not check it so that it can accept inputs that are larger than q, the curve's prime, or even are not properly formatted !!

#### The Fix

Constrain every output signal of components in lt to be equal to one.

```circom
template CoreVerifyPubkeyG1(n, k){
  ...
  var q[50] = get_BLS12_381_prime(n, k);
  
  component lt[10];
  ... // Loops same as before
 
  var r = 0;
  for(var i=0; i<10; i++){
    r += lt[i].out;
  }
  r === 10;
  ...
}
```

### 2. No distuiguishing leaf nodes and branch nodes






