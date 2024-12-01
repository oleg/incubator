package agent

import (
	"crypto/tls"
	"github.com/oleg/incubator/go/proglog/internal/discovery"
	"github.com/oleg/incubator/go/proglog/internal/log"
	"google.golang.org/grpc"
	"sync"
)

type Agent struct {
	Config

	log        *log.Log
	server     *grpc.Server
	membership *discovery.Membership
	replicator *log.Replicator

	shutdown     bool
	shutdowns    chan struct{}
	shutdownLock sync.Mutex
}

type Config struct {
	ServerTLSConfig *tls.Config
	PeerTLSConfig   *tls.Config
	DataDir         string
	BindAddr        string
	
	///
}
