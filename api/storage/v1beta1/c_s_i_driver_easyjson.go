// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package v1beta1

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

func easyjson86ba675aDecodeGithubComKubewardenK8sObjectsApiStorageV1beta1(in *jlexer.Lexer, out *CSIDriver) {
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
					out.Spec = new(CSIDriverSpec)
				}
				easyjson86ba675aDecodeGithubComKubewardenK8sObjectsApiStorageV1beta11(in, out.Spec)
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
func easyjson86ba675aEncodeGithubComKubewardenK8sObjectsApiStorageV1beta1(out *jwriter.Writer, in CSIDriver) {
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
			easyjson86ba675aEncodeGithubComKubewardenK8sObjectsApiStorageV1beta11(out, *in.Spec)
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v CSIDriver) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson86ba675aEncodeGithubComKubewardenK8sObjectsApiStorageV1beta1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v CSIDriver) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson86ba675aEncodeGithubComKubewardenK8sObjectsApiStorageV1beta1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *CSIDriver) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson86ba675aDecodeGithubComKubewardenK8sObjectsApiStorageV1beta1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *CSIDriver) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson86ba675aDecodeGithubComKubewardenK8sObjectsApiStorageV1beta1(l, v)
}
func easyjson86ba675aDecodeGithubComKubewardenK8sObjectsApiStorageV1beta11(in *jlexer.Lexer, out *CSIDriverSpec) {
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
		case "attachRequired":
			out.AttachRequired = bool(in.Bool())
		case "fsGroupPolicy":
			out.FsGroupPolicy = string(in.String())
		case "podInfoOnMount":
			out.PodInfoOnMount = bool(in.Bool())
		case "storageCapacity":
			out.StorageCapacity = bool(in.Bool())
		case "volumeLifecycleModes":
			if in.IsNull() {
				in.Skip()
				out.VolumeLifecycleModes = nil
			} else {
				in.Delim('[')
				if out.VolumeLifecycleModes == nil {
					if !in.IsDelim(']') {
						out.VolumeLifecycleModes = make([]string, 0, 4)
					} else {
						out.VolumeLifecycleModes = []string{}
					}
				} else {
					out.VolumeLifecycleModes = (out.VolumeLifecycleModes)[:0]
				}
				for !in.IsDelim(']') {
					var v1 string
					v1 = string(in.String())
					out.VolumeLifecycleModes = append(out.VolumeLifecycleModes, v1)
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
func easyjson86ba675aEncodeGithubComKubewardenK8sObjectsApiStorageV1beta11(out *jwriter.Writer, in CSIDriverSpec) {
	out.RawByte('{')
	first := true
	_ = first
	if in.AttachRequired {
		const prefix string = ",\"attachRequired\":"
		first = false
		out.RawString(prefix[1:])
		out.Bool(bool(in.AttachRequired))
	}
	if in.FsGroupPolicy != "" {
		const prefix string = ",\"fsGroupPolicy\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.FsGroupPolicy))
	}
	if in.PodInfoOnMount {
		const prefix string = ",\"podInfoOnMount\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Bool(bool(in.PodInfoOnMount))
	}
	if in.StorageCapacity {
		const prefix string = ",\"storageCapacity\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Bool(bool(in.StorageCapacity))
	}
	if len(in.VolumeLifecycleModes) != 0 {
		const prefix string = ",\"volumeLifecycleModes\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('[')
			for v2, v3 := range in.VolumeLifecycleModes {
				if v2 > 0 {
					out.RawByte(',')
				}
				out.String(string(v3))
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}
