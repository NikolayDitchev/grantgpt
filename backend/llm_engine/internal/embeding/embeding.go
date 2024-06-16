package embeding

import (
	"github.com/tmc/langchaingo/embeddings"
)

type Embedder interface {
	CreateModule() *embeddings.EmbedderImpl
	GetDimensions() int
}
