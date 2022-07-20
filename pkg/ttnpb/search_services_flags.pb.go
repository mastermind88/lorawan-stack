// Code generated by protoc-gen-go-flags. DO NOT EDIT.
// versions:
// - protoc-gen-go-flags v1.0.3
// - protoc              v3.9.1
// source: lorawan-stack/api/search_services.proto

package ttnpb

import (
	flagsplugin "github.com/TheThingsIndustries/protoc-gen-go-flags/flagsplugin"
	gogo "github.com/TheThingsIndustries/protoc-gen-go-flags/gogo"
	pflag "github.com/spf13/pflag"
)

// AddSetFlagsForSearchApplicationsRequest adds flags to select fields in SearchApplicationsRequest.
func AddSetFlagsForSearchApplicationsRequest(flags *pflag.FlagSet, prefix string, hidden bool) {
	flags.AddFlag(flagsplugin.NewStringFlag(flagsplugin.Prefix("query", prefix), "", flagsplugin.WithHidden(hidden)))
	flags.AddFlag(flagsplugin.NewStringFlag(flagsplugin.Prefix("id-contains", prefix), "", flagsplugin.WithHidden(hidden)))
	flags.AddFlag(flagsplugin.NewStringFlag(flagsplugin.Prefix("name-contains", prefix), "", flagsplugin.WithHidden(hidden)))
	flags.AddFlag(flagsplugin.NewStringFlag(flagsplugin.Prefix("description-contains", prefix), "", flagsplugin.WithHidden(hidden)))
	flags.AddFlag(flagsplugin.NewStringStringMapFlag(flagsplugin.Prefix("attributes-contain", prefix), "", flagsplugin.WithHidden(hidden)))
	flags.AddFlag(flagsplugin.NewStringSliceFlag(flagsplugin.Prefix("field-mask", prefix), "", flagsplugin.WithHidden(hidden)))
	flags.AddFlag(flagsplugin.NewStringFlag(flagsplugin.Prefix("order", prefix), "", flagsplugin.WithHidden(hidden)))
	flags.AddFlag(flagsplugin.NewUint32Flag(flagsplugin.Prefix("limit", prefix), "", flagsplugin.WithHidden(hidden)))
	flags.AddFlag(flagsplugin.NewUint32Flag(flagsplugin.Prefix("page", prefix), "", flagsplugin.WithHidden(hidden)))
	flags.AddFlag(flagsplugin.NewBoolFlag(flagsplugin.Prefix("deleted", prefix), "", flagsplugin.WithHidden(hidden)))
}

// SetFromFlags sets the SearchApplicationsRequest message from flags.
func (m *SearchApplicationsRequest) SetFromFlags(flags *pflag.FlagSet, prefix string) (paths []string, err error) {
	if val, changed, err := flagsplugin.GetString(flags, flagsplugin.Prefix("query", prefix)); err != nil {
		return nil, err
	} else if changed {
		m.Query = val
		paths = append(paths, flagsplugin.Prefix("query", prefix))
	}
	if val, changed, err := flagsplugin.GetString(flags, flagsplugin.Prefix("id_contains", prefix)); err != nil {
		return nil, err
	} else if changed {
		m.IdContains = val
		paths = append(paths, flagsplugin.Prefix("id_contains", prefix))
	}
	if val, changed, err := flagsplugin.GetString(flags, flagsplugin.Prefix("name_contains", prefix)); err != nil {
		return nil, err
	} else if changed {
		m.NameContains = val
		paths = append(paths, flagsplugin.Prefix("name_contains", prefix))
	}
	if val, changed, err := flagsplugin.GetString(flags, flagsplugin.Prefix("description_contains", prefix)); err != nil {
		return nil, err
	} else if changed {
		m.DescriptionContains = val
		paths = append(paths, flagsplugin.Prefix("description_contains", prefix))
	}
	if val, changed, err := flagsplugin.GetStringStringMap(flags, flagsplugin.Prefix("attributes_contain", prefix)); err != nil {
		return nil, err
	} else if changed {
		m.AttributesContain = val
		paths = append(paths, flagsplugin.Prefix("attributes_contain", prefix))
	}
	if val, changed, err := flagsplugin.GetStringSlice(flags, flagsplugin.Prefix("field_mask", prefix)); err != nil {
		return nil, err
	} else if changed {
		m.FieldMask = gogo.SetFieldMask(val)
		paths = append(paths, flagsplugin.Prefix("field_mask", prefix))
	}
	if val, changed, err := flagsplugin.GetString(flags, flagsplugin.Prefix("order", prefix)); err != nil {
		return nil, err
	} else if changed {
		m.Order = val
		paths = append(paths, flagsplugin.Prefix("order", prefix))
	}
	if val, changed, err := flagsplugin.GetUint32(flags, flagsplugin.Prefix("limit", prefix)); err != nil {
		return nil, err
	} else if changed {
		m.Limit = val
		paths = append(paths, flagsplugin.Prefix("limit", prefix))
	}
	if val, changed, err := flagsplugin.GetUint32(flags, flagsplugin.Prefix("page", prefix)); err != nil {
		return nil, err
	} else if changed {
		m.Page = val
		paths = append(paths, flagsplugin.Prefix("page", prefix))
	}
	if val, changed, err := flagsplugin.GetBool(flags, flagsplugin.Prefix("deleted", prefix)); err != nil {
		return nil, err
	} else if changed {
		m.Deleted = val
		paths = append(paths, flagsplugin.Prefix("deleted", prefix))
	}
	return paths, nil
}

