package main
import (
    "net/http"
    "io"
    // "io/ioutil"
    // "fmt"
    "log"
)

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        resp, err := http.Get("https://drive.google.com/uc?export=download&id=1u2NNJK-2pCe8_8PLc3k8eNJYtHsp5X0j")
        check(err)
        defer resp.Body.Close()
        for k, v := range resp.Header {
            for _, vv := range v {
                w.Header().Set(k, vv)
            }
        }
        io.Copy(w, resp.Body)
    })
    http.ListenAndServe(":8002", nil)
}

func check (err error) {
    if err != nil {
        log.Fatal(err)
    }
}
