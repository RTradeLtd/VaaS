package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	etcdv3 "github.com/coreos/etcd/clientv3"
	"github.com/lytics/grid"
)

type LeaderActor struct {
	client *grid.Client
}

const timeout = time.Second * 2

func (a *LeaderActor) Act(c context.Context) {
	ticker := time.NewTicker(timeout)
	defer ticker.Stop()
	existing := make(map[string]bool)

	for {

		select {
		case <-c.Done():
			return
		case <-ticker.C:
			// check for current peers
			peers, err := a.client.Query(timeout, grid.Peers)
			if err != nil {
				log.Fatal(err)
			}
			for _, peer := range peers {
				if existing[peer.Name()] {
					continue
				}
				// create a worker
				existing[peer.Name()] = true
				start := grid.NewActorStart("worker-%d", len(existing))
				start.Type = "worker"

				// for new peers start the worker
				_, err := a.client.Request(timeout, peer.Name(), start)
				if err != nil {
					log.Fatal(err)
				}
			}
		}
	}
}

// WorkerActor is started by our leader
type WorkerActor struct{}

func (a *WorkerActor) Act(ctx context.Context) {
	fmt.Println("Hello World")
	for {
		select {
		case <-ctx.Done():
			fmt.Println("good bye")
			return
		}
	}
}

func InitializeDistributor(address string) error {
	etcd, err := etcdv3.New(
		etcdv3.Config{Endpoints: []string{"localhost:2379"}},
	)
	if err != nil {
		return err
	}
	client, err := grid.NewClient(
		etcd,
		grid.ClientCfg{Namespace: "vaas"},
	)
	if err != nil {
		return err
	}
	server, err := grid.NewServer(
		etcd,
		grid.ServerCfg{Namespace: "vaas"},
	)
	if err != nil {
		return err
	}

	server.RegisterDef("leader", func(_ []byte) (grid.Actor, error) {
		return &LeaderActor{client: client}, nil
	})

	server.RegisterDef("worker", func(_ []byte) (grid.Actor, error) {
		return &WorkerActor{}, nil
	})

	go func() {
		sig := make(chan os.Signal, 1)
		signal.Notify(sig, os.Interrupt, syscall.SIGTERM)
		<-sig
		fmt.Println("shutting down")
		server.Stop()
		fmt.Println("shutdown complete")
	}()

	listen, err := net.Listen("tcp", address)
	if err != nil {
		return err
	}

	return server.Serve(listen)
}
