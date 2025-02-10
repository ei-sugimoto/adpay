package cmd

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/ei-sugimoto/adpay/apps/backend/controller"
	"github.com/ei-sugimoto/adpay/apps/backend/infra"
	"github.com/ei-sugimoto/adpay/apps/backend/infra/persistence"
	"github.com/ei-sugimoto/adpay/apps/backend/middleware"
	"github.com/ei-sugimoto/adpay/apps/backend/usecase"
	"github.com/uptrace/bun/extra/bundebug"
)

func Serve() {
	mux := http.NewServeMux()
	muxWithMiddleware := middleware.LoggingMiddleware(mux)
	muxWithMiddleware = middleware.AuthenticatingMiddleware(muxWithMiddleware)

	db := infra.NewDB()
	db.AddQueryHook(bundebug.NewQueryHook(bundebug.WithVerbose(true)))
	defer db.Close()

	userPersistence := persistence.NewUserPersistence(db)
	projectPersistence := persistence.NewProjectPersistence(db, userPersistence)

	userUsecase := usecase.NewUserUsecase(userPersistence)
	projectUsecase := usecase.NewProjectUsecase(projectPersistence)

	userController := controller.NewUserController(userUsecase)
	projectController := controller.NewProjectController(projectUsecase)

	Routes := map[string]func(http.ResponseWriter, *http.Request){
		"/register": userController.Register,
		"/login":    userController.Login,
		"/project":  projectController.Save,
	}

	for pattern, handler := range Routes {
		mux.HandleFunc(pattern, handler)
	}

	server := &http.Server{
		Addr:    ":8000",
		Handler: muxWithMiddleware,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Could not listen on :8000: %v\n", err)
		}
	}()
	log.Println("Server is ready to handle requests at :8000")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Server is shutting down...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Could not gracefully shutdown the server: %v\n", err)
	}
	log.Println("Server stopped")
}
