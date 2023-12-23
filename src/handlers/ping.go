package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tatsuya06068/go-gin-redis/drivers"
)

// @BasePath /api/v1

// PingExample godoc
// @Summary liveness probe
// @Schemes
// @Description do ping
// @Tags Ping
// @Accept json
// @Param  test    query     string  true  "QueryParams"
// @Produce json
// @Success 200 {string} Helloworld
// @Failure      500  {object}  httputil.HTTPError
// @Router /ping [get]
func Ping(g *gin.Context) {
	// ua := g.Query("test")
	driver := drivers.NewRedisDriver()
	result, err:= driver.Ping(g)

	if err != nil {
		g.JSON(http.StatusOK, err)	
	}

	g.JSON(http.StatusOK, result)
}
