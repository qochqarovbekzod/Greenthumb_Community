package postgres

import (
	pb "community-service/generated/community"
	"community-service/pkg"
	"database/sql"
	"fmt"
)

type CommunityRepo struct {
	DB *sql.DB
}

func NewCommunityRepo(db *sql.DB) *CommunityRepo {
	return &CommunityRepo{DB: db}
}

func (c *CommunityRepo) CreateCommunity(in *pb.CreateCommunityRequest) (*pb.CreateCommunityResponse, error) {
	rows, err := c.DB.Exec(`
			INSERS INTO
			communities(
				name,
				description,
				location
				)
			VALUES(
				$1,
				$2,
				$3)
			`, in.Name, in.Description, in.Location)

	if err != nil {
		return &pb.CreateCommunityResponse{Success: false}, err
	}

	rowsAffected, err := rows.RowsAffected()

	if err != nil || rowsAffected == 0 {
		return &pb.CreateCommunityResponse{Success: false}, err
	}

	return &pb.CreateCommunityResponse{Success: true}, nil
}

func (c *CommunityRepo) GetCommunity(in *pb.GetCommunityRequest) (*pb.GetCommunityResponse, error) {

	var resp pb.GetCommunityResponse
	err := c.DB.QueryRow(`
			SELECT
				id,
				name,
				description,
				location
			FROM communities
			WHERE
				id=$1 AND deleted_at=0
			`, in.Id).Scan(&resp.Id, &resp.Name, &resp.Description, &resp.Location)

	return &resp, err
}

func (c *CommunityRepo) UpdateCommunity(in *pb.UpdateCommunityRequest) (*pb.UpdateCommunityResponse, error) {

	params := make(map[string]interface{})

	var query = "UPDATE communities SET "
	if in.Name != "" {
		query += "name = :name, "
		params["name"] = in.Name
	}
	if in.Description != "" {
		query += "Description = :Description, "
		params["Description"] = in.Description
	}
	if in.Location != "" {
		query += "Location = :Location, "
		params["Location"] = in.Location
	}

	query += "updated_at = CURRENT_TIMESTAMP WHERE id = :id AND deleted_at = 0"
	params["id"] = in.Id
	query, args := pkg.ReplaceQueryParams(query, params)

	res, err := c.DB.Exec(query, args...)
	if err != nil {
		return nil, err
	}

	rowsAffected, _ := res.RowsAffected()
	if rowsAffected == 0 {
		return &pb.UpdateCommunityResponse{Succses: false}, fmt.Errorf("no rows affected, user with id %s not found", in.Id)
	}

	return &pb.UpdateCommunityResponse{Succses: true}, nil

}

func (c *CommunityRepo) DeleteCommunity(in *pb.DeleteCommunityRequest) (*pb.DeleteCommunityResponse, error) {
	rows, err := c.DB.Exec(`
			UPDATE
				communities	
			SET de
				deleted_ad=date_part('epoch', current_timestamp)::INT 
			WHERE
				id=$1
		`, in.Id)

	if err != nil {
		return &pb.DeleteCommunityResponse{Succses: false}, err
	}

	rowsAffected, err := rows.RowsAffected()

	if err != nil || rowsAffected == 0 {
		return &pb.DeleteCommunityResponse{Succses: false}, err
	}

	return &pb.DeleteCommunityResponse{Succses: true}, nil

}

