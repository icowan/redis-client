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
	panic("implement me")
}

func (c *cluster) Pipelined(ctx context.Context, fn func(pipeliner redis.Pipeliner) error) ([]redis.Cmder, error) {
	panic("implement me")
}

func (c *cluster) TxPipelined(ctx context.Context, fn func(pipeliner redis.Pipeliner) error) ([]redis.Cmder, error) {
	panic("implement me")
}

func (c *cluster) TxPipeline(ctx context.Context) redis.Pipeliner {
	panic("implement me")
}

func (c *cluster) Command(ctx context.Context) *redis.CommandsInfoCmd {
	panic("implement me")
}

func (c *cluster) ClientGetName(ctx context.Context) string {
	panic("implement me")
}

func (c *cluster) Echo(ctx context.Context, message interface{}) string {
	panic("implement me")
}

func (c *cluster) Quit(ctx context.Context) error {
	panic("implement me")
}

func (c *cluster) Unlink(ctx context.Context, keys ...string) int {
	panic("implement me")
}

func (c *cluster) Dump(ctx context.Context, key string) string {
	panic("implement me")
}

func (c *cluster) Expire(ctx context.Context, key string, expiration time.Duration) bool {
	panic("implement me")
}

func (c *cluster) ExpireAt(ctx context.Context, key string, tm time.Time) bool {
	panic("implement me")
}

func (c *cluster) Migrate(ctx context.Context, host, port, key string, db int64, timeout time.Duration) error {
	panic("implement me")
}

func (c *cluster) Move(ctx context.Context, key string, db int64) bool {
	panic("implement me")
}

func (c *cluster) ObjectRefCount(ctx context.Context, key string) int {
	panic("implement me")
}

func (c *cluster) ObjectEncoding(ctx context.Context, key string) string {
	panic("implement me")
}

func (c *cluster) ObjectIdleTime(ctx context.Context, key string) time.Duration {
	panic("implement me")
}

func (c *cluster) Persist(ctx context.Context, key string) bool {
	panic("implement me")
}

func (c *cluster) PExpire(ctx context.Context, key string, expiration time.Duration) bool {
	panic("implement me")
}

func (c *cluster) PExpireAt(ctx context.Context, key string, tm time.Time) bool {
	panic("implement me")
}

func (c *cluster) PTTL(ctx context.Context, key string) time.Duration {
	panic("implement me")
}

func (c *cluster) RandomKey(ctx context.Context) string {
	panic("implement me")
}

func (c *cluster) Rename(ctx context.Context, key, newkey string) *redis.StatusCmd {
	panic("implement me")
}

func (c *cluster) RenameNX(ctx context.Context, key, newkey string) bool {
	panic("implement me")
}

func (c *cluster) Restore(ctx context.Context, key string, ttl time.Duration, value string) error {
	panic("implement me")
}

func (c *cluster) RestoreReplace(ctx context.Context, key string, ttl time.Duration, value string) error {
	panic("implement me")
}

func (c *cluster) Sort(ctx context.Context, key string, sort *redis.Sort) []string {
	panic("implement me")
}

func (c *cluster) SortStore(ctx context.Context, key, store string, sort *redis.Sort) int {
	panic("implement me")
}

func (c *cluster) SortInterfaces(ctx context.Context, key string, sort *redis.Sort) []interface{} {
	panic("implement me")
}

func (c *cluster) Touch(ctx context.Context, keys ...string) int {
	panic("implement me")
}

func (c *cluster) Type(ctx context.Context, key string) error {
	panic("implement me")
}

func (c *cluster) Scan(ctx context.Context, cursor uint64, match string, count int64) *redis.ScanCmd {
	panic("implement me")
}

func (c *cluster) SScan(ctx context.Context, key string, cursor uint64, match string, count int64) *redis.ScanCmd {
	panic("implement me")
}

func (c *cluster) HScan(ctx context.Context, key string, cursor uint64, match string, count int64) *redis.ScanCmd {
	panic("implement me")
}

func (c *cluster) ZScan(ctx context.Context, key string, cursor uint64, match string, count int64) *redis.ScanCmd {
	panic("implement me")
}

func (c *cluster) Append(ctx context.Context, key, value string) int {
	panic("implement me")
}

func (c *cluster) BitCount(ctx context.Context, key string, bitCount *redis.BitCount) int {
	panic("implement me")
}

