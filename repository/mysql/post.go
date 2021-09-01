package mysql

import "github.com/Thewalkers2012/BlogBackend/models"

func CreatePost(req *models.Post) (err error) {
	query := `insert into post(post_id, title, content, author_id, community_id)
						values(?, ?, ?, ?, ?)`
	_, err = db.Exec(query, req.ID, req.Title, req.Content, req.AuthorID, req.CommunityID)
	return err
}
