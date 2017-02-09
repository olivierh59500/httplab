package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gchaincl/httplab"
)

func NewHandler(ui *httplab.UI) http.Handler {
	fn := func(w http.ResponseWriter, req *http.Request) {
		ui.Display(
			httplab.RequestView,
			httplab.DumpRequest(req),
		)

		resp := ui.Response()
		time.Sleep(resp.Delay)
		resp.Write(w)

	}
	return http.HandlerFunc(fn)
}

func main() {
	ui, err := httplab.NewUI()
	if err != nil {
		log.Fatalln(err)
	}

	go func() {
		http.Handle("/", NewHandler(ui))
		ui.Display(httplab.InfoView, []byte("Listening on :8000"))
		if err := http.ListenAndServe(":8000", nil); err != nil {
			log.Fatalln(err)
		}
	}()

	if err := ui.Loop(); err != nil {
		log.Fatalln(err)
	}

}
