package adapter

import (
	"context"
	"errors"
	"net/http"
)

type key string

const (
	ctxResWriterKey = key("resWriter")
)

func SetResWriter(parents context.Context, w http.ResponseWriter) context.Context {
	return context.WithValue(parents, ctxResWriterKey, w)
}

var errContext = errors.New("not found http response writer")

func GetResWriter(ctx context.Context) (http.ResponseWriter, error) {
	v := ctx.Value(ctxResWriterKey)
	w, ok := v.(http.ResponseWriter)

	if !ok {
		return nil, errContext
	}

	return w, nil
}
