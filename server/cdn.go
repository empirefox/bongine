package server

import (
	"github.com/golang/glog"
	"github.com/kataras/iris"
	"github.com/uber-go/zap"
)

func (s *Server) GetUptoken(ctx *iris.Context) {
	ctx.JSON(iris.StatusOK, iris.Map{
		"Uptoken": s.qiniu.Uptoken(),
	})
}

func (s *Server) PostList(ctx *iris.Context) {
	var data struct{ Prefix string }
	if err := ctx.ReadJSON(&data); err != nil {
		ctx.EmitError(iris.StatusBadRequest)
		return
	}

	items, err := s.qiniu.List(data.Prefix)
	if err != nil {
		s.log.Error("List from qiniu failed", zap.Error(err))
		ctx.EmitError(iris.StatusInternalServerError)
		return
	}

	ctx.JSON(iris.StatusOK, items)
}

func (s *Server) PostDelete(ctx *iris.Context) {
	var data struct{ Key string }
	if err := ctx.ReadJSON(&data); err != nil {
		ctx.EmitError(iris.StatusBadRequest)
		return
	}

	err := s.qiniu.Delete(data.Key)
	if err != nil {
		glog.Errorln(err)
		s.log.Error("Delete from qiniu failed", zap.Error(err))
		ctx.EmitError(iris.StatusInternalServerError)
		return
	}
}
