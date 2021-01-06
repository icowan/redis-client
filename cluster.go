/**
 * @Time: 2020/3/30 10:39
 * @Author: solacowa@gmail.com
 * @File: cluster
 * @Software: GoLand
 */

package redisclient

import (
	"context"
	"encoding/json"
	"time"

	"github.com/go-redis/redis"
)

type cluster struct {
	client *redis.ClusterClient
	//prefix func(s string) string
	prefix string
}

func (c *cluster) Pipeline(ctx context.Context) redis.Pipeliner {
	return c.client.Pipeline()
}

func (c *cluster) Pipelined(ctx context.Context, fn func(pipeliner redis.Pipeliner) error) ([]redis.Cmder, error) {
	return c.client.Pipelined(fn)
}

func (c *cluster) TxPipelined(ctx context.Context, fn func(pipeliner redis.Pipeliner) error) ([]redis.Cmder, error) {
	return c.client.TxPipelined(fn)
}

func (c *cluster) TxPipeline(ctx context.Context) redis.Pipeliner {
	return c.client.TxPipeline()
}

func (c *cluster) Command(ctx context.Context) *redis.CommandsInfoCmd {
	return c.client.Command()
}

func (c *cluster) ClientGetName(ctx context.Context) string {
	return c.client.ClientGetName().Val()
}

func (c *cluster) Echo(ctx context.Context, message interface{}) string {
	return c.client.Echo(message).Val()
}

func (c *cluster) Quit(ctx context.Context) error {
	return c.client.Quit().Err()
}

func (c *cluster) Unlink(ctx context.Context, keys ...string) int64 {
	return c.client.Unlink(keys...).Val()
}

func (c *cluster) Dump(ctx context.Context, key string) string {
	return c.client.Dump(key).Val()
}

func (c *cluster) Expire(ctx context.Context, key string, expiration time.Duration) bool {
	return c.client.Expire(key, expiration).Val()
}

func (c *cluster) ExpireAt(ctx context.Context, key string, tm time.Time) bool {
	return c.client.ExpireAt(key, tm).Val()
}

func (c *cluster) Migrate(ctx context.Context, host, port, key string, db int64, timeout time.Duration) error {
	return c.client.Migrate(host, port, key, db, timeout).Err()
}

func (c *cluster) Move(ctx context.Context, key string, db int64) bool {
	return c.client.Move(key, db).Val()
}

func (c *cluster) ObjectRefCount(ctx context.Context, key string) int64 {
	return c.client.ObjectRefCount(key).Val()
}

func (c *cluster) ObjectEncoding(ctx context.Context, key string) string {
	return c.client.ObjectEncoding(key).Val()
}

func (c *cluster) ObjectIdleTime(ctx context.Context, key string) time.Duration {
	return c.client.ObjectIdleTime(key).Val()
}

func (c *cluster) Persist(ctx context.Context, key string) bool {
	return c.client.Persist(key).Val()
}

func (c *cluster) PExpire(ctx context.Context, key string, expiration time.Duration) bool {
	return c.client.PExpire(key, expiration).Val()
}

func (c *cluster) PExpireAt(ctx context.Context, key string, tm time.Time) bool {
	return c.client.PExpireAt(key, tm).Val()
}

func (c *cluster) PTTL(ctx context.Context, key string) time.Duration {
	return c.client.PTTL(key).Val()
}

func (c *cluster) RandomKey(ctx context.Context) string {
	return c.client.RandomKey().Val()
}

func (c *cluster) Rename(ctx context.Context, key, newkey string) *redis.StatusCmd {
	return c.client.Rename(key, newkey)
}

func (c *cluster) RenameNX(ctx context.Context, key, newkey string) bool {
	return c.client.RenameNX(key, newkey).Val()
}

func (c *cluster) Restore(ctx context.Context, key string, ttl time.Duration, value string) error {
	return c.client.Restore(key, ttl, value).Err()
}

func (c *cluster) RestoreReplace(ctx context.Context, key string, ttl time.Duration, value string) error {
	return c.client.RestoreReplace(key, ttl, value).Err()
}

func (c *cluster) Sort(ctx context.Context, key string, sort *redis.Sort) []string {
	return c.client.Sort(key, sort).Val()
}

func (c *cluster) SortStore(ctx context.Context, key, store string, sort *redis.Sort) int64 {
	return c.client.SortStore(key, store, sort).Val()
}

func (c *cluster) SortInterfaces(ctx context.Context, key string, sort *redis.Sort) []interface{} {
	return c.client.SortInterfaces(key, sort).Val()
}

func (c *cluster) Touch(ctx context.Context, keys ...string) int64 {
	return c.client.Touch(keys...).Val()
}

