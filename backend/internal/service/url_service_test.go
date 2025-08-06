package service_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yusufbulac/byfood-case/backend/internal/dto"
	"github.com/yusufbulac/byfood-case/backend/internal/service"
)

func TestUrlService_ProcessURL(t *testing.T) {
	svc := service.NewUrlService()

	tests := []struct {
		name      string
		input     dto.UrlProcessRequest
		wantURL   string
		expectErr bool
		errMsg    string
	}{
		{
			name: "canonical - removes query, fragment, trailing slash",
			input: dto.UrlProcessRequest{
				URL:       "https://example.com/test/path/?utm_source=google#section",
				Operation: "canonical",
			},
			wantURL:   "https://example.com/test/path",
			expectErr: false,
		},
		{
			name: "redirection - adds www, forces https, lowercases",
			input: dto.UrlProcessRequest{
				URL:       "http://EXAMPLE.com/Some/Path",
				Operation: "redirection",
			},
			wantURL:   "https://www.example.com/some/path",
			expectErr: false,
		},
		{
			name: "all - canonical + redirection",
			input: dto.UrlProcessRequest{
				URL:       "http://EXAMPLE.com/Some/Path/?q=test#frag",
				Operation: "all",
			},
			wantURL:   "https://www.example.com/some/path",
			expectErr: false,
		},
		{
			name: "invalid operation",
			input: dto.UrlProcessRequest{
				URL:       "https://example.com/test",
				Operation: "invalid",
			},
			expectErr: true,
			errMsg:    "invalid operation type",
		},
		{
			name: "invalid url format",
			input: dto.UrlProcessRequest{
				URL:       "://not a valid url",
				Operation: "canonical",
			},
			expectErr: true,
			errMsg:    "invalid URL",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := svc.ProcessURL(tt.input)

			if tt.expectErr {
				assert.Nil(t, resp)
				assert.Error(t, err)
				assert.Equal(t, tt.errMsg, err.Error())
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, resp)
				assert.Equal(t, tt.wantURL, resp.ProcessedURL)
			}
		})
	}
}
