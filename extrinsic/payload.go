package extrinsic

import (
	"fmt"

	"github.com/centrifuge/go-substrate-rpc-client/v4/signature"
	"github.com/centrifuge/go-substrate-rpc-client/v4/types"
	"github.com/centrifuge/go-substrate-rpc-client/v4/types/codec"
	"github.com/vedhavyas/go-subkey/scale"
)

func (e ExtrinsicPayloadV4) Sign(signer signature.KeyringPair) (types.Signature, error) {
	b, err := codec.Encode(e)
	if err != nil {
		return types.Signature{}, err
	}

	sig, err := signature.Sign(b, signer.URI)
	return types.NewSignature(sig), err
}

func (e ExtrinsicPayloadV4) Encode(encoder scale.Encoder) error {
	err := encoder.Encode(e.Method)
	if err != nil {
		return err
	}

	err = encoder.Encode(e.Era)
	if err != nil {
		return err
	}

	err = encoder.Encode(e.Nonce)
	if err != nil {
		return err
	}

	err = encoder.Encode(e.Tip)
	if err != nil {
		return err
	}

	err = encoder.Encode(e.SpecVersion)
	if err != nil {
		return err
	}

	err = encoder.Encode(e.TransactionVersion)
	if err != nil {
		return err
	}

	err = encoder.Encode(e.GenesisHash)
	if err != nil {
		return err
	}

	err = encoder.Encode(e.BlockHash)
	if err != nil {
		return err
	}

	return nil
}

// Decode does nothing and always returns an error. ExtrinsicPayloadV4 is only used for encoding, not for decoding
func (e *ExtrinsicPayloadV4) Decode(decoder scale.Decoder) error {
	return fmt.Errorf("decoding of ExtrinsicPayloadV4 is not supported")
}
