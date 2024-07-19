// Package ticketsservice provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/oapi-codegen/oapi-codegen/v2 version v2.3.0 DO NOT EDIT.
package ticketsservice

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"net/http"
	"net/url"
	"path"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/labstack/echo/v4"
	"github.com/oapi-codegen/runtime"
)

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// List Tickets
	// (GET /tickets)
	ListTicket(ctx echo.Context, params ListTicketParams) error
	// Create a new Ticket
	// (POST /tickets)
	CreateTicket(ctx echo.Context) error
	// Deletes a Ticket by ID
	// (DELETE /tickets/{id})
	DeleteTicket(ctx echo.Context, id int) error
	// Find a Ticket by ID
	// (GET /tickets/{id})
	ReadTicket(ctx echo.Context, id int) error
	// Updates a Ticket
	// (PATCH /tickets/{id})
	UpdateTicket(ctx echo.Context, id int) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// ListTicket converts echo context to params.
func (w *ServerInterfaceWrapper) ListTicket(ctx echo.Context) error {
	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params ListTicketParams
	// ------------- Optional query parameter "status" -------------

	err = runtime.BindQueryParameter("form", true, false, "status", ctx.QueryParams(), &params.Status)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter status: %s", err))
	}

	// ------------- Optional query parameter "page" -------------

	err = runtime.BindQueryParameter("form", true, false, "page", ctx.QueryParams(), &params.Page)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter page: %s", err))
	}

	// ------------- Optional query parameter "itemsPerPage" -------------

	err = runtime.BindQueryParameter("form", true, false, "itemsPerPage", ctx.QueryParams(), &params.ItemsPerPage)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter itemsPerPage: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.ListTicket(ctx, params)
	return err
}

// CreateTicket converts echo context to params.
func (w *ServerInterfaceWrapper) CreateTicket(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.CreateTicket(ctx)
	return err
}

// DeleteTicket converts echo context to params.
func (w *ServerInterfaceWrapper) DeleteTicket(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "id" -------------
	var id int

	err = runtime.BindStyledParameterWithOptions("simple", "id", ctx.Param("id"), &id, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter id: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.DeleteTicket(ctx, id)
	return err
}

// ReadTicket converts echo context to params.
func (w *ServerInterfaceWrapper) ReadTicket(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "id" -------------
	var id int

	err = runtime.BindStyledParameterWithOptions("simple", "id", ctx.Param("id"), &id, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter id: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.ReadTicket(ctx, id)
	return err
}

// UpdateTicket converts echo context to params.
func (w *ServerInterfaceWrapper) UpdateTicket(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "id" -------------
	var id int

	err = runtime.BindStyledParameterWithOptions("simple", "id", ctx.Param("id"), &id, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter id: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.UpdateTicket(ctx, id)
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

	router.GET(baseURL+"/tickets", wrapper.ListTicket)
	router.POST(baseURL+"/tickets", wrapper.CreateTicket)
	router.DELETE(baseURL+"/tickets/:id", wrapper.DeleteTicket)
	router.GET(baseURL+"/tickets/:id", wrapper.ReadTicket)
	router.PATCH(baseURL+"/tickets/:id", wrapper.UpdateTicket)

}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"H4sIAAAAAAAC/+xXX2/bNhD/KgS3R8F2W+chfuuWDTCwh2DrnoogYMSzdJ1EquSxrhHouw9HWrJsKam7",
	"bl0T9CEIJZ7u7+93d76Xua0ba8CQl6t76cA31njoH5aLBR9zawgM8VE1TYW5IrRm/s5bw+98XkKt+NQ4",
	"24AjTBpyq4H/064BuZJoCApwss0kOGcdy7SZ9KQo+IGcJ4emkG2bSQfvAzrQcvU2aevFb7JO3N69g5xk",
	"y/IafO6wYe+iwQ+qQi3QNIEyoRUpsX8no3IOcPmEA3TgbXA5CGNJbGwwg7gun3BcuTWbCnNCU4guRt+F",
	"dvGkMRkMfGwgJ9AiGowqk7PR3hvM/wIa+628x8IATPiUSdTTIdkGDOjbu93kV41D65DipYaNChXJlazs",
	"VmYSTKg5wPRUg8ZQy0yWWJSDMA+6Dgk7aDKWbj0pR6AHGo/forltnC0ceC8zqa2BSfVUOlD69ijOwS1S",
	"BZ8uFrLBJDtMTe/8ICNDi9NVRbOxKdxhdd+U6AV6oYxQgawowIBTXOzX12sxkBW10iBsIGE3LP2LIZFQ",
	"IDRs0GBU2Icm+f6PdP/6ei0z+QGcTzYXsxezRVdt1aBcyVezxewVx6OojEWZU4RVPBcJXseO/4aeRMKe",
	"n6X0uEiotd7f7oHJSp2qgYC58vZUT7IjUkbFBisCF8ssV/J9AMepNapOVdpn/UDVz8FIjQZrFl9k47qf",
	"urUtFYlGFSDICgdGP+gVCx351Nt5kY0YNjaEBLXIbTB0sCQa/kt6p0zyN/4a3PXItPqYTL+8uMged+Qm",
	"Ox7eLz+zSUYf+PCjg41cyR/mh81gvm9P8z0E2t6+ck7tHhhMoeoAJSr08av9NjFlond+3q0dUX55rvwy",
	"yV+eK3/J8hfn+sOCsU+HulZud8IXpqkqmA1d977h7mr9BM9+dqAIvFDCwLbLjzKaMeLRkxcYkePJOlXA",
	"mIpJQU9G7m/g6Serd18wFB8dLt+HyHCI/Fvz42RuJCCQFXmsrxwaJReg/UJ+n0PrB91KPul/RuGvRsnE",
	"jCNmTTGzzfppOL9H3Sa4VUAwZutVfO8FldBxdYtUxuc980CL9dWYpunL82bm+oqXgIONbkzw9B5MCT0C",
	"xXBWnDERlhP7yiCoYUBiq7xIWdHPrHN3NVVdRe92Yn013cMnV6Vf0ehPQyI2dQcUnOGePkbI76D0t4WP",
	"r9hRpgHX/4J9RnBjtJyFtUZRXo7R9mej1RFcj5aFvFSm4Ab1yMaQNPzvUPu+p/zne8q5a8YhqQycEPHx",
	"bW0cyafn1gtOuTy5nrTt3wEAAP//XEpxfJYVAAA=",
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
