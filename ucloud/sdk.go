package ucloud

import "fmt"
import "net/http"
import "io/ioutil"
import "crypto/sha1"
import "sort"
import "encoding/json"
import "bytes"

func VerfyAc(params map[string]string, private_key string) string {
	params_data := ""

	sorted_keys := make([]string, 0)
	for key, _ := range params {
		sorted_keys = append(sorted_keys, key)
	}

	sort.Strings(sorted_keys)

	for _, v := range sorted_keys {
		params_data += v
		params_data += params[v]
	}

	params_data += private_key

	return fmt.Sprintf("%x", sha1.Sum([]byte(params_data)))
}

func Request(base_url string, params map[string]string) (string, error) {
	client := &http.Client{}
	b, err := json.Marshal(params)

	if err != nil {
		return "", err
	}

	req, err := http.NewRequest("POST", base_url, bytes.NewBuffer([]byte(b)))

	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)

	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}
