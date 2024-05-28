package curl

import (
	"fmt"
	"net/url"
	"testing"
)

func Test_MakeParams(t *testing.T) {
	params := url.Values{}
	params.Add("id", "1")
	params.Add("name", "admin")
	params.Add("pid", "0")

	body := MakeParams(params)

	fmt.Println(body)

}
