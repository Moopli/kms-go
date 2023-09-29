/*
Copyright SecureKey Technologies Inc. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package kms

// KeyOptions holds options for Create, Rotate and CreateAndExportPubKeyBytes.
type KeyOptions struct {
	attrs []string
}

// NewKeyOpt creates a new empty key option.
// Not to be used directly. It's intended for implementations of KeyManager interface
// Use WithAttrs() option function below instead.
func NewKeyOpt() *KeyOptions {
	return &KeyOptions{}
}

// Attrs gets the additional attributes to be used for a key creation.
// Not to be used directly. It's intended for implementations of KeyManager interface
// Use WithAttrs() option function below instead.
func (pk *KeyOptions) Attrs() []string {
	return pk.attrs
}

// KeyOpts are the create key option.
type KeyOpts func(opts *KeyOptions)

// WithAttrs option is for creating a key that requires extra attributes.
func WithAttrs(attrs []string) KeyOpts {
	return func(opts *KeyOptions) {
		opts.attrs = attrs
	}
}
