/**
 * @Time: 2020/3/30 10:39
 * @Author: solacowa@gmail.com
 * @File: single
 * @Software: GoLand
 */

package redisclient

import (
	"context"
	"encoding/json"
	"time"

	"github.com/go-redis/redis"
)

type single struct {
	client *redis.Client
	prefix string
}

func (c *single) Pipeline(ctx context.Context) redis.Pipeliner {
	panic("implement me")
}

func (c *single) Pipelined(ctx context.Context, fn func(pipeliner redis.Pipeliner) error) ([]redis.Cmder, error) {
	panic("implement me")
}

func (c *single) TxPipelined(ctx context.Context, fn func(pipeliner redis.Pipeliner) error) ([]redis.Cmder, error) {
	panic("implement me")
}

func (c *single) TxPipeline(ctx context.Context) redis.Pipeliner {
	panic("implement me")
}

func (c *single) Command(ctx context.Context) *redis.CommandsInfoCmd {
	panic("implement me")
}

func (c *single) ClientGetName(ctx context.Context) string {
	panic("implement me")
}

func (c *single) Echo(ctx context.Context, message interface{}) string {
	panic("implement me")
}

func (c *single) Quit(ctx context.Context) error {
	panic("implement me")
}

func (c *single) Unlink(ctx context.Context, keys ...string) int {
	panic("implement me")
}

func (c *single) Dump(ctx context.Context, key string) string {
	panic("implement me")
}

func (c *single) Expire(ctx context.Context, key string, expiration time.Duration) bool {
	panic("implement me")
}

func (c *single) ExpireAt(ctx context.Context, key string, tm time.Time) bool {
	panic("implement me")
}

func (c *single) Migrate(ctx context.Context, host, port, key string, db int64, timeout time.Duration) error {
	panic("implement me")
}

func (c *single) Move(ctx context.Context, key string, db int64) bool {
	panic("implement me")
}

func (c *single) ObjectRefCount(ctx context.Context, key string) int {
	panic("implement me")
}

func (c *single) ObjectEncoding(ctx context.Context, key string) string {
	panic("implement me")
}

func (c *single) ObjectIdleTime(ctx context.Context, key string) time.Duration {
	panic("implement me")
}

func (c *single) Persist(ctx context.Context, key string) bool {
	panic("implement me")
}

func (c *single) PExpire(ctx context.Context, key string, expiration time.Duration) bool {
	panic("implement me")
}

func (c *single) PExpireAt(ctx context.Context, key string, tm time.Time) bool {
	panic("implement me")
}

func (c *single) PTTL(ctx context.Context, key string) time.Duration {
	panic("implement me")
}

func (c *single) RandomKey(ctx context.Context) string {
	panic("implement me")
}

func (c *single) Rename(ctx context.Context, key, newkey string) *redis.StatusCmd {
	panic("implement me")
}

func (c *single) RenameNX(ctx context.Context, key, newkey string) bool {
	panic("implement me")
}

func (c *single) Restore(ctx context.Context, key string, ttl time.Duration, value string) error {
	panic("implement me")
}

func (c *single) RestoreReplace(ctx context.Context, key string, ttl time.Duration, value string) error {
	panic("implement me")
}

func (c *single) Sort(ctx context.Context, key string, sort *redis.Sort) []string {
	panic("implement me")
}

func (c *single) SortStore(ctx context.Context, key, store string, sort *redis.Sort) int {
	panic("implement me")
}

func (c *single) SortInterfaces(ctx context.Context, key string, sort *redis.Sort) []interface{} {
	panic("implement me")
}

func (c *single) Touch(ctx context.Context, keys ...string) int {
	panic("implement me")
}

func (c *single) Type(ctx context.Context, key string) error {
	panic("implement me")
}

func (c *single) Scan(ctx context.Context, cursor uint64, match string, count int64) *redis.ScanCmd {
	panic("implement me")
}

func (c *single) SScan(ctx context.Context, key string, cursor uint64, match string, count int64) *redis.ScanCmd {
	panic("implement me")
}

func (c *single) HScan(ctx context.Context, key string, cursor uint64, match string, count int64) *redis.ScanCmd {
	panic("implement me")
}

