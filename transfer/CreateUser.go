package transfer

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/google/uuid"
)

type AccountListProtoRequest struct {
	ID               string
	Domain           string
	BundleIdentifier string
	Anchor           bool
}

type UserCreationProtoRequest struct {
	Orgdomain      string
	OrganisationId string
	Name           string
	Email          string
	Usercount      int
	Token          string
	BundleId       string
	Admin          bool
	AccountsLists  []AccountListProtoRequest
}

type AccountsList []struct {
	ID               string `json:"id"`
	Domain           string `json:"domain"`
	BundleIdentifier string `json:"bundle_identifier,omitempty"`
	Anchor           bool   `json:"anchor,omitempty"`
}

type UserJsonRequest struct {
	FirstName string       `json:"first_name"`
	Email     string       `json:"email"`
	Admin     bool         `json:"admin,omitempty"`
	Accounts  AccountsList `json:"accounts,omitempty"`
}

type UserCreationJsonRequest struct {
	User             UserJsonRequest `json:"user"`
	BundleIdentifier string          `json:"bundle_identifier"`
}

func CreateUserForBundle(req *UserCreationProtoRequest) (string, int) {
	var number int
	url := "https://" + req.Orgdomain + "/api/v2/users"
	method := "POST"
	var statusCode int
	i := 0
	for i < req.Usercount {
		var accountsList AccountsList
		for _, v := range req.AccountsLists {
			accountsList = append(accountsList, struct {
				ID               string "json:\"id\""
				Domain           string "json:\"domain\""
				BundleIdentifier string "json:\"bundle_identifier,omitempty\""
				Anchor           bool   "json:\"anchor,omitempty\""
			}{ID: v.ID, Domain: v.Domain, BundleIdentifier: v.BundleIdentifier, Anchor: v.Anchor})
		}
		builder := strings.Builder{}
		builder.WriteString(req.Email)
		builder.WriteString(uuid.New().String())
		builder.WriteString("@gmail.com")
		userreq := UserJsonRequest{FirstName: req.Name, Email: builder.String(), Accounts: accountsList, Admin: req.Admin}
		userRequest := UserCreationJsonRequest{User: userreq, BundleIdentifier: req.BundleId}
		fmt.Println("user request:", userRequest)
		jsonbody, _ := json.Marshal(userRequest)
		payload := strings.NewReader(string(jsonbody))

		client := &http.Client{}
		reqs, err := http.NewRequest(method, url, payload)
		reqs.Header.Add("Content-Type", "application/json")
		reqs.Header.Add("Authorization", "Bearer "+req.Token)

		if err != nil {
			fmt.Println(err)
			return "", 0
		}
		res, err := client.Do(reqs)
		fmt.Println(res)
		if err != nil {
			fmt.Println(err)
			fmt.Println("Errored http request", reqs)
			return "", 0

		} else if res.StatusCode == 200 {
			number++
			i++
			// if user created in the org , calling create User API to add account mappings for anchor and seeder accounts
			statusCode = res.StatusCode
			for _, v := range req.AccountsLists {
				userRequest := UserUpdateProtoRequest{Orgdomain: req.Orgdomain, OrganisationId: req.OrganisationId, Email: builder.String(), Token: req.Token, BundleId: req.BundleId, AccountId: v.ID, AccountDomain: v.Domain}
				UpdateUserMapping(&userRequest)
			}

			fmt.Println("i:", i)
		}
		fmt.Println(res.StatusCode)
		defer res.Body.Close()
	}
	fmt.Println("user creation for loop call", number)
	return "user created with", statusCode
}
