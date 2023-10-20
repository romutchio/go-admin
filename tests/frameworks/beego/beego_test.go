package beego

import (
	"net/http"
	"testing"

	"github.com/gavv/httpexpect"
	"github.com/romutchio/go-admin/tests/common"
)

func TestBeego(t *testing.T) {
	common.ExtraTest(httpexpect.WithConfig(httpexpect.Config{
		Client: &http.Client{
			Transport: httpexpect.NewBinder(internalHandler()),
			Jar:       httpexpect.NewJar(),
		},
		Reporter: httpexpect.NewAssertReporter(t),
	}))
}