// AddSetFlagsForSearchClientsRequest adds flags to select fields in SearchClientsRequest.
func AddSetFlagsForSearchClientsRequest(flags *pflag.FlagSet, prefix string, hidden bool) {
	flags.AddFlag(flagsplugin.NewStringFlag(flagsplugin.Prefix("query", prefix), "", flagsplugin.WithHidden(hidden)))
	flags.AddFlag(flagsplugin.NewStringFlag(flagsplugin.Prefix("id-contains", prefix), "", flagsplugin.WithHidden(hidden)))
	flags.AddFlag(flagsplugin.NewStringFlag(flagsplugin.Prefix("name-contains", prefix), "", flagsplugin.WithHidden(hidden)))
	flags.AddFlag(flagsplugin.NewStringFlag(flagsplugin.Prefix("description-contains", prefix), "", flagsplugin.WithHidden(hidden)))
	flags.AddFlag(flagsplugin.NewStringStringMapFlag(flagsplugin.Prefix("attributes-contain", prefix), "", flagsplugin.WithHidden(hidden)))
	flags.AddFlag(flagsplugin.NewStringSliceFlag(flagsplugin.Prefix("state", prefix), flagsplugin.EnumValueDesc(State_value), flagsplugin.WithHidden(hidden)))
	flags.AddFlag(flagsplugin.NewStringSliceFlag(flagsplugin.Prefix("field-mask", prefix), "", flagsplugin.WithHidden(hidden)))
	flags.AddFlag(flagsplugin.NewStringFlag(flagsplugin.Prefix("order", prefix), "", flagsplugin.WithHidden(hidden)))
	flags.AddFlag(flagsplugin.NewUint32Flag(flagsplugin.Prefix("limit", prefix), "", flagsplugin.WithHidden(hidden)))
	flags.AddFlag(flagsplugin.NewUint32Flag(flagsplugin.Prefix("page", prefix), "", flagsplugin.WithHidden(hidden)))
	flags.AddFlag(flagsplugin.NewBoolFlag(flagsplugin.Prefix("deleted", prefix), "", flagsplugin.WithHidden(hidden)))
}

// SetFromFlags sets the SearchClientsRequest message from flags.
func (m *SearchClientsRequest) SetFromFlags(flags *pflag.FlagSet, prefix string) (paths []string, err error) {
	if val, changed, err := flagsplugin.GetString(flags, flagsplugin.Prefix("query", prefix)); err != nil {
		return nil, err
	} else if changed {
		m.Query = val
		paths = append(paths, flagsplugin.Prefix("query", prefix))
	}
	if val, changed, err := flagsplugin.GetString(flags, flagsplugin.Prefix("id_contains", prefix)); err != nil {
		return nil, err
	} else if changed {
		m.IdContains = val
		paths = append(paths, flagsplugin.Prefix("id_contains", prefix))
	}
	if val, changed, err := flagsplugin.GetString(flags, flagsplugin.Prefix("name_contains", prefix)); err != nil {
		return nil, err
	} else if changed {
		m.NameContains = val
		paths = append(paths, flagsplugin.Prefix("name_contains", prefix))
	}
	if val, changed, err := flagsplugin.GetString(flags, flagsplugin.Prefix("description_contains", prefix)); err != nil {
		return nil, err
	} else if changed {
		m.DescriptionContains = val
		paths = append(paths, flagsplugin.Prefix("description_contains", prefix))
	}
	if val, changed, err := flagsplugin.GetStringStringMap(flags, flagsplugin.Prefix("attributes_contain", prefix)); err != nil {
		return nil, err
	} else if changed {
		m.AttributesContain = val
		paths = append(paths, flagsplugin.Prefix("attributes_contain", prefix))
	}
	if val, changed, err := flagsplugin.GetStringSlice(flags, flagsplugin.Prefix("state", prefix)); err != nil {
		return nil, err
	} else if changed {
		m.State = make([]State, len(val))
		for i, v := range val {
			enumValue, err := flagsplugin.SetEnumString(v, State_value)
			if err != nil {
				return nil, err
			}
			m.State[i] = State(enumValue)
		}
		paths = append(paths, flagsplugin.Prefix("state", prefix))
	}
	if val, changed, err := flagsplugin.GetStringSlice(flags, flagsplugin.Prefix("field_mask", prefix)); err != nil {
		return nil, err
	} else if changed {
		m.FieldMask = gogo.SetFieldMask(val)
		paths = append(paths, flagsplugin.Prefix("field_mask", prefix))
	}
	if val, changed, err := flagsplugin.GetString(flags, flagsplugin.Prefix("order", prefix)); err != nil {
		return nil, err
	} else if changed {
		m.Order = val
		paths = append(paths, flagsplugin.Prefix("order", prefix))
	}
	if val, changed, err := flagsplugin.GetUint32(flags, flagsplugin.Prefix("limit", prefix)); err != nil {
		return nil, err
	} else if changed {
		m.Limit = val
		paths = append(paths, flagsplugin.Prefix("limit", prefix))
	}
	if val, changed, err := flagsplugin.GetUint32(flags, flagsplugin.Prefix("page", prefix)); err != nil {
		return nil, err
	} else if changed {
		m.Page = val
		paths = append(paths, flagsplugin.Prefix("page", prefix))
	}
	if val, changed, err := flagsplugin.GetBool(flags, flagsplugin.Prefix("deleted", prefix)); err != nil {
		return nil, err
	} else if changed {
		m.Deleted = val
		paths = append(paths, flagsplugin.Prefix("deleted", prefix))
	}
	return paths, nil
}

