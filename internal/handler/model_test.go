package handler_test

import (
	"testing"

	"github.com/saisai/newsapi/internal/handler"
	"github.com/stretchr/testify/assert"
)

func TestNewsPostReqBody_Validate(t *testing.T) {
	testCases := []struct {
		name        string
		req         handler.NewsPostReqBody
		expectedErr string
	}{
		{
			name:        "author empty",
			req:         handler.NewsPostReqBody{},
			expectedErr: "author is empty",
		},
		{
			name: "title empty",
			req: handler.NewsPostReqBody{
				Author: "test-author",
			},
			expectedErr: "title is empty",
		},
		{
			name: "summary empty",
			req: handler.NewsPostReqBody{
				Author: "test-author",
				Title:  "test-title",
			},
			expectedErr: "summary is empty",
		},
		{
			name: "time invalid",
			req: handler.NewsPostReqBody{
				Author:    "test-author",
				Title:     "test-title",
				Summary:   "test-summary",
				CreatedAt: "invalid",
			},
			expectedErr: `parsing time "invalid"`,
		},
		{
			name: "source invalid",
			req: handler.NewsPostReqBody{
				Author:    "test-author",
				Title:     "test-title",
				Summary:   "test-summary",
				CreatedAt: "2024-04-07T05:13:27+00:00",
			},
			expectedErr: "source is empty",
		},
		{
			name: "tags empty",
			req: handler.NewsPostReqBody{
				Author:    "test-author",
				Title:     "test-title",
				Summary:   "test-summary",
				CreatedAt: "2024-04-07T05:13:27+00:00",
				Source:    "https://test-news.com",
			},
			expectedErr: "tags cannot be empty",
		},
		{
			name: "validate",
			req: handler.NewsPostReqBody{
				Author:    "test-author",
				Title:     "test-title",
				Summary:   "test-summary",
				CreatedAt: "2024-04-07T05:13:27+00:00",
				Source:    "https://test-news.com",
				Tags:      []string{"test-tag"},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := tc.req.Validate()

			if tc.expectedErr != "" {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tc.expectedErr)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
