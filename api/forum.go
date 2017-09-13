package api

import (
	"fmt"
	"unicode/utf8"
)

// ForumController is the forum controller interface
type ForumController interface {
	Create(*ForumCreateRequest) (*ForumCreateResponse, error)
	Update(*ForumUpdateRequest) (*ForumUpdateResponse, error)
	List(*ForumListRequest) (*ForumListResponse, error)
	Delete()
	Get()
}

// ForumCreateRequest is the create request for forum controller
type ForumCreateRequest struct {
	Title string
}

// Validate validates request
func (req *ForumCreateRequest) Validate() error {
	if len(req.Title) == 0 {
		return fmt.Errorf("title required")
	}
	if utf8.RuneCountInString(req.Title) < 4 {
		return fmt.Errorf("title must have >= 4 chars")
	}
	return nil
}

// ForumCreateResponse is the create response for forum controller
type ForumCreateResponse struct {
	ID int
}

// ForumUpdateRequest is the update request for forum controller
type ForumUpdateRequest struct {
	ID    int
	Title string
}

// ForumUpdateResponse is the update response for forum controller
type ForumUpdateResponse struct {
}

// ForumListRequest is the list request for forum controller
type ForumListRequest struct {
}

// ForumListResponse is the list response for forum controller
type ForumListResponse struct {
	Forums []*ForumItem
}

// ForumItem is the forum list item
type ForumItem struct {
	ID    int
	Title string
}
