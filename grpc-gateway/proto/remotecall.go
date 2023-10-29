package grpcgateway

import (
	"context"
	"fmt"
	"github.com/jhump/protoreflect/grpcreflect"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/reflection/grpc_reflection_v1alpha"
	"net/http"
)

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
	cc, err := dial(ctx, address)
	if err != nil {
		return nil, err
	}
	refClient := grpcreflect.NewClient(refCtx, grpc_reflection_v1alpha.NewServerReflectionClient(cc))
	reflectSource := DescriptorSourceFromServer(ctx, refClient)
	return reflectSource, nil
}

// ListServiceMethod list all service method
func ListServiceMethod(ctx context.Context, address string, headerList ...http.Header) error {
	descSource, err := GetDescSource(ctx, address, headerList...)
	if err != nil {
		return err
	}

	symbols, err := descSource.ListServices()
	if err != nil {
		return err
	}

	for _, s := range symbols {
		if s[0] == '.' {
			s = s[1:]
		}
		dsc, err := descSource.FindSymbol(s)
		if err != nil {
			return err
		}

		txt, err := GetDescriptorText(dsc, descSource)
		if err != nil {
			return err
		}
		fmt.Println(txt)

		desc := ProtoInfoDesc(dsc)
		fmt.Println(desc)
	}

	return nil
}
