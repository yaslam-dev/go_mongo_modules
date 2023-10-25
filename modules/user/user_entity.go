package user

import (
	"strings"

	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson"
)

type User struct {
	mgm.DefaultModel `bson:",inline"`
	Name             string `json:"name" bson:"name"`
	Email            string `json:"email" bson:"email"`
	Bio              string `json:"bio" bson:"bio"`
	Image            string `json:"image" bson:"image"`
	LoggedBy         string `json:"loggedBy" bson:"loggedBy"`
}

func (model *User) CollectionName() string {
	return COLLECTION_NAME
}

func (u *User) GetByID() error {
	err := mgm.Coll(&User{}).FindByID(u.ID, &User{})
	if err != nil {
		return err
	}
	return nil
}

func (u *User) GetByEmail() error {
	err := mgm.Coll(&User{}).SimpleFind(&User{}, bson.D{{Key: "email", Value: u.Email}})
	if err != nil {
		return err
	}
	return nil
}

func (u *User) Create() error {
	err := mgm.Coll(&User{}).Create(u)
	if err != nil {
		return err
	}
	return nil
}

func (u *User) Update() error {
	err := mgm.Coll(&User{}).Update(u)
	if err != nil {
		return err
	}
	return nil
}

func (u *User) Prepare() {
	u.Email = strings.ToLower(u.Email)
}

func (u *User) Delete() error {
	err := mgm.Coll(&User{}).Delete(u)
	if err != nil {
		return err
	}
	return nil
}

func (u *User) SearchByName() ([]User, error) {
	var users []User
	err := mgm.Coll(&User{}).SimpleFind(&users, bson.D{{Key: "name", Value: u.Name}})
	if err != nil {
		return nil, err
	}
	return users, nil
}
