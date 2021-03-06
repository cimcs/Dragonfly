// Code generated by go-swagger; DO NOT EDIT.

package types

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PreheatCreateRequest Request option of creating a preheat task in supernode.
//
// swagger:model PreheatCreateRequest
type PreheatCreateRequest struct {

	// ID
	ID string `json:"ID,omitempty"`
}

// Validate validates this preheat create request
func (m *PreheatCreateRequest) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PreheatCreateRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PreheatCreateRequest) UnmarshalBinary(b []byte) error {
	var res PreheatCreateRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
