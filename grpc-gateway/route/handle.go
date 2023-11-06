package httprouter

import (
	"context"
	"fmt"
	"github.com/LCY2013/grpc-cloud/grpc-gateway/ack"
	grpcgateway "github.com/LCY2013/grpc-cloud/grpc-gateway/proto"
	"github.com/LCY2013/grpc-cloud/logger"
	"net/http"
	"strings"
)

var (
	R = &Route{
		tree: map[string]*node{},
	}
)

type Route struct {
	tree map[string]*node
}

// AddPath add path
func (r *Route) AddPath(path string, annotation *grpcgateway.Annotation) {
	if annotation.Method == "*" {
		for k := range grpcgateway.MethodMap {
			r.addUri(path, &grpcgateway.Annotation{
				Method:  k,
				Address: annotation.Address,
				Symbols: annotation.Symbols,
				Map:     annotation.Map,
			})
		}

		return
	}
	r.addUri(path, annotation)
}

func (r *Route) addUri(path string, annotation *grpcgateway.Annotation) {
	n, ok := R.tree[annotation.Method]
	if !ok {
		R.tree[annotation.Method] = &node{}
		n = R.tree[annotation.Method]
	}
	n.addRoute(path, HandlersChain{func(writer http.ResponseWriter, request *http.Request, params Params) {
		source, err := grpcgateway.GetDescSource(context.Background(), annotation.Address)
		if err != nil {
			_, _ = writer.Write([]byte(err.Error()))
			return
		}

		// if not verbose output, then also include record delimiters
		// between each message, so output could potentially be piped
		// to another gateway process
		options := grpcgateway.FormatOptions{
			EmitJSONDefaultFields: true,
			IncludeTextSeparator:  false,
			AllowUnknownFields:    true,
		}
		rf, formatter, err := grpcgateway.RequestParserAndFormatter(grpcgateway.FormatJSON, source, request.Body, options)
		if err != nil {
			logger.Log.Error(err, "Failed to construct request parser and formatter for %q", grpcgateway.FormatJSON)
		}

		cc, err := grpcgateway.Dial(context.Background(), annotation.Address)
		if err != nil {
			_, _ = writer.Write([]byte(err.Error()))
			return
		}
		h := &grpcgateway.DefaultEventHandler{
			Out:            writer,
			Formatter:      formatter,
			VerbosityLevel: 0, // 0,1,2
		}

		err = grpcgateway.InvokeRPC(context.Background(), source, cc, annotation.Symbols, header(request.Header), h, rf.Next)
		if err != nil {
			_, _ = writer.Write([]byte(ack.ToFailResponse(err.Error())))
			return
		}
	}})
}

func header(deader map[string][]string) (h []string) {
	for k, v := range deader {
		h = append(h, fmt.Sprintf("%s:%s", k, v[0]))
	}

	return
}

func Init() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		n, ok := R.tree[strings.ToUpper(request.Method)]
		if !ok {
			http.NotFound(writer, request)
			return
		}
		handle, ps, tsr := n.getValue(request.RequestURI, func() *Params {
			return &Params{}
		})
		if tsr {
			http.NotFound(writer, request)
			return
		}
		if ps == nil {
			ps = &Params{}
		}
		logger.Log.Info(ps)
		for _, h := range handle {
			h(writer, request, *ps)
		}
	})
}
