package golangCursoUdemyHtml

import (
	"io/ioutil"
	"net/http"
	"regexp"
)

// Titulo recebe um ou mais urls de sites e retorna o título da página
func Titulo(urls ...string) <-chan string {

	c := make(chan string)

	for _, url := range urls {

		go func(url string) {

			resp, _ := http.Get(url) // http.Get(url) returns a *http.Response

			html, _ := ioutil.ReadAll(resp.Body) // retorna um []byte

			r, _ := regexp.Compile("<title>(.*?)<\\/title>") // retorna um *regexp.Regexp

			c <- r.FindStringSubmatch(string(html))[1] // retorna um []string
		}(url)

	}

	return c // retorna um channel de strings (o canal não está sendo encerrado)
}
