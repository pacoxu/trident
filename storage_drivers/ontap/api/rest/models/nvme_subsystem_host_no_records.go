// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// NvmeSubsystemHostNoRecords The NVMe host provisioned to access NVMe namespaces mapped to a subsystem.
//
// swagger:model nvme_subsystem_host_no_records
type NvmeSubsystemHostNoRecords struct {

	// links
	Links *NvmeSubsystemHostNoRecordsLinks `json:"_links,omitempty"`

	// io queue
	IoQueue *NvmeSubsystemHostNoRecordsIoQueue `json:"io_queue,omitempty"`

	// The NVMe qualified name (NQN) used to identify the NVMe storage target. Not allowed in POST when the `records` property is used.
	//
	// Example: nqn.1992-01.example.com:string
	Nqn string `json:"nqn,omitempty"`

	// subsystem
	Subsystem *NvmeSubsystemHostNoRecordsSubsystem `json:"subsystem,omitempty"`
}

// Validate validates this nvme subsystem host no records
func (m *NvmeSubsystemHostNoRecords) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIoQueue(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSubsystem(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NvmeSubsystemHostNoRecords) validateLinks(formats strfmt.Registry) error {
	if swag.IsZero(m.Links) { // not required
		return nil
	}

	if m.Links != nil {
		if err := m.Links.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("_links")
			}
			return err
		}
	}

	return nil
}

func (m *NvmeSubsystemHostNoRecords) validateIoQueue(formats strfmt.Registry) error {
	if swag.IsZero(m.IoQueue) { // not required
		return nil
	}

	if m.IoQueue != nil {
		if err := m.IoQueue.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("io_queue")
			}
			return err
		}
	}

	return nil
}

func (m *NvmeSubsystemHostNoRecords) validateSubsystem(formats strfmt.Registry) error {
	if swag.IsZero(m.Subsystem) { // not required
		return nil
	}

	if m.Subsystem != nil {
		if err := m.Subsystem.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("subsystem")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this nvme subsystem host no records based on the context it is used
func (m *NvmeSubsystemHostNoRecords) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateIoQueue(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSubsystem(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NvmeSubsystemHostNoRecords) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

	if m.Links != nil {
		if err := m.Links.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("_links")
			}
			return err
		}
	}

	return nil
}

func (m *NvmeSubsystemHostNoRecords) contextValidateIoQueue(ctx context.Context, formats strfmt.Registry) error {

	if m.IoQueue != nil {
		if err := m.IoQueue.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("io_queue")
			}
			return err
		}
	}

	return nil
}

