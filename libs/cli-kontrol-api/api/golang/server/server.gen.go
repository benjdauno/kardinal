// Package server provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen/v2 version 2.1.0 DO NOT EDIT.
package server

import (
	"bytes"
	"compress/gzip"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"path"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
	. "github.com/kurtosis-tech/kardinal/libs/cli-kontrol-api/api/golang/types"
	"github.com/labstack/echo/v4"
	"github.com/oapi-codegen/runtime"
	strictecho "github.com/oapi-codegen/runtime/strictmiddleware/echo"
)

// ServerInterface represents all server handlers.
type ServerInterface interface {

	// (GET /health)
	GetHealth(ctx echo.Context) error

	// (POST /tenant/{uuid}/deploy)
	PostTenantUuidDeploy(ctx echo.Context, uuid Uuid) error

	// (POST /tenant/{uuid}/flow/create)
	PostTenantUuidFlowCreate(ctx echo.Context, uuid Uuid) error

	// (DELETE /tenant/{uuid}/flow/{flow-id})
	DeleteTenantUuidFlowFlowId(ctx echo.Context, uuid Uuid, flowId FlowId) error

	// (GET /tenant/{uuid}/flows)
	GetTenantUuidFlows(ctx echo.Context, uuid Uuid) error
	// Cluster resource definition in a manifest YAML response
	// (GET /tenant/{uuid}/manifest)
	GetTenantUuidManifest(ctx echo.Context, uuid Uuid) error

	// (GET /tenant/{uuid}/templates)
	GetTenantUuidTemplates(ctx echo.Context, uuid Uuid) error

	// (POST /tenant/{uuid}/templates/create)
	PostTenantUuidTemplatesCreate(ctx echo.Context, uuid Uuid) error

	// (DELETE /tenant/{uuid}/templates/{template-name})
	DeleteTenantUuidTemplatesTemplateName(ctx echo.Context, uuid Uuid, templateName TemplateName) error

	// (GET /tenant/{uuid}/topology)
	GetTenantUuidTopology(ctx echo.Context, uuid Uuid) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// GetHealth converts echo context to params.
func (w *ServerInterfaceWrapper) GetHealth(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetHealth(ctx)
	return err
}

// PostTenantUuidDeploy converts echo context to params.
func (w *ServerInterfaceWrapper) PostTenantUuidDeploy(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "uuid" -------------
	var uuid Uuid

	err = runtime.BindStyledParameterWithOptions("simple", "uuid", ctx.Param("uuid"), &uuid, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter uuid: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.PostTenantUuidDeploy(ctx, uuid)
	return err
}

// PostTenantUuidFlowCreate converts echo context to params.
func (w *ServerInterfaceWrapper) PostTenantUuidFlowCreate(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "uuid" -------------
	var uuid Uuid

	err = runtime.BindStyledParameterWithOptions("simple", "uuid", ctx.Param("uuid"), &uuid, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter uuid: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.PostTenantUuidFlowCreate(ctx, uuid)
	return err
}

// DeleteTenantUuidFlowFlowId converts echo context to params.
func (w *ServerInterfaceWrapper) DeleteTenantUuidFlowFlowId(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "uuid" -------------
	var uuid Uuid

	err = runtime.BindStyledParameterWithOptions("simple", "uuid", ctx.Param("uuid"), &uuid, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter uuid: %s", err))
	}

	// ------------- Path parameter "flow-id" -------------
	var flowId FlowId

	err = runtime.BindStyledParameterWithOptions("simple", "flow-id", ctx.Param("flow-id"), &flowId, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter flow-id: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.DeleteTenantUuidFlowFlowId(ctx, uuid, flowId)
	return err
}

// GetTenantUuidFlows converts echo context to params.
func (w *ServerInterfaceWrapper) GetTenantUuidFlows(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "uuid" -------------
	var uuid Uuid

	err = runtime.BindStyledParameterWithOptions("simple", "uuid", ctx.Param("uuid"), &uuid, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter uuid: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetTenantUuidFlows(ctx, uuid)
	return err
}

// GetTenantUuidManifest converts echo context to params.
func (w *ServerInterfaceWrapper) GetTenantUuidManifest(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "uuid" -------------
	var uuid Uuid

	err = runtime.BindStyledParameterWithOptions("simple", "uuid", ctx.Param("uuid"), &uuid, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter uuid: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetTenantUuidManifest(ctx, uuid)
	return err
}

// GetTenantUuidTemplates converts echo context to params.
func (w *ServerInterfaceWrapper) GetTenantUuidTemplates(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "uuid" -------------
	var uuid Uuid

	err = runtime.BindStyledParameterWithOptions("simple", "uuid", ctx.Param("uuid"), &uuid, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter uuid: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetTenantUuidTemplates(ctx, uuid)
	return err
}

// PostTenantUuidTemplatesCreate converts echo context to params.
func (w *ServerInterfaceWrapper) PostTenantUuidTemplatesCreate(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "uuid" -------------
	var uuid Uuid

	err = runtime.BindStyledParameterWithOptions("simple", "uuid", ctx.Param("uuid"), &uuid, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter uuid: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.PostTenantUuidTemplatesCreate(ctx, uuid)
	return err
}

// DeleteTenantUuidTemplatesTemplateName converts echo context to params.
func (w *ServerInterfaceWrapper) DeleteTenantUuidTemplatesTemplateName(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "uuid" -------------
	var uuid Uuid

	err = runtime.BindStyledParameterWithOptions("simple", "uuid", ctx.Param("uuid"), &uuid, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter uuid: %s", err))
	}

	// ------------- Path parameter "template-name" -------------
	var templateName TemplateName

	err = runtime.BindStyledParameterWithOptions("simple", "template-name", ctx.Param("template-name"), &templateName, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter template-name: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.DeleteTenantUuidTemplatesTemplateName(ctx, uuid, templateName)
	return err
}

// GetTenantUuidTopology converts echo context to params.
func (w *ServerInterfaceWrapper) GetTenantUuidTopology(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "uuid" -------------
	var uuid Uuid

	err = runtime.BindStyledParameterWithOptions("simple", "uuid", ctx.Param("uuid"), &uuid, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter uuid: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetTenantUuidTopology(ctx, uuid)
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

	router.GET(baseURL+"/health", wrapper.GetHealth)
	router.POST(baseURL+"/tenant/:uuid/deploy", wrapper.PostTenantUuidDeploy)
	router.POST(baseURL+"/tenant/:uuid/flow/create", wrapper.PostTenantUuidFlowCreate)
	router.DELETE(baseURL+"/tenant/:uuid/flow/:flow-id", wrapper.DeleteTenantUuidFlowFlowId)
	router.GET(baseURL+"/tenant/:uuid/flows", wrapper.GetTenantUuidFlows)
	router.GET(baseURL+"/tenant/:uuid/manifest", wrapper.GetTenantUuidManifest)
	router.GET(baseURL+"/tenant/:uuid/templates", wrapper.GetTenantUuidTemplates)
	router.POST(baseURL+"/tenant/:uuid/templates/create", wrapper.PostTenantUuidTemplatesCreate)
	router.DELETE(baseURL+"/tenant/:uuid/templates/:template-name", wrapper.DeleteTenantUuidTemplatesTemplateName)
	router.GET(baseURL+"/tenant/:uuid/topology", wrapper.GetTenantUuidTopology)

}

type ErrorJSONResponse struct {
	// Error Error type
	Error string `json:"error"`

	// Msg Error message
	Msg *string `json:"msg,omitempty"`
}

type NotFoundJSONResponse struct {
	// Id Resource ID
	Id string `json:"id"`

	// ResourceType Resource type
	ResourceType string `json:"resource-type"`
}

type GetHealthRequestObject struct {
}

type GetHealthResponseObject interface {
	VisitGetHealthResponse(w http.ResponseWriter) error
}

type GetHealth200JSONResponse string

func (response GetHealth200JSONResponse) VisitGetHealthResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type PostTenantUuidDeployRequestObject struct {
	Uuid Uuid `json:"uuid"`
	Body *PostTenantUuidDeployJSONRequestBody
}

type PostTenantUuidDeployResponseObject interface {
	VisitPostTenantUuidDeployResponse(w http.ResponseWriter) error
}

type PostTenantUuidDeploy200JSONResponse Flow

func (response PostTenantUuidDeploy200JSONResponse) VisitPostTenantUuidDeployResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type PostTenantUuidDeploy404JSONResponse struct{ NotFoundJSONResponse }

func (response PostTenantUuidDeploy404JSONResponse) VisitPostTenantUuidDeployResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(404)

	return json.NewEncoder(w).Encode(response)
}

type PostTenantUuidDeploy500JSONResponse struct{ ErrorJSONResponse }

func (response PostTenantUuidDeploy500JSONResponse) VisitPostTenantUuidDeployResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(500)

	return json.NewEncoder(w).Encode(response)
}

type PostTenantUuidFlowCreateRequestObject struct {
	Uuid Uuid `json:"uuid"`
	Body *PostTenantUuidFlowCreateJSONRequestBody
}

type PostTenantUuidFlowCreateResponseObject interface {
	VisitPostTenantUuidFlowCreateResponse(w http.ResponseWriter) error
}

type PostTenantUuidFlowCreate200JSONResponse Flow

func (response PostTenantUuidFlowCreate200JSONResponse) VisitPostTenantUuidFlowCreateResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type PostTenantUuidFlowCreate404JSONResponse struct{ NotFoundJSONResponse }

func (response PostTenantUuidFlowCreate404JSONResponse) VisitPostTenantUuidFlowCreateResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(404)

	return json.NewEncoder(w).Encode(response)
}

type PostTenantUuidFlowCreate500JSONResponse struct{ ErrorJSONResponse }

func (response PostTenantUuidFlowCreate500JSONResponse) VisitPostTenantUuidFlowCreateResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(500)

	return json.NewEncoder(w).Encode(response)
}

type DeleteTenantUuidFlowFlowIdRequestObject struct {
	Uuid   Uuid   `json:"uuid"`
	FlowId FlowId `json:"flow-id"`
}

type DeleteTenantUuidFlowFlowIdResponseObject interface {
	VisitDeleteTenantUuidFlowFlowIdResponse(w http.ResponseWriter) error
}

type DeleteTenantUuidFlowFlowId2xxResponse struct {
	StatusCode int
}

func (response DeleteTenantUuidFlowFlowId2xxResponse) VisitDeleteTenantUuidFlowFlowIdResponse(w http.ResponseWriter) error {
	w.WriteHeader(response.StatusCode)
	return nil
}

type DeleteTenantUuidFlowFlowId404JSONResponse struct{ NotFoundJSONResponse }

func (response DeleteTenantUuidFlowFlowId404JSONResponse) VisitDeleteTenantUuidFlowFlowIdResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(404)

	return json.NewEncoder(w).Encode(response)
}

type DeleteTenantUuidFlowFlowId500JSONResponse struct{ ErrorJSONResponse }

func (response DeleteTenantUuidFlowFlowId500JSONResponse) VisitDeleteTenantUuidFlowFlowIdResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(500)

	return json.NewEncoder(w).Encode(response)
}

type GetTenantUuidFlowsRequestObject struct {
	Uuid Uuid `json:"uuid"`
}

type GetTenantUuidFlowsResponseObject interface {
	VisitGetTenantUuidFlowsResponse(w http.ResponseWriter) error
}

type GetTenantUuidFlows200JSONResponse []Flow

func (response GetTenantUuidFlows200JSONResponse) VisitGetTenantUuidFlowsResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type GetTenantUuidFlows404JSONResponse struct{ NotFoundJSONResponse }

func (response GetTenantUuidFlows404JSONResponse) VisitGetTenantUuidFlowsResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(404)

	return json.NewEncoder(w).Encode(response)
}

type GetTenantUuidFlows500JSONResponse struct{ ErrorJSONResponse }

func (response GetTenantUuidFlows500JSONResponse) VisitGetTenantUuidFlowsResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(500)

	return json.NewEncoder(w).Encode(response)
}

type GetTenantUuidManifestRequestObject struct {
	Uuid Uuid `json:"uuid"`
}

type GetTenantUuidManifestResponseObject interface {
	VisitGetTenantUuidManifestResponse(w http.ResponseWriter) error
}

type GetTenantUuidManifest200ApplicationxYamlResponse struct {
	Body          io.Reader
	ContentLength int64
}

func (response GetTenantUuidManifest200ApplicationxYamlResponse) VisitGetTenantUuidManifestResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/x-yaml")
	if response.ContentLength != 0 {
		w.Header().Set("Content-Length", fmt.Sprint(response.ContentLength))
	}
	w.WriteHeader(200)

	if closer, ok := response.Body.(io.ReadCloser); ok {
		defer closer.Close()
	}
	_, err := io.Copy(w, response.Body)
	return err
}

type GetTenantUuidManifest404JSONResponse struct{ NotFoundJSONResponse }

func (response GetTenantUuidManifest404JSONResponse) VisitGetTenantUuidManifestResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(404)

	return json.NewEncoder(w).Encode(response)
}

