package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type UserInfo struct {
	Result []struct {
		Handle       string `json:"handle"`
		LastOnline   uint64 `json:"lastOnlineTimeSeconds"`
		Organization string `json:"organization"`
	} `json:"result"`
}

func checkerr(e error) {
	if e != nil {
		panic(e)
	}
}
func main() {

	res, err := http.Get("https://codeforces.com/api/user.info?handles=shadman_96&checkHistoricHandles=false")
	checkerr(err)

	defer res.Body.Close()

	if res.StatusCode != 200 {
		panic(" API is not working")
	}

	body, err := io.ReadAll(res.Body)
	checkerr(err)

	var holder UserInfo
	err = json.Unmarshal(body, &holder)
	checkerr(err)

	result := holder.Result[0]
	fmt.Println(result.Handle)

}
