// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package v1

import (
	json "encoding/json"
	_v1 "github.com/kubewarden/k8s-objects/apimachinery/pkg/apis/meta/v1"
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

func easyjsonA448722eDecodeGithubComKubewardenK8sObjectsApiAdmissionregistrationV1(in *jlexer.Lexer, out *MutatingWebhookConfiguration) {
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
		case "apiVersion":
			out.APIVersion = string(in.String())
		case "kind":
			out.Kind = string(in.String())
		case "metadata":
			if in.IsNull() {
				in.Skip()
				out.Metadata = nil
			} else {
				if out.Metadata == nil {
					out.Metadata = new(_v1.ObjectMeta)
				}
				(*out.Metadata).UnmarshalEasyJSON(in)
			}
		case "webhooks":
			if in.IsNull() {
				in.Skip()
				out.Webhooks = nil
			} else {
				in.Delim('[')
				if out.Webhooks == nil {
					if !in.IsDelim(']') {
						out.Webhooks = make([]*MutatingWebhook, 0, 8)
					} else {
						out.Webhooks = []*MutatingWebhook{}
					}
				} else {
					out.Webhooks = (out.Webhooks)[:0]
				}
				for !in.IsDelim(']') {
					var v1 *MutatingWebhook
					if in.IsNull() {
						in.Skip()
						v1 = nil
					} else {
						if v1 == nil {
							v1 = new(MutatingWebhook)
						}
						(*v1).UnmarshalEasyJSON(in)
					}
					out.Webhooks = append(out.Webhooks, v1)
					in.WantComma()
				}
				in.Delim(']')
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
func easyjsonA448722eEncodeGithubComKubewardenK8sObjectsApiAdmissionregistrationV1(out *jwriter.Writer, in MutatingWebhookConfiguration) {
	out.RawByte('{')
	first := true
	_ = first
	if in.APIVersion != "" {
		const prefix string = ",\"apiVersion\":"
		first = false
		out.RawString(prefix[1:])
		out.String(string(in.APIVersion))
	}
	if in.Kind != "" {
		const prefix string = ",\"kind\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Kind))
	}
	if in.Metadata != nil {
		const prefix string = ",\"metadata\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		(*in.Metadata).MarshalEasyJSON(out)
	}
	if len(in.Webhooks) != 0 {
		const prefix string = ",\"webhooks\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('[')
			for v2, v3 := range in.Webhooks {
				if v2 > 0 {
					out.RawByte(',')
				}
				if v3 == nil {
					out.RawString("null")
				} else {
					(*v3).MarshalEasyJSON(out)
				}
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v MutatingWebhookConfiguration) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonA448722eEncodeGithubComKubewardenK8sObjectsApiAdmissionregistrationV1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v MutatingWebhookConfiguration) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonA448722eEncodeGithubComKubewardenK8sObjectsApiAdmissionregistrationV1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *MutatingWebhookConfiguration) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonA448722eDecodeGithubComKubewardenK8sObjectsApiAdmissionregistrationV1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *MutatingWebhookConfiguration) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonA448722eDecodeGithubComKubewardenK8sObjectsApiAdmissionregistrationV1(l, v)
}