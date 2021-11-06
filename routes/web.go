// Package router_web used to handle web routes
package router_web

import (
	"fmt"
	"net/http"
	"pengaduan/controllers"
	"strconv"
)

// Routes Handler Function
func Routes(portInt int) {
	http.HandleFunc("/", controllers.PengaduanIndex)
	http.HandleFunc("/create", controllers.PengaduanCreate)

	// Handle Static File
	http.Handle("/static/",
	  http.StripPrefix("/static/",
	    http.FileServer(
	      http.Dir("assets"))))

	port := ":" + strconv.Itoa(portInt)

	fmt.Println("sever started at localhost" + port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		fmt.Println(err)
	}
}
