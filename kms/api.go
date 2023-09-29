/*
 Copyright SecureKey Technologies Inc. All Rights Reserved.

 SPDX-License-Identifier: Apache-2.0
*/

package kms

import (
	"errors"
)

// ErrKeyNotFound is an error type that a KMS expects from the Store.Get method if no key stored under the given
// key ID could be found.
var ErrKeyNotFound = errors.New("key not found")
