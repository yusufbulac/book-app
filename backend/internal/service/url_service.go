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

// Removes query parameters, fragments and trailing slashes
func getCanonicalURL(u *url.URL) string {
	u.RawQuery = ""
	u.Fragment = ""
	u.Path = strings.TrimSuffix(u.Path, "/")
	return u.String()
}

// Lowercase host/path and enforce https + www.
func getRedirectionURL(u *url.URL) string {
	u.Scheme = "https"

	// Ensure lowercase host and add www.
	host := strings.ToLower(u.Host)
	if !strings.HasPrefix(host, "www.") {
		host = "www." + host
	}
	u.Host = host

	// Lowercase path
	u.Path = strings.ToLower(u.Path)

	return u.String()
}

// Apply canonical then redirection
func getCombinedURL(u *url.URL) string {
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
