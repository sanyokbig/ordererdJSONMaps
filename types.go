package main

import (
	"sort"

	"github.com/mailru/easyjson/jlexer"
	"github.com/mailru/easyjson/jwriter"
)

/*
	keys := []string{}
	for key, _ := range v {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	for _, v2Name := range keys {
		if v2First {
			v2First = false
		} else {
			w.RawByte(',')
		}
		w.String(string(v2Name))
		w.RawByte(':')
		w.Int(int(v[v2Name]))
	}

*/

type OrderedMapStringToInt map[string]int

// MarshalJSON supports json.Marshaler interface
func (v OrderedMapStringToInt) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	v.encodeJSON(&w)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v OrderedMapStringToInt) MarshalEasyJSON(w *jwriter.Writer) {
	v.encodeJSON(w)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *OrderedMapStringToInt) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	v.decodeJSON(&r)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *OrderedMapStringToInt) UnmarshalEasyJSON(l *jlexer.Lexer) {
	v.decodeJSON(l)
}

// Internal encode and decode methods copied from easyjson generated code
func (v OrderedMapStringToInt) encodeJSON(out *jwriter.Writer) {
	if v == nil && (out.Flags&jwriter.NilMapAsEmpty) == 0 {
		out.RawString(`null`)
	} else {
		out.RawByte('{')
		v2First := true
		// This part edited to use ordered map keys instead of simple range
		keys := []string{}
		for key, _ := range v {
			keys = append(keys, key)
		}
		sort.Strings(keys)
		for _, v2Name := range keys {
			if v2First {
				v2First = false
			} else {
				out.RawByte(',')
			}
			out.String(string(v2Name))
			out.RawByte(':')
			out.Int(int(v[v2Name]))
		}
		out.RawByte('}')
	}
}
func (v *OrderedMapStringToInt) decodeJSON(in *jlexer.Lexer) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
	} else {
		in.Delim('{')
		if !in.IsDelim('}') {
			*v = make(OrderedMapStringToInt)
		} else {
			*v = nil
		}
		for !in.IsDelim('}') {
			key := string(in.String())
			in.WantColon()
			var v1 int
			v1 = int(in.Int())
			(*v)[key] = v1
			in.WantComma()
		}
		in.Delim('}')
	}
	if isTopLevel {
		in.Consumed()
	}
}

type OrderedMapStringToFilter map[string]Filter

// MarshalJSON supports json.Marshaler interface
func (v OrderedMapStringToFilter) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	v.encodeJSON(&w)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v OrderedMapStringToFilter) MarshalEasyJSON(w *jwriter.Writer) {
	v.encodeJSON(w)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *OrderedMapStringToFilter) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	v.decodeJSON(&r)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *OrderedMapStringToFilter) UnmarshalEasyJSON(l *jlexer.Lexer) {
	v.decodeJSON(l)
}

// Internal encode and decode methods copied from easyjson generated code
func (v OrderedMapStringToFilter) encodeJSON(out *jwriter.Writer) {
	if v == nil && (out.Flags&jwriter.NilMapAsEmpty) == 0 {
		out.RawString(`null`)
	} else {
		out.RawByte('{')
		v2First := true
		keys := []string{}
		for key, _ := range v {
			keys = append(keys, key)
		}
		sort.Strings(keys)
		for _, v2Name := range keys {
			if v2First {
				v2First = false
			} else {
				out.RawByte(',')
			}
			out.String(string(v2Name))
			out.RawByte(':')
			(v[v2Name]).MarshalEasyJSON(out)
		}
		out.RawByte('}')
	}
}
func (v *OrderedMapStringToFilter) decodeJSON(in *jlexer.Lexer) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
	} else {
		in.Delim('{')
		if !in.IsDelim('}') {
			*v = make(OrderedMapStringToFilter)
		} else {
			*v = nil
		}
		for !in.IsDelim('}') {
			key := string(in.String())
			in.WantColon()
			var v1 Filter
			(v1).UnmarshalEasyJSON(in)
			(*v)[key] = v1
			in.WantComma()
		}
		in.Delim('}')
	}
	if isTopLevel {
		in.Consumed()
	}
}
