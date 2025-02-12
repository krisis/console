// Code generated by go-swagger; DO NOT EDIT.

// This file is part of MinIO Console Server
// Copyright (c) 2023 MinIO, Inc.
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.
//

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PoolAffinity If specified, affinity will define the pod's scheduling constraints
//
// swagger:model poolAffinity
type PoolAffinity struct {

	// node affinity
	NodeAffinity *PoolAffinityNodeAffinity `json:"nodeAffinity,omitempty"`

	// pod affinity
	PodAffinity *PoolAffinityPodAffinity `json:"podAffinity,omitempty"`

	// pod anti affinity
	PodAntiAffinity *PoolAffinityPodAntiAffinity `json:"podAntiAffinity,omitempty"`
}

// Validate validates this pool affinity
func (m *PoolAffinity) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateNodeAffinity(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePodAffinity(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePodAntiAffinity(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PoolAffinity) validateNodeAffinity(formats strfmt.Registry) error {
	if swag.IsZero(m.NodeAffinity) { // not required
		return nil
	}

	if m.NodeAffinity != nil {
		if err := m.NodeAffinity.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("nodeAffinity")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("nodeAffinity")
			}
			return err
		}
	}

	return nil
}

func (m *PoolAffinity) validatePodAffinity(formats strfmt.Registry) error {
	if swag.IsZero(m.PodAffinity) { // not required
		return nil
	}

	if m.PodAffinity != nil {
		if err := m.PodAffinity.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("podAffinity")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("podAffinity")
			}
			return err
		}
	}

	return nil
}

func (m *PoolAffinity) validatePodAntiAffinity(formats strfmt.Registry) error {
	if swag.IsZero(m.PodAntiAffinity) { // not required
		return nil
	}

	if m.PodAntiAffinity != nil {
		if err := m.PodAntiAffinity.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("podAntiAffinity")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("podAntiAffinity")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this pool affinity based on the context it is used
func (m *PoolAffinity) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateNodeAffinity(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePodAffinity(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePodAntiAffinity(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PoolAffinity) contextValidateNodeAffinity(ctx context.Context, formats strfmt.Registry) error {

	if m.NodeAffinity != nil {
		if err := m.NodeAffinity.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("nodeAffinity")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("nodeAffinity")
			}
			return err
		}
	}

	return nil
}

func (m *PoolAffinity) contextValidatePodAffinity(ctx context.Context, formats strfmt.Registry) error {

	if m.PodAffinity != nil {
		if err := m.PodAffinity.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("podAffinity")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("podAffinity")
			}
			return err
		}
	}

	return nil
}

