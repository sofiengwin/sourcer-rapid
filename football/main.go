package football

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

// 2 Champions League
// 39 EPL
const Base string = "https://api-football-v1.p.rapidapi.com/v3"

type PredictionTeam struct {
	Id     int
	Name   string
	Last_5 struct {
		Form  string
		Att   string
		Def   string
		Goals struct {
			For struct {
				Total   int
				Average string
			}
			Against struct {
				Total   int
				Average string
			}
		}
	}
}

type IResponse struct {
	Response []struct {
		Predictions struct {
			Winner struct {
				Id      int
				Name    string
				Comment string
			}
			Advice     string
			WinOrDraw  bool `json:"win_or_draw"`
			UnderCover bool `json:"under_cover"`
		}
		Teams struct {
			Home PredictionTeam
			Away PredictionTeam
		}
	}
}

type Fixture struct {
	Response []struct {
		Fixture struct {
			Id int
		}
		Teams struct {
			Home struct {
				Name string
			}
			Away struct {
				Name string
			}
		}
	}
}

func GetFixture() {
	ids := fixtures()
	fmt.Printf("%v", ids)

	for _, fixtureId := range ids {
		body := client(http.MethodGet, fmt.Sprintf("/predictions?fixture=%d", fixtureId))

		iresponse := IResponse{}
		// var dat map[string]interface{}

		if err := json.Unmarshal(body, &iresponse); err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			os.Exit(1)
		}
		// fmt.Printf("%s", body)
		// fmt.Println(iresponse.Response[0].Predictions.Winner.Comment)
		fmt.Printf("%s\n", iresponse)
	}
}

func fixtures() []int {
	body := client(http.MethodGet, "/fixtures?date=2023-11-04&league=39&season=2023")

	response := Fixture{}
	if err := json.Unmarshal(body, &response); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}

	fmt.Println(response.Response)
	fixtureIds := []int{}
	for _, fixture := range response.Response {
		fixtureIds = append(fixtureIds, fixture.Fixture.Id)
	}
	fmt.Println(fixtureIds)
	return fixtureIds
}

func (ir IResponse) String() string {
	pr := ir.Response[0]
	homeTeam := fmt.Sprintf("HomeTeam: %v Prediction: %v Comment: %v Form: %v Att: %v Def: %v ForTotal: %v ForAverage: %v AgainstTotal: %v AgainstAverage: %v",
		pr.Teams.Home.Name,
		pr.Predictions.Winner.Name,
		pr.Predictions.Winner.Comment,
		pr.Teams.Home.Last_5.Form,
		pr.Teams.Home.Last_5.Att,
		pr.Teams.Home.Last_5.Def,
		pr.Teams.Home.Last_5.Goals.For.Total,
		pr.Teams.Home.Last_5.Goals.For.Average,
		pr.Teams.Home.Last_5.Goals.Against.Total,
		pr.Teams.Home.Last_5.Goals.Against.Average,
	)
	awayTeam := fmt.Sprintf("AwayTeam: %v Prediction: %v Comment: %v Form: %v Att: %v Def: %v ForTotal: %v ForAverage: %v AgainstTotal: %v AgainstAverage: %v",
		pr.Teams.Away.Name,
		pr.Predictions.Winner.Name,
		pr.Predictions.Winner.Comment,
		pr.Teams.Away.Last_5.Form,
		pr.Teams.Away.Last_5.Att,
		pr.Teams.Away.Last_5.Def,
		pr.Teams.Away.Last_5.Goals.For.Total,
		pr.Teams.Away.Last_5.Goals.For.Average,
		pr.Teams.Away.Last_5.Goals.Against.Total,
		pr.Teams.Away.Last_5.Goals.Against.Average,
	)

	return fmt.Sprintf("%v \n %v\n", homeTeam, awayTeam)
}

func client(http_method string, endpoint string) []byte {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%v%v", Base, endpoint), nil)
	req.Header.Set("X-RapidAPI-Key", os.Getenv("X_RapidAPI_KEY"))

	if err != nil {
		fmt.Fprintf(os.Stderr, "error creating request: %v\n", err)
		os.Exit(1)
	}

	client := http.Client{}

	resp, err := client.Do(req)

	if err != nil {
		fmt.Fprintf(os.Stderr, "client fetch: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Request Status: %d\n", resp.StatusCode)
	body, err := io.ReadAll(resp.Body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading response body %v\n ", err)
	}

	return body
}
