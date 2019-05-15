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

// NewGetDeviceDatasourceInstanceGraphDataParams creates a new GetDeviceDatasourceInstanceGraphDataParams object
// with the default values initialized.
func NewGetDeviceDatasourceInstanceGraphDataParams() *GetDeviceDatasourceInstanceGraphDataParams {
	var ()
	return &GetDeviceDatasourceInstanceGraphDataParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetDeviceDatasourceInstanceGraphDataParamsWithTimeout creates a new GetDeviceDatasourceInstanceGraphDataParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetDeviceDatasourceInstanceGraphDataParamsWithTimeout(timeout time.Duration) *GetDeviceDatasourceInstanceGraphDataParams {
	var ()
	return &GetDeviceDatasourceInstanceGraphDataParams{

		timeout: timeout,
	}
}

// NewGetDeviceDatasourceInstanceGraphDataParamsWithContext creates a new GetDeviceDatasourceInstanceGraphDataParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetDeviceDatasourceInstanceGraphDataParamsWithContext(ctx context.Context) *GetDeviceDatasourceInstanceGraphDataParams {
	var ()
	return &GetDeviceDatasourceInstanceGraphDataParams{

		Context: ctx,
	}
}

// NewGetDeviceDatasourceInstanceGraphDataParamsWithHTTPClient creates a new GetDeviceDatasourceInstanceGraphDataParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetDeviceDatasourceInstanceGraphDataParamsWithHTTPClient(client *http.Client) *GetDeviceDatasourceInstanceGraphDataParams {
	var ()
	return &GetDeviceDatasourceInstanceGraphDataParams{
		HTTPClient: client,
	}
}