func (c *single) ZScan(ctx context.Context, key string, cursor uint64, match string, count int64) *redis.ScanCmd {
	panic("implement me")
}

func (c *single) Append(ctx context.Context, key, value string) int {
	panic("implement me")
}

func (c *single) BitCount(ctx context.Context, key string, bitCount *redis.BitCount) int {
	panic("implement me")
}

func (c *single) BitOpAnd(ctx context.Context, destKey string, keys ...string) int {
	panic("implement me")
}

func (c *single) BitOpOr(ctx context.Context, destKey string, keys ...string) int {
	panic("implement me")
}

func (c *single) BitOpXor(ctx context.Context, destKey string, keys ...string) int {
	panic("implement me")
}

func (c *single) BitOpNot(ctx context.Context, destKey string, key string) int {
	panic("implement me")
}

func (c *single) BitPos(ctx context.Context, key string, bit int64, pos ...int64) int {
	panic("implement me")
}

func (c *single) Decr(ctx context.Context, key string) int {
	panic("implement me")
}

func (c *single) DecrBy(ctx context.Context, key string, decrement int64) int {
	panic("implement me")
}

func (c *single) GetBit(ctx context.Context, key string, offset int64) int {
	panic("implement me")
}

func (c *single) GetRange(ctx context.Context, key string, start, end int64) string {
	panic("implement me")
}

func (c *single) GetSet(ctx context.Context, key string, value interface{}) string {
	panic("implement me")
}

func (c *single) IncrBy(ctx context.Context, key string, value int64) int {
	panic("implement me")
}

func (c *single) IncrByFloat(ctx context.Context, key string, value float64) float64 {
	panic("implement me")
}

func (c *single) MGet(ctx context.Context, keys ...string) []interface{} {
	panic("implement me")
}

func (c *single) MSet(ctx context.Context, pairs ...interface{}) error {
	panic("implement me")
}

func (c *single) MSetNX(ctx context.Context, pairs ...interface{}) bool {
	panic("implement me")
}

func (c *single) SetBit(ctx context.Context, key string, offset int64, value int) int {
	panic("implement me")
}

func (c *single) SetNX(ctx context.Context, key string, value interface{}, expiration time.Duration) bool {
	panic("implement me")
}

func (c *single) SetXX(ctx context.Context, key string, value interface{}, expiration time.Duration) bool {
	panic("implement me")
}

func (c *single) SetRange(ctx context.Context, key string, offset int64, value string) int {
	panic("implement me")
}

func (c *single) StrLen(ctx context.Context, key string) int {
	panic("implement me")
}

func (c *single) HExists(ctx context.Context, key, field string) bool {
	panic("implement me")
}

func (c *single) HIncrBy(ctx context.Context, key, field string, incr int64) int {
	panic("implement me")
}

func (c *single) HIncrByFloat(ctx context.Context, key, field string, incr float64) float64 {
	panic("implement me")
}

func (c *single) HKeys(ctx context.Context, key string) []string {
	panic("implement me")
}

func (c *single) HMGet(ctx context.Context, key string, fields ...string) []interface{} {
	panic("implement me")
}

func (c *single) HMSet(ctx context.Context, key string, fields map[string]interface{}) error {
	panic("implement me")
}

func (c *single) HSetNX(ctx context.Context, key, field string, value interface{}) bool {
	panic("implement me")
}

func (c *single) HVals(ctx context.Context, key string) []string {
	panic("implement me")
}

func (c *single) BLPop(ctx context.Context, timeout time.Duration, keys ...string) []string {
	panic("implement me")
}

func (c *single) BRPop(ctx context.Context, timeout time.Duration, keys ...string) []string {
	panic("implement me")
}

func (c *single) BRPopLPush(ctx context.Context, source, destination string, timeout time.Duration) string {
	panic("implement me")
}

func (c *single) LIndex(ctx context.Context, key string, index int64) string {
	panic("implement me")
}

func (c *single) LInsert(ctx context.Context, key, op string, pivot, value interface{}) int {
	panic("implement me")
}

func (c *single) LInsertBefore(ctx context.Context, key string, pivot, value interface{}) int {
	panic("implement me")
}

