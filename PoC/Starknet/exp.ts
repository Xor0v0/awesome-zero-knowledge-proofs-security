import { hashChain } from "micro-starknet"

var h1 = hashChain([1, 2, 3])
console.log(h1)

var h2 = hashChain([2, 3])
console.log(h2)

var h3 = hashChain([1, h2])
console.log(h3)