func (c *cluster) BitOpAnd(ctx context.Context, destKey string, keys ...string) int {
	panic("implement me")
}

func (c *cluster) BitOpOr(ctx context.Context, destKey string, keys ...string) int {
	panic("implement me")
}

func (c *cluster) BitOpXor(ctx context.Context, destKey string, keys ...string) int {
	panic("implement me")
}

func (c *cluster) BitOpNot(ctx context.Context, destKey string, key string) int {
	panic("implement me")
}

func (c *cluster) BitPos(ctx context.Context, key string, bit int64, pos ...int64) int {
	panic("implement me")
}

func (c *cluster) Decr(ctx context.Context, key string) int {
	panic("implement me")
}

func (c *cluster) DecrBy(ctx context.Context, key string, decrement int64) int {
	panic("implement me")
}

func (c *cluster) GetBit(ctx context.Context, key string, offset int64) int {
	panic("implement me")
}

func (c *cluster) GetRange(ctx context.Context, key string, start, end int64) string {
	panic("implement me")
}

func (c *cluster) GetSet(ctx context.Context, key string, value interface{}) string {
	panic("implement me")
}

func (c *cluster) IncrBy(ctx context.Context, key string, value int64) int {
	panic("implement me")
}

func (c *cluster) IncrByFloat(ctx context.Context, key string, value float64) float64 {
	panic("implement me")
}

func (c *cluster) MGet(ctx context.Context, keys ...string) []interface{} {
	panic("implement me")
}

func (c *cluster) MSet(ctx context.Context, pairs ...interface{}) error {
	panic("implement me")
}

func (c *cluster) MSetNX(ctx context.Context, pairs ...interface{}) bool {
	panic("implement me")
}

func (c *cluster) SetBit(ctx context.Context, key string, offset int64, value int) int {
	panic("implement me")
}

func (c *cluster) SetNX(ctx context.Context, key string, value interface{}, expiration time.Duration) bool {
	panic("implement me")
}

func (c *cluster) SetXX(ctx context.Context, key string, value interface{}, expiration time.Duration) bool {
	panic("implement me")
}

func (c *cluster) SetRange(ctx context.Context, key string, offset int64, value string) int {
	panic("implement me")
}

func (c *cluster) StrLen(ctx context.Context, key string) int {
	panic("implement me")
}

func (c *cluster) HExists(ctx context.Context, key, field string) bool {
	panic("implement me")
}

func (c *cluster) HIncrBy(ctx context.Context, key, field string, incr int64) int {
	panic("implement me")
}

func (c *cluster) HIncrByFloat(ctx context.Context, key, field string, incr float64) float64 {
	panic("implement me")
}

func (c *cluster) HKeys(ctx context.Context, key string) []string {
	panic("implement me")
}

func (c *cluster) HMGet(ctx context.Context, key string, fields ...string) []interface{} {
	panic("implement me")
}

func (c *cluster) HMSet(ctx context.Context, key string, fields map[string]interface{}) error {
	panic("implement me")
}

func (c *cluster) HSetNX(ctx context.Context, key, field string, value interface{}) bool {
	panic("implement me")
}

func (c *cluster) HVals(ctx context.Context, key string) []string {
	panic("implement me")
}

func (c *cluster) BLPop(ctx context.Context, timeout time.Duration, keys ...string) []string {
	panic("implement me")
}

func (c *cluster) BRPop(ctx context.Context, timeout time.Duration, keys ...string) []string {
	panic("implement me")
}

func (c *cluster) BRPopLPush(ctx context.Context, source, destination string, timeout time.Duration) string {
	panic("implement me")
}

func (c *cluster) LIndex(ctx context.Context, key string, index int64) string {
	panic("implement me")
}

func (c *cluster) LInsert(ctx context.Context, key, op string, pivot, value interface{}) int {
	panic("implement me")
}

func (c *cluster) LInsertBefore(ctx context.Context, key string, pivot, value interface{}) int {
	panic("implement me")
}

func (c *cluster) LInsertAfter(ctx context.Context, key string, pivot, value interface{}) int {
	panic("implement me")
}

func (c *cluster) LPop(ctx context.Context, key string) string {
	panic("implement me")
}

func (c *cluster) LPushX(ctx context.Context, key string, value interface{}) int {
	panic("implement me")
}

