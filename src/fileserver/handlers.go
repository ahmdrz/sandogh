package fileserver

import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func getMD5Hash(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}

func (s *Server) getHandler(ctx *gin.Context) {
	filePath := s.getFilePath(ctx)

	checksum := getMD5Hash(filePath)
	eTagKey := `"` + checksum + `"`
	ctx.Header("Cache-Control", "max-age=1209600")
	ctx.Header("Etag", eTagKey)
	if match := ctx.GetHeader("If-None-Match"); match != "" {
		if strings.Contains(match, eTagKey) {
			ctx.Writer.WriteHeader(http.StatusNotModified)
			return
		}
	}

	bytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Printf("Error on %s: %v", filePath, err)
		ctx.String(http.StatusInternalServerError, "Something went wrong !")
		return
	}

	contentType := http.DetectContentType(bytes)
	ctx.Header("Content-Type", contentType)
	ctx.Header("Content-Length", strconv.Itoa(len(bytes)))
	ctx.Writer.Write(bytes)
}

func directoryExists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}

func (s *Server) postHandler(ctx *gin.Context) {
	p := s.getFilePath(ctx)

	file, _, err := ctx.Request.FormFile("file")
	if err != nil {
		log.Printf("Error on %s: %v", p, err)
		ctx.String(http.StatusInternalServerError, "Something went wrong !")
		return
	}
	out, err := os.Create(p)
	if err != nil {
		log.Printf("Error on %s: %v", p, err)
		ctx.String(http.StatusInternalServerError, "Something went wrong !")
		return
	}
	defer out.Close()
	_, err = io.Copy(out, file)
	if err != nil {
		log.Printf("Error on %s: %v", p, err)
		ctx.String(http.StatusInternalServerError, "Something went wrong !")
		return
	}
}

func (s *Server) deleteHandler(ctx *gin.Context) {
	filePath := s.getFilePath(ctx)
	err := os.Remove(filePath)
	if err != nil {
		log.Printf("Error on %s: %v", filePath, err)
		ctx.String(http.StatusInternalServerError, "Something went wrong !")
		return
	}

	ctx.String(http.StatusOK, "File removed !")
}

func (s *Server) getFilePath(ctx *gin.Context) string {
	return path.Join(s.config.BaseDirectory, ctx.Param("directory"), ctx.Param("name"))
}

func (s *Server) healthCheckHandler(ctx *gin.Context) {
	ctx.String(http.StatusOK, "Health: ok")
}
