package main

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"github.com/eaciit/clit"

	"git.eaciitapp.com/sebar/knot"

	c "eaciit/datacatalogue/webapp/controllers"
	// s "eaciit/datacatalogue/webapp/services"
)

func main() {
	clit.LoadConfigFromFlag("", "", filepath.Join(clit.ExeDir(), "..", "webapp", "config", "app.json"))
	if err := clit.Commit(); err != nil {
		kill(err)
	}
	defer clit.Close()

	// s.NewDSCService().CreateSystemDummyData()
	// s.NewDSCService().CreateMDTableDummyData()
	// s.NewDSCService().CreateMDColumnDummyData()
	// s.NewDSCService().CreateBusinessTermDummyData()
	// s.NewDSCService().CreateSubCategoryDummyData()
	// s.NewDSCService().CreateCategoryDummyData()
	// s.NewDSCService().CreatePolicyDummyData()
	// s.NewDSCService().CreateDSProcessesDummyData()
	// s.NewDSCService().CreateDSProcessesDetailDummyData()

	app := knot.NewApp()
	app.Register(c.NewUsersController(), "")
	app.Register(c.NewDSCController(), "")
	app.Register(c.NewDPOController(), "")

	app.Static("css", filepath.Join(clit.ExeDir(), "views", "dist", "css"))
	app.Static("js", filepath.Join(clit.ExeDir(), "views", "dist", "js"))
	app.Static("img", filepath.Join(clit.ExeDir(), "views", "dist", "img"))

	s := knot.NewServer()
	s.RegisterApp(app, "")

	fmt.Println("environment", clit.Config("default", "environment", "").(string))
	switch clit.Config("default", "environment", "").(string) {
	case "prod", "production":
		s.Route("/", func(ctx *knot.WebContext) {
			ctx.Writer.Header().Set("Access-Control-Allow-Origin", "*")
			ctx.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
			ctx.Writer.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

			http.ServeFile(ctx.Writer, ctx.Request, filepath.Join("views/dist", "/index.html"))
		})
	case "dev", "development":
		s.ReverseProxy("/", "http://localhost:8080", knot.VirtualDirectory)
	}

	s.Start(clit.Config("default", "host", "").(string) + ":" + clit.Config("default", "port", "").(string))
	s.Wait()
}

func kill(err error) {
	fmt.Printf("error. %s \n", err.Error())
	os.Exit(100)
}
