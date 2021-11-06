package controllers

import (
	"net/http"
)

func PengaduanIndex(w http.ResponseWriter, r *http.Request)  {
	var viewFiles = components
	viewFiles = append(viewFiles, "views/components/footer.html")
	viewFiles = append(viewFiles, "views/pengaduan/index.html")
	view(w, "pengaduan.index", compact{}, viewFiles...)
}

func PengaduanCreate(w http.ResponseWriter, r *http.Request) {
	var viewFiles = components
	viewFiles = append(viewFiles, "views/components/footer.html")
	viewFiles = append(viewFiles, "views/pengaduan/create.html")
	// fmt.Println(viewFiles)
	view(w, "pengaduan.create", compact{}, viewFiles...)
}
