package commentReaction

import (
	"database/sql"
	"errors"
	"fmt"
	"forum/internal/models"
)

type CommentReactionService struct {
	repo models.CommentReactionRepo
}

func NewCommentReactionService(repo models.CommentReactionRepo) *CommentReactionService {
	return &CommentReactionService{repo}
}

func (s *CommentReactionService) CreateCommentsReactions(reaction *models.CommentReactionDTO) error {
	// fmt.Println(reaction.CommentID)
	// fmt.Println(reaction.ID)
	// fmt.Println(reaction.Status)
	// fmt.Println(reaction.UserID)
	// fmt.Println("================================================================")
	r, err := s.repo.GetReactionByUserIDAndCommentID(reaction.UserID, reaction.CommentID)

	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			return err
		}
	} else {
		s.repo.DeleteCommentsReactions(r.ID)
		if r.Status == reaction.Status {
			return nil
		}
	}
	fmt.Println("================================================================")

	return s.CreateCommentsReactions(reaction)
}
