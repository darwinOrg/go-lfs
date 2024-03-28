package fs

import (
	"github.com/julienschmidt/httprouter"
	"io"
	"lfs/setting"
	"log"
	"net/http"
	"os"
	"strconv"
)

func Download(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	queryValues := req.URL.Query()
	fileKey := queryValues["fileKey"][0]
	log.Printf("dowload fileKey:%s\n", fileKey)
	hash := strconv.FormatInt(hashFileKey(fileKey), 10)
	downLoadFile := setting.GetLocalStorePath() + hash + "/" + fileKey

	file, err := os.Open(downLoadFile)
	if err != nil {
		errOutput(w)
		return
	}
	defer file.Close()

	w.Header().Set("Content-Type", "application/octet-stream")
	w.Header().Set("Content-Disposition", "attachment; filename="+fileKey)

	io.Copy(w, file)
}
