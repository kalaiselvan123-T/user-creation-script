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
// 	req.Header.Add("Authorization", "Bearer eyJraWQiOiIxMTQ2MzI4NTY1MzMyNzA1MjkiLCJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiJ9.eyJhdWQiOiI2MDg5MzgwOTUzNjUxMzc2MDkiLCJzdWIiOiI2MjYxMDE4NTc4MjMwOTk3NTkiLCJvcmdhbmlzYXRpb25faWQiOiI2MjYxMDE4MTUyNTAxNzg3NjEiLCJzY29wZSI6WyJST0xFX0NMSUVOVCJdLCJwcm9kdWN0X2lkIjoiNjA3NTA0MzMxMzUyOTE1OTY5IiwiaXNzIjoiaHR0cHM6Ly9wc3RtdGVzdC5waG9lbml4LmZyZXNod29ya3NhcGkuaW8iLCJvcmdhbmlzYXRpb25fZG9tYWluIjoicHN0bXRlc3QucGhvZW5peC5mcmVzaHdvcmtzYXBpLmlvIiwiZXhwIjoxNzI3MTk3NTE1LCJpYXQiOjE2OTU1NzUxMTUsImp0aSI6IjYyNjEwMTg1NzgyMzA5OTc2MCJ9.UUG46nEouZd97vyTwmrClrjrA7E4lE9-c5ojeEpQzxb_iJtWlczhVI8RkmidInBu4mOCQPXr8zQTHCmNdEQeRb8einkpN5DZEb6lJ1EMdl86CPgQTuMTC5qlIp6gUM0ZppXs-kvr03XHERjnuDIGt1faEAUMPeP_Vqj8e91x_foKLy7Qhr0W9shVlXFiitJnNAMPtmtPov-ev2YyKqHzi5C8crDSiJTdEN64huE9a6xM6SjUwGpRyXN5z5Xju9XHZkvWIP5Q2w732dkIInUZaiW9w-NoF9s5jzzHMjmiXlffbbkY5H64Ry0QfVct9b1Y_Ljuj7H6GZE4NrAXrB-8KQ")

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
