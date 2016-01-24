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

func FileUploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.ContentLength > 2*1024*1024 {
		http.Error(w, "Request too large, I'm too tired to be smart right now", http.StatusExpectationFailed)
		return
	}

	file, header, err := r.FormFile("file")
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	reader := io.LimitReader(file, 2*1024*1024)
	wav := wave.New(wave.Stereo, 22000)
	h := hasher.New(reader)
	sng := h.Hash()
	file.Close()

	if len(sng.Notes) == 0 {
		http.Error(w, "The file is too small, at least 17 bytes are required, and yes I'm that lazy!", 400)
		return
	}

	ks := karplus.Song{
		Song:         *sng,
		SamplingRate: 22000,
	}
	ks.Sound(&wav)

	w.Header().Add("Content-Type", "audio/x-wav")

	filename := url.QueryEscape(
		strings.TrimSuffix(header.Filename, filepath.Ext(header.Filename)) + ".wav")

	w.Header().Add("Content-Disposition", "attachment; filename="+filename)
	io.Copy(w, wav.Reader())
}
