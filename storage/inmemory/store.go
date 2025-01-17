package inmemory

import (
	"fmt"

	"github.com/programmierigel/voting/tools"
	"github.com/programmierigel/voting/voting"
)

type Store struct {
	name          string
	candidates    voting.Candidates
	votes         voting.CandidatesVoteStore
	votesCounting voting.CountingVotes
	votingActive  bool
	password      string
}

func New(password string) *Store {

	candidates := []string{
		"undefined",
	}

	votingCount := make(map[string]int, len(candidates))

	for i := 0; i < len(candidates); i++ {
		votingCount[candidates[i]] = 0

	}

	return &Store{
		name:          "",
		candidates:    candidates,
		votes:         make(voting.CandidatesVoteStore, len(candidates)),
		votesCounting: votingCount,
		votingActive:  false,
		password:      password,
	}
}

func (s *Store) GetCandidates() []string {
	return s.candidates
}

func (s *Store) CheckPassword(passwordToCheck string) bool {
	return passwordToCheck == s.password
}

func (s *Store) CountVoting() voting.CountingVotes {
	return s.votesCounting
}

func (s *Store) GetAllUndefinedVotes(password string) ([]voting.AllVotes, error) {
	if password != s.password {
		return nil, fmt.Errorf("unknown password")
	}

	if s.votingActive {
		return nil, fmt.Errorf("voting active")
	}

	allVotes := []voting.AllVotes{}
	PosOfUndefined, err := tools.FindInSlice(s.candidates, "undefined")
	if err != nil {
		return allVotes, fmt.Errorf("undefined not found")
	}

	for j := range s.votes[s.candidates[PosOfUndefined]] {

		allVotes = append(allVotes, voting.AllVotes{
			VoteName: s.votes[s.candidates[PosOfUndefined]][j],
			Value:    s.candidates[PosOfUndefined],
		})
	}

	return allVotes, nil
}

func (s *Store) InsertNewVotable(password string, votable string) error {
	if password != s.password {
		return fmt.Errorf("unknown password")
	}

	if s.votingActive {
		return fmt.Errorf("voting is active")
	}

	if s.IDExists(votable) {
		return fmt.Errorf("votable already exists")
	}
	PosUndef, err := tools.FindInSlice(s.candidates, "undefined")
	if err != nil {
		return err
	}
	s.votes[s.candidates[PosUndef]] = append(s.votes[s.candidates[PosUndef]], votable)
	s.votesCounting[s.candidates[PosUndef]] += 1
	return nil
}

func (s *Store) RemoveVotable(password string, id string) error {
	if password != s.password {
		return fmt.Errorf("unknown password")
	}

	if s.votingActive {
		return fmt.Errorf("voting is active")
	}
	PosUndef, err := tools.FindInSlice(s.candidates, "undefined")
	if err != nil {
		return err
	}

	if !tools.StringInSlice(id, s.votes[s.candidates[PosUndef]]) {
		return fmt.Errorf("votable not found")
	}

	PositionOfRemoval, err := tools.FindInSlice(s.votes[s.candidates[PosUndef]], id)

	if err != nil {
		return fmt.Errorf("internal error")
	}

	s.votes[s.candidates[PosUndef]] = tools.RemoveElementFromSlice(s.votes[s.candidates[PosUndef]], PositionOfRemoval)
	s.votesCounting[s.candidates[PosUndef]] -= 1
	return nil
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
		return fmt.Errorf("not valid vote")
	}

	if candidate == "undefined" {
		return fmt.Errorf("not valid vote")
	}

	PositionOfUndefinedInCandidates, err := tools.FindInSlice(s.candidates, "undefined")
	if err != nil {
		return err
	}
	if !tools.StringInSlice(id, s.votes[s.candidates[PositionOfUndefinedInCandidates]]) {
		return fmt.Errorf("not valid id // doublevote")
	}

	for _, v := range s.candidates {
		// fmt.Print(strconv.Itoa(i + 2))
		if v == candidate {
			PositionOfCandidate, err := tools.FindInSlice(s.candidates, candidate)

			if err != nil {
				return err
			}

			s.votes[s.candidates[PositionOfCandidate]] = append(s.votes[v], id)
			s.votesCounting[s.candidates[PositionOfCandidate]] += 1

		}
	}

	index, err := tools.FindInSlice(s.votes[s.candidates[PositionOfUndefinedInCandidates]], id)
	if err != nil {
		return fmt.Errorf("internal server error")
	}

	s.votes[s.candidates[PositionOfUndefinedInCandidates]] = tools.RemoveElementFromSlice(s.votes[s.candidates[PositionOfUndefinedInCandidates]], index)
	s.votesCounting[s.candidates[PositionOfUndefinedInCandidates]] -= 1
	return nil

}

func (s *Store) DeleteAll(password string) error {
	if password != s.password {
		return fmt.Errorf("unknown password")
	}
	newStore := New(s.password)
	s.votes = newStore.votes
	s.candidates = newStore.candidates
	s.votesCounting = newStore.votesCounting
	s.votingActive = newStore.votingActive
	return nil
}

func (s *Store) DeleteAllVotes(password string) error {
	if password != s.password {
		return fmt.Errorf("unknown password")
	}
	newStore := New(s.password)
	s.votes = newStore.votes
	s.votesCounting = newStore.votesCounting
	s.votingActive = newStore.votingActive
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
		return true, fmt.Errorf("voting already inactive")
	}

	if password == s.password {
		s.votingActive = false
		return s.votingActive, nil
	}

	return s.votingActive, fmt.Errorf("unknown password")
}

func (s *Store) IsVotingActive() bool {
	return s.votingActive
}

func (s *Store) AddCandidate(password string, candidate string) error {
	if password != s.password {
		return fmt.Errorf("unknown password")
	}

	if s.votingActive {
		return fmt.Errorf("voting active")
	}

	if tools.StringInSlice(candidate, s.candidates) {
		return fmt.Errorf("candidate already exists")
	}

	s.candidates = append(s.candidates, candidate)
	return nil
}

func (s *Store) RemoveCandidate(password string, candidate string) error {
	if password != s.password {
		return fmt.Errorf("unknown password")
	}

	if s.votingActive {
		return fmt.Errorf("voting active")
	}

	if candidate == "undefined" {
		return fmt.Errorf("cannot remove undefined")
	}

	if !tools.StringInSlice(candidate, s.candidates) {
		return fmt.Errorf("candidate not found")
	}

	Position, err := tools.FindInSlice(s.candidates, candidate)

	if err != nil {
		return fmt.Errorf("candidate has no position")
	}

	PositionOfUndefined, err := tools.FindInSlice(s.candidates, "undefined")

	if err != nil {
		return fmt.Errorf("candidate has no position")
	}

	count := s.votesCounting[s.candidates[Position]]

	if 0 < count {
		s.votesCounting[s.candidates[PositionOfUndefined]] += count
	}

	s.candidates = tools.RemoveElementFromSlice(s.candidates, Position)
	return nil
}

func (s *Store) GetName() string {
	copyOfString := ""

	for _, character := range s.name {
		copyOfString += string(character)
	}
	return copyOfString
}

func (s *Store) SetName(password string, name string) error {
	if password != s.password {
		return fmt.Errorf("unknown password")
	}

  if s.name == name{
    return fmt.Errorf("want the set same name")
  }
	s.name = name
	return nil
}

func (s *Store) IDExists(id string) bool {
	for _, candidate := range s.votes {
		for _, vote := range candidate {
			if vote == id {
				return true
			}
		}
	}
	return false
}