// AddSetFlagsForSearchGatewaysRequest adds flags to select fields in SearchGatewaysRequest.
func AddSetFlagsForSearchGatewaysRequest(flags *pflag.FlagSet, prefix string, hidden bool) {
	flags.AddFlag(flagsplugin.NewStringFlag(flagsplugin.Prefix("query", prefix), "", flagsplugin.WithHidden(hidden)))
	flags.AddFlag(flagsplugin.NewStringFlag(flagsplugin.Prefix("id-contains", prefix), "", flagsplugin.WithHidden(hidden)))
	flags.AddFlag(flagsplugin.NewStringFlag(flagsplugin.Prefix("name-contains", prefix), "", flagsplugin.WithHidden(hidden)))
	flags.AddFlag(flagsplugin.NewStringFlag(flagsplugin.Prefix("description-contains", prefix), "", flagsplugin.WithHidden(hidden)))
	flags.AddFlag(flagsplugin.NewStringStringMapFlag(flagsplugin.Prefix("attributes-contain", prefix), "", flagsplugin.WithHidden(hidden)))
	flags.AddFlag(flagsplugin.NewStringFlag(flagsplugin.Prefix("eui-contains", prefix), "", flagsplugin.WithHidden(hidden)))
	flags.AddFlag(flagsplugin.NewStringSliceFlag(flagsplugin.Prefix("field-mask", prefix), "", flagsplugin.WithHidden(hidden)))
	flags.AddFlag(flagsplugin.NewStringFlag(flagsplugin.Prefix("order", prefix), "", flagsplugin.WithHidden(hidden)))
	flags.AddFlag(flagsplugin.NewUint32Flag(flagsplugin.Prefix("limit", prefix), "", flagsplugin.WithHidden(hidden)))
	flags.AddFlag(flagsplugin.NewUint32Flag(flagsplugin.Prefix("page", prefix), "", flagsplugin.WithHidden(hidden)))
	flags.AddFlag(flagsplugin.NewBoolFlag(flagsplugin.Prefix("deleted", prefix), "", flagsplugin.WithHidden(hidden)))
}

