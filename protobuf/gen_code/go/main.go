package main

import (
	"fmt"
	"log"

	"github.com/RPGGO-AI/l-gos-spec/protobuf/app/go/entity"

	"google.golang.org/protobuf/proto"
)

func main() {
	ch := &entity.Chapter{
		ChapterId: "1",
		Name:      "Introduction",
		CoverUrl:  "https://example.com/cover.jpg",
		InitDialogue: []*entity.Dialogue{
			{Name: "Alice", Index: 1, Message: "Hello!"},
		},
	}

	data, err := proto.Marshal(ch)
	if err != nil {
		log.Fatalf("Failed to marshal: %v", err)
	}

	fmt.Println("Serialized data:", string(data))

	newCh := &entity.Chapter{}
	if err := proto.Unmarshal(data, newCh); err != nil {
		log.Fatalf("Failed to unmarshal: %v", err)
	}

	fmt.Printf("Deserialized Chapter: %+v\n", newCh)
}
