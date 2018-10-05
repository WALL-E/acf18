package cert

import (
	amino "github.com/tendermint/go-amino"
	"github.com/tendermint/tendermint/crypto/ed25519"
)

const (
	CsrAminoRoute = "certificate/csr"
	CrtAminoRoute = "certificate/crt"
)

type Serialization interface {
	Json(cdc *amino.Codec) []byte
	Bytes(cdc *amino.Codec) []byte
}

var _ Serialization = CertificateSigningRequest{}

type CertificateSigningRequest struct {
	Version   int8                  `json:"version"`
	CA        bool                  `json:"ca"`
	CN        string                `json:"cn"`
	IsBanker  bool                  `json:"is_banker"`
	PublicKey ed25519.PubKeyEd25519 `json:"public_key"`
}

func (csr CertificateSigningRequest) Json(cdc *amino.Codec) []byte {
	bz, err := cdc.MarshalJSON(csr)
	if err != nil {
		panic(err)
	}
	return bz
}

func (csr CertificateSigningRequest) Bytes(cdc *amino.Codec) []byte {
	bz, err := cdc.MarshalBinaryBare(csr)
	if err != nil {
		panic(err)
	}
	return bz
}

var _ Serialization = Certificate{}

type Certificate struct {
	CSR       CertificateSigningRequest `json:"csr"`
	Signature []byte                    `json:"signature"`
}

func (crt Certificate) Json(cdc *amino.Codec) []byte {
	bz, err := cdc.MarshalJSON(crt)
	if err != nil {
		panic(err)
	}
	return bz
}

func (crt Certificate) Bytes(cdc *amino.Codec) []byte {
	bz, err := cdc.MarshalBinaryBare(crt)
	if err != nil {
		panic(err)
	}
	return bz
}