func (c *cluster) Type(ctx context.Context, key string) string {
	return c.client.Type(key).Val()
}

func (c *cluster) Scan(ctx context.Context, cursor uint64, match string, count int64) ([]string, uint64) {
	return c.client.Scan(cursor, match, count).Val()
}

func (c *cluster) SScan(ctx context.Context, key string, cursor uint64, match string, count int64) ([]string, uint64) {
	return c.client.SScan(key, cursor, match, count).Val()
}

func (c *cluster) HScan(ctx context.Context, key string, cursor uint64, match string, count int64) ([]string, uint64) {
	return c.client.HScan(key, cursor, match, count).Val()
}

func (c *cluster) ZScan(ctx context.Context, key string, cursor uint64, match string, count int64) ([]string, uint64) {
	return c.client.ZScan(key, cursor, match, count).Val()
}

func (c *cluster) Append(ctx context.Context, key, value string) int64 {
	return c.client.Append(key, value).Val()
}

func (c *cluster) BitCount(ctx context.Context, key string, bitCount *redis.BitCount) int64 {
	return c.client.BitCount(key, bitCount).Val()
}

func (c *cluster) BitOpAnd(ctx context.Context, destKey string, keys ...string) int64 {
	return c.client.BitOpAnd(destKey, keys...).Val()
}

func (c *cluster) BitOpOr(ctx context.Context, destKey string, keys ...string) int64 {
	return c.client.BitOpOr(destKey, keys...).Val()
}

func (c *cluster) BitOpXor(ctx context.Context, destKey string, keys ...string) int64 {
	return c.client.BitOpXor(destKey, keys...).Val()
}

func (c *cluster) BitOpNot(ctx context.Context, destKey string, key string) int64 {
	return c.client.BitOpNot(destKey, key).Val()
}

func (c *cluster) BitPos(ctx context.Context, key string, bit int64, pos ...int64) int64 {
	return c.client.BitPos(key, bit, pos...).Val()
}

func (c *cluster) Decr(ctx context.Context, key string) int64 {
	return c.client.Decr(key).Val()
}

func (c *cluster) DecrBy(ctx context.Context, key string, decrement int64) int64 {
	return c.client.DecrBy(key, decrement).Val()
}

func (c *cluster) GetBit(ctx context.Context, key string, offset int64) int64 {
	return c.client.GetBit(key, offset).Val()
}

func (c *cluster) GetRange(ctx context.Context, key string, start, end int64) string {
	return c.client.GetRange(key, start, end).Val()
}

func (c *cluster) GetSet(ctx context.Context, key string, value interface{}) string {
	return c.client.GetSet(key, value).Val()
}

func (c *cluster) IncrBy(ctx context.Context, key string, value int64) int64 {
	return c.client.IncrBy(key, value).Val()
}

func (c *cluster) IncrByFloat(ctx context.Context, key string, value float64) float64 {
	return c.client.IncrByFloat(key, value).Val()
}

func (c *cluster) MGet(ctx context.Context, keys ...string) []interface{} {
	return c.client.MGet(keys...).Val()
}

func (c *cluster) MSet(ctx context.Context, pairs ...interface{}) error {
	return c.client.MSet(pairs...).Err()
}

func (c *cluster) MSetNX(ctx context.Context, pairs ...interface{}) bool {
	return c.client.MSetNX(pairs...).Val()
}

func (c *cluster) SetBit(ctx context.Context, key string, offset int64, value int) int64 {
	return c.client.SetBit(key, offset, value).Val()
}

func (c *cluster) SetNX(ctx context.Context, key string, value interface{}, expiration time.Duration) bool {
	return c.client.SetNX(key, value, expiration).Val()
}

func (c *cluster) SetXX(ctx context.Context, key string, value interface{}, expiration time.Duration) bool {
	return c.client.SetXX(key, value, expiration).Val()
}

func (c *cluster) SetRange(ctx context.Context, key string, offset int64, value string) int64 {
	return c.client.SetRange(key, offset, value).Val()
}

func (c *cluster) StrLen(ctx context.Context, key string) int64 {
	return c.client.StrLen(key).Val()
}

func (c *cluster) HExists(ctx context.Context, key, field string) bool {
	return c.client.HExists(key, field).Val()
}

func (c *cluster) HIncrBy(ctx context.Context, key, field string, incr int64) int64 {
	return c.client.HIncrBy(key, field, incr).Val()
}

func (c *cluster) HIncrByFloat(ctx context.Context, key, field string, incr float64) float64 {
	return c.client.HIncrByFloat(key, field, incr).Val()
}

func (c *cluster) HKeys(ctx context.Context, key string) []string {
	return c.client.HKeys(key).Val()
}

func (c *cluster) HMGet(ctx context.Context, key string, fields ...string) []interface{} {
	return c.client.HMGet(key, fields...).Val()
}