type GetTenantUuidManifest500JSONResponse struct{ ErrorJSONResponse }

func (response GetTenantUuidManifest500JSONResponse) VisitGetTenantUuidManifestResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(500)

	return json.NewEncoder(w).Encode(response)
}

type GetTenantUuidTemplatesRequestObject struct {
	Uuid Uuid `json:"uuid"`
}

type GetTenantUuidTemplatesResponseObject interface {
	VisitGetTenantUuidTemplatesResponse(w http.ResponseWriter) error
}

type GetTenantUuidTemplates200JSONResponse []Template

func (response GetTenantUuidTemplates200JSONResponse) VisitGetTenantUuidTemplatesResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type GetTenantUuidTemplates404JSONResponse struct{ NotFoundJSONResponse }

func (response GetTenantUuidTemplates404JSONResponse) VisitGetTenantUuidTemplatesResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(404)

	return json.NewEncoder(w).Encode(response)
}

type GetTenantUuidTemplates500JSONResponse struct{ ErrorJSONResponse }

func (response GetTenantUuidTemplates500JSONResponse) VisitGetTenantUuidTemplatesResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(500)

	return json.NewEncoder(w).Encode(response)
}

type PostTenantUuidTemplatesCreateRequestObject struct {
	Uuid Uuid `json:"uuid"`
	Body *PostTenantUuidTemplatesCreateJSONRequestBody
}

