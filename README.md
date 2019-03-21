# Votechain
Blockchain in Go for an election

## How it Works
Votechain is a parallel system to what is already being used. For this project, the active would be the voter's vote and the transaction would be the act of assigning that vote of the election of a candidate to the body that counts the votes at the end of the election. The block chain is transmitted to all nodes, which are the network computers. When a transaction occurs, these nodes check if it has already been done previously. With each new transaction, this data is propagated to all nodes of the network, this creates an unchanging record of the vote.

## Run Locally
1) Install dependencies
```
make setup
```
2) Set Environment Variable
3) Run project
```
make run
```

## Generate New Address
1) Build the command line interface
```
make build
```
2) Generate New Address
```
./votechain-cli getnewaddress
```

## Sending Data
1) Build the command line interface
```
make build
```
2) Send a transaction
```
./votechain-cli sendvote -from=<private_address> -to=<vote>
 ```
