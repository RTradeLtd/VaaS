package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/RTradeLtd/VaaS/ethereum"
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
type WorkerActor struct {
	EG     *ethereum.EthereumGenerator
	server *grid.Server
}

func (a *WorkerActor) Act(ctx context.Context) {
	name, err := grid.ContextActorName(ctx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("starting worker ", name)
	mailbox, err := grid.NewMailbox(a.server, name, 10)
	if err != nil {
		log.Fatal(err)
	}
	defer mailbox.Close()
	fmt.Println("waiting for messages")
	for {
		select {
		case <-ctx.Done():
			fmt.Println("good bye")
			return
		case req, ok := <-mailbox.C:
			if !ok {
				fmt.Println("erorr retrieving message from mailbox")
				return
			}
			switch i := req.Msg().(type) {
			case *GenerationRequest:
				eg := ethereum.InitializeEthereumGenerator(i.SearchPrefix, 10000000)
				suc, err := eg.Run()
				if err != nil {
					fmt.Println(err)
					continue
				}
				err = req.Respond(&GenerationResponse{Key: suc.Key, Address: suc.Address})
				if err != nil {
					fmt.Println("failed to send response ", err)
				}
			default:
				fmt.Printf("ERROR: wrong type %#v\n", req.Msg())
			}
		}
	}
}

func InitializeDistributor(address string) error {

	grid.Register(GenerationRequest{})
	grid.Register(GenerationResponse{})

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
		return &WorkerActor{server: server}, nil
	})

	go func() {
		sig := make(chan os.Signal, 1)
		signal.Notify(sig, os.Interrupt, syscall.SIGTERM)
		<-sig
		fmt.Println("shutting down")
		server.Stop()
		fmt.Println("shutdown complete")
	}()

	api := InitializeAPI(client)
	go api.Router.Run("127.0.0.1:6767")
	listen, err := net.Listen("tcp", address)
	if err != nil {
		return err
	}

	return server.Serve(listen)
}

type GAPI struct {
	c *grid.Client
}

func GenerateAPI(client *grid.Client) *GAPI {
	a := &GAPI{c: client}
	return a
}

func (a *GAPI) Run(address string) {
	http.HandleFunc("/api/v1/ethereum/generate/distributed", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("handling worker")

		req, err := a.c.Request(timeout, "worker-1", &GenerationRequest{SearchPrefix: "DD"})
		fmt.Printf("request worker:%q\nresponse: %#v\nerr%v\n", "worker-1", req, err)
		if resp, ok := req.(*GenerationResponse); ok {
			fmt.Fprintf(w, "Address %s\nKey %s\n", resp.Address, resp.Key)
		} else {
			fmt.Fprintf(w, "wrong resposne type")
		}
	})

	err := http.ListenAndServe(address, nil)
	if err != nil {
		log.Fatal(err)
	}
}
