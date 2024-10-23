# Mismatched Type or Length

## Introduction

Mismatched type or length is a bug developers do not perform sanitization or check on the the type or length of the inputs, which will result in unexpected behavior, even panic. This is something that pragues us always in securty especially in memory unsafe language. 

Even strongly typed languages like Rust can potentially lead to the aforementioned vulnerability. For example, in Rust, using `as` for type casting performs a truncation rather than raising a warning, which is a design decision balancing safety, performance, and developer control. Rust places the responsibility on the developer to ensure correct type conversions when using `as`. If a safer conversion is needed, Rust offers alternatives, like: `TryFrom` and `Tryinto` trait.

## Case

### 1: Scroll zkTrie - Unchecked usize to c_int casts allow hash collisions by length

| Identifier | Severity | Location | Status |
| :--------: | :------: | :------: | :----: |
| Trail of Bit |   High   | [lib.rs](https://github.com/scroll-tech/zktrie/blob/90179c19281670f41c54bd80ab01e4d64c860521/src/lib.rs#L134C1-L138C6) | [fixed](https://github.com/scroll-tech/zktrie/commit/9d28429589c4703d7d20e01d82f280c37e4022a6) |

#### Background

According to the [docs](https://docs.scroll.io/en/technology/sequencer/zktrie/), zkTrie is a sparse binary merkle patricia trie used to store key-value pairs efficiently. It explains the tree structure, construction, node hashing, and tree operations, including insertion and deletion.

- Merkle Tree: A Merkle Tree is a tree where each leaf node represents a hash of a data block, and each non-leaf node represents the hash of its child nodes.
- Patricia Trie: A Patricia Trie is a type of radix tree or compressed trie used to store key-value pairs efficiently. It encodes the nodes with same prefix of the key to share the common path, where the path is determined by the value of the node key.

Given a key-value pair, we first compute a secure key for the corresponding leaf node by hashing the original key (i.e., account address and storage key) using the Poseidon hash function. This can make the key uniformly distributed over the key space.

We encode the path of a new leaf node by traversing the secure key from Least Significant Bit (LSB) to the Most Significant Bit (MSB). At each step, if the bit is 0, we will traverse to the left child; otherwise, traverse to the right child.

We limit the maximum depth of zkTrie to 248, meaning that the tree will only traverse the lower 248 bits of the key. Because the secure key space is a finite field used by Poseidon hash that doesn’t occupy the full range of $2^{256}$, the bit representation of the key can be ambiguous in the finite field and thus results in a soundness issue in the zk circuit. After we truncate the key to lower 248 bits, the key space can fully occupy the range of $2^{248}$ and won’t have the ambiguity in the bit representation.

There are 3 types of nodes in zkTrie:

- Branch Node: Given the zkTrie is a binary tree, a branch node has two children.
- Leaf Node: A leaf node holds the data of a key-value pair.
- Empty Node: An empty node is a special type of node, indicating the sub-trie that shares the same prefix is empty.

Each type of nodes corresponds to different hashing computations. In Scroll, the Poseidon hash function is configured to take two field element inputs each time and a domain_value as the initial context for domain separation, denoted as follows.

```
h{domain_value}(input1, input2)
```

- Empty Node: The node hash of an empty node is 0.
- Branch Node: The branch node hash is computed as follows
  
  ```
  branchNodeHash = h{BranchNodeType}(leftChildHash, rightChildHash)
  ```
- Leaf Node: The node hash of a leaf node is computed as follows.

    ```
    leafNodeHash = h{LeafNodeType}(nodeKey, valueHash)
    ```

#### Vuln Description

We noticed that the implementation of zktrie uses FFI (Foreign Function Interface) to link C code compiled from Go code. In this scenario, we need to pay attention to casting issues, i.e. data truncation, pointer type errors, struct alligned issues, signedness conversion issues and character encoding inconsistencies.

The following snippet shows one of the casting issues:

```Rust
use std::ffi::{self, c_char, c_int, c_void};
...

#[repr(C)]
struct TrieNode {
    _data: [u8; 0],
    _marker: PhantomData<(*mut u8, PhantomPinned)>,
}

...

#[link(name = "zktrie")]
extern "C" {
    ...
    fn NewTrieNode(data: *const u8, data_sz: c_int) -> *const TrieNode;
    ...
}

...
impl ZkTrieNode {
    pub fn parse(data: &[u8]) -> Self {
        Self {
            trie_node: unsafe { NewTrieNode(data.as_ptr(), data.len() as c_int) },
        }
    }
```

The Rust library regularly needs to convert the input function’s byte array length from the
usize type to the c_int type. Both `usize` and `c_int` are architecture-dependent. So when using `as` operator to perform a casting, there is a possibility that the `usize` type is larger than `c_int` type, resulting in truncation. Therefore, attackers can exploit this to achieve a collision for nodes constructed from different byte arrays.

Specifically, in Rust, `c_int` is a signed integer type that typically has a maximum value defined by `c_int::MAX`. This value depends on the architecture (e.g., 32-bit or 64-bit). When you cast a larger usize value (which is an unsigned integer and can represent larger numbers) to `c_int`, if the value exceeds `c_int::MAX`, it will wrap around due to integer overflow. Note that, the range that usize can represent is always 2 times bigger than c_int type.

Observe that `(c_int::Max * 2 + 2) as c_int` is 0.  Thus, creating two nodes
that have the same prefix and are then padded with different bytes with that length will cause the Go library to interpret only the common prefix of these nodes.

```rust
    #[test]
    fn invalid_cast() {
        init_hash_scheme(hash_scheme);
        // common prefix
        let nd =
        &hex::decode("012098f5fb9e239eab3ceac3f27b81e481dc3124d55ffed523a839ee8446b64864010100000000000000000000000000000000000000000000000000000000018282256f8b00").unwrap();
        // create node1 with prefix padded by zeroes
        let mut vec_nd = nd.to_vec();
        let mut zero_padd_data = vec![0u8; (c_int::MAX as usize) * 2 + 2];
        vec_nd.append(&mut zero_padd_data);
        let node1 = ZkTrieNode::parse(&vec_nd);
        // create node2 with prefix padded by ones
        let mut vec_nd = nd.to_vec();
        let mut one_padd_data = vec![1u8; (c_int::MAX as usize) * 2 + 2];
        vec_nd.append(&mut one_padd_data);
        let node2 = ZkTrieNode::parse(&vec_nd);
        // create node3 with just the prefix
        let node3 =
        ZkTrieNode::parse(&hex::decode("012098f5fb9e239eab3ceac3f27b81e481dc3124d55ffed523a839ee8446b64864010100000000000000000000000000000000000000000000000000000000018282256f8b00").unwrap());
        // all hashes are equal
        assert_eq!(node1.node_hash(), node2.node_hash());
        assert_eq!(node1.node_hash(), node3.node_hash());
    }
```

This finding also allows an attacker to cause a runtime error by choosing the data array with the appropriate length that will cause the cast to result in a negative number.

```rust
    #[test]
    fn invalid_cast() {
        init_hash_scheme(hash_scheme);
        let data = vec![0u8; c_int::MAX as usize + 1];
        println!("{:?}", data.len() as c_int);
        let _nd = ZkTrieNode::parse(&data);
    }
```

**Impact**: An attacker provides two different byte arrays that will have the same node_hash, breaking the assumption that such nodes are hard to obtain.

#### The Fix

Consider  have the code perform the cast in a checked manner by using the
c_int::try_from method to allow validation if the conversion succeeds. Determine whether the Rust functions should allow arbitrary length inputs; document the length requirements and assumptions.
