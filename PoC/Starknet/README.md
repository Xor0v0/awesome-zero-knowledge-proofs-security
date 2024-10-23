# PoC for scure-starknet

## PoC for KS-SBCF-F-01

```bash
# install deps
npm install
# compile deps (idk if there is a convenient way to execute)
cd node_modules/micro-starknet/
tsc index.ts

# generate js script
cd ../..
tsc --init --target es2020
tsc exp.ts
# execute exp.js
node exp.js
```