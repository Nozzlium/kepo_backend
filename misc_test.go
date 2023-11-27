package main

import (
	"fmt"
	"nozzlium/kepo_backend/helper"
	"testing"
)

func TestStringFormat(t *testing.T) {
	text1 := "Lorem ipsum dolor sit ut."
	fmt.Println(helper.GetNotificationPreview(text1))
	text2 := "Lorem ipsum dolor sit amet viverra."
	fmt.Println(len(text2))
	fmt.Println(helper.GetNotificationPreview(text2))
}
