package main
import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "log"
    "net/http"
    "os"
)
type AboutYou struct {
    FirstName    string `json:"firstName"`
    LastName     string `json:"lastName"`
    DOB          string `json:"dob"`
    EmailAddress string `json:"emailAddress"`
}
func aboutYouHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
    w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
    if r.Method == http.MethodOptions {
        w.WriteHeader(http.StatusOK)
        return
    }
    if r.Method != http.MethodPost {
        http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
        return
    }
    body, err := ioutil.ReadAll(r.Body)
    if err != nil {
        http.Error(w, "Failed to read request body", http.StatusBadRequest)
        return
    }
    var aboutYou AboutYou
    err = json.Unmarshal(body, &aboutYou)
    if err != nil {
        http.Error(w, "Failed to parse JSON", http.StatusBadRequest)
        return
    }
    fileName := fmt.Sprintf("collected_data/%s_%s.json", aboutYou.FirstName, aboutYou.LastName)
    file, err := os.Create(fileName)
    if err != nil {
        http.Error(w, "Failed to write to file", http.StatusInternalServerError)
        return
    }
    defer file.Close()
    file.Write(body)
    w.WriteHeader(http.StatusOK)
    w.Write([]byte("Data saved successfully"))
}
func main() {
    http.HandleFunc("/v1/aboutYou", aboutYouHandler)
    fmt.Println("Starting server on port 9090...")
    log.Fatal(http.ListenAndServe(":9090", nil))
}