func (c *cluster) LRange(ctx context.Context, key string, start, stop int64) []string {
	panic("implement me")
}

func (c *cluster) LRem(ctx context.Context, key string, count int64, value interface{}) int {
	panic("implement me")
}

func (c *cluster) LSet(ctx context.Context, key string, index int64, value interface{}) error {
	panic("implement me")
}

func (c *cluster) LTrim(ctx context.Context, key string, start, stop int64) error {
	panic("implement me")
}

func (c *cluster) RPopLPush(ctx context.Context, source, destination string) string {
	panic("implement me")
}

func (c *cluster) RPush(ctx context.Context, key string, values ...interface{}) int {
	panic("implement me")
}

func (c *cluster) RPushX(ctx context.Context, key string, value interface{}) int {
	panic("implement me")
}

func (c *cluster) SAdd(ctx context.Context, key string, members ...interface{}) int {
	panic("implement me")
}

func (c *cluster) SCard(ctx context.Context, key string) int {
	panic("implement me")
}

func (c *cluster) SDiff(ctx context.Context, keys ...string) []string {
	panic("implement me")
}

func (c *cluster) SDiffStore(ctx context.Context, destination string, keys ...string) int {
	panic("implement me")
}

func (c *cluster) SInter(ctx context.Context, keys ...string) []string {
	panic("implement me")
}

func (c *cluster) SInterStore(ctx context.Context, destination string, keys ...string) int {
	panic("implement me")
}

func (c *cluster) SIsMember(ctx context.Context, key string, member interface{}) bool {
	panic("implement me")
}

func (c *cluster) SMembers(ctx context.Context, key string) []string {
	panic("implement me")
}

func (c *cluster) SMembersMap(ctx context.Context, key string) map[string]struct{} {
	panic("implement me")
}

func (c *cluster) SMove(ctx context.Context, source, destination string, member interface{}) bool {
	panic("implement me")
}

func (c *cluster) SPop(ctx context.Context, key string) string {
	panic("implement me")
}

func (c *cluster) SPopN(ctx context.Context, key string, count int64) []string {
	panic("implement me")
}

func (c *cluster) SRandMember(ctx context.Context, key string) string {
	panic("implement me")
}

func (c *cluster) SRandMemberN(ctx context.Context, key string, count int64) []string {
	panic("implement me")
}

func (c *cluster) SRem(ctx context.Context, key string, members ...interface{}) int {
	panic("implement me")
}

func (c *cluster) SUnion(ctx context.Context, keys ...string) []string {
	panic("implement me")
}

func (c *cluster) SUnionStore(ctx context.Context, destination string, keys ...string) int {
	panic("implement me")
}

func (c *cluster) XAdd(ctx context.Context, a *redis.XAddArgs) string {
	panic("implement me")
}

func (c *cluster) XDel(ctx context.Context, stream string, ids ...string) int {
	panic("implement me")
}

func (c *cluster) XLen(ctx context.Context, stream string) int {
	panic("implement me")
}

func (c *cluster) XRange(ctx context.Context, stream, start, stop string) []redis.XMessage {
	panic("implement me")
}

func (c *cluster) XRangeN(ctx context.Context, stream, start, stop string, count int64) []redis.XMessage {
	panic("implement me")
}

func (c *cluster) XRevRange(ctx context.Context, stream string, start, stop string) []redis.XMessage {
	panic("implement me")
}

func (c *cluster) XRevRangeN(ctx context.Context, stream string, start, stop string, count int64) []redis.XMessage {
	panic("implement me")
}

func (c *cluster) XRead(ctx context.Context, a *redis.XReadArgs) []redis.XStream {
	panic("implement me")
}

func (c *cluster) XReadStreams(ctx context.Context, streams ...string) []redis.XStream {
	panic("implement me")
}

func (c *cluster) XGroupCreate(ctx context.Context, stream, group, start string) error {
	panic("implement me")
}

func (c *cluster) XGroupCreateMkStream(ctx context.Context, stream, group, start string) error {
	panic("implement me")
}

func (c *cluster) XGroupSetID(ctx context.Context, stream, group, start string) error {
	panic("implement me")
}

func (c *cluster) XGroupDestroy(ctx context.Context, stream, group string) int {
	panic("implement me")
}

