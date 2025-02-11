// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// AzureKeyVaultKey azure key vault key
//
// swagger:model azure_key_vault_key
type AzureKeyVaultKey struct {

	// Key identifier of the AKV key encryption key.
	// Example: https://keyvault1.vault.azure.net/keys/key1
	// Format: uri
	KeyID strfmt.URI `json:"key_id,omitempty"`

	// Set to "svm" for interfaces owned by an SVM. Otherwise, set to "cluster".
	// Enum: [svm cluster]
	Scope string `json:"scope,omitempty"`

	// svm
	Svm *AzureKeyVaultKeySvm `json:"svm,omitempty"`
}

// Validate validates this azure key vault key
func (m *AzureKeyVaultKey) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateKeyID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateScope(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSvm(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AzureKeyVaultKey) validateKeyID(formats strfmt.Registry) error {
	if swag.IsZero(m.KeyID) { // not required
		return nil
	}

	if err := validate.FormatOf("key_id", "body", "uri", m.KeyID.String(), formats); err != nil {
		return err
	}

	return nil
}

var azureKeyVaultKeyTypeScopePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["svm","cluster"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		azureKeyVaultKeyTypeScopePropEnum = append(azureKeyVaultKeyTypeScopePropEnum, v)
	}
}

const (

	// BEGIN DEBUGGING
	// azure_key_vault_key
	// AzureKeyVaultKey
	// scope
	// Scope
	// svm
	// END DEBUGGING
	// AzureKeyVaultKeyScopeSvm captures enum value "svm"
	AzureKeyVaultKeyScopeSvm string = "svm"

	// BEGIN DEBUGGING
	// azure_key_vault_key
	// AzureKeyVaultKey
	// scope
	// Scope
	// cluster
	// END DEBUGGING
	// AzureKeyVaultKeyScopeCluster captures enum value "cluster"
	AzureKeyVaultKeyScopeCluster string = "cluster"
)

// prop value enum
func (m *AzureKeyVaultKey) validateScopeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, azureKeyVaultKeyTypeScopePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *AzureKeyVaultKey) validateScope(formats strfmt.Registry) error {
	if swag.IsZero(m.Scope) { // not required
		return nil
	}

	// value enum
	if err := m.validateScopeEnum("scope", "body", m.Scope); err != nil {
		return err
	}

	return nil
}

func (m *AzureKeyVaultKey) validateSvm(formats strfmt.Registry) error {
	if swag.IsZero(m.Svm) { // not required
		return nil
	}

	if m.Svm != nil {
		if err := m.Svm.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("svm")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this azure key vault key based on the context it is used
func (m *AzureKeyVaultKey) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSvm(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AzureKeyVaultKey) contextValidateSvm(ctx context.Context, formats strfmt.Registry) error {

	if m.Svm != nil {
		if err := m.Svm.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("svm")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AzureKeyVaultKey) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AzureKeyVaultKey) UnmarshalBinary(b []byte) error {
	var res AzureKeyVaultKey
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// AzureKeyVaultKeySvm azure key vault key svm
//
// swagger:model AzureKeyVaultKeySvm
type AzureKeyVaultKeySvm struct {

	// links
	Links *AzureKeyVaultKeySvmLinks `json:"_links,omitempty"`

	// The name of the SVM.
	//
	// Example: svm1
	Name string `json:"name,omitempty"`

	// The unique identifier of the SVM.
	//
	// Example: 02c9e252-41be-11e9-81d5-00a0986138f7
	UUID string `json:"uuid,omitempty"`
}

// Validate validates this azure key vault key svm
func (m *AzureKeyVaultKeySvm) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AzureKeyVaultKeySvm) validateLinks(formats strfmt.Registry) error {
	if swag.IsZero(m.Links) { // not required
		return nil
	}

	if m.Links != nil {
		if err := m.Links.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("svm" + "." + "_links")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this azure key vault key svm based on the context it is used
func (m *AzureKeyVaultKeySvm) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AzureKeyVaultKeySvm) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

	if m.Links != nil {
		if err := m.Links.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("svm" + "." + "_links")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AzureKeyVaultKeySvm) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AzureKeyVaultKeySvm) UnmarshalBinary(b []byte) error {
	var res AzureKeyVaultKeySvm
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// AzureKeyVaultKeySvmLinks azure key vault key svm links
//
// swagger:model AzureKeyVaultKeySvmLinks
type AzureKeyVaultKeySvmLinks struct {

	// self
	Self *Href `json:"self,omitempty"`
}

// Validate validates this azure key vault key svm links
func (m *AzureKeyVaultKeySvmLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AzureKeyVaultKeySvmLinks) validateSelf(formats strfmt.Registry) error {
	if swag.IsZero(m.Self) { // not required
		return nil
	}

	if m.Self != nil {
		if err := m.Self.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("svm" + "." + "_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this azure key vault key svm links based on the context it is used
func (m *AzureKeyVaultKeySvmLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AzureKeyVaultKeySvmLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

	if m.Self != nil {
		if err := m.Self.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("svm" + "." + "_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AzureKeyVaultKeySvmLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AzureKeyVaultKeySvmLinks) UnmarshalBinary(b []byte) error {
	var res AzureKeyVaultKeySvmLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