type PostTenantUuidTemplatesCreateResponseObject interface {
	VisitPostTenantUuidTemplatesCreateResponse(w http.ResponseWriter) error
}

type PostTenantUuidTemplatesCreate200JSONResponse Template

func (response PostTenantUuidTemplatesCreate200JSONResponse) VisitPostTenantUuidTemplatesCreateResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type PostTenantUuidTemplatesCreate404JSONResponse struct{ NotFoundJSONResponse }

func (response PostTenantUuidTemplatesCreate404JSONResponse) VisitPostTenantUuidTemplatesCreateResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(404)

	return json.NewEncoder(w).Encode(response)
}

type PostTenantUuidTemplatesCreate500JSONResponse struct{ ErrorJSONResponse }

func (response PostTenantUuidTemplatesCreate500JSONResponse) VisitPostTenantUuidTemplatesCreateResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(500)

	return json.NewEncoder(w).Encode(response)
}

type DeleteTenantUuidTemplatesTemplateNameRequestObject struct {
	Uuid         Uuid         `json:"uuid"`
	TemplateName TemplateName `json:"template-name"`
}

type DeleteTenantUuidTemplatesTemplateNameResponseObject interface {
	VisitDeleteTenantUuidTemplatesTemplateNameResponse(w http.ResponseWriter) error
}