func (c *cluster) XGroupDelConsumer(ctx context.Context, stream, group, consumer string) int {
	panic("implement me")
}

func (c *cluster) XReadGroup(ctx context.Context, a *redis.XReadGroupArgs) []redis.XStream {
	panic("implement me")
}

func (c *cluster) XAck(ctx context.Context, stream, group string, ids ...string) int {
	panic("implement me")
}

func (c *cluster) XPending(stream, group string) *redis.XPending {
	panic("implement me")
}

func (c *cluster) XPendingExt(ctx context.Context, a *redis.XPendingExtArgs) []redis.XPendingExt {
	panic("implement me")
}

func (c *cluster) XClaim(ctx context.Context, a *redis.XClaimArgs) []redis.XMessage {
	panic("implement me")
}

func (c *cluster) XClaimJustID(ctx context.Context, a *redis.XClaimArgs) []string {
	panic("implement me")
}

func (c *cluster) XTrim(ctx context.Context, key string, maxLen int64) int {
	panic("implement me")
}

func (c *cluster) XTrimApprox(ctx context.Context, key string, maxLen int64) int {
	panic("implement me")
}

func (c *cluster) BZPopMax(ctx context.Context, timeout time.Duration, keys ...string) redis.ZWithKey {
	panic("implement me")
}

func (c *cluster) BZPopMin(timeout time.Duration, keys ...string) redis.ZWithKey {
	panic("implement me")
}

func (c *cluster) ZAddNX(ctx context.Context, key string, members ...redis.Z) int {
	panic("implement me")
}

func (c *cluster) ZAddXX(ctx context.Context, key string, members ...redis.Z) int {
	panic("implement me")
}

func (c *cluster) ZAddCh(ctx context.Context, key string, members ...redis.Z) int {
	panic("implement me")
}

func (c *cluster) ZAddNXCh(ctx context.Context, key string, members ...redis.Z) int {
	panic("implement me")
}

func (c *cluster) ZAddXXCh(ctx context.Context, key string, members ...redis.Z) int {
	panic("implement me")
}

func (c *cluster) ZIncr(ctx context.Context, key string, member redis.Z) float64 {
	panic("implement me")
}

func (c *cluster) ZIncrNX(ctx context.Context, key string, member redis.Z) float64 {
	panic("implement me")
}

func (c *cluster) ZIncrXX(ctx context.Context, key string, member redis.Z) float64 {
	panic("implement me")
}

func (c *cluster) ZCount(ctx context.Context, key, min, max string) int {
	panic("implement me")
}

func (c *cluster) ZLexCount(ctx context.Context, key, min, max string) int {
	panic("implement me")
}

func (c *cluster) ZIncrBy(ctx context.Context, key string, increment float64, member string) float64 {
	panic("implement me")
}

func (c *cluster) ZInterStore(ctx context.Context, destination string, store redis.ZStore, keys ...string) int {
	panic("implement me")
}

func (c *cluster) ZPopMax(ctx context.Context, key string, count ...int64) []redis.Z {
	panic("implement me")
}

func (c *cluster) ZPopMin(ctx context.Context, key string, count ...int64) []redis.Z {
	panic("implement me")
}

func (c *cluster) ZRange(ctx context.Context, key string, start, stop int64) []string {
	panic("implement me")
}

func (c *cluster) ZRangeByScore(ctx context.Context, key string, opt redis.ZRangeBy) []string {
	panic("implement me")
}

func (c *cluster) ZRangeByLex(ctx context.Context, key string, opt redis.ZRangeBy) []string {
	panic("implement me")
}

func (c *cluster) ZRangeByScoreWithScores(ctx context.Context, key string, opt redis.ZRangeBy) []redis.Z {
	panic("implement me")
}

func (c *cluster) ZRank(ctx context.Context, key, member string) int {
	panic("implement me")
}

func (c *cluster) ZRem(ctx context.Context, key string, members ...interface{}) int {
	panic("implement me")
}

func (c *cluster) ZRemRangeByRank(ctx context.Context, key string, start, stop int64) int {
	panic("implement me")
}

func (c *cluster) ZRemRangeByScore(ctx context.Context, key, min, max string) int {
	panic("implement me")
}

func (c *cluster) ZRemRangeByLex(ctx context.Context, key, min, max string) int {
	panic("implement me")
}

