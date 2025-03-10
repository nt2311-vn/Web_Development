package main

import (
	"fmt"
	"net/http"
)

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
