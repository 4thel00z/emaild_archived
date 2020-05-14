package main

import (
	"context"
	"flag"
	"github.com/4thel00z/emaild/pkg/core"
	"github.com/monzo/typhon"
	"log"
	"net"
	"os"
	"os/signal"
	"path"
	"runtime"
	"syscall"
	"time"
)

var (
	pending      = flag.Int("queue-size", 100, "how many pending emails can the daemon have, before the queue blocks")
	accountsPath = flag.String("accountsPath", "", "where can we find the accountsPath file")
	unixSock     = flag.String("unix-socket", "/tmp/emaild.sock", "path to the unix domain socket")
)

func main() {
	flag.Parse()
	if !flag.Parsed() {
		flag.PrintDefaults()
		os.Exit(1)
	}
	if *accountsPath == "" {
		_, filename, _, ok := runtime.Caller(1)
		if !ok {
			log.Fatal("Could not get current file for default accounts.json!")
		}
		*accountsPath = path.Join(path.Dir(filename), "../config/accounts.json")
	}

	accounts, err := core.ParseAccountsFromPath(*accountsPath)

	if err != nil {
		log.Fatal(err.Error())
	}

	scheduler := core.NewScheduler(*pending, accounts)
	go scheduler.Run()

	svc := core.Router(scheduler).Serve().
		Filter(typhon.ErrorFilter).
		Filter(typhon.H2cFilter)

	unixListener, err := net.Listen("unix", *unixSock)
	if err != nil {
		log.Fatal(err.Error())
	}
	srv, err := typhon.Serve(svc, unixListener)
	if err != nil {
		log.Fatal(err.Error())
	}
	log.Printf("👋  Listening on %v", srv.Listener().Addr())

	done := make(chan os.Signal, 1)
	signal.Notify(done, syscall.SIGINT, syscall.SIGTERM)
	<-done
	scheduler.Exit()
	log.Printf("☠️  Shutting down")
	c, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	srv.Stop(c)
}
