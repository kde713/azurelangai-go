package v20230401

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

type LROResult struct {
	// Kind Enumeration of supported Text Analysis long-running operation task results.
	Kind               TaskKind    `json:"kind"`
	LastUpdateDateTime string      `json:"lastUpdateDateTime"`
	Results            interface{} `json:"-"`
	Status             string      `json:"status"`
	TaskName           string      `json:"taskName"`
}

type resultsOnly[T any] struct {
	Results T `json:"results"`
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
