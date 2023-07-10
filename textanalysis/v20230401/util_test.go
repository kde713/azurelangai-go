package v20230401_test

import (
	"testing"

	"github.com/kde713/azurelangai-go/textanalysis/v20230401"
)

func TestParseJobID(t *testing.T) {
	exampleOperationLocation := "https://this-is-example.cognitiveservices.azure.com/language/analyze-text/jobs/e9769bf0-e12f-4463-bf41-07d400bdf45c?api-version=2023-04-01"
	expectedJobID := "e9769bf0-e12f-4463-bf41-07d400bdf45c"
	gotJobID, err := v20230401.ParseJobID(exampleOperationLocation)
	if err != nil {
		t.Fatal(err)
	}
	if gotJobID != expectedJobID {
		t.Fatalf("Expected job ID %s, got %s", expectedJobID, gotJobID)
	}
}
