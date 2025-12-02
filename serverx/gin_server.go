package serverx

import (
	"donkey-ucenter/asset"
	"donkey-ucenter/router"
	"github.com/gin-gonic/gin"
	"io/fs"
	"net/http"
)

var (
	GinEngine *gin.Engine

	StaticServer *StaticFsServer
	DistServer   *DistFsServer
)

func init() {
	// dist file server, vue
	distFS, _ := fs.Sub(asset.Dist, "dist")
	DistServer = NewDistLocalOrEmbedServer("./asset/dist/", distFS)

	// dist file server, 普通静态站内容
	//staticFS, _ := fs.Sub(asset.Static, "static")
	//StaticServer = NewStaticLocalOrEmbedServer("./asset/static/", staticFS)

	// gin http api server
	GinEngine = gin.Default()
	router.InitRouter(GinEngine)

	GinEngine.NoRoute(func(ginCtx *gin.Context) {
		//upath := strings.TrimLeft(ginContext.Request.URL.Path, "/")
		if StaticServer != nil && StaticServer.CanServe(ginCtx.Request.URL.Path) {
			StaticServer.ServeHTTP(ginCtx.Writer, ginCtx.Request)
			return
		}

		// Dist前端存在时，NotFound 由前端项目维护
		if DistServer != nil {
			DistServer.ServeHTTP(ginCtx.Writer, ginCtx.Request)
			return
		}

		http.NotFound(ginCtx.Writer, ginCtx.Request)
	})
}
