package v20230401_test

import (
	"context"
	"crypto/rand"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/kde713/azurelangai-go/textanalysis/v20230401"
)

var c v20230401.Client

func init() {
	testingEndpoint := os.Getenv("AZURELANGAI_TEST_ENDPOINT")
	testingKey := os.Getenv("AZURELANGAI_TEST_KEY")
	if testingEndpoint == "" || testingKey == "" {
		panic("AZURELANGAI_TEST_ENDPOINT and AZURELANGAI_TEST_KEY must be set")
	}

	c = v20230401.NewClient(testingEndpoint, testingKey)
}

func generateUUID() string {
	uuidBytes := make([]byte, 16)
	_, err := rand.Read(uuidBytes)
	if err != nil {
		panic(err)
	}

	// Set the UUID version (4) and variant (2)
	uuidBytes[6] = (uuidBytes[6] & 0x0F) | 0x40
	uuidBytes[8] = (uuidBytes[8] & 0x3F) | 0x80

	uuid := fmt.Sprintf("%x-%x-%x-%x-%x",
		uuidBytes[0:4], uuidBytes[4:6], uuidBytes[6:8], uuidBytes[8:10], uuidBytes[10:])

	return uuid
}

func TestClient_AnalyzeTextLanguageDetection(t *testing.T) {
	docID := generateUUID()
	body := v20230401.LanguageDetectionAnalysisInput{
		Documents: []v20230401.LanguageInput{
			{
				ID:   docID,
				Text: "이 문서는 한국어로 쓰여졌습니다.",
			},
		},
	}
	result, err := c.AnalyzeTextLanguageDetection(context.TODO(), body, v20230401.LanguageDetectionTaskParameters{})
	if err != nil {
		t.Fatal(err)
	}
	if len(result.Errors) > 0 {
		t.Fatal(result.Errors[0])
	}
	if len(result.Documents) != 1 {
		t.Fatalf("Expected 1 document, got %d", len(result.Documents))
	}
	if result.Documents[0].ID != docID {
		t.Errorf("Expected document ID %s, got %s", docID, result.Documents[0].ID)
	}
	if result.Documents[0].DetectedLanguage.ISO6391Name != "ko" {
		t.Errorf("Expected detected language ko, got %s", result.Documents[0].DetectedLanguage.ISO6391Name)
	}
}

func TestClient_AnalyzeTextEntityRecognition(t *testing.T) {
	docID := generateUUID()
	body := v20230401.MultiLanguageAnalysisInput{
		Documents: []v20230401.MultiLanguageInput{
			{
				ID:       docID,
				Language: "en",
				Text:     "Hello, my name is Mateo Gomez.",
			},
		},
	}
	result, err := c.AnalyzeTextEntityRecognition(context.TODO(), body, v20230401.EntitiesTaskParameters{})
	if err != nil {
		t.Fatal(err)
	}
	if len(result.Errors) > 0 {
		t.Fatal(result.Errors[0])
	}
	if len(result.Documents) != 1 {
		t.Fatalf("Expected 1 document, got %d", len(result.Documents))
	}
	if result.Documents[0].ID != docID {
		t.Errorf("Expected document ID %s, got %s", docID, result.Documents[0].ID)
	}
	if len(result.Documents[0].Entities) == 0 {
		// At least, model must extract "Mateo Gomez" as a Person.
		t.Error("[Azure Not-working] Expected at least one entity")
	}
}

