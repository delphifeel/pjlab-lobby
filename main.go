package main

import (
	"log"
	"net/http"
	"fmt"
	"io"
)

func testToken(username string, tokenB64 string) bool {
	resp, err := http.Get(
		fmt.Sprintf("http://localhost:8090/testToken?username=%s&tokenB64=%s", username, tokenB64),
	)
	if err != nil {
		log.Println(err)
		return false
	}

	if resp.StatusCode > 299 {
		log.Printf("ERROR: got %v from pjlab_auth\n", resp.StatusCode)
		return false
	}

	body, err := io.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		log.Println(err)
		return false
	}

	return string(body) == "Yes"
}

func joinHandler(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Connection", "close")

	username := req.FormValue("username")
	if username == "" {
		http.Error(w, "No username", 400)
		return
	}
	tokenB64 := req.FormValue("tokenB64")
	if tokenB64 == "" {
		http.Error(w, "No token", 400)
		return
	}

	if ok := testToken(username, tokenB64); !ok {
		http.Error(w, "Wrong token", 403)
		return
	}

	if err := joinLobby(username); err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	fmt.Fprintf(w, "Yes")
}

func init() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
}

func main() {
	http.HandleFunc("/join", joinHandler)
	log.Println("pjlab_lobby service started.")

	if err := http.ListenAndServe(":8091", nil); err != nil {
		log.Fatal(err)
	}
}