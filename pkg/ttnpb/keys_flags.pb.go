// Code generated by protoc-gen-go-flags. DO NOT EDIT.
// versions:
// - protoc-gen-go-flags v1.0.6
// - protoc              v3.21.1
// source: lorawan-stack/api/keys.proto

package ttnpb

import (
	flagsplugin "github.com/TheThingsIndustries/protoc-gen-go-flags/flagsplugin"
	pflag "github.com/spf13/pflag"
	customflags "go.thethings.network/lorawan-stack/v3/cmd/ttn-lw-cli/customflags"
)

// AddSelectFlagsForKeyEnvelope adds flags to select fields in KeyEnvelope.
func AddSelectFlagsForKeyEnvelope(flags *pflag.FlagSet, prefix string, hidden bool) {
	flags.AddFlag(flagsplugin.NewBoolFlag(flagsplugin.Prefix("key", prefix), flagsplugin.SelectDesc(flagsplugin.Prefix("key", prefix), false), flagsplugin.WithHidden(hidden)))
	flags.AddFlag(flagsplugin.NewBoolFlag(flagsplugin.Prefix("kek-label", prefix), flagsplugin.SelectDesc(flagsplugin.Prefix("kek-label", prefix), false), flagsplugin.WithHidden(hidden)))
	flags.AddFlag(flagsplugin.NewBoolFlag(flagsplugin.Prefix("encrypted-key", prefix), flagsplugin.SelectDesc(flagsplugin.Prefix("encrypted-key", prefix), false), flagsplugin.WithHidden(hidden)))
}

// SelectFromFlags outputs the fieldmask paths forKeyEnvelope message from select flags.
func PathsFromSelectFlagsForKeyEnvelope(flags *pflag.FlagSet, prefix string) (paths []string, err error) {
	if val, selected, err := flagsplugin.GetBool(flags, flagsplugin.Prefix("key", prefix)); err != nil {
		return nil, err
	} else if selected && val {
		paths = append(paths, flagsplugin.Prefix("key", prefix))
	}
	if val, selected, err := flagsplugin.GetBool(flags, flagsplugin.Prefix("kek_label", prefix)); err != nil {
		return nil, err
	} else if selected && val {
		paths = append(paths, flagsplugin.Prefix("kek_label", prefix))
	}
	if val, selected, err := flagsplugin.GetBool(flags, flagsplugin.Prefix("encrypted_key", prefix)); err != nil {
		return nil, err
	} else if selected && val {
		paths = append(paths, flagsplugin.Prefix("encrypted_key", prefix))
	}
	return paths, nil
}

// AddSetFlagsForKeyEnvelope adds flags to select fields in KeyEnvelope.
func AddSetFlagsForKeyEnvelope(flags *pflag.FlagSet, prefix string, hidden bool) {
	flags.AddFlag(customflags.New16BytesFlag(flagsplugin.Prefix("key", prefix), "", flagsplugin.WithHidden(hidden)))
	flags.AddFlag(flagsplugin.NewStringFlag(flagsplugin.Prefix("kek-label", prefix), "", flagsplugin.WithHidden(hidden)))
	flags.AddFlag(flagsplugin.NewHexBytesFlag(flagsplugin.Prefix("encrypted-key", prefix), "", flagsplugin.WithHidden(hidden)))
}

// SetFromFlags sets the KeyEnvelope message from flags.
func (m *KeyEnvelope) SetFromFlags(flags *pflag.FlagSet, prefix string) (paths []string, err error) {
	if val, changed, err := customflags.GetExactBytes(flags, flagsplugin.Prefix("key", prefix)); err != nil {
		return nil, err
	} else if changed {
		m.Key = val
		paths = append(paths, flagsplugin.Prefix("key", prefix))
	}
	if val, changed, err := flagsplugin.GetString(flags, flagsplugin.Prefix("kek_label", prefix)); err != nil {
		return nil, err
	} else if changed {
		m.KekLabel = val
		paths = append(paths, flagsplugin.Prefix("kek_label", prefix))
	}
	if val, changed, err := flagsplugin.GetBytes(flags, flagsplugin.Prefix("encrypted_key", prefix)); err != nil {
		return nil, err
	} else if changed {
		m.EncryptedKey = val
		paths = append(paths, flagsplugin.Prefix("encrypted_key", prefix))
	}
	return paths, nil
}

