package main

import "fmt"
import "os"
import "strings"
import "./ucloud/config"
import "./ucloud/sdk"

func main() {

	cnf, err := config.NewConfig("ini", "./config.conf")
	if err != nil {
		fmt.Println("read config error", err)
		return
	}

	fmt.Println("test")

	public_key := cnf.String("public_key")
	private_key := cnf.String("private_key")
	base_url := cnf.String("base_url")
	project_id := cnf.String("project_id")

	fmt.Println(public_key, private_key, base_url, project_id)

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

		params["Signature"] = sdk.VerfyAc(params, private_key)

		data, err := sdk.Request(base_url, params)
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println(data)

	} else {
		fmt.Println("./send_sms '13764073xxx|13764073xxx' 'test' ")
	}
}