// SetFromFlags sets the SearchGatewaysRequest message from flags.
func (m *SearchGatewaysRequest) SetFromFlags(flags *pflag.FlagSet, prefix string) (paths []string, err error) {
	if val, changed, err := flagsplugin.GetString(flags, flagsplugin.Prefix("query", prefix)); err != nil {
		return nil, err
	} else if changed {
		m.Query = val
		paths = append(paths, flagsplugin.Prefix("query", prefix))
	}
	if val, changed, err := flagsplugin.GetString(flags, flagsplugin.Prefix("id_contains", prefix)); err != nil {
		return nil, err
	} else if changed {
		m.IdContains = val
		paths = append(paths, flagsplugin.Prefix("id_contains", prefix))
	}
	if val, changed, err := flagsplugin.GetString(flags, flagsplugin.Prefix("name_contains", prefix)); err != nil {
		return nil, err
	} else if changed {
		m.NameContains = val
		paths = append(paths, flagsplugin.Prefix("name_contains", prefix))
	}
	if val, changed, err := flagsplugin.GetString(flags, flagsplugin.Prefix("description_contains", prefix)); err != nil {
		return nil, err
	} else if changed {
		m.DescriptionContains = val
		paths = append(paths, flagsplugin.Prefix("description_contains", prefix))
	}
	if val, changed, err := flagsplugin.GetStringStringMap(flags, flagsplugin.Prefix("attributes_contain", prefix)); err != nil {
		return nil, err
	} else if changed {
		m.AttributesContain = val
		paths = append(paths, flagsplugin.Prefix("attributes_contain", prefix))
	}
	if val, changed, err := flagsplugin.GetString(flags, flagsplugin.Prefix("eui_contains", prefix)); err != nil {
		return nil, err
	} else if changed {
		m.EuiContains = val
		paths = append(paths, flagsplugin.Prefix("eui_contains", prefix))
	}
	if val, changed, err := flagsplugin.GetStringSlice(flags, flagsplugin.Prefix("field_mask", prefix)); err != nil {
		return nil, err
	} else if changed {
		m.FieldMask = gogo.SetFieldMask(val)
		paths = append(paths, flagsplugin.Prefix("field_mask", prefix))
	}
	if val, changed, err := flagsplugin.GetString(flags, flagsplugin.Prefix("order", prefix)); err != nil {
		return nil, err
	} else if changed {
		m.Order = val
		paths = append(paths, flagsplugin.Prefix("order", prefix))
	}
	if val, changed, err := flagsplugin.GetUint32(flags, flagsplugin.Prefix("limit", prefix)); err != nil {
		return nil, err
	} else if changed {
		m.Limit = val
		paths = append(paths, flagsplugin.Prefix("limit", prefix))
	}
	if val, changed, err := flagsplugin.GetUint32(flags, flagsplugin.Prefix("page", prefix)); err != nil {
		return nil, err
	} else if changed {
		m.Page = val
		paths = append(paths, flagsplugin.Prefix("page", prefix))
	}
	if val, changed, err := flagsplugin.GetBool(flags, flagsplugin.Prefix("deleted", prefix)); err != nil {
		return nil, err
	} else if changed {
		m.Deleted = val
		paths = append(paths, flagsplugin.Prefix("deleted", prefix))
	}
	return paths, nil
}

// AddSetFlagsForSearchOrganizationsRequest adds flags to select fields in SearchOrganizationsRequest.
func AddSetFlagsForSearchOrganizationsRequest(flags *pflag.FlagSet, prefix string, hidden bool) {
	flags.AddFlag(flagsplugin.NewStringFlag(flagsplugin.Prefix("query", prefix), "", flagsplugin.WithHidden(hidden)))
	flags.AddFlag(flagsplugin.NewStringFlag(flagsplugin.Prefix("id-contains", prefix), "", flagsplugin.WithHidden(hidden)))
	flags.AddFlag(flagsplugin.NewStringFlag(flagsplugin.Prefix("name-contains", prefix), "", flagsplugin.WithHidden(hidden)))
	flags.AddFlag(flagsplugin.NewStringFlag(flagsplugin.Prefix("description-contains", prefix), "", flagsplugin.WithHidden(hidden)))
	flags.AddFlag(flagsplugin.NewStringStringMapFlag(flagsplugin.Prefix("attributes-contain", prefix), "", flagsplugin.WithHidden(hidden)))
	flags.AddFlag(flagsplugin.NewStringSliceFlag(flagsplugin.Prefix("field-mask", prefix), "", flagsplugin.WithHidden(hidden)))
	flags.AddFlag(flagsplugin.NewStringFlag(flagsplugin.Prefix("order", prefix), "", flagsplugin.WithHidden(hidden)))
	flags.AddFlag(flagsplugin.NewUint32Flag(flagsplugin.Prefix("limit", prefix), "", flagsplugin.WithHidden(hidden)))
	flags.AddFlag(flagsplugin.NewUint32Flag(flagsplugin.Prefix("page", prefix), "", flagsplugin.WithHidden(hidden)))
	flags.AddFlag(flagsplugin.NewBoolFlag(flagsplugin.Prefix("deleted", prefix), "", flagsplugin.WithHidden(hidden)))
}

