package services

import (
	"context"
	"errors"

	"github.com/Avinashanakal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserServiceImpl struct {
	//pointer to the object
	UserCollection *mongo.Collection
	Ctx            context.Context
}

func NewUser(userCollection *mongo.Collection, ctx context.Context) *UserServiceImpl {
	return &UserServiceImpl{
		UserCollection: userCollection,
		Ctx:            ctx,
	}
}

func (u *UserServiceImpl) CreateUser(user *models.User) error {

	_, err := u.UserCollection.InsertOne(u.Ctx, user)
	if err != nil {
		return err
	}
	return nil
}

func (u *UserServiceImpl) GetUser(name *string) (*models.User, error) {
	var user *models.User

	query := bson.D{bson.E{Key: "name", Value: name}}
	if err := u.UserCollection.FindOne(u.Ctx, query).Decode(&user); err != nil {
		return user, err
	}
	return nil, nil
}

func (u *UserServiceImpl) GetAll() ([]*models.User, error) {

	var users []*models.User
	cursor, err := u.UserCollection.Find(u.Ctx, bson.D{{}})
	if err != nil {
		return nil, err
	}

	for cursor.Next(u.Ctx) {
		var user *models.User
		if err := cursor.Decode(&user); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}
	cursor.Close(u.Ctx)

	if len(users) == 0 {
		return nil, errors.New("documents not FOUND")
	}
	return users, nil
}

func (u *UserServiceImpl) UpdateUser(user *models.User) error {

	query := bson.D{bson.E{Key: "name", Value: user.Name}}
	update := bson.D{bson.E{Key: "$set", Value: bson.D{bson.E{Key: "name", Value: user.Name}, bson.E{Key: "age", Value: user.Age}, bson.E{Key: "address", Value: user.Address}}}}

	result, _ := u.UserCollection.UpdateOne(u.Ctx, query, update)
	if result.MatchedCount != 1 {
		return errors.New("no records found")
	}
	return nil
}

func (u *UserServiceImpl) DeleteUser(name *string) error {

	query := bson.D{bson.E{Key: "name", Value: name}}
	delete, _ := u.UserCollection.DeleteOne(u.Ctx, query)

	if delete.DeletedCount != 1 {
		return errors.New("no records deleted")
	}
	return nil
}
