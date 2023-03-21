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

func easyjson2f078869DecodeGithubComKubewardenK8sObjectsApiCoreV1(in *jlexer.Lexer, out *ProjectedVolumeSource) {
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
		case "defaultMode":
			out.DefaultMode = int32(in.Int32())
		case "sources":
			if in.IsNull() {
				in.Skip()
				out.Sources = nil
			} else {
				in.Delim('[')
				if out.Sources == nil {
					if !in.IsDelim(']') {
						out.Sources = make([]*VolumeProjection, 0, 8)
					} else {
						out.Sources = []*VolumeProjection{}
					}
				} else {
					out.Sources = (out.Sources)[:0]
				}
				for !in.IsDelim(']') {
					var v1 *VolumeProjection
					if in.IsNull() {
						in.Skip()
						v1 = nil
					} else {
						if v1 == nil {
							v1 = new(VolumeProjection)
						}
						easyjson2f078869DecodeGithubComKubewardenK8sObjectsApiCoreV11(in, v1)
					}
					out.Sources = append(out.Sources, v1)
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
func easyjson2f078869EncodeGithubComKubewardenK8sObjectsApiCoreV1(out *jwriter.Writer, in ProjectedVolumeSource) {
	out.RawByte('{')
	first := true
	_ = first
	if in.DefaultMode != 0 {
		const prefix string = ",\"defaultMode\":"
		first = false
		out.RawString(prefix[1:])
		out.Int32(int32(in.DefaultMode))
	}
	if len(in.Sources) != 0 {
		const prefix string = ",\"sources\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('[')
			for v2, v3 := range in.Sources {
				if v2 > 0 {
					out.RawByte(',')
				}
				if v3 == nil {
					out.RawString("null")
				} else {
					easyjson2f078869EncodeGithubComKubewardenK8sObjectsApiCoreV11(out, *v3)
				}
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v ProjectedVolumeSource) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson2f078869EncodeGithubComKubewardenK8sObjectsApiCoreV1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v ProjectedVolumeSource) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson2f078869EncodeGithubComKubewardenK8sObjectsApiCoreV1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *ProjectedVolumeSource) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson2f078869DecodeGithubComKubewardenK8sObjectsApiCoreV1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *ProjectedVolumeSource) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson2f078869DecodeGithubComKubewardenK8sObjectsApiCoreV1(l, v)
}
func easyjson2f078869DecodeGithubComKubewardenK8sObjectsApiCoreV11(in *jlexer.Lexer, out *VolumeProjection) {
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
		case "configMap":
			if in.IsNull() {
				in.Skip()
				out.ConfigMap = nil
			} else {
				if out.ConfigMap == nil {
					out.ConfigMap = new(ConfigMapProjection)
				}
				(*out.ConfigMap).UnmarshalEasyJSON(in)
			}
		case "downwardAPI":
			if in.IsNull() {
				in.Skip()
				out.DownwardAPI = nil
			} else {
				if out.DownwardAPI == nil {
					out.DownwardAPI = new(DownwardAPIProjection)
				}
				(*out.DownwardAPI).UnmarshalEasyJSON(in)
			}
		case "secret":
			if in.IsNull() {
				in.Skip()
				out.Secret = nil
			} else {
				if out.Secret == nil {
					out.Secret = new(SecretProjection)
				}
				easyjson2f078869DecodeGithubComKubewardenK8sObjectsApiCoreV12(in, out.Secret)
			}
		case "serviceAccountToken":
			if in.IsNull() {
				in.Skip()
				out.ServiceAccountToken = nil
			} else {
				if out.ServiceAccountToken == nil {
					out.ServiceAccountToken = new(ServiceAccountTokenProjection)
				}
				easyjson2f078869DecodeGithubComKubewardenK8sObjectsApiCoreV13(in, out.ServiceAccountToken)
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
func easyjson2f078869EncodeGithubComKubewardenK8sObjectsApiCoreV11(out *jwriter.Writer, in VolumeProjection) {
	out.RawByte('{')
	first := true
	_ = first
	if in.ConfigMap != nil {
		const prefix string = ",\"configMap\":"
		first = false
		out.RawString(prefix[1:])
		(*in.ConfigMap).MarshalEasyJSON(out)
	}
	if in.DownwardAPI != nil {
		const prefix string = ",\"downwardAPI\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		(*in.DownwardAPI).MarshalEasyJSON(out)
	}
	if in.Secret != nil {
		const prefix string = ",\"secret\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		easyjson2f078869EncodeGithubComKubewardenK8sObjectsApiCoreV12(out, *in.Secret)
	}
	if in.ServiceAccountToken != nil {
		const prefix string = ",\"serviceAccountToken\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		easyjson2f078869EncodeGithubComKubewardenK8sObjectsApiCoreV13(out, *in.ServiceAccountToken)
	}
	out.RawByte('}')
}
func easyjson2f078869DecodeGithubComKubewardenK8sObjectsApiCoreV13(in *jlexer.Lexer, out *ServiceAccountTokenProjection) {
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
		case "audience":
			out.Audience = string(in.String())
		case "expirationSeconds":
			out.ExpirationSeconds = int64(in.Int64())
		case "path":
			if in.IsNull() {
				in.Skip()
				out.Path = nil
			} else {
				if out.Path == nil {
					out.Path = new(string)
				}
				*out.Path = string(in.String())
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
func easyjson2f078869EncodeGithubComKubewardenK8sObjectsApiCoreV13(out *jwriter.Writer, in ServiceAccountTokenProjection) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Audience != "" {
		const prefix string = ",\"audience\":"
		first = false
		out.RawString(prefix[1:])
		out.String(string(in.Audience))
	}
	if in.ExpirationSeconds != 0 {
		const prefix string = ",\"expirationSeconds\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.ExpirationSeconds))
	}
	{
		const prefix string = ",\"path\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		if in.Path == nil {
			out.RawString("null")
		} else {
			out.String(string(*in.Path))
		}
	}
	out.RawByte('}')
}
func easyjson2f078869DecodeGithubComKubewardenK8sObjectsApiCoreV12(in *jlexer.Lexer, out *SecretProjection) {
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
		case "items":
			if in.IsNull() {
				in.Skip()
				out.Items = nil
			} else {
				in.Delim('[')
				if out.Items == nil {
					if !in.IsDelim(']') {
						out.Items = make([]*KeyToPath, 0, 8)
					} else {
						out.Items = []*KeyToPath{}
					}
				} else {
					out.Items = (out.Items)[:0]
				}
				for !in.IsDelim(']') {
					var v4 *KeyToPath
					if in.IsNull() {
						in.Skip()
						v4 = nil
					} else {
						if v4 == nil {
							v4 = new(KeyToPath)
						}
						(*v4).UnmarshalEasyJSON(in)
					}
					out.Items = append(out.Items, v4)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "name":
			out.Name = string(in.String())
		case "optional":
			out.Optional = bool(in.Bool())
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
func easyjson2f078869EncodeGithubComKubewardenK8sObjectsApiCoreV12(out *jwriter.Writer, in SecretProjection) {
	out.RawByte('{')
	first := true
	_ = first
	if len(in.Items) != 0 {
		const prefix string = ",\"items\":"
		first = false
		out.RawString(prefix[1:])
		{
			out.RawByte('[')
			for v5, v6 := range in.Items {
				if v5 > 0 {
					out.RawByte(',')
				}
				if v6 == nil {
					out.RawString("null")
				} else {
					(*v6).MarshalEasyJSON(out)
				}
			}
			out.RawByte(']')
		}
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
	if in.Optional {
		const prefix string = ",\"optional\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Bool(bool(in.Optional))
	}
	out.RawByte('}')
}