// SetFromFlags sets the SearchOrganizationsRequest message from flags.
func (m *SearchOrganizationsRequest) SetFromFlags(flags *pflag.FlagSet, prefix string) (paths []string, err error) {
	if val, changed, err := flagsplugin.GetString(flags, flagsplugin.Prefix("query", prefix)); err != nil {
		return nil, err
	} else if changed {
		m.Query = val
		paths = append(paths, flagsplugin.Prefix("query", prefix))
	}
	if val, changed, err := flagsplugin.GetString(flags, flagsplugin.Prefix("id_contains", prefix)); err != nil {
		return nil, err
	} else if changed {
		m.IdContains = val
		paths = append(paths, flagsplugin.Prefix("id_contains", prefix))
	}
	if val, changed, err := flagsplugin.GetString(flags, flagsplugin.Prefix("name_contains", prefix)); err != nil {
		return nil, err
	} else if changed {
		m.NameContains = val
		paths = append(paths, flagsplugin.Prefix("name_contains", prefix))
	}
	if val, changed, err := flagsplugin.GetString(flags, flagsplugin.Prefix("description_contains", prefix)); err != nil {
		return nil, err
	} else if changed {
		m.DescriptionContains = val
		paths = append(paths, flagsplugin.Prefix("description_contains", prefix))
	}
	if val, changed, err := flagsplugin.GetStringStringMap(flags, flagsplugin.Prefix("attributes_contain", prefix)); err != nil {
		return nil, err
	} else if changed {
		m.AttributesContain = val
		paths = append(paths, flagsplugin.Prefix("attributes_contain", prefix))
	}
	if val, changed, err := flagsplugin.GetStringSlice(flags, flagsplugin.Prefix("field_mask", prefix)); err != nil {
		return nil, err
	} else if changed {
		m.FieldMask = gogo.SetFieldMask(val)
		paths = append(paths, flagsplugin.Prefix("field_mask", prefix))
	}
	if val, changed, err := flagsplugin.GetString(flags, flagsplugin.Prefix("order", prefix)); err != nil {
		return nil, err
	} else if changed {
		m.Order = val
		paths = append(paths, flagsplugin.Prefix("order", prefix))
	}
	if val, changed, err := flagsplugin.GetUint32(flags, flagsplugin.Prefix("limit", prefix)); err != nil {
		return nil, err
	} else if changed {
		m.Limit = val
		paths = append(paths, flagsplugin.Prefix("limit", prefix))
	}
	if val, changed, err := flagsplugin.GetUint32(flags, flagsplugin.Prefix("page", prefix)); err != nil {
		return nil, err
	} else if changed {
		m.Page = val
		paths = append(paths, flagsplugin.Prefix("page", prefix))
	}
	if val, changed, err := flagsplugin.GetBool(flags, flagsplugin.Prefix("deleted", prefix)); err != nil {
		return nil, err
	} else if changed {
		m.Deleted = val
		paths = append(paths, flagsplugin.Prefix("deleted", prefix))
	}
	return paths, nil
}

// AddSetFlagsForSearchUsersRequest adds flags to select fields in SearchUsersRequest.
func AddSetFlagsForSearchUsersRequest(flags *pflag.FlagSet, prefix string, hidden bool) {
	flags.AddFlag(flagsplugin.NewStringFlag(flagsplugin.Prefix("query", prefix), "", flagsplugin.WithHidden(hidden)))
	flags.AddFlag(flagsplugin.NewStringFlag(flagsplugin.Prefix("id-contains", prefix), "", flagsplugin.WithHidden(hidden)))
	flags.AddFlag(flagsplugin.NewStringFlag(flagsplugin.Prefix("name-contains", prefix), "", flagsplugin.WithHidden(hidden)))
	flags.AddFlag(flagsplugin.NewStringFlag(flagsplugin.Prefix("description-contains", prefix), "", flagsplugin.WithHidden(hidden)))
	flags.AddFlag(flagsplugin.NewStringStringMapFlag(flagsplugin.Prefix("attributes-contain", prefix), "", flagsplugin.WithHidden(hidden)))
	flags.AddFlag(flagsplugin.NewStringSliceFlag(flagsplugin.Prefix("state", prefix), flagsplugin.EnumValueDesc(State_value), flagsplugin.WithHidden(hidden)))
	flags.AddFlag(flagsplugin.NewStringSliceFlag(flagsplugin.Prefix("field-mask", prefix), "", flagsplugin.WithHidden(hidden)))
	flags.AddFlag(flagsplugin.NewStringFlag(flagsplugin.Prefix("order", prefix), "", flagsplugin.WithHidden(hidden)))
	flags.AddFlag(flagsplugin.NewUint32Flag(flagsplugin.Prefix("limit", prefix), "", flagsplugin.WithHidden(hidden)))
	flags.AddFlag(flagsplugin.NewUint32Flag(flagsplugin.Prefix("page", prefix), "", flagsplugin.WithHidden(hidden)))
	flags.AddFlag(flagsplugin.NewBoolFlag(flagsplugin.Prefix("deleted", prefix), "", flagsplugin.WithHidden(hidden)))
}

