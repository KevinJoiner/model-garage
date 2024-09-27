package convert_test

import (
	"context"
	"testing"

	"github.com/DIMO-Network/model-garage/pkg/vss/convert"
)

func BenchmarkConvertFromV1DataConversion(b *testing.B) {
	getter := &tokenGetter{}
	inputData := []byte(fullInputJSON)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := convert.SignalsFromV1Payload(context.Background(), getter, inputData)
		if err != nil {
			b.Fatalf("error converting full input data: %v", err)
		}
	}
}

func BenchmarkConvertFromV2DataConversion(b *testing.B) {
	inputData := []byte(fullV2InputJSON)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := convert.SignalsFromV2Payload(inputData)
		if err != nil {
			b.Fatalf("error converting full input data: %v", err)
		}
	}
}
