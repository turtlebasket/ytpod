package main

import (
	"log"
	"net/http"
	"time"
    "fmt"
    "os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/ksinica/ytpod/pkg/youtube"
)

type transport struct{}

func (*transport) RoundTrip(r *http.Request) (*http.Response, error) {
	r.Header.Set(
		"User-Agent",
		"github.com/ksinica/ytpod",
	)
	return http.DefaultTransport.RoundTrip(r)
}

func main() {
	c := http.Client{
		Transport: new(transport),
	}

	r := chi.NewMux()
	r.Use(youtube.UseHTTPClient(&c))
	r.Use(middleware.Timeout(time.Second * 60))

	r.Route("/youtube", youtube.Router)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Content-Type", "text/html")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(fmt.Sprintf(`
        <!DOCTYPE html>
        <html>
            <head><title>ytpod</title></head>
            <body>
                <h1>fork of <a href="https://github.com/ksinica/ytpod">ytpod</a></h1>
                <p>Converts a YT channel into a podcatcher-compatible RSS feed.</p>
                <hr>
                <p>Supported feed url formats:</p>
                <ul>
                    <li>%[1]s/youtube/feed/@USERNAME</li>
                    <li>%[1]s/youtube/feed/user/USERNAME</li>
                    <li>%[1]s/youtube/feed/channel/CHANNEL_ID</li>
                </ul>
            </body>
        </html>
        `, os.Getenv("YTPOD_URL"))))
	})

	log.Fatal(http.ListenAndServe(":8080", r))
}
