package mysql

import (
	"database/sql"
	"errors"

	"github.com/Thewalkers2012/BlogBackend/models"
	"go.uber.org/zap"
)

var ErrorInValidID = errors.New("无效的id")

func GetCommunityList() (communityList []*models.Community, err error) {
	query := `select community_id, community_name from community`
	if err := db.Select(&communityList, query); err != nil {
		if err == sql.ErrNoRows {
			zap.L().Warn("there is no community in db")
			err = nil
		}
	}

	return
}

// GetCommunityDetailsByID 根据 id 查询社区详情
func GetCommunityDetailsByID(id int64) (community *models.CommunityDetails, err error) {
	community = new(models.CommunityDetails)
	query := `select community_id, community_name, introduction, create_time
		from community
		where community_id = ?
	`

	if err = db.Get(community, query, id); err != nil {
		if err == sql.ErrNoRows {
			err = ErrorInValidID
		}
	}

	return community, err
}
