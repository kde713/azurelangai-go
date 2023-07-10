package v20230401

type LROResult struct {
	// Kind Enumeration of supported Text Analysis long-running operation task results.
	Kind               string      `json:"kind"`
	LastUpdateDateTime string      `json:"lastUpdateDateTime"`
	Results            interface{} `json:"-"`
	Status             string      `json:"status"`
	TaskName           string      `json:"taskName"`
}

type resultsOnly[T any] struct {
	Results T `json:"results"`
}
