package post

import (
	"rest-in-go/models"
	"rest-in-go/utils"
)

type PostResponse struct {
	ID        uint     `json:"id"`
	Title     string   `json:"title"`
	Body      string   `json:"body"`
	Tags      []string `json:"tags"`
	UpVotes   int      `json:"upvotes"`
	DownVotes int      `json:"downvotes"`
}

func processPosts(posts []models.Post) []PostResponse {
	returnPost := utils.Map(posts, func(post models.Post) PostResponse {
		return PostResponse{
			ID:    post.ID,
			Title: post.Title,
			Body:  post.Body,
			Tags: utils.Map(post.Tags, func(tag models.Tag) string {
				return tag.Name
			}),
			UpVotes: utils.Reduce(post.Votes, func(acc int, vote models.Vote) int {
				if vote.Value == 1 {
					return acc + 1
				}
				return acc
			}, 0),
			DownVotes: utils.Reduce(post.Votes, func(acc int, vote models.Vote) int {
				if vote.Value == -1 {
					return acc + 1
				}
				return acc
			}, 0),
		}
	})

	return returnPost
}
