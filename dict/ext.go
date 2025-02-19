// Code generated by radius-dict-gen. DO NOT EDIT.

package dict

import (
	"layeh.com/radius"
)

const (
	UserID_Type    radius.Type = 222
	UserRealm_Type radius.Type = 223
	UserIMSI_Type  radius.Type = 224
)

func UserID_Add(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	p.Add(UserID_Type, a)
	return
}

func UserID_AddString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	p.Add(UserID_Type, a)
	return
}

func UserID_Get(p *radius.Packet) (value []byte) {
	value, _ = UserID_Lookup(p)
	return
}

func UserID_GetString(p *radius.Packet) (value string) {
	value, _ = UserID_LookupString(p)
	return
}

func UserID_Gets(p *radius.Packet) (values [][]byte, err error) {
	var i []byte
	for _, avp := range p.Attributes {
		if avp.Type != UserID_Type {
			continue
		}
		attr := avp.Attribute
		i = radius.Bytes(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func UserID_GetStrings(p *radius.Packet) (values []string, err error) {
	var i string
	for _, avp := range p.Attributes {
		if avp.Type != UserID_Type {
			continue
		}
		attr := avp.Attribute
		i = radius.String(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func UserID_Lookup(p *radius.Packet) (value []byte, err error) {
	a, ok := p.Lookup(UserID_Type)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.Bytes(a)
	return
}

func UserID_LookupString(p *radius.Packet) (value string, err error) {
	a, ok := p.Lookup(UserID_Type)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.String(a)
	return
}

func UserID_Set(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	p.Set(UserID_Type, a)
	return
}

func UserID_SetString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	p.Set(UserID_Type, a)
	return
}

func UserID_Del(p *radius.Packet) {
	p.Attributes.Del(UserID_Type)
}

func UserRealm_Add(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	p.Add(UserRealm_Type, a)
	return
}

func UserRealm_AddString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	p.Add(UserRealm_Type, a)
	return
}

func UserRealm_Get(p *radius.Packet) (value []byte) {
	value, _ = UserRealm_Lookup(p)
	return
}

func UserRealm_GetString(p *radius.Packet) (value string) {
	value, _ = UserRealm_LookupString(p)
	return
}

func UserRealm_Gets(p *radius.Packet) (values [][]byte, err error) {
	var i []byte
	for _, avp := range p.Attributes {
		if avp.Type != UserRealm_Type {
			continue
		}
		attr := avp.Attribute
		i = radius.Bytes(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func UserRealm_GetStrings(p *radius.Packet) (values []string, err error) {
	var i string
	for _, avp := range p.Attributes {
		if avp.Type != UserRealm_Type {
			continue
		}
		attr := avp.Attribute
		i = radius.String(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func UserRealm_Lookup(p *radius.Packet) (value []byte, err error) {
	a, ok := p.Lookup(UserRealm_Type)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.Bytes(a)
	return
}

func UserRealm_LookupString(p *radius.Packet) (value string, err error) {
	a, ok := p.Lookup(UserRealm_Type)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.String(a)
	return
}

func UserRealm_Set(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	p.Set(UserRealm_Type, a)
	return
}

func UserRealm_SetString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	p.Set(UserRealm_Type, a)
	return
}

func UserRealm_Del(p *radius.Packet) {
	p.Attributes.Del(UserRealm_Type)
}

func UserIMSI_Add(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	p.Add(UserIMSI_Type, a)
	return
}

func UserIMSI_AddString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	p.Add(UserIMSI_Type, a)
	return
}

func UserIMSI_Get(p *radius.Packet) (value []byte) {
	value, _ = UserIMSI_Lookup(p)
	return
}

func UserIMSI_GetString(p *radius.Packet) (value string) {
	value, _ = UserIMSI_LookupString(p)
	return
}

func UserIMSI_Gets(p *radius.Packet) (values [][]byte, err error) {
	var i []byte
	for _, avp := range p.Attributes {
		if avp.Type != UserIMSI_Type {
			continue
		}
		attr := avp.Attribute
		i = radius.Bytes(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func UserIMSI_GetStrings(p *radius.Packet) (values []string, err error) {
	var i string
	for _, avp := range p.Attributes {
		if avp.Type != UserIMSI_Type {
			continue
		}
		attr := avp.Attribute
		i = radius.String(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func UserIMSI_Lookup(p *radius.Packet) (value []byte, err error) {
	a, ok := p.Lookup(UserIMSI_Type)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.Bytes(a)
	return
}

func UserIMSI_LookupString(p *radius.Packet) (value string, err error) {
	a, ok := p.Lookup(UserIMSI_Type)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.String(a)
	return
}

func UserIMSI_Set(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	p.Set(UserIMSI_Type, a)
	return
}

func UserIMSI_SetString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	p.Set(UserIMSI_Type, a)
	return
}

func UserIMSI_Del(p *radius.Packet) {
	p.Attributes.Del(UserIMSI_Type)
}
