package adapter

import (
	"math/big"
	"github.com/ethereum/go-ethereum/common"
)

type Erc20Transfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int

}