// AddSelectFlagsForRootKeys adds flags to select fields in RootKeys.
func AddSelectFlagsForRootKeys(flags *pflag.FlagSet, prefix string, hidden bool) {
	flags.AddFlag(flagsplugin.NewBoolFlag(flagsplugin.Prefix("root-key-id", prefix), flagsplugin.SelectDesc(flagsplugin.Prefix("root-key-id", prefix), false), flagsplugin.WithHidden(hidden)))
	flags.AddFlag(flagsplugin.NewBoolFlag(flagsplugin.Prefix("app-key", prefix), flagsplugin.SelectDesc(flagsplugin.Prefix("app-key", prefix), true), flagsplugin.WithHidden(hidden)))
	AddSelectFlagsForKeyEnvelope(flags, flagsplugin.Prefix("app-key", prefix), hidden)
	flags.AddFlag(flagsplugin.NewBoolFlag(flagsplugin.Prefix("nwk-key", prefix), flagsplugin.SelectDesc(flagsplugin.Prefix("nwk-key", prefix), true), flagsplugin.WithHidden(hidden)))
	AddSelectFlagsForKeyEnvelope(flags, flagsplugin.Prefix("nwk-key", prefix), hidden)
}

// SelectFromFlags outputs the fieldmask paths forRootKeys message from select flags.
func PathsFromSelectFlagsForRootKeys(flags *pflag.FlagSet, prefix string) (paths []string, err error) {
	if val, selected, err := flagsplugin.GetBool(flags, flagsplugin.Prefix("root_key_id", prefix)); err != nil {
		return nil, err
	} else if selected && val {
		paths = append(paths, flagsplugin.Prefix("root_key_id", prefix))
	}
	if val, selected, err := flagsplugin.GetBool(flags, flagsplugin.Prefix("app_key", prefix)); err != nil {
		return nil, err
	} else if selected && val {
		paths = append(paths, flagsplugin.Prefix("app_key", prefix))
	}
	if selectPaths, err := PathsFromSelectFlagsForKeyEnvelope(flags, flagsplugin.Prefix("app_key", prefix)); err != nil {
		return nil, err
	} else {
		paths = append(paths, selectPaths...)
	}
	if val, selected, err := flagsplugin.GetBool(flags, flagsplugin.Prefix("nwk_key", prefix)); err != nil {
		return nil, err
	} else if selected && val {
		paths = append(paths, flagsplugin.Prefix("nwk_key", prefix))
	}
	if selectPaths, err := PathsFromSelectFlagsForKeyEnvelope(flags, flagsplugin.Prefix("nwk_key", prefix)); err != nil {
		return nil, err
	} else {
		paths = append(paths, selectPaths...)
	}
	return paths, nil
}

// AddSetFlagsForRootKeys adds flags to select fields in RootKeys.
func AddSetFlagsForRootKeys(flags *pflag.FlagSet, prefix string, hidden bool) {
	flags.AddFlag(flagsplugin.NewStringFlag(flagsplugin.Prefix("root-key-id", prefix), "", flagsplugin.WithHidden(hidden)))
	AddSetFlagsForKeyEnvelope(flags, flagsplugin.Prefix("app-key", prefix), hidden)
	AddSetFlagsForKeyEnvelope(flags, flagsplugin.Prefix("nwk-key", prefix), hidden)
}

// SetFromFlags sets the RootKeys message from flags.
func (m *RootKeys) SetFromFlags(flags *pflag.FlagSet, prefix string) (paths []string, err error) {
	if val, changed, err := flagsplugin.GetString(flags, flagsplugin.Prefix("root_key_id", prefix)); err != nil {
		return nil, err
	} else if changed {
		m.RootKeyId = val
		paths = append(paths, flagsplugin.Prefix("root_key_id", prefix))
	}
	if changed := flagsplugin.IsAnyPrefixSet(flags, flagsplugin.Prefix("app_key", prefix)); changed {
		if m.AppKey == nil {
			m.AppKey = &KeyEnvelope{}
		}
		if setPaths, err := m.AppKey.SetFromFlags(flags, flagsplugin.Prefix("app_key", prefix)); err != nil {
			return nil, err
		} else {
			paths = append(paths, setPaths...)
		}
	}
	if changed := flagsplugin.IsAnyPrefixSet(flags, flagsplugin.Prefix("nwk_key", prefix)); changed {
		if m.NwkKey == nil {
			m.NwkKey = &KeyEnvelope{}
		}
		if setPaths, err := m.NwkKey.SetFromFlags(flags, flagsplugin.Prefix("nwk_key", prefix)); err != nil {
			return nil, err
		} else {
			paths = append(paths, setPaths...)
		}
	}
	return paths, nil
}

