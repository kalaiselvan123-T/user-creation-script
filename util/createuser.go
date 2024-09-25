package util

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
)

type CommonRequest struct {
	orgdomain       string
	organisation_id string
	product_id      string
	domain          string
	account_id      string
	name            string
	email           string
	usercount       int
	token           string
	bundle_id       string
	vari            int
}
type AccountresponseStruct struct {
	Account struct {
		ID string `json:"id"`
	} `json:"account"`
}

func MapUsertosatandaloneaccount(a *CommonRequest) (int, error) {

	//variables
	orgdomain := a.orgdomain
	org_id := a.organisation_id
	id := a.account_id
	userCount := a.usercount
	var j int = 0

	var i int = a.vari

	// api
	url := "https://" + orgdomain + "/api/v2/accounts/" + id + "/users"
	method := "POST"

	type userCreationrequest struct {
		Organisation_id string `json:"organisation_id"`
		Name            string `json:"first_name"`
		Email           string `json:"email"`
		Admin           bool   `json:"admin"`
	}

	type requestBody struct {
		User userCreationrequest `json:"user"`
	}

	for j < userCount {
		builder := strings.Builder{}
		builder.WriteString(a.email)
		builder.WriteString(strconv.Itoa(i))
		builder.WriteString("@gmail.com")

		user := userCreationrequest{Organisation_id: org_id, Email: builder.String(), Name: a.name, Admin: true}
		body := requestBody{User: user}
		jsonbody, err := json.Marshal(body)
		if err != nil {
			// Handle the error appropriately
			fmt.Println("Error marshaling body:", err)
			return j, fmt.Errorf("error marshaling body: %w", err)

		}
		payload := strings.NewReader(string(jsonbody))

		client := &http.Client{}
		req, err := http.NewRequest(method, url, payload)

		if err != nil {
			fmt.Println(err)
			return j, fmt.Errorf("error creating request: %w", err)
		}
		req.Header.Add("Content-Type", "application/json")
		req.Header.Add("Authorization", "Bearer "+a.token)
		res, err := client.Do(req)
		if err != nil {
			fmt.Println(err)
			return j, fmt.Errorf("error making HTTP request: %w", err)
		} else if res.StatusCode == 200 {
			i++
			j++

		} else {
			i++
			if i == userCount*userCount {
				break
			}
		}
		defer res.Body.Close()

		responebody, err := io.ReadAll(res.Body)
		if err != nil {
			fmt.Println(err)
			return j, fmt.Errorf("HTTP request failed with status: %s", res.Status)
		}
		var out AccountresponseStruct
		if err := json.Unmarshal(responebody, &out); err != nil {
			fmt.Println(err)
			return j, err
		}
		fmt.Println(string(responebody))

	}
	return j, nil
}
