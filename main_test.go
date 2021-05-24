package main

import (
	"os"
	"requester/internal"
	"testing"
)

func TestMainFunc(t *testing.T) {
	wc := &internal.Requester{
		NumOfRoutines: 10,
		Urls:          nil,
	}

	os.Args = append(os.Args, "-parallel 3 adjust.com google.com facebook.com yahoo.com yandex.com twitter.com reddit.com/r/funny reddit.com/r/notfunny baroquemusiclibrary.com")

	res, err := wc.Run()
	if err != nil {
		t.Error(err)
	}
	sampleResult := [9]string{
		"google.com f1f262d8b1ad2d1af781cfa6d0fd3ae1",
		"facebook.com cb37ba09000643c04da645bf4adc1b6d",
		"adjust.com a2844c6c59e5e037c455179f8be9f9c2",
		"yandex.com 91785d6e7206fc90bf666f5e9aa8e951",
		"twitter.com 4721ffea9b9be9b0484f25a9388c10ef",
		"yahoo.com cc284be38587dbf47bd38632f84bd94a",
		"reddit.com/r/notfunny 1Furfhe8BFbkEhXn1xcYPr8jYAACNpfV7p",
		"reddit.com/r/funny 7dd1af8550e3e18a41054e8dadd66102",
		"baroquemusiclibrary.com 70a2d332a7e78ec29e9dbab7ae878240",
	}
	if len(sampleResult) == len(res) {
		for _, line := range sampleResult {
			_, found := internal.Find(res, line)
			if !found {
				t.Error("something doesn't fit")
			}
		}

	}
	// Test results here, and decide pass/fail.
}
