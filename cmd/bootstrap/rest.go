package bootstrap

import (
	"compress/flate"
	"context"
	"fmt"
	chim "github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	middleware "github.com/rianekacahya/boilerplate/middleware/rest"
	"github.com/rianekacahya/boilerplate/pkg/response"
	"github.com/spf13/cobra"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/rianekacahya/boilerplate/transport/rest/auth"
	"github.com/rianekacahya/boilerplate/transport/rest/oauth"
)

var (
	rest = &cobra.Command{
		Use:   "rest",
		Short: "Starting API Server",
		Run: func(cmd *cobra.Command, args []string) {
			// init context
			ctx, cancel := context.WithCancel(context.Background())

			// init rest server
			r := chi.NewRouter()

			// custom not found
			r.NotFound(response.NotFound)

			// custom method not allowed
			r.MethodNotAllowed(response.MethodNotAllowed)

			// set middleware
			r.Use(
				chim.RedirectSlashes,
				chim.Heartbeat("/ping"),
				render.SetContentType(render.ContentTypeJSON),
				middleware.RequestCORS(),
				middleware.RequestHeader(),
				middleware.Debug(dependency.Cfg.GetBool("debug")),
				chim.NoCache,
				chim.RequestID,
				chim.RealIP,
				chim.Recoverer,
				chim.Compress(flate.DefaultCompression),
			)

			// init router group
			r.Route("/v1", func(r chi.Router) {
				r.Use(middleware.RequestLogger(dependency.Logger, dependency.Cfg.GetBool("debug")))

				// bootstrap auth transport
				auth.NewHandler(r, usecase)

				// bootstrap oauth transport
				oauth.NewHandler(r, usecase)
			})

			// start chi server with gracefully shutdown
			srv := &http.Server{
				Addr:         fmt.Sprintf(":%s", dependency.Cfg.GetString("rest.port")),
				Handler:      r,
				ReadTimeout:  time.Duration(dependency.Cfg.GetInt("rest.read_timeout")) * time.Second,
				WriteTimeout: time.Duration(dependency.Cfg.GetInt("rest.write_timeout")) * time.Second,
				IdleTimeout:  time.Duration(dependency.Cfg.GetInt("rest.idle_timeout")) * time.Second,
			}

			idleConnsClosed := make(chan struct{})
			go func() {
				sigint := make(chan os.Signal, 1)
				signal.Notify(sigint, os.Interrupt)
				signal.Notify(sigint, syscall.SIGTERM)
				<-sigint

				// We received an interrupt signal, shut down.
				if err := srv.Shutdown(ctx); err != nil {
					log.Fatalf("HTTP server shutdown error: %v", err)
				}

				log.Printf("Shutdown HTTP server on port : %v", dependency.Cfg.GetString("rest.port"))

				cancel()
				close(idleConnsClosed)
			}()

			log.Printf("HTTP server run on port : %v", dependency.Cfg.GetString("rest.port"))
			if err := srv.ListenAndServe(); err != http.ErrServerClosed {
				log.Fatalf("HTTP server listen and serve error: %v", err)
			}

			<-idleConnsClosed
		},
	}
)
