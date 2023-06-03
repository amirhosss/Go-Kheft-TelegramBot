package languages

import (
	"encoding/json"
	"log"
	"os"
)

type Languages struct {
	Messages struct {
		Default struct {
			Query    []string `json:"query"`
			Response []string `json:"response"`
			Btns     []struct {
				Text     string `json:"text"`
				Callback string `json:"callback"`
			}
		}
	}
}

var Response Languages

func init() {
	data, err := os.ReadFile("languages/fa.json")
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(data, &Response)
	if err != nil {
		log.Fatal(err)
	}
}