func TestClient_AnalyzeTextKeyPhraseExtraction(t *testing.T) {
	docID := generateUUID()
	body := v20230401.MultiLanguageAnalysisInput{
		Documents: []v20230401.MultiLanguageInput{
			{
				ID:       docID,
				Language: "en",
				Text:     "Hello, my name is Mateo Gomez.",
			},
		},
	}
	result, err := c.AnalyzeTextKeyPhraseExtraction(context.TODO(), body, v20230401.KeyPhraseTaskParameters{})
	if err != nil {
		t.Fatal(err)
	}
	if len(result.Errors) > 0 {
		t.Fatal(result.Errors[0])
	}
	if len(result.Documents) != 1 {
		t.Fatalf("Expected 1 document, got %d", len(result.Documents))
	}
	if result.Documents[0].ID != docID {
		t.Errorf("Expected document ID %s, got %s", docID, result.Documents[0].ID)
	}
	if len(result.Documents[0].KeyPhrases) == 0 {
		// At least, model must extract "Mateo Gomez", "name" as Key Phrase.
		t.Error("[Azure Not-working] Expected at least one entity")
	}
}

func TestClient_AnalyzeTextSentimentAnalysis(t *testing.T) {
	docID := generateUUID()
	body := v20230401.MultiLanguageAnalysisInput{
		Documents: []v20230401.MultiLanguageInput{
			{
				ID:       docID,
				Language: "en",
				Text:     "Amazingly comfortable!",
			},
		},
	}
	result, err := c.AnalyzeTextSentimentAnalysis(context.TODO(), body, v20230401.SentimentAnalysisTaskParameters{})
	if err != nil {
		t.Fatal(err)
	}
	if len(result.Errors) > 0 {
		t.Fatal(result.Errors[0])
	}
	if len(result.Documents) != 1 {
		t.Fatalf("Expected 1 document, got %d", len(result.Documents))
	}
	if result.Documents[0].ID != docID {
		t.Errorf("Expected document ID %s, got %s", docID, result.Documents[0].ID)
	}
	if result.Documents[0].Sentiment != v20230401.SentimentPositive {
		t.Error("[Azure Not-working] Expected positive sentiment")
	}
	if len(result.Documents[0].Sentences) != 1 {
		t.Errorf("Expected 1 sentences, got %d", len(result.Documents))
	}
}

