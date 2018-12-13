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
)

// NewGetWidgetListParams creates a new GetWidgetListParams object
// with the default values initialized.
func NewGetWidgetListParams() *GetWidgetListParams {
	var (
		offsetDefault = int32(0)
		sizeDefault   = int32(50)
	)
	return &GetWidgetListParams{
		Offset: &offsetDefault,
		Size:   &sizeDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetWidgetListParamsWithTimeout creates a new GetWidgetListParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetWidgetListParamsWithTimeout(timeout time.Duration) *GetWidgetListParams {
	var (
		offsetDefault = int32(0)
		sizeDefault   = int32(50)
	)
	return &GetWidgetListParams{
		Offset: &offsetDefault,
		Size:   &sizeDefault,

		timeout: timeout,
	}
}

// NewGetWidgetListParamsWithContext creates a new GetWidgetListParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetWidgetListParamsWithContext(ctx context.Context) *GetWidgetListParams {
	var (
		offsetDefault = int32(0)
		sizeDefault   = int32(50)
	)
	return &GetWidgetListParams{
		Offset: &offsetDefault,
		Size:   &sizeDefault,

		Context: ctx,
	}
}

// NewGetWidgetListParamsWithHTTPClient creates a new GetWidgetListParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetWidgetListParamsWithHTTPClient(client *http.Client) *GetWidgetListParams {
	var (
		offsetDefault = int32(0)
		sizeDefault   = int32(50)
	)
	return &GetWidgetListParams{
		Offset:     &offsetDefault,
		Size:       &sizeDefault,
		HTTPClient: client,
	}
}

/*GetWidgetListParams contains all the parameters to send to the API endpoint
for the get widget list operation typically these are written to a http.Request
*/
type GetWidgetListParams struct {

	/*Fields*/
	Fields *string
	/*Filter*/
	Filter *string
	/*Offset*/
	Offset *int32
	/*Size*/
	Size *int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get widget list params
func (o *GetWidgetListParams) WithTimeout(timeout time.Duration) *GetWidgetListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get widget list params
func (o *GetWidgetListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get widget list params
func (o *GetWidgetListParams) WithContext(ctx context.Context) *GetWidgetListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get widget list params
func (o *GetWidgetListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get widget list params
func (o *GetWidgetListParams) WithHTTPClient(client *http.Client) *GetWidgetListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get widget list params
func (o *GetWidgetListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFields adds the fields to the get widget list params
func (o *GetWidgetListParams) WithFields(fields *string) *GetWidgetListParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the get widget list params
func (o *GetWidgetListParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithFilter adds the filter to the get widget list params
func (o *GetWidgetListParams) WithFilter(filter *string) *GetWidgetListParams {
	o.SetFilter(filter)
	return o
}

// SetFilter adds the filter to the get widget list params
func (o *GetWidgetListParams) SetFilter(filter *string) {
	o.Filter = filter
}

// WithOffset adds the offset to the get widget list params
func (o *GetWidgetListParams) WithOffset(offset *int32) *GetWidgetListParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the get widget list params
func (o *GetWidgetListParams) SetOffset(offset *int32) {
	o.Offset = offset
}

// WithSize adds the size to the get widget list params
func (o *GetWidgetListParams) WithSize(size *int32) *GetWidgetListParams {
	o.SetSize(size)
	return o
}

// SetSize adds the size to the get widget list params
func (o *GetWidgetListParams) SetSize(size *int32) {
	o.Size = size
}

// WriteToRequest writes these params to a swagger request
func (o *GetWidgetListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Fields != nil {

		// query param fields
		var qrFields string
		if o.Fields != nil {
			qrFields = *o.Fields
		}
		qFields := qrFields
		if qFields != "" {
			if err := r.SetQueryParam("fields", qFields); err != nil {
				return err
			}
		}

	}

	if o.Filter != nil {

		// query param filter
		var qrFilter string
		if o.Filter != nil {
			qrFilter = *o.Filter
		}
		qFilter := qrFilter
		if qFilter != "" {
			if err := r.SetQueryParam("filter", qFilter); err != nil {
				return err
			}
		}

	}

	if o.Offset != nil {

		// query param offset
		var qrOffset int32
		if o.Offset != nil {
			qrOffset = *o.Offset
		}
		qOffset := swag.FormatInt32(qrOffset)
		if qOffset != "" {
			if err := r.SetQueryParam("offset", qOffset); err != nil {
				return err
			}
		}

	}

	if o.Size != nil {

		// query param size
		var qrSize int32
		if o.Size != nil {
			qrSize = *o.Size
		}
		qSize := swag.FormatInt32(qrSize)
		if qSize != "" {
			if err := r.SetQueryParam("size", qSize); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
