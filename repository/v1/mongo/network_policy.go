package mongo

import (
	"context"
	v1 "github.com/klovercloud-ci-cd/light-house-query/core/v1"
	"github.com/klovercloud-ci-cd/light-house-query/core/v1/repository"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

var (
	NetworkPolicyCollection = "networkPolicyCollection"
)

type networkPolicyRepository struct {
	manager *dmManager
	timeout time.Duration
}

func (n networkPolicyRepository) Get(agent string, option v1.ResourceQueryOption) ([]v1.NetworkPolicy, int64) {
	var results []v1.NetworkPolicy
	query := bson.M{
		"$and": []bson.M{{"agent_name": agent}},
	}
	coll := n.manager.Db.Collection(NetworkPolicyCollection)
	skip := option.Pagination.Page * option.Pagination.Limit
	findOptions := options.FindOptions{
		Limit: &option.Pagination.Limit,
		Skip:  &skip,
		Sort:  bson.M{"created_at": -1},
	}
	if option.AscendingSort {
		findOptions.Sort = bson.M{"created_at": 1}
	}
	result, err := coll.Find(n.manager.Ctx, query, &findOptions)
	if err != nil {
		log.Println(err.Error())
	}
	for result.Next(context.TODO()) {
		elemValue := new(v1.NetworkPolicy)
		err := result.Decode(elemValue)
		if err != nil {
			log.Println("[ERROR]", err)
			break
		}
		results = append(results, *elemValue)
	}
	count, err := coll.CountDocuments(n.manager.Ctx, query)
	if err != nil {
		log.Println(err.Error())
	}
	return results, count
}

// NewNetworkPolicyRepository returns repository.NetworkPolicy type repository
func NewNetworkPolicyRepository(timeout int) repository.NetworkPolicy {
	return &networkPolicyRepository{
		manager: GetDmManager(),
		timeout: time.Duration(timeout),
	}
}
