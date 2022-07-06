// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package v1alpha1

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

func easyjson976908b2DecodeGithubComKubewardenK8sObjectsApiAuditregistrationV1alpha1(in *jlexer.Lexer, out *WebhookClientConfig) {
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
		case "caBundle":
			if in.IsNull() {
				in.Skip()
				out.CaBundle = nil
			} else {
				out.CaBundle = in.Bytes()
			}
		case "service":
			if in.IsNull() {
				in.Skip()
				out.Service = nil
			} else {
				if out.Service == nil {
					out.Service = new(ServiceReference)
				}
				(*out.Service).UnmarshalEasyJSON(in)
			}
		case "url":
			out.URL = string(in.String())
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
func easyjson976908b2EncodeGithubComKubewardenK8sObjectsApiAuditregistrationV1alpha1(out *jwriter.Writer, in WebhookClientConfig) {
	out.RawByte('{')
	first := true
	_ = first
	if len(in.CaBundle) != 0 {
		const prefix string = ",\"caBundle\":"
		first = false
		out.RawString(prefix[1:])
		out.Base64Bytes(in.CaBundle)
	}
	if in.Service != nil {
		const prefix string = ",\"service\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		(*in.Service).MarshalEasyJSON(out)
	}
	if in.URL != "" {
		const prefix string = ",\"url\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.URL))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v WebhookClientConfig) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson976908b2EncodeGithubComKubewardenK8sObjectsApiAuditregistrationV1alpha1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v WebhookClientConfig) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson976908b2EncodeGithubComKubewardenK8sObjectsApiAuditregistrationV1alpha1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *WebhookClientConfig) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson976908b2DecodeGithubComKubewardenK8sObjectsApiAuditregistrationV1alpha1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *WebhookClientConfig) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson976908b2DecodeGithubComKubewardenK8sObjectsApiAuditregistrationV1alpha1(l, v)
}