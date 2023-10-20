package gorilla

import (
	// add gorilla adapter
	_ "github.com/romutchio/go-admin/adapter/gorilla"
	"github.com/romutchio/go-admin/modules/config"
	"github.com/romutchio/go-admin/modules/language"
	"github.com/romutchio/go-admin/plugins/admin/modules/table"

	// add mysql driver
	_ "github.com/romutchio/go-admin/modules/db/drivers/mysql"
	// add postgresql driver
	_ "github.com/romutchio/go-admin/modules/db/drivers/postgres"
	// add sqlite driver
	_ "github.com/romutchio/go-admin/modules/db/drivers/sqlite"
	// add mssql driver
	_ "github.com/romutchio/go-admin/modules/db/drivers/mssql"
	// add adminlte ui theme
	"github.com/GoAdminGroup/themes/adminlte"

	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/romutchio/go-admin/engine"
	"github.com/romutchio/go-admin/plugins/admin"
	"github.com/romutchio/go-admin/plugins/example"
	"github.com/romutchio/go-admin/template"
	"github.com/romutchio/go-admin/template/chartjs"
	"github.com/romutchio/go-admin/tests/tables"
)

func internalHandler() http.Handler {
	app := mux.NewRouter()
	eng := engine.Default()

	examplePlugin := example.NewExample()
	template.AddComp(chartjs.NewChart())

	if err := eng.AddConfigFromJSON(os.Args[len(os.Args)-1]).
		AddPlugins(admin.NewAdmin(tables.Generators).
			AddGenerator("user", tables.GetUserTable), examplePlugin).
		Use(app); err != nil {
		panic(err)
	}

	eng.HTML("GET", "/admin", tables.GetContent)

	return app
}

func NewHandler(dbs config.DatabaseList, gens table.GeneratorList) http.Handler {
	app := mux.NewRouter()
	eng := engine.Default()

	template.AddComp(chartjs.NewChart())

	if err := eng.AddConfig(&config.Config{
		Databases: dbs,
		UrlPrefix: "admin",
		Store: config.Store{
			Path:   "./uploads",
			Prefix: "uploads",
		},
		Language:    language.EN,
		IndexUrl:    "/",
		Debug:       true,
		ColorScheme: adminlte.ColorschemeSkinBlack,
	}).
		AddPlugins(admin.NewAdmin(gens)).
		Use(app); err != nil {
		panic(err)
	}

	eng.HTML("GET", "/admin", tables.GetContent)

	return app
}
