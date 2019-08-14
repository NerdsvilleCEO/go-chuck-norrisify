package utils

import "net/http"

// LimitNumClients is HTTP handling middleware that ensures no more than
// maxClients requests are passed concurrently to the given handler f.
func LimitNumClients(f http.HandlerFunc, maxClients int) http.HandlerFunc {
  // Counting semaphore using a buffered channel
  sema := make(chan struct{}, maxClients)

  return func(w http.ResponseWriter, req *http.Request) {
    sema <- struct{}{}
    defer func() { <-sema }()
    f(w, req)
  }
}
