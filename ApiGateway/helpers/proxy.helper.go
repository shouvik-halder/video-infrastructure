package helpers

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func ProxyRequest(
	w http.ResponseWriter,
	req *http.Request,
	targetURL string,
	targetPath string,
) {
	target, err := url.Parse(targetURL)
	if err != nil {
		http.Error(w, `{"message":"invalid service url"}`, http.StatusInternalServerError)
		return
	}

	proxy := &httputil.ReverseProxy{
		Rewrite: func(proxyReq *httputil.ProxyRequest) {
			proxyReq.SetURL(target)
			proxyReq.Out.URL.Path = targetPath
			proxyReq.SetXForwarded()

			// Mark as coming from the gateway
			proxyReq.Out.Header.Set("X-Forwarded-By", "vi-api-gateway")

			// Strip headers that should never come from external callers
			proxyReq.Out.Header.Del("X-Internal-Secret")
			proxyReq.Out.Header.Del("X-User-ID")
		},
		ErrorHandler: func(w http.ResponseWriter, r *http.Request, err error) {
			fmt.Printf("proxy error → %s%s: %v\n", targetURL, targetPath, err)
			w.Header().Set("Content-Type", "application/json")
			http.Error(w, `{"message":"upstream service unavailable"}`, http.StatusBadGateway)
		
		},
	}

	proxy.ServeHTTP(w, req)
}