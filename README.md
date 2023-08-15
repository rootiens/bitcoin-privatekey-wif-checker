# Purpose of the Application:

The purpose of this application is to generate all possible combinations of a Bitcoin private key WIF by appending a second part (in a scenario that you only know characters, not order) to a known first part. It then checks each generated key to verify its validity using the Bitcoin WIF checksum algorithm.

# How to use it
`git clone https://github.com/rootiens/bitcoin-privatekey-wif-checker.git`
`cd bitcoin-privatekey-wif-checker`
`go mod tidy`
`go run main.go`

# Example
Correct address : 
`L3nJu1EWNPSL7JEBtzAurvVY9dTGps3jbBYMpb25usXbMGpVbcS9`

Known first part:
`L3nJu1EWNPSL7JEBtzAurvVY9dTGps3jbBYMpb25usXbMGp`

Second part (Unordered, but with correct chars):
`cbVS9`

Output:
`Correct key found: L3nJu1EWNPSL7JEBtzAurvVY9dTGps3jbBYMpb25usXbMGpVbcS9`
