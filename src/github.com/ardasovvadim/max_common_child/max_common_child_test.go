package max_common_child

import (
	"testing"
)

func TestShouldGetMaxCommonChild(t *testing.T) {
	result, err := GetMaxCommonChild("harry", "robby")
	if result != "ry" || err != nil {
		t.Error(
			"For", "[harry, robby]",
			"expected", "ry",
			"got", result,
		)
	}
}

func TestShouldThrowErrorOnGettingMaxCommonChild(t *testing.T) {
	_, err := GetMaxCommonChild("xx", "yy")
	if err != nil {
		t.Error(
			"For", "[xx, yy] there is no error",
		)
	}
}
