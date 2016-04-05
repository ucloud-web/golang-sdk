package main

import "fmt"
import "os"
import "strings"
import "./ucloud"

func main() {

	public_key := "xxxxxxxxxxxxxxxxxxx"
	private_key := "xxxxxxxxxxxxxxxxxx"
	base_url := "https://api.ucloud.cn"
	project_id := "xxxxx"

	args_len := len(os.Args)
	if args_len >= 3 {

		params := map[string]string{}
		params["PublicKey"] = public_key
		params["Action"] = "SendSms"
		params["Content"] = os.Args[2]

		phones := strings.Split(os.Args[1], "|")

		for key, val := range phones {
			params["Phone."+string(key)] = val
		}
		if project_id != "" {
			params["ProjectId"] = project_id
		}

		params["Signature"] = ucloud.VerfyAc(params, private_key)

		data, err := ucloud.Request(base_url, params)
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println(data)

	} else {
		fmt.Println("./send_sms '13764073xxx|13764073xxx' 'test' ")
	}
}
