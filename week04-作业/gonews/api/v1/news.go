package v1

import (
	"app/internal/models"
	"app/pkg/e"

	"github.com/gin-gonic/gin"
)

func AddNews(c *gin.Context) {
	news := &models.News{}
	if err := c.ShouldBind(news); err != nil {
		e.Err("bind", err)
	}

	if err := news.Save(); err != nil {
		e.Err("添加新闻失败", err)
	}
	e.Ok(c, "添加新闻成功", nil)

}

func GetSection(c *gin.Context) {

	var retSection []models.RetSection
	var err error
	if retSection, err = models.GetSectionList(); err != nil {
		e.Err("获取地区列表失败", err)
	}
	e.NewsOk(c, "获取地区列表成功", "regions", retSection)

}
