package middleware

import (
	"net/http"

	"github.com/lius-new/blog-backend/rpc/analyzer/analyzerclient"
)

type AnalyzerMiddleware struct {
	analyzer analyzerclient.Analyzer
}

func NewAnalyzerMiddleware(analyzer analyzerclient.Analyzer) *AnalyzerMiddleware {
	return &AnalyzerMiddleware{
		analyzer,
	}
}

func (m *AnalyzerMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO generate middleware implement function, delete after code implementation

		// Passthrough to next handler if need
		next(w, r)
	}
}
