package sdk

import (
	"fmt"

	"github.com/centrifuge/go-substrate-rpc-client/v4/types"
	"github.com/vedhavyas/go-subkey"
)

func convertMultiAddress(receiver string) (types.MultiAddress, error) {
	_, pubkeyBytes, _ := subkey.SS58Decode(receiver)
	address := subkey.EncodeHex(pubkeyBytes)

	dest, err := types.NewMultiAddressFromHexAccountID(address)
	if err != nil {
		_ = fmt.Errorf("cannot create address from given hex:%w", err)
		return types.MultiAddress{}, err
	}
	return dest, nil
}