/*GetDeviceDatasourceInstanceGraphDataParams contains all the parameters to send to the API endpoint
for the get device datasource instance graph data operation typically these are written to a http.Request
*/
type GetDeviceDatasourceInstanceGraphDataParams struct {

	/*DeviceID*/
	DeviceID int32
	/*End*/
	End *int64
	/*Format*/
	Format *string
	/*GraphID*/
	GraphID int32
	/*HdsID
	  The device-datasource ID

	*/
	HdsID int32
	/*ID*/
	ID int32
	/*Start*/
	Start *int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get device datasource instance graph data params
func (o *GetDeviceDatasourceInstanceGraphDataParams) WithTimeout(timeout time.Duration) *GetDeviceDatasourceInstanceGraphDataParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get device datasource instance graph data params
func (o *GetDeviceDatasourceInstanceGraphDataParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get device datasource instance graph data params
func (o *GetDeviceDatasourceInstanceGraphDataParams) WithContext(ctx context.Context) *GetDeviceDatasourceInstanceGraphDataParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get device datasource instance graph data params
func (o *GetDeviceDatasourceInstanceGraphDataParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get device datasource instance graph data params
func (o *GetDeviceDatasourceInstanceGraphDataParams) WithHTTPClient(client *http.Client) *GetDeviceDatasourceInstanceGraphDataParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get device datasource instance graph data params
func (o *GetDeviceDatasourceInstanceGraphDataParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDeviceID adds the deviceID to the get device datasource instance graph data params
func (o *GetDeviceDatasourceInstanceGraphDataParams) WithDeviceID(deviceID int32) *GetDeviceDatasourceInstanceGraphDataParams {
	o.SetDeviceID(deviceID)
	return o
}

// SetDeviceID adds the deviceId to the get device datasource instance graph data params
func (o *GetDeviceDatasourceInstanceGraphDataParams) SetDeviceID(deviceID int32) {
	o.DeviceID = deviceID
}

// WithEnd adds the end to the get device datasource instance graph data params
func (o *GetDeviceDatasourceInstanceGraphDataParams) WithEnd(end *int64) *GetDeviceDatasourceInstanceGraphDataParams {
	o.SetEnd(end)
	return o
}

// SetEnd adds the end to the get device datasource instance graph data params
func (o *GetDeviceDatasourceInstanceGraphDataParams) SetEnd(end *int64) {
	o.End = end
}

// WithFormat adds the format to the get device datasource instance graph data params
func (o *GetDeviceDatasourceInstanceGraphDataParams) WithFormat(format *string) *GetDeviceDatasourceInstanceGraphDataParams {
	o.SetFormat(format)
	return o
}

// SetFormat adds the format to the get device datasource instance graph data params
func (o *GetDeviceDatasourceInstanceGraphDataParams) SetFormat(format *string) {
	o.Format = format
}

// WithGraphID adds the graphID to the get device datasource instance graph data params
func (o *GetDeviceDatasourceInstanceGraphDataParams) WithGraphID(graphID int32) *GetDeviceDatasourceInstanceGraphDataParams {
	o.SetGraphID(graphID)
	return o
}

// SetGraphID adds the graphId to the get device datasource instance graph data params
func (o *GetDeviceDatasourceInstanceGraphDataParams) SetGraphID(graphID int32) {
	o.GraphID = graphID
}

// WithHdsID adds the hdsID to the get device datasource instance graph data params
func (o *GetDeviceDatasourceInstanceGraphDataParams) WithHdsID(hdsID int32) *GetDeviceDatasourceInstanceGraphDataParams {
	o.SetHdsID(hdsID)
	return o
}

// SetHdsID adds the hdsId to the get device datasource instance graph data params
func (o *GetDeviceDatasourceInstanceGraphDataParams) SetHdsID(hdsID int32) {
	o.HdsID = hdsID
}

// WithID adds the id to the get device datasource instance graph data params
func (o *GetDeviceDatasourceInstanceGraphDataParams) WithID(id int32) *GetDeviceDatasourceInstanceGraphDataParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get device datasource instance graph data params
func (o *GetDeviceDatasourceInstanceGraphDataParams) SetID(id int32) {
	o.ID = id
}

// WithStart adds the start to the get device datasource instance graph data params
func (o *GetDeviceDatasourceInstanceGraphDataParams) WithStart(start *int64) *GetDeviceDatasourceInstanceGraphDataParams {
	o.SetStart(start)
	return o
}

// SetStart adds the start to the get device datasource instance graph data params
func (o *GetDeviceDatasourceInstanceGraphDataParams) SetStart(start *int64) {
	o.Start = start
}

// WriteToRequest writes these params to a swagger request
func (o *GetDeviceDatasourceInstanceGraphDataParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param deviceId
	if err := r.SetPathParam("deviceId", swag.FormatInt32(o.DeviceID)); err != nil {
		return err
	}

	if o.End != nil {

		// query param end
		var qrEnd int64
		if o.End != nil {
			qrEnd = *o.End
		}
		qEnd := swag.FormatInt64(qrEnd)
		if qEnd != "" {
			if err := r.SetQueryParam("end", qEnd); err != nil {
				return err
			}
		}

	}

	if o.Format != nil {

		// query param format
		var qrFormat string
		if o.Format != nil {
			qrFormat = *o.Format
		}
		qFormat := qrFormat
		if qFormat != "" {
			if err := r.SetQueryParam("format", qFormat); err != nil {
				return err
			}
		}

	}

	// path param graphId
	if err := r.SetPathParam("graphId", swag.FormatInt32(o.GraphID)); err != nil {
		return err
	}

	// path param hdsId
	if err := r.SetPathParam("hdsId", swag.FormatInt32(o.HdsID)); err != nil {
		return err
	}

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt32(o.ID)); err != nil {
		return err
	}

	if o.Start != nil {

		// query param start
		var qrStart int64
		if o.Start != nil {
			qrStart = *o.Start
		}
		qStart := swag.FormatInt64(qrStart)
		if qStart != "" {
			if err := r.SetQueryParam("start", qStart); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}