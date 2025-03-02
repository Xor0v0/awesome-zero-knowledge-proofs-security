# Checklist for Zero Knowledge Circuits

> No security by obscurity.

## Base checks

1. Completeness: Are all constraints correctly defined and covering all input variables?
   - Make sure circuits pass correct computation and reject incorrect computation.
   - Each variable should contribute to either directly or indirectly to a constraint (which lead to assigned but not constrained bug).
2. Input range validation: Do input variables fall within expected ranges to prevent overflow or undefined behaviour?
3. Over/under flow check: both input and operation need adequate checks.
4. Consistency check: If multiple paths compute the same variable, do they yield the same result?
5. Zero knowledge property check: Does the circuit leak private information?

## Integration checks

1. Document all assumptions especially implication assumption.
   - Libraries built for high-performance circuits do not enforce preconditions on inputs.
   - No underconstrained signals during library integration.
2. When circuits integrate with smart contract part, system maybe incur extra bug.

## Architecture design checks

- Double-spending prevention, e.g. if protocol needs nullifier mechanism?
- Front-running

## Circuit Optimizations

1. No redundant constraints: Are there any unnecessary constraints that add computational burden? (Minimum constraint rule)

## Misc

1. Third-party library security:
   - Understanding implicit assumptions: Do not overlook any hidden preconditions required by the library.
   - Avoiding unsafe parameters: Ensure that cryptographic components are initialized with secure and recommended parameters, e.g. `component m = MiMCSponge(1,2,1)` in Circom, which generates am increment output hasher.
2. If customizing a cryptographic scheme, ensure that it has undergone security proof and strictly follows security guidelines during implementation
3. Code readability: Is the circuit code well-structured and easy to audit?

