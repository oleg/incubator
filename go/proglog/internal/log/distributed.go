package log

import (
	"github.com/hashicorp/raft"
	raftboltdb "github.com/hashicorp/raft-boltdb"
	"io"
	"net"
	"os"
	"path/filepath"
	"time"
)

type DistributedLog struct {
	config Config
	log    *Log
	raft   *raft.Raft
}

func NewDistributedLog(dataDir string, config Config) (*DistributedLog, error) {
	l := &DistributedLog{config: config}
	if err := l.setupLog(dataDir); err != nil {
		return nil, err
	}
	if err := l.setupRaft(dataDir); err != nil {
		return nil, err
	}
	return l, nil
}

func (l *DistributedLog) setupLog(dataDir string) error {
	logDir := filepath.Join(dataDir, "log")
	if err := os.MkdirAll(logDir, 0755); err != nil {
		return err
	}
	log, err := NewLog(logDir, l.config)
	if err != nil {
		return err
	}
	l.log = log
	return nil
}

func (l *DistributedLog) setupRaft(dataDir string) error {
	fsm := &fsm{log: l.log}

	logDir := filepath.Join(dataDir, "raft", "log")
	if err := os.MkdirAll(logDir, 0755); err != nil {
		return err
	}
	logConfig := l.config
	logConfig.Segment.InitialOffset = 1
	logStore, err := newLogStore(logDir, logConfig)
	if err != nil {
		return err
	}
	stableStore, err := raftboltdb.NewBoltStore(filepath.Join(dataDir, "raft", "stable"))
	if err != nil {
		return err
	}
	retain := 1
	snapshotStore, err := raft.NewFileSnapshotStore(filepath.Join(dataDir, "raft"), retain, os.Stderr)
	if err != nil {
		return err
	}
	maxPool := 5
	timeout := 10 * time.Second
	transport := raft.NewNetworkTransport(l.config.Raft.StreamLayer, maxPool, timeout, os.Stderr)

	config := raft.DefaultConfig()
	config.LocalID = l.config.Raft.LocalID
	if l.config.Raft.HeartbeatTimeout != 0 {
		config.HeartbeatTimeout = l.config.Raft.HeartbeatTimeout
	}
	if l.config.Raft.ElectionTimeout != 0 {
		config.ElectionTimeout = l.config.Raft.ElectionTimeout
	}
	if l.config.Raft.LeaderLeaseTimeout != 0 {
		config.LeaderLeaseTimeout = l.config.Raft.LeaderLeaseTimeout
	}
	if l.config.Raft.CommitTimeout != 0 {
		config.CommitTimeout = l.config.Raft.CommitTimeout
	}

	l.raft, err = raft.NewRaft(config, fsm, logStore, stableStore, snapshotStore, transport)
	if err != nil {
		return err
	}
	hasState, err := raft.HasExistingState(logStore, stableStore, snapshotStore)
	if err != nil {
		return err
	}
	if l.config.Raft.Bootstrap && !hasState {
		config := raft.Configuration{
			Servers: []raft.Server{{
				ID:      config.LocalID,
				Address: transport.LocalAddr(),
			}},
		}
		err = l.raft.BootstrapCluster(config).Error()
	}
	return err
}

var _ raft.FSM = (*fsm)(nil)

type fsm struct {
	log *Log
}

func (f fsm) Apply(log *raft.Log) interface{} {
	//TODO implement me
	panic("implement me")
}

func (f fsm) Snapshot() (raft.FSMSnapshot, error) {
	//TODO implement me
	panic("implement me")
}

func (f fsm) Restore(snapshot io.ReadCloser) error {
	//TODO implement me
	panic("implement me")
}

var _ raft.LogStore = (*logStore)(nil)

type logStore struct {
	*Log
}

func newLogStore(dir string, config Config) (*logStore, error) {
	l, err := NewLog(dir, config)
	if err != nil {
		return nil, err
	}
	return &logStore{l}, nil
}

func (l logStore) FirstIndex() (uint64, error) {
	//TODO implement me
	panic("implement me")
}

func (l logStore) LastIndex() (uint64, error) {
	//TODO implement me
	panic("implement me")
}

func (l logStore) GetLog(index uint64, log *raft.Log) error {
	//TODO implement me
	panic("implement me")
}

func (l logStore) StoreLog(log *raft.Log) error {
	//TODO implement me
	panic("implement me")
}

func (l logStore) StoreLogs(logs []*raft.Log) error {
	//TODO implement me
	panic("implement me")
}

func (l logStore) DeleteRange(min, max uint64) error {
	//TODO implement me
	panic("implement me")
}

var _ raft.StreamLayer = (*StreamLayer)(nil)

type StreamLayer struct {
}

func (s *StreamLayer) Accept() (net.Conn, error) {
	//TODO implement me
	panic("implement me")
}

func (s *StreamLayer) Close() error {
	//TODO implement me
	panic("implement me")
}

func (s *StreamLayer) Addr() net.Addr {
	//TODO implement me
	panic("implement me")
}

func (s *StreamLayer) Dial(address raft.ServerAddress, timeout time.Duration) (net.Conn, error) {
	//TODO implement me
	panic("implement me")
}
