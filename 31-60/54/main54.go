package main

import (
	"fmt"
	"time"
)

func TimeAgo(pastTime time.Time) string {
	now := time.Now()
	diff := now.Sub(pastTime)

	switch {
	case diff.Hours() < 1:
		minutes := int(diff.Minutes())
		if minutes == 1 {
			return "1 minute ago"
		}
		return fmt.Sprintf("%d minutes ago", minutes)
	case diff.Hours() < 24:
		hours := int(diff.Hours())
		if hours == 1 {
			return "1 hour ago"
		}
		return fmt.Sprintf("%d hours ago", hours)
	case diff.Hours() < 48:
		return "1 day ago"
	case diff.Hours() < 24*7:
		days := int(diff.Hours() / 24)
		if days == 1 {
			return "1 day ago"
		}
		return fmt.Sprintf("%d days ago", days)
	case diff.Hours() < 24*30:
		months := int(diff.Hours() / (24 * 7 * 4))
		if months == 1 {
			return "1 month ago"
		}
		return fmt.Sprintf("%d months ago", months)
	default:
		years := int(diff.Hours() / (24 * 30 * 12))
		if years == 1 {
			return "1 year ago"
		}
		return fmt.Sprintf("%d years ago", years)
	}
}

func main() {

}
