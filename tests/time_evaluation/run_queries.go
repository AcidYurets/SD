package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	grqphqlEndpoint = "http://localhost:4040/graphql"
	token           = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2ODU5Mzg0MTgsImlhdCI6MTY4NDEzODQxOCwianRpIjoiOTk0N2I0MzctZmE0Zi00ODkwLThjMzAtYzY3Mjk4NWM3MDUwIiwic2Vzc2lvbiI6eyJTSUQiOiI3MWI5ZTRjNy1kMTg3LTRmZjYtYjdhNS03NzNlNmMzNGIyNDMiLCJVc2VyVXVpZCI6ImM2MzVlNTA1LWY5NTctNGUzMy1iMDY1LTY5NGU5Y2RmNGU1YyIsIlJvbGUiOiJhZG1pbiJ9fQ.FC90LPmzzDzl5SjthPJCEUKyyjoyfoUc1JelUnm1dUQ"
	requestedEvents = 2000
)

func main() {
	// Define the GraphQL query as a string
	query := `
query SearchEvent($paginate: PaginationInput!, $filter: EventFilter, $sort: EventSort) {
  SearchEvent(paginate: $paginate, filter: $filter, sort: $sort) {
    Uuid
    Name
    Description
    IsWholeDay
    Creator {
      Login
      Phone
      Uuid
    }
    Type
    Timestamp
    Invitations {
      Uuid
      AccessRight {
        Code
        Description
      }
      User {
        Login
        Phone
        Role
        Uuid
      }
    }
  }
}
`

	variables := fmt.Sprintf(`
{
  "paginate": {
    "Page": 1,
    "PageSize": %d
  },
  "filter":{
    "FTSearchStr": {
      "Str": "Соб"
    },
    "CreatorLogin": {
      "Ts": "Us"
    },
    "Description": {
      "Ts": "Опи"
    },
    "TagName": {
      "Ts": "Тег"
    }
  },
  "sort": {
    "CreatorLogin": "ASC"
  }
}
`, requestedEvents)

	// Define the request body as a JSON object
	requestBody := map[string]interface{}{
		"query":         query,
		"variables":     json.RawMessage(variables),
		"operationName": "SearchEvent",
	}

	requestBodyBytes, err := json.Marshal(requestBody)
	if err != nil {
		panic(err)
	}

	req, err := http.NewRequest("POST", grqphqlEndpoint, bytes.NewBuffer(requestBodyBytes))
	if err != nil {
		panic(err)
	}

	// Set headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", token)

	// Send the request
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println(resp.StatusCode)

	// Print response
	//body, err := io.ReadAll(resp.Body)
	//if err != nil {
	//	panic(err)
	//}
	//
	//fmt.Println(string(body))
}
