package service

import (
	"errors"
	"github.com/yusufbulac/byfood-case/backend/internal/dto"
	"net/url"
	"strings"
)

type UrlService interface {
	ProcessURL(input dto.UrlProcessRequest) (*dto.UrlProcessResponse, error)
}

type urlService struct{}

func NewUrlService() UrlService {
	return &urlService{}
}

func (s *urlService) ProcessURL(input dto.UrlProcessRequest) (*dto.UrlProcessResponse, error) {
	parsedURL, err := url.Parse(input.URL)
	if err != nil {
		return nil, errors.New("invalid URL")
	}

	switch strings.ToLower(input.Operation) {
	case "canonical":
		return &dto.UrlProcessResponse{
			ProcessedURL: getCanonicalURL(parsedURL),
		}, nil
	case "redirection":
		return &dto.UrlProcessResponse{
			ProcessedURL: getRedirectionURL(parsedURL),
		}, nil
	case "all":
		return &dto.UrlProcessResponse{
			ProcessedURL: getCombinedURL(parsedURL),
		}, nil
	default:
		return nil, errors.New("invalid operation type")
	}
}

func getCanonicalURL(u *url.URL) string {
	// Remove query, fragment, trailing slash etc.
	u.RawQuery = ""
	u.Fragment = ""
	u.Path = strings.TrimSuffix(u.Path, "/")
	return u.String()
}

func getRedirectionURL(u *url.URL) string {
	// Example: enforce HTTPS and www
	u.Scheme = "https"
	if !strings.HasPrefix(u.Host, "www.") {
		u.Host = "www." + u.Host
	}
	return u.String()
}

func getCombinedURL(u *url.URL) string {
	// Apply both canonical and redirection transformations
	u = cloneURL(u)
	u = parseURL(getCanonicalURL(u))
	return getRedirectionURL(u)
}

func parseURL(raw string) *url.URL {
	u, _ := url.Parse(raw)
	return u
}

func cloneURL(u *url.URL) *url.URL {
	copied := *u
	return &copied
}