func TestClient_LongRunningOperation(t *testing.T) {
	docID := generateUUID()
	builder := v20230401.NewJobBuilder(
		fmt.Sprintf("TestClient_LongRunningOperation %s", docID),
		v20230401.MultiLanguageAnalysisInput{
			Documents: []v20230401.MultiLanguageInput{
				{
					ID:       docID,
					Language: "en",
					Text:     "I can describe my experience of this game in two words: \"TECHNOLOGY RECHARGED!\"",
				},
			},
		},
	)
	builder.AddEntityRecognitionTask("test-entityrecognition", v20230401.EntitiesTaskParameters{})
	builder.AddKeyPhraseExtractionTask("test-keyphraseextraction", v20230401.KeyPhraseTaskParameters{})
	builder.AddSentimentAnalysisTask("test-sentimentanalysis", v20230401.SentimentAnalysisTaskParameters{})
	builder.AddExtractiveSummarizationTask("test-extractivesummarization", v20230401.ExtractiveSummarizationTaskParameters{
		SentenceCount: 1,
	})
	builder.AddAbstractiveSummarizationTask("test-abstractivesummarization", v20230401.AbstractiveSummarizationTaskParameters{
		SentenceCount: 1,
	})
	req, err := builder.Build()
	if err != nil {
		t.Fatal(err)
	}
	jobID, err := c.SubmitTextAnalyticsJob(context.TODO(), *req)
	if err != nil {
		t.Fatal(err)
	}

	var taskResults v20230401.Tasks
	for taskResults.Completed == 0 || taskResults.Completed != taskResults.Total {
		time.Sleep(500 * time.Millisecond)
		jobResult, err := c.GetTextAnalyticsJobResult(context.TODO(), jobID)
		if err != nil {
			t.Fatal(err)
		}
		if jobResult.Tasks.Total != 5 {
			t.Fatalf("Expected 5 tasks, got %d", jobResult.Tasks.Total)
		}
		taskResults = jobResult.Tasks
	}

	for _, task := range taskResults.Items {
		if task.Status != v20230401.StatusSucceeded {
			t.Fatalf("Unexpected %s task status: %s", task.Kind, task.Status)
		}
		switch task.Kind {
		case v20230401.LROKindEntityRecognition:
			r, ok := task.Results.(v20230401.EntitiesResult)
			if !ok {
				t.Fatalf("Unexpected result type (failed to decode): %T", task.Results)
			}
			if len(r.Errors) != 0 {
				t.Fatalf("Unexpected error: %s", r.Errors[0])
			}
			if len(r.Documents) != 1 {
				t.Fatalf("Expected 1 document, got %d", len(r.Documents))
			}
			if r.Documents[0].ID != docID {
				t.Errorf("Expected document ID %s, got %s", docID, r.Documents[0].ID)
			}
			if len(r.Documents[0].Entities) == 0 {
				t.Error("Expected at least one entity")
			}
		case v20230401.LROKindKeyPhraseExtraction:
			r, ok := task.Results.(v20230401.KeyPhraseResult)
			if !ok {
				t.Fatalf("Unexpected result type (failed to decode): %T", task.Results)
			}
			if len(r.Errors) != 0 {
				t.Fatalf("Unexpected error: %s", r.Errors[0])
			}
			if len(r.Documents) != 1 {
				t.Fatalf("Expected 1 document, got %d", len(r.Documents))
			}
			if r.Documents[0].ID != docID {
				t.Errorf("Expected document ID %s, got %s", docID, r.Documents[0].ID)
			}
			if len(r.Documents[0].KeyPhrases) == 0 {
				t.Error("Expected at least one key phrase")
			}
		case v20230401.LROKindSentimentAnalysis:
			r, ok := task.Results.(v20230401.SentimentResponse)
			if !ok {
				t.Fatalf("Unexpected result type (failed to decode): %T", task.Results)
			}
			if len(r.Errors) != 0 {
				t.Fatalf("Unexpected error: %s", r.Errors[0])
			}
			if len(r.Documents) != 1 {
				t.Fatalf("Expected 1 document, got %d", len(r.Documents))
			}
			if r.Documents[0].ID != docID {
				t.Errorf("Expected document ID %s, got %s", docID, r.Documents[0].ID)
			}
			if r.Documents[0].Sentiment != v20230401.SentimentPositive {
				t.Error("[Azure Not-working] Expected positive sentiment")
			}
		case v20230401.LROKindAbstractiveSummarization:
			r, ok := task.Results.(v20230401.AbstractiveSummarizationResult)
			if !ok {
				t.Fatalf("Unexpected result type (failed to decode): %T", task.Results)
			}
			if len(r.Errors) != 0 {
				t.Fatalf("Unexpected error: %s", r.Errors[0])
			}
			if len(r.Documents) != 1 {
				t.Fatalf("Expected 1 document, got %d", len(r.Documents))
			}
			if r.Documents[0].ID != docID {
				t.Errorf("Expected document ID %s, got %s", docID, r.Documents[0].ID)
			}
			if len(r.Documents[0].Summaries) == 0 {
				t.Error("Expected at least one summary")
			}
		case v20230401.LROKindExtractiveSummarization:
			r, ok := task.Results.(v20230401.ExtractiveSummarizationResult)
			if !ok {
				t.Fatalf("Unexpected result type (failed to decode): %T", task.Results)
			}
			if len(r.Errors) != 0 {
				t.Fatalf("Unexpected error: %s", r.Errors[0])
			}
			if len(r.Documents) != 1 {
				t.Fatalf("Expected 1 document, got %d", len(r.Documents))
			}
			if r.Documents[0].ID != docID {
				t.Errorf("Expected document ID %s, got %s", docID, r.Documents[0].ID)
			}
			if len(r.Documents[0].Sentences) == 0 {
				t.Error("Expected at least one sentence")
			}
		}
	}
}
