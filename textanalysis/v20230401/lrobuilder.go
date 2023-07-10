package v20230401

import (
	"errors"
)

type LROBuilder struct {
	body SubmitJobRequestBody
}

func NewJobBuilder(displayName string, input MultiLanguageAnalysisInput) *LROBuilder {
	return &LROBuilder{body: SubmitJobRequestBody{
		AnalysisInput: input,
		Tasks:         make([]TaskRequest, 0),
		DisplayName:   displayName,
	}}
}

func (b *LROBuilder) Build() (*SubmitJobRequestBody, error) {
	if len(b.body.Tasks) == 0 {
		return nil, errors.New("no tasks added to job")
	}
	return &b.body, nil
}

func (b *LROBuilder) AddEntityRecognitionTask(taskName string, parameters EntitiesTaskParameters) {
	b.body.Tasks = append(b.body.Tasks, TaskRequest{
		Kind:       TaskKindEntityRecognition,
		Parameters: parameters,
		TaskName:   taskName,
	})
}

func (b *LROBuilder) AddKeyPhraseExtractionTask(taskName string, parameters KeyPhraseTaskParameters) {
	b.body.Tasks = append(b.body.Tasks, TaskRequest{
		Kind:       TaskKindKeyPhraseExtraction,
		Parameters: parameters,
		TaskName:   taskName,
	})
}

func (b *LROBuilder) AddSentimentAnalysisTask(taskName string, parameters SentimentAnalysisTaskParameters) {
	b.body.Tasks = append(b.body.Tasks, TaskRequest{
		Kind:       TaskKindSentimentAnalysis,
		Parameters: parameters,
		TaskName:   taskName,
	})
}

func (b *LROBuilder) AddAbstractiveSummarizationTask(taskName string, parameters AbstractiveSummarizationTaskParameters) {
	b.body.Tasks = append(b.body.Tasks, TaskRequest{
		Kind:       TaskKindAbstractiveSummarization,
		Parameters: parameters,
		TaskName:   taskName,
	})
}

func (b *LROBuilder) AddExtractiveSummarizationTask(taskName string, parameters ExtractiveSummarizationTaskParameters) {
	b.body.Tasks = append(b.body.Tasks, TaskRequest{
		Kind:       TaskKindExtractiveSummarization,
		Parameters: parameters,
		TaskName:   taskName,
	})
}
