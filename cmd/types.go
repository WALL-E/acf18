package cmd

import (
	amino "github.com/tendermint/go-amino"
	"github.com/tendermint/tendermint/crypto/ed25519"
)

var privateKeyFile string
var publicKeyFile string
var csrFile string

type Serialization interface {
	ToJson() []byte
}

type Csr struct {
	version   int8                  `json:"version"`
	ca        bool                  `json:"ca"`
	cn        string                `json:"cn"`
	publicKey ed25519.PubKeyEd25519 `json:"public_key"`
}

func (csr Csr) ToJson() []byte {
	bz, err := cdc.MarshalJSON(csr)
	if err != nil {
		panic(err)
	}
	return bz
}

type Crt struct {
	csr       Csr                         `json:"csr"`
	signature [ed25519.SignatureSize]byte `json:"signature"`
}

func (crt Crt) ToJson() []byte {
	bz, err := cdc.MarshalJSON(crt)
	if err != nil {
		panic(err)
	}
	return bz
}

const (
	CsrAminoRoute = "ed25519/csr"
	CrtAminoRoute = "ed25519/crt"
)

var cdc = amino.NewCodec()

func init() {
	cdc.RegisterInterface((*Serialization)(nil), nil)

	cdc.RegisterConcrete(Csr{},
		CsrAminoRoute, nil)
	cdc.RegisterConcrete(Crt{},
		CrtAminoRoute, nil)
}