func (c *cluster) HMSet(ctx context.Context, key string, fields map[string]interface{}) error {
	return c.client.HMSet(key, fields).Err()
}

func (c *cluster) HSetNX(ctx context.Context, key, field string, value interface{}) bool {
	return c.client.HSetNX(key, field, value).Val()
}

func (c *cluster) HVals(ctx context.Context, key string) []string {
	return c.client.HVals(key).Val()
}

func (c *cluster) BLPop(ctx context.Context, timeout time.Duration, keys ...string) []string {
	return c.client.BLPop(timeout, keys...).Val()
}

func (c *cluster) BRPop(ctx context.Context, timeout time.Duration, keys ...string) []string {
	return c.client.BRPop(timeout, keys...).Val()
}

func (c *cluster) BRPopLPush(ctx context.Context, source, destination string, timeout time.Duration) string {
	return c.client.BRPopLPush(source, destination, timeout).Val()
}

func (c *cluster) LIndex(ctx context.Context, key string, index int64) string {
	return c.client.LIndex(key, index).Val()
}

func (c *cluster) LInsert(ctx context.Context, key, op string, pivot, value interface{}) int64 {
	return c.client.LInsert(key, op, pivot, value).Val()
}

func (c *cluster) LInsertBefore(ctx context.Context, key string, pivot, value interface{}) int64 {
	return c.client.LInsertBefore(key, pivot, value).Val()
}

func (c *cluster) LInsertAfter(ctx context.Context, key string, pivot, value interface{}) int64 {
	return c.client.LInsertAfter(key, pivot, value).Val()
}

func (c *cluster) LPop(ctx context.Context, key string) string {
	return c.client.LPop(key).Val()
}

func (c *cluster) LPushX(ctx context.Context, key string, value interface{}) int64 {
	return c.client.LPushX(key, value).Val()
}

func (c *cluster) LRange(ctx context.Context, key string, start, stop int64) []string {
	return c.client.LRange(key, start, stop).Val()
}

func (c *cluster) LRem(ctx context.Context, key string, count int64, value interface{}) int64 {
	return c.client.LRem(key, count, value).Val()
}

func (c *cluster) LSet(ctx context.Context, key string, index int64, value interface{}) error {
	return c.client.LSet(key, index, value).Err()
}

func (c *cluster) LTrim(ctx context.Context, key string, start, stop int64) error {
	return c.client.LTrim(key, start, stop).Err()
}

func (c *cluster) RPopLPush(ctx context.Context, source, destination string) string {
	return c.client.RPopLPush(source, destination).Val()
}

func (c *cluster) RPush(ctx context.Context, key string, values ...interface{}) int64 {
	return c.client.RPush(key, values...).Val()
}

func (c *cluster) RPushX(ctx context.Context, key string, value interface{}) int64 {
	return c.client.RPushX(key, value).Val()
}

func (c *cluster) SAdd(ctx context.Context, key string, members ...interface{}) int64 {
	return c.client.SAdd(key, members...).Val()
}

func (c *cluster) SCard(ctx context.Context, key string) int64 {
	return c.client.SCard(key).Val()
}

func (c *cluster) SDiff(ctx context.Context, keys ...string) []string {
	return c.client.SDiff(keys...).Val()
}

func (c *cluster) SDiffStore(ctx context.Context, destination string, keys ...string) int64 {
	return c.client.SDiffStore(destination, keys...).Val()
}

func (c *cluster) SInter(ctx context.Context, keys ...string) []string {
	return c.client.SInter(keys...).Val()
}

func (c *cluster) SInterStore(ctx context.Context, destination string, keys ...string) int64 {
	return c.client.SInterStore(destination, keys...).Val()
}

func (c *cluster) SIsMember(ctx context.Context, key string, member interface{}) bool {
	return c.client.SIsMember(key, member).Val()
}

func (c *cluster) SMembers(ctx context.Context, key string) []string {
	return c.client.SMembers(key).Val()
}

func (c *cluster) SMembersMap(ctx context.Context, key string) map[string]struct{} {
	return c.client.SMembersMap(key).Val()
}

func (c *cluster) SMove(ctx context.Context, source, destination string, member interface{}) bool {
	return c.client.SMove(source, destination, member).Val()
}

func (c *cluster) SPop(ctx context.Context, key string) string {
	return c.client.SPop(key).Val()
}

func (c *cluster) SPopN(ctx context.Context, key string, count int64) []string {
	return c.client.SPopN(key, count).Val()
}

func (c *cluster) SRandMember(ctx context.Context, key string) string {
	return c.client.SRandMember(key).Val()
}

func (c *cluster) SRandMemberN(ctx context.Context, key string, count int64) []string {
	return c.client.SRandMemberN(key, count).Val()
}

