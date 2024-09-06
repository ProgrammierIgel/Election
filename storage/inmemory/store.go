package inmemory

import (
	"fmt"
	"reflect"

	"github.com/programmierigel/voting/storage"
	"github.com/programmierigel/voting/tools"
	"github.com/programmierigel/voting/voting"
)

type Store struct {
	votes         storage.Candidates
	votesCounting storage.CountingVotes
	votingActive  bool
	password      string
}

func New(password string) *Store {
	return &Store{
		votes: storage.Candidates{
			Undefined:  []string{"123"},
			Candidate1: []string{},
			Candidate2: []string{},
			Candidate3: []string{},
			Candidate4: []string{},
			Candidate5: []string{},
		},
		votesCounting: storage.CountingVotes{
			Undefined:  0,
			Candidate1: 0,
			Candidate2: 0,
			Candidate3: 0,
			Candidate4: 0,
			Candidate5: 0,
		},
		votingActive: false,
		password:     password,
	}
}

func (s *Store) GetCandidates() []string {
	// (i) By Gemini
	t := reflect.TypeOf(storage.Candidates{})

	keys := make([]string, 0, t.NumField())

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		if field.Name != "" {
			keys = append(keys, field.Name)
		}
	}

	return keys
}

func (s *Store) CountVoting() storage.CountingVotes {
	return s.votesCounting
}

func (s *Store) InsertVote(insertedVote voting.Vote) error {
	allCandidates := s.GetCandidates()

	candidate := insertedVote.Candidate

	candidate = tools.FormatToValidString(candidate)

	id := insertedVote.ID
	if !s.votingActive {
		return fmt.Errorf("no voting active")
	}

	if !tools.StringInSlice(candidate, allCandidates) {
		fmt.Println("LOL2")
		return fmt.Errorf("not valid vote")
	}

	if candidate == "Undefined" {
		return fmt.Errorf("not valid vote")
	}

	if !tools.StringInSlice(id, s.votes.Undefined) {
		return fmt.Errorf("not valid id // doublevote")
	}

	values := reflect.ValueOf(s.votes)
	types := values.Type()

	for i := 0; i < values.NumField(); i++ {
		if types.Field(i).Name == candidate {
			slice := reflect.Append(values.Field(i), reflect.ValueOf(id)).Interface().([]string)

			switch types.Field(i).Name {

			case "Candidate1":
				s.votes.Candidate1 = slice
				s.votesCounting.Candidate1 += 1
			case "Candidate2":
				s.votes.Candidate2 = slice
				s.votesCounting.Candidate2 += 1
			case "Candidate3":
				s.votes.Candidate3 = slice
				s.votesCounting.Candidate3 += 1
			case "Candidate4":
				s.votes.Candidate4 = slice
				s.votesCounting.Candidate4 += 1
			case "Candidate5":
				s.votes.Candidate5 = slice
				s.votesCounting.Candidate5 += 1
			}

		}
	}

	index, err := tools.GetIndexOfElementInSlice(s.votes.Undefined, id)
	if err != nil {
		return fmt.Errorf("internal server error")
	}

	s.votes.Undefined = tools.RemoveElementFromSlice(s.votes.Undefined, index)

	return nil

}

func (s *Store) DeleteAll(password string) error {
	if password != s.password {
		return fmt.Errorf("unknown password")
	}

	s.votes = storage.Candidates{
		Undefined:  []string{},
		Candidate1: []string{},
		Candidate2: []string{},
		Candidate3: []string{},
		Candidate4: []string{},
		Candidate5: []string{},
	}

	s.votesCounting = storage.CountingVotes{
		Undefined:  0,
		Candidate1: 0,
		Candidate2: 0,
		Candidate3: 0,
		Candidate4: 0,
		Candidate5: 0,
	}

	return nil
}

func (s *Store) ActivateVoting(password string) (bool, error) {
	if s.votingActive {
		return true, fmt.Errorf("voting already active")
	}

	if password == s.password {
		s.votingActive = true
		return s.votingActive, nil
	}

	return s.votingActive, fmt.Errorf("unknown password")
}

func (s *Store) DeactivateVoting(password string) (bool, error) {
	if !s.votingActive {
		return true, fmt.Errorf("voting already inactiveW")
	}

	if password == s.password {
		s.votingActive = false
		return s.votingActive, nil
	}

	return s.votingActive, fmt.Errorf("unknown password")
}

func (s *Store) InsertNewVotable(password string, votabel string) error {
	if password != s.password {
		return fmt.Errorf("unkown password")
	}

	s.votes.Undefined = append(s.votes.Undefined, votabel)
	return nil
}

func (s *Store) IsVotingActive() bool {
	return s.votingActive

}
