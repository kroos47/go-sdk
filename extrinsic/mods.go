package extrinsic

import "github.com/centrifuge/go-substrate-rpc-client/v4/types"

type SignatureOptions struct {
	Era                types.ExtrinsicEra
	Nonce              types.UCompact
	Tip                types.UCompact
	SpecVersion        types.U32
	GenesisHash        types.Hash
	BlockHash          types.Hash
	TransactionVersion types.U32
	AppID              types.UCompact
}

type ExtrinsicPayloadV4 struct {
	types.ExtrinsicPayloadV3
	AppID              types.UCompact
	TransactionVersion types.U32
}

type ExtrinsicSignatureV4 struct {
	Signer    types.MultiAddress
	Signature types.MultiSignature
	Era       types.ExtrinsicEra // extra via system::CheckEra
	Nonce     types.UCompact     // extra via system::CheckNonce (Compact<Index> where Index is u32))
	Tip       types.UCompact     // extra via balances::TakeFees (Compact<Balance> where Balance is u128))
	AppID     types.UCompact     // Avail specific AppID
}

type Extrinsic struct {
	// Version is the encoded version flag (which encodes the raw transaction version and signing information in one byte)
	Version byte
	// Signature is the ExtrinsicSignatureV4, it's presence depends on the Version flag
	Signature ExtrinsicSignatureV4
	// Method is the call this extrinsic wraps
	Method types.Call
}
