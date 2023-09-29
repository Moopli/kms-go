/*
Copyright SecureKey Technologies Inc. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package kms

// PrivateKeyOptions holds options for ImportPrivateKey.
type PrivateKeyOptions struct {
	ksID string
}

// NewOpt creates a new empty private key option.
// Not to be used directly. It's intended for implementations of KeyManager interface
// Use WithKeyID() option function below instead.
func NewOpt() *PrivateKeyOptions {
	return &PrivateKeyOptions{}
}

// KsID gets the KsID to be used for import a private key.
// Not to be used directly. It's intended for implementations of KeyManager interface
// Use WithKeyID() option function below instead.
func (pk *PrivateKeyOptions) KsID() string {
	return pk.ksID
}

// PrivateKeyOpts are the import private key option.
type PrivateKeyOpts func(opts *PrivateKeyOptions)

// WithKeyID option is for importing a private key with a specified KeyID.
func WithKeyID(keyID string) PrivateKeyOpts {
	return func(opts *PrivateKeyOptions) {
		opts.ksID = keyID
	}
}
