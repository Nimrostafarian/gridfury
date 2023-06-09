---
order: 5
---

# Bech32 on FURYhub

Bech32 is a new Bitcoin address format proposed by Pieter Wuille and Greg Maxwell. Besides Bitcoin addresses, Bech32 can encode any short binary data. In the FURY network, keys and addresses may refer to a number of different roles in the network like accounts, validators etc. The FURY network is designed to use the Bech32 address format to provide robust integrity checks on data. The human readable part(HRP) makes it more efficient to read and the users could see error messages. More details in [bip-0173](https://github.com/bitcoin/bips/blob/master/bip-0173.mediawiki)

## Human Readable Part Table

| HRP | Definition                              |
| --- | --------------------------------------- |
| iaa | FURYhub Account Address                 |
| iap | FURYhub Account Public Key              |
| iva | FURYhub Validator's Operator Address    |
| ivp | FURYhub Validator's Operator Public Key |
| ica | Tendermint Consensus Address            |
| icp | Tendermint Consensus Public Key         |

## Encoding

Not all interfaces to FURYhub users should be exposed as bech32 interfaces. Many addresses are still in hex or base64 encoded form.

To covert between other binary representation of addresses and keys, it is important to first apply the Amino encoding process before bech32 encoding.

## Account Key Example

Account Key, aka. [Application Key](validator-faq.md#application-key). Once you create a new address, you should see the following:

```bash
NAME:    TYPE:           ADDRESS:                                PUBKEY:
test1    local    did:grid:iaa18ekc4dswwrh2a6lfyev4tr25h5y76jkpclyxkz    iap1addwnpepqgxa40ww28uy9q46gg48g6ulqdzwupyjcwfumgfjpvz7krmg5mrnwk5xq9l
```

This means you have created a new address `did:grid:iaa18ekc4dswwrh2a6lfyev4tr25h5y76jkpclyxkz`, with the HRP `iaa`. And its public key could be encoded into `iap1addwnpepqgxa40ww28uy9q46gg48g6ulqdzwupyjcwfumgfjpvz7krmg5mrnwk5xq9l`, with the HRP `iap`.

## Validator Key Example

Validator Key, aka. [Tendermint Key](validator-faq.md#tendermint-key). A Tendermint Consensus Public key is generated when the node is created with  `grid init`.
You can get this value with

```bash
grid tendermint show-validator --home=<grid-home>
```

Example Output:

```bash
did:grid:icp1zcjduepqzuz420weqehs3mq0qny54umfk5r78yup6twtdt7mxafrprms5zqsjeuxvx
```
