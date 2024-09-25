package main

import (
	"fmt"
	"main/transfer"
	"sync"
)

func main() {

	orgdomain := "{org_domain}"
	organisation_id := "{org_id}"
	token := "{token}"
	bundle_id := "{bunlde_id}" // if user need to be mapped in bundle accounts
	name := "{user_name}"
	email := "{user_email_prefix}"
	count := 1 // no of user need to created for each thread
	admin := false
	// Here we have created a FS-Fr bundle , so included the respective product's id to create respective accounts in the given bundle
	FreshService_product_id := "{product_id1}"
	FreshRelease_product_id := "{product_id2}"

	var accountslist []transfer.AccountListProtoRequest

	/* account creation */

	accountreq := transfer.AccountCreationStructRequest{OrgDomain: orgdomain, OrganisationId: organisation_id, ProductId: FreshService_product_id, BundleId: bundle_id, Domain: "anchorAccount", Anchor: true, Token: token}
	seederaccountreq := transfer.AccountCreationStructRequest{OrgDomain: orgdomain, OrganisationId: organisation_id, ProductId: FreshRelease_product_id, BundleId: bundle_id, Domain: "seederAccount", Anchor: false, Token: token}
	account_id1, domain1 := transfer.CreateStandaloneAcount(&accountreq)
	account_id2, domain2 := transfer.CreateStandaloneAcount(&seederaccountreq)

	accountslist = append(accountslist, transfer.AccountListProtoRequest{
		ID:               account_id1,
		Domain:           domain1,
		BundleIdentifier: bundle_id,
	})
	accountslist = append(accountslist, transfer.AccountListProtoRequest{
		ID:               account_id2,
		Domain:           domain2,
		BundleIdentifier: bundle_id,
	})

	// creating users with multiple go routines
	users := transfer.UserCreationProtoRequest{Orgdomain: orgdomain, OrganisationId: organisation_id, Name: name, Email: email, Usercount: count, Token: token, BundleId: bundle_id, AccountsLists: accountslist, Admin: admin}
	var wg sync.WaitGroup

	N := 100 // Number of concurrent threads needed
	var number int
	wg.Add(N)

	for i := 0; i < N; i++ {
		go func() {
			defer wg.Done()
			// Call createUser method with the user template
			responseString, statusCode := transfer.CreateUserForBundle(&users)
			number++
			fmt.Printf("Response String: %s, Status Code: %d\n", responseString, statusCode)
		}()
	}

	wg.Wait()

	fmt.Println("number call were made", number)
}
