package v20230401

const APIVersion = "2023-04-01"
const AnalyzeTextAPIPath = "/language/:analyze-text"
const SubmitJobAPIPath = "/language/analyze-text/jobs"
const JobStatusAPIPath = "/language/analyze-text/jobs/{jobId}"

type TaskKind string

const (
	TaskKindLanguageDetection        TaskKind = "LanguageDetection"
	TaskKindEntityRecognition        TaskKind = "EntityRecognition"
	TaskKindKeyPhraseExtraction      TaskKind = "KeyPhraseExtraction"
	TaskKindSentimentAnalysis        TaskKind = "SentimentAnalysis"
	TaskKindExtractiveSummarization  TaskKind = "ExtractiveSummarization"
	TaskKindAbstractiveSummarization TaskKind = "AbstractiveSummarization"
)

type Sentiment string

const (
	SentimentPositive Sentiment = "positive"
	SentimentNeutral  Sentiment = "neutral"
	SentimentNegative Sentiment = "negative"
	SentimentMixed    Sentiment = "mixed"
)
