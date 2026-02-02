package markdown

import "regexp"

// stripSlackEmojiSkinTones removes Slack-style emoji skin tone suffixes.
// Examples:
//   :thumbsup_tone2:
//   :thumbsup::skin-tone-3:
func stripSlackEmojiSkinTones(text string) string {
	re := regexp.MustCompile(`(_tone|::skin-tone-)[1-6]`)
	return re.ReplaceAllString(text, "")
}
