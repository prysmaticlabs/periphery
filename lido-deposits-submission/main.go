package main

import (
	"encoding/hex"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/pkg/errors"
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
	itemFlag                    = flag.String("deposit-data", "", "Path to the deposit data json file")
	operatorRegistryAddressFlag = flag.String("operator-registry-address", "", "Address of the Lido node operator registry. Mainnet=0x55032650b14df07b85bF18A3a3eC8E0Af2e028d5 Goerli=0x9D4AF1Ee19Dad8857db3a45B0374c81c8A1C6320")
	jsonRPCEndpointFlag         = flag.String("web3", "http://localhost:8545", "JSON RPC endpoint of the Ethereum node")
	operatorNameFlag            = flag.String("operator-name", "Prylabs", "Name of the Lido operator")
	operatorIDFlag              = flag.Int("operator-id", -1, "ID of the operator to make deposits. If unset then the operator ID will be looked up by the operator name.")
)

func main() {
	flag.Parse()

	if len(*itemFlag) == 0 {
		panic("No deposit data provided")
	}
	if len(*operatorRegistryAddressFlag) == 0 {
		panic("No operator registry address provided")
	}

	fmt.Printf("Opening deposit data file: %s\n", *itemFlag)
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
	numItems := uint64(len(items))
	fmt.Printf("Found %d key signature pairs in deposit data.\n", numItems)

	r, err := rpc.Dial(*jsonRPCEndpointFlag)
	if err != nil {
		panic(err)
	}
	client := ethclient.NewClient(r)
	registry, err := bindings.NewNodeOperatorsRegistryCaller(common.HexToAddress(*operatorRegistryAddressFlag), client)
	if err != nil {
		panic(err)
	}
	operatorID := uint64(*operatorIDFlag)
	if *operatorIDFlag == -1 {
		fmt.Printf("No operator ID provided. Looking up operator ID by name \"%s\"\n", *operatorNameFlag)
		operatorID, err = determineLidoOperatorId(registry, *operatorNameFlag)
		if err != nil {
			panic(errors.Wrapf(err, "failed to find operator ID for operator %s", *operatorNameFlag))
		}
	}
	fmt.Printf("Operator ID: %d\n", operatorID)

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

	registryABI, err := bindings.NodeOperatorsRegistryMetaData.GetAbi()
	if err != nil {
		panic(err)
	}

	/**
	 * @notice Add `_quantity` validator signing keys of operator #`_id` to the set of usable keys. Concatenated keys are: `_pubkeys`. Can be done by node operator in question by using the designated rewards address.
	 * @dev Along with each key the DAO has to provide a signatures for the
	 *      (pubkey, withdrawal_credentials, 32000000000) message.
	 *      Given that information, the contract'll be able to call
	 *      deposit_contract.deposit on-chain.
	 * @param _operator_id Node Operator id
	 * @param _quantity Number of signing keys provided
	 * @param _pubkeys Several concatenated validator signing keys
	 * @param _signatures Several concatenated signatures for (pubkey, withdrawal_credentials, 32000000000) messages
	 */
	// function addSigningKeysOperatorBH(uint256 _operator_id, uint256 _quantity, bytes _pubkeys, bytes _signatures) external;
	calldata, err := registryABI.Pack("addSigningKeysOperatorBH", big.NewInt(int64(operatorID)), big.NewInt(int64(numItems)), pubKeys, sigs)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Send the following call data to the Lido node operator contract at address %s\n", *operatorRegistryAddressFlag)
	fmt.Printf("Calldata = %s\n", hexutil.Encode(calldata))
}

func determineLidoOperatorId(caller *bindings.NodeOperatorsRegistryCaller, operatorName string) (uint64, error) {
	numOps, err := caller.GetNodeOperatorsCount(nil)
	if err != nil {
		return 0, err
	}
	total := numOps.Uint64()
	var operatorID uint64
	var found bool
	for i := uint64(0); i < total; i++ {
		op, err := caller.GetNodeOperator(nil, big.NewInt(int64(i)), true)
		if err != nil {
			return 0, err
		}
		if op.Name == operatorName {
			operatorID = i
			found = true
		}
	}
	if !found {
		return 0, errors.New("operator not found")
	}
	return operatorID, nil
}
