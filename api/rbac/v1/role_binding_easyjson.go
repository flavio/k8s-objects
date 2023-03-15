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

func easyjsonCff2cfa6DecodeGithubComKubewardenK8sObjectsApiRbacV1(in *jlexer.Lexer, out *RoleBinding) {
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
		case "roleRef":
			if in.IsNull() {
				in.Skip()
				out.RoleRef = nil
			} else {
				if out.RoleRef == nil {
					out.RoleRef = new(RoleRef)
				}
				easyjsonCff2cfa6DecodeGithubComKubewardenK8sObjectsApiRbacV11(in, out.RoleRef)
			}
		case "subjects":
			if in.IsNull() {
				in.Skip()
				out.Subjects = nil
			} else {
				in.Delim('[')
				if out.Subjects == nil {
					if !in.IsDelim(']') {
						out.Subjects = make([]*Subject, 0, 8)
					} else {
						out.Subjects = []*Subject{}
					}
				} else {
					out.Subjects = (out.Subjects)[:0]
				}
				for !in.IsDelim(']') {
					var v1 *Subject
					if in.IsNull() {
						in.Skip()
						v1 = nil
					} else {
						if v1 == nil {
							v1 = new(Subject)
						}
						easyjsonCff2cfa6DecodeGithubComKubewardenK8sObjectsApiRbacV12(in, v1)
					}
					out.Subjects = append(out.Subjects, v1)
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
func easyjsonCff2cfa6EncodeGithubComKubewardenK8sObjectsApiRbacV1(out *jwriter.Writer, in RoleBinding) {
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
	{
		const prefix string = ",\"roleRef\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		if in.RoleRef == nil {
			out.RawString("null")
		} else {
			easyjsonCff2cfa6EncodeGithubComKubewardenK8sObjectsApiRbacV11(out, *in.RoleRef)
		}
	}
	if len(in.Subjects) != 0 {
		const prefix string = ",\"subjects\":"
		out.RawString(prefix)
		{
			out.RawByte('[')
			for v2, v3 := range in.Subjects {
				if v2 > 0 {
					out.RawByte(',')
				}
				if v3 == nil {
					out.RawString("null")
				} else {
					easyjsonCff2cfa6EncodeGithubComKubewardenK8sObjectsApiRbacV12(out, *v3)
				}
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v RoleBinding) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonCff2cfa6EncodeGithubComKubewardenK8sObjectsApiRbacV1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v RoleBinding) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonCff2cfa6EncodeGithubComKubewardenK8sObjectsApiRbacV1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *RoleBinding) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonCff2cfa6DecodeGithubComKubewardenK8sObjectsApiRbacV1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *RoleBinding) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonCff2cfa6DecodeGithubComKubewardenK8sObjectsApiRbacV1(l, v)
}
func easyjsonCff2cfa6DecodeGithubComKubewardenK8sObjectsApiRbacV12(in *jlexer.Lexer, out *Subject) {
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
		case "apiGroup":
			out.APIGroup = string(in.String())
		case "kind":
			if in.IsNull() {
				in.Skip()
				out.Kind = nil
			} else {
				if out.Kind == nil {
					out.Kind = new(string)
				}
				*out.Kind = string(in.String())
			}
		case "name":
			if in.IsNull() {
				in.Skip()
				out.Name = nil
			} else {
				if out.Name == nil {
					out.Name = new(string)
				}
				*out.Name = string(in.String())
			}
		case "namespace":
			out.Namespace = string(in.String())
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
func easyjsonCff2cfa6EncodeGithubComKubewardenK8sObjectsApiRbacV12(out *jwriter.Writer, in Subject) {
	out.RawByte('{')
	first := true
	_ = first
	if in.APIGroup != "" {
		const prefix string = ",\"apiGroup\":"
		first = false
		out.RawString(prefix[1:])
		out.String(string(in.APIGroup))
	}
	{
		const prefix string = ",\"kind\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		if in.Kind == nil {
			out.RawString("null")
		} else {
			out.String(string(*in.Kind))
		}
	}
	{
		const prefix string = ",\"name\":"
		out.RawString(prefix)
		if in.Name == nil {
			out.RawString("null")
		} else {
			out.String(string(*in.Name))
		}
	}
	if in.Namespace != "" {
		const prefix string = ",\"namespace\":"
		out.RawString(prefix)
		out.String(string(in.Namespace))
	}
	out.RawByte('}')
}
func easyjsonCff2cfa6DecodeGithubComKubewardenK8sObjectsApiRbacV11(in *jlexer.Lexer, out *RoleRef) {
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
		case "apiGroup":
			if in.IsNull() {
				in.Skip()
				out.APIGroup = nil
			} else {
				if out.APIGroup == nil {
					out.APIGroup = new(string)
				}
				*out.APIGroup = string(in.String())
			}
		case "kind":
			if in.IsNull() {
				in.Skip()
				out.Kind = nil
			} else {
				if out.Kind == nil {
					out.Kind = new(string)
				}
				*out.Kind = string(in.String())
			}
		case "name":
			if in.IsNull() {
				in.Skip()
				out.Name = nil
			} else {
				if out.Name == nil {
					out.Name = new(string)
				}
				*out.Name = string(in.String())
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
func easyjsonCff2cfa6EncodeGithubComKubewardenK8sObjectsApiRbacV11(out *jwriter.Writer, in RoleRef) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"apiGroup\":"
		out.RawString(prefix[1:])
		if in.APIGroup == nil {
			out.RawString("null")
		} else {
			out.String(string(*in.APIGroup))
		}
	}
	{
		const prefix string = ",\"kind\":"
		out.RawString(prefix)
		if in.Kind == nil {
			out.RawString("null")
		} else {
			out.String(string(*in.Kind))
		}
	}
	{
		const prefix string = ",\"name\":"
		out.RawString(prefix)
		if in.Name == nil {
			out.RawString("null")
		} else {
			out.String(string(*in.Name))
		}
	}
	out.RawByte('}')
}