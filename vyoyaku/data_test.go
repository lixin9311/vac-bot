package vyoyaku

import (
	"fmt"
	"testing"
)

func TestData(t *testing.T) {
	infos, err := institutionsData.ReadDir(".")
	if err != nil {
		t.Error(err)
	}
	for _, v := range infos {
		fmt.Println(v.Name())
	}
}
