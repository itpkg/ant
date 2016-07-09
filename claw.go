package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func parse(_url string) (string, string, error) {
	uri, err := url.Parse(_url)
	if err != nil {
		return "", "", err
	}
	return uri.Host, uri.Path, nil
}

func claw(_url string) error {
	host, home, err := parse(_url)
	if err != nil {
		return err
	}
	log.Infof("GET %s", _url)
	log.Infof("host=%s, path=%s", host, home)
	doc, err := goquery.NewDocument(_url)
	if err != nil {
		return err
	}

	// if err := os.MkdirAll(path.Join("tmp", "name", 0755)); err != nil {
	// 	return err
	// }
	doc.Find("a").Each(func(i int, s *goquery.Selection) {
		if href, ok := s.Attr("href"); ok {
			log.Infof("FIND %s", href)
			if !strings.HasPrefix(href, "http") {

			}

			hos, hom, err := parse(href)
			if err != nil {
				log.Errorf("bad parse %s: %v", href, err)
				return
			}

			log.Infof("host=%s, path=%s", hos, hom)

			if (host == "" || hos == host) && strings.HasPrefix(hom, home) {
				//claw(href)
				return
			} else {
				log.Warningf("INGNORE")
			}
		}
	})
	return nil
}

func claw1(url string) error {
	res, err := http.Get(url)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	buf, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}
	fmt.Printf("%s", buf)
	return nil
}
