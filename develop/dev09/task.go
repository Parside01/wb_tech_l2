package main

import (
	"flag"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
)

/*
=== Утилита wget ===

Реализовать утилиту wget с возможностью скачивать сайты целиком

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

type WgetConfig struct {
	URL string
	Dir string
}

var config = &WgetConfig{}

// Суть простая, указываем папку(опционально) и url.
// Программа просто скачивает html и все css, js, img штуки.
// => для каких-т сложных сайтов не будет работать. (wget для них тоже не работает как надо)

func main() {
	flag.Usage = func() {
		log.Fatal("Usage: wget [url]")
	}
	flag.StringVar(&config.Dir, "p", "wget_download", "Path to download files")

	flag.Parse()
	if flag.NArg() != 1 {
		flag.Usage()
	}

	url := strings.TrimSpace(flag.Arg(0))
	if url == "" {
		flag.Usage()
	}
	config.URL = url
	fmt.Println(url)

	if err := os.Mkdir(config.Dir, os.ModePerm); err != nil {
		log.Fatal(err)
	}

	err := DownloadFile(config.URL, filepath.Join(config.Dir, "index.html"))
	if err != nil {
		log.Fatal(err)
	}

	html, err := os.Open(filepath.Join(config.Dir, "index.html"))
	if err != nil {
		log.Fatal(err)
	}

	if err = DownloadSources(html); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Site downloaded to %s", config.URL)
}

func DownloadSources(site io.Reader) error {
	doc, err := goquery.NewDocumentFromReader(site)
	if err != nil {
		return err
	}

	doc.Find("link[rel='stylesheet']").Each(func(i int, s *goquery.Selection) {
		href, ok := s.Attr("href")
		if ok {
			if !strings.HasPrefix(href, "http") {
				href, _ = url.JoinPath(config.URL, href)
			}
		}
		name := filepath.Base(href)
		err = DownloadFile(href, filepath.Join(config.Dir, name))
		if err != nil {
			return
		}
	})

	doc.Find("script[src]").Each(func(i int, s *goquery.Selection) {
		src, ok := s.Attr("src")
		if ok {
			if !strings.HasPrefix(src, "http") {
				src, _ = url.JoinPath(config.URL, src)
			}
		}
		name := filepath.Base(src)
		err = DownloadFile(src, filepath.Join(config.Dir, name))
		if err != nil {
			return
		}
	})

	doc.Find("img[src]").Each(func(i int, s *goquery.Selection) {
		src, ok := s.Attr("src")
		if ok {
			if !strings.HasPrefix(src, "http") {
				src, _ = url.JoinPath(config.URL, src)
			}
		}
		name := filepath.Base(src)
		err = DownloadFile(src, filepath.Join(config.Dir, name))
		if err != nil {
			return
		}
	})
	return nil
}

func DownloadFile(url string, filename string) error {
	resp, err := Get(url)
	if err != nil {
		return err
	}

	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.Copy(file, resp.Body)
	return err
}

func Get(url string) (*http.Response, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("bad status: %s", resp.Status)
	}
	return resp, nil
}
