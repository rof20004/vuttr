package controllers

import (
	"encoding/json"
	"log"

	"github.com/rof20004/vuttr/models"
	"github.com/rof20004/vuttr/repositories"
	"github.com/valyala/fasthttp"
)

// ListTools - return all tools
func ListTools(ctx *fasthttp.RequestCtx) {
	tag := string(ctx.QueryArgs().Peek("tag"))

	if tag != "" {
		filterTools(ctx, tag)
		return
	}

	tools, err := repositories.FindAllTools()
	if err != nil {
		log.Println(err)
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		return
	}

	ctx.SetStatusCode(fasthttp.StatusOK)
	ctx.SetContentType("application/json;charset=utf-8")
	json.NewEncoder(ctx).Encode(tools)
}

func filterTools(ctx *fasthttp.RequestCtx, tag string) error {
	tools, err := repositories.FindToolsByTag(tag)
	if err != nil {
		log.Println(err)
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		return nil
	}

	ctx.SetStatusCode(fasthttp.StatusOK)
	ctx.SetContentType("application/json;charset=utf-8")
	return json.NewEncoder(ctx).Encode(tools)
}

// InsertTool - create tool
func InsertTool(ctx *fasthttp.RequestCtx) {
	var tool models.Tool
	err := json.Unmarshal(ctx.PostBody(), &tool)
	if err != nil {
		log.Println(err)
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}

	newTool, err := repositories.CreateTool(tool)
	if err != nil {
		log.Println(err)
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		return
	}

	ctx.SetStatusCode(fasthttp.StatusOK)
	ctx.SetContentType("application/json;charset=utf-8")
	json.NewEncoder(ctx).Encode(newTool)
}

// RemoveTool - remove tool
func RemoveTool(ctx *fasthttp.RequestCtx) {
	id, ok := ctx.UserValue("toolID").(string)
	if !ok {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}

	err := repositories.DeleteTool(id)
	if err != nil {
		log.Println(err)
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		return
	}

	ctx.SetStatusCode(fasthttp.StatusNoContent)
}
