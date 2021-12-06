package automation

// This file is auto-generated.
//
// Changes to this file may cause incorrect behavior and will be lost if
// the code is regenerated.
//
// Definitions file that controls how this file is generated:
// automation/automation/jwt_handler.yaml

import (
	"context"
	atypes "github.com/cortezaproject/corteza-server/automation/types"
	"github.com/cortezaproject/corteza-server/pkg/expr"
	"github.com/cortezaproject/corteza-server/pkg/wfexec"
	"io"
)

var _ wfexec.ExecResponse

type (
	jwtHandlerRegistry interface {
		AddFunctions(ff ...*atypes.Function)
		Type(ref string) expr.Type
	}
)

func (h jwtHandler) register() {
	h.reg.AddFunctions(
		h.Generate(),
	)
}

type (
	jwtGenerateArgs struct {
		hasScp    bool
		Scp       interface{}
		scpString string
		scpArray  []expr.TypedValue

		hasHeader    bool
		Header       interface{}
		headerKV     map[string]string
		headerString string
		headerStream io.Reader

		hasPayload    bool
		Payload       interface{}
		payloadKV     map[string]string
		payloadString string
		payloadStream io.Reader

		hasSecret    bool
		Secret       interface{}
		secretString string
		secretStream io.Reader
	}

	jwtGenerateResults struct {
		Token string
	}
)

func (a jwtGenerateArgs) GetScp() (bool, string, []expr.TypedValue) {
	return a.hasScp, a.scpString, a.scpArray
}

func (a jwtGenerateArgs) GetHeader() (bool, map[string]string, string, io.Reader) {
	return a.hasHeader, a.headerKV, a.headerString, a.headerStream
}

func (a jwtGenerateArgs) GetPayload() (bool, map[string]string, string, io.Reader) {
	return a.hasPayload, a.payloadKV, a.payloadString, a.payloadStream
}

func (a jwtGenerateArgs) GetSecret() (bool, string, io.Reader) {
	return a.hasSecret, a.secretString, a.secretStream
}

// Generate function Generate JWT
//
// expects implementation of generate function:
// func (h jwtHandler) generate(ctx context.Context, args *jwtGenerateArgs) (results *jwtGenerateResults, err error) {
//    return
// }
func (h jwtHandler) Generate() *atypes.Function {
	return &atypes.Function{
		Ref:    "jwtGenerate",
		Kind:   "function",
		Labels: map[string]string(nil),
		Meta: &atypes.FunctionMeta{
			Short: "Generate JWT",
		},

		Parameters: []*atypes.Param{
			{
				Name:  "scp",
				Types: []string{"String", "Array"},
			},
			{
				Name:  "header",
				Types: []string{"KV", "String", "Reader"}, Required: true,
			},
			{
				Name:  "payload",
				Types: []string{"KV", "String", "Reader"}, Required: true,
			},
			{
				Name:  "secret",
				Types: []string{"String", "Reader"}, Required: true,
			},
		},

		Results: []*atypes.Param{

			{
				Name:  "token",
				Types: []string{"String"},
			},
		},

		Handler: func(ctx context.Context, in *expr.Vars) (out *expr.Vars, err error) {
			var (
				args = &jwtGenerateArgs{
					hasScp:     in.Has("scp"),
					hasHeader:  in.Has("header"),
					hasPayload: in.Has("payload"),
					hasSecret:  in.Has("secret"),
				}
			)

			if err = in.Decode(args); err != nil {
				return
			}

			// Converting Scp argument
			if args.hasScp {
				aux := expr.Must(expr.Select(in, "scp"))
				switch aux.Type() {
				case h.reg.Type("String").Type():
					args.scpString = aux.Get().(string)
				case h.reg.Type("Array").Type():
					args.scpArray = aux.Get().([]expr.TypedValue)
				}
			}

			// Converting Header argument
			if args.hasHeader {
				aux := expr.Must(expr.Select(in, "header"))
				switch aux.Type() {
				case h.reg.Type("KV").Type():
					args.headerKV = aux.Get().(map[string]string)
				case h.reg.Type("String").Type():
					args.headerString = aux.Get().(string)
				case h.reg.Type("Reader").Type():
					args.headerStream = aux.Get().(io.Reader)
				}
			}

			// Converting Payload argument
			if args.hasPayload {
				aux := expr.Must(expr.Select(in, "payload"))
				switch aux.Type() {
				case h.reg.Type("KV").Type():
					args.payloadKV = aux.Get().(map[string]string)
				case h.reg.Type("String").Type():
					args.payloadString = aux.Get().(string)
				case h.reg.Type("Reader").Type():
					args.payloadStream = aux.Get().(io.Reader)
				}
			}

			// Converting Secret argument
			if args.hasSecret {
				aux := expr.Must(expr.Select(in, "secret"))
				switch aux.Type() {
				case h.reg.Type("String").Type():
					args.secretString = aux.Get().(string)
				case h.reg.Type("Reader").Type():
					args.secretStream = aux.Get().(io.Reader)
				}
			}

			var results *jwtGenerateResults
			if results, err = h.generate(ctx, args); err != nil {
				return
			}

			out = &expr.Vars{}

			{
				// converting results.Token (string) to String
				var (
					tval expr.TypedValue
				)

				if tval, err = h.reg.Type("String").Cast(results.Token); err != nil {
					return
				} else if err = expr.Assign(out, "token", tval); err != nil {
					return
				}
			}

			return
		},
	}
}