func (c *cluster) SRem(ctx context.Context, key string, members ...interface{}) int64 {
	return c.client.SRem(key, members...).Val()
}

func (c *cluster) SUnion(ctx context.Context, keys ...string) []string {
	return c.client.SUnion(keys...).Val()
}

func (c *cluster) SUnionStore(ctx context.Context, destination string, keys ...string) int64 {
	return c.client.SUnionStore(destination, keys...).Val()
}

func (c *cluster) XAdd(ctx context.Context, a *redis.XAddArgs) string {
	return c.client.XAdd(a).Val()
}

func (c *cluster) XDel(ctx context.Context, stream string, ids ...string) int64 {
	return c.client.XDel(stream, ids...).Val()
}

func (c *cluster) XLen(ctx context.Context, stream string) int64 {
	return c.client.XLen(stream).Val()
}

func (c *cluster) XRange(ctx context.Context, stream, start, stop string) []redis.XMessage {
	return c.client.XRange(stream, start, stop).Val()
}

func (c *cluster) XRangeN(ctx context.Context, stream, start, stop string, count int64) []redis.XMessage {
	return c.client.XRangeN(stream, start, stop, count).Val()
}

func (c *cluster) XRevRange(ctx context.Context, stream string, start, stop string) []redis.XMessage {
	return c.client.XRevRange(stream, start, stop).Val()
}

func (c *cluster) XRevRangeN(ctx context.Context, stream string, start, stop string, count int64) []redis.XMessage {
	return c.client.XRevRangeN(stream, start, stop, count).Val()
}

func (c *cluster) XRead(ctx context.Context, a *redis.XReadArgs) []redis.XStream {
	return c.client.XRead(a).Val()
}

func (c *cluster) XReadStreams(ctx context.Context, streams ...string) []redis.XStream {
	return c.client.XReadStreams(streams...).Val()
}

func (c *cluster) XGroupCreate(ctx context.Context, stream, group, start string) error {
	return c.client.XGroupCreate(stream, group, start).Err()
}

func (c *cluster) XGroupCreateMkStream(ctx context.Context, stream, group, start string) error {
	return c.client.XGroupCreateMkStream(stream, group, start).Err()
}

func (c *cluster) XGroupSetID(ctx context.Context, stream, group, start string) error {
	return c.client.XGroupSetID(stream, group, start).Err()
}

func (c *cluster) XGroupDestroy(ctx context.Context, stream, group string) int64 {
	return c.client.XGroupDestroy(stream, group).Val()
}

func (c *cluster) XGroupDelConsumer(ctx context.Context, stream, group, consumer string) int64 {
	return c.client.XGroupDelConsumer(stream, group, consumer).Val()
}

func (c *cluster) XReadGroup(ctx context.Context, a *redis.XReadGroupArgs) []redis.XStream {
	return c.client.XReadGroup(a).Val()
}

func (c *cluster) XAck(ctx context.Context, stream, group string, ids ...string) int64 {
	return c.client.XAck(stream, group, ids...).Val()
}

func (c *cluster) XPending(stream, group string) *redis.XPending {
	return c.client.XPending(stream, group).Val()
}

func (c *cluster) XPendingExt(ctx context.Context, a *redis.XPendingExtArgs) []redis.XPendingExt {
	return c.client.XPendingExt(a).Val()
}

func (c *cluster) XClaim(ctx context.Context, a *redis.XClaimArgs) []redis.XMessage {
	return c.client.XClaim(a).Val()
}

func (c *cluster) XClaimJustID(ctx context.Context, a *redis.XClaimArgs) []string {
	return c.client.XClaimJustID(a).Val()
}

func (c *cluster) XTrim(ctx context.Context, key string, maxLen int64) int64 {
	return c.client.XTrim(key, maxLen).Val()
}

func (c *cluster) XTrimApprox(ctx context.Context, key string, maxLen int64) int64 {
	return c.client.XTrimApprox(key, maxLen).Val()
}

func (c *cluster) BZPopMax(ctx context.Context, timeout time.Duration, keys ...string) redis.ZWithKey {
	return c.client.BZPopMax(timeout, keys...).Val()
}

func (c *cluster) BZPopMin(timeout time.Duration, keys ...string) redis.ZWithKey {
	return c.client.BZPopMin(timeout, keys...).Val()
}

func (c *cluster) ZAddNX(ctx context.Context, key string, members ...redis.Z) int64 {
	return c.client.ZAddNX(key, members...).Val()
}

func (c *cluster) ZAddXX(ctx context.Context, key string, members ...redis.Z) int64 {
	return c.client.ZAddXX(key, members...).Val()
}

