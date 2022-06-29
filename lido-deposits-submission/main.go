package main

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/prysmaticlabs/periphery/lido-deposits-submission/bindings"
)

type Deposits struct {
	Deposits []*Deposit
}

type Deposit struct {
	Pubkey    string `json:"pubkey"`
	Signature string `json:"signature"`
}

var (
	itemFlag = flag.String("path", "", "")
)

func main() {
	flag.Parse()
	f, err := os.Open(*itemFlag)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	enc, err := io.ReadAll(f)
	if err != nil {
		panic(err)
	}
	items := []*Deposit{}
	if err := json.Unmarshal(enc, &items); err != nil {
		panic(err)
	}
	r, err := rpc.Dial("")
	if err != nil {
		panic(err)
	}
	client := ethclient.NewClient(r)
	addr := common.HexToAddress("")     // node operator registry address.
	prylabsOperatorId := big.NewInt(44) // hardcoded, use the helper func below instead.
	transactor, err := bindings.NewNodeOperatorsRegistryTransactor(addr, client)
	if err != nil {
		panic(err)
	}
	quantity := big.NewInt(10)
	var pubKeys []byte
	var sigs []byte

	for _, item := range items {
		pub, err := hex.DecodeString(item.Pubkey)
		if err != nil {
			panic(err)
		}
		sig, err := hex.DecodeString(item.Signature)
		if err != nil {
			panic(err)
		}
		pubKeys = append(pubKeys, pub...)
		sigs = append(sigs, sig...)
	}
	fmt.Println(len(pubKeys))
	fmt.Println(len(sigs))

	from := common.HexToAddress("")          // gnosis safe prylabs multisig addr
	privateKey, err := crypto.HexToECDSA("") // dont use raw priv key
	if err != nil {
		panic(err)
	}
	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		panic(err)
	}
	tx, err := transactor.AddSigningKeysOperatorBH(&bind.TransactOpts{
		From: from,
		Signer: func(_ common.Address, tx *types.Transaction) (*types.Transaction, error) {
			return types.SignTx(tx, types.NewLondonSigner(chainID), privateKey)
		},
	}, prylabsOperatorId, quantity, pubKeys, sigs)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", tx)
}

func determinePrylabsOperatorId(no *bindings.NodeOperatorsRegistryCaller) (uint64, error) {
	numOps, err := no.GetNodeOperatorsCount(nil)
	if err != nil {
		return 0, err
	}
	total := numOps.Uint64()
	var prylabsOperatorNumber uint64
	var found bool
	for i := uint64(0); i < total; i++ {
		op, err := no.GetNodeOperator(nil, big.NewInt(int64(i)), true)
		if err != nil {
			return 0, err
		}
		if op.Name == "Prylabs" { // do not hardcode
			prylabsOperatorNumber = i
			found = true
		}
	}
	if !found {
		return 0, errors.New("prylabs not found")
	}
	return prylabsOperatorNumber, nil
}
