package languages

import (
	"encoding/json"
	"log"
	"os"
)

type Languages struct {
	Messages struct {
		NonMember struct {
			Query    string   `json:"query"`
			Response []string `json:"response"`
			Btns     []string `json:"btns"`
			Failed   string   `json:"failed"`
		}
		Member struct {
			Query    string   `json:"query"`
			Response []string `json:"response"`
			Btns     []string `json:"btns"`
			Failed   string   `json:"failed"`
		}
		Registration struct {
			Query    string   `json:"query"`
			Response []string `json:"response"`
			Btns     []string `json:"btns"`
		}
		Rules struct {
			Query    string   `json:"query"`
			Response []string `json:"response"`
		}
	}
}

var Response Languages

func init() {
	data, err := os.ReadFile("bot/languages/fa.json")
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(data, &Response)
	if err != nil {
		log.Fatal(err)
	}
}
