package raft

import "go.etcd.io/etcd/raft"

type Node struct {
	raftNode  raft.Node
}

