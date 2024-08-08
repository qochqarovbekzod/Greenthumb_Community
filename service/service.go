package service

import (
	pb "community-service/generated/community"
	"community-service/storage/postgres"
	"context"
)

type CommunityServer struct {
	pb.UnimplementedComunityServiceServer
	Community *postgres.CommunityRepo
}

func (c *CommunityServer) CreateCommunity(ctx context.Context, in *pb.CreateCommunityRequest) (*pb.CreateCommunityResponse, error) {
	resp, err := c.Community.CreateCommunity(in)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *CommunityServer) GetCommunity(ctx context.Context, in *pb.GetCommunityRequest) (*pb.GetCommunityResponse, error) {
	resp, err := c.Community.GetCommunity(in)
	if err != nil {
		return nil, err
	}

	return resp,nil
}

func (c *CommunityServer) UpdateCommunity(ctx context.Context, in *pb.UpdateCommunityRequest) (*pb.UpdateCommunityResponse, error) {
	resp, err := c.Community.UpdateCommunity(in)
	if err != nil {
		return nil, err
	}

	return resp,nil
}

func (c *CommunityServer) DeleteCommunity(ctx context.Context, in *pb.DeleteCommunityRequest) (*pb.DeleteCommunityResponse, error) {
	resp, err := c.Community.DeleteCommunity(in)
	if err != nil {
		return nil, err
	}

	return resp,nil
}

func (c *CommunityServer) ListCommunities(ctx context.Context, in *pb.ListCommunitiesRequest) (*pb.ListCommunitiesResponse, error) {
	resp, err := c.Community.ListCommunities(in)
	if err != nil {
		return nil, err
	}

	return resp,nil
}

func (c *CommunityServer) JoinCommunity(ctx context.Context, in *pb.JoinCommunityRequest) (*pb.JoinCommunityResponse, error) {
	resp, err := c.Community.JoinCommunity(in)
	if err != nil {
		return nil, err
	}

	return resp,nil
}

func (c *CommunityServer) LeaveCommunity(ctx context.Context, in *pb.LeaveCommunityRequest) (*pb.LeaveCommunityResponse, error) {
	resp, err := c.Community.LeaveCommunity(in)
	if err != nil {
		return nil, err
	}

	return resp,nil
}

func (c *CommunityServer) CreateCommunityEvent(ctx context.Context, in *pb.CreateCommunityEventRequest) (*pb.CreateCommunityEventResponse, error) {
	resp, err := c.Community.CreateCommunityEvent(in)
	if err != nil {
		return nil, err
	}

	return resp,nil
}


func (c *CommunityServer) ListCommunityEvents(ctx context.Context, in *pb.ListCommunityEventsRequest) (*pb.ListCommunityEventsResponse, error) {
	resp, err := c.Community.ListCommunityEvents(in)
	if err != nil {
		return nil, err
	}

	return resp,nil
}

func (c *CommunityServer) CreateCommunityForumPost(ctx context.Context, in *pb.CreateCommunityForumPostRequest) (*pb.CreateCommunityForumPostRespnse, error) {
	resp, err := c.Community.CreateCommunityForumPost(in)
	if err != nil {
		return nil, err
	}

	return resp,nil
}

func (c *CommunityServer) ListCommunityForumPosts(ctx context.Context, in *pb.ListCommunityForumPostsRequest) (*pb.ListCommunityForumPostsResponse, error) {
	resp, err := c.Community.ListCommunityForumPosts(in)
	if err != nil {
		return nil, err
	}

	return resp,nil
}

func (c *CommunityServer) AddForumPostComment(ctx context.Context, in *pb.AddForumPostCommentRequest) (*pb.AddForumPostCommentResponse, error) {
	resp, err := c.Community.AddForumPostComment(in)
	if err != nil {
		return nil, err
	}

	return resp,nil
}

func (c *CommunityServer) ListForumPostComments(ctx context.Context, in *pb.ListForumPostCommentsRequest) (*pb.ListForumPostCommentsResponse, error) {
	resp, err := c.Community.ListForumPostComments(in)
	if err != nil {
		return nil, err
	}

	return resp,nil
}



