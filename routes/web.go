package routes

import (
	"github.com/bluewebtech/harper"
)

func Routes(app *harper.Application) *harper.Application {
	app.Get("/", root)
	app.Get("/about", about)

	return app
}

func root(ctx *harper.Context) {
	ctx.Send("Harper!")
}

func about(ctx *harper.Context) {
	ctx.Send("About Harper!")
}
