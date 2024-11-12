# Voting Server

Query

- get Candidates
  - /getCandidates
  - Response: ```{"candidates": []string}```
- ping
  - /ping
  - Response: OK
- show Voting
  - /showVoting
  - Response: ``` {
      "Undefined":  int,
      "Candidate1": int,
      "Candidate2": int,
      "Candidate3": int,
      "Candidate4": int,
      "Candidate5": int,
    }```
- is voting Active
  - /votingActive
  - Response:```
    {
      "votingActive": bool
    }
  ```
Commands
- activate Voting
  - /activateVoting
  - Request: ```{"password": string}```
  - Response: ```{"votingActive": bool}```
- deactivate Voting
  - /deactivateVoting
  - Request: ```{"password": string}```
  - Response: ```{"votingActive": bool}```
- delete all entries
  - /deleteAll
  - Request: ```{"password": string}```
  - Response: OK / 500 (Internal Server Error)
- insert New Votable
  - /insertNewVotable
  - Request: ```{"password": string, "vote-id": string}```
  - Response: OK / 500 (Internal Server Error)
- make vote
  - /makeVote
  - Request: ```{"id": string, "candidate": string}```
  - Response: OK / 500 (Internal Server Error)
