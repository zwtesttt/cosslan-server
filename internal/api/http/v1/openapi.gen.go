// Package v1 provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen/v2 version v2.1.0 DO NOT EDIT.
package v1

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"net/url"
	"path"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/labstack/echo/v4"
)

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Say Hello
	// (GET /api/v1/line/)
	Line(ctx echo.Context) error
	// 创建网络
	// (POST /api/v1/network)
	CreateNetwork(ctx echo.Context) error
	// Say Hello
	// (GET /api/v1/node/)
	Node(ctx echo.Context) error
	// 用户登录
	// (POST /api/v1/user/login)
	Login(ctx echo.Context) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// Line converts echo context to params.
func (w *ServerInterfaceWrapper) Line(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.Line(ctx)
	return err
}

// CreateNetwork converts echo context to params.
func (w *ServerInterfaceWrapper) CreateNetwork(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.CreateNetwork(ctx)
	return err
}

// Node converts echo context to params.
func (w *ServerInterfaceWrapper) Node(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.Node(ctx)
	return err
}

// Login converts echo context to params.
func (w *ServerInterfaceWrapper) Login(ctx echo.Context) error {
	var err error

	ctx.Set(BearerAuthScopes, []string{})

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.Login(ctx)
	return err
}

// This is a simple interface which specifies echo.Route addition functions which
// are present on both echo.Echo and echo.Group, since we want to allow using
// either of them for path registration
type EchoRouter interface {
	CONNECT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	DELETE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	GET(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	HEAD(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	OPTIONS(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PATCH(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	POST(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PUT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	TRACE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
}

// RegisterHandlers adds each server route to the EchoRouter.
func RegisterHandlers(router EchoRouter, si ServerInterface) {
	RegisterHandlersWithBaseURL(router, si, "")
}

// Registers handlers, and prepends BaseURL to the paths, so that the paths
// can be served under a prefix.
func RegisterHandlersWithBaseURL(router EchoRouter, si ServerInterface, baseURL string) {

	wrapper := ServerInterfaceWrapper{
		Handler: si,
	}

	router.GET(baseURL+"/api/v1/line/", wrapper.Line)
	router.POST(baseURL+"/api/v1/network", wrapper.CreateNetwork)
	router.GET(baseURL+"/api/v1/node/", wrapper.Node)
	router.POST(baseURL+"/api/v1/user/login", wrapper.Login)

}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"H4sIAAAAAAAC/8xUQWsUSxD+K0u9d5zd2c0jj/fmlghiJARJFA9hkc5M7WwnM92d7p4kQxjQW/AgCehB",
	"jOhRD4I3wZCfM9m/IdWzm51xYkASIYeF3pquqq+++vo7hFCmSgoU1kBwCCYcY8rc8Z5GZnEN7b7UOxRQ",
	"WirUlqP7zKJIo3HHCE2oubJcCghgcn4yOTstT7+W75+DB3jAUpUgBDD4f6E3+Pe/3qDXBw9srihorOYi",
	"hsIDk20JtM9SZnbaRcsvx5Pzk4tXnycfXzSKLiwu9ma/K8oWlxG5tY2hpUbraJQUBtszhTJy0WkKFxZj",
	"1ODBQTeWXYp2zQ5XXelgsaSrJN3REFidIV2TKbeYKptDMGKJwcKDiFlWKzrHkZq4Fp8ivkmvq4Z9YlCv",
	"ypiLddzN0Nj20JgyntBhJHXKLATTiHeLyDxQzJh9qaNbn1jjbsY1RhBsXgK/bDZsMUJCwzDT3OYbpPWK",
	"g2VkGvVSZsf0b8v9uz+j4+HTx+BVL4MqVV/n9IytVVBQYS5G0s3HrdNmKI1JmOgsPVoBD/ZQm0rMg16/",
	"16flSIWCKQ4B/ONChNyOHSKfKe7vDfyEC/QpEKPbHW2OET8rEQSwygUCcVAp2mUu9PuVloVF4XKYUgkP",
	"XZa/bQjC7KG3t+EGab69pY7JwhCNGWVJZ9aqYjJLU6ZzCGCD5Z0HmCSSeGGxoW0QdBjSvdkwomYm0lwx",
	"T9NzquWiscsyyn9rpr81jiCAv/y5u/lTa/ObPYqmhkhpxQ0Jva75+py+Fs9TSJ3QAYxqrCf5T3SXR+/K",
	"s++V09YYn/HbJF1G1yhojSzvriqIoDeHyQxqPyE/+7WInN39IfG07PQO6Wfy9qw8f3NxdFy+/NAwOgg2",
	"mxa3OSyGde4nrz9dHH2r8mv0E9lEf1H8CAAA//9EqDNTJwgAAA==",
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
