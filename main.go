package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"
)

type Note struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Body  string `json:"body"`
}

type Response struct {
	Code    int         `json:"code"`
	Status  string      `json:"status"`
	Data    interface{} `json:"data,omitempty"` // omitempty -> datanya tidak akan tampil, kalau nil
	Message string      `json:"message"`
}

var (
	notes = []Note{}
)

func write(w http.ResponseWriter, code int, message string, status string, data interface{}) {
	var response = Response{
		Code:    code,
		Message: message,
		Status:  status,
		Data:    data,
	}
	resp, _ := json.Marshal(response)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(resp)
}

func remove(slice []Note, s int) []Note {
	return append(slice[:s], slice[s+1:]...)
}

func main() {

	handler := http.NewServeMux()
	baseURl := "localhost"
	port := "8000"

	handler.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Pong"))
	})

	handler.HandleFunc("/api/v1/notes", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			// Create Note
			rBody, err := io.ReadAll(r.Body)
			if err != nil {
				fmt.Println("error :", err)
				write(w, http.StatusInternalServerError, "Sistem sedang sibuk", "error", nil)
				return
			}

			var note Note
			err = json.Unmarshal(rBody, &note)
			if err != nil {
				fmt.Println("error :", err)
				write(w, http.StatusInternalServerError, "Sistem sedang sibuk", "error", nil)
				return
			}

			if note.Title == "" || note.Body == "" {
				fmt.Println("error :", "title/body is null")
				write(w, http.StatusBadRequest, "Salah input", "error", nil)
				return
			}

			note.ID = len(notes) + 1

			notes = append(notes, note)

			write(w, http.StatusCreated, "Note Baru Berhasil Ditambahkan", "Success", nil)
			return
		}

		if r.Method == http.MethodGet {
			write(w, http.StatusOK, "Success get list notes", "Success", notes)
			return
		}

		if r.Method == http.MethodPut {
			// get parameter dari url
			pID := r.URL.Query().Get("id")
			id, err := strconv.Atoi(pID)
			if err != nil {
				fmt.Println("error :", err)
				write(w, http.StatusBadRequest, "salah parameter", "error", nil)
				return
			}

			// Update Note
			rBody, err := io.ReadAll(r.Body)
			if err != nil {
				fmt.Println("error :", err)
				write(w, http.StatusInternalServerError, "Sistem sedang sibuk", "error", nil)
				return
			}

			var note Note
			err = json.Unmarshal(rBody, &note)
			if err != nil {
				fmt.Println("error :", err)
				write(w, http.StatusInternalServerError, "Sistem sedang sibuk", "error", nil)
				return
			}

			if note.Title == "" || note.Body == "" {
				fmt.Println("error :", "title/body is null")
				write(w, http.StatusBadRequest, "Salah input", "error", nil)
				return
			}

			for i, oldNote := range notes {
				if oldNote.ID == id {
					notes[i].Title = note.Title
					notes[i].Body = note.Body
				}
			}

			write(w, http.StatusCreated, "Note Berhasil Diubah", "Success", nil)
			return
		}

		if r.Method == http.MethodDelete {
			// get parameter dari url
			pID := r.URL.Query().Get("id")
			id, err := strconv.Atoi(pID)
			if err != nil {
				fmt.Println("error :", err)
				write(w, http.StatusBadRequest, "salah parameter", "error", nil)
				return
			}

			for i, oldNote := range notes {
				if oldNote.ID == id {
					notes = remove(notes, i)
				}
			}

			write(w, http.StatusCreated, "Note Berhasil Dihapus", "Success", nil)
			return
		}

		write(w, http.StatusMethodNotAllowed, "Method Not Allowed", "error", nil)
		return
	})

	addr := fmt.Sprintf("%s:%s", baseURl, port)
	server := &http.Server{
		Addr:         addr,
		Handler:      handler,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	fmt.Printf("server run on -> %s ", addr)
	server.ListenAndServe()
}
