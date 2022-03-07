package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"encoding/base64"
	"errors"
	"gqlpagination1/graph/generated"
	"gqlpagination1/graph/model"
)

func (r *chatRoomResolver) MessagesConnection(ctx context.Context, obj *model.ChatRoom, first *int, after *string) (*model.MessagesConnection, error) {
	var decodedCursor string

	if after != nil {
		b, err := base64.StdEncoding.DecodeString(*after)
		if err != nil {
			return nil, err
		}
		decodedCursor = string(b)
	}

	edges := make([]*model.MessagesEdge, *first)
	count := 0
	currentPage := false

	if decodedCursor == "" {
		currentPage = true
	}

	hasNextPage := false

	for i, v := range r.Messages[obj.ID] {

		if v.ID == decodedCursor {
			currentPage = true
		}

		if currentPage && count < *first {
			edges[count] = &model.MessagesEdge{
				Cursor: base64.StdEncoding.EncodeToString([]byte(v.ID)),
				Node:   &v,
			}
			count++
		}

		if count == *first && i < len(r.Messages[obj.ID]) {
			hasNextPage = true
		}

	}

	pageInfo := model.PageInfo{
		StartCursor: base64.StdEncoding.EncodeToString([]byte(edges[0].Node.ID)),
		EndCursor:   base64.StdEncoding.EncodeToString([]byte(edges[count-1].Node.ID)),
		HasNextPage: &hasNextPage,
	}

	mc := model.MessagesConnection{
		Edges:    edges[:count],
		PageInfo: &pageInfo,
	}

	return &mc, nil
}

func (r *queryResolver) ChatRoom(ctx context.Context, id string) (*model.ChatRoom, error) {
	if t, ok := r.ChatRooms[id]; ok {
		return &t, nil
	}
	return nil, errors.New("chat room not found")
}

// ChatRoom returns generated.ChatRoomResolver implementation.
func (r *Resolver) ChatRoom() generated.ChatRoomResolver { return &chatRoomResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type chatRoomResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
