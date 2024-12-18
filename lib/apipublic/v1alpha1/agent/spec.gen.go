// Package v1alpha1 provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/oapi-codegen/oapi-codegen/v2 version v2.3.0 DO NOT EDIT.
package v1alpha1

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"net/url"
	"path"
	"strings"

	externalRef0 "github.com/flightctl/flightctl/lib/apipublic/v1alpha1"
	"github.com/getkin/kin-openapi/openapi3"
)

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"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",
}

// GetSwagger returns the content of the embedded swagger specification file
// or error if failed to decode
func decodeSpec() ([]byte, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %w", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
	}

	return buf.Bytes(), nil
}

var rawSpec = decodeSpecCached()

// a naive cached of a decoded swagger spec
func decodeSpecCached() func() ([]byte, error) {
	data, err := decodeSpec()
	return func() ([]byte, error) {
		return data, err
	}
}

// Constructs a synthetic filesystem for resolving external references when loading openapi specifications.
func PathToRawSpec(pathToFile string) map[string]func() ([]byte, error) {
	res := make(map[string]func() ([]byte, error))
	if len(pathToFile) > 0 {
		res[pathToFile] = rawSpec
	}

	pathPrefix := path.Dir(pathToFile)

	for rawPath, rawFunc := range externalRef0.PathToRawSpec(path.Join(pathPrefix, "../openapi.yaml")) {
		if _, ok := res[rawPath]; ok {
			// it is not possible to compare functions in golang, so always overwrite the old value
		}
		res[rawPath] = rawFunc
	}
	return res
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file. The external references of Swagger specification are resolved.
// The logic of resolving external references is tightly connected to "import-mapping" feature.
// Externally referenced files must be embedded in the corresponding golang packages.
// Urls can be supported but this task was out of the scope.
func GetSwagger() (swagger *openapi3.T, err error) {
	resolvePath := PathToRawSpec("")

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, url *url.URL) ([]byte, error) {
		pathToFile := url.String()
		pathToFile = path.Clean(pathToFile)
		getSpec, ok := resolvePath[pathToFile]
		if !ok {
			err1 := fmt.Errorf("path not found: %s", pathToFile)
			return nil, err1
		}
		return getSpec()
	}
	var specData []byte
	specData, err = rawSpec()
	if err != nil {
		return
	}
	swagger, err = loader.LoadFromData(specData)
	if err != nil {
		return
	}
	return
}
