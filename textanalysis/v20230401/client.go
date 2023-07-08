package v20230401

import (
	"context"
	"fmt"
	"github.com/go-resty/resty/v2"
)

type Client interface {
	AnalyzeTextLanguageDetection(ctx context.Context, input LanguageDetectionAnalysisInput, parameters LanguageDetectionTaskParameters) (*LanguageDetectionResult, error)
	AnalyzeTextEntityRecognition(ctx context.Context, input MultiLanguageAnalysisInput, parameters EntitiesTaskParameters) (*EntitiesResult, error)
	AnalyzeTextKeyPhraseExtraction(ctx context.Context, input MultiLanguageAnalysisInput, parameters KeyPhraseTaskParameters) (*KeyPhraseResult, error)
	AnalyzeTextSentimentAnalysis(ctx context.Context, input MultiLanguageAnalysisInput, parameters SentimentAnalysisTaskParameters) (*SentimentResponse, error)
}

var _ Client = (*client)(nil)

type client struct {
	r *resty.Client
}

func (c client) AnalyzeTextSentimentAnalysis(ctx context.Context, input MultiLanguageAnalysisInput, parameters SentimentAnalysisTaskParameters) (*SentimentResponse, error) {
	body := RequestBody[MultiLanguageAnalysisInput, SentimentAnalysisTaskParameters]{
		Kind:          "KeyPhraseExtraction",
		AnalysisInput: input,
		Parameters:    parameters,
	}
	req, err := c.r.R().
		SetContext(ctx).
		SetQueryParam("api-version", APIVersion).
		SetBody(body).
		SetResult(TaskResponse[SentimentResponse]{}).
		SetError(ErrorResponse{}).
		Post(APIPath)
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
		Kind:          "KeyPhraseExtraction",
		AnalysisInput: input,
		Parameters:    parameters,
	}
	req, err := c.r.R().
		SetContext(ctx).
		SetQueryParam("api-version", APIVersion).
		SetBody(body).
		SetResult(TaskResponse[KeyPhraseResult]{}).
		SetError(ErrorResponse{}).
		Post(APIPath)
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
		Kind:          "EntityRecognition",
		AnalysisInput: input,
		Parameters:    parameters,
	}
	req, err := c.r.R().
		SetContext(ctx).
		SetQueryParam("api-version", APIVersion).
		SetBody(body).
		SetResult(TaskResponse[EntitiesResult]{}).
		SetError(ErrorResponse{}).
		Post(APIPath)
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
		Kind:          "LanguageDetection",
		AnalysisInput: input,
		Parameters:    parameters,
	}
	req, err := c.r.R().
		SetContext(ctx).
		SetQueryParam("api-version", APIVersion).
		SetBody(body).
		SetResult(TaskResponse[LanguageDetectionResult]{}).
		SetError(ErrorResponse{}).
		Post(APIPath)
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

func NewClient(endpoint string, key string) Client {
	return &client{
		r: resty.New().SetBaseURL(endpoint).SetHeader("Ocp-Apim-Subscription-Key", key),
	}
}
