package main

// 前面属于引用的包
import (
	// "encoding/base64"
	"flag"
	"fmt"
	// "io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)
//这个是加密的密码，也就是解密PDF文件的密码
// var key = []byte("12580")

// 判断文件是否存在，如果存在返回true，如果不存在返回false
func IsFileExist(filename string) bool {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return false
	}
	return true
}

// 验证页面，返回 PDF 
func Query(c *gin.Context) {
	file := c.Query("file")
	code := c.Query("code")

	filename := "./data/" + code
	
	if file != "validatePZ" || !IsFileExist(filename) {
		c.Status(http.StatusNotFound)
		return
	}

	// bytes, _ := ioutil.ReadFile(filename)
	// if len(key) != 0 {
	// 	bytes = AESDecrypt(bytes, key)
	// }

	// pdf := base64.StdEncoding.EncodeToString(bytes)

	// c.HTML(http.StatusOK, "index.html", gin.H{
	// 	"PDF": pdf,
	// })

	// c.HTML(http.StatusOK, "S2I4DUAOS5.html" , gin.H{
	// 	"html": filename,
	// })

	c.HTML(http.StatusOK, "index.html" , gin.H{
		// "html": filename,
	})
}

func main() {
	var port int

	flag.IntVar(&port, "port", 80, "端口号")

	flag.Parse()

	r := gin.New()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	// r.Static("/static", "./static")
	// r.LoadHTMLFiles("./static/index.html")

	r.Static("/data", "./data")
	r.LoadHTMLFiles("./data/index.html")

	r.GET("/ggfwwt/assets/plugins/pdfjs/web/viewer.html", Query)

	srv := http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: r,
	}
	log.Printf("HTTP Service Started at %s", srv.Addr)
	srv.ListenAndServe()
}
