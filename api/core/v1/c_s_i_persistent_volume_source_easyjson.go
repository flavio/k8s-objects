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

func easyjson92d5e37cDecodeGithubComKubewardenK8sObjectsApiCoreV1(in *jlexer.Lexer, out *CSIPersistentVolumeSource) {
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
		case "controllerPublishSecretRef":
			if in.IsNull() {
				in.Skip()
				out.ControllerPublishSecretRef = nil
			} else {
				if out.ControllerPublishSecretRef == nil {
					out.ControllerPublishSecretRef = new(SecretReference)
				}
				easyjson92d5e37cDecodeGithubComKubewardenK8sObjectsApiCoreV11(in, out.ControllerPublishSecretRef)
			}
		case "driver":
			if in.IsNull() {
				in.Skip()
				out.Driver = nil
			} else {
				if out.Driver == nil {
					out.Driver = new(string)
				}
				*out.Driver = string(in.String())
			}
		case "fsType":
			out.FsType = string(in.String())
		case "nodePublishSecretRef":
			if in.IsNull() {
				in.Skip()
				out.NodePublishSecretRef = nil
			} else {
				if out.NodePublishSecretRef == nil {
					out.NodePublishSecretRef = new(SecretReference)
				}
				easyjson92d5e37cDecodeGithubComKubewardenK8sObjectsApiCoreV11(in, out.NodePublishSecretRef)
			}
		case "nodeStageSecretRef":
			if in.IsNull() {
				in.Skip()
				out.NodeStageSecretRef = nil
			} else {
				if out.NodeStageSecretRef == nil {
					out.NodeStageSecretRef = new(SecretReference)
				}
				easyjson92d5e37cDecodeGithubComKubewardenK8sObjectsApiCoreV11(in, out.NodeStageSecretRef)
			}
		case "readOnly":
			out.ReadOnly = bool(in.Bool())
		case "volumeAttributes":
			if in.IsNull() {
				in.Skip()
			} else {
				in.Delim('{')
				if !in.IsDelim('}') {
					out.VolumeAttributes = make(map[string]string)
				} else {
					out.VolumeAttributes = nil
				}
				for !in.IsDelim('}') {
					key := string(in.String())
					in.WantColon()
					var v1 string
					v1 = string(in.String())
					(out.VolumeAttributes)[key] = v1
					in.WantComma()
				}
				in.Delim('}')
			}
		case "volumeHandle":
			if in.IsNull() {
				in.Skip()
				out.VolumeHandle = nil
			} else {
				if out.VolumeHandle == nil {
					out.VolumeHandle = new(string)
				}
				*out.VolumeHandle = string(in.String())
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
func easyjson92d5e37cEncodeGithubComKubewardenK8sObjectsApiCoreV1(out *jwriter.Writer, in CSIPersistentVolumeSource) {
	out.RawByte('{')
	first := true
	_ = first
	if in.ControllerPublishSecretRef != nil {
		const prefix string = ",\"controllerPublishSecretRef\":"
		first = false
		out.RawString(prefix[1:])
		easyjson92d5e37cEncodeGithubComKubewardenK8sObjectsApiCoreV11(out, *in.ControllerPublishSecretRef)
	}
	{
		const prefix string = ",\"driver\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		if in.Driver == nil {
			out.RawString("null")
		} else {
			out.String(string(*in.Driver))
		}
	}
	if in.FsType != "" {
		const prefix string = ",\"fsType\":"
		out.RawString(prefix)
		out.String(string(in.FsType))
	}
	if in.NodePublishSecretRef != nil {
		const prefix string = ",\"nodePublishSecretRef\":"
		out.RawString(prefix)
		easyjson92d5e37cEncodeGithubComKubewardenK8sObjectsApiCoreV11(out, *in.NodePublishSecretRef)
	}
	if in.NodeStageSecretRef != nil {
		const prefix string = ",\"nodeStageSecretRef\":"
		out.RawString(prefix)
		easyjson92d5e37cEncodeGithubComKubewardenK8sObjectsApiCoreV11(out, *in.NodeStageSecretRef)
	}
	if in.ReadOnly {
		const prefix string = ",\"readOnly\":"
		out.RawString(prefix)
		out.Bool(bool(in.ReadOnly))
	}
	if len(in.VolumeAttributes) != 0 {
		const prefix string = ",\"volumeAttributes\":"
		out.RawString(prefix)
		{
			out.RawByte('{')
			v2First := true
			for v2Name, v2Value := range in.VolumeAttributes {
				if v2First {
					v2First = false
				} else {
					out.RawByte(',')
				}
				out.String(string(v2Name))
				out.RawByte(':')
				out.String(string(v2Value))
			}
			out.RawByte('}')
		}
	}
	{
		const prefix string = ",\"volumeHandle\":"
		out.RawString(prefix)
		if in.VolumeHandle == nil {
			out.RawString("null")
		} else {
			out.String(string(*in.VolumeHandle))
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v CSIPersistentVolumeSource) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson92d5e37cEncodeGithubComKubewardenK8sObjectsApiCoreV1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v CSIPersistentVolumeSource) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson92d5e37cEncodeGithubComKubewardenK8sObjectsApiCoreV1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *CSIPersistentVolumeSource) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson92d5e37cDecodeGithubComKubewardenK8sObjectsApiCoreV1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *CSIPersistentVolumeSource) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson92d5e37cDecodeGithubComKubewardenK8sObjectsApiCoreV1(l, v)
}
func easyjson92d5e37cDecodeGithubComKubewardenK8sObjectsApiCoreV11(in *jlexer.Lexer, out *SecretReference) {
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
			out.Name = string(in.String())
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
func easyjson92d5e37cEncodeGithubComKubewardenK8sObjectsApiCoreV11(out *jwriter.Writer, in SecretReference) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Name != "" {
		const prefix string = ",\"name\":"
		first = false
		out.RawString(prefix[1:])
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
	out.RawByte('}')
}
