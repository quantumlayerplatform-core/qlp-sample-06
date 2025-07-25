package views

import (
	"fmt"
	"path/to/project/frontend/components"
	"path/to/project/frontend/store"
)

// UserView represents the user interface
type UserView struct {
	UserComponent *components.UserComponent
}

// NewUserView creates a new UserView
func NewUserView(client *store.GRPCClient) *UserView {
	userComponent := components.NewUserComponent(client)
	return &UserView{
		UserComponent: userComponent,
	}
}

// RenderUser displays the user information
func (uv *UserView) RenderUser(userID string) {
	user, err := uv.UserComponent.FetchUser(userID)
	if err != nil {
		fmt.Printf("Error rendering user: %v\n", err)
		return
	}
	fmt.Printf("User ID: %s, Name: %s\n", user.Id, user.Name)
}