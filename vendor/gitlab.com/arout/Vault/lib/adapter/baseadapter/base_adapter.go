package baseadapter

import (
	log "github.com/mgutz/logxi/v1"
	"gitlab.com/arout/Vault/lib"
)

// IBlockchainAdapter Blockchain Adapter Interface
// contains common methods for all Blockchain Adapter variants
type IBlockchainAdapter interface {
	DerivePrivateKey(log.Logger) (string, error)
	DerivePublicKey(log.Logger) (string, error)
	DeriveAddress(log.Logger) (string, error)
	GetBlockchainNetwork() string
	CreateSignedTransaction(lib.IRawTx, log.Logger) (string, error)
}

// BlockchainAdapter contains common fields for
// all Blockchain Adapter variants
type BlockchainAdapter struct {
	Seed           []byte
	DerivationPath string
	PrivateKey     string
	IsDev          bool
	IBlockchainAdapter
}