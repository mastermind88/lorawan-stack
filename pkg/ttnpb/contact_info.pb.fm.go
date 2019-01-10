// Code generated by protoc-gen-fieldmask. DO NOT EDIT.

package ttnpb

import fmt "fmt"

var ContactInfoFieldPathsNested = []string{
	"contact_method",
	"contact_type",
	"public",
	"validated_at",
	"value",
}

var ContactInfoFieldPathsTopLevel = []string{
	"contact_method",
	"contact_type",
	"public",
	"validated_at",
	"value",
}

func (dst *ContactInfo) SetFields(src *ContactInfo, paths ...string) error {
	for name, subs := range _processPaths(append(paths[:0:0], paths...)) {
		switch name {
		case "contact_type":
			if len(subs) > 0 {
				return fmt.Errorf("'contact_type' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.ContactType = src.ContactType
			} else {
				var zero ContactType
				dst.ContactType = zero
			}
		case "contact_method":
			if len(subs) > 0 {
				return fmt.Errorf("'contact_method' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.ContactMethod = src.ContactMethod
			} else {
				var zero ContactMethod
				dst.ContactMethod = zero
			}
		case "value":
			if len(subs) > 0 {
				return fmt.Errorf("'value' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.Value = src.Value
			} else {
				var zero string
				dst.Value = zero
			}
		case "public":
			if len(subs) > 0 {
				return fmt.Errorf("'public' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.Public = src.Public
			} else {
				var zero bool
				dst.Public = zero
			}
		case "validated_at":
			if len(subs) > 0 {
				return fmt.Errorf("'validated_at' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.ValidatedAt = src.ValidatedAt
			} else {
				dst.ValidatedAt = nil
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

var ContactInfoValidationFieldPathsNested = []string{
	"contact_info",
	"created_at",
	"entity",
	"entity.ids.application_ids",
	"entity.ids.application_ids.application_id",
	"entity.ids.client_ids",
	"entity.ids.client_ids.client_id",
	"entity.ids.device_ids",
	"entity.ids.device_ids.application_ids",
	"entity.ids.device_ids.application_ids.application_id",
	"entity.ids.device_ids.dev_addr",
	"entity.ids.device_ids.dev_eui",
	"entity.ids.device_ids.device_id",
	"entity.ids.device_ids.join_eui",
	"entity.ids.gateway_ids",
	"entity.ids.gateway_ids.eui",
	"entity.ids.gateway_ids.gateway_id",
	"entity.ids.organization_ids",
	"entity.ids.organization_ids.organization_id",
	"entity.ids.user_ids",
	"entity.ids.user_ids.email",
	"entity.ids.user_ids.user_id",
	"expires_at",
	"id",
	"ids",
	"token",
}

var ContactInfoValidationFieldPathsTopLevel = []string{
	"contact_info",
	"created_at",
	"entity",
	"expires_at",
	"id",
	"ids",
	"token",
}

func (dst *ContactInfoValidation) SetFields(src *ContactInfoValidation, paths ...string) error {
	for name, subs := range _processPaths(append(paths[:0:0], paths...)) {
		switch name {
		case "id":
			if len(subs) > 0 {
				return fmt.Errorf("'id' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.ID = src.ID
			} else {
				var zero string
				dst.ID = zero
			}
		case "token":
			if len(subs) > 0 {
				return fmt.Errorf("'token' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.Token = src.Token
			} else {
				var zero string
				dst.Token = zero
			}
		case "entity":
			if len(subs) > 0 {
				newDst := dst.Entity
				if newDst == nil {
					newDst = &EntityIdentifiers{}
					dst.Entity = newDst
				}
				var newSrc *EntityIdentifiers
				if src != nil {
					newSrc = src.Entity
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.Entity = src.Entity
				} else {
					dst.Entity = nil
				}
			}
		case "contact_info":
			if len(subs) > 0 {
				return fmt.Errorf("'contact_info' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.ContactInfo = src.ContactInfo
			} else {
				dst.ContactInfo = nil
			}
		case "created_at":
			if len(subs) > 0 {
				return fmt.Errorf("'created_at' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.CreatedAt = src.CreatedAt
			} else {
				dst.CreatedAt = nil
			}
		case "expires_at":
			if len(subs) > 0 {
				return fmt.Errorf("'expires_at' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.ExpiresAt = src.ExpiresAt
			} else {
				dst.ExpiresAt = nil
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}
