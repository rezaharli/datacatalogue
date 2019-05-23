package main

import (
	"eaciit/datacatalogue/webapp/helpers"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/fsnotify/fsnotify"

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
	app.Register(c.NewDashboardController(), "")
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

	conn := helpers.Database()
	if conn != nil {
		if clit.Config("default", "dataDummy", false).(bool) == true {
			s.NewDSCService().CreateSystemDummyData()
			s.NewDSCService().CreateMDResourceDummyData()
			s.NewDSCService().CreatePeopleDummyData()
			s.NewDSCService().CreateRoleDummyData()
			s.NewDSCService().CreateCategoryDummyData()
			s.NewDSCService().CreateSubCategoryDummyData()
			s.NewDSCService().CreatePolicyDummyData()
			s.NewDSCService().CreateBusinessTermDummyData()
			s.NewDSCService().CreateMDTableDummyData()
			s.NewDSCService().CreateMDColumnDummyData()
			s.NewDSCService().CreateMDColumnDetailsDummyData()
			s.NewDSCService().CreateLinkColumnBusinessTermDummyData()
			s.NewDSCService().CreateDSProcessesDummyData()
			s.NewDSCService().CreateSegmentDummyData()
			s.NewDSCService().CreateDSProcessesDetailDummyData()
			s.NewDSCService().CreateLinkDSProcessSegmentDummyData()
			s.NewDSCService().CreateLinkColumnInterfaceDummyData()
			s.NewDSCService().CreateLinkColumnGoldenSourceDummyData()
			s.NewDSCService().CreatePriorityReportsDummyData()
			s.NewDSCService().CreateCRMDummyData()
			s.NewDSCService().CreateCDEDummyData()
			s.NewDSCService().CreateLinkRolePeopleDummyData()
		}

		s.NewDSCService().CreateUserDummyData()

		// creates a new file watcher
		filename := "nohup.out"

		err := waitUntilFind(filename)
		if err != nil {
			log.Fatalln(err)
		}

		watcher, err := fsnotify.NewWatcher()
		if err != nil {
			log.Fatalln(err)
		}
		defer watcher.Close()

		err = watcher.Add(filename)
		if err != nil {
			log.Fatalln(err)
		}

		renameCh := make(chan bool)
		removeCh := make(chan bool)
		errCh := make(chan error)

		go func() {
			for {
				select {
				case event := <-watcher.Events:
					switch {
					case event.Op&fsnotify.Write == fsnotify.Write:
						writeLog()
					}
				case err := <-watcher.Errors:
					errCh <- err
				}
			}
		}()

		go func() {
			for {
				select {
				case <-renameCh:
					err = waitUntilFind(filename)
					if err != nil {
						log.Fatalln(err)
					}
					err = watcher.Add(filename)
					if err != nil {
						log.Fatalln(err)
					}
				case <-removeCh:
					err = waitUntilFind(filename)
					if err != nil {
						log.Fatalln(err)
					}
					err = watcher.Add(filename)
					if err != nil {
						log.Fatalln(err)
					}
				}
			}
		}()

		log.Fatalln(<-errCh)
	}

	toolkit.Println("Done. Waiting.")

	server.Wait()
}

func waitUntilFind(filename string) error {
	for {
		time.Sleep(1 * time.Second)
		_, err := os.Stat(filename)
		if err != nil {
			if os.IsNotExist(err) {
				continue
			} else {
				return err
			}
		}
		break
	}
	return nil
}

func writeLog() error {
	copyFile := func(src, dst string) (int64, error) {
		sourceFileStat, err := os.Stat(src)
		if err != nil {
			return 0, err
		}

		if !sourceFileStat.Mode().IsRegular() {
			return 0, fmt.Errorf("%s is not a regular file", src)
		}

		source, err := os.Open(src)
		if err != nil {
			return 0, err
		}
		defer source.Close()

		destination, err := os.Create(dst)
		if err != nil {
			return 0, err
		}
		defer destination.Close()
		nBytes, err := io.Copy(destination, source)
		return nBytes, err
	}

	deleteFile := func(path string) error {
		return os.Remove(path)
	}

	path := "logfiles"
	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.Mkdir(path, os.ModePerm)
	}

	currentTime := time.Now()
	time4daysAgo := currentTime.AddDate(0, 0, -4)

	_ = deleteFile("logfiles/nohup_" + time4daysAgo.Format("2006-01-02") + ".out")
	_, err := copyFile("nohup.out", "logfiles/nohup_"+currentTime.Format("2006-01-02")+".out")

	return err
}

func kill(err error) {
	fmt.Printf("error. %s \n", err.Error())
	os.Exit(100)
}
