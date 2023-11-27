package helper

import "fmt"

func GetNotificationPreview(text string) string {
	if len(text) > 25 {
		return fmt.Sprintf("%.22s...", text)
	}
	return text
}
