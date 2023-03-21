// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package v1alpha1

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

func easyjsonC09fa914DecodeGithubComKubewardenK8sObjectsApiAdmissionregistrationV1alpha1(in *jlexer.Lexer, out *ValidatingAdmissionPolicy) {
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
		case "spec":
			if in.IsNull() {
				in.Skip()
				out.Spec = nil
			} else {
				if out.Spec == nil {
					out.Spec = new(ValidatingAdmissionPolicySpec)
				}
				easyjsonC09fa914DecodeGithubComKubewardenK8sObjectsApiAdmissionregistrationV1alpha11(in, out.Spec)
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
func easyjsonC09fa914EncodeGithubComKubewardenK8sObjectsApiAdmissionregistrationV1alpha1(out *jwriter.Writer, in ValidatingAdmissionPolicy) {
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
	if in.Spec != nil {
		const prefix string = ",\"spec\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		easyjsonC09fa914EncodeGithubComKubewardenK8sObjectsApiAdmissionregistrationV1alpha11(out, *in.Spec)
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v ValidatingAdmissionPolicy) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC09fa914EncodeGithubComKubewardenK8sObjectsApiAdmissionregistrationV1alpha1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v ValidatingAdmissionPolicy) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC09fa914EncodeGithubComKubewardenK8sObjectsApiAdmissionregistrationV1alpha1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *ValidatingAdmissionPolicy) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC09fa914DecodeGithubComKubewardenK8sObjectsApiAdmissionregistrationV1alpha1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *ValidatingAdmissionPolicy) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC09fa914DecodeGithubComKubewardenK8sObjectsApiAdmissionregistrationV1alpha1(l, v)
}
func easyjsonC09fa914DecodeGithubComKubewardenK8sObjectsApiAdmissionregistrationV1alpha11(in *jlexer.Lexer, out *ValidatingAdmissionPolicySpec) {
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
		case "failurePolicy":
			out.FailurePolicy = string(in.String())
		case "matchConstraints":
			if in.IsNull() {
				in.Skip()
				out.MatchConstraints = nil
			} else {
				if out.MatchConstraints == nil {
					out.MatchConstraints = new(MatchResources)
				}
				(*out.MatchConstraints).UnmarshalEasyJSON(in)
			}
		case "paramKind":
			if in.IsNull() {
				in.Skip()
				out.ParamKind = nil
			} else {
				if out.ParamKind == nil {
					out.ParamKind = new(ParamKind)
				}
				(*out.ParamKind).UnmarshalEasyJSON(in)
			}
		case "validations":
			if in.IsNull() {
				in.Skip()
				out.Validations = nil
			} else {
				in.Delim('[')
				if out.Validations == nil {
					if !in.IsDelim(']') {
						out.Validations = make([]*Validation, 0, 8)
					} else {
						out.Validations = []*Validation{}
					}
				} else {
					out.Validations = (out.Validations)[:0]
				}
				for !in.IsDelim(']') {
					var v1 *Validation
					if in.IsNull() {
						in.Skip()
						v1 = nil
					} else {
						if v1 == nil {
							v1 = new(Validation)
						}
						easyjsonC09fa914DecodeGithubComKubewardenK8sObjectsApiAdmissionregistrationV1alpha12(in, v1)
					}
					out.Validations = append(out.Validations, v1)
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
func easyjsonC09fa914EncodeGithubComKubewardenK8sObjectsApiAdmissionregistrationV1alpha11(out *jwriter.Writer, in ValidatingAdmissionPolicySpec) {
	out.RawByte('{')
	first := true
	_ = first
	if in.FailurePolicy != "" {
		const prefix string = ",\"failurePolicy\":"
		first = false
		out.RawString(prefix[1:])
		out.String(string(in.FailurePolicy))
	}
	if in.MatchConstraints != nil {
		const prefix string = ",\"matchConstraints\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		(*in.MatchConstraints).MarshalEasyJSON(out)
	}
	if in.ParamKind != nil {
		const prefix string = ",\"paramKind\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		(*in.ParamKind).MarshalEasyJSON(out)
	}
	{
		const prefix string = ",\"validations\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		if in.Validations == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v2, v3 := range in.Validations {
				if v2 > 0 {
					out.RawByte(',')
				}
				if v3 == nil {
					out.RawString("null")
				} else {
					easyjsonC09fa914EncodeGithubComKubewardenK8sObjectsApiAdmissionregistrationV1alpha12(out, *v3)
				}
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}
func easyjsonC09fa914DecodeGithubComKubewardenK8sObjectsApiAdmissionregistrationV1alpha12(in *jlexer.Lexer, out *Validation) {
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
		case "expression":
			if in.IsNull() {
				in.Skip()
				out.Expression = nil
			} else {
				if out.Expression == nil {
					out.Expression = new(string)
				}
				*out.Expression = string(in.String())
			}
		case "message":
			out.Message = string(in.String())
		case "reason":
			out.Reason = string(in.String())
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
func easyjsonC09fa914EncodeGithubComKubewardenK8sObjectsApiAdmissionregistrationV1alpha12(out *jwriter.Writer, in Validation) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"expression\":"
		out.RawString(prefix[1:])
		if in.Expression == nil {
			out.RawString("null")
		} else {
			out.String(string(*in.Expression))
		}
	}
	if in.Message != "" {
		const prefix string = ",\"message\":"
		out.RawString(prefix)
		out.String(string(in.Message))
	}
	if in.Reason != "" {
		const prefix string = ",\"reason\":"
		out.RawString(prefix)
		out.String(string(in.Reason))
	}
	out.RawByte('}')
}