// AddSelectFlagsForSessionKeys adds flags to select fields in SessionKeys.
func AddSelectFlagsForSessionKeys(flags *pflag.FlagSet, prefix string, hidden bool) {
	flags.AddFlag(flagsplugin.NewBoolFlag(flagsplugin.Prefix("session-key-id", prefix), flagsplugin.SelectDesc(flagsplugin.Prefix("session-key-id", prefix), false), flagsplugin.WithHidden(hidden)))
	flags.AddFlag(flagsplugin.NewBoolFlag(flagsplugin.Prefix("f-nwk-s-int-key", prefix), flagsplugin.SelectDesc(flagsplugin.Prefix("f-nwk-s-int-key", prefix), true), flagsplugin.WithHidden(hidden)))
	AddSelectFlagsForKeyEnvelope(flags, flagsplugin.Prefix("f-nwk-s-int-key", prefix), hidden)
	flags.AddFlag(flagsplugin.NewBoolFlag(flagsplugin.Prefix("s-nwk-s-int-key", prefix), flagsplugin.SelectDesc(flagsplugin.Prefix("s-nwk-s-int-key", prefix), true), flagsplugin.WithHidden(hidden)))
	AddSelectFlagsForKeyEnvelope(flags, flagsplugin.Prefix("s-nwk-s-int-key", prefix), hidden)
	flags.AddFlag(flagsplugin.NewBoolFlag(flagsplugin.Prefix("nwk-s-enc-key", prefix), flagsplugin.SelectDesc(flagsplugin.Prefix("nwk-s-enc-key", prefix), true), flagsplugin.WithHidden(hidden)))
	AddSelectFlagsForKeyEnvelope(flags, flagsplugin.Prefix("nwk-s-enc-key", prefix), hidden)
	flags.AddFlag(flagsplugin.NewBoolFlag(flagsplugin.Prefix("app-s-key", prefix), flagsplugin.SelectDesc(flagsplugin.Prefix("app-s-key", prefix), true), flagsplugin.WithHidden(hidden)))
	AddSelectFlagsForKeyEnvelope(flags, flagsplugin.Prefix("app-s-key", prefix), hidden)
}

// SelectFromFlags outputs the fieldmask paths forSessionKeys message from select flags.
func PathsFromSelectFlagsForSessionKeys(flags *pflag.FlagSet, prefix string) (paths []string, err error) {
	if val, selected, err := flagsplugin.GetBool(flags, flagsplugin.Prefix("session_key_id", prefix)); err != nil {
		return nil, err
	} else if selected && val {
		paths = append(paths, flagsplugin.Prefix("session_key_id", prefix))
	}
	if val, selected, err := flagsplugin.GetBool(flags, flagsplugin.Prefix("f_nwk_s_int_key", prefix)); err != nil {
		return nil, err
	} else if selected && val {
		paths = append(paths, flagsplugin.Prefix("f_nwk_s_int_key", prefix))
	}
	if selectPaths, err := PathsFromSelectFlagsForKeyEnvelope(flags, flagsplugin.Prefix("f_nwk_s_int_key", prefix)); err != nil {
		return nil, err
	} else {
		paths = append(paths, selectPaths...)
	}
	if val, selected, err := flagsplugin.GetBool(flags, flagsplugin.Prefix("s_nwk_s_int_key", prefix)); err != nil {
		return nil, err
	} else if selected && val {
		paths = append(paths, flagsplugin.Prefix("s_nwk_s_int_key", prefix))
	}
	if selectPaths, err := PathsFromSelectFlagsForKeyEnvelope(flags, flagsplugin.Prefix("s_nwk_s_int_key", prefix)); err != nil {
		return nil, err
	} else {
		paths = append(paths, selectPaths...)
	}
	if val, selected, err := flagsplugin.GetBool(flags, flagsplugin.Prefix("nwk_s_enc_key", prefix)); err != nil {
		return nil, err
	} else if selected && val {
		paths = append(paths, flagsplugin.Prefix("nwk_s_enc_key", prefix))
	}
	if selectPaths, err := PathsFromSelectFlagsForKeyEnvelope(flags, flagsplugin.Prefix("nwk_s_enc_key", prefix)); err != nil {
		return nil, err
	} else {
		paths = append(paths, selectPaths...)
	}
	if val, selected, err := flagsplugin.GetBool(flags, flagsplugin.Prefix("app_s_key", prefix)); err != nil {
		return nil, err
	} else if selected && val {
		paths = append(paths, flagsplugin.Prefix("app_s_key", prefix))
	}
	if selectPaths, err := PathsFromSelectFlagsForKeyEnvelope(flags, flagsplugin.Prefix("app_s_key", prefix)); err != nil {
		return nil, err
	} else {
		paths = append(paths, selectPaths...)
	}
	return paths, nil
}

