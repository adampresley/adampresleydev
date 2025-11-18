package home

import (
	"net/http"
	"time"

	"github.com/adampresley/adamgokit/httphelpers"
	"github.com/adampresley/adamgokit/rendering"
)

func (c HomeController) HomePage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		httphelpers.WriteText(w, http.StatusNotFound, "Not Found")
		return
	}

	pageName := "pages/home"

	viewData := HomePage{
		BaseViewModel: rendering.BaseViewModel{
			IsHtmx:             httphelpers.IsHtmx(r),
			JavascriptIncludes: []rendering.JavascriptInclude{},
		},
		NumYears: time.Now().Year() - 2000,
	}

	c.renderer.Render(pageName, viewData, w)
}
