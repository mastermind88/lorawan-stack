// Code generated by protoc-gen-fieldmask. DO NOT EDIT.

package ttnpb

import fmt "fmt"

var KeyEnvelopeFieldPathsNested = []string{
	"kek_label",
	"key",
}

var KeyEnvelopeFieldPathsTopLevel = []string{
	"kek_label",
	"key",
}

func (dst *KeyEnvelope) SetFields(src *KeyEnvelope, paths ...string) error {
	for name, subs := range _processPaths(append(paths[:0:0], paths...)) {
		switch name {
		case "key":
			if len(subs) > 0 {
				return fmt.Errorf("'key' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.Key = src.Key
			} else {
				var zero []byte
				dst.Key = zero
			}
		case "kek_label":
			if len(subs) > 0 {
				return fmt.Errorf("'kek_label' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.KEKLabel = src.KEKLabel
			} else {
				var zero string
				dst.KEKLabel = zero
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

var RootKeysFieldPathsNested = []string{
	"app_key",
	"app_key.kek_label",
	"app_key.key",
	"nwk_key",
	"nwk_key.kek_label",
	"nwk_key.key",
	"root_key_id",
}

var RootKeysFieldPathsTopLevel = []string{
	"app_key",
	"nwk_key",
	"root_key_id",
}

func (dst *RootKeys) SetFields(src *RootKeys, paths ...string) error {
	for name, subs := range _processPaths(append(paths[:0:0], paths...)) {
		switch name {
		case "root_key_id":
			if len(subs) > 0 {
				return fmt.Errorf("'root_key_id' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.RootKeyID = src.RootKeyID
			} else {
				var zero string
				dst.RootKeyID = zero
			}
		case "app_key":
			if len(subs) > 0 {
				newDst := dst.AppKey
				if newDst == nil {
					newDst = &KeyEnvelope{}
					dst.AppKey = newDst
				}
				var newSrc *KeyEnvelope
				if src != nil {
					newSrc = src.AppKey
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.AppKey = src.AppKey
				} else {
					dst.AppKey = nil
				}
			}
		case "nwk_key":
			if len(subs) > 0 {
				newDst := dst.NwkKey
				if newDst == nil {
					newDst = &KeyEnvelope{}
					dst.NwkKey = newDst
				}
				var newSrc *KeyEnvelope
				if src != nil {
					newSrc = src.NwkKey
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.NwkKey = src.NwkKey
				} else {
					dst.NwkKey = nil
				}
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

var SessionKeysFieldPathsNested = []string{
	"app_s_key",
	"app_s_key.kek_label",
	"app_s_key.key",
	"f_nwk_s_int_key",
	"f_nwk_s_int_key.kek_label",
	"f_nwk_s_int_key.key",
	"nwk_s_enc_key",
	"nwk_s_enc_key.kek_label",
	"nwk_s_enc_key.key",
	"s_nwk_s_int_key",
	"s_nwk_s_int_key.kek_label",
	"s_nwk_s_int_key.key",
	"session_key_id",
}

var SessionKeysFieldPathsTopLevel = []string{
	"app_s_key",
	"f_nwk_s_int_key",
	"nwk_s_enc_key",
	"s_nwk_s_int_key",
	"session_key_id",
}

func (dst *SessionKeys) SetFields(src *SessionKeys, paths ...string) error {
	for name, subs := range _processPaths(append(paths[:0:0], paths...)) {
		switch name {
		case "session_key_id":
			if len(subs) > 0 {
				return fmt.Errorf("'session_key_id' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.SessionKeyID = src.SessionKeyID
			} else {
				var zero []byte
				dst.SessionKeyID = zero
			}
		case "f_nwk_s_int_key":
			if len(subs) > 0 {
				newDst := dst.FNwkSIntKey
				if newDst == nil {
					newDst = &KeyEnvelope{}
					dst.FNwkSIntKey = newDst
				}
				var newSrc *KeyEnvelope
				if src != nil {
					newSrc = src.FNwkSIntKey
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.FNwkSIntKey = src.FNwkSIntKey
				} else {
					dst.FNwkSIntKey = nil
				}
			}
		case "s_nwk_s_int_key":
			if len(subs) > 0 {
				newDst := dst.SNwkSIntKey
				if newDst == nil {
					newDst = &KeyEnvelope{}
					dst.SNwkSIntKey = newDst
				}
				var newSrc *KeyEnvelope
				if src != nil {
					newSrc = src.SNwkSIntKey
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.SNwkSIntKey = src.SNwkSIntKey
				} else {
					dst.SNwkSIntKey = nil
				}
			}
		case "nwk_s_enc_key":
			if len(subs) > 0 {
				newDst := dst.NwkSEncKey
				if newDst == nil {
					newDst = &KeyEnvelope{}
					dst.NwkSEncKey = newDst
				}
				var newSrc *KeyEnvelope
				if src != nil {
					newSrc = src.NwkSEncKey
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.NwkSEncKey = src.NwkSEncKey
				} else {
					dst.NwkSEncKey = nil
				}
			}
		case "app_s_key":
			if len(subs) > 0 {
				newDst := dst.AppSKey
				if newDst == nil {
					newDst = &KeyEnvelope{}
					dst.AppSKey = newDst
				}
				var newSrc *KeyEnvelope
				if src != nil {
					newSrc = src.AppSKey
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.AppSKey = src.AppSKey
				} else {
					dst.AppSKey = nil
				}
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}
