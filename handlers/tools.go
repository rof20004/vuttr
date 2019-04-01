package handlers

import (
	"github.com/buaazp/fasthttprouter"
	"github.com/rof20004/vuttr/auth"
	"github.com/rof20004/vuttr/controllers"
)

// InitializeToolsHandlers - start tools handlers
func InitializeToolsHandlers(router *fasthttprouter.Router) {
	var v1 = "/tools"
	router.GET(v1, controllers.ListTools)
	router.POST(v1, controllers.InsertTool)
	router.DELETE(v1+"/:toolID", auth.BasicAuth(controllers.RemoveTool, "admin", "bossabox"))
}
