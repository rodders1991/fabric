/*
Copyright IBM Corp. 2016 All Rights Reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

		 http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package sw

import (
	"crypto/ecdsa"
	"crypto/x509"
	"fmt"

	"github.com/hyperledger/fabric/core/crypto/bccsp"
	"github.com/hyperledger/fabric/core/crypto/primitives"
)

type ecdsaPrivateKey struct {
	k *ecdsa.PrivateKey
}

// FIXME: remove as soon as there's a way to import the key more properly
func NewEcdsaPrivateKey(k *ecdsa.PrivateKey) bccsp.Key {
	return &ecdsaPrivateKey{k: k}
}

// Bytes converts this key to its byte representation,
// if this operation is allowed.
func (k *ecdsaPrivateKey) Bytes() (raw []byte, err error) {
	return
}

// SKI returns the subject key identifier of this key.
func (k *ecdsaPrivateKey) SKI() (ski []byte) {
	raw, _ := primitives.PrivateKeyToDER(k.k)
	// TODO: Error should not be thrown. Anyway, move the marshalling at initialization.

	return primitives.Hash(raw)
}

// Symmetric returns true if this key is a symmetric key,
// false if this key is asymmetric
func (k *ecdsaPrivateKey) Symmetric() bool {
	return false
}

// Private returns true if this key is a private key,
// false otherwise.
func (k *ecdsaPrivateKey) Private() bool {
	return true
}

// PublicKey returns the corresponding public key part of an asymmetric public/private key pair.
// This method returns an error in symmetric key schemes.
func (k *ecdsaPrivateKey) PublicKey() (bccsp.Key, error) {
	return &ecdsaPublicKey{&k.k.PublicKey}, nil
}

type ecdsaPublicKey struct {
	k *ecdsa.PublicKey
}

// FIXME: remove as soon as there's a way to import the key more properly
func NewEcdsaPublicKey(k *ecdsa.PublicKey) bccsp.Key {
	return &ecdsaPublicKey{k: k}
}

// Bytes converts this key to its byte representation,
// if this operation is allowed.
func (k *ecdsaPublicKey) Bytes() (raw []byte, err error) {
	raw, err = x509.MarshalPKIXPublicKey(k.k)
	if err != nil {
		return nil, fmt.Errorf("Failed marshalling key [%s]", err)
	}
	return
}

// SKI returns the subject key identifier of this key.
func (k *ecdsaPublicKey) SKI() (ski []byte) {
	raw, _ := primitives.PublicKeyToPEM(k.k, nil)
	// TODO: Error should not be thrown. Anyway, move the marshalling at initialization.

	return primitives.Hash(raw)
}

// Symmetric returns true if this key is a symmetric key,
// false if this key is asymmetric
func (k *ecdsaPublicKey) Symmetric() bool {
	return false
}

// Private returns true if this key is a private key,
// false otherwise.
func (k *ecdsaPublicKey) Private() bool {
	return false
}

// PublicKey returns the corresponding public key part of an asymmetric public/private key pair.
// This method returns an error in symmetric key schemes.
func (k *ecdsaPublicKey) PublicKey() (bccsp.Key, error) {
	return k, nil
}