func (c *cluster) ZAddCh(ctx context.Context, key string, members ...redis.Z) int64 {
	return c.client.ZAddCh(key, members...).Val()
}

func (c *cluster) ZAddNXCh(ctx context.Context, key string, members ...redis.Z) int64 {
	return c.client.ZAddNXCh(key, members...).Val()
}

func (c *cluster) ZAddXXCh(ctx context.Context, key string, members ...redis.Z) int64 {
	return c.client.ZAddXXCh(key, members...).Val()
}

func (c *cluster) ZIncr(ctx context.Context, key string, member redis.Z) float64 {
	return c.client.ZIncr(key, member).Val()
}

func (c *cluster) ZIncrNX(ctx context.Context, key string, member redis.Z) float64 {
	return c.client.ZIncrNX(key, member).Val()
}

func (c *cluster) ZIncrXX(ctx context.Context, key string, member redis.Z) float64 {
	return c.client.ZIncrXX(key, member).Val()
}

func (c *cluster) ZCount(ctx context.Context, key, min, max string) int64 {
	return c.client.ZCount(key, min, max).Val()
}

func (c *cluster) ZLexCount(ctx context.Context, key, min, max string) int64 {
	return c.client.ZLexCount(key, min, max).Val()
}

func (c *cluster) ZIncrBy(ctx context.Context, key string, increment float64, member string) float64 {
	return c.client.ZIncrBy(key, increment, member).Val()
}

func (c *cluster) ZInterStore(ctx context.Context, destination string, store redis.ZStore, keys ...string) int64 {
	return c.client.ZInterStore(destination, store, keys...).Val()
}

func (c *cluster) ZPopMax(ctx context.Context, key string, count ...int64) []redis.Z {
	return c.client.ZPopMax(key, count...).Val()
}

func (c *cluster) ZPopMin(ctx context.Context, key string, count ...int64) []redis.Z {
	return c.client.ZPopMin(key, count...).Val()
}

func (c *cluster) ZRange(ctx context.Context, key string, start, stop int64) []string {
	return c.client.ZRange(key, start, stop).Val()
}

func (c *cluster) ZRangeByScore(ctx context.Context, key string, opt redis.ZRangeBy) []string {
	return c.client.ZRangeByScore(key, opt).Val()
}

func (c *cluster) ZRangeByLex(ctx context.Context, key string, opt redis.ZRangeBy) []string {
	return c.client.ZRangeByLex(key, opt).Val()
}

func (c *cluster) ZRangeByScoreWithScores(ctx context.Context, key string, opt redis.ZRangeBy) []redis.Z {
	return c.client.ZRangeByScoreWithScores(key, opt).Val()
}

func (c *cluster) ZRank(ctx context.Context, key, member string) (int64, error) {
	return c.client.ZRank(key, member).Result()
}

func (c *cluster) ZRem(ctx context.Context, key string, members ...interface{}) int64 {
	return c.client.ZRem(key, members...).Val()
}

func (c *cluster) ZRemRangeByRank(ctx context.Context, key string, start, stop int64) (int64, error) {
	return c.client.ZRemRangeByRank(key, start, stop).Result()
}

func (c *cluster) ZRemRangeByScore(ctx context.Context, key, min, max string) int64 {
	return c.client.ZRemRangeByScore(key, min, max).Val()
}

func (c *cluster) ZRemRangeByLex(ctx context.Context, key, min, max string) int64 {
	return c.client.ZRemRangeByLex(key, min, max).Val()
}

func (c *cluster) ZRevRange(ctx context.Context, key string, start, stop int64) []string {
	return c.client.ZRevRange(key, start, stop).Val()
}

func (c *cluster) ZRevRangeWithScores(ctx context.Context, key string, start, stop int64) []redis.Z {
	return c.client.ZRevRangeWithScores(key, start, stop).Val()
}

func (c *cluster) ZRevRangeByScore(ctx context.Context, key string, opt redis.ZRangeBy) []string {
	return c.client.ZRevRangeByScore(key, opt).Val()
}

func (c *cluster) ZRevRangeByLex(ctx context.Context, key string, opt redis.ZRangeBy) []string {
	return c.client.ZRevRangeByLex(key, opt).Val()
}

func (c *cluster) ZRevRangeByScoreWithScores(ctx context.Context, key string, opt redis.ZRangeBy) []redis.Z {
	return c.client.ZRevRangeByScoreWithScores(key, opt).Val()
}

func (c *cluster) ZRevRank(ctx context.Context, key, member string) (int64, error) {
	return c.client.ZRevRank(key, member).Result()
}

func (c *cluster) ZScore(ctx context.Context, key, member string) float64 {
	return c.client.ZScore(key, member).Val()
}

