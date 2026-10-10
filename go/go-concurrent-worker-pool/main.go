package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"example.com/go-concurrent-worker-pool/jobprocessor"
)
type JobProcessor = jobprocessor.JobProcessor

type Server struct {
	jobProcessor *JobProcessor
}

func (server *Server) NewServer(jp *JobProcessor) *Server{
	server.jobProcessor = jp
	return server
}

func (server *Server) handleSubmitJob(w http.ResponseWriter, r *http.Request){
	if r.Method != http.MethodPost {
		http.Error(w,"Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req jobprocessor.Job
	err:= json.NewDecoder(r.Body).Decode(&req)
	if err != nil{
		http.Error(w, "Bad request payload", http.StatusBadRequest)
		return
	}

	


}

func main() {
	// test()
	// test1()
	// test2()
	// daisyChain()
	// PingPong()
	// contextMain()
	
	server:= &Server{jobProcessor: jobprocessor.NewJobProcessor(100)}
	mux := http.NewServeMux()
	mux.HandleFunc("/submit_job", server.handleSubmitJob)
	fmt.Println("Server starting on :8080")
	if err :=http.ListenAndServe(":8080",mux); err != nil{
		panic(err)
	}
	os.Exit(0)
}
