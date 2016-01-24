package api

import (
	"io"
	"net/http"
	"net/url"
	"path/filepath"
	"strings"

	"github.com/gophergala2016/chopher/hasher"
	"github.com/gophergala2016/chopher/karplus"
	"github.com/gophergala2016/chopher/wave"
)

func Index(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}

func FileUploadHandler(w http.ResponseWriter, r *http.Request) {
	file, header, err := r.FormFile("file")
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	reader := io.LimitReader(file, 1024*1024)
	wav := wave.New(wave.Stereo, 44000)
	h := hasher.New(reader)
	sng := h.Hash()
	file.Close()

	ks := karplus.Song{
		Song:         *sng,
		SamplingRate: 44000,
	}
	ks.Sound(&wav)

	w.Header().Add("Content-Type", "audio/x-wav")

	filename := url.QueryEscape(
		strings.TrimSuffix(header.Filename, filepath.Ext(header.Filename)) + ".wav")

	w.Header().Add("Content-Disposition", "filename="+filename)
	io.Copy(w, wav.Reader())
}
