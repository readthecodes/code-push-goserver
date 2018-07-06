package main

import (
	"fmt"
	"net/http"
	"runtime"

	"github.com/gin-gonic/gin"

	"github.com/readthecodes/code-push-goserver/routes"
)

func main() {
	// 使用多核
	runtime.GOMAXPROCS(runtime.NumCPU())

	// 生成一个gin实例
	app := gin.Default()
	app.LoadHTMLGlob("templates/*")

	// 配置路由
	// general
	app.GET("/", routes.IndexView)
	app.GET("/tokens", routes.TokenGetView)
	app.GET("/updateCheck", routes.UpdateCheckView)
	app.POST("/reportStatus/download", routes.ReportDownloadView)
	app.POST("/reportStatus/deploy", routes.ReportDeployView)
	app.POST("/authenticated", routes.AuthenticatedView)
	// auth
	app.GET("/auth/login", routes.LoginWebView)
	app.POST("/auth/login", routes.LoginPostView)
	app.POST("/auth/logout", routes.LogoutView)
	app.GET("/auth/link", routes.LinkView)
	app.GET("/auth/register", routes.RegisterView)
	// accessKeys
	app.GET("/accessKeys/", routes.AccessKeysGetView)
	app.POST("/accessKeys/", routes.AccessKeysPostView)
	app.DELETE("/accessKeys/:name", routes.AccessKeysDeleteView)
	app.PATCH("/accessKeys/:name", routes.AccessKeysPatchView)
	// sessions
	app.DELETE("/sessions/:machineName", routes.SessionDeleteView)
	// account
	app.GET("/account", routes.AccountGetView)
	// users
	app.GET("/users/", routes.UsersGetView)
	app.POST("/users/", routes.UsersCreateView)
	app.GET("/users/exists", routes.UsersExistsView)
	app.POST("/users/registerCode", routes.RegisterCodeCreateView)
	app.POST("/users/registerCode/exists", routes.RegisterCodeExistsView)
	app.POST("/users/password", routes.UsersPasswordUpdateView)
	// apps
	app.GET("/apps/", routes.AppsListView)
	app.POST("/apps/", routes.AppCreateView)
	app.PATCH("/apps/:appName", routes.AppUpdateView)
	app.DELETE("/apps/:appName", routes.AppDeleteView)
	app.POST("/apps/:appName/transfer/:email", routes.AppTransferView)
	app.GET("/apps/:appName/deployments", routes.DeploymentsListView)
	app.POST("/apps/:appName/deployments", routes.DeploymentsAddView)
	app.GET("/apps/:appName/deployments/:deploymentName", routes.DeploymentGetView)
	app.PATCH("/apps/:appName/deployments/:deploymentName", routes.DeploymentUpdateView)
	app.DELETE("/apps/:appName/deployments/:deploymentName", routes.DeploymentDeleteView)
	app.POST("/apps/:appName/deployments/:deploymentName/release", routes.DeploymentReleaseCreateView)
	app.PATCH("/apps/:appName/deployments/:deploymentName/release", routes.DeploymentReleaseUpdateView)
	app.POST("/apps/:appName/deployments/:sourceDeploymentName/promote/:destDeploymentName", routes.DeploymentPromoteView)
	app.POST("/apps/:appName/deployments/:deploymentName/rollback", routes.DeploymentRollbackView)
	app.POST("/apps/:appName/deployments/:deploymentName/rollback/label", routes.DeploymentRollbackView)
	app.GET("/apps/:appName/deployments/:deploymentName/metrics", routes.DeploymentMetricsView)
	app.GET("/apps/:appName/deployments/:deploymentName/history", routes.DeploymentHistoryView)
	app.DELETE("/apps/:appName/deployments/:deploymentName/history", routes.DeploymentHistoryDeleteView)
	app.GET("/apps/:appName/collaborators", routes.CollaboratorsListView)
	app.POST("/apps/:appName/collaborators/:email", routes.CollaboratorCreateView)
	app.DELETE("/apps/:appName/collaborators/:email", routes.CollaboratorDeleteView)

	server := &http.Server{
		Addr:           ":6666",
		Handler:        app,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	server.ListenAndServe()
}