func (c *single) LInsertAfter(ctx context.Context, key string, pivot, value interface{}) int {
	panic("implement me")
}

func (c *single) LPop(ctx context.Context, key string) string {
	panic("implement me")
}

func (c *single) LPushX(ctx context.Context, key string, value interface{}) int {
	panic("implement me")
}

func (c *single) LRange(ctx context.Context, key string, start, stop int64) []string {
	panic("implement me")
}

func (c *single) LRem(ctx context.Context, key string, count int64, value interface{}) int {
	panic("implement me")
}

func (c *single) LSet(ctx context.Context, key string, index int64, value interface{}) error {
	panic("implement me")
}

func (c *single) LTrim(ctx context.Context, key string, start, stop int64) error {
	panic("implement me")
}

func (c *single) RPopLPush(ctx context.Context, source, destination string) string {
	panic("implement me")
}

func (c *single) RPush(ctx context.Context, key string, values ...interface{}) int {
	panic("implement me")
}

func (c *single) RPushX(ctx context.Context, key string, value interface{}) int {
	panic("implement me")
}

func (c *single) SAdd(ctx context.Context, key string, members ...interface{}) int {
	panic("implement me")
}

func (c *single) SCard(ctx context.Context, key string) int {
	panic("implement me")
}

func (c *single) SDiff(ctx context.Context, keys ...string) []string {
	panic("implement me")
}

func (c *single) SDiffStore(ctx context.Context, destination string, keys ...string) int {
	panic("implement me")
}

func (c *single) SInter(ctx context.Context, keys ...string) []string {
	panic("implement me")
}

func (c *single) SInterStore(ctx context.Context, destination string, keys ...string) int {
	panic("implement me")
}

func (c *single) SIsMember(ctx context.Context, key string, member interface{}) bool {
	panic("implement me")
}

func (c *single) SMembers(ctx context.Context, key string) []string {
	panic("implement me")
}

func (c *single) SMembersMap(ctx context.Context, key string) map[string]struct{} {
	panic("implement me")
}

func (c *single) SMove(ctx context.Context, source, destination string, member interface{}) bool {
	panic("implement me")
}

func (c *single) SPop(ctx context.Context, key string) string {
	panic("implement me")
}

func (c *single) SPopN(ctx context.Context, key string, count int64) []string {
	panic("implement me")
}

func (c *single) SRandMember(ctx context.Context, key string) string {
	panic("implement me")
}

func (c *single) SRandMemberN(ctx context.Context, key string, count int64) []string {
	panic("implement me")
}

func (c *single) SRem(ctx context.Context, key string, members ...interface{}) int {
	panic("implement me")
}

func (c *single) SUnion(ctx context.Context, keys ...string) []string {
	panic("implement me")
}

func (c *single) SUnionStore(ctx context.Context, destination string, keys ...string) int {
	panic("implement me")
}

func (c *single) XAdd(ctx context.Context, a *redis.XAddArgs) string {
	panic("implement me")
}

func (c *single) XDel(ctx context.Context, stream string, ids ...string) int {
	panic("implement me")
}

func (c *single) XLen(ctx context.Context, stream string) int {
	panic("implement me")
}

func (c *single) XRange(ctx context.Context, stream, start, stop string) []redis.XMessage {
	panic("implement me")
}

func (c *single) XRangeN(ctx context.Context, stream, start, stop string, count int64) []redis.XMessage {
	panic("implement me")
}

func (c *single) XRevRange(ctx context.Context, stream string, start, stop string) []redis.XMessage {
	panic("implement me")
}

func (c *single) XRevRangeN(ctx context.Context, stream string, start, stop string, count int64) []redis.XMessage {
	panic("implement me")
}

func (c *single) XRead(ctx context.Context, a *redis.XReadArgs) []redis.XStream {
	panic("implement me")
}

func (c *single) XReadStreams(ctx context.Context, streams ...string) []redis.XStream {
	panic("implement me")
}

func (c *single) XGroupCreate(ctx context.Context, stream, group, start string) error {
	panic("implement me")
}

