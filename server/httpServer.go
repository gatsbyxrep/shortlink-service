package server

import (
	"context"
	"errors"
	"io"
	"log"
	"main/links"
	"net/http"
)

type HttpServer struct {
	httpServer   *http.Server
	linksService links.LinksServiceServer
	errCh        chan error
}

func NewHttpServer(linksServ links.LinksServiceServer, port string) HttpServer {
	router := http.NewServeMux()
	httpServer := HttpServer{
		httpServer: &http.Server{
			Addr:    ":" + port,
			Handler: router,
		},
		linksService: linksServ,
		errCh:        make(chan error, 1),
	}
	router.HandleFunc("/", httpServer.performIndexRequest)
	return httpServer
}

func (httpServer HttpServer) Start() {
	go func() {
		if err := httpServer.httpServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			httpServer.errCh <- err
		}
	}()
}

func (httpServer HttpServer) Stop() error {
	return httpServer.httpServer.Shutdown(context.Background())
}

func (httpServer HttpServer) Error() chan error {
	return httpServer.errCh
}

func (httpServer HttpServer) performIndexRequest(w http.ResponseWriter, req *http.Request) {
	body, err := io.ReadAll(req.Body)
	if err != nil {
		log.Println("[WARNING]:no body request")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	bodyStr := string(body)
	log.Printf("[%s]:%s", req.Method, bodyStr)
	switch req.Method {
	case "POST":
		{
			response, _ := httpServer.linksService.Create(context.TODO(), &links.CreateShortLinkRequest{
				OriginalLink: &links.Link{
					Url: bodyStr,
				},
			})
			if err != nil {
				log.Println("[ERROR]:" + err.Error())
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(httpServer.httpServer.Addr + "/" + response.ShortLink.Url))
		}
	case "GET":
		{
			response, err := httpServer.linksService.Retrive(context.TODO(), &links.RetriveOriginalLinkRequest{
				ShortLink: &links.Link{
					Url: bodyStr,
				},
			})
			if err != nil {
				log.Println("[ERROR]:" + err.Error())
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(response.OriginalLink.Url))
		}
	default:
		{
			w.WriteHeader(http.StatusMethodNotAllowed)
			w.Write([]byte("[ERROR]: invalid method"))
		}
	}
}
