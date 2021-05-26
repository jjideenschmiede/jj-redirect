//**********************************************************
//
// Copyright (C) 2018 - 2021 J&J Ideenschmiede UG (haftungsbeschränkt) <info@jj-ideenschmiede.de>
//
// This file is part of jj-redirect.
// All code may be used. Feel free and maybe code something better.
//
// Author: Jonas Kwiedor
//
//**********************************************************

package main

import (
	"net/http"
	"os"
)

func main() {

	// Get all requests
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {

		// Redirect to target
		http.Redirect(writer, request, os.Getenv("REDIRECT_TARGET"), http.StatusSeeOther)

	})

	// Run server
	http.ListenAndServe(":80", nil)
}
