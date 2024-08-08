package header

// Go Substrate RPC Client (GSRPC) provides APIs and types around Polkadot and any Substrate-based chain RPC calls
//
// Copyright 2019 Centrifuge GmbH
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

import (
	"github.com/centrifuge/go-substrate-rpc-client/v4/scale"
	. "github.com/centrifuge/go-substrate-rpc-client/v4/types"
)

type DataLookupIndexItem struct {
	AppId UCompact `json:"appId"`
	Start UCompact `json:"start"`
}
type DataLookup struct {
	Size  UCompact              `json:"size"`
	Index []DataLookupIndexItem `json:"index"`
}

type KateCommitment struct {
	Rows       UCompact `json:"rows"`
	Cols       UCompact `json:"cols"`
	Commitment []U8     `json:"commitment"`
	DataRoot   Hash     `json:"dataRoot"`
}

type V3HeaderExtension struct {
	AppLookup  DataLookup     `json:"appLookup"`
	Commitment KateCommitment `json:"commitment"`
}

// type ExtensionType int

// const (
// 	ExtensionTypeNone ExtensionType = iota
// 	ExtensionTypeV1
// 	ExtensionTypeV2
// )

type HeaderExtensionEnum struct {
	V3 V3HeaderExtension `json:"V3"`
}

type AvailHeader struct {
	ParentHash     Hash                `json:"parentHash"`
	Number         BlockNumber         `json:"number"`
	StateRoot      Hash                `json:"stateRoot"`
	ExtrinsicsRoot Hash                `json:"extrinsicsRoot"`
	Digest         Digest              `json:"digest"`
	Extension      HeaderExtensionEnum `json:"extension"`
}

func (m HeaderExtensionEnum) Encode(encoder scale.Encoder) error {
	var err, err1 error

	err = encoder.PushByte(2)

	if err != nil {
		return err
	}
	err1 = encoder.Encode(m.V3)
	if err1 != nil {
		return err1
	}
	return nil
}

func (m *HeaderExtensionEnum) Decode(decoder scale.Decoder) error {
	b, err := decoder.ReadOneByte()

	if err != nil {
		return err
	}

	if b == 2 {
		err = decoder.Decode(&m.V3)
	}
	if err != nil {
		return err
	}

	return nil
}
