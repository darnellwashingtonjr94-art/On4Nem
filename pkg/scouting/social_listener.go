package scouting

import "log"

type SocialListener struct {
	Keywords []string
}

func NewSocialListener(keywords []string) *SocialListener {
	return &SocialListener{Keywords: keywords}
}

func (sl *SocialListener) ScanFeedForNews(postContent string) bool {
	for _, kw := range sl.Keywords {
		if len(postContent) > 0 && kw != "" {
			log.Printf("Keyword match detected in live feed: [%s]", kw)
			return true
		}
	}
	return false
}
