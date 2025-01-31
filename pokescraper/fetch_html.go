package pokescraper

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"

	"golang.org/x/net/html"
)

// returns a string of the given page's html response
func FetchHTML(url string) (string, error) {
	resp, reqErr := http.Get(url)
	if reqErr != nil {
		return "", fmt.Errorf("unable to retrieve HTML response: %v", reqErr)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("unexpected response status: %v", resp.StatusCode)
	}

	htmlBuilder := new(strings.Builder)
	_, ioErr := io.Copy(htmlBuilder, resp.Body)
	if ioErr != nil {
		return "", fmt.Errorf("unexpected I/O copy error: %v", ioErr)
	}

	return htmlBuilder.String(), nil
}

// returns pointer to html Node struct of passed html content
func ParseHTML(htmlContent string) (*html.Node, error) {
	if htmlContent == "" || htmlContent == " " {
		return nil, errors.New("empty string passed")
	}

	document, err := html.Parse(strings.NewReader(htmlContent))
	if err != nil {
		return nil, fmt.Errorf("unable to parse HTML from passed string:\n%v", err)
	}

	return document, nil
}
