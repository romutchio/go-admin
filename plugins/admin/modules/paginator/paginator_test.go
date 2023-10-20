package paginator

import (
	"testing"

	_ "github.com/GoAdminGroup/themes/sword"

	"github.com/romutchio/go-admin/modules/config"
	"github.com/romutchio/go-admin/plugins/admin/modules/parameter"
)

func TestGet(t *testing.T) {
	config.Initialize(&config.Config{Theme: "sword"})
	param := parameter.BaseParam()
	Get(Config{
		Size:         105,
		Param:        param.SetPage("7"),
		PageSizeList: []string{"10", "20", "50", "100"},
	})
}
