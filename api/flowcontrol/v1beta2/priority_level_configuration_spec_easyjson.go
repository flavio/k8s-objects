// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package v1beta2

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

func easyjson7155abb8DecodeGithubComKubewardenK8sObjectsApiFlowcontrolV1beta2(in *jlexer.Lexer, out *PriorityLevelConfigurationSpec) {
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
		case "limited":
			if in.IsNull() {
				in.Skip()
				out.Limited = nil
			} else {
				if out.Limited == nil {
					out.Limited = new(LimitedPriorityLevelConfiguration)
				}
				(*out.Limited).UnmarshalEasyJSON(in)
			}
		case "type":
			if in.IsNull() {
				in.Skip()
				out.Type = nil
			} else {
				if out.Type == nil {
					out.Type = new(string)
				}
				*out.Type = string(in.String())
			}
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
func easyjson7155abb8EncodeGithubComKubewardenK8sObjectsApiFlowcontrolV1beta2(out *jwriter.Writer, in PriorityLevelConfigurationSpec) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Limited != nil {
		const prefix string = ",\"limited\":"
		first = false
		out.RawString(prefix[1:])
		(*in.Limited).MarshalEasyJSON(out)
	}
	{
		const prefix string = ",\"type\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		if in.Type == nil {
			out.RawString("null")
		} else {
			out.String(string(*in.Type))
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v PriorityLevelConfigurationSpec) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson7155abb8EncodeGithubComKubewardenK8sObjectsApiFlowcontrolV1beta2(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v PriorityLevelConfigurationSpec) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson7155abb8EncodeGithubComKubewardenK8sObjectsApiFlowcontrolV1beta2(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *PriorityLevelConfigurationSpec) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson7155abb8DecodeGithubComKubewardenK8sObjectsApiFlowcontrolV1beta2(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *PriorityLevelConfigurationSpec) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson7155abb8DecodeGithubComKubewardenK8sObjectsApiFlowcontrolV1beta2(l, v)
}
