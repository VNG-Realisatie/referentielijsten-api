package v1

import "net/http"

const (
	productionHostName string = "referentielijsten-api.vng.cloud"
	// currently we do not have a test server but this is for future proofing
	testHostName string = "referentielijsten-api.test.vng.cloud"
)

// getScheme determines the scheme used for building up a url
// currently the apis do not run on https -> they will always return http when running on kube
// for production this scheme should always be https
func getScheme(req *http.Request) string {
	scheme := "http://"
	if req.TLS != nil || req.Host == productionHostName || req.Host == testHostName {
		scheme = "https://"
	}

	return scheme
}
