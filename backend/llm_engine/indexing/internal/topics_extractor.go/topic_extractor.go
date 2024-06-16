package topics_extractor

import (
	"indexing/internal/eu_api_client"
)

type TopicChunk struct {
	Id           int
	ChunkContent string
	Metadata     map[string]string
}

/*
	Metadata = {
		"topicId" => string
		"budget"  => float
		""
	}
*/

func GetTopicChunkChan() (tcc chan TopicChunk, err error) {

	euClient := eu_api_client.New()
	tcc = make(chan TopicChunk)

	topicIdChan := make(chan string)
	errChan := make(chan error)

	go euClient.GetTopicIDs(topicIdChan, errChan, 30)

	return

}
