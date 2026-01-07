package home

import (
	"net/http"
	"time"

	"github.com/adampresley/httphelpers/requests"
	"github.com/adampresley/httphelpers/responses"
	"github.com/adampresley/rendering"
)

func (c HomeController) HomePage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		responses.Text(w, http.StatusNotFound, "Not Found")
		return
	}

	pageName := "pages/home"

	viewData := HomePage{
		BaseViewModel: rendering.BaseViewModel{
			IsHtmx:             requests.IsHtmx(r),
			JavascriptIncludes: []rendering.JavascriptInclude{},
		},
		NumYears: time.Now().Year() - 2000,
	}

	c.renderer.Render(pageName, viewData, w)
}
