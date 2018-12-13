// Code generated by go-swagger; DO NOT EDIT.

package lm

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/logicmonitor/lm-sdk-go/models"
)

// NewPatchWebsiteByIDParams creates a new PatchWebsiteByIDParams object
// with the default values initialized.
func NewPatchWebsiteByIDParams() *PatchWebsiteByIDParams {
	var ()
	return &PatchWebsiteByIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPatchWebsiteByIDParamsWithTimeout creates a new PatchWebsiteByIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPatchWebsiteByIDParamsWithTimeout(timeout time.Duration) *PatchWebsiteByIDParams {
	var ()
	return &PatchWebsiteByIDParams{

		timeout: timeout,
	}
}

// NewPatchWebsiteByIDParamsWithContext creates a new PatchWebsiteByIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewPatchWebsiteByIDParamsWithContext(ctx context.Context) *PatchWebsiteByIDParams {
	var ()
	return &PatchWebsiteByIDParams{

		Context: ctx,
	}
}

// NewPatchWebsiteByIDParamsWithHTTPClient creates a new PatchWebsiteByIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPatchWebsiteByIDParamsWithHTTPClient(client *http.Client) *PatchWebsiteByIDParams {
	var ()
	return &PatchWebsiteByIDParams{
		HTTPClient: client,
	}
}

/*PatchWebsiteByIDParams contains all the parameters to send to the API endpoint
for the patch website by Id operation typically these are written to a http.Request
*/
type PatchWebsiteByIDParams struct {

	/*Body*/
	Body models.Website
	/*ID*/
	ID int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the patch website by Id params
func (o *PatchWebsiteByIDParams) WithTimeout(timeout time.Duration) *PatchWebsiteByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch website by Id params
func (o *PatchWebsiteByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch website by Id params
func (o *PatchWebsiteByIDParams) WithContext(ctx context.Context) *PatchWebsiteByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch website by Id params
func (o *PatchWebsiteByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch website by Id params
func (o *PatchWebsiteByIDParams) WithHTTPClient(client *http.Client) *PatchWebsiteByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch website by Id params
func (o *PatchWebsiteByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the patch website by Id params
func (o *PatchWebsiteByIDParams) WithBody(body models.Website) *PatchWebsiteByIDParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the patch website by Id params
func (o *PatchWebsiteByIDParams) SetBody(body models.Website) {
	o.Body = body
}

// WithID adds the id to the patch website by Id params
func (o *PatchWebsiteByIDParams) WithID(id int32) *PatchWebsiteByIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the patch website by Id params
func (o *PatchWebsiteByIDParams) SetID(id int32) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *PatchWebsiteByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt32(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
