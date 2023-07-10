package v20230401

type RequestBody[AnalysisInput any, Parameters any] struct {
	// Kind Enumeration of supported Text Analysis tasks.
	Kind          TaskKind      `json:"kind"`
	AnalysisInput AnalysisInput `json:"analysisInput"`
	// Parameters Supported parameters for requesting analysis task.
	Parameters Parameters `json:"parameters"`
}

type LanguageInput struct {
	CountryHint string `json:"countryHint,omitempty"`
	// ID Unique, non-empty document identifier.
	ID   string `json:"id"`
	Text string `json:"text"`
}

type LanguageDetectionAnalysisInput struct {
	Documents []LanguageInput `json:"documents"`
}

type LanguageDetectionTaskParameters struct {
	LoggingOptOut bool   `json:"loggingOptOut,omitempty"`
	ModelVersion  string `json:"modelVersion,omitempty"`
}

type MultiLanguageInput struct {
	// ID A unique, non-empty document identifier.
	ID string `json:"id"`
	// Language (Optional) This is the 2 letter ISO 639-1 representation of a language. For example, use "en" for English; "es" for Spanish etc. If not set, use "en" for English as default.
	Language string `json:"language,omitempty"`
	// Text The input text to process.
	Text string `json:"text"`
}

type MultiLanguageAnalysisInput struct {
	// Documents Contains an input document to be analyzed by the service.
	Documents []MultiLanguageInput `json:"documents"`
}

type EntitiesTaskParameters struct {
	LoggingOptOut bool   `json:"loggingOptOut,omitempty"`
	ModelVersion  string `json:"modelVersion,omitempty"`
	// StringIndexType Specifies the method used to interpret string offsets. Defaults to Text Elements (Graphemes) according to Unicode v8.0.0. For additional information see https://aka.ms/text-analytics-offsets.
	StringIndexType string `json:"stringIndexType,omitempty"`
}

type KeyPhraseTaskParameters struct {
	LoggingOptOut bool   `json:"loggingOptOut,omitempty"`
	ModelVersion  string `json:"modelVersion,omitempty"`
}

type SentimentAnalysisTaskParameters struct {
	LoggingOptOut bool   `json:"loggingOptOut,omitempty"`
	ModelVersion  string `json:"modelVersion,omitempty"`
	OpinionMining bool   `json:"opinionMining,omitempty"`
	// StringIndexType Specifies the method used to interpret string offsets. Defaults to Text Elements (Graphemes) according to Unicode v8.0.0. For additional information see https://aka.ms/text-analytics-offsets.
	StringIndexType string `json:"stringIndexType,omitempty"`
}

type ExtractiveSummarizationTaskParameters struct {
	LoggingOptOut bool   `json:"loggingOptOut,omitempty"`
	ModelVersion  string `json:"modelVersion,omitempty"`
	SentenceCount int    `json:"sentenceCount,omitempty"`
	// SortBy The sorting criteria to use for the results of Extractive Summarization.
	// "Offset" (Default): Indicates that results should be sorted in order of appearance in the text.
	// "Rank": Indicates that results should be sorted in order of importance (i.e. rank score) according to the model.
	SortBy string `json:"sortBy,omitempty"`
	// StringIndexType Specifies the method used to interpret string offsets. Defaults to Text Elements (Graphemes) according to Unicode v8.0.0. For additional information see https://aka.ms/text-analytics-offsets.
	StringIndexType string `json:"stringIndexType,omitempty"`
}

type AbstractiveSummarizationTaskParameters struct {
	LoggingOptOut bool   `json:"loggingOptOut,omitempty"`
	ModelVersion  string `json:"modelVersion,omitempty"`
	SentenceCount int    `json:"sentenceCount,omitempty"`
	// StringIndexType Specifies the method used to interpret string offsets. Defaults to Text Elements (Graphemes) according to Unicode v8.0.0. For additional information see https://aka.ms/text-analytics-offsets.
	StringIndexType string `json:"stringIndexType,omitempty"`
}

type InputError struct {
	// Error Error encountered.
	Error ErrorInformation `json:"error"`
	// ID The ID of the input.
	ID string `json:"id"`
}

type ErrorInformation struct {
	// Code One of a server-defined set of error codes.
	Code string `json:"code"`
	// Message A human-readable representation of the error.
	Message string `json:"message"`
	// Target The target of the error.
	Target string `json:"target"`
}

type DocumentWarning struct {
	// Code Error code.
	Code string `json:"code"`
	// Message Warning message.
	Message string `json:"message"`
	// TargetRef A JSON pointer reference indicating the target object.
	TargetRef string `json:"targetRef"`
}

type ErrorResponse struct {
	// Error The error object.
	Error ErrorInformation `json:"error"`
}

type TaskResponse[Results any] struct {
	// Kind Enumeration of supported Text Analysis task results.
	Kind    TaskKind `json:"kind"`
	Results Results  `json:"results"`
}

type DetectedLanguage struct {
	// ConfidenceScore A confidence score between 0 and 1. Scores close to 1 indicate 100% certainty that the identified language is true.
	ConfidenceScore float64 `json:"confidenceScore"`
	// ISO6391Name A two letter representation of the detected language according to the ISO 639-1 standard (e.g. en, fr).
	ISO6391Name string `json:"iso6391Name"`
	// Name Long name of a detected language (e.g. English, French).
	Name string `json:"name"`
	// Script Identifies the script of the input document.
	Script string `json:"script"`
}

