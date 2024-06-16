package eu_api_client

import (
	"fmt"
	"testing"
)

// func TestGetResults(t *testing.T) {

// 	bodyParams := make(map[string][]byte)
// 	urlParams := make(map[string][]string)

// 	bodyParams["query"] = GetTopicQuery()
// 	bodyParams["language"] = []byte(`["en"]`)

// 	urlParams["apiKey"] = []string{"SEDIA"}
// 	urlParams["text"] = []string{"***"}
// 	urlParams["pageSize"] = []string{"100"}

// 	apicaller, err := NewAPI_Caller(bodyParams, urlParams)
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}

// 	resultsChan, err := apicaller.GetResults()
// 	if err != nil {
// 		log.Fatalln(err)
// 	}

// 	_ = resultsChan

// 	// for resultArr := range resultsChan {
// 	// 	fmt.Printf("results in this array: %v\n\n\n----------------\n\n\n-----------", len(resultArr))
// 	// 	for _, result := range resultArr {
// 	// 		fmt.Println(result.Metadata["identifier"])
// 	// 	}
// 	// }
// }

// func TestNewAPI_Caller(t *testing.T) {

// 	bodyParams := make(map[string][]byte)
// 	urlParams := make(map[string][]string)

// 	bodyParams["query"] = GetTopicQuery()
// 	bodyParams["language"] = []byte(`["en"]`)

// 	urlParams["apiKey"] = []string{"SEDIA"}
// 	urlParams["text"] = []string{"***"}
// 	urlParams["pageSize"] = []string{"50"}

// 	apicaller, err := NewAPI_Caller(bodyParams, urlParams)
// 	if err != nil {
// 		log.Fatalln(err)
// 		return
// 	}

// 	_ = apicaller

// }

// func TestGetTopicsChan(t *testing.T) {
// 	topics, err := GetTopicsChan(50)
// 	if err != nil {
// 		t.Error(err)
// 	}

// 	_ = topics
// }

func TestGetTopicIDs(t *testing.T) {

	apc := New()

	topicIDsChan := make(chan string)
	errChan := make(chan error)
	done := make(chan struct{})

	go apc.GetTopicIDs(topicIDsChan, errChan, 30)

	go func() {
		for err := range errChan {
			t.Error(err)
		}

		done <- struct{}{}
	}()

	for id := range topicIDsChan {
		fmt.Println("unique topic: " + id)
	}

	<-done
}
