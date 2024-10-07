package cloudevent_test

import (
	"testing"

	"github.com/DIMO-Network/model-garage/pkg/cloudevent"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"
)

func TestDecodeDID(t *testing.T) {
	tests := []struct {
		name          string
		input         string
		expectedDID   cloudevent.DID
		expectedError bool
	}{
		{
			name:  "valid DID",
			input: "did:nft:137:0xbA5738a18d83D41847dfFbDC6101d37C69c9B0cF_123",
			expectedDID: cloudevent.DID{
				ChainID:         "137",
				ContractAddress: common.HexToAddress("0xbA5738a18d83D41847dfFbDC6101d37C69c9B0cF"),
				TokenID:         123,
			},
		},
		{
			name:          "invalid format - wrong part count",
			input:         "did:nft:1",
			expectedDID:   cloudevent.DID{},
			expectedError: true,
		},
		{
			name:          "invalid format - wrong token part count",
			input:         "did:nft:1:0xbA5738a18d83D41847dfFbDC6101d37C69c9B0cF",
			expectedDID:   cloudevent.DID{},
			expectedError: true,
		},
		{
			name:          "invalid tokenID",
			input:         "did:nft:1:0xbA5738a18d83D41847dfFbDC6101d37C69c9B0cF_notanumber",
			expectedDID:   cloudevent.DID{},
			expectedError: true,
		},
		{
			name:          "invalid DID string - wrong prefix",
			input:         "invalidprefix:nft:1:0xbA5738a18d83D41847dfFbDC6101d37C69c9B0cF_1",
			expectedDID:   cloudevent.DID{},
			expectedError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			did, err := cloudevent.DecodeDID(tt.input)

			// Check if the error matches the expected error
			if tt.expectedError {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}

			// Check if the DID struct matches the expected DID
			require.Equal(t, tt.expectedDID, did)
		})
	}
}
