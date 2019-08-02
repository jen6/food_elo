package app

import (
	//"database/sql"
	//"github.com/coopernurse/gorp"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jen6/letzy/app/models"
	rgorp "github.com/revel/modules/orm/gorp/app"
	"github.com/revel/revel"
	//"github.com/valyala/fasthttp"
	//"gopkg.in/gorp.v2"
)

var (
	// AppVersion revel app version (ldflags)
	AppVersion string

	// BuildTime revel app build-time (ldflags)
	BuildTime string
)

func init() {
	// Filters is the default set of global filters.
	revel.Filters = []revel.Filter{
		revel.PanicFilter,             // Recover from panics and display an error page instead.
		revel.RouterFilter,            // Use the routing table to select the right Action
		revel.FilterConfiguringFilter, // A hook for adding or removing per-Action filters.
		revel.ParamsFilter,            // Parse parameters into Controller.Params.
		revel.SessionFilter,           // Restore and write the session cookie.
		revel.FlashFilter,             // Restore and write the flash cookie.
		revel.ValidationFilter,        // Restore kept validation errors and save new ones from cookie.
		revel.I18nFilter,              // Resolve the requested language
		HeaderFilter,                  // Add some security based headers
		revel.InterceptorFilter,       // Run interceptors around the action.
		revel.CompressFilter,          // Compress the result.
		revel.BeforeAfterFilter,       // Call the before and after filter functions
		revel.ActionInvoker,           // Invoke the action.
	}
	revel.OnAppStart(func() {
		rgorp.Db.SetDbInit(func(dbGorp *rgorp.DbGorp) error {
			t := rgorp.Db.Map.AddTableWithName(models.Food{}, "foods").SetKeys(true, "Idx")
			t.ColMap("Rating").Transient = true

			t = rgorp.Db.Map.AddTableWithName(models.Action{}, "elo_action").SetKeys(true, "Idx")
			t.ColMap("FoodA").Transient = true
			t.ColMap("FoodB").Transient = true
			t.ColMap("IsAWin").Transient = true
			return nil
		})
		Dbm := rgorp.Db.Map

		if Dbm == nil {
			panic("asdf")
		}

		rgorp.Db.TraceOn(revel.AppLog)
	})

}

// HeaderFilter adds common security headers
// There is a full implementation of a CSRF filter in
// https://github.com/revel/modules/tree/master/csrf
var HeaderFilter = func(c *revel.Controller, fc []revel.Filter) {
	c.Response.Out.Header().Add("X-Frame-Options", "SAMEORIGIN")
	c.Response.Out.Header().Add("X-XSS-Protection", "1; mode=block")
	c.Response.Out.Header().Add("X-Content-Type-Options", "nosniff")
	c.Response.Out.Header().Add("Referrer-Policy", "strict-origin-when-cross-origin")

	fc[0](c, fc[1:]) // Execute the next filter stage.
}

//func ExampleStartupScript() {
//	// revel.DevMod and revel.RunMode work here
//	// Use this script to check for dev mode and set dev/prod startup scripts here!
//	if revel.DevMode == true {
//		// Dev mode
//	}
//}