type DeleteTenantUuidTemplatesTemplateName2xxResponse struct {
	StatusCode int
}

func (response DeleteTenantUuidTemplatesTemplateName2xxResponse) VisitDeleteTenantUuidTemplatesTemplateNameResponse(w http.ResponseWriter) error {
	w.WriteHeader(response.StatusCode)
	return nil
}

type DeleteTenantUuidTemplatesTemplateName404JSONResponse struct{ NotFoundJSONResponse }

func (response DeleteTenantUuidTemplatesTemplateName404JSONResponse) VisitDeleteTenantUuidTemplatesTemplateNameResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(404)

	return json.NewEncoder(w).Encode(response)
}

type DeleteTenantUuidTemplatesTemplateName500JSONResponse struct{ ErrorJSONResponse }

func (response DeleteTenantUuidTemplatesTemplateName500JSONResponse) VisitDeleteTenantUuidTemplatesTemplateNameResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(500)

	return json.NewEncoder(w).Encode(response)
}

type GetTenantUuidTopologyRequestObject struct {
	Uuid Uuid `json:"uuid"`
}

type GetTenantUuidTopologyResponseObject interface {
	VisitGetTenantUuidTopologyResponse(w http.ResponseWriter) error
}

type GetTenantUuidTopology200JSONResponse ClusterTopology

func (response GetTenantUuidTopology200JSONResponse) VisitGetTenantUuidTopologyResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type GetTenantUuidTopology404JSONResponse struct{ NotFoundJSONResponse }

func (response GetTenantUuidTopology404JSONResponse) VisitGetTenantUuidTopologyResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(404)

	return json.NewEncoder(w).Encode(response)
}

type GetTenantUuidTopology500JSONResponse struct{ ErrorJSONResponse }