// SetFromFlags sets the SearchUsersRequest message from flags.
func (m *SearchUsersRequest) SetFromFlags(flags *pflag.FlagSet, prefix string) (paths []string, err error) {
	if val, changed, err := flagsplugin.GetString(flags, flagsplugin.Prefix("query", prefix)); err != nil {
		return nil, err
	} else if changed {
		m.Query = val
		paths = append(paths, flagsplugin.Prefix("query", prefix))
	}
	if val, changed, err := flagsplugin.GetString(flags, flagsplugin.Prefix("id_contains", prefix)); err != nil {
		return nil, err
	} else if changed {
		m.IdContains = val
		paths = append(paths, flagsplugin.Prefix("id_contains", prefix))
	}
	if val, changed, err := flagsplugin.GetString(flags, flagsplugin.Prefix("name_contains", prefix)); err != nil {
		return nil, err
	} else if changed {
		m.NameContains = val
		paths = append(paths, flagsplugin.Prefix("name_contains", prefix))
	}
	if val, changed, err := flagsplugin.GetString(flags, flagsplugin.Prefix("description_contains", prefix)); err != nil {
		return nil, err
	} else if changed {
		m.DescriptionContains = val
		paths = append(paths, flagsplugin.Prefix("description_contains", prefix))
	}
	if val, changed, err := flagsplugin.GetStringStringMap(flags, flagsplugin.Prefix("attributes_contain", prefix)); err != nil {
		return nil, err
	} else if changed {
		m.AttributesContain = val
		paths = append(paths, flagsplugin.Prefix("attributes_contain", prefix))
	}
	if val, changed, err := flagsplugin.GetStringSlice(flags, flagsplugin.Prefix("state", prefix)); err != nil {
		return nil, err
	} else if changed {
		m.State = make([]State, len(val))
		for i, v := range val {
			enumValue, err := flagsplugin.SetEnumString(v, State_value)
			if err != nil {
				return nil, err
			}
			m.State[i] = State(enumValue)
		}
		paths = append(paths, flagsplugin.Prefix("state", prefix))
	}
	if val, changed, err := flagsplugin.GetStringSlice(flags, flagsplugin.Prefix("field_mask", prefix)); err != nil {
		return nil, err
	} else if changed {
		m.FieldMask = gogo.SetFieldMask(val)
		paths = append(paths, flagsplugin.Prefix("field_mask", prefix))
	}
	if val, changed, err := flagsplugin.GetString(flags, flagsplugin.Prefix("order", prefix)); err != nil {
		return nil, err
	} else if changed {
		m.Order = val
		paths = append(paths, flagsplugin.Prefix("order", prefix))
	}
	if val, changed, err := flagsplugin.GetUint32(flags, flagsplugin.Prefix("limit", prefix)); err != nil {
		return nil, err
	} else if changed {
		m.Limit = val
		paths = append(paths, flagsplugin.Prefix("limit", prefix))
	}
	if val, changed, err := flagsplugin.GetUint32(flags, flagsplugin.Prefix("page", prefix)); err != nil {
		return nil, err
	} else if changed {
		m.Page = val
		paths = append(paths, flagsplugin.Prefix("page", prefix))
	}
	if val, changed, err := flagsplugin.GetBool(flags, flagsplugin.Prefix("deleted", prefix)); err != nil {
		return nil, err
	} else if changed {
		m.Deleted = val
		paths = append(paths, flagsplugin.Prefix("deleted", prefix))
	}
	return paths, nil
}

// AddSetFlagsForSearchAccountsRequest adds flags to select fields in SearchAccountsRequest.
func AddSetFlagsForSearchAccountsRequest(flags *pflag.FlagSet, prefix string, hidden bool) {
	flags.AddFlag(flagsplugin.NewStringFlag(flagsplugin.Prefix("query", prefix), "", flagsplugin.WithHidden(hidden)))
	flags.AddFlag(flagsplugin.NewBoolFlag(flagsplugin.Prefix("only-users", prefix), "", flagsplugin.WithHidden(hidden)))
	AddSetFlagsForApplicationIdentifiers(flags, flagsplugin.Prefix("collaborator-of.application-ids", prefix), hidden)
	AddSetFlagsForClientIdentifiers(flags, flagsplugin.Prefix("collaborator-of.client-ids", prefix), hidden)
	AddSetFlagsForGatewayIdentifiers(flags, flagsplugin.Prefix("collaborator-of.gateway-ids", prefix), hidden)
	AddSetFlagsForOrganizationIdentifiers(flags, flagsplugin.Prefix("collaborator-of.organization-ids", prefix), hidden)
}

