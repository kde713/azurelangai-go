package v20230401

const APIVersion = "2023-04-01"
const APIPath = "/language/:analyze-text"

type TaskKind string

const (
	TaskKindLanguageDetection        TaskKind = "LanguageDetection"
	TaskKindEntityRecognition        TaskKind = "EntityRecognition"
	TaskKindKeyPhraseExtraction      TaskKind = "KeyPhraseExtraction"
	TaskKindSentimentAnalysis        TaskKind = "SentimentAnalysis"
	TaskKindExtractiveSummarization  TaskKind = "ExtractiveSummarization"
	TaskKindAbstractiveSummarization TaskKind = "AbstractiveSummarization"
)
