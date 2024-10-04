package convert_test

import (
	"testing"

	"github.com/DIMO-Network/model-garage/pkg/convert"
	"github.com/stretchr/testify/require"
)

func TestDecodeDID(t *testing.T) {
	tests := []struct {
		name          string
		input         string
		expectedDID   convert.DID
		expectedError bool
	}{
		{
			name:  "valid DID",
			input: "did:nft:137:0xbA5738a18d83D41847dfFbDC6101d37C69c9B0cF_123",
			expectedDID: convert.DID{
				ChainID:         "137",
				ContractAddress: "0xbA5738a18d83D41847dfFbDC6101d37C69c9B0cF",
				TokenID:         123,
			},
		},
		{
			name:          "invalid format - wrong part count",
			input:         "did:nft:1",
			expectedDID:   convert.DID{},
			expectedError: true,
		},
		{
			name:          "invalid format - wrong token part count",
			input:         "did:nft:1:0xbA5738a18d83D41847dfFbDC6101d37C69c9B0cF",
			expectedDID:   convert.DID{},
			expectedError: true,
		},
		{
			name:          "invalid tokenID",
			input:         "did:nft:1:0xbA5738a18d83D41847dfFbDC6101d37C69c9B0cF_notanumber",
			expectedDID:   convert.DID{},
			expectedError: true,
		},
		{
			name:          "invalid DID string - wrong prefix",
			input:         "invalidprefix:nft:1:0xbA5738a18d83D41847dfFbDC6101d37C69c9B0cF_1",
			expectedDID:   convert.DID{},
			expectedError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			did, err := convert.DecodeDID(tt.input)

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
