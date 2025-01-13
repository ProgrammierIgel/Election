package storage

import (
	"github.com/programmierigel/voting/voting"
)

type Store interface {
	CountVoting() CountingVotes
	GetCandidates() []string
	IsVotingActive() bool

	CheckPassword(passwordToCheck string) bool
	InsertVote(vote voting.Vote) error

	ActivateVoting(password string) (bool, error)
	GetAllUndefinedVotes() ([]AllVotes, error)
	DeactivateVoting(password string) (bool, error)
	DeleteAll(password string) error
	DeleteAllVotes(password string) error
	InsertNewVotable(password string, votabel string) error
	AddCandidate(password string, candidate string) error
}
