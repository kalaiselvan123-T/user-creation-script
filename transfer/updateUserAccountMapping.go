package transfer

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type UserUpdateProtoRequest struct {
	Orgdomain      string
	OrganisationId string
	Email          string
	Token          string
	BundleId       string
	AccountId      string
	AccountDomain  string
}
type UpdateUserJsonrequest struct {
	AccountIdentifier struct {
		ID string `json:"id"`
	} `json:"account_identifier"`
	User struct {
		Email    string `json:"email"`
		Accounts []struct {
			ID               string `json:"id"`
			BundleIdentifier string `json:"bundle_identifier"`
		} `json:"accounts"`
	} `json:"user"`
	BundleIdentifier string `json:"bundle_identifier"`
}

func UpdateUserMapping(r *UserUpdateProtoRequest) {
	url := "https://" + r.Orgdomain + "/api/v2/users"
	method := "POST"
	var userupdaterequest UpdateUserJsonrequest
	userupdaterequest.AccountIdentifier.ID = r.AccountId
	userupdaterequest.User.Email = r.Email
	userupdaterequest.BundleIdentifier = r.BundleId
	userupdaterequest.User.Accounts = []struct {
		ID               string `json:"id"`
		BundleIdentifier string `json:"bundle_identifier"`
	}{
		{ID: r.AccountId, BundleIdentifier: r.BundleId},
	}
	jsonData, err := json.MarshalIndent(userupdaterequest, "", "  ")
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		return
	}
	payload := strings.NewReader(string(jsonData))

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+r.Token)

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer res.Body.Close()
	responebody, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}
	if res.StatusCode != 200 {
		fmt.Println("Error while making http call", string(responebody))
	} else if res.StatusCode != 200 {
		fmt.Println("Account mapping response", string(responebody))
	}

}
