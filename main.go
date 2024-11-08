package main

import (
    "encoding/json"
    "log"
    "math/rand"
    "net/http"
    "time"
)

type PasswordRequest struct {
    Length    int  `json:"length"`
    Upper     bool `json:"upper"`
    Lower     bool `json:"lower"`
    Numbers   bool `json:"numbers"`
    Symbols   bool `json:"symbols"`
}

type PasswordResponse struct {
    Password string `json:"password"`
}

func generatePassword(req PasswordRequest) string {
    var charset string

    if req.Upper {
        charset += "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
    }
    if req.Lower {
        charset += "abcdefghijklmnopqrstuvwxyz"
    }
    if req.Numbers {
        charset += "0123456789"
    }
    if req.Symbols {
        charset += "!@#$%^&*()_+-=[]{}|;:,.<>?"
    }

    if charset == "" {
        charset = "abcdefghijklmnopqrstuvwxyz" // fallback to lowercase if nothing selected
    }

    password := make([]byte, req.Length)
    for i := range password {
        password[i] = charset[rand.Intn(len(charset))]
    }

    return string(password)
}

func handleGenerate(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }

    var req PasswordRequest
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    if req.Length < 7 || req.Length > 100 {
        http.Error(w, "Password length must be between 7 and 100", http.StatusBadRequest)
        return
    }

    password := generatePassword(req)

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(PasswordResponse{Password: password})
}

func main() {
    rand.Seed(time.Now().UnixNano())

    fs := http.FileServer(http.Dir("static"))
    http.Handle("/", fs)
    http.HandleFunc("/generate", handleGenerate)

    log.Println("Server starting on :80")
    if err := http.ListenAndServe(":80", nil); err != nil {
        log.Fatal(err)
    }
}