func (c *single) XGroupCreateMkStream(ctx context.Context, stream, group, start string) error {
	panic("implement me")
}

func (c *single) XGroupSetID(ctx context.Context, stream, group, start string) error {
	panic("implement me")
}

func (c *single) XGroupDestroy(ctx context.Context, stream, group string) int {
	panic("implement me")
}

func (c *single) XGroupDelConsumer(ctx context.Context, stream, group, consumer string) int {
	panic("implement me")
}

func (c *single) XReadGroup(ctx context.Context, a *redis.XReadGroupArgs) []redis.XStream {
	panic("implement me")
}

func (c *single) XAck(ctx context.Context, stream, group string, ids ...string) int {
	panic("implement me")
}

func (c *single) XPending(stream, group string) *redis.XPending {
	panic("implement me")
}

func (c *single) XPendingExt(ctx context.Context, a *redis.XPendingExtArgs) []redis.XPendingExt {
	panic("implement me")
}

func (c *single) XClaim(ctx context.Context, a *redis.XClaimArgs) []redis.XMessage {
	panic("implement me")
}

func (c *single) XClaimJustID(ctx context.Context, a *redis.XClaimArgs) []string {
	panic("implement me")
}

func (c *single) XTrim(ctx context.Context, key string, maxLen int64) int {
	panic("implement me")
}

func (c *single) XTrimApprox(ctx context.Context, key string, maxLen int64) int {
	panic("implement me")
}

func (c *single) BZPopMax(ctx context.Context, timeout time.Duration, keys ...string) redis.ZWithKey {
	panic("implement me")
}

func (c *single) BZPopMin(timeout time.Duration, keys ...string) redis.ZWithKey {
	panic("implement me")
}

func (c *single) ZAddNX(ctx context.Context, key string, members ...redis.Z) int {
	panic("implement me")
}

func (c *single) ZAddXX(ctx context.Context, key string, members ...redis.Z) int {
	panic("implement me")
}

func (c *single) ZAddCh(ctx context.Context, key string, members ...redis.Z) int {
	panic("implement me")
}

func (c *single) ZAddNXCh(ctx context.Context, key string, members ...redis.Z) int {
	panic("implement me")
}

func (c *single) ZAddXXCh(ctx context.Context, key string, members ...redis.Z) int {
	panic("implement me")
}

func (c *single) ZIncr(ctx context.Context, key string, member redis.Z) float64 {
	panic("implement me")
}

func (c *single) ZIncrNX(ctx context.Context, key string, member redis.Z) float64 {
	panic("implement me")
}

func (c *single) ZIncrXX(ctx context.Context, key string, member redis.Z) float64 {
	panic("implement me")
}

func (c *single) ZCount(ctx context.Context, key, min, max string) int {
	panic("implement me")
}

func (c *single) ZLexCount(ctx context.Context, key, min, max string) int {
	panic("implement me")
}

func (c *single) ZIncrBy(ctx context.Context, key string, increment float64, member string) float64 {
	panic("implement me")
}

func (c *single) ZInterStore(ctx context.Context, destination string, store redis.ZStore, keys ...string) int {
	panic("implement me")
}

func (c *single) ZPopMax(ctx context.Context, key string, count ...int64) []redis.Z {
	panic("implement me")
}

func (c *single) ZPopMin(ctx context.Context, key string, count ...int64) []redis.Z {
	panic("implement me")
}

func (c *single) ZRange(ctx context.Context, key string, start, stop int64) []string {
	panic("implement me")
}

func (c *single) ZRangeByScore(ctx context.Context, key string, opt redis.ZRangeBy) []string {
	panic("implement me")
}

func (c *single) ZRangeByLex(ctx context.Context, key string, opt redis.ZRangeBy) []string {
	panic("implement me")
}

func (c *single) ZRangeByScoreWithScores(ctx context.Context, key string, opt redis.ZRangeBy) []redis.Z {
	panic("implement me")
}

func (c *single) ZRank(ctx context.Context, key, member string) int {
	panic("implement me")
}

func (c *single) ZRem(ctx context.Context, key string, members ...interface{}) int {
	panic("implement me")
}

