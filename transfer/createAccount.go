package transfer

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/google/uuid"
)

type AccountCreationStructRequest struct {
	OrgDomain      string
	OrganisationId string
	ProductId      string
	BundleId       string
	Domain         string
	Token          string
	Anchor         bool
}

type Account struct {
	OrganisationID   string `json:"organisation_id"`
	ProductID        string `json:"product_id"`
	Domain           string `json:"domain"`
	BundleIdentifier string `json:"bundle_identifier,omitempty"`
	Anchor           bool   `json:"anchor,omitempty"`
}
type AccountCreationJsonRequest struct {
	Account Account `json:"account"`
}
type AccountresponseStruct struct {
	Account struct {
		ID     string `json:"id"`
		Domain string `json:"domain"`
	} `json:"account"`
}

func CreateStandaloneAcount(r *AccountCreationStructRequest) (string, string) {
	url := "https://" + r.OrgDomain + "/api/v2/accounts"
	method := "POST"
	accountDomainwithUUID := r.Domain + uuid.New().String()
	var accountrequest Account

	if len(r.BundleId) > 0 {
		accountrequest = Account{OrganisationID: r.OrganisationId, ProductID: r.ProductId, Domain: accountDomainwithUUID, BundleIdentifier: r.BundleId, Anchor: r.Anchor}

	} else {
		accountrequest = Account{OrganisationID: r.OrganisationId, ProductID: r.ProductId, Domain: accountDomainwithUUID}

	}

	account := AccountCreationJsonRequest{accountrequest}

	jsonbody, err := json.Marshal(account)
	if err != nil {
		// Handle the error appropriately
		fmt.Println("Error marshaling body:", err)
	}
	payload := strings.NewReader(string(jsonbody))

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return "", ""
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+r.Token)

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return "", ""
	}
	defer res.Body.Close()

	responebody, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return "", ""
	}
	var out AccountresponseStruct
	if err := json.Unmarshal(responebody, &out); err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(responebody))

	return out.Account.ID, out.Account.Domain
}
