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
	PVCCollection = "pvcCollection"
)

type persistentVolumeClaimRepository struct {
	manager *dmManager
	timeout time.Duration
}

func (p persistentVolumeClaimRepository) GetById(id, agent, processId string) v1.PersistentVolumeClaim {
	query := bson.M{
		"$and": []bson.M{
			{"agent_name": agent},
			{"obj.metadata.labels.process_id": processId},
			{"obj.metadata.uid": id},
		},
	}
	coll := p.manager.Db.Collection(PVCCollection)
	result := coll.FindOne(p.manager.Ctx, query, nil)
	elemValue := new(v1.PersistentVolumeClaim)
	err := result.Decode(elemValue)
	if err != nil {
		log.Println("[ERROR]", err)
	}
	return *elemValue
}

func (p persistentVolumeClaimRepository) GetByAgentAndProcessIdWithoutPagination(agent, processId string) []v1.PersistentVolumeClaim {
	var results []v1.PersistentVolumeClaim
	query := bson.M{
		"$and": []bson.M{
			{"agent_name": agent},
			{"obj.metadata.labels.process_id": processId},
		},
	}
	coll := p.manager.Db.Collection(PVCCollection)
	findOptions := options.FindOptions{
		Sort: bson.M{"created_at": -1},
	}
	result, err := coll.Find(p.manager.Ctx, query, &findOptions)
	if err != nil {
		log.Println(err.Error())
		return results
	}
	for result.Next(context.TODO()) {
		elemValue := new(v1.PersistentVolumeClaim)
		err := result.Decode(elemValue)
		if err != nil {
			log.Println("[ERROR]", err)
			break
		}
		results = append(results, *elemValue)
	}
	return results
}

func (p persistentVolumeClaimRepository) GetByAgentAndProcessId(agent, processId string, option v1.ResourceQueryOption) ([]v1.PersistentVolumeClaim, int64) {
	var results []v1.PersistentVolumeClaim
	query := bson.M{
		"$and": []bson.M{
			{"agent_name": agent},
			{"obj.metadata.labels.process_id": processId},
		},
	}
	coll := p.manager.Db.Collection(PVCCollection)
	skip := option.Pagination.Page * option.Pagination.Limit
	findOptions := options.FindOptions{
		Limit: &option.Pagination.Limit,
		Skip:  &skip,
		Sort:  bson.M{"created_at": -1},
	}
	if option.AscendingSort {
		findOptions.Sort = bson.M{"created_at": 1}
	}
	result, err := coll.Find(p.manager.Ctx, query, &findOptions)
	if err != nil {
		log.Println(err.Error())
		return results, 0
	}
	for result.Next(context.TODO()) {
		elemValue := new(v1.PersistentVolumeClaim)
		err := result.Decode(elemValue)
		if err != nil {
			log.Println("[ERROR]", err)
			break
		}
		results = append(results, *elemValue)
	}
	count, err := coll.CountDocuments(p.manager.Ctx, query)
	if err != nil {
		log.Println(err.Error())
	}
	return results, count
}

func (p persistentVolumeClaimRepository) GetByAgentAndProcessIdAndOwnerReference(agent, ownerReference, processId string, option v1.ResourceQueryOption) ([]v1.PersistentVolumeClaim, int64) {
	var results []v1.PersistentVolumeClaim
	query := bson.M{
		"$and": []bson.M{
			{"agent_name": agent},
			{"obj.metadata.ownerReferences.uid": ownerReference},
			{"obj.metadata.labels.process_id": processId},
		},
	}
	coll := p.manager.Db.Collection(PVCCollection)
	skip := option.Pagination.Page * option.Pagination.Limit
	findOptions := options.FindOptions{
		Limit: &option.Pagination.Limit,
		Skip:  &skip,
		Sort:  bson.M{"created_at": -1},
	}
	if option.AscendingSort {
		findOptions.Sort = bson.M{"created_at": 1}
	}
	result, err := coll.Find(p.manager.Ctx, query, &findOptions)
	if err != nil {
		log.Println(err.Error())
		return results, 0
	}
	for result.Next(context.TODO()) {
		elemValue := new(v1.PersistentVolumeClaim)
		err := result.Decode(elemValue)
		if err != nil {
			log.Println("[ERROR]", err)
			break
		}
		results = append(results, *elemValue)
	}
	count, err := coll.CountDocuments(p.manager.Ctx, query)
	if err != nil {
		log.Println(err.Error())
	}
	return results, count
}

// NewPersistentVolumeClaimRepository returns repository.PersistentVolumeClaim type repository
func NewPersistentVolumeClaimRepository(timeout int) repository.PersistentVolumeClaim {
	return &persistentVolumeClaimRepository{
		manager: GetDmManager(),
		timeout: time.Duration(timeout),
	}
}
