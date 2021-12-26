package v1

import (
	"app/pkg/e"

	"github.com/gin-gonic/gin"
)

func AuthTest(c *gin.Context) {

	e.Ok(c, "通过测试", nil)

}
