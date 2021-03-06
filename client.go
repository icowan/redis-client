/**
 * @Time: 2020/3/30 10:38
 * @Author: solacowa@gmail.com
 * @File: client
 * @Software: GoLand
 */

package redisclient

import (
	"context"
	"strings"
	"time"

	"github.com/go-redis/redis"
)

type RedisClient interface {
	Set(ctx context.Context, k string, v interface{}, expir ...time.Duration) (err error)
	Get(ctx context.Context, k string) (v string, err error)
	Del(ctx context.Context, k string) (err error)
	Exists(ctx context.Context, keys ...string) int64
	HSet(ctx context.Context, k string, field string, v interface{}) (err error)
	HGet(ctx context.Context, k string, field string) (res string, err error)
	HGetAll(ctx context.Context, k string) (res map[string]string, err error)
	HLen(ctx context.Context, k string) (res int64, err error)
	ZCard(ctx context.Context, k string) (res int64, err error)
	ZRangeWithScores(ctx context.Context, k string, start, stop int64) (res []redis.Z, err error)
	ZAdd(ctx context.Context, k string, score float64, member interface{}) (err error)
	HDelAll(ctx context.Context, k string) (err error)
	HDel(ctx context.Context, k string, field string) (err error)
	Keys(ctx context.Context, pattern string) (res []string, err error)
	LLen(ctx context.Context, key string) int64
	RPop(ctx context.Context, key string) (res string, err error)
	LPush(ctx context.Context, key string, val interface{}) (err error)
	TypeOf(ctx context.Context, key string) (res string, err error)
	Close(ctx context.Context) error
	Subscribe(ctx context.Context, channels ...string) *redis.PubSub
	Publish(ctx context.Context, channel string, message interface{}) error
	Incr(ctx context.Context, key string, exp time.Duration) error
	SetPrefix(ctx context.Context, prefix string) RedisClient
	TTL(ctx context.Context, key string) time.Duration
	Ping(ctx context.Context) error
	Pipeline(ctx context.Context) redis.Pipeliner
	Pipelined(ctx context.Context, fn func(pipeliner redis.Pipeliner) error) ([]redis.Cmder, error)

	TxPipelined(ctx context.Context, fn func(pipeliner redis.Pipeliner) error) ([]redis.Cmder, error)
	TxPipeline(ctx context.Context) redis.Pipeliner

	Command(ctx context.Context) *redis.CommandsInfoCmd
	ClientGetName(ctx context.Context) string
	Echo(ctx context.Context, message interface{}) string
	Quit(ctx context.Context) error
	Unlink(ctx context.Context, keys ...string) int64
	Dump(ctx context.Context, key string) string
	Expire(ctx context.Context, key string, expiration time.Duration) bool
	ExpireAt(ctx context.Context, key string, tm time.Time) bool
	Migrate(ctx context.Context, host, port, key string, db int64, timeout time.Duration) error
	Move(ctx context.Context, key string, db int64) bool
	ObjectRefCount(ctx context.Context, key string) int64
	ObjectEncoding(ctx context.Context, key string) string
	ObjectIdleTime(ctx context.Context, key string) time.Duration
	Persist(ctx context.Context, key string) bool
	PExpire(ctx context.Context, key string, expiration time.Duration) bool
	PExpireAt(ctx context.Context, key string, tm time.Time) bool
	PTTL(ctx context.Context, key string) time.Duration
	RandomKey(ctx context.Context) string
	Rename(ctx context.Context, key, newkey string) *redis.StatusCmd
	RenameNX(ctx context.Context, key, newkey string) bool
	Restore(ctx context.Context, key string, ttl time.Duration, value string) error
	RestoreReplace(ctx context.Context, key string, ttl time.Duration, value string) error
	Sort(ctx context.Context, key string, sort *redis.Sort) []string
	SortStore(ctx context.Context, key, store string, sort *redis.Sort) int64
	SortInterfaces(ctx context.Context, key string, sort *redis.Sort) []interface{}
	Touch(ctx context.Context, keys ...string) int64
	Type(ctx context.Context, key string) string
	Scan(ctx context.Context, cursor uint64, match string, count int64) ([]string, uint64)
	SScan(ctx context.Context, key string, cursor uint64, match string, count int64) ([]string, uint64)
	HScan(ctx context.Context, key string, cursor uint64, match string, count int64) ([]string, uint64)
	ZScan(ctx context.Context, key string, cursor uint64, match string, count int64) ([]string, uint64)
	Append(ctx context.Context, key, value string) int64
	BitCount(ctx context.Context, key string, bitCount *redis.BitCount) int64
	BitOpAnd(ctx context.Context, destKey string, keys ...string) int64
	BitOpOr(ctx context.Context, destKey string, keys ...string) int64
	BitOpXor(ctx context.Context, destKey string, keys ...string) int64
	BitOpNot(ctx context.Context, destKey string, key string) int64
	BitPos(ctx context.Context, key string, bit int64, pos ...int64) int64
	Decr(ctx context.Context, key string) int64
	DecrBy(ctx context.Context, key string, decrement int64) int64
	GetBit(ctx context.Context, key string, offset int64) int64
	GetRange(ctx context.Context, key string, start, end int64) string
	GetSet(ctx context.Context, key string, value interface{}) string
	IncrBy(ctx context.Context, key string, value int64) int64
	IncrByFloat(ctx context.Context, key string, value float64) float64
	MGet(ctx context.Context, keys ...string) []interface{}
	MSet(ctx context.Context, pairs ...interface{}) error
	MSetNX(ctx context.Context, pairs ...interface{}) bool
	SetBit(ctx context.Context, key string, offset int64, value int) int64
	SetNX(ctx context.Context, key string, value interface{}, expiration time.Duration) bool
	SetXX(ctx context.Context, key string, value interface{}, expiration time.Duration) bool
	SetRange(ctx context.Context, key string, offset int64, value string) int64
	StrLen(ctx context.Context, key string) int64
	HExists(ctx context.Context, key, field string) bool
	HIncrBy(ctx context.Context, key, field string, incr int64) int64
	HIncrByFloat(ctx context.Context, key, field string, incr float64) float64
	HKeys(ctx context.Context, key string) []string
	HMGet(ctx context.Context, key string, fields ...string) []interface{}
	HMSet(ctx context.Context, key string, fields map[string]interface{}) error
	HSetNX(ctx context.Context, key, field string, value interface{}) bool
	HVals(ctx context.Context, key string) []string
	BLPop(ctx context.Context, timeout time.Duration, keys ...string) []string
	BRPop(ctx context.Context, timeout time.Duration, keys ...string) []string
	BRPopLPush(ctx context.Context, source, destination string, timeout time.Duration) string
	LIndex(ctx context.Context, key string, index int64) string
	LInsert(ctx context.Context, key, op string, pivot, value interface{}) int64
	LInsertBefore(ctx context.Context, key string, pivot, value interface{}) int64
	LInsertAfter(ctx context.Context, key string, pivot, value interface{}) int64
	LPop(ctx context.Context, key string) string
	LPushX(ctx context.Context, key string, value interface{}) int64
	LRange(ctx context.Context, key string, start, stop int64) []string
	LRem(ctx context.Context, key string, count int64, value interface{}) int64
	LSet(ctx context.Context, key string, index int64, value interface{}) error
	LTrim(ctx context.Context, key string, start, stop int64) error
	RPopLPush(ctx context.Context, source, destination string) string
	RPush(ctx context.Context, key string, values ...interface{}) int64
	RPushX(ctx context.Context, key string, value interface{}) int64
	SAdd(ctx context.Context, key string, members ...interface{}) int64
	SCard(ctx context.Context, key string) int64
	SDiff(ctx context.Context, keys ...string) []string
	SDiffStore(ctx context.Context, destination string, keys ...string) int64
	SInter(ctx context.Context, keys ...string) []string
	SInterStore(ctx context.Context, destination string, keys ...string) int64
	SIsMember(ctx context.Context, key string, member interface{}) bool
	SMembers(ctx context.Context, key string) []string
	SMembersMap(ctx context.Context, key string) map[string]struct{}
	SMove(ctx context.Context, source, destination string, member interface{}) bool
	SPop(ctx context.Context, key string) string
	SPopN(ctx context.Context, key string, count int64) []string
	SRandMember(ctx context.Context, key string) string
	SRandMemberN(ctx context.Context, key string, count int64) []string
	SRem(ctx context.Context, key string, members ...interface{}) int64
	SUnion(ctx context.Context, keys ...string) []string
	SUnionStore(ctx context.Context, destination string, keys ...string) int64
	XAdd(ctx context.Context, a *redis.XAddArgs) string
	XDel(ctx context.Context, stream string, ids ...string) int64
	XLen(ctx context.Context, stream string) int64
	XRange(ctx context.Context, stream, start, stop string) []redis.XMessage
	XRangeN(ctx context.Context, stream, start, stop string, count int64) []redis.XMessage
	XRevRange(ctx context.Context, stream string, start, stop string) []redis.XMessage
	XRevRangeN(ctx context.Context, stream string, start, stop string, count int64) []redis.XMessage
	XRead(ctx context.Context, a *redis.XReadArgs) []redis.XStream
	XReadStreams(ctx context.Context, streams ...string) []redis.XStream
	XGroupCreate(ctx context.Context, stream, group, start string) error
	XGroupCreateMkStream(ctx context.Context, stream, group, start string) error
	XGroupSetID(ctx context.Context, stream, group, start string) error
	XGroupDestroy(ctx context.Context, stream, group string) int64
	XGroupDelConsumer(ctx context.Context, stream, group, consumer string) int64
	XReadGroup(ctx context.Context, a *redis.XReadGroupArgs) []redis.XStream
	XAck(ctx context.Context, stream, group string, ids ...string) int64
	XPending(stream, group string) *redis.XPending
	XPendingExt(ctx context.Context, a *redis.XPendingExtArgs) []redis.XPendingExt
	XClaim(ctx context.Context, a *redis.XClaimArgs) []redis.XMessage
	XClaimJustID(ctx context.Context, a *redis.XClaimArgs) []string
	XTrim(ctx context.Context, key string, maxLen int64) int64
	XTrimApprox(ctx context.Context, key string, maxLen int64) int64
	BZPopMax(ctx context.Context, timeout time.Duration, keys ...string) redis.ZWithKey
	BZPopMin(timeout time.Duration, keys ...string) redis.ZWithKey
	ZAddNX(ctx context.Context, key string, members ...redis.Z) int64
	ZAddXX(ctx context.Context, key string, members ...redis.Z) int64
	ZAddCh(ctx context.Context, key string, members ...redis.Z) int64
	ZAddNXCh(ctx context.Context, key string, members ...redis.Z) int64
	ZAddXXCh(ctx context.Context, key string, members ...redis.Z) int64
	ZIncr(ctx context.Context, key string, member redis.Z) float64
	ZIncrNX(ctx context.Context, key string, member redis.Z) float64
	ZIncrXX(ctx context.Context, key string, member redis.Z) float64
	ZCount(ctx context.Context, key, min, max string) int64
	ZLexCount(ctx context.Context, key, min, max string) int64
	ZIncrBy(ctx context.Context, key string, increment float64, member string) float64
	ZInterStore(ctx context.Context, destination string, store redis.ZStore, keys ...string) int64
	ZPopMax(ctx context.Context, key string, count ...int64) []redis.Z
	ZPopMin(ctx context.Context, key string, count ...int64) []redis.Z
	ZRange(ctx context.Context, key string, start, stop int64) []string
	ZRangeByScore(ctx context.Context, key string, opt redis.ZRangeBy) []string
	ZRangeByLex(ctx context.Context, key string, opt redis.ZRangeBy) []string
	ZRangeByScoreWithScores(ctx context.Context, key string, opt redis.ZRangeBy) []redis.Z
	ZRank(ctx context.Context, key, member string) (int64, error)
	ZRem(ctx context.Context, key string, members ...interface{}) int64
	ZRemRangeByRank(ctx context.Context, key string, start, stop int64) (int64, error)
	ZRemRangeByScore(ctx context.Context, key, min, max string) int64
	ZRemRangeByLex(ctx context.Context, key, min, max string) int64
	ZRevRange(ctx context.Context, key string, start, stop int64) []string
	ZRevRangeWithScores(ctx context.Context, key string, start, stop int64) []redis.Z
	ZRevRangeByScore(ctx context.Context, key string, opt redis.ZRangeBy) []string
	ZRevRangeByLex(ctx context.Context, key string, opt redis.ZRangeBy) []string
	ZRevRangeByScoreWithScores(ctx context.Context, key string, opt redis.ZRangeBy) []redis.Z
	ZRevRank(ctx context.Context, key, member string) (int64, error)
	ZScore(ctx context.Context, key, member string) float64
	ZUnionStore(ctx context.Context, dest string, store redis.ZStore, keys ...string) int64
	PFAdd(ctx context.Context, key string, els ...interface{}) int64
	PFCount(ctx context.Context, keys ...string) int64
	PFMerge(ctx context.Context, dest string, keys ...string) error
	BgRewriteAOF(ctx context.Context) error
	BgSave(ctx context.Context) error
	ClientKill(ctx context.Context, ipPort string) error
	ClientKillByFilter(ctx context.Context, keys ...string) int64
	ClientList(ctx context.Context) string
	ClientPause(ctx context.Context, dur time.Duration) bool
	ClientID(ctx context.Context) int64
	ConfigGet(ctx context.Context, parameter string) []interface{}
	ConfigResetStat(ctx context.Context) error
	ConfigSet(ctx context.Context, parameter, value string) error
	ConfigRewrite(ctx context.Context) error
	DBSize(ctx context.Context) int64
	FlushAll(ctx context.Context) error
	FlushAllAsync(ctx context.Context) error
	FlushDB(ctx context.Context) error
	FlushDBAsync(ctx context.Context) error
	Info(ctx context.Context, section ...string) string
	LastSave(ctx context.Context) int64
	Save(ctx context.Context) error
	Shutdown(ctx context.Context) error
	ShutdownSave(ctx context.Context) error
	ShutdownNoSave(ctx context.Context) error
	SlaveOf(ctx context.Context, host, port string) error
	Time(ctx context.Context) time.Time
	Eval(ctx context.Context, script string, keys []string, args ...interface{}) *redis.Cmd
	EvalSha(ctx context.Context, sha1 string, keys []string, args ...interface{}) *redis.Cmd
	ScriptExists(ctx context.Context, hashes ...string) []bool
	ScriptFlush(ctx context.Context) error
	ScriptKill(ctx context.Context) error
	ScriptLoad(ctx context.Context, script string) string
	DebugObject(ctx context.Context, key string) string
	PubSubChannels(ctx context.Context, pattern string) []string
	PubSubNumSub(ctx context.Context, channels ...string) map[string]int64
	PubSubNumPat(ctx context.Context) int64
	ClusterSlots(ctx context.Context) []redis.ClusterSlot
	ClusterNodes(ctx context.Context) string
	ClusterMeet(ctx context.Context, host, port string) error
	ClusterForget(ctx context.Context, nodeID string) error
	ClusterReplicate(ctx context.Context, nodeID string) error
	ClusterResetSoft(ctx context.Context) error
	ClusterResetHard(ctx context.Context) error
	ClusterInfo(ctx context.Context) string
	ClusterKeySlot(ctx context.Context, key string) int64
	ClusterGetKeysInSlot(ctx context.Context, slot int, count int) []string
	ClusterCountFailureReports(ctx context.Context, nodeID string) int64
	ClusterCountKeysInSlot(ctx context.Context, slot int) int64
	ClusterDelSlots(ctx context.Context, slots ...int) error
	ClusterDelSlotsRange(ctx context.Context, min, max int) error
	ClusterSaveConfig(ctx context.Context) error
	ClusterSlaves(ctx context.Context, nodeID string) []string
	ClusterFailover(ctx context.Context) error
	ClusterAddSlots(ctx context.Context, slots ...int) error
	ClusterAddSlotsRange(ctx context.Context, min, max int) error
	GeoAdd(ctx context.Context, key string, geoLocation ...*redis.GeoLocation) int64
	GeoPos(ctx context.Context, key string, members ...string) []*redis.GeoPos
	GeoRadius(ctx context.Context, key string, longitude, latitude float64, query *redis.GeoRadiusQuery) []redis.GeoLocation
	GeoRadiusRO(ctx context.Context, key string, longitude, latitude float64, query *redis.GeoRadiusQuery) []redis.GeoLocation
	GeoRadiusByMember(ctx context.Context, key, member string, query *redis.GeoRadiusQuery) []redis.GeoLocation
	GeoRadiusByMemberRO(ctx context.Context, key, member string, query *redis.GeoRadiusQuery) []redis.GeoLocation
	GeoDist(ctx context.Context, key string, member1, member2, unit string) float64
	GeoHash(ctx context.Context, key string, members ...string) []string
	ReadOnly(ctx context.Context) error
	ReadWrite(ctx context.Context) error
	MemoryUsage(ctx context.Context, key string, samples ...int) int64
}

const (
	RedisCluster = "cluster"
	RedisSingle  = "single"
	expiration   = 600 * time.Second
)

func NewRedisClient(hosts, password, prefix string, db int) (rds RedisClient, err error) {
	h := strings.Split(hosts, ",")
	if len(h) > 1 {
		rds = NewRedisCluster(
			h,
			password,
			prefix,
		)
	} else {
		rds = NewRedisSingle(
			hosts,
			password,
			prefix,
			db,
		)
	}

	if err = rds.Ping(context.Background()); err != nil {
		return nil, err
	}
	return rds, nil
}
