package grpcgateway

import (
	"context"
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/jhump/protoreflect/grpcreflect"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/reflection/grpc_reflection_v1alpha"
	"net/http"
	"strings"
)

var (
	ignoreService = map[string]struct{}{
		"grpc.reflection.v1.ServerReflection":      {},
		"grpc.reflection.v1alpha.ServerReflection": {},
	}
	methodMap = map[string]string{
		http.MethodGet:     http.MethodGet,
		http.MethodHead:    http.MethodHead,
		http.MethodPost:    http.MethodPost,
		http.MethodPut:     http.MethodPut,
		http.MethodPatch:   http.MethodPatch,
		http.MethodDelete:  http.MethodDelete,
		http.MethodConnect: http.MethodConnect,
		http.MethodOptions: http.MethodOptions,
		http.MethodTrace:   http.MethodTrace,
	}
)

type Annotation struct {
	Method  string            `json:"method,omitempty"`
	Address string            `json:"address,omitempty"`
	Symbols string            `json:"symbols,omitempty"`
	Map     map[string]string `json:"map,omitempty"`
}

// GetDescSource is a function that returns a string describing the source of a
func GetDescSource(ctx context.Context, address string, headerList ...http.Header) (DescriptorSource, error) {
	var headers []string
	for _, header := range headerList {
		for k, v := range header {
			headers = append(headers, k+":"+v[0])
		}
	}
	md := MetadataFromHeaders(headers)
	refCtx := metadata.NewOutgoingContext(ctx, md)
	cc, err := Dial(ctx, address)
	if err != nil {
		return nil, err
	}
	refClient := grpcreflect.NewClient(refCtx, grpc_reflection_v1alpha.NewServerReflectionClient(cc))
	reflectSource := DescriptorSourceFromServer(ctx, refClient)
	return reflectSource, nil
}

// ListServiceMethod list all service method
func ListServiceMethod(ctx context.Context, address string, headerList ...http.Header) (annotation map[string]*Annotation, err error) {
	descSource, err := GetDescSource(ctx, address, headerList...)
	if err != nil {
		return nil, err
	}

	annotation = map[string]*Annotation{}

	symbols, err := descSource.ListServices()
	if err != nil {
		return nil, err
	}

	for _, s := range symbols {
		if _, ok := ignoreService[s]; ok {
			continue
		}
		if s[0] == '.' {
			s = s[1:]
		}
		dsc, err := descSource.FindSymbol(s)
		if err != nil {
			return nil, err
		}

		desc := ProtoInfoDesc(dsc)
		for method, options := range desc.MethodOptions {
			for _, op := range options {
				optVal, ok := op.val.(proto.Message)
				if !ok {
					continue
				}
				anno := proto.CompactTextString(optVal)

				annota := &Annotation{
					Address: address,
					Map:     make(map[string]string),
					Symbols: fmt.Sprintf("%s/%s", s, method.GetName()),
				}
				kvs := strings.Split(anno, " ")
				for _, a := range kvs {
					kv := strings.Split(a, ":")
					annota.Map[kv[0]] = kv[1]
					if _, ok := methodMap[strings.ToUpper(kv[0])]; ok {
						annota.Method = methodMap[strings.ToUpper(kv[0])]
						annotation[kv[1]] = annota
					}
				}
			}
		}
	}

	return
}
