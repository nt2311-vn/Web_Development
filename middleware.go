package main

import (
	"fmt"
	"net/http"
)

type AfterMiddleware struct {
	handler http.Handler
}

func middleware1(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Executing middleware1()...")
		next.ServeHTTP(w, r)
		fmt.Fprintln(w, "Executing middleware1() again...")
	})
}

func middleware2(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Executing middleware2()...")
		next.ServeHTTP(w, r)
		fmt.Fprintln(w, "Executing middleware2() again...")
	})
}

func final(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprintln(w, "Executing final()...")
	fmt.Fprintln(w, "Done")
}

func (a *AfterMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	a.handler.ServeHTTP(w, r)
	w.Write([]byte(" +++ Hello from middleware! +++ "))
}

func myHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(" *** Hello from myHandler! *** "))
}
