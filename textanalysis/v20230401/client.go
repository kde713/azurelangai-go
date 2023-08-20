package v20230401

import (
	"context"
	"fmt"
	"net/http"

	"github.com/go-resty/resty/v2"
)

type Client interface {
	AnalyzeTextLanguageDetection(ctx context.Context, input LanguageDetectionAnalysisInput, parameters LanguageDetectionTaskParameters) (*LanguageDetectionResult, error)
	AnalyzeTextEntityRecognition(ctx context.Context, input MultiLanguageAnalysisInput, parameters EntitiesTaskParameters) (*EntitiesResult, error)
	AnalyzeTextKeyPhraseExtraction(ctx context.Context, input MultiLanguageAnalysisInput, parameters KeyPhraseTaskParameters) (*KeyPhraseResult, error)
	AnalyzeTextSentimentAnalysis(ctx context.Context, input MultiLanguageAnalysisInput, parameters SentimentAnalysisTaskParameters) (*SentimentResponse, error)
	SubmitTextAnalyticsJob(ctx context.Context, input SubmitJobRequestBody) (string, error)
	GetTextAnalyticsJobResult(ctx context.Context, jobID string) (*JobStatusResponse, error)
}

var _ Client = (*client)(nil)

type client struct {
	r *resty.Client
}

func (c client) SubmitTextAnalyticsJob(ctx context.Context, input SubmitJobRequestBody) (string, error) {
	req, err := c.r.R().
		SetContext(ctx).
		SetQueryParam("api-version", APIVersion).
		SetBody(input).
		SetError(ErrorResponse{}).
		Post(SubmitJobAPIPath)
	if err != nil {
		return "", err
	}
	if req.IsError() {
		errorResp := req.Error().(*ErrorResponse)
		if errorResp == nil {
			return "", fmt.Errorf("error response parse failed: status %d", req.StatusCode())
		}
		return "", &TaskError{Information: errorResp.Error}
	}
	jobLocation := req.Header().Get("Operation-Location")
	if jobLocation == "" {
		return "", fmt.Errorf("missing Operation-Location: status %d", req.StatusCode())
	}
	jobID, err := ParseJobID(jobLocation)
	if err != nil {
		return "", fmt.Errorf("failed to parse jobID: %w", err)
	}
	return jobID, nil
}

func (c client) GetTextAnalyticsJobResult(ctx context.Context, jobID string) (*JobStatusResponse, error) {
	req, err := c.r.R().
		SetContext(ctx).
		SetQueryParam("api-version", APIVersion).
		SetPathParam("jobId", jobID).
		SetResult(JobStatusResponse{}).
		SetError(ErrorResponse{}).
		Get(JobStatusAPIPath)
	if err != nil {
		return nil, err
	}
	if req.IsError() {
		errorResp := req.Error().(*ErrorResponse)
		if errorResp == nil {
			return nil, fmt.Errorf("error response parse failed: status %d", req.StatusCode())
		}
		return nil, &TaskError{Information: errorResp.Error}
	}
	jobResp := req.Result().(*JobStatusResponse)
	if jobResp == nil {
		return nil, fmt.Errorf("job response parse failed: status %d", req.StatusCode())
	}
	return jobResp, nil
}

func (c client) AnalyzeTextSentimentAnalysis(ctx context.Context, input MultiLanguageAnalysisInput, parameters SentimentAnalysisTaskParameters) (*SentimentResponse, error) {
	body := RequestBody[MultiLanguageAnalysisInput, SentimentAnalysisTaskParameters]{
		Kind:          TaskKindSentimentAnalysis,
		AnalysisInput: input,
		Parameters:    parameters,
	}
	req, err := c.r.R().
		SetContext(ctx).
		SetQueryParam("api-version", APIVersion).
		SetBody(body).
		SetResult(TaskResponse[SentimentResponse]{}).
		SetError(ErrorResponse{}).
		Post(AnalyzeTextAPIPath)
	if err != nil {
		return nil, err
	}
	if req.IsError() {
		errorResp := req.Error().(*ErrorResponse)
		if errorResp == nil {
			return nil, fmt.Errorf("error response parse failed: status %d", req.StatusCode())
		}
		return nil, &TaskError{Information: errorResp.Error}
	}
	taskResp := req.Result().(*TaskResponse[SentimentResponse])
	if taskResp == nil {
		return nil, fmt.Errorf("task response parse failed: status %d", req.StatusCode())
	}
	return &taskResp.Results, nil
}

