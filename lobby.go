package main

import "errors"

const maxPlayersInLobby = 1
const maxLobbies = 1024

var lobby [][]string
var userInLobby = make(map[string] bool)

func addNewLobby() error {
	if (len(lobby) == maxLobbies) {
		return errors.New("Max lobbies reached")
	}

	lobby = append(lobby, []string {})
	return nil
}

func joinLobby(username string) error {
	// for now we just adding username to array until array is full.
	// Then creating new sub array
	
	if userInLobby[username] {
		return errors.New("Already in lobby")
	}

	if len(lobby) == 0 {
		addNewLobby()
	}
	lobbiesCount := len(lobby)
	playersInLobby := len(lobby[lobbiesCount-1])
	if playersInLobby == maxPlayersInLobby {
		if err := addNewLobby(); err != nil {
			return err
		}
	}
	lobby[lobbiesCount-1] = append(lobby[lobbiesCount-1], username)
	userInLobby[username] = true
	return nil
}