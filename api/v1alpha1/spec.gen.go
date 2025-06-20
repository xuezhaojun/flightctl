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