// AddSetFlagsForSessionKeys adds flags to select fields in SessionKeys.
func AddSetFlagsForSessionKeys(flags *pflag.FlagSet, prefix string, hidden bool) {
	flags.AddFlag(flagsplugin.NewBytesFlag(flagsplugin.Prefix("session-key-id", prefix), "", flagsplugin.WithHidden(hidden)))
	AddSetFlagsForKeyEnvelope(flags, flagsplugin.Prefix("f-nwk-s-int-key", prefix), hidden)
	AddSetFlagsForKeyEnvelope(flags, flagsplugin.Prefix("s-nwk-s-int-key", prefix), hidden)
	AddSetFlagsForKeyEnvelope(flags, flagsplugin.Prefix("nwk-s-enc-key", prefix), hidden)
	AddSetFlagsForKeyEnvelope(flags, flagsplugin.Prefix("app-s-key", prefix), hidden)
}

// SetFromFlags sets the SessionKeys message from flags.
func (m *SessionKeys) SetFromFlags(flags *pflag.FlagSet, prefix string) (paths []string, err error) {
	if val, changed, err := flagsplugin.GetBytes(flags, flagsplugin.Prefix("session_key_id", prefix)); err != nil {
		return nil, err
	} else if changed {
		m.SessionKeyId = val
		paths = append(paths, flagsplugin.Prefix("session_key_id", prefix))
	}
	if changed := flagsplugin.IsAnyPrefixSet(flags, flagsplugin.Prefix("f_nwk_s_int_key", prefix)); changed {
		if m.FNwkSIntKey == nil {
			m.FNwkSIntKey = &KeyEnvelope{}
		}
		if setPaths, err := m.FNwkSIntKey.SetFromFlags(flags, flagsplugin.Prefix("f_nwk_s_int_key", prefix)); err != nil {
			return nil, err
		} else {
			paths = append(paths, setPaths...)
		}
	}
	if changed := flagsplugin.IsAnyPrefixSet(flags, flagsplugin.Prefix("s_nwk_s_int_key", prefix)); changed {
		if m.SNwkSIntKey == nil {
			m.SNwkSIntKey = &KeyEnvelope{}
		}
		if setPaths, err := m.SNwkSIntKey.SetFromFlags(flags, flagsplugin.Prefix("s_nwk_s_int_key", prefix)); err != nil {
			return nil, err
		} else {
			paths = append(paths, setPaths...)
		}
	}
	if changed := flagsplugin.IsAnyPrefixSet(flags, flagsplugin.Prefix("nwk_s_enc_key", prefix)); changed {
		if m.NwkSEncKey == nil {
			m.NwkSEncKey = &KeyEnvelope{}
		}
		if setPaths, err := m.NwkSEncKey.SetFromFlags(flags, flagsplugin.Prefix("nwk_s_enc_key", prefix)); err != nil {
			return nil, err
		} else {
			paths = append(paths, setPaths...)
		}
	}
	if changed := flagsplugin.IsAnyPrefixSet(flags, flagsplugin.Prefix("app_s_key", prefix)); changed {
		if m.AppSKey == nil {
			m.AppSKey = &KeyEnvelope{}
		}
		if setPaths, err := m.AppSKey.SetFromFlags(flags, flagsplugin.Prefix("app_s_key", prefix)); err != nil {
			return nil, err
		} else {
			paths = append(paths, setPaths...)
		}
	}
	return paths, nil
}
