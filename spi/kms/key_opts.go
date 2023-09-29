/*
Copyright SecureKey Technologies Inc. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package kms

import (
	"github.com/trustbloc/kms-go/internal/api/kms"
)

// NewKeyOpt creates a new empty key option.
// Not to be used directly. It's intended for implementations of KeyManager interface
// Use WithAttrs() option function below instead.
func NewKeyOpt() *kms.KeyOptions { // nolint
	return kms.NewKeyOpt()
}

// KeyOpts are the create key option.
type KeyOpts = kms.KeyOpts

// WithAttrs option is for creating a key that requires extra attributes.
func WithAttrs(attrs []string) KeyOpts {
	return kms.WithAttrs(attrs)
}
