/*
Copyright SecureKey Technologies Inc. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package crypto

import (
	"github.com/trustbloc/kms-go/internal/api/crypto"
)

// NewOpt creates a new empty wrap key option.
// Not to be used directly. It's intended for implementations of Crypto interface
// Use WithSender() option function below instead.
func NewOpt() *crypto.WrapKeyOptions {
	return crypto.NewOpt()
}

// WrapKeyOpts are the crypto.Wrap key options.
type WrapKeyOpts = crypto.WrapKeyOpts

// WithSender option is for setting a sender key with crypto wrapping (eg: AuthCrypt). For Anoncrypt,
// this option must not be set.
// Sender is a key used for ECDH-1PU key agreement for authenticating the sender.
// senderkey can be of the following there types:
//   - *keyset.Handle (requires private key handle for crypto.WrapKey())
//   - *crypto.PublicKey (available for UnwrapKey() only)
//   - *ecdsa.PublicKey (available for UnwrapKey() only)
func WithSender(senderKey interface{}) WrapKeyOpts {
	return crypto.WithSender(senderKey)
}

// WithXC20PKW option is a flag option for crypto wrapping. When used, key wrapping will use XChacha20Poly1305
// encryption as key wrapping. The absence of this option (default) uses AES256-GCM encryption as key wrapping. The KDF
// used in the crypto wrapping function is selected based on the type of recipient key argument of KeyWrap(), it is
// independent of this option.
func WithXC20PKW() WrapKeyOpts {
	return crypto.WithXC20PKW()
}

// WithTag option is to instruct the key wrapping function of the authentication tag to be used in the wrapping process.
// It is mainly used with CBC+HMAC content encryption to authenticate the sender of an encrypted JWE message (ie
// authcrypt/ECDH-1PU). The absence of this option means the sender's identity is not revealed (ie anoncrypt/ECDH-ES).
func WithTag(tag []byte) WrapKeyOpts {
	return crypto.WithTag(tag)
}

// WithEPK option is to instruct the key wrapping function of the ephemeral key to be used in the wrapping process.
// It is mainly used for ECDH-1PU during KDF. This option allows passing a predefined EPK instead of generating a new
// one when wrapping. It is useful for Wrap() call only since Unwrap() already uses a predefined EPK. The absence of
// this option means a new EPK will be generated internally.
func WithEPK(epk *PrivateKey) WrapKeyOpts {
	return crypto.WithEPK(epk)
}
