// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package v1beta1

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

func easyjson4188fa4DecodeGithubComKubewardenK8sObjectsApiStorageV1beta1(in *jlexer.Lexer, out *CSINode) {
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
			(out.Metadata).UnmarshalEasyJSON(in)
		case "spec":
			if in.IsNull() {
				in.Skip()
				out.Spec = nil
			} else {
				if out.Spec == nil {
					out.Spec = new(CSINodeSpec)
				}
				easyjson4188fa4DecodeGithubComKubewardenK8sObjectsApiStorageV1beta11(in, out.Spec)
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
func easyjson4188fa4EncodeGithubComKubewardenK8sObjectsApiStorageV1beta1(out *jwriter.Writer, in CSINode) {
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
	if true {
		const prefix string = ",\"metadata\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		(in.Metadata).MarshalEasyJSON(out)
	}
	{
		const prefix string = ",\"spec\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		if in.Spec == nil {
			out.RawString("null")
		} else {
			easyjson4188fa4EncodeGithubComKubewardenK8sObjectsApiStorageV1beta11(out, *in.Spec)
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v CSINode) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson4188fa4EncodeGithubComKubewardenK8sObjectsApiStorageV1beta1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v CSINode) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson4188fa4EncodeGithubComKubewardenK8sObjectsApiStorageV1beta1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *CSINode) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson4188fa4DecodeGithubComKubewardenK8sObjectsApiStorageV1beta1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *CSINode) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson4188fa4DecodeGithubComKubewardenK8sObjectsApiStorageV1beta1(l, v)
}
func easyjson4188fa4DecodeGithubComKubewardenK8sObjectsApiStorageV1beta11(in *jlexer.Lexer, out *CSINodeSpec) {
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
		case "drivers":
			if in.IsNull() {
				in.Skip()
				out.Drivers = nil
			} else {
				in.Delim('[')
				if out.Drivers == nil {
					if !in.IsDelim(']') {
						out.Drivers = make([]*CSINodeDriver, 0, 8)
					} else {
						out.Drivers = []*CSINodeDriver{}
					}
				} else {
					out.Drivers = (out.Drivers)[:0]
				}
				for !in.IsDelim(']') {
					var v1 *CSINodeDriver
					if in.IsNull() {
						in.Skip()
						v1 = nil
					} else {
						if v1 == nil {
							v1 = new(CSINodeDriver)
						}
						easyjson4188fa4DecodeGithubComKubewardenK8sObjectsApiStorageV1beta12(in, v1)
					}
					out.Drivers = append(out.Drivers, v1)
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
func easyjson4188fa4EncodeGithubComKubewardenK8sObjectsApiStorageV1beta11(out *jwriter.Writer, in CSINodeSpec) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"drivers\":"
		out.RawString(prefix[1:])
		if in.Drivers == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v2, v3 := range in.Drivers {
				if v2 > 0 {
					out.RawByte(',')
				}
				if v3 == nil {
					out.RawString("null")
				} else {
					easyjson4188fa4EncodeGithubComKubewardenK8sObjectsApiStorageV1beta12(out, *v3)
				}
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}
func easyjson4188fa4DecodeGithubComKubewardenK8sObjectsApiStorageV1beta12(in *jlexer.Lexer, out *CSINodeDriver) {
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
		case "nodeID":
			if in.IsNull() {
				in.Skip()
				out.NodeID = nil
			} else {
				if out.NodeID == nil {
					out.NodeID = new(string)
				}
				*out.NodeID = string(in.String())
			}
		case "topologyKeys":
			if in.IsNull() {
				in.Skip()
				out.TopologyKeys = nil
			} else {
				in.Delim('[')
				if out.TopologyKeys == nil {
					if !in.IsDelim(']') {
						out.TopologyKeys = make([]string, 0, 4)
					} else {
						out.TopologyKeys = []string{}
					}
				} else {
					out.TopologyKeys = (out.TopologyKeys)[:0]
				}
				for !in.IsDelim(']') {
					var v4 string
					v4 = string(in.String())
					out.TopologyKeys = append(out.TopologyKeys, v4)
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
func easyjson4188fa4EncodeGithubComKubewardenK8sObjectsApiStorageV1beta12(out *jwriter.Writer, in CSINodeDriver) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"name\":"
		out.RawString(prefix[1:])
		if in.Name == nil {
			out.RawString("null")
		} else {
			out.String(string(*in.Name))
		}
	}
	{
		const prefix string = ",\"nodeID\":"
		out.RawString(prefix)
		if in.NodeID == nil {
			out.RawString("null")
		} else {
			out.String(string(*in.NodeID))
		}
	}
	{
		const prefix string = ",\"topologyKeys\":"
		out.RawString(prefix)
		if in.TopologyKeys == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v5, v6 := range in.TopologyKeys {
				if v5 > 0 {
					out.RawByte(',')
				}
				out.String(string(v6))
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}
