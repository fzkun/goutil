package magicapi

import (
	"fmt"
	"testing"
)

func TestMagicApi(t *testing.T) {
	sdk := NewMagicSdk(NewContext(Config{
		MagicUrl: "http://10.8.0.165:9990",
		IsDebug:  true,
	}, nil, nil))
	var (
		list []map[string]any
	)
	if err := sdk.NewMagicRequest("/xt_oracle/random_nh_card", map[string]any{
		"num": 1,
	}, &list); err != nil {
		t.Fatal(err)
	}
	fmt.Println(list)
}
