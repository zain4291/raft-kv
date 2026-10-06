package raft

import (
	"sync"
)

// Raft structure for the algorithm
type Raft struct {
	mu          sync.Mutex
	currentTerm int        // persistent: latest term server has seen (initialized to 0)
	votedFor    string     // persistent: candidateId that received vote in current term (null if none)
	log         []logEntry // persistent: log entries
	commitIndex int        // volatile: index of highest entry known to be committed (initialized to 0)
	lastApplied int        // volatile: index of highest log entry applied to state machine
	nextIndex   []int      // volatile (leaders only): for each server, index of the next log entry to send to that server
	matchIndex  []int      // volatile (leaders only): for each server, index of highest log entry known to be replicated on server
	id          int        // this server's ID
	peers       []string   // addresses of other servers
	state       string     // state of the server (leader, follower, candidate)
}

type logEntry struct {
	term    int    // the term when the leader created this entry
	command string // command for state machine e.g. Put 10
}
