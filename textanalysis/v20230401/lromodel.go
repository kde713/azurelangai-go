package v20230401

import (
	"encoding/json"
)

type JobStatus string

const (
	StatusCancelled          JobStatus = "cancelled"
	StatusCancelling         JobStatus = "cancelling"
	StatusFailed             JobStatus = "failed"
	StatusNotStarted         JobStatus = "notStarted"
	StatusPartiallyCompleted JobStatus = "partiallyCompleted"
	StatusRunning            JobStatus = "running"
	StatusSucceeded          JobStatus = "succeeded"
)

type LROKind string

const (
	LROKindEntityRecognition        LROKind = "EntityRecognitionLROResults"
	LROKindKeyPhraseExtraction      LROKind = "KeyPhraseExtractionLROResults"
	LROKindSentimentAnalysis        LROKind = "SentimentAnalysisLROResults"
	LROKindExtractiveSummarization  LROKind = "ExtractiveSummarizationLROResults"
	LROKindAbstractiveSummarization LROKind = "AbstractiveSummarizationLROResults"
)

type commonLROResult struct {
	Kind               LROKind   `json:"kind"`
	LastUpdateDateTime string    `json:"lastUpdateDateTime"`
	Status             JobStatus `json:"status"`
	TaskName           string    `json:"taskName"`
}

type resultsOnly[T any] struct {
	Results T `json:"results"`
}

var _ json.Unmarshaler = (*LROResult)(nil)

type LROResult struct {
	// Kind Enumeration of supported Text Analysis long-running operation task results.
	Kind               LROKind
	LastUpdateDateTime string
	Results            interface{}
	Status             JobStatus
	TaskName           string
}

func (r *LROResult) UnmarshalJSON(bytes []byte) error {
	// Map common keys
	var common commonLROResult
	if err := json.Unmarshal(bytes, &common); err != nil {
		return err
	}
	r.Kind = common.Kind
	r.LastUpdateDateTime = common.LastUpdateDateTime
	r.Status = common.Status
	r.TaskName = common.TaskName

	if r.Status != StatusSucceeded {
		// If the task is not complete, return immediately
		return nil
	}

	switch r.Kind {
	case LROKindEntityRecognition:
		var results resultsOnly[EntitiesResult]
		if err := json.Unmarshal(bytes, &results); err != nil {
			return err
		}
		r.Results = results.Results
	case LROKindKeyPhraseExtraction:
		var results resultsOnly[KeyPhraseResult]
		if err := json.Unmarshal(bytes, &results); err != nil {
			return err
		}
		r.Results = results.Results
	case LROKindSentimentAnalysis:
		var results resultsOnly[SentimentResponse]
		if err := json.Unmarshal(bytes, &results); err != nil {
			return err
		}
		r.Results = results.Results
	case LROKindExtractiveSummarization:
		var results resultsOnly[ExtractiveSummarizationResult]
		if err := json.Unmarshal(bytes, &results); err != nil {
			return err
		}
		r.Results = results.Results
	case LROKindAbstractiveSummarization:
		var results resultsOnly[AbstractiveSummarizationResult]
		if err := json.Unmarshal(bytes, &results); err != nil {
			return err
		}
		r.Results = results.Results
	}
	return nil
}

type Tasks struct {
	Completed  int         `json:"completed"`
	Failed     int         `json:"failed"`
	InProgress int         `json:"inProgress"`
	Items      []LROResult `json:"items"`
	Total      int         `json:"total"`
}

type TaskRequest struct {
	// Kind Enumeration of supported long-running Text Analysis tasks.
	Kind       TaskKind    `json:"kind"`
	Parameters interface{} `json:"parameters"`
	TaskName   string      `json:"taskName"`
}

type SubmitJobRequestBody struct {
	AnalysisInput MultiLanguageAnalysisInput `json:"analysisInput"`
	Tasks         []TaskRequest              `json:"tasks"`
	DisplayName   string                     `json:"displayName"`
}

type JobStatusResponse struct {
	CreatedDateTime    string             `json:"createdDateTime"`
	DisplayName        string             `json:"displayName"`
	Errors             []ErrorInformation `json:"errors"`
	ExpirationDateTime string             `json:"expirationDateTime"`
	JobID              string             `json:"jobId"`
	LastUpdateDateTime string             `json:"lastUpdateDateTime"`
	NextLink           string             `json:"nextLink"`
	Status             JobStatus          `json:"status"`
	Tasks              Tasks              `json:"tasks"`
}
