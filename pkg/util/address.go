package util

import (
	"bytes"

	"github.com/ipfs/fs-repo-migrations/ipfs-6-to-7/gx/ipfs/QmWFAMPqsEyUX7gDUsRVmMWz59FxSpJ1b2v6bJ1yYzo7jY/go-base58-fast/base58"
)

func EncodeAddress(publicKeyBytes []byte) string {
	bs := bytes.Join([][]byte{
		[]byte{0x6F},
		Hash160(publicKeyBytes),
	}, []byte{})
	checksum := Hash256(bs)[:4]
	return base58.Encode(bytes.Join([][]bytes{bs, checksum}, []byte{}))
}
