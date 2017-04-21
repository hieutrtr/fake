package fake

import (
	"testing"
)

func TestAds(t *testing.T) {
	for _, lang := range GetLangs() {
		SetLang(lang)

		v := AdState()
		if v == "" {
			t.Errorf("Ad's state failed with lang %s", lang)
		}
	}
}
