package reverse

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http/httputil"
	"net/url"
)

// 反向代理本地 反向
func ReverseProxy(c *gin.Context) {
	target := "localhost:8090"

	u := &url.URL{}
	u.Scheme = "http"
	u.Host = target

	proxy := httputil.NewSingleHostReverseProxy(u)
	fmt.Println(c.Request.URL.Path)
	c.Request.URL.Path = "redirect"
	proxy.ServeHTTP(c.Writer, c.Request)
}
