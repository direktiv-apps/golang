// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PostParamsBodyCommandsItems post params body commands items
//
// swagger:model postParamsBodyCommandsItems
type PostParamsBodyCommandsItems struct {

	// Command to run
	// Example: go version
	Command string `json:"command,omitempty"`

	// Stops excecution if command fails, otherwise proceeds with next command
	Continue bool `json:"continue,omitempty"`

	// If set to false the command will not print the full command with arguments to logs.
	Print *bool `json:"print,omitempty"`

	// If set to false the command will not print output to logs.
	Silent *bool `json:"silent,omitempty"`
}

// Validate validates this post params body commands items
func (m *PostParamsBodyCommandsItems) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this post params body commands items based on context it is used
func (m *PostParamsBodyCommandsItems) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PostParamsBodyCommandsItems) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PostParamsBodyCommandsItems) UnmarshalBinary(b []byte) error {
	var res PostParamsBodyCommandsItems
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