type LanguageDetectionDocumentResult struct {
	// DetectedLanguage Detected Language.
	DetectedLanguage DetectedLanguage `json:"detectedLanguage"`
	// ID Unique, non-empty document identifier.
	ID string `json:"id"`
	// Warnings Warnings encountered while processing document.
	Warnings []DocumentWarning `json:"warnings"`
}

type LanguageDetectionResult struct {
	// Documents Response by document
	Documents []LanguageDetectionDocumentResult `json:"documents"`
	// Errors Errors by document id.
	Errors []InputError `json:"errors"`
	// ModelVersion This field indicates which model is used for scoring.
	ModelVersion string `json:"modelVersion"`
}

type DocumentError struct {
	// Error Error encountered.
	Error ErrorInformation `json:"error"`
	// ID The ID of the input document.
	ID string `json:"id"`
}

type Entity struct {
	// Category Entity type.
	Category string `json:"category"`
	// ConfidenceScore Confidence score between 0 and 1 of the extracted entity.
	ConfidenceScore float64 `json:"confidenceScore"`
	// Length Length for the entity text. Use of different 'stringIndexType' values can affect the length returned.
	Length int `json:"length"`
	// Offset Start position for the entity text. Use of different 'stringIndexType' values can affect the offset returned.
	Offset int `json:"offset"`
	// SubCategory (Optional) Entity sub type.
	SubCategory string `json:"subcategory"`
	// Text Entity text as appears in the request.
	Text string `json:"text"`
}

type Documents struct {
	// Entities Recognized entities in the document.
	Entities []Entity `json:"entities"`
	// ID Unique, non-empty document identifier.
	ID string `json:"id"`
	// Warnings Warnings encountered while processing document.
	Warnings []DocumentWarning `json:"warnings"`
}

type EntitiesResult struct {
	// Documents Response by document
	Documents []Documents `json:"documents"`
	// Errors Errors by document id.
	Errors []DocumentError `json:"errors"`
	// ModelVersion This field indicates which model is used for scoring.
	ModelVersion string `json:"modelVersion"`
}

type KeyPhraseResult struct {
	// Documents Response by document
	Documents []Documents `json:"documents"`
	// Errors Errors by document id.
	Errors []DocumentError `json:"errors"`
	// ModelVersion This field indicates which model is used for scoring.
	ModelVersion string `json:"modelVersion"`
}

type SentimentResponse struct {
	// Documents Response by document
	Documents []Documents `json:"documents"`
	// Errors Errors by document id.
	Errors []DocumentError `json:"errors"`
	// ModelVersion This field indicates which model is used for scoring.
	ModelVersion string `json:"modelVersion"`
}

type SummaryContext struct {
	// Length The length of the context. Use of different 'stringIndexType' values can affect the length returned.
	Length int `json:"length"`
	// Offset Start position for the context. Use of different 'stringIndexType' values can affect the offset returned.
	Offset int `json:"offset"`
}

type AbstractiveSummary struct {
	// Contexts The context list of the summary.
	Contexts []SummaryContext `json:"contexts"`
	// Text The text of the summary.
	Text string `json:"text"`
}

type AbstractiveSummaryDocumentResult struct {
	// ID Unique, non-empty document identifier.
	ID string `json:"id"`
	// Summaries A list of abstractive summaries.
	Summaries []AbstractiveSummary `json:"summaries"`
	// Warnings Warnings encountered while processing document.
	Warnings []DocumentWarning `json:"warnings"`
}

type AbstractiveSummarizationResult struct {
	// Documents Response by document
	Documents []AbstractiveSummaryDocumentResult `json:"documents"`
	// Errors Errors by document id.
	Errors []DocumentError `json:"errors"`
	// ModelVersion This field indicates which model is used for scoring.
	ModelVersion string `json:"modelVersion"`
}

type ExtractedSummarySentence struct {
	// Length The length of the sentence.
	Length int `json:"length"`
	// Offset The sentence offset from the start of the document, based on the value of the parameter StringIndexType.
	Offset int `json:"offset"`
	// RankScore A double value representing the relevance of the sentence within the summary. Higher values indicate higher importance.
	RankScore float64 `json:"rankScore"`
	// Text The extracted sentence text.
	Text string `json:"text"`
}

type ExtractedSummaryDocumentResult struct {
	// ID Unique, non-empty document identifier.
	ID string `json:"id"`
	// Sentences A ranked list of sentences representing the extracted summary.
	Sentences []ExtractedSummarySentence `json:"sentences"`
	// Warnings Warnings encountered while processing document.
	Warnings []DocumentWarning `json:"warnings"`
}

type ExtractiveSummarizationResult struct {
	// Documents Response by document
	Documents []ExtractedSummaryDocumentResult `json:"documents"`
	// Errors Errors by document id.
	Errors []DocumentError `json:"errors"`
	// ModelVersion This field indicates which model is used for scoring.
	ModelVersion string `json:"modelVersion"`
}
