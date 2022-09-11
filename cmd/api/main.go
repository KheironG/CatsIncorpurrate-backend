package main

import (
    "flag"
    "fmt"
    "log"
    "net/http"
    "encoding/json"
    "io/ioutil"
)

const version = "1.0.0";

type config struct {
    port int
    env string
}

type Cats struct {
	Cats []Cat `json:"cats"`
}

type Cat struct {
	Name string `json:"name"`
	Image string `json:"image"`
    CutenessLevel int `json:"cutenessLevel"`
    AllergyInducingFur bool `json:"allergyInducingFur"`
    LivesLeft int `json:"livesLeft"`
}

func main() {

    var cfg config

    flag.IntVar( &cfg.port, "port", 4000, "server port to listen on" )
    flag.StringVar( &cfg.env, "env", "development", "Application environment" )
    flag.Parse()

    fmt.Println("Running")

    http.HandleFunc( "/cats", func(w http.ResponseWriter, r *http.Request ) {

        file, _ := ioutil.ReadFile("catdata.json")

	    data := Cats{}

	    _ = json.Unmarshal([]byte(file), &data)

        js, err := json.MarshalIndent(data, "", "\t")
        if err != nil {
            log.Println(err)
        }

        w.Header().Set("Content-Type", "application/json")
        w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000" )
        w.WriteHeader(http.StatusOK)
        w.Write(js)

    })

    err := http.ListenAndServe(fmt.Sprintf(":%d", cfg.port), nil)
    if err != nil {
        log.Println(err)
    }

}
