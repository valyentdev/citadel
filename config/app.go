package config

import (
	"citadel/app/controllers"
	apiControllers "citadel/app/controllers/api"
	authControllers "citadel/app/controllers/auth"
	"citadel/app/drivers"
	dockerDriver "citadel/app/drivers/docker_driver"
	"citadel/app/listeners"
	"citadel/app/repositories"
	"citadel/app/services"
	"citadel/public"
	"log/slog"
	"os"

	"github.com/caesar-rocks/core"
	"github.com/caesar-rocks/events"
)

func ProvideApp(env *EnvironmentVariables) *core.App {
	app := core.NewApp(&core.AppConfig{
		Addr: os.Getenv("ADDR"),
	})

	app.RegisterProviders(
		repositories.NewWebsiteVisitsRepository,
		controllers.NewAnalyticsWebsitesController,
		repositories.NewAnalyticsWebsitesRepository,
		apiControllers.NewEmailsController,
		controllers.NewMailDomainsController,
		controllers.NewMailApiKeysController,
		controllers.NewStorageController,
		controllers.NewDatabasesController,
		controllers.NewStripeController,
		authControllers.NewCliController,
		authControllers.NewResetPwdController,
		controllers.NewGithubController,
		controllers.NewDeploymentsController,
		controllers.NewEnvController,
		controllers.NewLogsController,
		controllers.NewCertsController,
		authControllers.NewSignOutController,
		controllers.NewBillingController,
		controllers.NewSettingsController,
		controllers.NewAppsController,
		authControllers.NewGithubController,
		authControllers.NewForgotPwdController,
		authControllers.NewSignInController,
		authControllers.NewSignUpController,
	)

	app.RegisterProviders(
		listeners.NewUsersListener,
		listeners.NewDeploymentsListener,
	)

	app.RegisterProviders(
		services.NewUsersService,
		services.NewAppsService,
		dockerDriver.NewDockerDriver,
	)

	app.RegisterProviders(
		repositories.NewUsersRepository,
		repositories.NewApplicationsRepository,
		repositories.NewCertificatesRepository,
		repositories.NewDeploymentsRepository,
		repositories.NewStorageBucketsRepository,
		repositories.NewDatabasesRepository,
		repositories.NewMailDomainsRepository,
		repositories.NewMailApiKeysRepository,
	)

	app.RegisterProviders(
		RegisterRoutes,
		ProvideStripe,
		ProvideEnvironmentVariables,
		ProvideDatabase,
		ProvideErrorHandler,
		RegisterEventsEmitter,
		ProvideAuth,
		ProvideDrive,
		ProvideDriver,
		ProvideMailer,
		ProvideRedis,
		ProvideVexillum,
	)

	app.RegisterInvokers(
		core.ServeStaticAssets(public.FS),
		events.ListenForEvents,
		func(driver drivers.Driver) {
			go func() {
				if err := driver.Init(); err != nil {
					slog.Error("Failed to initialize driver", "error", err)
				}
			}()
		},
	)

	return app
}
