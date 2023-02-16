package entry_model

import (
	"context"
	"wordle-server/configs"
	"wordle-server/entry"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var EntryCollection = configs.DB.Collection("entries")

// type Entry struct {
// 	ID       primitive.ObjectID `bson:"_id,omitempty"`
// 	PuzzleId primitive.ObjectID `bson:"puzzleId"`
// 	UserId   primitive.ObjectID `bson:"userId"`
// 	Guesses  [5]struct {
// 		Guess  string          `bson:"guess"`
// 		Result [5]LetterResult `bson:"result"`
// 	} `bson:"guesses"`
// }

// type EntryInput struct {
// 	Guess    string `json:"guess" validate:"required,len=5,alpha"`
// 	PuzzleId string `json:"puzzleId" validate:"required,len=24,hexadecimal"`
// 	UserId   string `json:"userId" validate:"required,len=24,hexadecimal"`
// }

func GetEntryByID(ctx context.Context, ID primitive.ObjectID) (entry.Entry, error) {
	var entry entry.Entry
	filter := bson.D{
		{Key: "_id", Value: ID},
	}
	err := EntryCollection.FindOne(ctx, filter).Decode(&entry)
	return entry, err
}

func GetEntry(ctx context.Context, entryInput entry.EntryInput) (entry.Entry, error) {
	var entry entry.Entry
	puzzleObjId, _ := primitive.ObjectIDFromHex(entryInput.PuzzleId)
	userObjId, _ := primitive.ObjectIDFromHex(entryInput.PuzzleId)
	filter := bson.D{
		{"puzzleId", puzzleObjId},
		{"userId", userObjId},
	}
	err := EntryCollection.FindOne(ctx, filter).Decode(&entry)
	return entry, err
}

func CreateEntry(ctx context.Context, entryInput entry.EntryInput) (*mongo.InsertOneResult, error) {
	puzzleObjId, _ := primitive.ObjectIDFromHex(entryInput.PuzzleId)
	userObjId, _ := primitive.ObjectIDFromHex(entryInput.PuzzleId)
	entry := entry.Entry{
		PuzzleId: puzzleObjId,
		UserId:   userObjId,
		Guesses:  []string{},
		Results:  [][]entry.LetterResult{},
	}
	return EntryCollection.InsertOne(ctx, entry)
}

func UpdateEntry(ctx context.Context, ID primitive.ObjectID, update bson.D) entry.Entry {
	var entry entry.Entry
	filter := bson.D{{"_id", ID}}
	opts := options.FindOneAndUpdate().SetReturnDocument(options.After)
	EntryCollection.FindOneAndUpdate(ctx, filter, update, opts).Decode(&entry)
	return entry
}
