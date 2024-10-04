package convert

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

var errInvalidDID = errors.New("invalid DID")

// DID is a Decentralized Identifier for NFTs.
type DID struct {
	ChainID         string `json:"chainId"`
	ContractAddress string `json:"contract"`
	TokenID         int    `json:"tokenId"`
}

// DecodeDID decodes a DID string into a DID struct.
func DecodeDID(did string) (DID, error) {
	// sample did "did:nft:1:0xbA5738a18d83D41847dfFbDC6101d37C69c9B0cF_1"
	parts := strings.Split(did, ":")
	if len(parts) != 4 {
		return DID{}, errInvalidDID
	}
	if parts[0] != "did" || parts[1] != "nft" {
		return DID{}, errInvalidDID
	}
	nftParts := strings.Split(parts[3], "_")
	if len(nftParts) != 2 {
		return DID{}, errInvalidDID
	}
	tokenID, err := strconv.Atoi(nftParts[1])
	if err != nil {
		return DID{}, fmt.Errorf("invalid tokenID: %w", err)
	}
	return DID{
		ChainID:         parts[2],
		ContractAddress: nftParts[0],
		TokenID:         tokenID,
	}, nil
}
