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

func easyjson2dce4aeDecodeGithubComKubewardenK8sObjectsApiCoreV1(in *jlexer.Lexer, out *SecurityContext) {
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
		case "allowPrivilegeEscalation":
			out.AllowPrivilegeEscalation = bool(in.Bool())
		case "capabilities":
			if in.IsNull() {
				in.Skip()
				out.Capabilities = nil
			} else {
				if out.Capabilities == nil {
					out.Capabilities = new(Capabilities)
				}
				(*out.Capabilities).UnmarshalEasyJSON(in)
			}
		case "privileged":
			out.Privileged = bool(in.Bool())
		case "procMount":
			out.ProcMount = string(in.String())
		case "readOnlyRootFilesystem":
			out.ReadOnlyRootFilesystem = bool(in.Bool())
		case "runAsGroup":
			out.RunAsGroup = int64(in.Int64())
		case "runAsNonRoot":
			out.RunAsNonRoot = bool(in.Bool())
		case "runAsUser":
			out.RunAsUser = int64(in.Int64())
		case "seLinuxOptions":
			if in.IsNull() {
				in.Skip()
				out.SeLinuxOptions = nil
			} else {
				if out.SeLinuxOptions == nil {
					out.SeLinuxOptions = new(SELinuxOptions)
				}
				(*out.SeLinuxOptions).UnmarshalEasyJSON(in)
			}
		case "seccompProfile":
			if in.IsNull() {
				in.Skip()
				out.SeccompProfile = nil
			} else {
				if out.SeccompProfile == nil {
					out.SeccompProfile = new(SeccompProfile)
				}
				(*out.SeccompProfile).UnmarshalEasyJSON(in)
			}
		case "windowsOptions":
			if in.IsNull() {
				in.Skip()
				out.WindowsOptions = nil
			} else {
				if out.WindowsOptions == nil {
					out.WindowsOptions = new(WindowsSecurityContextOptions)
				}
				easyjson2dce4aeDecodeGithubComKubewardenK8sObjectsApiCoreV11(in, out.WindowsOptions)
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
func easyjson2dce4aeEncodeGithubComKubewardenK8sObjectsApiCoreV1(out *jwriter.Writer, in SecurityContext) {
	out.RawByte('{')
	first := true
	_ = first
	if in.AllowPrivilegeEscalation {
		const prefix string = ",\"allowPrivilegeEscalation\":"
		first = false
		out.RawString(prefix[1:])
		out.Bool(bool(in.AllowPrivilegeEscalation))
	}
	if in.Capabilities != nil {
		const prefix string = ",\"capabilities\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		(*in.Capabilities).MarshalEasyJSON(out)
	}
	if in.Privileged {
		const prefix string = ",\"privileged\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Bool(bool(in.Privileged))
	}
	if in.ProcMount != "" {
		const prefix string = ",\"procMount\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.ProcMount))
	}
	if in.ReadOnlyRootFilesystem {
		const prefix string = ",\"readOnlyRootFilesystem\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Bool(bool(in.ReadOnlyRootFilesystem))
	}
	if in.RunAsGroup != 0 {
		const prefix string = ",\"runAsGroup\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.RunAsGroup))
	}
	if in.RunAsNonRoot {
		const prefix string = ",\"runAsNonRoot\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Bool(bool(in.RunAsNonRoot))
	}
	if in.RunAsUser != 0 {
		const prefix string = ",\"runAsUser\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.RunAsUser))
	}
	if in.SeLinuxOptions != nil {
		const prefix string = ",\"seLinuxOptions\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		(*in.SeLinuxOptions).MarshalEasyJSON(out)
	}
	if in.SeccompProfile != nil {
		const prefix string = ",\"seccompProfile\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		(*in.SeccompProfile).MarshalEasyJSON(out)
	}
	if in.WindowsOptions != nil {
		const prefix string = ",\"windowsOptions\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		easyjson2dce4aeEncodeGithubComKubewardenK8sObjectsApiCoreV11(out, *in.WindowsOptions)
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v SecurityContext) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson2dce4aeEncodeGithubComKubewardenK8sObjectsApiCoreV1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v SecurityContext) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson2dce4aeEncodeGithubComKubewardenK8sObjectsApiCoreV1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *SecurityContext) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson2dce4aeDecodeGithubComKubewardenK8sObjectsApiCoreV1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *SecurityContext) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson2dce4aeDecodeGithubComKubewardenK8sObjectsApiCoreV1(l, v)
}
func easyjson2dce4aeDecodeGithubComKubewardenK8sObjectsApiCoreV11(in *jlexer.Lexer, out *WindowsSecurityContextOptions) {
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
		case "gmsaCredentialSpec":
			out.GmsaCredentialSpec = string(in.String())
		case "gmsaCredentialSpecName":
			out.GmsaCredentialSpecName = string(in.String())
		case "runAsUserName":
			out.RunAsUserName = string(in.String())
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
func easyjson2dce4aeEncodeGithubComKubewardenK8sObjectsApiCoreV11(out *jwriter.Writer, in WindowsSecurityContextOptions) {
	out.RawByte('{')
	first := true
	_ = first
	if in.GmsaCredentialSpec != "" {
		const prefix string = ",\"gmsaCredentialSpec\":"
		first = false
		out.RawString(prefix[1:])
		out.String(string(in.GmsaCredentialSpec))
	}
	if in.GmsaCredentialSpecName != "" {
		const prefix string = ",\"gmsaCredentialSpecName\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.GmsaCredentialSpecName))
	}
	if in.RunAsUserName != "" {
		const prefix string = ",\"runAsUserName\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.RunAsUserName))
	}
	out.RawByte('}')
}