func (c *cluster) ZUnionStore(ctx context.Context, dest string, store redis.ZStore, keys ...string) int64 {
	return c.client.ZUnionStore(dest, store, keys...).Val()
}

func (c *cluster) PFAdd(ctx context.Context, key string, els ...interface{}) int64 {
	return c.client.PFAdd(key, els...).Val()
}

func (c *cluster) PFCount(ctx context.Context, keys ...string) int64 {
	return c.client.PFCount(keys...).Val()
}

func (c *cluster) PFMerge(ctx context.Context, dest string, keys ...string) error {
	return c.client.PFMerge(dest, keys...).Err()
}

func (c *cluster) BgRewriteAOF(ctx context.Context) error {
	return c.client.BgRewriteAOF().Err()
}

func (c *cluster) BgSave(ctx context.Context) error {
	return c.client.BgSave().Err()
}

func (c *cluster) ClientKill(ctx context.Context, ipPort string) error {
	return c.client.ClientKill(ipPort).Err()
}

func (c *cluster) ClientKillByFilter(ctx context.Context, keys ...string) int64 {
	return c.client.ClientKillByFilter(keys...).Val()
}

func (c *cluster) ClientList(ctx context.Context) string {
	return c.client.ClientList().Val()
}

func (c *cluster) ClientPause(ctx context.Context, dur time.Duration) bool {
	return c.client.ClientPause(dur).Val()
}

func (c *cluster) ClientID(ctx context.Context) int64 {
	return c.client.ClientID().Val()
}

func (c *cluster) ConfigGet(ctx context.Context, parameter string) []interface{} {
	return c.client.ConfigGet(parameter).Val()
}

func (c *cluster) ConfigResetStat(ctx context.Context) error {
	return c.client.ConfigResetStat().Err()
}

func (c *cluster) ConfigSet(ctx context.Context, parameter, value string) error {
	return c.client.ConfigSet(parameter, value).Err()
}

func (c *cluster) ConfigRewrite(ctx context.Context) error {
	return c.client.ConfigRewrite().Err()
}

func (c *cluster) DBSize(ctx context.Context) int64 {
	return c.client.DBSize().Val()
}

func (c *cluster) FlushAll(ctx context.Context) error {
	return c.client.FlushAll().Err()
}

func (c *cluster) FlushAllAsync(ctx context.Context) error {
	return c.client.FlushAllAsync().Err()
}

func (c *cluster) FlushDB(ctx context.Context) error {
	return c.client.FlushDB().Err()
}

func (c *cluster) FlushDBAsync(ctx context.Context) error {
	return c.client.FlushDBAsync().Err()
}

func (c *cluster) Info(ctx context.Context, section ...string) string {
	return c.client.Info(section...).Val()
}

func (c *cluster) LastSave(ctx context.Context) int64 {
	return c.client.LastSave().Val()
}

func (c *cluster) Save(ctx context.Context) error {
	return c.client.Save().Err()
}

func (c *cluster) Shutdown(ctx context.Context) error {
	return c.client.Shutdown().Err()
}

func (c *cluster) ShutdownSave(ctx context.Context) error {
	return c.client.ShutdownSave().Err()
}

func (c *cluster) ShutdownNoSave(ctx context.Context) error {
	return c.client.ShutdownNoSave().Err()
}

func (c *cluster) SlaveOf(ctx context.Context, host, port string) error {
	return c.client.SlaveOf(host, port).Err()
}

func (c *cluster) Time(ctx context.Context) time.Time {
	return c.client.Time().Val()
}

func (c *cluster) Eval(ctx context.Context, script string, keys []string, args ...interface{}) *redis.Cmd {
	return c.client.Eval(script, keys, args...)
}

func (c *cluster) EvalSha(ctx context.Context, sha1 string, keys []string, args ...interface{}) *redis.Cmd {
	return c.client.EvalSha(sha1, keys, args...)
}

func (c *cluster) ScriptExists(ctx context.Context, hashes ...string) []bool {
	return c.client.ScriptExists(hashes...).Val()
}

func (c *cluster) ScriptFlush(ctx context.Context) error {
	return c.client.ScriptFlush().Err()
}

func (c *cluster) ScriptKill(ctx context.Context) error {
	return c.client.ScriptKill().Err()
}

func (c *cluster) ScriptLoad(ctx context.Context, script string) string {
	return c.client.ScriptLoad(script).Val()
}

func (c *cluster) DebugObject(ctx context.Context, key string) string {
	return c.client.DebugObject(key).Val()
}

func (c *cluster) PubSubChannels(ctx context.Context, pattern string) []string {
	return c.client.PubSubChannels(pattern).Val()
}

func (c *cluster) PubSubNumSub(ctx context.Context, channels ...string) map[string]int64 {
	return c.client.PubSubNumSub(channels...).Val()
}

