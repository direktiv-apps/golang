// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"

	"app/models"
)

// NewPostParams creates a new PostParams object
// with the default values initialized.
func NewPostParams() PostParams {

	var (
		// initialize parameters with default values

		direktivActionIDDefault = string("development")
		direktivTempDirDefault  = string("/tmp")
	)

	return PostParams{
		DirektivActionID: &direktivActionIDDefault,

		DirektivTempDir: &direktivTempDirDefault,
	}
}

// PostParams contains all the bound params for the post operation
// typically these are obtained from a http.Request
//
// swagger:parameters Post
type PostParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*direktiv action id is an UUID.
	For development it can be set to 'development'

	  In: header
	  Default: "development"
	*/
	DirektivActionID *string
	/*direktiv temp dir is the working directory for that request
	For development it can be set to e.g. '/tmp'

	  In: header
	  Default: "/tmp"
	*/
	DirektivTempDir *string
	/*
	  In: body
	*/
	Body *models.PostParamsBody
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewPostParams() beforehand.
func (o *PostParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	if err := o.bindDirektivActionID(r.Header[http.CanonicalHeaderKey("Direktiv-ActionID")], true, route.Formats); err != nil {
		res = append(res, err)
	}

	if err := o.bindDirektivTempDir(r.Header[http.CanonicalHeaderKey("Direktiv-TempDir")], true, route.Formats); err != nil {
		res = append(res, err)
	}

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body models.PostParamsBody
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			res = append(res, errors.NewParseError("body", "body", "", err))
		} else {
			// validate body object
			if err := body.Validate(route.Formats); err != nil {
				res = append(res, err)
			}

			// direktiv: add request to context so we can use in validate
			baseCtx := context.WithValue(r.Context(), "req", r)

			ctx := validate.WithOperationRequest(baseCtx)
			if err := body.ContextValidate(ctx, route.Formats); err != nil {
				res = append(res, err)
			}

			if len(res) == 0 {
				o.Body = &body
			}
		}
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindDirektivActionID binds and validates parameter DirektivActionID from header.
func (o *PostParams) bindDirektivActionID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false

	if raw == "" { // empty values pass all other validations
		// Default values have been previously initialized by NewPostParams()
		return nil
	}
	o.DirektivActionID = &raw

	return nil
}

// bindDirektivTempDir binds and validates parameter DirektivTempDir from header.
func (o *PostParams) bindDirektivTempDir(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false

	if raw == "" { // empty values pass all other validations
		// Default values have been previously initialized by NewPostParams()
		return nil
	}
	o.DirektivTempDir = &raw

	return nil
}
