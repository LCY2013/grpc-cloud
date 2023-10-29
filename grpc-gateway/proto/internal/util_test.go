package internal_test

import (
	"github.com/LCY2013/grpc-cloud/grpc-gateway/proto/internal"
	testutil "github.com/LCY2013/grpc-cloud/grpc-gateway/utils"
	"testing"
)

func TestCreatePrefixList(t *testing.T) {
	list := internal.CreatePrefixList("")
	testutil.Eq(t, []string{""}, list)

	list = internal.CreatePrefixList("pkg")
	testutil.Eq(t, []string{"pkg", ""}, list)

	list = internal.CreatePrefixList("fully.qualified.pkg.name")
	testutil.Eq(t, []string{"fully.qualified.pkg.name", "fully.qualified.pkg", "fully.qualified", "fully", ""}, list)
}
