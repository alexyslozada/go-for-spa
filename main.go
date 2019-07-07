package main

import (
	"log"
	"net/http"
	"os"
	"path"
	"strings"

	"github.com/alexyslozada/config-go"
)

func main() {
	cnfg, err := config.New("./config.json")
	if err != nil {
		log.Printf("no se pudo leer el archivo de configuración: %v", err)
		os.Exit(1)
	}

	port, err := cnfg.Get("port")
	if err != nil {
		log.Printf("no se pudo leer el puerto: %v", err)
		os.Exit(1)
	}

	publicDir, err := cnfg.Get("public_dir")
	if err != nil {
		log.Printf("no se pudo leer el directorio: %v", err)
		os.Exit(1)
	}

	index, err := cnfg.Get("index")
	if err != nil {
		log.Printf("no se pudo leer el index: %v", err)
		os.Exit(1)
	}

	http.Handle("/", handler(publicDir, index))
	http.HandleFunc("/health", health)

	log.Printf("Servidor iniciado en el puerto: %s...\n", port)
	err = http.ListenAndServe(":" + port, nil)
	if err != nil {
		log.Printf("error al subir el servidor: %v", err)
		os.Exit(1)
	}
}

func handler(publicDir, index string) http.Handler {
	handler := http.FileServer(http.Dir(publicDir))

	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		_path := req.URL.Path

		if strings.Contains(_path, ".") || _path == "/" {
			handler.ServeHTTP(w, req)
			return
		}

		http.ServeFile(w, req, path.Join(publicDir, "/" + index))
	})
}

func health(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}