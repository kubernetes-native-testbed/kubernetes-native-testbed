package comment

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type CommentRepository interface {
	FindByUUID(string) (*Comment, error)
	Store(*Comment) (string, error)
	Update(*Comment) error
	DeleteByUUID(string) error
}

type commentRepositoryMongo struct {
	client *mongo.Client
	dbName string
}

func (cr *commentRepositoryMongo) FindByUUID(uuid string) (*Comment, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	if err := updateConnection(cr.client, ctx); err != nil {
		return nil, err
	}
	coll := cr.client.Database(cr.dbName).Collection("comments")

	var c Comment
	filter := bson.M{"uuid": uuid}
	if err := coll.FindOne(ctx, filter, options.FindOne()).Decode(&c); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("record(%s) not found: %w", uuid, err)
		}
		return nil, err
	}

	if c.DeletedAt != nil {
		return nil, fmt.Errorf("%s is deleted", uuid)
	}

	return &c, nil
}

func (cr *commentRepositoryMongo) Store(c *Comment) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	if err := updateConnection(cr.client, ctx); err != nil {
		return "", err
	}
	coll := cr.client.Database(cr.dbName).Collection("comments")

	c.UUID = uuid.New().String()
	c.CreatedAt = time.Now()

	// exists check
	var ret Comment
	filter := bson.M{"uuid": c.UUID}
	if err := coll.FindOne(ctx, filter, options.FindOne()).Decode(&ret); err != nil {
		if err != mongo.ErrNoDocuments {
			return "", err
		}
	} else {
		return "", fmt.Errorf("record already exists: %s", c.UUID)
	}

	_, err := coll.InsertOne(ctx, c)
	if err != nil {
		return "", err
	}

	return c.UUID, nil
}

func (cr *commentRepositoryMongo) Update(c *Comment) error {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	if err := updateConnection(cr.client, ctx); err != nil {
		return err
	}
	coll := cr.client.Database(cr.dbName).Collection("comments")

	filter := bson.M{"uuid": c.UUID}
	update := bson.M{"$set": bson.M{
		"message":             c.Message,
		"user_uuid":           c.UserUUID,
		"parent_comment_uuid": c.ParentCommentUUID,
		"updatedat":           time.Now(),
	}}
	if _, err := coll.UpdateOne(ctx, filter, update); err != nil {
		return err
	}

	return nil
}

func (cr *commentRepositoryMongo) DeleteByUUID(uuid string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	if err := updateConnection(cr.client, ctx); err != nil {
		return err
	}
	coll := cr.client.Database(cr.dbName).Collection("comments")

	filter := bson.M{"uuid": uuid}
	update := bson.M{"$set": bson.M{
		"deletedat": time.Now(),
	}}
	_, err := coll.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}

	return nil
}

func updateConnection(client *mongo.Client, ctx context.Context) error {
	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		if err := client.Connect(ctx); err != nil {
			return err
		}
	}
	return nil
}

type CommentRepositoryMongoConfig struct {
	Host     string
	Port     int
	Username string
	Password string
	DBName   string
}

func (c *CommentRepositoryMongoConfig) Connect() (CommentRepository, error) {
	uri := fmt.Sprintf("mongodb://%s:%s@%s:%d", c.Username, c.Password, c.Host, c.Port)
	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}
	return &commentRepositoryMongo{client: client, dbName: c.DBName}, nil
}