func (m *NvmeSubsystemHostNoRecords) contextValidateSubsystem(ctx context.Context, formats strfmt.Registry) error {

	if m.Subsystem != nil {
		if err := m.Subsystem.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("subsystem")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *NvmeSubsystemHostNoRecords) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NvmeSubsystemHostNoRecords) UnmarshalBinary(b []byte) error {
	var res NvmeSubsystemHostNoRecords
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// NvmeSubsystemHostNoRecordsIoQueue The properties of the submission queue used to submit I/O commands for execution by the NVMe controller.
//
// swagger:model NvmeSubsystemHostNoRecordsIoQueue
type NvmeSubsystemHostNoRecordsIoQueue struct {

	// The number of I/O queue pairs. The default value is inherited from the owning NVMe subsystem.
	//
	// Example: 4
	// Read Only: true
	// Maximum: 15
	// Minimum: 1
	Count int64 `json:"count,omitempty"`

	// The I/O queue depth. The default value is inherited from the owning NVMe subsystem.
	//
	// Example: 32
	// Read Only: true
	// Maximum: 128
	// Minimum: 16
	Depth int64 `json:"depth,omitempty"`
}

// Validate validates this nvme subsystem host no records io queue
func (m *NvmeSubsystemHostNoRecordsIoQueue) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDepth(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NvmeSubsystemHostNoRecordsIoQueue) validateCount(formats strfmt.Registry) error {
	if swag.IsZero(m.Count) { // not required
		return nil
	}

	if err := validate.MinimumInt("io_queue"+"."+"count", "body", m.Count, 1, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("io_queue"+"."+"count", "body", m.Count, 15, false); err != nil {
		return err
	}

	return nil
}

func (m *NvmeSubsystemHostNoRecordsIoQueue) validateDepth(formats strfmt.Registry) error {
	if swag.IsZero(m.Depth) { // not required
		return nil
	}

	if err := validate.MinimumInt("io_queue"+"."+"depth", "body", m.Depth, 16, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("io_queue"+"."+"depth", "body", m.Depth, 128, false); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this nvme subsystem host no records io queue based on the context it is used
func (m *NvmeSubsystemHostNoRecordsIoQueue) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCount(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDepth(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NvmeSubsystemHostNoRecordsIoQueue) contextValidateCount(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "io_queue"+"."+"count", "body", int64(m.Count)); err != nil {
		return err
	}

	return nil
}

func (m *NvmeSubsystemHostNoRecordsIoQueue) contextValidateDepth(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "io_queue"+"."+"depth", "body", int64(m.Depth)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *NvmeSubsystemHostNoRecordsIoQueue) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NvmeSubsystemHostNoRecordsIoQueue) UnmarshalBinary(b []byte) error {
	var res NvmeSubsystemHostNoRecordsIoQueue
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// NvmeSubsystemHostNoRecordsLinks nvme subsystem host no records links
//
// swagger:model NvmeSubsystemHostNoRecordsLinks
type NvmeSubsystemHostNoRecordsLinks struct {

	// self
	Self *Href `json:"self,omitempty"`
}

// Validate validates this nvme subsystem host no records links
func (m *NvmeSubsystemHostNoRecordsLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NvmeSubsystemHostNoRecordsLinks) validateSelf(formats strfmt.Registry) error {
	if swag.IsZero(m.Self) { // not required
		return nil
	}

	if m.Self != nil {
		if err := m.Self.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this nvme subsystem host no records links based on the context it is used
func (m *NvmeSubsystemHostNoRecordsLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NvmeSubsystemHostNoRecordsLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

	if m.Self != nil {
		if err := m.Self.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *NvmeSubsystemHostNoRecordsLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NvmeSubsystemHostNoRecordsLinks) UnmarshalBinary(b []byte) error {
	var res NvmeSubsystemHostNoRecordsLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// NvmeSubsystemHostNoRecordsSubsystem The NVMe subsystem to which the NVMe host has been provisioned.
//
// swagger:model NvmeSubsystemHostNoRecordsSubsystem
type NvmeSubsystemHostNoRecordsSubsystem struct {

	// links
	Links *NvmeSubsystemHostNoRecordsSubsystemLinks `json:"_links,omitempty"`

	// The unique identifier of the NVMe subsystem.
	//
	// Example: 1cd8a442-86d1-11e0-ae1c-123478563412
	// Read Only: true
	UUID string `json:"uuid,omitempty"`
}

// Validate validates this nvme subsystem host no records subsystem
func (m *NvmeSubsystemHostNoRecordsSubsystem) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NvmeSubsystemHostNoRecordsSubsystem) validateLinks(formats strfmt.Registry) error {
	if swag.IsZero(m.Links) { // not required
		return nil
	}

	if m.Links != nil {
		if err := m.Links.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("subsystem" + "." + "_links")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this nvme subsystem host no records subsystem based on the context it is used
func (m *NvmeSubsystemHostNoRecordsSubsystem) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUUID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NvmeSubsystemHostNoRecordsSubsystem) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

	if m.Links != nil {
		if err := m.Links.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("subsystem" + "." + "_links")
			}
			return err
		}
	}

	return nil
}

func (m *NvmeSubsystemHostNoRecordsSubsystem) contextValidateUUID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "subsystem"+"."+"uuid", "body", string(m.UUID)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *NvmeSubsystemHostNoRecordsSubsystem) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NvmeSubsystemHostNoRecordsSubsystem) UnmarshalBinary(b []byte) error {
	var res NvmeSubsystemHostNoRecordsSubsystem
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// NvmeSubsystemHostNoRecordsSubsystemLinks nvme subsystem host no records subsystem links
//
// swagger:model NvmeSubsystemHostNoRecordsSubsystemLinks
type NvmeSubsystemHostNoRecordsSubsystemLinks struct {

	// self
	Self *Href `json:"self,omitempty"`
}

// Validate validates this nvme subsystem host no records subsystem links
func (m *NvmeSubsystemHostNoRecordsSubsystemLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NvmeSubsystemHostNoRecordsSubsystemLinks) validateSelf(formats strfmt.Registry) error {
	if swag.IsZero(m.Self) { // not required
		return nil
	}

	if m.Self != nil {
		if err := m.Self.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("subsystem" + "." + "_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this nvme subsystem host no records subsystem links based on the context it is used
func (m *NvmeSubsystemHostNoRecordsSubsystemLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NvmeSubsystemHostNoRecordsSubsystemLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

	if m.Self != nil {
		if err := m.Self.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("subsystem" + "." + "_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *NvmeSubsystemHostNoRecordsSubsystemLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NvmeSubsystemHostNoRecordsSubsystemLinks) UnmarshalBinary(b []byte) error {
	var res NvmeSubsystemHostNoRecordsSubsystemLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
