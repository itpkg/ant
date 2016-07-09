package minify

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/PuerkitoBio/goquery"
	_minify "github.com/tdewolff/minify"
	"github.com/tdewolff/minify/html"
	"github.com/urfave/cli"
)

//Command minify
var Command = cli.Command{

	Name:    "minify",
	Aliases: []string{"m"},
	Usage:   "minify html",
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:  "src, s",
			Usage: "source folder",
		},
		// cli.BoolFlag{
		// 	Name:  "css, c",
		// 	Usage: "keep css block",
		// },
		// cli.BoolFlag{
		// 	Name:  "js, j",
		// 	Usage: "keep javascript block",
		// },
		cli.StringSliceFlag{
			Name:  "ext, e",
			Value: &cli.StringSlice{".html", ".xhtml"},
			Usage: "file extensions",
		},
	},
	Action: func(c *cli.Context) error {
		src := c.String("src")
		if len(src) == 0 {
			cli.ShowCommandHelp(c, "minify")
			return nil
		}

		m := _minify.New()
		// if !c.Bool("css") {
		// 	m.AddFunc("text/css", css.Minify)
		// }
		// if !c.Bool("js") {
		// 	m.AddFunc("text/javascript", js.Minify)
		// }
		m.AddFunc("text/html", html.Minify)

		return filepath.Walk(src, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if info.IsDir() {
				return nil
			}
			for _, f := range c.StringSlice("ext") {
				if filepath.Ext(info.Name()) == f {
					return min(m, src, path)
				}
			}
			return nil
		})
	},
}

func min(m *_minify.M, s, p string) error {

	// m.AddFunc("image/svg+xml", svg.Minify)
	// m.AddFuncRegexp(regexp.MustCompile("[/+]json$"), json.Minify)
	// m.AddFuncRegexp(regexp.MustCompile("[/+]xml$"), xml.Minify)

	fnt := s + "-min/" + p[len(s)+1:]
	fmt.Printf("%s => %s\n", p, fnt)

	fds, err := os.Open(p)
	if err != nil {
		return err
	}
	defer fds.Close()
	// var h HTML
	// dec := xml.NewDecoder(fds)
	// if err = dec.Decode(&h); err != nil {
	// 	return err
	// }
	//fmt.Printf("%+v\n", h)
	doc, err := goquery.NewDocumentFromReader(fds)
	if err != nil {
		return err
	}

	if err = os.MkdirAll(filepath.Dir(fnt), 0755); err != nil {
		return err
	}
	fdt, err := os.OpenFile(fnt, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer fdt.Close()

	body, err := doc.Find("body").Html()
	if err != nil {
		return err
	}
	body, err = m.String("text/html", body)
	if err != nil {
		return err
	}
	fdt.WriteString("<html><head><title>")
	fdt.WriteString(doc.Find("head").Find("title").Text())
	fdt.WriteString("</title></head><body>")
	fdt.WriteString(body)
	return nil
	// enc := xml.NewEncoder(fdt)
	// return enc.Encode(&h)
}
