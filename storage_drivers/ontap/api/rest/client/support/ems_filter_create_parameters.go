// Code generated by go-swagger; DO NOT EDIT.

package support

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/netapp/trident/storage_drivers/ontap/api/rest/models"
)

// NewEmsFilterCreateParams creates a new EmsFilterCreateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewEmsFilterCreateParams() *EmsFilterCreateParams {
	return &EmsFilterCreateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewEmsFilterCreateParamsWithTimeout creates a new EmsFilterCreateParams object
// with the ability to set a timeout on a request.
func NewEmsFilterCreateParamsWithTimeout(timeout time.Duration) *EmsFilterCreateParams {
	return &EmsFilterCreateParams{
		timeout: timeout,
	}
}

// NewEmsFilterCreateParamsWithContext creates a new EmsFilterCreateParams object
// with the ability to set a context for a request.
func NewEmsFilterCreateParamsWithContext(ctx context.Context) *EmsFilterCreateParams {
	return &EmsFilterCreateParams{
		Context: ctx,
	}
}

// NewEmsFilterCreateParamsWithHTTPClient creates a new EmsFilterCreateParams object
// with the ability to set a custom HTTPClient for a request.
func NewEmsFilterCreateParamsWithHTTPClient(client *http.Client) *EmsFilterCreateParams {
	return &EmsFilterCreateParams{
		HTTPClient: client,
	}
}

/*
EmsFilterCreateParams contains all the parameters to send to the API endpoint

	for the ems filter create operation.

	Typically these are written to a http.Request.
*/
type EmsFilterCreateParams struct {

	/* Info.

	   Information specification
	*/
	Info *models.EmsFilter

	/* ReturnRecords.

	   The default is false.  If set to true, the records are returned.
	*/
	ReturnRecordsQueryParameter *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the ems filter create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *EmsFilterCreateParams) WithDefaults() *EmsFilterCreateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the ems filter create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *EmsFilterCreateParams) SetDefaults() {
	var (
		returnRecordsQueryParameterDefault = bool(false)
	)

	val := EmsFilterCreateParams{
		ReturnRecordsQueryParameter: &returnRecordsQueryParameterDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the ems filter create params
func (o *EmsFilterCreateParams) WithTimeout(timeout time.Duration) *EmsFilterCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the ems filter create params
func (o *EmsFilterCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the ems filter create params
func (o *EmsFilterCreateParams) WithContext(ctx context.Context) *EmsFilterCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the ems filter create params
func (o *EmsFilterCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the ems filter create params
func (o *EmsFilterCreateParams) WithHTTPClient(client *http.Client) *EmsFilterCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the ems filter create params
func (o *EmsFilterCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInfo adds the info to the ems filter create params
func (o *EmsFilterCreateParams) WithInfo(info *models.EmsFilter) *EmsFilterCreateParams {
	o.SetInfo(info)
	return o
}

// SetInfo adds the info to the ems filter create params
func (o *EmsFilterCreateParams) SetInfo(info *models.EmsFilter) {
	o.Info = info
}

// WithReturnRecordsQueryParameter adds the returnRecords to the ems filter create params
func (o *EmsFilterCreateParams) WithReturnRecordsQueryParameter(returnRecords *bool) *EmsFilterCreateParams {
	o.SetReturnRecordsQueryParameter(returnRecords)
	return o
}

// SetReturnRecordsQueryParameter adds the returnRecords to the ems filter create params
func (o *EmsFilterCreateParams) SetReturnRecordsQueryParameter(returnRecords *bool) {
	o.ReturnRecordsQueryParameter = returnRecords
}

// WriteToRequest writes these params to a swagger request
func (o *EmsFilterCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Info != nil {
		if err := r.SetBodyParam(o.Info); err != nil {
			return err
		}
	}

	if o.ReturnRecordsQueryParameter != nil {

		// query param return_records
		var qrReturnRecords bool

		if o.ReturnRecordsQueryParameter != nil {
			qrReturnRecords = *o.ReturnRecordsQueryParameter
		}
		qReturnRecords := swag.FormatBool(qrReturnRecords)
		if qReturnRecords != "" {

			if err := r.SetQueryParam("return_records", qReturnRecords); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
