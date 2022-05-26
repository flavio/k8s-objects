// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package v1

import (
	json "encoding/json"
	easyjson "github.com/mailru/easyjson"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
)

// suppress unused package warning
var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

func easyjsonE311e2eeDecodeGithubComKubewardenK8sObjectsApiAuthorizationV1(in *jlexer.Lexer, out *ResourceAttributes) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "group":
			out.Group = string(in.String())
		case "name":
			out.Name = string(in.String())
		case "namespace":
			out.Namespace = string(in.String())
		case "resource":
			out.Resource = string(in.String())
		case "subresource":
			out.Subresource = string(in.String())
		case "verb":
			out.Verb = string(in.String())
		case "version":
			out.Version = string(in.String())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonE311e2eeEncodeGithubComKubewardenK8sObjectsApiAuthorizationV1(out *jwriter.Writer, in ResourceAttributes) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Group != "" {
		const prefix string = ",\"group\":"
		first = false
		out.RawString(prefix[1:])
		out.String(string(in.Group))
	}
	if in.Name != "" {
		const prefix string = ",\"name\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Name))
	}
	if in.Namespace != "" {
		const prefix string = ",\"namespace\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Namespace))
	}
	if in.Resource != "" {
		const prefix string = ",\"resource\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Resource))
	}
	if in.Subresource != "" {
		const prefix string = ",\"subresource\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Subresource))
	}
	if in.Verb != "" {
		const prefix string = ",\"verb\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Verb))
	}
	if in.Version != "" {
		const prefix string = ",\"version\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Version))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v ResourceAttributes) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonE311e2eeEncodeGithubComKubewardenK8sObjectsApiAuthorizationV1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v ResourceAttributes) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonE311e2eeEncodeGithubComKubewardenK8sObjectsApiAuthorizationV1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *ResourceAttributes) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonE311e2eeDecodeGithubComKubewardenK8sObjectsApiAuthorizationV1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *ResourceAttributes) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonE311e2eeDecodeGithubComKubewardenK8sObjectsApiAuthorizationV1(l, v)
}
