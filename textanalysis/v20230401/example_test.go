package v20230401_test

import (
	"context"
	"errors"
	"fmt"

	azuretextanalysis "github.com/kde713/azurelangai-go/textanalysis/v20230401"
)

var azureTextAnalysisClient azuretextanalysis.Client

func ExampleNewClient() {
	// Get your endpoint and key from the Azure portal.
	endpoint := "https://<this-is-example>.cognitiveservices.azure.com/"
	key := "<this-is-example>"

	// Create a client.
	azureTextAnalysisClient = azuretextanalysis.NewClient(endpoint, key)
}

func ExampleClient_AnalyzeTextLanguageDetection() {
	// Create an input.
	input := azuretextanalysis.LanguageDetectionAnalysisInput{
		Documents: []azuretextanalysis.LanguageInput{
			{
				ID:   "<unique-id>",
				Text: "이 문서는 한국어로 쓰여졌습니다.",
			},
		},
	}

	// Create a task parameters.
	params := azuretextanalysis.LanguageDetectionTaskParameters{}

	// Call the service.
	result, err := azureTextAnalysisClient.AnalyzeTextLanguageDetection(context.TODO(), input, params)
	var taskErr *azuretextanalysis.TaskError
	if err != nil {
		if errors.As(err, &taskErr) {
			// Azure Task Failed with general error response
			// You can get azuretextanalysis.ErrorInformation data from here
			fmt.Println(taskErr.Information)
		} else {
			// Other errors like network error, parse error, etc.
			fmt.Println(err)
		}
	}

	// Do something with the result.
	fmt.Println(result)
}

func ExampleClient_SubmitTextAnalyticsJob() {
	// Create an document input
	input := azuretextanalysis.MultiLanguageAnalysisInput{
		Documents: []azuretextanalysis.MultiLanguageInput{
			{
				ID:       "<unique-id>",
				Language: "en",
				Text:     "More than a year and a half after Covid-19 swept across the globe, .....",
			},
		},
	}

	// Start a job builder
	builder := azuretextanalysis.NewJobBuilder("<job-display-name>", input)

	// Add jobs you want
	builder.AddKeyPhraseExtractionTask("<task-name>", azuretextanalysis.KeyPhraseTaskParameters{})
	builder.AddAbstractiveSummarizationTask("<task-name>", azuretextanalysis.AbstractiveSummarizationTaskParameters{
		SentenceCount: 1,
	})

	// Submit the job
	jobReq, err := builder.Build()
	if err != nil {
		panic(err)
	}
	jobID, err := azureTextAnalysisClient.SubmitTextAnalyticsJob(context.TODO(), *jobReq)
	if err != nil {
		panic(err)
	}

	// Get the job status
	jobResult, err := azureTextAnalysisClient.GetTextAnalyticsJobResult(context.TODO(), jobID)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Job status: expect=%s, got=%s\n", azuretextanalysis.StatusSucceeded, jobResult.Status) // Wait for job to be completed

	// Extract result (when completed)
	for _, task := range jobResult.Tasks.Items {
		switch task.Kind {
		case azuretextanalysis.LROKindKeyPhraseExtraction:
			result, ok := task.Results.(azuretextanalysis.KeyPhraseResult)
			fmt.Printf("KeyPhraseExtracted (%v): %v\n", ok, result) // Do something with the result
		case azuretextanalysis.LROKindAbstractiveSummarization:
			result, ok := task.Results.(azuretextanalysis.AbstractiveSummarizationResult)
			fmt.Printf("AbstractiveSummarized (%v): %v\n", ok, result) // Do something with the result
		}
	}
}
