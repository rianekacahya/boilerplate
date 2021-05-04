package bootstrap

import (
	"compress/flate"
	"context"
	"fmt"
	chim "github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/spf13/cobra"
	"log"
	"net/http"
	"os"
	"os/signal"
	middleware "github.com/rianekacahya/boilerplate/middleware/rest"
	"github.com/rianekacahya/boilerplate/pkg/response"
	"syscall"
	"time"

	auu "github.com/rianekacahya/boilerplate/internal/auth"
	aur "github.com/rianekacahya/boilerplate/internal/auth/repository"
	aut "github.com/rianekacahya/boilerplate/transport/rest/auth"

	ouu "github.com/rianekacahya/boilerplate/internal/oauth"
	our "github.com/rianekacahya/boilerplate/internal/oauth/repository"
	out "github.com/rianekacahya/boilerplate/transport/rest/oauth"
)

var (
	rest = &cobra.Command{
		Use:   "rest",
		Short: "Starting API Server",
		Run: func(cmd *cobra.Command, args []string) {
			// init context
			ctx, cancel := context.WithCancel(context.Background())

			// init repository
			authRepo := aur.NewAuthRepository(dbw, dbr)
			oauthRepo := our.NewOauthRepository(dbw, dbr, rdb)

			// init usecase
			authUsecase := auu.NewAuthUsecase(authRepo)
			oauthUsecase := ouu.NewOauthUsecase(oauthRepo, jwt, cfg)

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
				middleware.Debug(cfg.GetBool("debug")),
				chim.NoCache,
				chim.RequestID,
				chim.RealIP,
				chim.Recoverer,
				chim.Compress(flate.DefaultCompression),
			)

			// init router group
			r.Route("/v1", func(r chi.Router) {
				r.Use(middleware.RequestLogger(logger, cfg.GetBool("debug")))

				// bootstrap auth transport
				aut.NewHandler(r, authUsecase)

				// bootstrap oauth transport
				out.NewHandler(r, oauthUsecase)
			})

			// start chi server with gracefully shutdown
			srv := &http.Server{
				Addr:         fmt.Sprintf(":%s", cfg.GetString("rest.port")),
				Handler:      r,
				ReadTimeout:  time.Duration(cfg.GetInt("rest.read_timeout")) * time.Second,
				WriteTimeout: time.Duration(cfg.GetInt("rest.write_timeout")) * time.Second,
				IdleTimeout:  time.Duration(cfg.GetInt("rest.idle_timeout")) * time.Second,
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

				log.Printf("Shutdown HTTP server on port : %v", cfg.GetString("rest.port"))

				cancel()
				close(idleConnsClosed)
			}()

			log.Printf("HTTP server run on port : %v", cfg.GetString("rest.port"))
			if err := srv.ListenAndServe(); err != http.ErrServerClosed {
				log.Fatalf("HTTP server listen and serve error: %v", err)
			}

			<-idleConnsClosed
		},
	}
)
