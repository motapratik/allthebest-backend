package resty_client

import (
	"context"
	"net/http"
)

type RestClient interface {
	Post(path string, headers http.Header, payload interface{}) (body []byte, statusCode int, err error)
	PostWithContext(ctx context.Context, path string, headers http.Header, payload interface{}) (body []byte, statusCode int, err error)
	PostFormData(path string, headers http.Header, payload map[string]string) (body []byte, statusCode int, err error)
	Put(path string, headers http.Header, payload interface{}) (body []byte, statusCode int, err error)
	PutWithContext(ctx context.Context, path string, headers http.Header, payload interface{}) (body []byte, statusCode int, err error)
	Get(path string, headers http.Header) (body []byte, statusCode int, err error)
	GetWithContext(ctx context.Context, path string, headers http.Header) (body []byte, statusCode int, err error)
	GetWithQueryParam(path string, headers http.Header, queryParam map[string]string) (body []byte, statusCode int, err error)
	GetWithQueryParamAndContext(ctx context.Context, path string, headers http.Header, queryParam map[string]string) (body []byte, statusCode int, err error)
	Delete(path string, headers http.Header, payload interface{}) (body []byte, statusCode int, err error)
	DeleteWithContext(ctx context.Context, path string, headers http.Header, payload interface{}) (body []byte, statusCode int, err error)
	Patch(path string, headers http.Header, payload interface{}) (body []byte, statusCode int, err error)
	PatchWithContext(ctx context.Context, path string, headers http.Header, payload interface{}) (body []byte, statusCode int, err error)
}
