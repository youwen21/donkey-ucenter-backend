package serverx

import (
	"errors"
	"io/fs"
	"net/http"
	"os"
	"strings"
)

//http.StripPrefix("/admin/", http.FileServer(http.FS(sub)))

type StaticFsServer struct {
	Fs fs.FS

	Files map[string]bool

	fileServer http.Handler
}

func (s StaticFsServer) CanServe(urlPath string) bool {
	upath := strings.TrimLeft(urlPath, "/")
	if _, ok := s.Files[upath]; ok {
		return true
	}

	return false
}

func (s StaticFsServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// css, js, img等文件
	// 需要输出正确的header头，所以借助http.FileServer
	s.fileServer.ServeHTTP(w, r)
}

func NewStaticEmbedServer(distFs fs.FS) *StaticFsServer {
	if distFs == nil {
		panic(errors.New("init NewEmbedServer fail"))
	}

	// 不存在本地 ./asset/static/目录
	return &StaticFsServer{
		distFs,
		initPathMap(distFs),
		http.FileServer(http.FS(distFs)),
	}
}

func NewStaticLocalServer(localDir string) *StaticFsServer {
	_, err := os.Stat(localDir)
	if err != nil {
		panic(errors.New("init NewLocalServer fail:" + localDir))
	}

	return &StaticFsServer{
		os.DirFS(localDir),
		initPathMap(os.DirFS(localDir)),
		http.FileServer(http.Dir(localDir)),
	}
}

func initPathMap(distFs fs.FS) map[string]bool {
	pathMap := make(map[string]bool, 0)

	// 遍历 distFs 所有文件, 将文件路径添加到 hsRouters 中
	fs.WalkDir(distFs, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			return nil
		}
		pathMap[path] = true
		return nil
	})

	return pathMap
}

func NewStaticLocalOrEmbedServer(localDir string, distFs fs.FS) *StaticFsServer {
	if localDir != "" {
		if _, err := os.Stat(localDir); err == nil {
			return NewStaticLocalServer(localDir)
		}
	}

	if distFs != nil {
		return NewStaticEmbedServer(distFs)
	}

	panic(errors.New("init NewLocalOrEmbedServer fail"))
}