func (c *CommunityRepo) ListCommunities(in *pb.ListCommunitiesRequest) (*pb.ListCommunitiesResponse, error) {

	var (
		params = make(map[string]interface{})
		arr    []interface{}
		filter string
	)

	query := `SELECT name, description,location
	FROM communities WHERE true `

	if len(in.Name) > 0 {
		params["name"] = in.Name
		filter += " and name = :name "
	}

	if in.Description != "" {
		params["Description"] = in.Description
		filter += " and Description = :Description "
	}

	if in.Location != "" {
		params["lacation"] = in.Location
		filter += " and location = :location "
	}

	if in.Offset > 0 {
		params["offset"] = in.Offset
		filter += " OFFSET :offset"
	}

	if in.Limit > 0 {
		params["limit"] = in.Limit
		filter += " LIMIT :limit"
	}
	query = query + filter

	query, arr = pkg.ReplaceQueryParams(query, params)
	fmt.Println(query, arr)
	rows, err := c.DB.Query(query, arr...)
	fmt.Println(err, query)
	if err != nil {
		return nil, err
	}

	var cpmmunityes []*pb.Comunity
	for rows.Next() {
		var community pb.Comunity
		err := rows.Scan(&community.Name, &community.Description, &community.Location)

		if err != nil {
			return nil, err
		}

		cpmmunityes = append(cpmmunityes, &community)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return &pb.ListCommunitiesResponse{Comunitys: cpmmunityes}, nil

}

func (c *CommunityRepo) JoinCommunity(in *pb.JoinCommunityRequest) (*pb.JoinCommunityResponse, error) {
	rows, err := c.DB.Exec(`
			INSETR INTO
			community_members(
				communityi_id,
				user_id)
			VALUES(
				$1,
				$2)`, in.CommunityId, in.UserId)
	if err != nil {
		return &pb.JoinCommunityResponse{Success: false}, err
	}
	rowsAffected, err := rows.RowsAffected()
	if err != nil || rowsAffected == 0 {
		return &pb.JoinCommunityResponse{Success: false}, err
	}

	return &pb.JoinCommunityResponse{Success: true}, nil
}

func (c *CommunityRepo) LeaveCommunity(in *pb.LeaveCommunityRequest) (*pb.LeaveCommunityResponse, error) {

	row, err := c.DB.Exec(`
			UPDATE 
			community_members
			SET
			deleted_at=date_part('epoch', current_timestamp)::INT 
			WHERE community_id=$1
			`, in.CommunityId)

	if err != nil {
		return &pb.LeaveCommunityResponse{Success: false}, err
	}

	rowsAffected, err := row.RowsAffected()

	if err != nil || rowsAffected == 0 {
		return &pb.LeaveCommunityResponse{Success: false}, err
	}

	return &pb.LeaveCommunityResponse{Success: true}, nil
}

func (c *CommunityRepo) CreateCommunityEvent(in *pb.CreateCommunityEventRequest) (*pb.CreateCommunityEventResponse, error) {

	rows, err := c.DB.Exec(`
			INSETR INTO
			events(
				id,
				community_id,
				name,
				description,
				type,
				start_type,
				end_type,
				location)
			VALUES(
				$1,
				$2,
				$3,
				$4,
				$5,
				$6,
				$7,
				$8)`, in.Id, in.ComunityId, in.Name, in.Description, in.Type, in.StartType, in.EndType, in.Location)
	if err != nil {
		return &pb.CreateCommunityEventResponse{Success: false}, err
	}
	rowsAffected, err := rows.RowsAffected()
	if err != nil || rowsAffected == 0 {
		return &pb.CreateCommunityEventResponse{Success: false}, err
	}

	return &pb.CreateCommunityEventResponse{Success: true}, nil

}

func (c *CommunityRepo) ListCommunityEvents(in *pb.ListCommunityEventsRequest) (*pb.ListCommunityEventsResponse, error) {
	rows, err := c.DB.Query(`
			SELECT 
				id,
				name,
				description,
				type,
				start_type,
				end_type,
				location
			FROM events
			WHERE
				community_id=$1
			`, in.CommunityId)
	if err != nil {
		return nil, err
	}

	var communityEvents []*pb.CommunityEvent

	for rows.Next() {
		var communityEvent pb.CommunityEvent
		err = rows.Scan(&communityEvent.Id, &communityEvent.Name, &communityEvent.Description, &communityEvent.Type,
			&communityEvent.StartType, &communityEvent.EndType, &communityEvent.Location)

		if err != nil {
			return nil, err
		}
		communityEvents = append(communityEvents, &communityEvent)

	}
	return &pb.ListCommunityEventsResponse{CommunityEvents: communityEvents}, nil

}
func (c *CommunityRepo) CreateCommunityForumPost(in *pb.CreateCommunityForumPostRequest) (*pb.CreateCommunityForumPostRespnse, error) {
	rows, err := c.DB.Exec(`
			INSETR INTO
			forum_posts(
				id,
				community_id,
				user_id,
				title,
				content)
			VALUES(
				$1,
				$2,
				$3,
				$4,
				$5
			)`, in.Id, in.CommunityId, in.UserId, in.Title, in.Content)
	if err != nil {
		return &pb.CreateCommunityForumPostRespnse{Success: false}, err
	}
	rowsAffected, err := rows.RowsAffected()
	if err != nil || rowsAffected == 0 {
		return &pb.CreateCommunityForumPostRespnse{Success: false}, err
	}

	return &pb.CreateCommunityForumPostRespnse{Success: true}, nil

}

func (c *CommunityRepo) ListCommunityForumPosts(in *pb.ListCommunityForumPostsRequest) (*pb.ListCommunityForumPostsResponse, error) {

	rows, err := c.DB.Query(`
			SELECT 
				id,
				user_id,
				title,
				content,
			FROM form_post
			WHERE
				community_id=$1
			`, in.ComunityId)
	if err != nil {
		return nil, err
	}

	var forumPosts []*pb.ForumPost

	for rows.Next() {
		var forumPost pb.ForumPost
		err = rows.Scan(&forumPost.Id, &forumPost.UserId, &forumPost.Title, &forumPost.Content)

		if err != nil {
			return nil, err
		}
		forumPosts = append(forumPosts, &forumPost)

	}

	return &pb.ListCommunityForumPostsResponse{ForumPosts: forumPosts}, nil

}

func (c *CommunityRepo) AddForumPostComment(in *pb.AddForumPostCommentRequest) (*pb.AddForumPostCommentResponse, error) {

	rows, err := c.DB.Exec(`
			INSETR INTO
			forum_comments(
				id,
				post_id,
				user_id,
				comment)
			VALUES(
				$1,
				$2,
				$3,
				$4
			)`, in.Id, in.PostId, in.UserId, in.Comment)
	if err != nil {
		return &pb.AddForumPostCommentResponse{Success: false}, err
	}
	rowsAffected, err := rows.RowsAffected()
	if err != nil || rowsAffected == 0 {
		return &pb.AddForumPostCommentResponse{Success: false}, err
	}

	return &pb.AddForumPostCommentResponse{Success: true}, nil

}

func (c *CommunityRepo) ListForumPostComments(in *pb.ListForumPostCommentsRequest) (*pb.ListForumPostCommentsResponse, error) {

	rows, err := c.DB.Query(`
			SELECT 
				id,
				user_id,
				comment
			FROM forum_comments
			WHERE
				post_id=$1
			`, in.PostId)
	if err != nil {
		return nil, err
	}

	var listForumPostComments []*pb.ListForumPostComment

	for rows.Next() {
		var listForumPostComment pb.ListForumPostComment
		err = rows.Scan(&listForumPostComment.Id, &listForumPostComment.UserId, &listForumPostComment.Comment)

		if err != nil {
			return nil, err
		}
		listForumPostComments = append(listForumPostComments, &listForumPostComment)

	}

	return &pb.ListForumPostCommentsResponse{ListForumPostComments: listForumPostComments}, nil

}
