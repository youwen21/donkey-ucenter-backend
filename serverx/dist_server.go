package serverx

import (
	"errors"
	"io"
	"io/fs"
	"net/http"
	"os"
	"strings"
)

//http.StripPrefix("/admin/", http.FileServer(http.FS(sub)))

type DistFsServer struct {
	Fs fs.FS

	fileServer http.Handler
}

func (s DistFsServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	upath := strings.TrimLeft(r.URL.Path, "/")
	ck, err := s.Fs.Open(upath)
	if err == nil {
		defer ck.Close()
		// css, js, img等文件
		// 需要输出正确的header头，所以借助http.FileServer
		s.fileServer.ServeHTTP(w, r)
		return
	}

	// 前端维护 404路由
	fi, _ := s.Fs.Open("index.html")
	defer fi.Close()
	content, _ := io.ReadAll(fi)
	w.Write(content)
}

func NewDistEmbedServer(distFs fs.FS) *DistFsServer {
	if distFs == nil {
		panic(errors.New("init NewEmbedServer fail"))
	}

	// 不存在本地 ./asset/static/目录
	return &DistFsServer{
		distFs,
		http.FileServer(http.FS(distFs)),
	}
}

func NewDistLocalServer(localDir string) *DistFsServer {
	_, err := os.Stat(localDir)
	if err != nil {
		panic(errors.New("init NewLocalServer fail:" + localDir))
	}

	return &DistFsServer{
		os.DirFS(localDir),
		http.FileServer(http.Dir(localDir)),
	}
}

func NewDistLocalOrEmbedServer(localDir string, distFs fs.FS) *DistFsServer {
	if localDir != "" {
		if _, err := os.Stat(localDir); err == nil {
			return NewDistLocalServer(localDir)
		}
	}

	if distFs != nil {
		return NewDistEmbedServer(distFs)
	}

	panic(errors.New("init NewLocalOrEmbedServer fail"))
}
