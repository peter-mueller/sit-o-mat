package user

import (
	"context"
	"testing"

	"gocloud.dev/docstore"

	_ "gocloud.dev/docstore/memdocstore"
)

func userCollection() *docstore.Collection {
	ctx := context.Background()
	coll, err := docstore.OpenCollection(ctx, "mem://user/name")
	if err != nil {
		panic(err)
	}
	return coll
}

func TestRegisterAndLogin(t *testing.T) {
	coll := userCollection()
	ctx := context.Background()
	defer coll.Close()
	s := Service{coll}

	u, err := s.RegisterUser(ctx, "p.mueller")
	if err != nil {
		t.Errorf("cannot register user: %v", err)
	}

	if u.Name != "p.mueller" {
		t.Error("name does not match")
	}
	if u.Password == "" {
		t.Errorf("no password was generated")
	}

	ok, err := s.LoginUser(ctx, u)
	if err != nil {
		t.Errorf("failed to login user: %v", err)
	}
	if !ok {
		t.Error("failed to login user with correct password")
	}

	u.Password = "wrongpassword"
	ok, err = s.LoginUser(ctx, u)
	if ok {
		t.Error("failed to reject user with wrong password")
	}
}

func TestRegisterAndDelete(t *testing.T) {

}
