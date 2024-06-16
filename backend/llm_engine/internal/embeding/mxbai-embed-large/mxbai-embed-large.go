package mxbaiembedlarge

import (
	"github.com/tmc/langchaingo/embeddings"
	"github.com/tmc/langchaingo/llms/ollama"
)

type MxbaiEmbedLangeWrap struct {
	dimensions     int
	embedingModule *embeddings.EmbedderImpl
}

func New() (*MxbaiEmbedLangeWrap, error) {
	return &MxbaiEmbedLangeWrap{
		dimensions:     1024,
		embedingModule: nil,
	}, nil
}

func (mlm *MxbaiEmbedLangeWrap) CreateModule() (*embeddings.EmbedderImpl, error) {

	llm, err := ollama.New(ollama.WithModel("mxbai-embed-large:latest"))
	if err != nil {
		return nil, err
	}

	e, err := embeddings.NewEmbedder(llm)
	if err != nil {
		return nil, err
	}

	mlm.embedingModule = e

	return e, nil
}

func (mlm *MxbaiEmbedLangeWrap) GetDimensions() int {
	return mlm.dimensions
}