func (c *single) ZRemRangeByRank(ctx context.Context, key string, start, stop int64) int {
	panic("implement me")
}

func (c *single) ZRemRangeByScore(ctx context.Context, key, min, max string) int {
	panic("implement me")
}

func (c *single) ZRemRangeByLex(ctx context.Context, key, min, max string) int {
	panic("implement me")
}

func (c *single) ZRevRange(ctx context.Context, key string, start, stop int64) []string {
	panic("implement me")
}

func (c *single) ZRevRangeWithScores(ctx context.Context, key string, start, stop int64) []redis.Z {
	panic("implement me")
}

func (c *single) ZRevRangeByScore(ctx context.Context, key string, opt redis.ZRangeBy) []string {
	panic("implement me")
}

func (c *single) ZRevRangeByLex(ctx context.Context, key string, opt redis.ZRangeBy) []string {
	panic("implement me")
}

func (c *single) ZRevRangeByScoreWithScores(ctx context.Context, key string, opt redis.ZRangeBy) []redis.Z {
	panic("implement me")
}

func (c *single) ZRevRank(ctx context.Context, key, member string) int {
	panic("implement me")
}

func (c *single) ZScore(ctx context.Context, key, member string) float64 {
	panic("implement me")
}

func (c *single) ZUnionStore(ctx context.Context, dest string, store redis.ZStore, keys ...string) int {
	panic("implement me")
}

func (c *single) PFAdd(ctx context.Context, key string, els ...interface{}) int {
	panic("implement me")
}

func (c *single) PFCount(ctx context.Context, keys ...string) int {
	panic("implement me")
}

func (c *single) PFMerge(ctx context.Context, dest string, keys ...string) error {
	panic("implement me")
}

func (c *single) BgRewriteAOF(ctx context.Context) error {
	panic("implement me")
}

func (c *single) BgSave(ctx context.Context) error {
	panic("implement me")
}

func (c *single) ClientKill(ctx context.Context, ipPort string) error {
	panic("implement me")
}

func (c *single) ClientKillByFilter(ctx context.Context, keys ...string) int {
	panic("implement me")
}

func (c *single) ClientList(ctx context.Context) string {
	panic("implement me")
}

func (c *single) ClientPause(ctx context.Context, dur time.Duration) bool {
	panic("implement me")
}

func (c *single) ClientID(ctx context.Context) int {
	panic("implement me")
}

func (c *single) ConfigGet(ctx context.Context, parameter string) []interface{} {
	panic("implement me")
}

func (c *single) ConfigResetStat(ctx context.Context) error {
	panic("implement me")
}

func (c *single) ConfigSet(ctx context.Context, parameter, value string) error {
	panic("implement me")
}

func (c *single) ConfigRewrite(ctx context.Context) error {
	panic("implement me")
}

func (c *single) DBSize(ctx context.Context) int {
	panic("implement me")
}

func (c *single) FlushAll(ctx context.Context) error {
	panic("implement me")
}

func (c *single) FlushAllAsync(ctx context.Context) error {
	panic("implement me")
}

func (c *single) FlushDB(ctx context.Context) error {
	panic("implement me")
}

func (c *single) FlushDBAsync(ctx context.Context) error {
	panic("implement me")
}

func (c *single) Info(ctx context.Context, section ...string) string {
	panic("implement me")
}

func (c *single) LastSave(ctx context.Context) int {
	panic("implement me")
}

func (c *single) Save(ctx context.Context) error {
	panic("implement me")
}

func (c *single) Shutdown(ctx context.Context) error {
	panic("implement me")
}

func (c *single) ShutdownSave(ctx context.Context) error {
	panic("implement me")
}

func (c *single) ShutdownNoSave(ctx context.Context) error {
	panic("implement me")
}

func (c *single) SlaveOf(ctx context.Context, host, port string) error {
	panic("implement me")
}

func (c *single) Time(ctx context.Context) time.Time {
	panic("implement me")
}

func (c *single) Eval(ctx context.Context, script string, keys []string, args ...interface{}) *redis.Cmd {
	panic("implement me")
}

func (c *single) EvalSha(ctx context.Context, sha1 string, keys []string, args ...interface{}) *redis.Cmd {
	panic("implement me")
}

