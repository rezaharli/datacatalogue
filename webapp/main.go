package main

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/eaciit/clit"
	"github.com/eaciit/toolkit"

	"git.eaciitapp.com/sebar/knot"

	c "eaciit/datacatalogue/webapp/controllers"
	s "eaciit/datacatalogue/webapp/services"
)

func main() {
	toolkit.Println(time.Now().String())
	clit.LoadConfigFromFlag("", "", filepath.Join(clit.ExeDir(), "..", "webapp", "config", "app.json"))
	if err := clit.Commit(); err != nil {
		kill(err)
	}
	defer clit.Close()

	app := knot.NewApp()
	app.Register(c.NewUsersController(), "")
	app.Register(c.NewDSCController(), "")
	app.Register(c.NewDPOController(), "")
	app.Register(c.NewDDOController(), "")
	app.Register(c.NewRFOController(), "")

	app.Static("css", filepath.Join(clit.ExeDir(), "views", "dist", "css"))
	app.Static("js", filepath.Join(clit.ExeDir(), "views", "dist", "js"))
	app.Static("img", filepath.Join(clit.ExeDir(), "views", "dist", "img"))

	server := knot.NewServer()
	server.RegisterApp(app, "")

	fmt.Println("environment", clit.Config("default", "environment", "").(string))
	switch clit.Config("default", "environment", "").(string) {
	case "prod", "production":
		server.Route("/", func(ctx *knot.WebContext) {
			ctx.Writer.Header().Set("Access-Control-Allow-Origin", "*")
			ctx.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
			ctx.Writer.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

			htmlPath := filepath.Join(clit.ExeDir(), "..", "webapp", "views", "dist", "index.html")
			toolkit.Println("htmlPath:", htmlPath)
			http.ServeFile(ctx.Writer, ctx.Request, htmlPath)
		})
	case "dev", "development":
		server.ReverseProxy("/", "http://localhost:8080", knot.VirtualDirectory)
	}

	server.Start(clit.Config("default", "host", "").(string) + ":" + clit.Config("default", "port", "").(string))

	s.NewDSCService().CreateSystemDummyData()
	s.NewDSCService().CreatePolicyDummyData()
	s.NewDSCService().CreateCategoryDummyData()
	s.NewDSCService().CreateSubCategoryDummyData()
	s.NewDSCService().CreateBusinessTermDummyData()
	s.NewDSCService().CreateMDResourceDummyData()
	s.NewDSCService().CreateMDTableDummyData()
	s.NewDSCService().CreateMDColumnDummyData()
	s.NewDSCService().CreatePeopleDummyData()
	s.NewDSCService().CreateUserDummyData()
	s.NewDSCService().CreateDSProcessesDummyData()
	s.NewDSCService().CreateSegmentDummyData()
	s.NewDSCService().CreateDSProcessesDetailDummyData()
	s.NewDSCService().CreateLinkSubcategoryPeopleDummyData()
	s.NewDSCService().CreateLinkCategoryPeopleDummyData()
	s.NewDSCService().CreatePriorityReportsDummyData()
	s.NewDSCService().CreateCRMDummyData()
	s.NewDSCService().CreateLinkCRMCDEDummyData()
	s.NewDSCService().CreateRoleDummyData()
	s.NewDSCService().CreateLinkRolePeopleDummyData()
	s.NewDSCService().CreateLinkColumnInterfaceDummyData()

	server.Wait()
}

func kill(err error) {
	fmt.Printf("error. %s \n", err.Error())
	os.Exit(100)
}
