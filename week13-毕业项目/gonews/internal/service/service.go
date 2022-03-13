package service

import (
	pb "gonews/api/gonews/v1"
	"gonews/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewBlogService)

type GonewsServer struct {
	pb.UnimplementedGonewsServerServer

	log *log.Helper

	article *biz.ArticleUsecase
}