func (m *PoolAffinity) contextValidatePodAntiAffinity(ctx context.Context, formats strfmt.Registry) error {

	if m.PodAntiAffinity != nil {
		if err := m.PodAntiAffinity.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("podAntiAffinity")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("podAntiAffinity")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PoolAffinity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PoolAffinity) UnmarshalBinary(b []byte) error {
	var res PoolAffinity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PoolAffinityNodeAffinity Describes node affinity scheduling rules for the pod.
//
// swagger:model PoolAffinityNodeAffinity
type PoolAffinityNodeAffinity struct {

	// The scheduler will prefer to schedule pods to nodes that satisfy the affinity expressions specified by this field, but it may choose a node that violates one or more of the expressions. The node that is most preferred is the one with the greatest sum of weights, i.e. for each node that meets all of the scheduling requirements (resource request, requiredDuringScheduling affinity expressions, etc.), compute a sum by iterating through the elements of this field and adding "weight" to the sum if the node matches the corresponding matchExpressions; the node(s) with the highest sum are the most preferred.
	PreferredDuringSchedulingIgnoredDuringExecution []*PoolAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionItems0 `json:"preferredDuringSchedulingIgnoredDuringExecution"`

	// required during scheduling ignored during execution
	RequiredDuringSchedulingIgnoredDuringExecution *PoolAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecution `json:"requiredDuringSchedulingIgnoredDuringExecution,omitempty"`
}

// Validate validates this pool affinity node affinity
func (m *PoolAffinityNodeAffinity) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePreferredDuringSchedulingIgnoredDuringExecution(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRequiredDuringSchedulingIgnoredDuringExecution(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PoolAffinityNodeAffinity) validatePreferredDuringSchedulingIgnoredDuringExecution(formats strfmt.Registry) error {
	if swag.IsZero(m.PreferredDuringSchedulingIgnoredDuringExecution) { // not required
		return nil
	}

	for i := 0; i < len(m.PreferredDuringSchedulingIgnoredDuringExecution); i++ {
		if swag.IsZero(m.PreferredDuringSchedulingIgnoredDuringExecution[i]) { // not required
			continue
		}

		if m.PreferredDuringSchedulingIgnoredDuringExecution[i] != nil {
			if err := m.PreferredDuringSchedulingIgnoredDuringExecution[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("nodeAffinity" + "." + "preferredDuringSchedulingIgnoredDuringExecution" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("nodeAffinity" + "." + "preferredDuringSchedulingIgnoredDuringExecution" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PoolAffinityNodeAffinity) validateRequiredDuringSchedulingIgnoredDuringExecution(formats strfmt.Registry) error {
	if swag.IsZero(m.RequiredDuringSchedulingIgnoredDuringExecution) { // not required
		return nil
	}

	if m.RequiredDuringSchedulingIgnoredDuringExecution != nil {
		if err := m.RequiredDuringSchedulingIgnoredDuringExecution.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("nodeAffinity" + "." + "requiredDuringSchedulingIgnoredDuringExecution")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("nodeAffinity" + "." + "requiredDuringSchedulingIgnoredDuringExecution")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this pool affinity node affinity based on the context it is used
func (m *PoolAffinityNodeAffinity) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidatePreferredDuringSchedulingIgnoredDuringExecution(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRequiredDuringSchedulingIgnoredDuringExecution(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PoolAffinityNodeAffinity) contextValidatePreferredDuringSchedulingIgnoredDuringExecution(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.PreferredDuringSchedulingIgnoredDuringExecution); i++ {

		if m.PreferredDuringSchedulingIgnoredDuringExecution[i] != nil {
			if err := m.PreferredDuringSchedulingIgnoredDuringExecution[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("nodeAffinity" + "." + "preferredDuringSchedulingIgnoredDuringExecution" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("nodeAffinity" + "." + "preferredDuringSchedulingIgnoredDuringExecution" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PoolAffinityNodeAffinity) contextValidateRequiredDuringSchedulingIgnoredDuringExecution(ctx context.Context, formats strfmt.Registry) error {

	if m.RequiredDuringSchedulingIgnoredDuringExecution != nil {
		if err := m.RequiredDuringSchedulingIgnoredDuringExecution.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("nodeAffinity" + "." + "requiredDuringSchedulingIgnoredDuringExecution")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("nodeAffinity" + "." + "requiredDuringSchedulingIgnoredDuringExecution")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PoolAffinityNodeAffinity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PoolAffinityNodeAffinity) UnmarshalBinary(b []byte) error {
	var res PoolAffinityNodeAffinity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PoolAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionItems0 An empty preferred scheduling term matches all objects with implicit weight 0 (i.e. it's a no-op). A null preferred scheduling term matches no objects (i.e. is also a no-op).
//
// swagger:model PoolAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionItems0
type PoolAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionItems0 struct {

	// A node selector term, associated with the corresponding weight.
	// Required: true
	Preference *NodeSelectorTerm `json:"preference"`

	// Weight associated with matching the corresponding nodeSelectorTerm, in the range 1-100.
	// Required: true
	Weight *int32 `json:"weight"`
}

// Validate validates this pool affinity node affinity preferred during scheduling ignored during execution items0
func (m *PoolAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePreference(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWeight(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PoolAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionItems0) validatePreference(formats strfmt.Registry) error {

	if err := validate.Required("preference", "body", m.Preference); err != nil {
		return err
	}

	if m.Preference != nil {
		if err := m.Preference.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("preference")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("preference")
			}
			return err
		}
	}

	return nil
}

func (m *PoolAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionItems0) validateWeight(formats strfmt.Registry) error {

	if err := validate.Required("weight", "body", m.Weight); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this pool affinity node affinity preferred during scheduling ignored during execution items0 based on the context it is used
func (m *PoolAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidatePreference(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PoolAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionItems0) contextValidatePreference(ctx context.Context, formats strfmt.Registry) error {

	if m.Preference != nil {
		if err := m.Preference.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("preference")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("preference")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PoolAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PoolAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionItems0) UnmarshalBinary(b []byte) error {
	var res PoolAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PoolAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecution If the affinity requirements specified by this field are not met at scheduling time, the pod will not be scheduled onto the node. If the affinity requirements specified by this field cease to be met at some point during pod execution (e.g. due to an update), the system may or may not try to eventually evict the pod from its node.
//
// swagger:model PoolAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecution
type PoolAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecution struct {

	// Required. A list of node selector terms. The terms are ORed.
	// Required: true
	NodeSelectorTerms []*NodeSelectorTerm `json:"nodeSelectorTerms"`
}

// Validate validates this pool affinity node affinity required during scheduling ignored during execution
func (m *PoolAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecution) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateNodeSelectorTerms(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PoolAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecution) validateNodeSelectorTerms(formats strfmt.Registry) error {

	if err := validate.Required("nodeAffinity"+"."+"requiredDuringSchedulingIgnoredDuringExecution"+"."+"nodeSelectorTerms", "body", m.NodeSelectorTerms); err != nil {
		return err
	}

	for i := 0; i < len(m.NodeSelectorTerms); i++ {
		if swag.IsZero(m.NodeSelectorTerms[i]) { // not required
			continue
		}

		if m.NodeSelectorTerms[i] != nil {
			if err := m.NodeSelectorTerms[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("nodeAffinity" + "." + "requiredDuringSchedulingIgnoredDuringExecution" + "." + "nodeSelectorTerms" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("nodeAffinity" + "." + "requiredDuringSchedulingIgnoredDuringExecution" + "." + "nodeSelectorTerms" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this pool affinity node affinity required during scheduling ignored during execution based on the context it is used
func (m *PoolAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecution) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateNodeSelectorTerms(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PoolAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecution) contextValidateNodeSelectorTerms(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.NodeSelectorTerms); i++ {

		if m.NodeSelectorTerms[i] != nil {
			if err := m.NodeSelectorTerms[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("nodeAffinity" + "." + "requiredDuringSchedulingIgnoredDuringExecution" + "." + "nodeSelectorTerms" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("nodeAffinity" + "." + "requiredDuringSchedulingIgnoredDuringExecution" + "." + "nodeSelectorTerms" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *PoolAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecution) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PoolAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecution) UnmarshalBinary(b []byte) error {
	var res PoolAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecution
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PoolAffinityPodAffinity Describes pod affinity scheduling rules (e.g. co-locate this pod in the same node, pool, etc. as some other pod(s)).
//
// swagger:model PoolAffinityPodAffinity
type PoolAffinityPodAffinity struct {

	// The scheduler will prefer to schedule pods to nodes that satisfy the affinity expressions specified by this field, but it may choose a node that violates one or more of the expressions. The node that is most preferred is the one with the greatest sum of weights, i.e. for each node that meets all of the scheduling requirements (resource request, requiredDuringScheduling affinity expressions, etc.), compute a sum by iterating through the elements of this field and adding "weight" to the sum if the node has pods which matches the corresponding podAffinityTerm; the node(s) with the highest sum are the most preferred.
	PreferredDuringSchedulingIgnoredDuringExecution []*PoolAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionItems0 `json:"preferredDuringSchedulingIgnoredDuringExecution"`

	// If the affinity requirements specified by this field are not met at scheduling time, the pod will not be scheduled onto the node. If the affinity requirements specified by this field cease to be met at some point during pod execution (e.g. due to a pod label update), the system may or may not try to eventually evict the pod from its node. When there are multiple elements, the lists of nodes corresponding to each podAffinityTerm are intersected, i.e. all terms must be satisfied.
	RequiredDuringSchedulingIgnoredDuringExecution []*PodAffinityTerm `json:"requiredDuringSchedulingIgnoredDuringExecution"`
}

// Validate validates this pool affinity pod affinity
func (m *PoolAffinityPodAffinity) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePreferredDuringSchedulingIgnoredDuringExecution(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRequiredDuringSchedulingIgnoredDuringExecution(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PoolAffinityPodAffinity) validatePreferredDuringSchedulingIgnoredDuringExecution(formats strfmt.Registry) error {
	if swag.IsZero(m.PreferredDuringSchedulingIgnoredDuringExecution) { // not required
		return nil
	}

	for i := 0; i < len(m.PreferredDuringSchedulingIgnoredDuringExecution); i++ {
		if swag.IsZero(m.PreferredDuringSchedulingIgnoredDuringExecution[i]) { // not required
			continue
		}

		if m.PreferredDuringSchedulingIgnoredDuringExecution[i] != nil {
			if err := m.PreferredDuringSchedulingIgnoredDuringExecution[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("podAffinity" + "." + "preferredDuringSchedulingIgnoredDuringExecution" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("podAffinity" + "." + "preferredDuringSchedulingIgnoredDuringExecution" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PoolAffinityPodAffinity) validateRequiredDuringSchedulingIgnoredDuringExecution(formats strfmt.Registry) error {
	if swag.IsZero(m.RequiredDuringSchedulingIgnoredDuringExecution) { // not required
		return nil
	}

	for i := 0; i < len(m.RequiredDuringSchedulingIgnoredDuringExecution); i++ {
		if swag.IsZero(m.RequiredDuringSchedulingIgnoredDuringExecution[i]) { // not required
			continue
		}

		if m.RequiredDuringSchedulingIgnoredDuringExecution[i] != nil {
			if err := m.RequiredDuringSchedulingIgnoredDuringExecution[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("podAffinity" + "." + "requiredDuringSchedulingIgnoredDuringExecution" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("podAffinity" + "." + "requiredDuringSchedulingIgnoredDuringExecution" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this pool affinity pod affinity based on the context it is used
func (m *PoolAffinityPodAffinity) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidatePreferredDuringSchedulingIgnoredDuringExecution(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRequiredDuringSchedulingIgnoredDuringExecution(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PoolAffinityPodAffinity) contextValidatePreferredDuringSchedulingIgnoredDuringExecution(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.PreferredDuringSchedulingIgnoredDuringExecution); i++ {

		if m.PreferredDuringSchedulingIgnoredDuringExecution[i] != nil {
			if err := m.PreferredDuringSchedulingIgnoredDuringExecution[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("podAffinity" + "." + "preferredDuringSchedulingIgnoredDuringExecution" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("podAffinity" + "." + "preferredDuringSchedulingIgnoredDuringExecution" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PoolAffinityPodAffinity) contextValidateRequiredDuringSchedulingIgnoredDuringExecution(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.RequiredDuringSchedulingIgnoredDuringExecution); i++ {

		if m.RequiredDuringSchedulingIgnoredDuringExecution[i] != nil {
			if err := m.RequiredDuringSchedulingIgnoredDuringExecution[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("podAffinity" + "." + "requiredDuringSchedulingIgnoredDuringExecution" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("podAffinity" + "." + "requiredDuringSchedulingIgnoredDuringExecution" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *PoolAffinityPodAffinity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PoolAffinityPodAffinity) UnmarshalBinary(b []byte) error {
	var res PoolAffinityPodAffinity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PoolAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionItems0 The weights of all of the matched WeightedPodAffinityTerm fields are added per-node to find the most preferred node(s)
//
// swagger:model PoolAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionItems0
type PoolAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionItems0 struct {

	// pod affinity term
	// Required: true
	PodAffinityTerm *PodAffinityTerm `json:"podAffinityTerm"`

	// weight associated with matching the corresponding podAffinityTerm, in the range 1-100.
	// Required: true
	Weight *int32 `json:"weight"`
}

// Validate validates this pool affinity pod affinity preferred during scheduling ignored during execution items0
func (m *PoolAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePodAffinityTerm(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWeight(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PoolAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionItems0) validatePodAffinityTerm(formats strfmt.Registry) error {

	if err := validate.Required("podAffinityTerm", "body", m.PodAffinityTerm); err != nil {
		return err
	}

	if m.PodAffinityTerm != nil {
		if err := m.PodAffinityTerm.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("podAffinityTerm")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("podAffinityTerm")
			}
			return err
		}
	}

	return nil
}

func (m *PoolAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionItems0) validateWeight(formats strfmt.Registry) error {

	if err := validate.Required("weight", "body", m.Weight); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this pool affinity pod affinity preferred during scheduling ignored during execution items0 based on the context it is used
func (m *PoolAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidatePodAffinityTerm(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PoolAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionItems0) contextValidatePodAffinityTerm(ctx context.Context, formats strfmt.Registry) error {

	if m.PodAffinityTerm != nil {
		if err := m.PodAffinityTerm.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("podAffinityTerm")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("podAffinityTerm")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PoolAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PoolAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionItems0) UnmarshalBinary(b []byte) error {
	var res PoolAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PoolAffinityPodAntiAffinity Describes pod anti-affinity scheduling rules (e.g. avoid putting this pod in the same node, pool, etc. as some other pod(s)).
//
// swagger:model PoolAffinityPodAntiAffinity
type PoolAffinityPodAntiAffinity struct {

	// The scheduler will prefer to schedule pods to nodes that satisfy the anti-affinity expressions specified by this field, but it may choose a node that violates one or more of the expressions. The node that is most preferred is the one with the greatest sum of weights, i.e. for each node that meets all of the scheduling requirements (resource request, requiredDuringScheduling anti-affinity expressions, etc.), compute a sum by iterating through the elements of this field and adding "weight" to the sum if the node has pods which matches the corresponding podAffinityTerm; the node(s) with the highest sum are the most preferred.
	PreferredDuringSchedulingIgnoredDuringExecution []*PoolAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionItems0 `json:"preferredDuringSchedulingIgnoredDuringExecution"`

	// If the anti-affinity requirements specified by this field are not met at scheduling time, the pod will not be scheduled onto the node. If the anti-affinity requirements specified by this field cease to be met at some point during pod execution (e.g. due to a pod label update), the system may or may not try to eventually evict the pod from its node. When there are multiple elements, the lists of nodes corresponding to each podAffinityTerm are intersected, i.e. all terms must be satisfied.
	RequiredDuringSchedulingIgnoredDuringExecution []*PodAffinityTerm `json:"requiredDuringSchedulingIgnoredDuringExecution"`
}

// Validate validates this pool affinity pod anti affinity
func (m *PoolAffinityPodAntiAffinity) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePreferredDuringSchedulingIgnoredDuringExecution(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRequiredDuringSchedulingIgnoredDuringExecution(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PoolAffinityPodAntiAffinity) validatePreferredDuringSchedulingIgnoredDuringExecution(formats strfmt.Registry) error {
	if swag.IsZero(m.PreferredDuringSchedulingIgnoredDuringExecution) { // not required
		return nil
	}

	for i := 0; i < len(m.PreferredDuringSchedulingIgnoredDuringExecution); i++ {
		if swag.IsZero(m.PreferredDuringSchedulingIgnoredDuringExecution[i]) { // not required
			continue
		}

		if m.PreferredDuringSchedulingIgnoredDuringExecution[i] != nil {
			if err := m.PreferredDuringSchedulingIgnoredDuringExecution[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("podAntiAffinity" + "." + "preferredDuringSchedulingIgnoredDuringExecution" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("podAntiAffinity" + "." + "preferredDuringSchedulingIgnoredDuringExecution" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PoolAffinityPodAntiAffinity) validateRequiredDuringSchedulingIgnoredDuringExecution(formats strfmt.Registry) error {
	if swag.IsZero(m.RequiredDuringSchedulingIgnoredDuringExecution) { // not required
		return nil
	}

	for i := 0; i < len(m.RequiredDuringSchedulingIgnoredDuringExecution); i++ {
		if swag.IsZero(m.RequiredDuringSchedulingIgnoredDuringExecution[i]) { // not required
			continue
		}

		if m.RequiredDuringSchedulingIgnoredDuringExecution[i] != nil {
			if err := m.RequiredDuringSchedulingIgnoredDuringExecution[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("podAntiAffinity" + "." + "requiredDuringSchedulingIgnoredDuringExecution" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("podAntiAffinity" + "." + "requiredDuringSchedulingIgnoredDuringExecution" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this pool affinity pod anti affinity based on the context it is used
func (m *PoolAffinityPodAntiAffinity) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidatePreferredDuringSchedulingIgnoredDuringExecution(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRequiredDuringSchedulingIgnoredDuringExecution(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PoolAffinityPodAntiAffinity) contextValidatePreferredDuringSchedulingIgnoredDuringExecution(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.PreferredDuringSchedulingIgnoredDuringExecution); i++ {

		if m.PreferredDuringSchedulingIgnoredDuringExecution[i] != nil {
			if err := m.PreferredDuringSchedulingIgnoredDuringExecution[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("podAntiAffinity" + "." + "preferredDuringSchedulingIgnoredDuringExecution" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("podAntiAffinity" + "." + "preferredDuringSchedulingIgnoredDuringExecution" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PoolAffinityPodAntiAffinity) contextValidateRequiredDuringSchedulingIgnoredDuringExecution(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.RequiredDuringSchedulingIgnoredDuringExecution); i++ {

		if m.RequiredDuringSchedulingIgnoredDuringExecution[i] != nil {
			if err := m.RequiredDuringSchedulingIgnoredDuringExecution[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("podAntiAffinity" + "." + "requiredDuringSchedulingIgnoredDuringExecution" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("podAntiAffinity" + "." + "requiredDuringSchedulingIgnoredDuringExecution" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *PoolAffinityPodAntiAffinity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PoolAffinityPodAntiAffinity) UnmarshalBinary(b []byte) error {
	var res PoolAffinityPodAntiAffinity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PoolAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionItems0 The weights of all of the matched WeightedPodAffinityTerm fields are added per-node to find the most preferred node(s)
//
// swagger:model PoolAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionItems0
type PoolAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionItems0 struct {

	// pod affinity term
	// Required: true
	PodAffinityTerm *PodAffinityTerm `json:"podAffinityTerm"`

	// weight associated with matching the corresponding podAffinityTerm, in the range 1-100.
	// Required: true
	Weight *int32 `json:"weight"`
}

// Validate validates this pool affinity pod anti affinity preferred during scheduling ignored during execution items0
func (m *PoolAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePodAffinityTerm(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWeight(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PoolAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionItems0) validatePodAffinityTerm(formats strfmt.Registry) error {

	if err := validate.Required("podAffinityTerm", "body", m.PodAffinityTerm); err != nil {
		return err
	}

	if m.PodAffinityTerm != nil {
		if err := m.PodAffinityTerm.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("podAffinityTerm")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("podAffinityTerm")
			}
			return err
		}
	}

	return nil
}

func (m *PoolAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionItems0) validateWeight(formats strfmt.Registry) error {

	if err := validate.Required("weight", "body", m.Weight); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this pool affinity pod anti affinity preferred during scheduling ignored during execution items0 based on the context it is used
func (m *PoolAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidatePodAffinityTerm(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PoolAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionItems0) contextValidatePodAffinityTerm(ctx context.Context, formats strfmt.Registry) error {

	if m.PodAffinityTerm != nil {
		if err := m.PodAffinityTerm.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("podAffinityTerm")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("podAffinityTerm")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PoolAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PoolAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionItems0) UnmarshalBinary(b []byte) error {
	var res PoolAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
