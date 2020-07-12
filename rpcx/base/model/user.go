package model

import (
	"context"
	"fmt"
)

type UserInfo struct {
	ID   int
	Name string
}

type UserReply struct {
	Message string
}

type User int

func (u User) GetInfo(ctx context.Context, ui *UserInfo, r *UserReply) error {
	r.Message = fmt.Sprintf("id: %d, name: %s", ui.ID, ui.Name)
	return nil
}
