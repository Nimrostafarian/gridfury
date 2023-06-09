# Random

Rand module allows you to post a random number request to the FURY Hub and query the random numbers or the pending random number requests.

## Available Commands

| Name                                | Description                                                      |
| ----------------------------------- | ---------------------------------------------------------------- |
| [request](#grid-tx-random-request)  | Request a random number with an optional block interval          |
| [random](#grid-query-random-random) | Query the generated random number by the request id              |
| [queue](#grid-query-random-queue)   | Query the pending random number requests with an optional height |

## grid tx random request

Request a random number.

```bash
grid tx random request [flags]
```

**Flags:**

| Name, shorthand   | Type   | Required | Default | Description                                                                  |
| ----------------- | ------ | -------- | ------- | ---------------------------------------------------------------------------- |
| --block-interval  | uint64 | true     | 10      | The block interval after which the requested random number will be generated |
| --oracle          | bool   |          | false   | Whether to use the oracle method                                             |
| --service-fee-cap | string |          | ""      | Max service fee, required if "oracle" is true                                |

### Request a random number

Post a random number request to the FURY Hub, the random number will be generated after `--block-interval` blocks.

```bash
# without oracle
grid tx random request --block-interval=100 --from=<key-name> --chain-id=gridfury --fees=0.3grid

# with oracle
grid tx random request --block-interval=100 --oracle=true --service-fee-cap=1grid --from=<key-name> --chain-id=gridfury --fees=0.3grid
```

:::tip
You will get a unique request id if the tx is committed, which can be used to query the status of the request. You can also [query the tx detail](./tx.md#grid-query-tx) to get the request id.
:::

## grid query random random

Query the generated random number by the request id.

```bash
grid query random random <request-id> [flags]
```

## grid query random queue

Query the pending random number requests with an optional block height.

```bash
grid query random queue <gen-height> [flags]
```

### Query random number request queue

Query the pending random number requests with an optional block height at which random numbers will be generated or request service.

```bash
grid query random queue 100000
```
