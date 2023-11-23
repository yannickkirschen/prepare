package prepper

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func DownloadFile(filepath string, url string) error {
	r, err := http.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, r.Body)
	return err
}

type Download struct {
	Volume  string   `json:"volume"`
	Cron    string   `json:"cron"`
	Sources []Source `json:"sources"`
}

type Source struct {
	Filename string `json:"file_name"`
	URL      string `json:"url"`
}

func ReadDownloadFile(filename string) (*Download, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	var download Download
	err = json.NewDecoder(file).Decode(&download)
	if err != nil {
		return nil, err
	}

	return &download, nil
}

func (d *Download) Download() error {
	timeNoColons := strings.ReplaceAll(time.Now().Format(time.RFC3339), ":", "-")
	timeNoPlus := strings.ReplaceAll(timeNoColons, "+", "-")

	for _, source := range d.Sources {
		err := source.Download(d.Volume + "/" + timeNoPlus)
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *Source) Download(volume string) error {
	path := volume + "/" + s.Filename
	dir := filepath.Dir(path)

	os.MkdirAll(volume, os.ModePerm)
	os.MkdirAll(dir, os.ModePerm)

	if err := DownloadFile(path, s.URL); err != nil {
		log.Println("Error downloading", s.Filename)
		return err
	}

	log.Println("Downloaded file", s.Filename)
	return nil
}
