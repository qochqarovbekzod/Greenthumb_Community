package postgres

import (
	pb "community-service/generated/community"
	"fmt"
	"reflect"
	"testing"
)

func TestCreateCommunity(t *testing.T) {
	db, err := ConnectDb()
	if err != nil {
		panic(err)
	}

	community := NewCommunityRepo(db)

	rescreate, err := community.CreateCommunity(&pb.CreateCommunityRequest{
		Name:        "Nur",
		Description: "Any",
		Location:    "Tashkent",
	})
	if err != nil {
		panic(err)
	}
	waitcreate := pb.CreateCommunityResponse{
		Success: true,
	}
	if !reflect.DeepEqual(rescreate, &waitcreate) {
		t.Errorf("have %v , wont %v", rescreate, &waitcreate)
	}
}

func TestGetCommunity(t *testing.T) {
	db, err := ConnectDb()
	if err != nil {
		panic(err)
	}

	community := NewCommunityRepo(db)
	resget, err := community.GetCommunity(&pb.GetCommunityRequest{
		Id: "30332b3e-433a-40b2-8066-8056bfcce188",
	})
	if err != nil {
		fmt.Println(err)
	}
	waitget := pb.GetCommunityResponse{
		Id:          "30332b3e-433a-40b2-8066-8056bfcce188",
		Name:        "Nur",
		Description: "Any",
		Location:    "Tashkent",
	}
	if !reflect.DeepEqual(resget, &waitget) {
		t.Errorf("have %v , wont %v", resget, &waitget)
	}
}

func TestUpdateCommunity(t *testing.T) {
	db, err := ConnectDb()
	if err != nil {
		panic(err)
	}

	community := NewCommunityRepo(db)
	resupdate, err := community.UpdateCommunity(&pb.UpdateCommunityRequest{
		Id:          "30332b3e-433a-40b2-8066-8056bfcce188",
		Name:        "NurMel",
		Description: "Any1",
		Location:    "Tashkent",
	})
	if err != nil {
		fmt.Println(err)
	}
	waitupdate := pb.UpdateCommunityResponse{
		Succses: true,
	}
	if !reflect.DeepEqual(resupdate, &waitupdate) {
		t.Errorf("have %v , wont %v", resupdate, &waitupdate)
	}
}

func TestDeleteCommunity(t *testing.T) {
	db, err := ConnectDb()
	if err != nil {
		panic(err)
	}
	community := NewCommunityRepo(db)
	resdelete, err := community.DeleteCommunity(&pb.DeleteCommunityRequest{
		Id: "30332b3e-433a-40b2-8066-8056bfcce188",
	})
	if err != nil {
		fmt.Println(err)
	}
	waitdelete := pb.DeleteCommunityResponse{
		Succses: true,
	}
	if !reflect.DeepEqual(resdelete, &waitdelete) {
		t.Errorf("have %v , wont %v", resdelete, &waitdelete)
	}
}

func TestListCommunity(t *testing.T) {
	db, err := ConnectDb()
	if err != nil {
		panic(err)
	}
	community := NewCommunityRepo(db)
	reslist, err := community.ListCommunities(&pb.ListCommunitiesRequest{
		Name: "Community_14",
		Limit: 1,
		Offset: 0,
	})
	if err != nil {
		fmt.Println(err)
	}
	waitlist := pb.ListCommunitiesResponse{
		Comunitys: []*pb.Comunity{
			{Name: "Community_14",
				Description: "This is the description for Community_14.",
				Location:    "Location_100",
			},
		},
	}
	if !reflect.DeepEqual(reslist, &waitlist) {
		t.Errorf("have %v , wont %v", reslist, &waitlist)
	}
}
