package cloudevent

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/ethereum/go-ethereum/common"
)

var errInvalidDID = errors.New("invalid DID")

// NFTDID is a Decentralized Identifier for NFTs.
type NFTDID struct {
	ChainID         string         `json:"chainId"`
	ContractAddress common.Address `json:"contract"`
	TokenID         uint32         `json:"tokenId"`
}

// DecodeNFTDID decodes a DID string into a DID struct.
func DecodeNFTDID(did string) (NFTDID, error) {
	// sample did "did:nft:1:0xbA5738a18d83D41847dfFbDC6101d37C69c9B0cF_1"
	parts := strings.Split(did, ":")
	if len(parts) != 4 {
		return NFTDID{}, errInvalidDID
	}
	if parts[0] != "did" || parts[1] != "nft" {
		return NFTDID{}, errInvalidDID
	}
	nftParts := strings.Split(parts[3], "_")
	if len(nftParts) != 2 {
		return NFTDID{}, errInvalidDID
	}
	tokenID, err := strconv.ParseUint(nftParts[1], 10, 32)
	if err != nil {
		return NFTDID{}, fmt.Errorf("invalid tokenID: %w", err)
	}
	addrBytes := nftParts[0]
	if !common.IsHexAddress(addrBytes) {
		return NFTDID{}, errors.New("invalid contract address")
	}

	return NFTDID{
		ChainID:         parts[2],
		ContractAddress: common.HexToAddress(string(addrBytes)),
		TokenID:         uint32(tokenID),
	}, nil
}