// SetFromFlags sets the SearchAccountsRequest message from flags.
func (m *SearchAccountsRequest) SetFromFlags(flags *pflag.FlagSet, prefix string) (paths []string, err error) {
	if val, changed, err := flagsplugin.GetString(flags, flagsplugin.Prefix("query", prefix)); err != nil {
		return nil, err
	} else if changed {
		m.Query = val
		paths = append(paths, flagsplugin.Prefix("query", prefix))
	}
	if val, changed, err := flagsplugin.GetBool(flags, flagsplugin.Prefix("only_users", prefix)); err != nil {
		return nil, err
	} else if changed {
		m.OnlyUsers = val
		paths = append(paths, flagsplugin.Prefix("only_users", prefix))
	}
	if changed := flagsplugin.IsAnyPrefixSet(flags, flagsplugin.Prefix("collaborator_of.application_ids", prefix)); changed {
		ov := &SearchAccountsRequest_ApplicationIds{}
		if ov.ApplicationIds == nil {
			ov.ApplicationIds = &ApplicationIdentifiers{}
		}
		if setPaths, err := ov.ApplicationIds.SetFromFlags(flags, flagsplugin.Prefix("collaborator_of.application_ids", prefix)); err != nil {
			return nil, err
		} else {
			paths = append(paths, setPaths...)
		}
		m.CollaboratorOf = ov
	}
	if changed := flagsplugin.IsAnyPrefixSet(flags, flagsplugin.Prefix("collaborator_of.client_ids", prefix)); changed {
		ov := &SearchAccountsRequest_ClientIds{}
		if ov.ClientIds == nil {
			ov.ClientIds = &ClientIdentifiers{}
		}
		if setPaths, err := ov.ClientIds.SetFromFlags(flags, flagsplugin.Prefix("collaborator_of.client_ids", prefix)); err != nil {
			return nil, err
		} else {
			paths = append(paths, setPaths...)
		}
		m.CollaboratorOf = ov
	}
	if changed := flagsplugin.IsAnyPrefixSet(flags, flagsplugin.Prefix("collaborator_of.gateway_ids", prefix)); changed {
		ov := &SearchAccountsRequest_GatewayIds{}
		if ov.GatewayIds == nil {
			ov.GatewayIds = &GatewayIdentifiers{}
		}
		if setPaths, err := ov.GatewayIds.SetFromFlags(flags, flagsplugin.Prefix("collaborator_of.gateway_ids", prefix)); err != nil {
			return nil, err
		} else {
			paths = append(paths, setPaths...)
		}
		m.CollaboratorOf = ov
	}
	if changed := flagsplugin.IsAnyPrefixSet(flags, flagsplugin.Prefix("collaborator_of.organization_ids", prefix)); changed {
		ov := &SearchAccountsRequest_OrganizationIds{}
		if ov.OrganizationIds == nil {
			ov.OrganizationIds = &OrganizationIdentifiers{}
		}
		if setPaths, err := ov.OrganizationIds.SetFromFlags(flags, flagsplugin.Prefix("collaborator_of.organization_ids", prefix)); err != nil {
			return nil, err
		} else {
			paths = append(paths, setPaths...)
		}
		m.CollaboratorOf = ov
	}
	return paths, nil
}

// AddSetFlagsForSearchEndDevicesRequest adds flags to select fields in SearchEndDevicesRequest.
func AddSetFlagsForSearchEndDevicesRequest(flags *pflag.FlagSet, prefix string, hidden bool) {
	AddSetFlagsForApplicationIdentifiers(flags, flagsplugin.Prefix("application-ids", prefix), hidden)
	flags.AddFlag(flagsplugin.NewStringFlag(flagsplugin.Prefix("query", prefix), "", flagsplugin.WithHidden(hidden)))
	flags.AddFlag(flagsplugin.NewStringFlag(flagsplugin.Prefix("id-contains", prefix), "", flagsplugin.WithHidden(hidden)))
	flags.AddFlag(flagsplugin.NewStringFlag(flagsplugin.Prefix("name-contains", prefix), "", flagsplugin.WithHidden(hidden)))
	flags.AddFlag(flagsplugin.NewStringFlag(flagsplugin.Prefix("description-contains", prefix), "", flagsplugin.WithHidden(hidden)))
	flags.AddFlag(flagsplugin.NewStringStringMapFlag(flagsplugin.Prefix("attributes-contain", prefix), "", flagsplugin.WithHidden(hidden)))
	flags.AddFlag(flagsplugin.NewStringFlag(flagsplugin.Prefix("dev-eui-contains", prefix), "", flagsplugin.WithHidden(hidden)))
	flags.AddFlag(flagsplugin.NewStringFlag(flagsplugin.Prefix("join-eui-contains", prefix), "", flagsplugin.WithHidden(hidden)))
	flags.AddFlag(flagsplugin.NewStringFlag(flagsplugin.Prefix("dev-addr-contains", prefix), "", flagsplugin.WithHidden(hidden)))
	flags.AddFlag(flagsplugin.NewStringSliceFlag(flagsplugin.Prefix("field-mask", prefix), "", flagsplugin.WithHidden(hidden)))
	flags.AddFlag(flagsplugin.NewStringFlag(flagsplugin.Prefix("order", prefix), "", flagsplugin.WithHidden(hidden)))
	flags.AddFlag(flagsplugin.NewUint32Flag(flagsplugin.Prefix("limit", prefix), "", flagsplugin.WithHidden(hidden)))
	flags.AddFlag(flagsplugin.NewUint32Flag(flagsplugin.Prefix("page", prefix), "", flagsplugin.WithHidden(hidden)))
}

