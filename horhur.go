package horhur

import (
	"net/url"
	"time"

	"github.com/playwright-community/playwright-go"
)

type RedirectInfo struct {
	NewURL       string
	IsSameDomain bool
}

func DoesItRedirect(toLoad string) (*RedirectInfo, bool, error) {
	pw, err := playwright.Run()
	if err != nil {
		return nil, false, err
	}

	browser, err := pw.Chromium.Launch()
	if err != nil {
		return nil, false, err
	}
	page, err := browser.NewPage()
	if err != nil {
		return nil, false, err
	}

	if _, err = page.Goto(toLoad); err != nil {
		return nil, false, err
	}

	// DIDNTDO(ttacon): parameterize this to take a different timeout, or only trigger
	// this when we detect a meta-refresh value that is set.
	<-time.After(time.Second * 2)

	if toLoad == page.URL() {
		return nil, false, nil
	}

	originalURL, _ := url.Parse(toLoad)
	newURL, _ := url.Parse(page.URL())

	return &RedirectInfo{
		NewURL:       page.URL(),
		IsSameDomain: originalURL.Hostname() == newURL.Hostname(),
	}, true, nil
}
