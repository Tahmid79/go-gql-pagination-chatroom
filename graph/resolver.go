package graph

import (
	"fmt"
	"gqlpagination1/graph/generated"
	"gqlpagination1/graph/model"
	"math/rand"
	"strconv"
	"time"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	ChatRooms map[string]model.ChatRoom
	Messages  map[string][]model.Message
}

func NewResolver() generated.Config {
	const nChatRooms = 20
	const nMessagesPerChatroom = 100
	r := Resolver{}
	r.ChatRooms = make(map[string]model.ChatRoom, nChatRooms)
	r.Messages = make(map[string][]model.Message, nChatRooms)

	rand.Seed(time.Now().UnixNano())
	for i := 0; i < nChatRooms; i++ {
		id := strconv.Itoa(i + 1)
		mockChatRoom := model.ChatRoom{
			ID:   id,
			Name: fmt.Sprint("Chatroom %d", id),
		}
		r.ChatRooms[id] = mockChatRoom
		r.Messages[id] = make([]model.Message, nMessagesPerChatroom)

		for k := 0; k < nMessagesPerChatroom; k++ {
			id = strconv.Itoa(k + 1)
			text := fmt.Sprintf("Message %d", k)

			mockMessage := model.Message{
				ID:   id,
				Text: &text,
			}

			r.Messages[id][k] = mockMessage

		}

	}

	return generated.Config{
		Resolvers: &r,
	}
}
