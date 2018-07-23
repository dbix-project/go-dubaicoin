// Copyright 2015 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package params

// MainnetBootnodes are the enode URLs of the P2P bootstrap nodes running on
// the main Ethereum network.
var MainnetBootnodes = []string{
	// Mainnet Dubaicoin-DBIX Go Bootnodes
	"enode://516180cabd015454788a6a0a4f3508a4c021bee56064640e6430e418513e3e9d428f8dc5eda01511baf6d39ae2cbb8f2b32117137ca0d5d9435414397a9fa913@138.68.175.161:57955",
	"enode://313b0f81d49c9158c6bb49d82858be09eb903a6ba04bf726c4bd073638049094f0df0ca2e2fb83a18f3815730c4f59522370cb76fafd4009d706a9040bb0a4d0@188.226.143.29:57955",
	"enode://7607042debf4b2de9c894186c2f683ad23e2597612738dc72f5527a4480ed6ec696d6ad97f4bfecb5daa0c2920755350a1b226ff10b588dbbf07037528d1f76f@128.199.229.117:57955",
	"enode://053e556a29aa708b402229578bd714fd4675ca42b8fd4905c2d198bc28eb280dcf95c2981e2471e54dc1c7ee8dcbc9605c49953be09c49ad669828b97d3e9dc3@159.203.58.168:57955",
	"enode://ccbccdaaca76d53e5494dcffe832a0dad2bb2812cdd3febac3a72586943eddbc04b8f92b2029333f96baa6145adbbf651c0cf309fc3cb18dc5667c3242e119d9@139.59.75.90:57955",
	"enode://44afd7dd16d8c4e33fcecc91bf357102b82a5039bdc4f7a6b65b777c9aada8a748c8ad5cc714161a5c7fbdb03d14d69c6bab2a66a34ce323d6e0c0655db8d46b@198.211.103.52:57955",
	"enode://b1cc94fd01d626574cc8e238613801896d7f1928ea240bc767c7c7bccfb9615b574cf4dceba3201a8aef5512bcba21b28552fcacaccc6d08eb631cff6674bb11@104.236.135.29:57955",
}

// TestnetBootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Morden test network.
var TestnetBootnodes = []string{
	// Testnet Dubaicoin-DBIX Go Bootnodes
	/*
	"enode://e4533109cc9bd7604e4ff6c095f7a1d807e15b38e9bfeb05d3b7c423ba86af0a9e89abbf40bd9dde4250fef114cd09270fa4e224cbeef8b7bf05a51e8260d6b8@94.242.229.4:57955",
	"enode://8c336ee6f03e99613ad21274f269479bf4413fb294d697ef15ab897598afb931f56beb8e97af530aee20ce2bcba5776f4a312bc168545de4d43736992c814592@94.242.229.203:57955",
	*/
}

// DiscoveryV5Bootnodes are the enode URLs of the P2P bootstrap nodes for the
// experimental RLPx v5 topic-discovery network.
var DiscoveryV5Bootnodes = []string{
	/*
	"enode://0cc5f5ffb5d9098c8b8c62325f3797f56509bff942704687b6530992ac706e2cb946b90a34f1f19548cd3c7baccbcaea354531e5983c7d1bc0dee16ce4b6440b@40.118.3.223:30305",
	"enode://1c7a64d76c0334b0418c004af2f67c50e36a3be60b5e4790bdac0439d21603469a85fad36f2473c9a80eb043ae60936df905fa28f1ff614c3e5dc34f15dcd2dc@40.118.3.223:30308",
	"enode://85c85d7143ae8bb96924f2b54f1b3e70d8c4d367af305325d30a61385a432f247d2c75c45c6b4a60335060d072d7f5b35dd1d4c45f76941f62a4f83b6e75daaf@40.118.3.223:30309",
	*/
}
