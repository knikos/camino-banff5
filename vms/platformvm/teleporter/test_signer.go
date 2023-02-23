// Copyright (C) 2019-2022, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package teleporter

import (
	"github.com/ava-labs/avalanchego/utils/crypto/blsavax"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/ava-labs/avalanchego/ids"
)

// SignerTests is a list of all signer tests
var SignerTests = []func(t *testing.T, s Signer, sk *blsavax.SecretKey, chainID ids.ID){
	TestSignerWrongChainID,
	TestSignerVerifies,
}

// Test that using a random SourceChainID results in an error
func TestSignerWrongChainID(t *testing.T, s Signer, _ *blsavax.SecretKey, _ ids.ID) {
	require := require.New(t)

	msg, err := NewUnsignedMessage(
		ids.GenerateTestID(),
		ids.GenerateTestID(),
		[]byte("payload"),
	)
	require.NoError(err)

	_, err = s.Sign(msg)
	require.Error(err)
}

// Test that a signature generated with the signer verifies correctly
func TestSignerVerifies(t *testing.T, s Signer, sk *blsavax.SecretKey, chainID ids.ID) {
	require := require.New(t)

	msg, err := NewUnsignedMessage(
		chainID,
		ids.GenerateTestID(),
		[]byte("payload"),
	)
	require.NoError(err)

	sigBytes, err := s.Sign(msg)
	require.NoError(err)

	sig, err := blsavax.SignatureFromBytes(sigBytes)
	require.NoError(err)

	pk := blsavax.PublicFromSecretKey(sk)
	msgBytes := msg.Bytes()
	valid := blsavax.Verify(pk, sig, msgBytes)
	require.True(valid)
}