func (c *cluster) ZRevRange(ctx context.Context, key string, start, stop int64) []string {
	panic("implement me")
}

func (c *cluster) ZRevRangeWithScores(ctx context.Context, key string, start, stop int64) []redis.Z {
	panic("implement me")
}

func (c *cluster) ZRevRangeByScore(ctx context.Context, key string, opt redis.ZRangeBy) []string {
	panic("implement me")
}

func (c *cluster) ZRevRangeByLex(ctx context.Context, key string, opt redis.ZRangeBy) []string {
	panic("implement me")
}

func (c *cluster) ZRevRangeByScoreWithScores(ctx context.Context, key string, opt redis.ZRangeBy) []redis.Z {
	panic("implement me")
}

func (c *cluster) ZRevRank(ctx context.Context, key, member string) int {
	panic("implement me")
}

func (c *cluster) ZScore(ctx context.Context, key, member string) float64 {
	panic("implement me")
}

func (c *cluster) ZUnionStore(ctx context.Context, dest string, store redis.ZStore, keys ...string) int {
	panic("implement me")
}

func (c *cluster) PFAdd(ctx context.Context, key string, els ...interface{}) int {
	panic("implement me")
}

func (c *cluster) PFCount(ctx context.Context, keys ...string) int {
	panic("implement me")
}

func (c *cluster) PFMerge(ctx context.Context, dest string, keys ...string) error {
	panic("implement me")
}

func (c *cluster) BgRewriteAOF(ctx context.Context) error {
	panic("implement me")
}

func (c *cluster) BgSave(ctx context.Context) error {
	panic("implement me")
}

func (c *cluster) ClientKill(ctx context.Context, ipPort string) error {
	panic("implement me")
}

func (c *cluster) ClientKillByFilter(ctx context.Context, keys ...string) int {
	panic("implement me")
}

func (c *cluster) ClientList(ctx context.Context) string {
	panic("implement me")
}

func (c *cluster) ClientPause(ctx context.Context, dur time.Duration) bool {
	panic("implement me")
}

func (c *cluster) ClientID(ctx context.Context) int {
	panic("implement me")
}

func (c *cluster) ConfigGet(ctx context.Context, parameter string) []interface{} {
	panic("implement me")
}

func (c *cluster) ConfigResetStat(ctx context.Context) error {
	panic("implement me")
}

func (c *cluster) ConfigSet(ctx context.Context, parameter, value string) error {
	panic("implement me")
}

func (c *cluster) ConfigRewrite(ctx context.Context) error {
	panic("implement me")
}

func (c *cluster) DBSize(ctx context.Context) int {
	panic("implement me")
}

func (c *cluster) FlushAll(ctx context.Context) error {
	panic("implement me")
}

func (c *cluster) FlushAllAsync(ctx context.Context) error {
	panic("implement me")
}

func (c *cluster) FlushDB(ctx context.Context) error {
	panic("implement me")
}

func (c *cluster) FlushDBAsync(ctx context.Context) error {
	panic("implement me")
}

func (c *cluster) Info(ctx context.Context, section ...string) string {
	panic("implement me")
}

func (c *cluster) LastSave(ctx context.Context) int {
	panic("implement me")
}

func (c *cluster) Save(ctx context.Context) error {
	panic("implement me")
}

func (c *cluster) Shutdown(ctx context.Context) error {
	panic("implement me")
}

func (c *cluster) ShutdownSave(ctx context.Context) error {
	panic("implement me")
}

func (c *cluster) ShutdownNoSave(ctx context.Context) error {
	panic("implement me")
}

func (c *cluster) SlaveOf(ctx context.Context, host, port string) error {
	panic("implement me")
}

func (c *cluster) Time(ctx context.Context) time.Time {
	panic("implement me")
}

func (c *cluster) Eval(ctx context.Context, script string, keys []string, args ...interface{}) *redis.Cmd {
	panic("implement me")
}

func (c *cluster) EvalSha(ctx context.Context, sha1 string, keys []string, args ...interface{}) *redis.Cmd {
	panic("implement me")
}

func (c *cluster) ScriptExists(ctx context.Context, hashes ...string) []bool {
	panic("implement me")
}

func (c *cluster) ScriptFlush(ctx context.Context) error {
	panic("implement me")
}

