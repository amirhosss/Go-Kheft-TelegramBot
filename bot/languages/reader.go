package languages

import (
	"encoding/json"
	"log"
	"os"
)

type Messages struct {
	Messages struct {
		NonMember struct {
			Response []string `json:"response"`
			Btns     []string `json:"btns"`
			Failed   string   `json:"failed"`
		}
		Member struct {
			Response []string `json:"response"`
			Btns     []string `json:"btns"`
			Failed   string   `json:"failed"`
		}
	}
	Conversations struct {
		Exit struct {
			Query string `json:"query"`
		}
		Registration struct {
			Query    string   `json:"query"`
			Response []string `json:"response"`
			Btns     []string `json:"btns"`
		}
		Rules struct {
			Query    string   `json:"query"`
			Response []string `json:"response"`
			Failed   string   `json:"failed"`
		}
		Username struct {
			Response []string `json:"response"`
		}
		Price struct {
			Response []string `json:"response"`
			Failed   string   `json:"failed"`
		}
		Advertise struct {
			Response    []string `json:"response"`
			Failed      string   `json:"failed"`
			FailedLimit string   `json:"failedLimit"`
		}
		Description struct {
			Response []string `json:"response"`
		}
	}
}

var Response Messages

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
