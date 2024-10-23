# Front Running

## Introduction

Front-running attacks take on a sophisticated form in cryptocurrency. The essence of `Front-running` is to utilize the transparency and time difference of transactions on the blockchain, and to "jump the queue" through higher transaction fees (such as gas fees on Ethereum), in order to complete similar operations before others' transactions.

With front-running attacks, hackers can do evil at a higher cost. It is usually an architectural design problem.

## Case

###  1: RLN: Spammers may slash themselves

| Identifier | Severity | Location | Status |
| :--------: | :------: | :------: | :----: |
| Veridise | Medium | [withdraw.circom](https://github.com/Rate-Limiting-Nullifier/circom-rln/blob/022b690b5615d1e26874013cf216136875d8f3ab/circuits/withdraw.circom)|Acknowledged|

#### Background

Applications using RLN and economic stake to implement spam resistance may suffer
from adversaries with large collateral.
Spammers who are able to slash themselves before others (or successfully front-run others) will
be able to recover their economic stake.

Identified in the [audit report](https://github.com/nullity00/zk-security-reviews/blob/main/RLN/VAR-RLN.pdf) by [Veridise](https://veridise.com/) 

#### Description

[withdraw.circom](https://github.com/Rate-Limiting-Nullifier/circom-rln/blob/022b690b5615d1e26874013cf216136875d8f3ab/circuits/withdraw.circom) is a template that used for withdrawal/slashing and is needed to prevent frontrun while withdrawing the stake from the smart-contract/registry.

The withdraw template consists of a proof of knowledge of an `identityCommitment`'s pre-image. 

```circom
template Withdraw() {
    signal input identitySecret;
    signal input addressHash;
    signal output identityCimmitment <== Poseidon(1)([identitySecret]);
}
```

While this does prevent frontrunners from simply replaying a transaction with an address they own as the beneficiary,it does not prevent front-running from the user targeted by the slashing.

For instance, suppose Alice is intending to spam a permissionless chat application which is using RLN for spam filtering via an economic stake.

1. Alice deposits 1 coin to register.
2. Alice sends as many messages as possible until she sees someone submit a slash request
on her identityCommitment.
3. Alice front-runs the request, slashing herself (since she knows her own identitySecret)
and recovering her 1 coin.

#### The fix

One simple fix is to add documentation describing the potential issue, and recommend that applications **only give a portion of the staked amount to the slasher**. For
instance, half of the stake could be given to the slasher, and the remaining half split amongst all non-malicious protocol participants.

A second approach is to disallow self-slashing by requiring Withdraws to provide evidence
of two identity_secret_hashes, one identity which is the slasher (along with a proof of tree
membership) and one identity to be slashed. This doubles the stake required to self-slash.