func (c *cluster) PubSubNumPat(ctx context.Context) int64 {
	return c.client.PubSubNumPat().Val()
}

func (c *cluster) ClusterSlots(ctx context.Context) []redis.ClusterSlot {
	return c.client.ClusterSlots().Val()
}

func (c *cluster) ClusterNodes(ctx context.Context) string {
	return c.client.ClusterNodes().Val()
}

func (c *cluster) ClusterMeet(ctx context.Context, host, port string) error {
	return c.client.ClusterMeet(host, port).Err()
}

func (c *cluster) ClusterForget(ctx context.Context, nodeID string) error {
	return c.client.ClusterForget(nodeID).Err()
}

func (c *cluster) ClusterReplicate(ctx context.Context, nodeID string) error {
	return c.client.ClusterReplicate(nodeID).Err()
}

func (c *cluster) ClusterResetSoft(ctx context.Context) error {
	return c.client.ClusterResetSoft().Err()
}

func (c *cluster) ClusterResetHard(ctx context.Context) error {
	return c.client.ClusterResetHard().Err()
}

func (c *cluster) ClusterInfo(ctx context.Context) string {
	return c.client.ClusterInfo().Val()
}

func (c *cluster) ClusterKeySlot(ctx context.Context, key string) int64 {
	return c.client.ClusterKeySlot(key).Val()
}

func (c *cluster) ClusterGetKeysInSlot(ctx context.Context, slot int, count int) []string {
	return c.client.ClusterGetKeysInSlot(slot, count).Val()
}

func (c *cluster) ClusterCountFailureReports(ctx context.Context, nodeID string) int64 {
	return c.client.ClusterCountFailureReports(nodeID).Val()
}

func (c *cluster) ClusterCountKeysInSlot(ctx context.Context, slot int) int64 {
	return c.client.ClusterCountKeysInSlot(slot).Val()
}

func (c *cluster) ClusterDelSlots(ctx context.Context, slots ...int) error {
	return c.client.ClusterDelSlots(slots...).Err()
}

func (c *cluster) ClusterDelSlotsRange(ctx context.Context, min, max int) error {
	return c.client.ClusterDelSlotsRange(min, max).Err()
}

func (c *cluster) ClusterSaveConfig(ctx context.Context) error {
	return c.client.ClusterSaveConfig().Err()
}

func (c *cluster) ClusterSlaves(ctx context.Context, nodeID string) []string {
	return c.client.ClusterSlaves(nodeID).Val()
}

func (c *cluster) ClusterFailover(ctx context.Context) error {
	return c.client.ClusterFailover().Err()
}

func (c *cluster) ClusterAddSlots(ctx context.Context, slots ...int) error {
	return c.client.ClusterAddSlots(slots...).Err()
}

func (c *cluster) ClusterAddSlotsRange(ctx context.Context, min, max int) error {
	return c.client.ClusterAddSlotsRange(min, max).Err()
}

func (c *cluster) GeoAdd(ctx context.Context, key string, geoLocation ...*redis.GeoLocation) int64 {
	return c.client.GeoAdd(key, geoLocation...).Val()
}

func (c *cluster) GeoPos(ctx context.Context, key string, members ...string) []*redis.GeoPos {
	return c.client.GeoPos(key, members...).Val()
}

func (c *cluster) GeoRadius(ctx context.Context, key string, longitude, latitude float64, query *redis.GeoRadiusQuery) []redis.GeoLocation {
	return c.client.GeoRadius(key, longitude, latitude, query).Val()
}

func (c *cluster) GeoRadiusRO(ctx context.Context, key string, longitude, latitude float64, query *redis.GeoRadiusQuery) []redis.GeoLocation {
	return c.client.GeoRadiusRO(key, longitude, latitude, query).Val()
}

func (c *cluster) GeoRadiusByMember(ctx context.Context, key, member string, query *redis.GeoRadiusQuery) []redis.GeoLocation {
	return c.client.GeoRadiusByMember(key, member, query).Val()
}

func (c *cluster) GeoRadiusByMemberRO(ctx context.Context, key, member string, query *redis.GeoRadiusQuery) []redis.GeoLocation {
	return c.client.GeoRadiusByMemberRO(key, member, query).Val()
}

func (c *cluster) GeoDist(ctx context.Context, key string, member1, member2, unit string) float64 {
	return c.client.GeoDist(key, member1, member2, unit).Val()
}

func (c *cluster) GeoHash(ctx context.Context, key string, members ...string) []string {
	return c.client.GeoHash(key, members...).Val()
}

func (c *cluster) ReadOnly(ctx context.Context) error {
	return c.client.ReadOnly().Err()
}

func (c *cluster) ReadWrite(ctx context.Context) error {
	return c.client.ReadWrite().Err()
}

