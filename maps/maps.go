package main

import "fmt"

func main() {
	// slices (index array)
	//websites := []string{"google.com", "facebook.com", "twitter.com"}

	// array with key value pairs
	websites := map[string]string{
		"Google": "https://www.google.com",
		"AWS":    "https://aws.com",
		"Azure":  "https://azure.com",
	}
	fmt.Println(websites)
	fmt.Println(websites["AWS"])
	websites["linkedin"] = "https://www.linkedin.com" // add a new item
	fmt.Println(websites)

	delete(websites, "Azure") // delete a key value pair
	fmt.Println(websites)
}