func (c *cluster) ScriptKill(ctx context.Context) error {
	panic("implement me")
}

func (c *cluster) ScriptLoad(ctx context.Context, script string) string {
	panic("implement me")
}

func (c *cluster) DebugObject(ctx context.Context, key string) string {
	panic("implement me")
}

func (c *cluster) PubSubChannels(ctx context.Context, pattern string) []string {
	panic("implement me")
}

func (c *cluster) PubSubNumSub(ctx context.Context, channels ...string) map[string]int64 {
	panic("implement me")
}

func (c *cluster) PubSubNumPat(ctx context.Context) int {
	panic("implement me")
}

func (c *cluster) ClusterSlots(ctx context.Context) []redis.ClusterSlot {
	panic("implement me")
}

func (c *cluster) ClusterNodes(ctx context.Context) string {
	panic("implement me")
}

func (c *cluster) ClusterMeet(ctx context.Context, host, port string) error {
	panic("implement me")
}

func (c *cluster) ClusterForget(ctx context.Context, nodeID string) error {
	panic("implement me")
}

func (c *cluster) ClusterReplicate(ctx context.Context, nodeID string) error {
	panic("implement me")
}

func (c *cluster) ClusterResetSoft(ctx context.Context) error {
	panic("implement me")
}

func (c *cluster) ClusterResetHard(ctx context.Context) error {
	panic("implement me")
}

func (c *cluster) ClusterInfo(ctx context.Context) string {
	panic("implement me")
}

func (c *cluster) ClusterKeySlot(ctx context.Context, key string) int {
	panic("implement me")
}

func (c *cluster) ClusterGetKeysInSlot(ctx context.Context, slot int, count int) []string {
	panic("implement me")
}

func (c *cluster) ClusterCountFailureReports(ctx context.Context, nodeID string) int {
	panic("implement me")
}

func (c *cluster) ClusterCountKeysInSlot(ctx context.Context, slot int) int {
	panic("implement me")
}

func (c *cluster) ClusterDelSlots(ctx context.Context, slots ...int) error {
	panic("implement me")
}

func (c *cluster) ClusterDelSlotsRange(ctx context.Context, min, max int) error {
	panic("implement me")
}

func (c *cluster) ClusterSaveConfig(ctx context.Context) error {
	panic("implement me")
}

func (c *cluster) ClusterSlaves(ctx context.Context, nodeID string) []string {
	panic("implement me")
}

func (c *cluster) ClusterFailover(ctx context.Context) error {
	panic("implement me")
}

func (c *cluster) ClusterAddSlots(ctx context.Context, slots ...int) error {
	panic("implement me")
}

func (c *cluster) ClusterAddSlotsRange(ctx context.Context, min, max int) error {
	panic("implement me")
}

func (c *cluster) GeoAdd(ctx context.Context, key string, geoLocation ...*redis.GeoLocation) int {
	panic("implement me")
}

func (c *cluster) GeoPos(ctx context.Context, key string, members ...string) []*redis.GeoPos {
	panic("implement me")
}

func (c *cluster) GeoRadius(ctx context.Context, key string, longitude, latitude float64, query *redis.GeoRadiusQuery) *redis.GeoLocationCmd {
	panic("implement me")
}

func (c *cluster) GeoRadiusRO(ctx context.Context, key string, longitude, latitude float64, query *redis.GeoRadiusQuery) *redis.GeoLocationCmd {
	panic("implement me")
}

func (c *cluster) GeoRadiusByMember(ctx context.Context, key, member string, query *redis.GeoRadiusQuery) *redis.GeoLocationCmd {
	panic("implement me")
}

func (c *cluster) GeoRadiusByMemberRO(ctx context.Context, key, member string, query *redis.GeoRadiusQuery) *redis.GeoLocationCmd {
	panic("implement me")
}

func (c *cluster) GeoDist(ctx context.Context, key string, member1, member2, unit string) float64 {
	panic("implement me")
}

func (c *cluster) GeoHash(ctx context.Context, key string, members ...string) []string {
	panic("implement me")
}

func (c *cluster) ReadOnly(ctx context.Context) error {
	panic("implement me")
}

func (c *cluster) ReadWrite(ctx context.Context) error {
	panic("implement me")
}

func (c *cluster) MemoryUsage(ctx context.Context, key string, samples ...int) int {
	panic("implement me")
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