func (response GetTenantUuidTopology500JSONResponse) VisitGetTenantUuidTopologyResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(500)

	return json.NewEncoder(w).Encode(response)
}

// StrictServerInterface represents all server handlers.
type StrictServerInterface interface {

	// (GET /health)
	GetHealth(ctx context.Context, request GetHealthRequestObject) (GetHealthResponseObject, error)

	// (POST /tenant/{uuid}/deploy)
	PostTenantUuidDeploy(ctx context.Context, request PostTenantUuidDeployRequestObject) (PostTenantUuidDeployResponseObject, error)

	// (POST /tenant/{uuid}/flow/create)
	PostTenantUuidFlowCreate(ctx context.Context, request PostTenantUuidFlowCreateRequestObject) (PostTenantUuidFlowCreateResponseObject, error)

	// (DELETE /tenant/{uuid}/flow/{flow-id})
	DeleteTenantUuidFlowFlowId(ctx context.Context, request DeleteTenantUuidFlowFlowIdRequestObject) (DeleteTenantUuidFlowFlowIdResponseObject, error)

	// (GET /tenant/{uuid}/flows)
	GetTenantUuidFlows(ctx context.Context, request GetTenantUuidFlowsRequestObject) (GetTenantUuidFlowsResponseObject, error)
	// Cluster resource definition in a manifest YAML response
	// (GET /tenant/{uuid}/manifest)
	GetTenantUuidManifest(ctx context.Context, request GetTenantUuidManifestRequestObject) (GetTenantUuidManifestResponseObject, error)

	// (GET /tenant/{uuid}/templates)
	GetTenantUuidTemplates(ctx context.Context, request GetTenantUuidTemplatesRequestObject) (GetTenantUuidTemplatesResponseObject, error)

	// (POST /tenant/{uuid}/templates/create)
	PostTenantUuidTemplatesCreate(ctx context.Context, request PostTenantUuidTemplatesCreateRequestObject) (PostTenantUuidTemplatesCreateResponseObject, error)

	// (DELETE /tenant/{uuid}/templates/{template-name})
	DeleteTenantUuidTemplatesTemplateName(ctx context.Context, request DeleteTenantUuidTemplatesTemplateNameRequestObject) (DeleteTenantUuidTemplatesTemplateNameResponseObject, error)

	// (GET /tenant/{uuid}/topology)
	GetTenantUuidTopology(ctx context.Context, request GetTenantUuidTopologyRequestObject) (GetTenantUuidTopologyResponseObject, error)
}

type StrictHandlerFunc = strictecho.StrictEchoHandlerFunc
type StrictMiddlewareFunc = strictecho.StrictEchoMiddlewareFunc

func NewStrictHandler(ssi StrictServerInterface, middlewares []StrictMiddlewareFunc) ServerInterface {
	return &strictHandler{ssi: ssi, middlewares: middlewares}
}

type strictHandler struct {
	ssi         StrictServerInterface
	middlewares []StrictMiddlewareFunc
}

