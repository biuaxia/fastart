package support

import (
	"flag"
	"fmt"
	"github.com/biuaxia/fart/code/core"
	"github.com/biuaxia/fart/code/tool/result"
	"log"
	"net/http"
	"strings"
)

const (
	// start web. This is the default mode.
	MODE_WEB = "web"
	// Current version.
	MODE_VERSION = "version"
)

type FartApplication struct {
	// mode
	mode string

	// host and port  default: http://127.0.0.1:core.DEFAULT_SERVER_PORT
	host string
	// username
	username string
	// password
	password string

	// source file/directory different mode has different usage.
	src string
	// destination directory path(relative to root)
	dest string
	// true: overwrite, false:skip
	overwrite bool
	filename  string
}

// Start the application.
func (this *FartApplication) Start() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("ERROR:%v\r\n", err)
		}
	}()

	modePtr := flag.String("mode", this.mode, "cli mode web/version")
	hostPtr := flag.String("host", this.username, "fart host")
	usernamePtr := flag.String("username", this.username, "username")
	passwordPtr := flag.String("password", this.password, "password")
	srcPtr := flag.String("src", this.src, "src absolute path")
	destPtr := flag.String("dest", this.dest, "destination path in fart.")
	overwritePtr := flag.Bool("overwrite", this.overwrite, "whether same file overwrite")
	filenamePtr := flag.String("filename", this.filename, "filename when crawl")

	// flag.Parse() must invoke before use.
	flag.Parse()

	this.mode = *modePtr
	this.host = *hostPtr
	this.username = *usernamePtr
	this.password = *passwordPtr
	this.src = *srcPtr
	this.dest = *destPtr
	this.overwrite = *overwritePtr
	this.filename = *filenamePtr

	// default start as web.
	if this.mode == "" || strings.ToLower(this.mode) == MODE_WEB {
		this.HandleWeb()
	} else if strings.ToLower(this.mode) == MODE_VERSION {
		this.HandleVersion()
	} else {
		panic(result.BadRequest("cannot handle mode %s \r\n", this.mode))
	}

}

func (this *FartApplication) HandleWeb() {
	// Step 1. Logger
	fartLogger := &FartLogger{}
	core.LOGGER = fartLogger
	fartLogger.Init()
	defer fartLogger.Destroy()

	// Step 2. Configuration
	fartConfig := &FartConfig{}
	core.CONFIG = fartConfig
	fartConfig.Init()

	// Step 3. Global Context
	fartContext := &FartContext{}
	core.CONTEXT = fartContext
	fartContext.Init()
	defer fartContext.Destroy()

	// Step 4. Start http
	http.Handle("/", core.CONTEXT)
	core.LOGGER.Info("App started at http://localhost:%v", core.CONFIG.ServerPort())

	dotPort := fmt.Sprintf(":%v", core.CONFIG.ServerPort())
	err1 := http.ListenAndServe(dotPort, nil)
	if err1 != nil {
		log.Fatal("ListenAndServe: ", err1)
	}
}

// fetch the application version
func (this *FartApplication) HandleVersion() {
	fmt.Printf("Fart %s\r\n", core.VERSION)
}
