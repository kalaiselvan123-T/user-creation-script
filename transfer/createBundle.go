package transfer

type BundleCreationRequest struct {
	Orgdomain       string
	Organisation_id string
	Token           string
	Bundle_type     string
}

// func CreateBundle(bundleCreationRequst *BundleCreationRequest) string {
// 	url := "https://" + bundleCreationRequst.Orgdomain + "/api/v2/bundles"
// 	method := "POST"

// 	type BundleCreationHttpRequest struct {
// 		organisation_id string
// 		token           string
// 		bundle_type     string
// 	}
// 	payload := string("kalai")
// 	client := &http.Client{}
// 	req, err := http.NewRequest(method, url)

// 	if err != nil {
// 		fmt.Println(err)
// 		return ""
// 	}
// 	req.Header.Add("Content-Type", "application/json")

// 	res, err := client.Do(req)
// 	if err != nil {
// 		fmt.Println(err)
// 		return ""
// 	}
// 	defer res.Body.Close()

// 	body, err := io.ReadAll(res.Body)
// 	if err != nil {
// 		fmt.Println(err)
// 		return ""
// 	}
// 	fmt.Println(body)
// 	return ""
// }
