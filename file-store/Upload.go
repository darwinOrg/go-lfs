package fs

import (
	"errors"
	"github.com/julienschmidt/httprouter"
	"io"
	"lfs/setting"
	"log"
	"mime/multipart"
	"net/http"
	"net/textproto"
	"os"
	"strconv"
)

var ErrMessageTooLarge = errors.New("multipart: message too large")
var errInvalidWrite = errors.New("invalid write result")
var ErrShortWrite = errors.New("short write")
var ErrExceedMax = errors.New("exceed max file")

type MyFileHeader struct {
	Filename string
	Header   textproto.MIMEHeader
	Size     int64

	Part    *multipart.Part
	Tmpfile string
}

type MyForm struct {
	Value map[string][]string
	File  map[string][]*MyFileHeader
}

func Upload(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	file, header, err := req.FormFile("file")
	if err != nil {
		errOutput(w)
		return
	}

	hash := strconv.FormatInt(hashFileKey(header.Filename), 10)
	dir := setting.GetLocalStorePath() + hash
	_ = os.MkdirAll(dir, os.ModePerm)
	filePath := dir + "/" + header.Filename
	b, err := io.ReadAll(file)

	_ = os.WriteFile(filePath, b, os.ModePerm)
	log.Printf("filePath:%s\n", filePath)
	normalOutput(w)
}

func errOutput(w http.ResponseWriter) {
	w.WriteHeader(500)
}

func normalOutput(w http.ResponseWriter) {
	w.WriteHeader(200)
}
