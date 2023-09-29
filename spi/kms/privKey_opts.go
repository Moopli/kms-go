/*
Copyright SecureKey Technologies Inc. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package kms

import (
	"github.com/trustbloc/kms-go/internal/api/kms"
)

// NewOpt creates a new empty private key option.
// Not to be used directly. It's intended for implementations of KeyManager interface
// Use WithKeyID() option function below instead.
func NewOpt() *kms.PrivateKeyOptions { // nolint
	return kms.NewOpt()
}

// PrivateKeyOpts are the import private key option.
type PrivateKeyOpts = kms.PrivateKeyOpts

// WithKeyID option is for importing a private key with a specified KeyID.
func WithKeyID(keyID string) PrivateKeyOpts {
	return kms.WithKeyID(keyID)
}