func (c *single) ScriptExists(ctx context.Context, hashes ...string) []bool {
	panic("implement me")
}

func (c *single) ScriptFlush(ctx context.Context) error {
	panic("implement me")
}

func (c *single) ScriptKill(ctx context.Context) error {
	panic("implement me")
}

func (c *single) ScriptLoad(ctx context.Context, script string) string {
	panic("implement me")
}

func (c *single) DebugObject(ctx context.Context, key string) string {
	panic("implement me")
}

func (c *single) PubSubChannels(ctx context.Context, pattern string) []string {
	panic("implement me")
}

func (c *single) PubSubNumSub(ctx context.Context, channels ...string) map[string]int64 {
	panic("implement me")
}

func (c *single) PubSubNumPat(ctx context.Context) int {
	panic("implement me")
}

func (c *single) ClusterSlots(ctx context.Context) []redis.ClusterSlot {
	panic("implement me")
}

func (c *single) ClusterNodes(ctx context.Context) string {
	panic("implement me")
}

func (c *single) ClusterMeet(ctx context.Context, host, port string) error {
	panic("implement me")
}

func (c *single) ClusterForget(ctx context.Context, nodeID string) error {
	panic("implement me")
}

func (c *single) ClusterReplicate(ctx context.Context, nodeID string) error {
	panic("implement me")
}

func (c *single) ClusterResetSoft(ctx context.Context) error {
	panic("implement me")
}

func (c *single) ClusterResetHard(ctx context.Context) error {
	panic("implement me")
}

func (c *single) ClusterInfo(ctx context.Context) string {
	panic("implement me")
}

func (c *single) ClusterKeySlot(ctx context.Context, key string) int {
	panic("implement me")
}

func (c *single) ClusterGetKeysInSlot(ctx context.Context, slot int, count int) []string {
	panic("implement me")
}

func (c *single) ClusterCountFailureReports(ctx context.Context, nodeID string) int {
	panic("implement me")
}

func (c *single) ClusterCountKeysInSlot(ctx context.Context, slot int) int {
	panic("implement me")
}

func (c *single) ClusterDelSlots(ctx context.Context, slots ...int) error {
	panic("implement me")
}

func (c *single) ClusterDelSlotsRange(ctx context.Context, min, max int) error {
	panic("implement me")
}

func (c *single) ClusterSaveConfig(ctx context.Context) error {
	panic("implement me")
}

func (c *single) ClusterSlaves(ctx context.Context, nodeID string) []string {
	panic("implement me")
}

func (c *single) ClusterFailover(ctx context.Context) error {
	panic("implement me")
}

func (c *single) ClusterAddSlots(ctx context.Context, slots ...int) error {
	panic("implement me")
}

func (c *single) ClusterAddSlotsRange(ctx context.Context, min, max int) error {
	panic("implement me")
}

func (c *single) GeoAdd(ctx context.Context, key string, geoLocation ...*redis.GeoLocation) int {
	panic("implement me")
}

func (c *single) GeoPos(ctx context.Context, key string, members ...string) []*redis.GeoPos {
	panic("implement me")
}

func (c *single) GeoRadius(ctx context.Context, key string, longitude, latitude float64, query *redis.GeoRadiusQuery) *redis.GeoLocationCmd {
	panic("implement me")
}

func (c *single) GeoRadiusRO(ctx context.Context, key string, longitude, latitude float64, query *redis.GeoRadiusQuery) *redis.GeoLocationCmd {
	panic("implement me")
}

func (c *single) GeoRadiusByMember(ctx context.Context, key, member string, query *redis.GeoRadiusQuery) *redis.GeoLocationCmd {
	panic("implement me")
}

func (c *single) GeoRadiusByMemberRO(ctx context.Context, key, member string, query *redis.GeoRadiusQuery) *redis.GeoLocationCmd {
	panic("implement me")
}

func (c *single) GeoDist(ctx context.Context, key string, member1, member2, unit string) float64 {
	panic("implement me")
}

func (c *single) GeoHash(ctx context.Context, key string, members ...string) []string {
	panic("implement me")
}

