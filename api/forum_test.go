package api_test

import (
	"testing"

	"github.com/acoshift/gowebboard/api"
)

func TestValidateForumCreateRequest(t *testing.T) {
	cases := []struct {
		req      api.ForumCreateRequest
		hasError bool
	}{
		{api.ForumCreateRequest{}, true},
		{api.ForumCreateRequest{Title: "a"}, true},
		{api.ForumCreateRequest{Title: "123"}, true},
		{api.ForumCreateRequest{Title: "test1234"}, false},
		{api.ForumCreateRequest{Title: "121212121"}, false},
	}

	for _, c := range cases {
		err := c.req.Validate()
		if c.hasError && err == nil {
			t.Errorf("expected has error; got nil")
		}
		if !c.hasError && err != nil {
			t.Errorf("expected not error; got %v", err)
		}
	}
}