func (c *cluster) MemoryUsage(ctx context.Context, key string, samples ...int) int64 {
	return c.client.MemoryUsage(key, samples...).Val()
}

func (c *cluster) Ping(_ context.Context) error {
	return c.client.Ping().Err()
}

func (c *cluster) LPush(_ context.Context, key string, val interface{}) (err error) {
	return c.client.LPush(key, val).Err()
}

func (c *cluster) RPop(_ context.Context, key string) (res string, err error) {
	return c.client.RPop(key).Result()
}

func (c *cluster) LLen(_ context.Context, key string) int64 {
	return c.client.LLen(key).Val()
}

func (c *cluster) TypeOf(_ context.Context, key string) (res string, err error) {
	return c.client.Type(key).Result()
}

func (c *cluster) Keys(_ context.Context, pattern string) (res []string, err error) {
	return c.client.Keys(pattern).Result()
}

func (c *cluster) ZAdd(_ context.Context, k string, score float64, member interface{}) (err error) {
	return c.client.ZAdd(c.setPrefix(k), redis.Z{
		Score:  score,
		Member: member,
	}).Err()
}

func (c *cluster) ZRangeWithScores(_ context.Context, k string, start, stop int64) (res []redis.Z, err error) {
	return c.client.ZRangeWithScores(c.setPrefix(k), start, stop).Result()
}

func (c *cluster) ZCard(_ context.Context, k string) (res int64, err error) {
	return c.client.ZCard(c.setPrefix(k)).Result()
}

func (c *cluster) HLen(_ context.Context, k string) (res int64, err error) {
	return c.client.HLen(c.setPrefix(k)).Result()
}

func (c *cluster) HGetAll(_ context.Context, k string) (res map[string]string, err error) {
	return c.client.HGetAll(c.setPrefix(k)).Result()
}

func (c *cluster) Incr(_ context.Context, key string, exp time.Duration) error {
	defer func() {
		c.client.Expire(c.setPrefix(key), exp)
	}()
	return c.client.Incr(c.setPrefix(key)).Err()
}

func NewRedisCluster(hosts []string, password, prefix string) RedisClient {
	return &cluster{client: redis.NewClusterClient(&redis.ClusterOptions{
		Addrs:    hosts,
		Password: password,
	}), prefix: prefix + ""}
}

func (c *cluster) Exists(_ context.Context, keys ...string) int64 {
	return c.client.Exists(keys...).Val()
}

func (c *cluster) TTL(_ context.Context, key string) time.Duration {
	return c.client.TTL(key).Val()
}

func (c *cluster) Set(_ context.Context, k string, v interface{}, expir ...time.Duration) (err error) {
	var val string
	switch v.(type) {
	case string:
		val = v.(string)
	default:
		b, _ := json.Marshal(v)
		val = string(b)
	}

	exp := expiration
	if len(expir) == 1 {
		exp = expir[0]
	}

	return c.client.Set(c.setPrefix(k), val, exp).Err()
}

func (c *cluster) Get(_ context.Context, k string) (v string, err error) {
	return c.client.Get(c.setPrefix(k)).Result()
}

func (c *cluster) Del(_ context.Context, k string) (err error) {
	return c.client.Del(c.setPrefix(k)).Err()
}

func (c *cluster) HSet(_ context.Context, k string, field string, v interface{}) (err error) {
	var val string
	switch v.(type) {
	case string:
		val = v.(string)
	default:
		b, _ := json.Marshal(v)
		val = string(b)
	}
	return c.client.HSet(c.setPrefix(k), field, val).Err()
}

func (c *cluster) HGet(_ context.Context, k string, field string) (res string, err error) {
	return c.client.HGet(c.setPrefix(k), field).Result()
}

func (c *cluster) HDelAll(_ context.Context, k string) (err error) {
	res, err := c.client.HKeys(c.setPrefix(k)).Result()
	if err != nil {
		return
	}
	return c.client.HDel(c.setPrefix(k), res...).Err()
}

func (c *cluster) HDel(_ context.Context, k string, field string) (err error) {
	return c.client.HDel(c.setPrefix(k), field).Err()
}

func (c *cluster) setPrefix(s string) string {
	return c.prefix + s
}

func (c *cluster) Close(_ context.Context) error {
	return c.client.Close()
}

func (c *cluster) Subscribe(_ context.Context, channels ...string) *redis.PubSub {
	return c.client.Subscribe(channels...)
}

func (c *cluster) Publish(_ context.Context, channel string, message interface{}) error {
	return c.client.Publish(channel, message).Err()
}

func (c *cluster) SetPrefix(_ context.Context, prefix string) RedisClient {
	c.prefix = prefix
	return c
}