func (c *single) ReadOnly(ctx context.Context) error {
	panic("implement me")
}

func (c *single) ReadWrite(ctx context.Context) error {
	panic("implement me")
}

func (c *single) MemoryUsage(ctx context.Context, key string, samples ...int) int {
	panic("implement me")
}

func (c *single) Ping(_ context.Context) error {
	return c.client.Ping().Err()
}

func (c *single) LPush(_ context.Context, key string, val interface{}) (err error) {
	return c.client.LPush(key, val).Err()
}

func (c *single) RPop(_ context.Context, key string) (res string, err error) {
	return c.client.RPop(key).Result()
}

func (c *single) LLen(_ context.Context, key string) int64 {
	return c.client.LLen(key).Val()
}

func (c *single) TypeOf(_ context.Context, key string) (res string, err error) {
	return c.client.Type(key).Result()
}

func (c *single) Keys(_ context.Context, pattern string) (res []string, err error) {
	return c.client.Keys(pattern).Result()
}

func (c *single) ZAdd(_ context.Context, k string, score float64, member interface{}) (err error) {
	return c.client.ZAdd(c.setPrefix(k), redis.Z{
		Score:  score,
		Member: member,
	}).Err()
}

func (c *single) ZCard(_ context.Context, k string) (res int64, err error) {
	return c.client.ZCard(c.setPrefix(k)).Result()
}

func (c *single) ZRangeWithScores(_ context.Context, k string, start, stop int64) (res []redis.Z, err error) {
	return c.client.ZRangeWithScores(c.setPrefix(k), start, stop).Result()
}

func (c *single) HLen(_ context.Context, k string) (res int64, err error) {
	return c.client.HLen(c.setPrefix(k)).Result()
}

func (c *single) HGetAll(_ context.Context, k string) (res map[string]string, err error) {
	return c.client.HGetAll(c.setPrefix(k)).Result()
}

func (c *single) Exists(_ context.Context, keys ...string) int64 {
	return c.client.Exists(keys...).Val()
}

func (c *single) TTL(_ context.Context, key string) time.Duration {
	return c.client.TTL(key).Val()
}

func NewRedisSingle(host, password, prefix string, db int) RedisClient {
	client := redis.NewClient(&redis.Options{
		Addr:     host,
		Password: password, // no password set
		DB:       db,       // use default DB
	})

	return &single{client: client, prefix: prefix}
}

func (c *single) Set(_ context.Context, k string, v interface{}, expir ...time.Duration) (err error) {
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

func (c *single) Get(_ context.Context, k string) (v string, err error) {
	return c.client.Get(c.setPrefix(k)).Result()
}

func (c *single) Del(_ context.Context, k string) (err error) {
	return c.client.Del(c.setPrefix(k)).Err()
}

func (c *single) HSet(_ context.Context, k string, field string, v interface{}) (err error) {
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

func (c *single) HGet(_ context.Context, k string, field string) (res string, err error) {
	return c.client.HGet(c.setPrefix(k), field).Result()
}

func (c *single) HDelAll(_ context.Context, k string) (err error) {
	res, err := c.client.HKeys(c.setPrefix(k)).Result()
	if err != nil {
		return
	}
	return c.client.HDel(c.setPrefix(k), res...).Err()
}

func (c *single) HDel(_ context.Context, k string, field string) (err error) {
	return c.client.HDel(c.setPrefix(k), field).Err()
}

func (c *single) setPrefix(s string) string {
	return c.prefix + s
}

func (c *single) Close(_ context.Context) error {
	return c.client.Close()
}

func (c *single) Subscribe(_ context.Context, channels ...string) *redis.PubSub {
	return c.client.Subscribe(channels...)
}

func (c *single) Publish(_ context.Context, channel string, message interface{}) error {
	return c.client.Publish(channel, message).Err()
}

func (c *single) Incr(_ context.Context, key string, expiration time.Duration) error {
	defer func() {
		c.client.Expire(c.setPrefix(key), expiration)
	}()
	return c.client.Incr(c.setPrefix(key)).Err()
}

func (c *single) SetPrefix(_ context.Context, prefix string) RedisClient {
	c.prefix = prefix
	return c
}
