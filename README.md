
## User creation script

Script to manage to create users in both organisation and accounts in bundle

### Clone the project
```
$ git clone https://github.com/kalaiselvan123-T/user-creation-script
```

#### Script setup

1. `cd user-creation-script` into the respective directory
2. Run `go mod init <SOME-NAME>` to initialize a module
3. Run `go mod tidy` to add required imports

### Fill the below defined variables in `main.go` 
```go
orgdomain := "{org_domain}"
organisation_id := "{org_id}"
token := "{token}"
bundle_id := "{bundle_id}"
name := "{user_name}"
email := "{user_email_prefix}"
count := 1 
admin := false
FreshService_product_id := "{product_id1}"
FreshRelease_product_id := "{product_id2}"
``` 

Note:
1. If the user needs to be created only in the organization, 

   `comment the below lines in the main.go`

```go
/* account creation */

accountreq := transfer.AccountCreationStructRequest{
    OrgDomain: orgdomain,
    OrganisationId: organisation_id,
    ProductId: FreshService_product_id,
    BundleId: bundle_id,
    Domain: "anchorAccount",
    Anchor: true,
    Token: token,
}
seederaccountreq := transfer.AccountCreationStructRequest{
    OrgDomain: orgdomain,
    OrganisationId: organisation_id,
    ProductId: FreshRelease_product_id,
    BundleId: bundle_id,
    Domain: "seederAccount",
    Anchor: false,
    Token: token,
}
account_id1, domain1 := transfer.CreateStandaloneAccount(&accountreq)
account_id2, domain2 := transfer.CreateStandaloneAccount(&seederaccountreq)

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
```

### Running the script 
To run the script, use the following command in the terminal:

```
go run <SOME-NAME>.go
```
