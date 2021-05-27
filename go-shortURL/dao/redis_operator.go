package dao

import (
	"encoding/json"
	"fmt"
	"go-shortURL/constants"
	"go-shortURL/models"
	"go-shortURL/utils"
	"time"

	"github.com/go-redis/redis"
	"github.com/mattheath/base62"
)

type UrlConverter interface {
	Shorten(string, int64) (string, error)
	Unshorten(string) (string, error)
	GetUrlDetail(string) (string, error)
}

type RedisClient struct {
	Client *redis.Client
}

func NewRedisClient(addr, password string, db int) *RedisClient {
	conn := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})

	if _, err := conn.Ping().Result(); err != nil {
		panic(err)
	}

	return &RedisClient{Client: conn}
}

func (rc *RedisClient) Shorten(urlStr string, expiration int64) (string, error) {
	var err error

	// hash and get url hash
	urlHash := utils.ToSha1(urlStr)
	urlHashVal, err := rc.Client.Get(fmt.Sprintf(constants.UrlHashKey, urlHash)).Result()
	if err != nil {
		// It's ok that kvpair doesn't exist.

		if err != redis.Nil {
			return "", err
		}
	} else {
		// It's ok that kvpair has expired.

		if urlHashVal != "{}" {
			return urlHashVal, nil
		}
	}

	detailJson, err := json.Marshal(
		models.UrlDetail{
			Url:          urlStr,
			CreateAt:     time.Now().String(),
			ExpireWithin: time.Duration(expiration),
		},
	)
	if err != nil {
		return "", err
	}

	// increase and encode url id
	err = rc.Client.Incr(constants.UrlIdKey).Err()
	if err != nil {
		return "", err
	}
	oid, err := rc.Client.Get(constants.UrlIdKey).Int64()
	if err != nil {
		return "", err
	}
	eid := base62.EncodeInt64(oid)

	// url hash refers to eid
	err = rc.Client.Set(fmt.Sprintf(constants.UrlHashKey, urlHash), eid, time.Minute*time.Duration(expiration)).Err()
	if err != nil {
		return "", err
	}

	// eid refers to url str
	err = rc.Client.Set(fmt.Sprintf(constants.ShortLinkUrlKey, eid), urlStr, time.Minute*time.Duration(expiration)).Err()
	if err != nil {
		return "", err
	}

	// eid refers to detail json
	err = rc.Client.Set(fmt.Sprintf(constants.ShortLinkDetailKey, eid), detailJson, time.Minute*time.Duration(expiration)).Err()
	if err != nil {
		return "", err
	}

	return eid, nil
}

func (rc *RedisClient) Unshorten(eid string) (string, error) {
	urlStr, err := rc.Client.Get(fmt.Sprintf(constants.ShortLinkUrlKey, eid)).Result()
	if err != nil {
		return "", err
	}
	return urlStr, nil
}

func (rc *RedisClient) GetUrlDetail(eid string) (string, error) {
	detailJson, err := rc.Client.Get(fmt.Sprintf(constants.ShortLinkDetailKey, eid)).Result()
	if err != nil {
		return "", err
	}
	return detailJson, nil
}
