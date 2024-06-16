package main

import "testing"

func TestIndexQdrantLlama3(t *testing.T) {

	err := IndexQdrantLlama3()
	if err != nil {
		t.Error(err)
	}
}

func TestCreateCollection(t *testing.T) {

	CreateQdrantCollection()

}
