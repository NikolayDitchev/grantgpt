package main

import (
	"context"
	"net/url"

	"github.com/tmc/langchaingo/embeddings"
	"github.com/tmc/langchaingo/llms/ollama"
	"github.com/tmc/langchaingo/vectorstores/qdrant"
)

func IndexQdrantLlama3() error {

	llm, err := ollama.New(ollama.WithModel("mxbai-embed-large:latest"))
	if err != nil {
		return err
	}

	e, err := embeddings.NewEmbedder(llm)
	if err != nil {
		return err
	}

	ctx := context.Background()

	url, err := url.Parse("http://localhost:6333")
	if err != nil {
		return err
	}

	store, err := qdrant.New(qdrant.WithURL(*url), qdrant.WithEmbedder(e), qdrant.WithCollectionName("TOPIC_DETAILS"))
	if err != nil {
		return err
	}

	_ = ctx

	_ = store

	return nil
}
