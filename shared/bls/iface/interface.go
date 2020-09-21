// Package iface provides the BLS interfaces that are implemented by the various BLS wrappers.
//
// This package should not be used by downstream consumers. These interfaces are re-exporter by
// github.com/prysmaticlabs/prysm/shared/bls. This package exists to prevent an import circular
// dependency.
package iface

// SecretKey represents a BLS secret or private key.
type SecretKey interface {
	PublicKey() PublicKey
	Sign(msg []byte) Signature
	Marshal() []byte
}

// PublicKey represents a BLS public key.
type PublicKey interface {
	Marshal() []byte
	Copy() PublicKey
	Aggregate(p2 PublicKey) PublicKey
	Equals(p2 PublicKey) bool
}

// Signature represents a BLS signature.
type Signature interface {
	Verify(pubKey PublicKey, msg []byte) bool
	AggregateVerify(pubKeys []PublicKey, msgs [][32]byte) bool
	FastAggregateVerify(pubKeys []PublicKey, msg [32]byte) bool
	Marshal() []byte
	Copy() Signature
}