func (c client) AnalyzeTextKeyPhraseExtraction(ctx context.Context, input MultiLanguageAnalysisInput, parameters KeyPhraseTaskParameters) (*KeyPhraseResult, error) {
	body := RequestBody[MultiLanguageAnalysisInput, KeyPhraseTaskParameters]{
		Kind:          TaskKindKeyPhraseExtraction,
		AnalysisInput: input,
		Parameters:    parameters,
	}
	req, err := c.r.R().
		SetContext(ctx).
		SetQueryParam("api-version", APIVersion).
		SetBody(body).
		SetResult(TaskResponse[KeyPhraseResult]{}).
		SetError(ErrorResponse{}).
		Post(AnalyzeTextAPIPath)
	if err != nil {
		return nil, err
	}
	if req.IsError() {
		errorResp := req.Error().(*ErrorResponse)
		if errorResp == nil {
			return nil, fmt.Errorf("error response parse failed: status %d", req.StatusCode())
		}
		return nil, &TaskError{Information: errorResp.Error}
	}
	taskResp := req.Result().(*TaskResponse[KeyPhraseResult])
	if taskResp == nil {
		return nil, fmt.Errorf("task response parse failed: status %d", req.StatusCode())
	}
	return &taskResp.Results, nil
}

func (c client) AnalyzeTextEntityRecognition(ctx context.Context, input MultiLanguageAnalysisInput, parameters EntitiesTaskParameters) (*EntitiesResult, error) {
	body := RequestBody[MultiLanguageAnalysisInput, EntitiesTaskParameters]{
		Kind:          TaskKindEntityRecognition,
		AnalysisInput: input,
		Parameters:    parameters,
	}
	req, err := c.r.R().
		SetContext(ctx).
		SetQueryParam("api-version", APIVersion).
		SetBody(body).
		SetResult(TaskResponse[EntitiesResult]{}).
		SetError(ErrorResponse{}).
		Post(AnalyzeTextAPIPath)
	if err != nil {
		return nil, err
	}
	if req.IsError() {
		errorResp := req.Error().(*ErrorResponse)
		if errorResp == nil {
			return nil, fmt.Errorf("error response parse failed: status %d", req.StatusCode())
		}
		return nil, &TaskError{Information: errorResp.Error}
	}
	taskResp := req.Result().(*TaskResponse[EntitiesResult])
	if taskResp == nil {
		return nil, fmt.Errorf("task response parse failed: status %d", req.StatusCode())
	}
	return &taskResp.Results, nil
}

func (c client) AnalyzeTextLanguageDetection(ctx context.Context, input LanguageDetectionAnalysisInput, parameters LanguageDetectionTaskParameters) (*LanguageDetectionResult, error) {
	body := RequestBody[LanguageDetectionAnalysisInput, LanguageDetectionTaskParameters]{
		Kind:          TaskKindLanguageDetection,
		AnalysisInput: input,
		Parameters:    parameters,
	}
	req, err := c.r.R().
		SetContext(ctx).
		SetQueryParam("api-version", APIVersion).
		SetBody(body).
		SetResult(TaskResponse[LanguageDetectionResult]{}).
		SetError(ErrorResponse{}).
		Post(AnalyzeTextAPIPath)
	if err != nil {
		return nil, err
	}
	if req.IsError() {
		errorResp := req.Error().(*ErrorResponse)
		if errorResp == nil {
			return nil, fmt.Errorf("error response parse failed: status %d", req.StatusCode())
		}
		return nil, &TaskError{Information: errorResp.Error}
	}
	taskResp := req.Result().(*TaskResponse[LanguageDetectionResult])
	if taskResp == nil {
		return nil, fmt.Errorf("task response parse failed: status %d", req.StatusCode())
	}
	return &taskResp.Results, nil
}

func NewClient(endpoint string, key string, optAppliers ...Option) Client {
	o := options{}
	for _, applier := range optAppliers {
		applier(&o)
	}

	r := resty.New().SetBaseURL(endpoint).SetHeader("Ocp-Apim-Subscription-Key", key)

	// Handle retry option
	if o.retryCount != 0 {
		r = r.SetRetryCount(o.retryCount).SetRetryWaitTime(o.retryWaitTime).SetRetryMaxWaitTime(o.retryMaxWaitTime)
		r = r.AddRetryCondition(func(r *resty.Response, err error) bool {
			statusCode := r.StatusCode()
			return statusCode == http.StatusTooManyRequests || statusCode == http.StatusInternalServerError || statusCode == http.StatusBadGateway || statusCode == http.StatusServiceUnavailable
		})
	}

	return &client{
		r: r,
	}
}