// GetHealth operation middleware
func (sh *strictHandler) GetHealth(ctx echo.Context) error {
	var request GetHealthRequestObject

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.GetHealth(ctx.Request().Context(), request.(GetHealthRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "GetHealth")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(GetHealthResponseObject); ok {
		return validResponse.VisitGetHealthResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}

// PostTenantUuidDeploy operation middleware
func (sh *strictHandler) PostTenantUuidDeploy(ctx echo.Context, uuid Uuid) error {
	var request PostTenantUuidDeployRequestObject

	request.Uuid = uuid

	var body PostTenantUuidDeployJSONRequestBody
	if err := ctx.Bind(&body); err != nil {
		return err
	}
	request.Body = &body

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.PostTenantUuidDeploy(ctx.Request().Context(), request.(PostTenantUuidDeployRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "PostTenantUuidDeploy")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(PostTenantUuidDeployResponseObject); ok {
		return validResponse.VisitPostTenantUuidDeployResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}

// PostTenantUuidFlowCreate operation middleware
func (sh *strictHandler) PostTenantUuidFlowCreate(ctx echo.Context, uuid Uuid) error {
	var request PostTenantUuidFlowCreateRequestObject

	request.Uuid = uuid

	var body PostTenantUuidFlowCreateJSONRequestBody
	if err := ctx.Bind(&body); err != nil {
		return err
	}
	request.Body = &body

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.PostTenantUuidFlowCreate(ctx.Request().Context(), request.(PostTenantUuidFlowCreateRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "PostTenantUuidFlowCreate")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(PostTenantUuidFlowCreateResponseObject); ok {
		return validResponse.VisitPostTenantUuidFlowCreateResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}

// DeleteTenantUuidFlowFlowId operation middleware
func (sh *strictHandler) DeleteTenantUuidFlowFlowId(ctx echo.Context, uuid Uuid, flowId FlowId) error {
	var request DeleteTenantUuidFlowFlowIdRequestObject

	request.Uuid = uuid
	request.FlowId = flowId

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.DeleteTenantUuidFlowFlowId(ctx.Request().Context(), request.(DeleteTenantUuidFlowFlowIdRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "DeleteTenantUuidFlowFlowId")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(DeleteTenantUuidFlowFlowIdResponseObject); ok {
		return validResponse.VisitDeleteTenantUuidFlowFlowIdResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}

// GetTenantUuidFlows operation middleware
func (sh *strictHandler) GetTenantUuidFlows(ctx echo.Context, uuid Uuid) error {
	var request GetTenantUuidFlowsRequestObject

	request.Uuid = uuid

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.GetTenantUuidFlows(ctx.Request().Context(), request.(GetTenantUuidFlowsRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "GetTenantUuidFlows")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(GetTenantUuidFlowsResponseObject); ok {
		return validResponse.VisitGetTenantUuidFlowsResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}

// GetTenantUuidManifest operation middleware
func (sh *strictHandler) GetTenantUuidManifest(ctx echo.Context, uuid Uuid) error {
	var request GetTenantUuidManifestRequestObject

	request.Uuid = uuid

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.GetTenantUuidManifest(ctx.Request().Context(), request.(GetTenantUuidManifestRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "GetTenantUuidManifest")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(GetTenantUuidManifestResponseObject); ok {
		return validResponse.VisitGetTenantUuidManifestResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}

// GetTenantUuidTemplates operation middleware
func (sh *strictHandler) GetTenantUuidTemplates(ctx echo.Context, uuid Uuid) error {
	var request GetTenantUuidTemplatesRequestObject

	request.Uuid = uuid

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.GetTenantUuidTemplates(ctx.Request().Context(), request.(GetTenantUuidTemplatesRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "GetTenantUuidTemplates")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(GetTenantUuidTemplatesResponseObject); ok {
		return validResponse.VisitGetTenantUuidTemplatesResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}

// PostTenantUuidTemplatesCreate operation middleware
func (sh *strictHandler) PostTenantUuidTemplatesCreate(ctx echo.Context, uuid Uuid) error {
	var request PostTenantUuidTemplatesCreateRequestObject

	request.Uuid = uuid

	var body PostTenantUuidTemplatesCreateJSONRequestBody
	if err := ctx.Bind(&body); err != nil {
		return err
	}
	request.Body = &body

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.PostTenantUuidTemplatesCreate(ctx.Request().Context(), request.(PostTenantUuidTemplatesCreateRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "PostTenantUuidTemplatesCreate")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(PostTenantUuidTemplatesCreateResponseObject); ok {
		return validResponse.VisitPostTenantUuidTemplatesCreateResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}

// DeleteTenantUuidTemplatesTemplateName operation middleware
func (sh *strictHandler) DeleteTenantUuidTemplatesTemplateName(ctx echo.Context, uuid Uuid, templateName TemplateName) error {
	var request DeleteTenantUuidTemplatesTemplateNameRequestObject

	request.Uuid = uuid
	request.TemplateName = templateName

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.DeleteTenantUuidTemplatesTemplateName(ctx.Request().Context(), request.(DeleteTenantUuidTemplatesTemplateNameRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "DeleteTenantUuidTemplatesTemplateName")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(DeleteTenantUuidTemplatesTemplateNameResponseObject); ok {
		return validResponse.VisitDeleteTenantUuidTemplatesTemplateNameResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}

// GetTenantUuidTopology operation middleware
func (sh *strictHandler) GetTenantUuidTopology(ctx echo.Context, uuid Uuid) error {
	var request GetTenantUuidTopologyRequestObject

	request.Uuid = uuid

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.GetTenantUuidTopology(ctx.Request().Context(), request.(GetTenantUuidTopologyRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "GetTenantUuidTopology")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(GetTenantUuidTopologyResponseObject); ok {
		return validResponse.VisitGetTenantUuidTopologyResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"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",
	"C+vlNYq+2/0dAAD//0dBE5meIAAA",
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
