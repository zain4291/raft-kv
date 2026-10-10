package raft

import (
	"fmt"
	"sync"
)

// Raft structure for the algorithm
type Raft struct {
	mu          sync.Mutex
	currentTerm int        // persistent: latest term server has seen (initialized to 0)
	votedFor    int        // persistent: candidateId that received vote in current term (-1 if none)
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
	Term    int    // the term when the leader created this entry
	Command string // command for state machine e.g. Put 10
}

type RequestVoteArgs struct {
	Term         int // candidate's term
	CandidateId  int // candidate's requesting vote
	LastLogIndex int // index of candidate's last log entry
	LastLogTerm  int // term of candidate's last log entry
}

type RequestVoteReply struct {
	Term        int  // currentTerm for candidate to update itself
	VoteGranted bool // true means candidate received vote
}

func (rf *Raft) RequestVote(args *RequestVoteArgs, reply *RequestVoteReply) {
	rf.mu.Lock()
	defer rf.mu.Unlock()

	// If request or response term is greater than current term, than change current term to that term
	if args.Term > rf.currentTerm {
		rf.currentTerm = args.Term
		rf.votedFor = -1
		reply.VoteGranted = false
	}

	if reply.Term > rf.currentTerm {
		rf.currentTerm = reply.Term
		rf.votedFor = -1
		reply.VoteGranted = false
	}
	// If term is less than the current term, reply with false
	if reply.Term < rf.currentTerm {
		fmt.Println("Out of date. Revert to follower")
		reply.VoteGranted = false
	}
	// If votedfor is -1 or its the candidate itself, and its candidate log is at least as up to date as receiver's log, give vote

	if rf.votedFor == -1 || args.CandidateId == rf.id {
		rf.votedFor = args.CandidateId
		reply.VoteGranted = true
	}
}