// SetFromFlags sets the SearchEndDevicesRequest message from flags.
func (m *SearchEndDevicesRequest) SetFromFlags(flags *pflag.FlagSet, prefix string) (paths []string, err error) {
	if changed := flagsplugin.IsAnyPrefixSet(flags, flagsplugin.Prefix("application_ids", prefix)); changed {
		if m.ApplicationIds == nil {
			m.ApplicationIds = &ApplicationIdentifiers{}
		}
		if setPaths, err := m.ApplicationIds.SetFromFlags(flags, flagsplugin.Prefix("application_ids", prefix)); err != nil {
			return nil, err
		} else {
			paths = append(paths, setPaths...)
		}
	}
	if val, changed, err := flagsplugin.GetString(flags, flagsplugin.Prefix("query", prefix)); err != nil {
		return nil, err
	} else if changed {
		m.Query = val
		paths = append(paths, flagsplugin.Prefix("query", prefix))
	}
	if val, changed, err := flagsplugin.GetString(flags, flagsplugin.Prefix("id_contains", prefix)); err != nil {
		return nil, err
	} else if changed {
		m.IdContains = val
		paths = append(paths, flagsplugin.Prefix("id_contains", prefix))
	}
	if val, changed, err := flagsplugin.GetString(flags, flagsplugin.Prefix("name_contains", prefix)); err != nil {
		return nil, err
	} else if changed {
		m.NameContains = val
		paths = append(paths, flagsplugin.Prefix("name_contains", prefix))
	}
	if val, changed, err := flagsplugin.GetString(flags, flagsplugin.Prefix("description_contains", prefix)); err != nil {
		return nil, err
	} else if changed {
		m.DescriptionContains = val
		paths = append(paths, flagsplugin.Prefix("description_contains", prefix))
	}
	if val, changed, err := flagsplugin.GetStringStringMap(flags, flagsplugin.Prefix("attributes_contain", prefix)); err != nil {
		return nil, err
	} else if changed {
		m.AttributesContain = val
		paths = append(paths, flagsplugin.Prefix("attributes_contain", prefix))
	}
	if val, changed, err := flagsplugin.GetString(flags, flagsplugin.Prefix("dev_eui_contains", prefix)); err != nil {
		return nil, err
	} else if changed {
		m.DevEuiContains = val
		paths = append(paths, flagsplugin.Prefix("dev_eui_contains", prefix))
	}
	if val, changed, err := flagsplugin.GetString(flags, flagsplugin.Prefix("join_eui_contains", prefix)); err != nil {
		return nil, err
	} else if changed {
		m.JoinEuiContains = val
		paths = append(paths, flagsplugin.Prefix("join_eui_contains", prefix))
	}
	if val, changed, err := flagsplugin.GetString(flags, flagsplugin.Prefix("dev_addr_contains", prefix)); err != nil {
		return nil, err
	} else if changed {
		m.DevAddrContains = val
		paths = append(paths, flagsplugin.Prefix("dev_addr_contains", prefix))
	}
	if val, changed, err := flagsplugin.GetStringSlice(flags, flagsplugin.Prefix("field_mask", prefix)); err != nil {
		return nil, err
	} else if changed {
		m.FieldMask = gogo.SetFieldMask(val)
		paths = append(paths, flagsplugin.Prefix("field_mask", prefix))
	}
	if val, changed, err := flagsplugin.GetString(flags, flagsplugin.Prefix("order", prefix)); err != nil {
		return nil, err
	} else if changed {
		m.Order = val
		paths = append(paths, flagsplugin.Prefix("order", prefix))
	}
	if val, changed, err := flagsplugin.GetUint32(flags, flagsplugin.Prefix("limit", prefix)); err != nil {
		return nil, err
	} else if changed {
		m.Limit = val
		paths = append(paths, flagsplugin.Prefix("limit", prefix))
	}
	if val, changed, err := flagsplugin.GetUint32(flags, flagsplugin.Prefix("page", prefix)); err != nil {
		return nil, err
	} else if changed {
		m.Page = val
		paths = append(paths, flagsplugin.Prefix("page", prefix))
	}
	return paths, nil
}
