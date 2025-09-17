package redis

import (
	"context"
	"log"
	"sync"
	"time"

	redis "github.com/redis/go-redis/v9"
)

var (
	client *Redis
	once   sync.Once
)

type Redis struct {
	client redis.UniversalClient
	conf   Conf
}

func NewRedis(conf Conf) (*Redis, error) {
	client := redis.NewUniversalClient(&redis.UniversalOptions{
		Addrs:    conf.Addrs,
		Password: conf.Password,
		DB:       conf.DB,
	})
	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		panic(err)
	}
	log.Println("connected to redis:", client)
	return &Redis{
		client: client,
		conf:   conf,
	}, nil
}

func RedisSingleton(conf Conf) (*Redis, error) {
	var err error
	once.Do(func() {
		client, err = NewRedis(conf)
	})
	return client, err
}

// ACLLog implements redis.Cmdable.
func (r *Redis) ACLLog(ctx context.Context, count int64) *redis.ACLLogCmd {
	return r.client.ACLLog(ctx, count)
}

// ACLLogReset implements redis.Cmdable.
func (r *Redis) ACLLogReset(ctx context.Context) *redis.StatusCmd {
	return r.client.ACLLogReset(ctx)
}

// BFAdd implements redis.Cmdable.
func (r *Redis) BFAdd(ctx context.Context, key string, element interface{}) *redis.BoolCmd {
	return r.client.BFAdd(ctx, key, element)
}

// BFCard implements redis.Cmdable.
func (r *Redis) BFCard(ctx context.Context, key string) *redis.IntCmd {
	return r.client.BFCard(ctx, key)
}

// BFExists implements redis.Cmdable.
func (r *Redis) BFExists(ctx context.Context, key string, element interface{}) *redis.BoolCmd {
	return r.client.BFExists(ctx, key, element)
}

// BFInfo implements redis.Cmdable.
func (r *Redis) BFInfo(ctx context.Context, key string) *redis.BFInfoCmd {
	return r.client.BFInfo(ctx, key)
}

// BFInfoArg implements redis.Cmdable.
func (r *Redis) BFInfoArg(ctx context.Context, key string, option string) *redis.BFInfoCmd {
	return r.client.BFInfoArg(ctx, key, option)
}

// BFInfoCapacity implements redis.Cmdable.
func (r *Redis) BFInfoCapacity(ctx context.Context, key string) *redis.BFInfoCmd {
	return r.client.BFInfoCapacity(ctx, key)
}

// BFInfoExpansion implements redis.Cmdable.
func (r *Redis) BFInfoExpansion(ctx context.Context, key string) *redis.BFInfoCmd {
	return r.client.BFInfoExpansion(ctx, key)
}

// BFInfoFilters implements redis.Cmdable.
func (r *Redis) BFInfoFilters(ctx context.Context, key string) *redis.BFInfoCmd {
	return r.client.BFInfoFilters(ctx, key)
}

// BFInfoItems implements redis.Cmdable.
func (r *Redis) BFInfoItems(ctx context.Context, key string) *redis.BFInfoCmd {
	return r.client.BFInfoItems(ctx, key)
}

// BFInfoSize implements redis.Cmdable.
func (r *Redis) BFInfoSize(ctx context.Context, key string) *redis.BFInfoCmd {
	return r.client.BFInfoSize(ctx, key)
}

// BFInsert implements redis.Cmdable.
func (r *Redis) BFInsert(ctx context.Context, key string, options *redis.BFInsertOptions, elements ...interface{}) *redis.BoolSliceCmd {
	return r.client.BFInsert(ctx, key, options, elements...)
}

// BFLoadChunk implements redis.Cmdable.
func (r *Redis) BFLoadChunk(ctx context.Context, key string, iterator int64, data interface{}) *redis.StatusCmd {
	return r.client.BFLoadChunk(ctx, key, iterator, data)
}

// BFMAdd implements redis.Cmdable.
func (r *Redis) BFMAdd(ctx context.Context, key string, elements ...interface{}) *redis.BoolSliceCmd {
	return r.client.BFMAdd(ctx, key, elements...)
}

// BFMExists implements redis.Cmdable.
func (r *Redis) BFMExists(ctx context.Context, key string, elements ...interface{}) *redis.BoolSliceCmd {
	return r.client.BFMExists(ctx, key, elements...)
}

// BFReserve implements redis.Cmdable.
func (r *Redis) BFReserve(ctx context.Context, key string, errorRate float64, capacity int64) *redis.StatusCmd {
	return r.client.BFReserve(ctx, key, errorRate, capacity)
}

// BFReserveExpansion implements redis.Cmdable.
func (r *Redis) BFReserveExpansion(ctx context.Context, key string, errorRate float64, capacity int64, expansion int64) *redis.StatusCmd {
	return r.client.BFReserveExpansion(ctx, key, errorRate, capacity, expansion)
}

// BFReserveNonScaling implements redis.Cmdable.
func (r *Redis) BFReserveNonScaling(ctx context.Context, key string, errorRate float64, capacity int64) *redis.StatusCmd {
	return r.client.BFReserveNonScaling(ctx, key, errorRate, capacity)
}

// BFScanDump implements redis.Cmdable.
func (r *Redis) BFScanDump(ctx context.Context, key string, iterator int64) *redis.ScanDumpCmd {
	return r.client.BFScanDump(ctx, key, iterator)
}

// CFAdd implements redis.Cmdable.
func (r *Redis) CFAdd(ctx context.Context, key string, element interface{}) *redis.BoolCmd {
	return r.client.CFAdd(ctx, key, element)
}

// CFAddNX implements redis.Cmdable.
func (r *Redis) CFAddNX(ctx context.Context, key string, element interface{}) *redis.BoolCmd {
	return r.client.CFAddNX(ctx, key, element)
}

// CFCount implements redis.Cmdable.
func (r *Redis) CFCount(ctx context.Context, key string, element interface{}) *redis.IntCmd {
	return r.client.CFCount(ctx, key, element)
}

// CFDel implements redis.Cmdable.
func (r *Redis) CFDel(ctx context.Context, key string, element interface{}) *redis.BoolCmd {
	return r.client.CFDel(ctx, key, element)
}

// CFExists implements redis.Cmdable.
func (r *Redis) CFExists(ctx context.Context, key string, element interface{}) *redis.BoolCmd {
	return r.client.CFExists(ctx, key, element)
}

// CFInfo implements redis.Cmdable.
func (r *Redis) CFInfo(ctx context.Context, key string) *redis.CFInfoCmd {
	return r.client.CFInfo(ctx, key)
}

// CFInsert implements redis.Cmdable.
func (r *Redis) CFInsert(ctx context.Context, key string, options *redis.CFInsertOptions, elements ...interface{}) *redis.BoolSliceCmd {
	return r.client.CFInsert(ctx, key, options, elements...)
}

// CFInsertNX implements redis.Cmdable.
func (r *Redis) CFInsertNX(ctx context.Context, key string, options *redis.CFInsertOptions, elements ...interface{}) *redis.IntSliceCmd {
	return r.client.CFInsertNX(ctx, key, options, elements...)
}

// CFLoadChunk implements redis.Cmdable.
func (r *Redis) CFLoadChunk(ctx context.Context, key string, iterator int64, data interface{}) *redis.StatusCmd {
	return r.client.CFLoadChunk(ctx, key, iterator, data)
}

// CFMExists implements redis.Cmdable.
func (r *Redis) CFMExists(ctx context.Context, key string, elements ...interface{}) *redis.BoolSliceCmd {
	return r.client.CFMExists(ctx, key, elements...)
}

// CFReserve implements redis.Cmdable.
func (r *Redis) CFReserve(ctx context.Context, key string, capacity int64) *redis.StatusCmd {
	return r.client.CFReserve(ctx, key, capacity)
}

// CFReserveBucketSize implements redis.Cmdable.
func (r *Redis) CFReserveBucketSize(ctx context.Context, key string, capacity int64, bucketsize int64) *redis.StatusCmd {
	return r.client.CFReserveBucketSize(ctx, key, capacity, bucketsize)
}

// CFReserveExpansion implements redis.Cmdable.
func (r *Redis) CFReserveExpansion(ctx context.Context, key string, capacity int64, expansion int64) *redis.StatusCmd {
	return r.client.CFReserveExpansion(ctx, key, capacity, expansion)
}

// CFReserveMaxIterations implements redis.Cmdable.
func (r *Redis) CFReserveMaxIterations(ctx context.Context, key string, capacity int64, maxiterations int64) *redis.StatusCmd {
	return r.client.CFReserveMaxIterations(ctx, key, capacity, maxiterations)
}

// CFScanDump implements redis.Cmdable.
func (r *Redis) CFScanDump(ctx context.Context, key string, iterator int64) *redis.ScanDumpCmd {
	return r.client.CFScanDump(ctx, key, iterator)
}

// CMSIncrBy implements redis.Cmdable.
func (r *Redis) CMSIncrBy(ctx context.Context, key string, elements ...interface{}) *redis.IntSliceCmd {
	return r.client.CMSIncrBy(ctx, key, elements...)
}

// CMSInfo implements redis.Cmdable.
func (r *Redis) CMSInfo(ctx context.Context, key string) *redis.CMSInfoCmd {
	return r.client.CMSInfo(ctx, key)
}

// CMSInitByDim implements redis.Cmdable.
func (r *Redis) CMSInitByDim(ctx context.Context, key string, width int64, height int64) *redis.StatusCmd {
	return r.client.CMSInitByDim(ctx, key, width, height)
}

// CMSInitByProb implements redis.Cmdable.
func (r *Redis) CMSInitByProb(ctx context.Context, key string, errorRate float64, probability float64) *redis.StatusCmd {
	return r.client.CMSInitByProb(ctx, key, errorRate, probability)
}

// CMSMerge implements redis.Cmdable.
func (r *Redis) CMSMerge(ctx context.Context, destKey string, sourceKeys ...string) *redis.StatusCmd {
	return r.client.CMSMerge(ctx, destKey, sourceKeys...)
}

// CMSMergeWithWeight implements redis.Cmdable.
func (r *Redis) CMSMergeWithWeight(ctx context.Context, destKey string, sourceKeys map[string]int64) *redis.StatusCmd {
	return r.client.CMSMergeWithWeight(ctx, destKey, sourceKeys)
}

// CMSQuery implements redis.Cmdable.
func (r *Redis) CMSQuery(ctx context.Context, key string, elements ...interface{}) *redis.IntSliceCmd {
	return r.client.CMSQuery(ctx, key, elements...)
}

// TDigestAdd implements redis.Cmdable.
func (r *Redis) TDigestAdd(ctx context.Context, key string, elements ...float64) *redis.StatusCmd {
	return r.client.TDigestAdd(ctx, key, elements...)
}

// TDigestByRank implements redis.Cmdable.
func (r *Redis) TDigestByRank(ctx context.Context, key string, rank ...uint64) *redis.FloatSliceCmd {
	return r.client.TDigestByRank(ctx, key, rank...)
}

// TDigestByRevRank implements redis.Cmdable.
func (r *Redis) TDigestByRevRank(ctx context.Context, key string, rank ...uint64) *redis.FloatSliceCmd {
	return r.client.TDigestByRevRank(ctx, key, rank...)
}

// TDigestCDF implements redis.Cmdable.
func (r *Redis) TDigestCDF(ctx context.Context, key string, elements ...float64) *redis.FloatSliceCmd {
	return r.client.TDigestCDF(ctx, key, elements...)
}

// TDigestCreate implements redis.Cmdable.
func (r *Redis) TDigestCreate(ctx context.Context, key string) *redis.StatusCmd {
	return r.client.TDigestCreate(ctx, key)
}

// TDigestCreateWithCompression implements redis.Cmdable.
func (r *Redis) TDigestCreateWithCompression(ctx context.Context, key string, compression int64) *redis.StatusCmd {
	return r.client.TDigestCreateWithCompression(ctx, key, compression)
}

// TDigestInfo implements redis.Cmdable.
func (r *Redis) TDigestInfo(ctx context.Context, key string) *redis.TDigestInfoCmd {
	return r.client.TDigestInfo(ctx, key)
}

// TDigestMax implements redis.Cmdable.
func (r *Redis) TDigestMax(ctx context.Context, key string) *redis.FloatCmd {
	return r.client.TDigestMax(ctx, key)
}

// TDigestMerge implements redis.Cmdable.
func (r *Redis) TDigestMerge(ctx context.Context, destKey string, options *redis.TDigestMergeOptions, sourceKeys ...string) *redis.StatusCmd {
	return r.client.TDigestMerge(ctx, destKey, options, sourceKeys...)
}

// TDigestMin implements redis.Cmdable.
func (r *Redis) TDigestMin(ctx context.Context, key string) *redis.FloatCmd {
	return r.client.TDigestMin(ctx, key)
}

// TDigestQuantile implements redis.Cmdable.
func (r *Redis) TDigestQuantile(ctx context.Context, key string, elements ...float64) *redis.FloatSliceCmd {
	return r.client.TDigestQuantile(ctx, key, elements...)
}

// TDigestRank implements redis.Cmdable.
func (r *Redis) TDigestRank(ctx context.Context, key string, values ...float64) *redis.IntSliceCmd {
	return r.client.TDigestRank(ctx, key, values...)
}

// TDigestReset implements redis.Cmdable.
func (r *Redis) TDigestReset(ctx context.Context, key string) *redis.StatusCmd {
	return r.client.TDigestReset(ctx, key)
}

// TDigestRevRank implements redis.Cmdable.
func (r *Redis) TDigestRevRank(ctx context.Context, key string, values ...float64) *redis.IntSliceCmd {
	return r.client.TDigestRevRank(ctx, key, values...)
}

// TDigestTrimmedMean implements redis.Cmdable.
func (r *Redis) TDigestTrimmedMean(ctx context.Context, key string, lowCutQuantile float64, highCutQuantile float64) *redis.FloatCmd {
	return r.client.TDigestTrimmedMean(ctx, key, lowCutQuantile, highCutQuantile)
}

// TopKAdd implements redis.Cmdable.
func (r *Redis) TopKAdd(ctx context.Context, key string, elements ...interface{}) *redis.StringSliceCmd {
	return r.client.TopKAdd(ctx, key, elements...)
}

// TopKCount implements redis.Cmdable.
func (r *Redis) TopKCount(ctx context.Context, key string, elements ...interface{}) *redis.IntSliceCmd {
	return r.client.TopKCount(ctx, key, elements...)
}

// TopKIncrBy implements redis.Cmdable.
func (r *Redis) TopKIncrBy(ctx context.Context, key string, elements ...interface{}) *redis.StringSliceCmd {
	return r.client.TopKIncrBy(ctx, key, elements...)
}

// TopKInfo implements redis.Cmdable.
func (r *Redis) TopKInfo(ctx context.Context, key string) *redis.TopKInfoCmd {
	return r.client.TopKInfo(ctx, key)
}

// TopKList implements redis.Cmdable.
func (r *Redis) TopKList(ctx context.Context, key string) *redis.StringSliceCmd {
	return r.client.TopKList(ctx, key)
}

// TopKListWithCount implements redis.Cmdable.
func (r *Redis) TopKListWithCount(ctx context.Context, key string) *redis.MapStringIntCmd {
	return r.client.TopKListWithCount(ctx, key)
}

// TopKQuery implements redis.Cmdable.
func (r *Redis) TopKQuery(ctx context.Context, key string, elements ...interface{}) *redis.BoolSliceCmd {
	return r.client.TopKQuery(ctx, key, elements...)
}

// TopKReserve implements redis.Cmdable.
func (r *Redis) TopKReserve(ctx context.Context, key string, k int64) *redis.StatusCmd {
	return r.client.TopKReserve(ctx, key, k)
}

// TopKReserveWithOptions implements redis.Cmdable.
func (r *Redis) TopKReserveWithOptions(ctx context.Context, key string, k int64, width int64, depth int64, decay float64) *redis.StatusCmd {
	return r.client.TopKReserveWithOptions(ctx, key, k, width, depth, decay)
}

var _ redis.Cmdable = &Redis{}

// ACLDryRun implements redis.Cmdable
func (r *Redis) ACLDryRun(ctx context.Context, username string, command ...interface{}) *redis.StringCmd {
	return r.client.ACLDryRun(ctx, username, command...)
}

// BLMPop implements redis.Cmdable
func (r *Redis) BLMPop(ctx context.Context, timeout time.Duration, direction string, count int64, keys ...string) *redis.KeyValuesCmd {
	return r.client.BLMPop(ctx, timeout, direction, count, keys...)
}

// BZMPop implements redis.Cmdable
func (r *Redis) BZMPop(ctx context.Context, timeout time.Duration, order string, count int64, keys ...string) *redis.ZSliceWithKeyCmd {
	return r.client.BZMPop(ctx, timeout, order, count, keys...)
}

// BitPosSpan implements redis.Cmdable
func (r *Redis) BitPosSpan(ctx context.Context, key string, bit int8, start int64, end int64, span string) *redis.IntCmd {
	return r.client.BitPosSpan(ctx, key, bit, start, end, span)
}

// ClientInfo implements redis.Cmdable
func (r *Redis) ClientInfo(ctx context.Context) *redis.ClientInfoCmd {
	return r.client.ClientInfo(ctx)
}

// ClusterLinks implements redis.Cmdable
func (r *Redis) ClusterLinks(ctx context.Context) *redis.ClusterLinksCmd {
	return r.client.ClusterLinks(ctx)
}

// ClusterMyShardID implements redis.Cmdable
func (r *Redis) ClusterMyShardID(ctx context.Context) *redis.StringCmd {
	return r.client.ClusterMyShardID(ctx)
}

// ClusterShards implements redis.Cmdable
func (r *Redis) ClusterShards(ctx context.Context) *redis.ClusterShardsCmd {
	return r.client.ClusterShards(ctx)
}

// CommandGetKeys implements redis.Cmdable
func (r *Redis) CommandGetKeys(ctx context.Context, commands ...interface{}) *redis.StringSliceCmd {
	return r.client.CommandGetKeys(ctx, commands...)
}

// CommandGetKeysAndFlags implements redis.Cmdable
func (r *Redis) CommandGetKeysAndFlags(ctx context.Context, commands ...interface{}) *redis.KeyFlagsCmd {
	return r.client.CommandGetKeysAndFlags(ctx, commands...)
}

// CommandList implements redis.Cmdable
func (r *Redis) CommandList(ctx context.Context, filter *redis.FilterBy) *redis.StringSliceCmd {
	return r.client.CommandList(ctx, filter)
}

// ExpireTime implements redis.Cmdable
func (r *Redis) ExpireTime(ctx context.Context, key string) *redis.DurationCmd {
	return r.client.ExpireTime(ctx, key)
}

// FCall implements redis.Cmdable
func (r *Redis) FCall(ctx context.Context, function string, keys []string, args ...interface{}) *redis.Cmd {
	return r.client.FCall(ctx, function, keys, args...)
}

// FCallRO implements redis.Cmdable
func (r *Redis) FCallRO(ctx context.Context, function string, keys []string, args ...interface{}) *redis.Cmd {
	return r.client.FCallRO(ctx, function, keys, args...)
}

// FCallRo implements redis.Cmdable
func (r *Redis) FCallRo(ctx context.Context, function string, keys []string, args ...interface{}) *redis.Cmd {
	return r.client.FCallRo(ctx, function, keys, args...)
}

// FunctionDelete implements redis.Cmdable
func (r *Redis) FunctionDelete(ctx context.Context, libName string) *redis.StringCmd {
	return r.client.FunctionDelete(ctx, libName)
}

// FunctionDump implements redis.Cmdable
func (r *Redis) FunctionDump(ctx context.Context) *redis.StringCmd {
	return r.client.FunctionDump(ctx)
}

// FunctionFlush implements redis.Cmdable
func (r *Redis) FunctionFlush(ctx context.Context) *redis.StringCmd {
	return r.client.FunctionFlush(ctx)
}

// FunctionFlushAsync implements redis.Cmdable
func (r *Redis) FunctionFlushAsync(ctx context.Context) *redis.StringCmd {
	return r.client.FunctionFlushAsync(ctx)
}

// FunctionKill implements redis.Cmdable
func (r *Redis) FunctionKill(ctx context.Context) *redis.StringCmd {
	return r.client.FunctionKill(ctx)
}

// FunctionList implements redis.Cmdable
func (r *Redis) FunctionList(ctx context.Context, q redis.FunctionListQuery) *redis.FunctionListCmd {
	return r.client.FunctionList(ctx, q)
}

// FunctionLoad implements redis.Cmdable
func (r *Redis) FunctionLoad(ctx context.Context, code string) *redis.StringCmd {
	return r.client.FunctionLoad(ctx, code)
}

// FunctionLoadReplace implements redis.Cmdable
func (r *Redis) FunctionLoadReplace(ctx context.Context, code string) *redis.StringCmd {
	return r.client.FunctionLoadReplace(ctx, code)
}

// FunctionRestore implements redis.Cmdable
func (r *Redis) FunctionRestore(ctx context.Context, libDump string) *redis.StringCmd {
	return r.client.FunctionRestore(ctx, libDump)
}

// FunctionStats implements redis.Cmdable
func (r *Redis) FunctionStats(ctx context.Context) *redis.FunctionStatsCmd {
	return r.client.FunctionStats(ctx)
}

// LCS implements redis.Cmdable
func (r *Redis) LCS(ctx context.Context, q *redis.LCSQuery) *redis.LCSCmd {
	return r.client.LCS(ctx, q)
}

// LMPop implements redis.Cmdable
func (r *Redis) LMPop(ctx context.Context, direction string, count int64, keys ...string) *redis.KeyValuesCmd {
	return r.client.LMPop(ctx, direction, count, keys...)
}

// ModuleLoadex implements redis.Cmdable
func (r *Redis) ModuleLoadex(ctx context.Context, conf *redis.ModuleLoadexConfig) *redis.StringCmd {
	return r.client.ModuleLoadex(ctx, conf)
}

// PExpireTime implements redis.Cmdable
func (r *Redis) PExpireTime(ctx context.Context, key string) *redis.DurationCmd {
	return r.client.PExpireTime(ctx, key)
}

// ZAddGT implements redis.Cmdable
func (r *Redis) ZAddGT(ctx context.Context, key string, members ...redis.Z) *redis.IntCmd {
	return r.client.ZAddGT(ctx, key, members...)
}

// ZAddLT implements redis.Cmdable
func (r *Redis) ZAddLT(ctx context.Context, key string, members ...redis.Z) *redis.IntCmd {
	return r.client.ZAddLT(ctx, key, members...)
}

// ZMPop implements redis.Cmdable
func (r *Redis) ZMPop(ctx context.Context, order string, count int64, keys ...string) *redis.ZSliceWithKeyCmd {
	return r.client.ZMPop(ctx, order, count, keys...)
}

// ZRankWithScore implements redis.Cmdable
func (r *Redis) ZRankWithScore(ctx context.Context, key string, member string) *redis.RankWithScoreCmd {
	return r.client.ZRankWithScore(ctx, key, member)
}

// ZRevRankWithScore implements redis.Cmdable
func (r *Redis) ZRevRankWithScore(ctx context.Context, key string, member string) *redis.RankWithScoreCmd {
	return r.client.ZRevRankWithScore(ctx, key, member)
}

// ClientUnblock implements redis.Cmdable
func (r *Redis) ClientUnblock(ctx context.Context, id int64) *redis.IntCmd {
	return r.client.ClientUnblock(ctx, id)
}

// ClientUnblockWithError implements redis.Cmdable
func (r *Redis) ClientUnblockWithError(ctx context.Context, id int64) *redis.IntCmd {
	return r.client.ClientUnblockWithError(ctx, id)
}

// ClientUnpause implements redis.Cmdable
func (r *Redis) ClientUnpause(ctx context.Context) *redis.BoolCmd {
	return r.client.ClientUnpause(ctx)
}

// ConfigGet implements redis.Cmdable
func (r *Redis) ConfigGet(ctx context.Context, parameter string) *redis.MapStringStringCmd {
	return r.client.ConfigGet(ctx, parameter)
}

// EvalRO implements redis.Cmdable
func (r *Redis) EvalRO(ctx context.Context, script string, keys []string, args ...interface{}) *redis.Cmd {
	return r.client.EvalRO(ctx, script, keys, args...)
}

// EvalShaRO implements redis.Cmdable
func (r *Redis) EvalShaRO(ctx context.Context, sha1 string, keys []string, args ...interface{}) *redis.Cmd {
	return r.client.EvalShaRO(ctx, sha1, keys, args...)
}

// HRandFieldWithValues implements redis.Cmdable
func (r *Redis) HRandFieldWithValues(ctx context.Context, key string, count int) *redis.KeyValueSliceCmd {
	return r.client.HRandFieldWithValues(ctx, key, count)
}

// PubSubShardChannels implements redis.Cmdable
func (r *Redis) PubSubShardChannels(ctx context.Context, pattern string) *redis.StringSliceCmd {
	return r.client.PubSubShardChannels(ctx, pattern)
}

// PubSubShardNumSub implements redis.Cmdable
func (r *Redis) PubSubShardNumSub(ctx context.Context, channels ...string) *redis.MapStringIntCmd {
	return r.client.PubSubShardNumSub(ctx, channels...)
}

// SInterCard implements redis.Cmdable
func (r *Redis) SInterCard(ctx context.Context, limit int64, keys ...string) *redis.IntCmd {
	return r.client.SInterCard(ctx, limit, keys...)
}

// SPublish implements redis.Cmdable
func (r *Redis) SPublish(ctx context.Context, channel string, message interface{}) *redis.IntCmd {
	return r.client.SPublish(ctx, channel, message)
}

// SetEx implements redis.Cmdable
func (r *Redis) SetEx(ctx context.Context, key string, value interface{}, expiration time.Duration) *redis.StatusCmd {
	return r.client.SetEx(ctx, key, value, expiration)
}

// SlowLogGet implements redis.Cmdable
func (r *Redis) SlowLogGet(ctx context.Context, num int64) *redis.SlowLogCmd {
	return r.client.SlowLogGet(ctx, num)
}

// SortRO implements redis.Cmdable
func (r *Redis) SortRO(ctx context.Context, key string, sort *redis.Sort) *redis.StringSliceCmd {
	return r.client.SortRO(ctx, key, sort)
}

// ZInterCard implements redis.Cmdable
func (r *Redis) ZInterCard(ctx context.Context, limit int64, keys ...string) *redis.IntCmd {
	return r.client.ZInterCard(ctx, limit, keys...)
}

// ZRandMemberWithScores implements redis.Cmdable
func (r *Redis) ZRandMemberWithScores(ctx context.Context, key string, count int) *redis.ZSliceCmd {
	return r.client.ZRandMemberWithScores(ctx, key, count)
}

// BLMove implements redis.Cmdable
func (r *Redis) BLMove(ctx context.Context, source string, destination string, srcpos string, destpos string, timeout time.Duration) *redis.StringCmd {
	return r.client.BLMove(ctx, source, destination, source, destpos, timeout)
}

// BitField implements redis.Cmdable
func (r *Redis) BitField(ctx context.Context, key string, args ...interface{}) *redis.IntSliceCmd {
	return r.client.BitField(ctx, key, args...)
}

// Copy implements redis.Cmdable
func (r *Redis) Copy(ctx context.Context, sourceKey string, destKey string, db int, replace bool) *redis.IntCmd {
	return r.client.Copy(ctx, sourceKey, destKey, db, replace)
}

// ExpireGT implements redis.Cmdable
func (r *Redis) ExpireGT(ctx context.Context, key string, expiration time.Duration) *redis.BoolCmd {
	return r.client.ExpireGT(ctx, key, expiration)
}

// ExpireLT implements redis.Cmdable
func (r *Redis) ExpireLT(ctx context.Context, key string, expiration time.Duration) *redis.BoolCmd {
	return r.client.ExpireLT(ctx, key, expiration)
}

// ExpireNX implements redis.Cmdable
func (r *Redis) ExpireNX(ctx context.Context, key string, expiration time.Duration) *redis.BoolCmd {
	return r.client.ExpireNX(ctx, key, expiration)
}

// ExpireXX implements redis.Cmdable
func (r *Redis) ExpireXX(ctx context.Context, key string, expiration time.Duration) *redis.BoolCmd {
	return r.client.ExpireXX(ctx, key, expiration)
}

// GeoRadiusByMemberStore implements redis.Cmdable
func (r *Redis) GeoRadiusByMemberStore(ctx context.Context, key string, member string, query *redis.GeoRadiusQuery) *redis.IntCmd {
	return r.client.GeoRadiusByMemberStore(ctx, key, member, query)
}

// GeoRadiusStore implements redis.Cmdable
func (r *Redis) GeoRadiusStore(ctx context.Context, key string, longitude float64, latitude float64, query *redis.GeoRadiusQuery) *redis.IntCmd {
	return r.client.GeoRadiusStore(ctx, key, latitude, latitude, query)
}

// GeoSearch implements redis.Cmdable
func (r *Redis) GeoSearch(ctx context.Context, key string, q *redis.GeoSearchQuery) *redis.StringSliceCmd {
	return r.client.GeoSearch(ctx, key, q)
}

// GeoSearchLocation implements redis.Cmdable
func (r *Redis) GeoSearchLocation(ctx context.Context, key string, q *redis.GeoSearchLocationQuery) *redis.GeoSearchLocationCmd {
	return r.client.GeoSearchLocation(ctx, key, q)
}

// GeoSearchStore implements redis.Cmdable
func (r *Redis) GeoSearchStore(ctx context.Context, key string, store string, q *redis.GeoSearchStoreQuery) *redis.IntCmd {
	return r.client.GeoSearchStore(ctx, key, store, q)
}

// GetDel implements redis.Cmdable
func (r *Redis) GetDel(ctx context.Context, key string) *redis.StringCmd {
	return r.client.GetDel(ctx, key)
}

// GetEx implements redis.Cmdable
func (r *Redis) GetEx(ctx context.Context, key string, expiration time.Duration) *redis.StringCmd {
	return r.client.GetEx(ctx, key, expiration)
}

// HRandField implements redis.Cmdable
func (r *Redis) HRandField(ctx context.Context, key string, count int) *redis.StringSliceCmd {
	return r.client.HRandField(ctx, key, count)
}

// LMove implements redis.Cmdable
func (r *Redis) LMove(ctx context.Context, source string, destination string, srcpos string, destpos string) *redis.StringCmd {
	return r.client.LMove(ctx, source, destination, srcpos, destpos)
}

// LPopCount implements redis.Cmdable
func (r *Redis) LPopCount(ctx context.Context, key string, count int) *redis.StringSliceCmd {
	return r.client.LPopCount(ctx, key, count)
}

// LPos implements redis.Cmdable
func (r *Redis) LPos(ctx context.Context, key string, value string, args redis.LPosArgs) *redis.IntCmd {
	return r.client.LPos(ctx, key, value, args)
}

// LPosCount implements redis.Cmdable
func (r *Redis) LPosCount(ctx context.Context, key string, value string, count int64, args redis.LPosArgs) *redis.IntSliceCmd {
	return r.client.LPosCount(ctx, key, value, count, args)
}

// RPopCount implements redis.Cmdable
func (r *Redis) RPopCount(ctx context.Context, key string, count int) *redis.StringSliceCmd {
	return r.client.RPopCount(ctx, key, count)
}

// SMIsMember implements redis.Cmdable
func (r *Redis) SMIsMember(ctx context.Context, key string, members ...interface{}) *redis.BoolSliceCmd {
	return r.client.SMIsMember(ctx, key, members...)
}

// ScanType implements redis.Cmdable
func (r *Redis) ScanType(ctx context.Context, cursor uint64, match string, count int64, keyType string) *redis.ScanCmd {
	return r.client.ScanType(ctx, cursor, match, count, keyType)
}

// SetArgs implements redis.Cmdable
func (r *Redis) SetArgs(ctx context.Context, key string, value interface{}, a redis.SetArgs) *redis.StatusCmd {
	return r.client.SetArgs(ctx, key, value, a)
}

// XAutoClaim implements redis.Cmdable
func (r *Redis) XAutoClaim(ctx context.Context, a *redis.XAutoClaimArgs) *redis.XAutoClaimCmd {
	return r.client.XAutoClaim(ctx, a)
}

// XAutoClaimJustID implements redis.Cmdable
func (r *Redis) XAutoClaimJustID(ctx context.Context, a *redis.XAutoClaimArgs) *redis.XAutoClaimJustIDCmd {
	return r.client.XAutoClaimJustID(ctx, a)
}

// XGroupCreateConsumer implements redis.Cmdable
func (r *Redis) XGroupCreateConsumer(ctx context.Context, stream string, group string, consumer string) *redis.IntCmd {
	return r.client.XGroupCreateConsumer(ctx, stream, group, consumer)
}

// XInfoConsumers implements redis.Cmdable
func (r *Redis) XInfoConsumers(ctx context.Context, key string, group string) *redis.XInfoConsumersCmd {
	return r.client.XInfoConsumers(ctx, key, group)
}

// XInfoGroups implements redis.Cmdable
func (r *Redis) XInfoGroups(ctx context.Context, key string) *redis.XInfoGroupsCmd {
	return r.client.XInfoGroups(ctx, key)
}

// XInfoStream implements redis.Cmdable
func (r *Redis) XInfoStream(ctx context.Context, key string) *redis.XInfoStreamCmd {
	return r.client.XInfoStream(ctx, key)
}

// XInfoStreamFull implements redis.Cmdable
func (r *Redis) XInfoStreamFull(ctx context.Context, key string, count int) *redis.XInfoStreamFullCmd {
	return r.client.XInfoStreamFull(ctx, key, count)

}

// XTrimMaxLen implements redis.Cmdable
func (r *Redis) XTrimMaxLen(ctx context.Context, key string, maxLen int64) *redis.IntCmd {
	return r.client.XTrimMaxLen(ctx, key, maxLen)
}

// XTrimMaxLenApprox implements redis.Cmdable
func (r *Redis) XTrimMaxLenApprox(ctx context.Context, key string, maxLen int64, limit int64) *redis.IntCmd {
	return r.client.XTrimMaxLenApprox(ctx, key, maxLen, limit)
}

// XTrimMinID implements redis.Cmdable
func (r *Redis) XTrimMinID(ctx context.Context, key string, minID string) *redis.IntCmd {
	return r.client.XTrimMinID(ctx, key, minID)
}

// XTrimMinIDApprox implements redis.Cmdable
func (r *Redis) XTrimMinIDApprox(ctx context.Context, key string, minID string, limit int64) *redis.IntCmd {
	return r.client.XTrimMinIDApprox(ctx, key, minID, limit)
}

// ZAddArgs implements redis.Cmdable
func (r *Redis) ZAddArgs(ctx context.Context, key string, args redis.ZAddArgs) *redis.IntCmd {
	return r.client.ZAddArgs(ctx, key, args)
}

// ZAddArgsIncr implements redis.Cmdable
func (r *Redis) ZAddArgsIncr(ctx context.Context, key string, args redis.ZAddArgs) *redis.FloatCmd {
	return r.client.ZAddArgsIncr(ctx, key, args)
}

// ZDiff implements redis.Cmdable
func (r *Redis) ZDiff(ctx context.Context, keys ...string) *redis.StringSliceCmd {
	return r.client.ZDiff(ctx, keys...)
}

// ZDiffStore implements redis.Cmdable
func (r *Redis) ZDiffStore(ctx context.Context, destination string, keys ...string) *redis.IntCmd {
	return r.client.ZDiffStore(ctx, destination, keys...)
}

// ZDiffWithScores implements redis.Cmdable
func (r *Redis) ZDiffWithScores(ctx context.Context, keys ...string) *redis.ZSliceCmd {
	return r.client.ZDiffWithScores(ctx, keys...)
}

// ZInter implements redis.Cmdable
func (r *Redis) ZInter(ctx context.Context, store *redis.ZStore) *redis.StringSliceCmd {
	return r.client.ZInter(ctx, store)
}

// ZInterWithScores implements redis.Cmdable
func (r *Redis) ZInterWithScores(ctx context.Context, store *redis.ZStore) *redis.ZSliceCmd {
	return r.client.ZInterWithScores(ctx, store)
}

// ZMScore implements redis.Cmdable
func (r *Redis) ZMScore(ctx context.Context, key string, members ...string) *redis.FloatSliceCmd {
	return r.client.ZMScore(ctx, key, members...)
}

// ZRandMember implements redis.Cmdable
func (r *Redis) ZRandMember(ctx context.Context, key string, count int) *redis.StringSliceCmd {
	return r.client.ZRandMember(ctx, key, count)
}

// ZRangeArgs implements redis.Cmdable
func (r *Redis) ZRangeArgs(ctx context.Context, z redis.ZRangeArgs) *redis.StringSliceCmd {
	return r.client.ZRangeArgs(ctx, z)
}

// ZRangeArgsWithScores implements redis.Cmdable
func (r *Redis) ZRangeArgsWithScores(ctx context.Context, z redis.ZRangeArgs) *redis.ZSliceCmd {
	return r.client.ZRangeArgsWithScores(ctx, z)
}

// ZRangeStore implements redis.Cmdable
func (r *Redis) ZRangeStore(ctx context.Context, dst string, z redis.ZRangeArgs) *redis.IntCmd {
	return r.client.ZRangeStore(ctx, dst, z)
}

// ZUnion implements redis.Cmdable
func (r *Redis) ZUnion(ctx context.Context, store redis.ZStore) *redis.StringSliceCmd {
	return r.client.ZUnion(ctx, store)
}

// ZUnionWithScores implements redis.Cmdable
func (r *Redis) ZUnionWithScores(ctx context.Context, store redis.ZStore) *redis.ZSliceCmd {
	return r.client.ZUnionWithScores(ctx, store)
}

func (r *Redis) Command(ctx context.Context) *redis.CommandsInfoCmd {
	return r.client.Command(ctx)
}

func (r *Redis) ClientGetName(ctx context.Context) *redis.StringCmd {
	return r.client.ClientGetName(ctx)
}

func (r *Redis) Echo(ctx context.Context, message interface{}) *redis.StringCmd {
	return r.client.Echo(ctx, message)
}

func (r *Redis) Ping(ctx context.Context) *redis.StatusCmd {
	return r.client.Ping(ctx)
}

func (r *Redis) Quit(ctx context.Context) *redis.StatusCmd {
	return r.client.Quit(ctx)
}

func (r *Redis) Del(ctx context.Context, keys ...string) *redis.IntCmd {
	return r.client.Del(ctx, keys...)
}

func (r *Redis) Unlink(ctx context.Context, keys ...string) *redis.IntCmd {
	return r.client.Unlink(ctx, keys...)
}

func (r *Redis) Dump(ctx context.Context, key string) *redis.StringCmd {
	return r.client.Dump(ctx, key)
}

func (r *Redis) Exists(ctx context.Context, keys ...string) *redis.IntCmd {
	return r.client.Exists(ctx, keys...)
}

func (r *Redis) Expire(ctx context.Context, key string, expiration time.Duration) *redis.BoolCmd {
	return r.client.Expire(ctx, key, expiration)
}

func (r *Redis) ExpireAt(ctx context.Context, key string, tm time.Time) *redis.BoolCmd {
	return r.client.ExpireAt(ctx, key, tm)
}

func (r *Redis) Keys(ctx context.Context, pattern string) *redis.StringSliceCmd {
	return r.client.Keys(ctx, pattern)
}

func (r *Redis) Migrate(ctx context.Context, host, port, key string, db int, timeout time.Duration) *redis.StatusCmd {
	return r.client.Migrate(ctx, host, port, key, db, timeout)
}

func (r *Redis) Move(ctx context.Context, key string, db int) *redis.BoolCmd {
	return r.client.Move(ctx, key, db)
}

func (r *Redis) ObjectRefCount(ctx context.Context, key string) *redis.IntCmd {
	return r.client.ObjectRefCount(ctx, key)
}

func (r *Redis) ObjectEncoding(ctx context.Context, key string) *redis.StringCmd {
	return r.client.ObjectEncoding(ctx, key)
}

func (r *Redis) ObjectIdleTime(ctx context.Context, key string) *redis.DurationCmd {
	return r.client.ObjectIdleTime(ctx, key)
}

func (r *Redis) Persist(ctx context.Context, key string) *redis.BoolCmd {
	return r.client.Persist(ctx, key)
}

func (r *Redis) PExpire(ctx context.Context, key string, expiration time.Duration) *redis.BoolCmd {
	return r.client.PExpire(ctx, key, expiration)
}

func (r *Redis) PExpireAt(ctx context.Context, key string, tm time.Time) *redis.BoolCmd {
	return r.client.PExpireAt(ctx, key, tm)
}

func (r *Redis) PTTL(ctx context.Context, key string) *redis.DurationCmd {
	return r.client.PTTL(ctx, key)
}

func (r *Redis) RandomKey(ctx context.Context) *redis.StringCmd {
	return r.client.RandomKey(ctx)
}

func (r *Redis) Rename(ctx context.Context, key, newkey string) *redis.StatusCmd {
	return r.client.Rename(ctx, key, newkey)
}

func (r *Redis) RenameNX(ctx context.Context, key, newkey string) *redis.BoolCmd {
	return r.client.RenameNX(ctx, key, newkey)
}

func (r *Redis) Restore(ctx context.Context, key string, ttl time.Duration, value string) *redis.StatusCmd {
	return r.client.Restore(ctx, key, ttl, value)
}

func (r *Redis) RestoreReplace(ctx context.Context, key string, ttl time.Duration, value string) *redis.StatusCmd {
	return r.client.RestoreReplace(ctx, key, ttl, value)
}

func (r *Redis) Sort(ctx context.Context, key string, sort *redis.Sort) *redis.StringSliceCmd {
	return r.client.Sort(ctx, key, sort)
}

func (r *Redis) SortStore(ctx context.Context, key, store string, sort *redis.Sort) *redis.IntCmd {
	return r.client.SortStore(ctx, key, store, sort)
}

func (r *Redis) SortInterfaces(ctx context.Context, key string, sort *redis.Sort) *redis.SliceCmd {
	return r.client.SortInterfaces(ctx, key, sort)
}

func (r *Redis) Touch(ctx context.Context, keys ...string) *redis.IntCmd {
	return r.client.Touch(ctx, keys...)
}

func (r *Redis) TTL(ctx context.Context, key string) *redis.DurationCmd {
	return r.client.TTL(ctx, key)
}

func (r *Redis) Type(ctx context.Context, key string) *redis.StatusCmd {
	return r.client.Type(ctx, key)
}

func (r *Redis) Scan(ctx context.Context, cursor uint64, match string, count int64) *redis.ScanCmd {
	return r.client.Scan(ctx, cursor, match, count)
}

func (r *Redis) SScan(ctx context.Context, key string, cursor uint64, match string, count int64) *redis.ScanCmd {
	return r.client.SScan(ctx, key, cursor, match, count)
}

func (r *Redis) HScan(ctx context.Context, key string, cursor uint64, match string, count int64) *redis.ScanCmd {
	return r.client.HScan(ctx, key, cursor, match, count)
}

func (r *Redis) ZScan(ctx context.Context, key string, cursor uint64, match string, count int64) *redis.ScanCmd {
	return r.client.ZScan(ctx, key, cursor, match, count)
}

func (r *Redis) Append(ctx context.Context, key, value string) *redis.IntCmd {
	return r.client.Append(ctx, key, value)
}

func (r *Redis) BitCount(ctx context.Context, key string, bitCount *redis.BitCount) *redis.IntCmd {
	return r.client.BitCount(ctx, key, bitCount)
}

func (r *Redis) BitOpAnd(ctx context.Context, destKey string, keys ...string) *redis.IntCmd {
	return r.client.BitOpAnd(ctx, destKey, keys...)
}

func (r *Redis) BitOpOr(ctx context.Context, destKey string, keys ...string) *redis.IntCmd {
	return r.client.BitOpOr(ctx, destKey, keys...)
}

func (r *Redis) BitOpXor(ctx context.Context, destKey string, keys ...string) *redis.IntCmd {
	return r.client.BitOpXor(ctx, destKey, keys...)
}

func (r *Redis) BitOpNot(ctx context.Context, destKey string, key string) *redis.IntCmd {
	return r.client.BitOpNot(ctx, destKey, key)
}

func (r *Redis) BitPos(ctx context.Context, key string, bit int64, pos ...int64) *redis.IntCmd {
	return r.client.BitPos(ctx, key, bit, pos...)
}

func (r *Redis) Decr(ctx context.Context, key string) *redis.IntCmd {
	return r.client.Decr(ctx, key)
}

func (r *Redis) DecrBy(ctx context.Context, key string, decrement int64) *redis.IntCmd {
	return r.client.DecrBy(ctx, key, decrement)
}

func (r *Redis) Get(ctx context.Context, key string) *redis.StringCmd {
	return r.client.Get(ctx, key)
}

func (r *Redis) GetBit(ctx context.Context, key string, offset int64) *redis.IntCmd {
	return r.client.GetBit(ctx, key, offset)
}

func (r *Redis) GetRange(ctx context.Context, key string, start, end int64) *redis.StringCmd {
	return r.client.GetRange(ctx, key, start, end)
}

func (r *Redis) GetSet(ctx context.Context, key string, value interface{}) *redis.StringCmd {
	return r.client.GetSet(ctx, key, value)
}

func (r *Redis) Incr(ctx context.Context, key string) *redis.IntCmd {
	return r.client.Incr(ctx, key)
}

func (r *Redis) IncrBy(ctx context.Context, key string, value int64) *redis.IntCmd {
	return r.client.IncrBy(ctx, key, value)
}

func (r *Redis) IncrByFloat(ctx context.Context, key string, value float64) *redis.FloatCmd {
	return r.client.IncrByFloat(ctx, key, value)
}

func (r *Redis) MGet(ctx context.Context, keys ...string) *redis.SliceCmd {
	return r.client.MGet(ctx, keys...)
}

func (r *Redis) MSet(ctx context.Context, pairs ...interface{}) *redis.StatusCmd {
	return r.client.MSet(ctx, pairs...)
}

func (r *Redis) MSetNX(ctx context.Context, pairs ...interface{}) *redis.BoolCmd {
	return r.client.MSetNX(ctx, pairs...)
}

func (r *Redis) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) *redis.StatusCmd {
	return r.client.Set(ctx, key, value, expiration)
}

func (r *Redis) SetBit(ctx context.Context, key string, offset int64, value int) *redis.IntCmd {
	return r.client.SetBit(ctx, key, offset, value)
}

func (r *Redis) SetNX(ctx context.Context, key string, value interface{}, expiration time.Duration) *redis.BoolCmd {
	return r.client.SetNX(ctx, key, value, expiration)
}

func (r *Redis) SetXX(ctx context.Context, key string, value interface{}, expiration time.Duration) *redis.BoolCmd {
	return r.client.SetXX(ctx, key, value, expiration)
}

func (r *Redis) SetRange(ctx context.Context, key string, offset int64, value string) *redis.IntCmd {
	return r.client.SetRange(ctx, key, offset, value)
}

func (r *Redis) StrLen(ctx context.Context, key string) *redis.IntCmd {
	return r.client.StrLen(ctx, key)
}

func (r *Redis) HDel(ctx context.Context, key string, fields ...string) *redis.IntCmd {
	return r.client.HDel(ctx, key, fields...)
}

func (r *Redis) HExists(ctx context.Context, key, field string) *redis.BoolCmd {
	return r.client.HExists(ctx, key, field)
}

func (r *Redis) HGet(ctx context.Context, key, field string) *redis.StringCmd {
	return r.client.HGet(ctx, key, field)
}

func (r *Redis) HGetAll(ctx context.Context, key string) *redis.MapStringStringCmd {
	return r.client.HGetAll(ctx, key)
}

func (r *Redis) HIncrBy(ctx context.Context, key, field string, incr int64) *redis.IntCmd {
	return r.client.HIncrBy(ctx, key, field, incr)
}

func (r *Redis) HIncrByFloat(ctx context.Context, key, field string, incr float64) *redis.FloatCmd {
	return r.client.HIncrByFloat(ctx, key, field, incr)
}

func (r *Redis) HKeys(ctx context.Context, key string) *redis.StringSliceCmd {
	return r.client.HKeys(ctx, key)
}

func (r *Redis) HLen(ctx context.Context, key string) *redis.IntCmd {
	return r.client.HLen(ctx, key)
}

func (r *Redis) HMGet(ctx context.Context, key string, fields ...string) *redis.SliceCmd {
	return r.client.HMGet(ctx, key, fields...)
}

func (r *Redis) HMSet(ctx context.Context, key string, values ...interface{}) *redis.BoolCmd {
	return r.client.HMSet(ctx, key, values...)
}

func (r *Redis) HSet(ctx context.Context, key string, values ...interface{}) *redis.IntCmd {
	return r.client.HSet(ctx, key, values...)
}

func (r *Redis) HSetNX(ctx context.Context, key, field string, value interface{}) *redis.BoolCmd {
	return r.client.HSetNX(ctx, key, field, value)
}

func (r *Redis) HVals(ctx context.Context, key string) *redis.StringSliceCmd {
	return r.client.HVals(ctx, key)
}

func (r *Redis) BLPop(ctx context.Context, timeout time.Duration, keys ...string) *redis.StringSliceCmd {
	return r.client.BLPop(ctx, timeout, keys...)
}

func (r *Redis) BRPop(ctx context.Context, timeout time.Duration, keys ...string) *redis.StringSliceCmd {
	return r.client.BRPop(ctx, timeout, keys...)
}

func (r *Redis) BRPopLPush(ctx context.Context, source, destination string, timeout time.Duration) *redis.StringCmd {
	return r.client.BRPopLPush(ctx, source, destination, timeout)
}

func (r *Redis) LIndex(ctx context.Context, key string, index int64) *redis.StringCmd {
	return r.client.LIndex(ctx, key, index)
}

func (r *Redis) LInsert(ctx context.Context, key, op string, pivot, value interface{}) *redis.IntCmd {
	return r.client.LInsert(ctx, key, op, pivot, value)
}

func (r *Redis) LInsertBefore(ctx context.Context, key string, pivot, value interface{}) *redis.IntCmd {
	return r.client.LInsertBefore(ctx, key, pivot, value)
}

func (r *Redis) LInsertAfter(ctx context.Context, key string, pivot, value interface{}) *redis.IntCmd {
	return r.client.LInsertAfter(ctx, key, pivot, value)
}

func (r *Redis) LLen(ctx context.Context, key string) *redis.IntCmd {
	return r.client.LLen(ctx, key)
}

func (r *Redis) LPop(ctx context.Context, key string) *redis.StringCmd {
	return r.client.LPop(ctx, key)
}

func (r *Redis) LPush(ctx context.Context, key string, values ...interface{}) *redis.IntCmd {
	return r.client.LPush(ctx, key, values...)
}

func (r *Redis) LPushX(ctx context.Context, key string, values ...interface{}) *redis.IntCmd {
	return r.client.LPushX(ctx, key, values...)
}

func (r *Redis) LRange(ctx context.Context, key string, start, stop int64) *redis.StringSliceCmd {
	return r.client.LRange(ctx, key, start, stop)
}

func (r *Redis) LRem(ctx context.Context, key string, count int64, value interface{}) *redis.IntCmd {
	return r.client.LRem(ctx, key, count, value)
}

func (r *Redis) LSet(ctx context.Context, key string, index int64, value interface{}) *redis.StatusCmd {
	return r.client.LSet(ctx, key, index, value)
}

func (r *Redis) LTrim(ctx context.Context, key string, start, stop int64) *redis.StatusCmd {
	return r.client.LTrim(ctx, key, start, stop)
}

func (r *Redis) RPop(ctx context.Context, key string) *redis.StringCmd {
	return r.client.RPop(ctx, key)
}

func (r *Redis) RPopLPush(ctx context.Context, source, destination string) *redis.StringCmd {
	return r.client.RPopLPush(ctx, source, destination)
}

func (r *Redis) RPush(ctx context.Context, key string, values ...interface{}) *redis.IntCmd {
	return r.client.RPush(ctx, key, values...)
}

func (r *Redis) RPushX(ctx context.Context, key string, values ...interface{}) *redis.IntCmd {
	return r.client.RPushX(ctx, key, values...)
}

func (r *Redis) SAdd(ctx context.Context, key string, members ...interface{}) *redis.IntCmd {
	return r.client.SAdd(ctx, key, members...)
}

func (r *Redis) SCard(ctx context.Context, key string) *redis.IntCmd {
	return r.client.SCard(ctx, key)
}

func (r *Redis) SDiff(ctx context.Context, keys ...string) *redis.StringSliceCmd {
	return r.client.SDiff(ctx, keys...)
}

func (r *Redis) SDiffStore(ctx context.Context, destination string, keys ...string) *redis.IntCmd {
	return r.client.SDiffStore(ctx, destination, keys...)
}

func (r *Redis) SInter(ctx context.Context, keys ...string) *redis.StringSliceCmd {
	return r.client.SInter(ctx, keys...)
}

func (r *Redis) SInterStore(ctx context.Context, destination string, keys ...string) *redis.IntCmd {
	return r.client.SInterStore(ctx, destination, keys...)
}

func (r *Redis) SIsMember(ctx context.Context, key string, member interface{}) *redis.BoolCmd {
	return r.client.SIsMember(ctx, key, member)
}

func (r *Redis) SMembers(ctx context.Context, key string) *redis.StringSliceCmd {
	return r.client.SMembers(ctx, key)
}

func (r *Redis) SMembersMap(ctx context.Context, key string) *redis.StringStructMapCmd {
	return r.client.SMembersMap(ctx, key)
}

func (r *Redis) SMove(ctx context.Context, source, destination string, member interface{}) *redis.BoolCmd {
	return r.client.SMove(ctx, source, destination, member)
}

func (r *Redis) SPop(ctx context.Context, key string) *redis.StringCmd {
	return r.client.SPop(ctx, key)
}

func (r *Redis) SPopN(ctx context.Context, key string, count int64) *redis.StringSliceCmd {
	return r.client.SPopN(ctx, key, count)
}

func (r *Redis) SRandMember(ctx context.Context, key string) *redis.StringCmd {
	return r.client.SRandMember(ctx, key)
}

func (r *Redis) SRandMemberN(ctx context.Context, key string, count int64) *redis.StringSliceCmd {
	return r.client.SRandMemberN(ctx, key, count)
}

func (r *Redis) SRem(ctx context.Context, key string, members ...interface{}) *redis.IntCmd {
	return r.client.SRem(ctx, key, members...)
}

func (r *Redis) SUnion(ctx context.Context, keys ...string) *redis.StringSliceCmd {
	return r.client.SUnion(ctx, keys...)
}

func (r *Redis) SUnionStore(ctx context.Context, destination string, keys ...string) *redis.IntCmd {
	return r.client.SUnionStore(ctx, destination, keys...)
}

func (r *Redis) XAdd(ctx context.Context, a *redis.XAddArgs) *redis.StringCmd {
	return r.client.XAdd(ctx, a)
}

func (r *Redis) XDel(ctx context.Context, stream string, ids ...string) *redis.IntCmd {
	return r.client.XDel(ctx, stream, ids...)
}

func (r *Redis) XLen(ctx context.Context, stream string) *redis.IntCmd {
	return r.client.XLen(ctx, stream)
}

func (r *Redis) XRange(ctx context.Context, stream, start, stop string) *redis.XMessageSliceCmd {
	return r.client.XRange(ctx, stream, start, stop)
}

func (r *Redis) XRangeN(ctx context.Context, stream, start, stop string, count int64) *redis.XMessageSliceCmd {
	return r.client.XRangeN(ctx, stream, start, stop, count)
}

func (r *Redis) XRevRange(ctx context.Context, stream string, start, stop string) *redis.XMessageSliceCmd {
	return r.client.XRevRange(ctx, stream, start, stop)
}

func (r *Redis) XRevRangeN(ctx context.Context, stream string, start, stop string, count int64) *redis.XMessageSliceCmd {
	return r.client.XRevRangeN(ctx, stream, start, stop, count)
}

func (r *Redis) XRead(ctx context.Context, a *redis.XReadArgs) *redis.XStreamSliceCmd {
	return r.client.XRead(ctx, a)
}

func (r *Redis) XReadStreams(ctx context.Context, streams ...string) *redis.XStreamSliceCmd {
	return r.client.XReadStreams(ctx, streams...)
}

func (r *Redis) XGroupCreate(ctx context.Context, stream, group, start string) *redis.StatusCmd {
	return r.client.XGroupCreate(ctx, stream, group, start)
}

func (r *Redis) XGroupCreateMkStream(ctx context.Context, stream, group, start string) *redis.StatusCmd {
	return r.client.XGroupCreateMkStream(ctx, stream, group, start)
}

func (r *Redis) XGroupSetID(ctx context.Context, stream, group, start string) *redis.StatusCmd {
	return r.client.XGroupSetID(ctx, stream, group, start)
}

func (r *Redis) XGroupDestroy(ctx context.Context, stream, group string) *redis.IntCmd {
	return r.client.XGroupDestroy(ctx, stream, group)
}

func (r *Redis) XGroupDelConsumer(ctx context.Context, stream, group, consumer string) *redis.IntCmd {
	return r.client.XGroupDelConsumer(ctx, stream, group, consumer)
}

func (r *Redis) XReadGroup(ctx context.Context, a *redis.XReadGroupArgs) *redis.XStreamSliceCmd {
	return r.client.XReadGroup(ctx, a)
}

func (r *Redis) XAck(ctx context.Context, stream, group string, ids ...string) *redis.IntCmd {
	return r.client.XAck(ctx, stream, group, ids...)
}

func (r *Redis) XPending(ctx context.Context, stream, group string) *redis.XPendingCmd {
	return r.client.XPending(ctx, stream, group)
}

func (r *Redis) XPendingExt(ctx context.Context, a *redis.XPendingExtArgs) *redis.XPendingExtCmd {
	return r.client.XPendingExt(ctx, a)
}

func (r *Redis) XClaim(ctx context.Context, a *redis.XClaimArgs) *redis.XMessageSliceCmd {
	return r.client.XClaim(ctx, a)
}

func (r *Redis) XClaimJustID(ctx context.Context, a *redis.XClaimArgs) *redis.StringSliceCmd {
	return r.client.XClaimJustID(ctx, a)
}

func (r *Redis) BZPopMax(ctx context.Context, timeout time.Duration, keys ...string) *redis.ZWithKeyCmd {
	return r.client.BZPopMax(ctx, timeout, keys...)
}

func (r *Redis) BZPopMin(ctx context.Context, timeout time.Duration, keys ...string) *redis.ZWithKeyCmd {
	return r.client.BZPopMin(ctx, timeout, keys...)
}

func (r *Redis) ZAdd(ctx context.Context, key string, members ...redis.Z) *redis.IntCmd {
	return r.client.ZAdd(ctx, key, members...)
}

func (r *Redis) ZAddNX(ctx context.Context, key string, members ...redis.Z) *redis.IntCmd {
	return r.client.ZAddNX(ctx, key, members...)
}

func (r *Redis) ZAddXX(ctx context.Context, key string, members ...redis.Z) *redis.IntCmd {
	return r.client.ZAddXX(ctx, key, members...)
}

func (r *Redis) ZCard(ctx context.Context, key string) *redis.IntCmd {
	return r.client.ZCard(ctx, key)
}

func (r *Redis) ZCount(ctx context.Context, key, min, max string) *redis.IntCmd {
	return r.client.ZCount(ctx, key, min, max)
}

func (r *Redis) ZLexCount(ctx context.Context, key, min, max string) *redis.IntCmd {
	return r.client.ZLexCount(ctx, key, min, max)
}

func (r *Redis) ZIncrBy(ctx context.Context, key string, increment float64, member string) *redis.FloatCmd {
	return r.client.ZIncrBy(ctx, key, increment, member)
}

func (r *Redis) ZInterStore(ctx context.Context, destination string, store *redis.ZStore) *redis.IntCmd {
	return r.client.ZInterStore(ctx, destination, store)
}

func (r *Redis) ZPopMax(ctx context.Context, key string, count ...int64) *redis.ZSliceCmd {
	return r.client.ZPopMax(ctx, key, count...)
}

func (r *Redis) ZPopMin(ctx context.Context, key string, count ...int64) *redis.ZSliceCmd {
	return r.client.ZPopMin(ctx, key, count...)
}

func (r *Redis) ZRange(ctx context.Context, key string, start, stop int64) *redis.StringSliceCmd {
	return r.client.ZRange(ctx, key, start, stop)
}

func (r *Redis) ZRangeWithScores(ctx context.Context, key string, start, stop int64) *redis.ZSliceCmd {
	return r.client.ZRangeWithScores(ctx, key, start, stop)
}

func (r *Redis) ZRangeByScore(ctx context.Context, key string, opt *redis.ZRangeBy) *redis.StringSliceCmd {
	return r.client.ZRangeByScore(ctx, key, opt)
}

func (r *Redis) ZRangeByLex(ctx context.Context, key string, opt *redis.ZRangeBy) *redis.StringSliceCmd {
	return r.client.ZRangeByLex(ctx, key, opt)
}

func (r *Redis) ZRangeByScoreWithScores(ctx context.Context, key string, opt *redis.ZRangeBy) *redis.ZSliceCmd {
	return r.client.ZRangeByScoreWithScores(ctx, key, opt)
}

func (r *Redis) ZRank(ctx context.Context, key, member string) *redis.IntCmd {
	return r.client.ZRank(ctx, key, member)
}

func (r *Redis) ZRem(ctx context.Context, key string, members ...interface{}) *redis.IntCmd {
	return r.client.ZRem(ctx, key, members)
}

func (r *Redis) ZRemRangeByRank(ctx context.Context, key string, start, stop int64) *redis.IntCmd {
	return r.client.ZRemRangeByRank(ctx, key, start, stop)
}

func (r *Redis) ZRemRangeByScore(ctx context.Context, key, min, max string) *redis.IntCmd {
	return r.client.ZRemRangeByScore(ctx, key, min, max)
}

func (r *Redis) ZRemRangeByLex(ctx context.Context, key, min, max string) *redis.IntCmd {
	return r.client.ZRemRangeByLex(ctx, key, min, max)
}

func (r *Redis) ZRevRange(ctx context.Context, key string, start, stop int64) *redis.StringSliceCmd {
	return r.client.ZRevRange(ctx, key, start, stop)
}

func (r *Redis) ZRevRangeWithScores(ctx context.Context, key string, start, stop int64) *redis.ZSliceCmd {
	return r.client.ZRevRangeWithScores(ctx, key, start, stop)
}

func (r *Redis) ZRevRangeByScore(ctx context.Context, key string, opt *redis.ZRangeBy) *redis.StringSliceCmd {
	return r.client.ZRevRangeByScore(ctx, key, opt)
}

func (r *Redis) ZRevRangeByLex(ctx context.Context, key string, opt *redis.ZRangeBy) *redis.StringSliceCmd {
	return r.client.ZRevRangeByLex(ctx, key, opt)
}

func (r *Redis) ZRevRangeByScoreWithScores(ctx context.Context, key string, opt *redis.ZRangeBy) *redis.ZSliceCmd {
	return r.client.ZRevRangeByScoreWithScores(ctx, key, opt)
}

func (r *Redis) ZRevRank(ctx context.Context, key, member string) *redis.IntCmd {
	return r.client.ZRevRank(ctx, key, member)
}

func (r *Redis) ZScore(ctx context.Context, key, member string) *redis.FloatCmd {
	return r.client.ZScore(ctx, key, member)
}

func (r *Redis) ZUnionStore(ctx context.Context, dest string, store *redis.ZStore) *redis.IntCmd {
	return r.client.ZUnionStore(ctx, dest, store)
}

func (r *Redis) PFAdd(ctx context.Context, key string, els ...interface{}) *redis.IntCmd {
	return r.client.PFAdd(ctx, key, els...)
}

func (r *Redis) PFCount(ctx context.Context, keys ...string) *redis.IntCmd {
	return r.client.PFCount(ctx, keys...)
}

func (r *Redis) PFMerge(ctx context.Context, dest string, keys ...string) *redis.StatusCmd {
	return r.client.PFMerge(ctx, dest, keys...)
}

func (r *Redis) BgRewriteAOF(ctx context.Context) *redis.StatusCmd {
	return r.client.BgRewriteAOF(ctx)
}

func (r *Redis) BgSave(ctx context.Context) *redis.StatusCmd {
	return r.client.BgSave(ctx)
}

func (r *Redis) ClientKill(ctx context.Context, ipPort string) *redis.StatusCmd {
	return r.client.ClientKill(ctx, ipPort)
}

func (r *Redis) ClientKillByFilter(ctx context.Context, keys ...string) *redis.IntCmd {
	return r.client.ClientKillByFilter(ctx, keys...)
}

func (r *Redis) ClientList(ctx context.Context) *redis.StringCmd {
	return r.client.ClientList(ctx)
}

func (r *Redis) ClientPause(ctx context.Context, dur time.Duration) *redis.BoolCmd {
	return r.client.ClientPause(ctx, dur)
}

func (r *Redis) ClientID(ctx context.Context) *redis.IntCmd {
	return r.client.ClientID(ctx)
}

func (r *Redis) ConfigResetStat(ctx context.Context) *redis.StatusCmd {
	return r.client.ConfigResetStat(ctx)
}

func (r *Redis) ConfigSet(ctx context.Context, parameter, value string) *redis.StatusCmd {
	return r.client.ConfigSet(ctx, parameter, value)
}

func (r *Redis) ConfigRewrite(ctx context.Context) *redis.StatusCmd {
	return r.client.ConfigRewrite(ctx)
}

func (r *Redis) DBSize(ctx context.Context) *redis.IntCmd {
	return r.client.DBSize(ctx)
}

func (r *Redis) FlushAll(ctx context.Context) *redis.StatusCmd {
	return r.client.FlushAll(ctx)
}

func (r *Redis) FlushAllAsync(ctx context.Context) *redis.StatusCmd {
	return r.client.FlushAllAsync(ctx)
}

func (r *Redis) FlushDB(ctx context.Context) *redis.StatusCmd {
	return r.client.FlushDB(ctx)
}

func (r *Redis) FlushDBAsync(ctx context.Context) *redis.StatusCmd {
	return r.client.FlushDBAsync(ctx)
}

func (r *Redis) Info(ctx context.Context, section ...string) *redis.StringCmd {
	return r.client.Info(ctx, section...)
}

func (r *Redis) LastSave(ctx context.Context) *redis.IntCmd {
	return r.client.LastSave(ctx)
}

func (r *Redis) Save(ctx context.Context) *redis.StatusCmd {
	return r.client.Save(ctx)
}

func (r *Redis) Shutdown(ctx context.Context) *redis.StatusCmd {
	return r.client.Shutdown(ctx)
}

func (r *Redis) ShutdownSave(ctx context.Context) *redis.StatusCmd {
	return r.client.ShutdownSave(ctx)
}

func (r *Redis) ShutdownNoSave(ctx context.Context) *redis.StatusCmd {
	return r.client.ShutdownNoSave(ctx)
}

func (r *Redis) SlaveOf(ctx context.Context, host, port string) *redis.StatusCmd {
	return r.client.SlaveOf(ctx, host, port)
}

func (r *Redis) Time(ctx context.Context) *redis.TimeCmd {
	return r.client.Time(ctx)
}

func (r *Redis) Eval(ctx context.Context, script string, keys []string, args ...interface{}) *redis.Cmd {
	return r.client.Eval(ctx, script, keys, args...)
}

func (r *Redis) EvalSha(ctx context.Context, sha1 string, keys []string, args ...interface{}) *redis.Cmd {
	return r.client.EvalSha(ctx, sha1, keys, args...)
}

func (r *Redis) ScriptExists(ctx context.Context, hashes ...string) *redis.BoolSliceCmd {
	return r.client.ScriptExists(ctx, hashes...)
}

func (r *Redis) ScriptFlush(ctx context.Context) *redis.StatusCmd {
	return r.client.ScriptFlush(ctx)
}

func (r *Redis) ScriptKill(ctx context.Context) *redis.StatusCmd {
	return r.client.ScriptKill(ctx)
}

func (r *Redis) ScriptLoad(ctx context.Context, script string) *redis.StringCmd {
	return r.client.ScriptLoad(ctx, script)
}

func (r *Redis) DebugObject(ctx context.Context, key string) *redis.StringCmd {
	return r.client.DebugObject(ctx, key)
}

func (r *Redis) Publish(ctx context.Context, channel string, message interface{}) *redis.IntCmd {
	return r.client.Publish(ctx, channel, message)
}

func (r *Redis) PubSubChannels(ctx context.Context, pattern string) *redis.StringSliceCmd {
	return r.client.PubSubChannels(ctx, pattern)
}

func (r *Redis) PubSubNumSub(ctx context.Context, channels ...string) *redis.MapStringIntCmd {
	return r.client.PubSubNumSub(ctx, channels...)
}

func (r *Redis) PubSubNumPat(ctx context.Context) *redis.IntCmd {
	return r.client.PubSubNumPat(ctx)
}

func (r *Redis) ClusterSlots(ctx context.Context) *redis.ClusterSlotsCmd {
	return r.client.ClusterSlots(ctx)
}

func (r *Redis) ClusterNodes(ctx context.Context) *redis.StringCmd {
	return r.client.ClusterNodes(ctx)
}

func (r *Redis) ClusterMeet(ctx context.Context, host, port string) *redis.StatusCmd {
	return r.client.ClusterMeet(ctx, host, port)
}

func (r *Redis) ClusterForget(ctx context.Context, nodeID string) *redis.StatusCmd {
	return r.client.ClusterForget(ctx, nodeID)
}

func (r *Redis) ClusterReplicate(ctx context.Context, nodeID string) *redis.StatusCmd {
	return r.client.ClusterReplicate(ctx, nodeID)
}

func (r *Redis) ClusterResetSoft(ctx context.Context) *redis.StatusCmd {
	return r.client.ClusterResetSoft(ctx)
}

func (r *Redis) ClusterResetHard(ctx context.Context) *redis.StatusCmd {
	return r.client.ClusterResetHard(ctx)
}

func (r *Redis) ClusterInfo(ctx context.Context) *redis.StringCmd {
	return r.client.ClusterInfo(ctx)
}

func (r *Redis) ClusterKeySlot(ctx context.Context, key string) *redis.IntCmd {
	return r.client.ClusterKeySlot(ctx, key)
}

func (r *Redis) ClusterGetKeysInSlot(ctx context.Context, slot int, count int) *redis.StringSliceCmd {
	return r.client.ClusterGetKeysInSlot(ctx, slot, count)
}

func (r *Redis) ClusterCountFailureReports(ctx context.Context, nodeID string) *redis.IntCmd {
	return r.client.ClusterCountFailureReports(ctx, nodeID)
}

func (r *Redis) ClusterCountKeysInSlot(ctx context.Context, slot int) *redis.IntCmd {
	return r.client.ClusterCountKeysInSlot(ctx, slot)
}

func (r *Redis) ClusterDelSlots(ctx context.Context, slots ...int) *redis.StatusCmd {
	return r.client.ClusterDelSlots(ctx, slots...)
}

func (r *Redis) ClusterDelSlotsRange(ctx context.Context, min, max int) *redis.StatusCmd {
	return r.client.ClusterDelSlotsRange(ctx, min, max)
}

func (r *Redis) ClusterSaveConfig(ctx context.Context) *redis.StatusCmd {
	return r.client.ClusterSaveConfig(ctx)
}

func (r *Redis) ClusterSlaves(ctx context.Context, nodeID string) *redis.StringSliceCmd {
	return r.client.ClusterSlaves(ctx, nodeID)
}

func (r *Redis) ClusterFailover(ctx context.Context) *redis.StatusCmd {
	return r.client.ClusterFailover(ctx)
}

func (r *Redis) ClusterAddSlots(ctx context.Context, slots ...int) *redis.StatusCmd {
	return r.client.ClusterAddSlots(ctx, slots...)
}

func (r *Redis) ClusterAddSlotsRange(ctx context.Context, min, max int) *redis.StatusCmd {
	return r.client.ClusterAddSlotsRange(ctx, min, max)
}

func (r *Redis) GeoAdd(ctx context.Context, key string, geoLocation ...*redis.GeoLocation) *redis.IntCmd {
	return r.client.GeoAdd(ctx, key, geoLocation...)
}

func (r *Redis) GeoPos(ctx context.Context, key string, members ...string) *redis.GeoPosCmd {
	return r.client.GeoPos(ctx, key, members...)
}

func (r *Redis) GeoRadius(ctx context.Context, key string, longitude, latitude float64, query *redis.GeoRadiusQuery) *redis.GeoLocationCmd {
	return r.client.GeoRadius(ctx, key, longitude, latitude, query)
}

func (r *Redis) GeoRadiusByMember(ctx context.Context, key, member string, query *redis.GeoRadiusQuery) *redis.GeoLocationCmd {
	return r.client.GeoRadiusByMember(ctx, key, member, query)
}

func (r *Redis) GeoDist(ctx context.Context, key string, member1, member2, unit string) *redis.FloatCmd {
	return r.client.GeoDist(ctx, key, member1, member2, unit)
}

func (r *Redis) GeoHash(ctx context.Context, key string, members ...string) *redis.StringSliceCmd {
	return r.client.GeoHash(ctx, key, members...)
}

func (r *Redis) ReadOnly(ctx context.Context) *redis.StatusCmd {
	return r.client.ReadOnly(ctx)
}

func (r *Redis) ReadWrite(ctx context.Context) *redis.StatusCmd {
	return r.client.ReadWrite(ctx)
}

func (r *Redis) MemoryUsage(ctx context.Context, key string, samples ...int) *redis.IntCmd {
	return r.client.MemoryUsage(ctx, key, samples...)
}

// PoolStats returns connection pool stats.
func (r *Redis) PoolStats() *redis.PoolStats {
	return r.client.PoolStats()
}

func (r *Redis) Pipelined(ctx context.Context, fn func(redis.Pipeliner) error) ([]redis.Cmder, error) {
	return r.client.Pipelined(ctx, fn)
}

func (r *Redis) Pipeline() redis.Pipeliner {
	return r.client.Pipeline()
}

func (r *Redis) TxPipelined(ctx context.Context, fn func(redis.Pipeliner) error) ([]redis.Cmder, error) {
	return r.client.TxPipelined(ctx, fn)
}

// TxPipeline acts like Pipeline, but wraps queued commands with MULTI/EXEC.
func (r *Redis) TxPipeline() redis.Pipeliner {
	return r.client.TxPipeline()
}

// Subscribe subscribes the client to the specified channels.
// Channels can be omitted to create empty subscription.
// Note that this method does not wait on a response from Redis, so the
// subscription may not be active immediately. To force the connection to wait,
// you may call the Receive() method on the returned *PubSub like so:
//
//	sub := client.Subscribe(queryResp)
//	iface, err := sub.Receive()
//	if err != nil {
//	    // handle error
//	}
//
//	// Should be *Subscription, but others are possible if other actions have been
//	// taken on sub since it was created.
//	switch iface.(type) {
//	case *Subscription:
//	    // subscribe succeeded
//	case *Message:
//	    // received first message
//	case *Pong:
//	    // pong received
//	default:
//	    // handle error
//	}
//
//	ch := sub.Channel()
func (r *Redis) Subscribe(ctx context.Context, channels ...string) *redis.PubSub {
	return r.client.Subscribe(ctx, channels...)
}

// PSubscribe subscribes the client to the given patterns.
// Patterns can be omitted to create empty subscription.
func (r *Redis) PSubscribe(ctx context.Context, channels ...string) *redis.PubSub {
	return r.client.PSubscribe(ctx, channels...)
}

// ACLCat implements redis.Cmdable.
func (r *Redis) ACLCat(ctx context.Context) *redis.StringSliceCmd {
	return r.client.ACLCat(ctx)
}

// ACLCatArgs implements redis.Cmdable.
func (r *Redis) ACLCatArgs(ctx context.Context, options *redis.ACLCatArgs) *redis.StringSliceCmd {
	return r.client.ACLCatArgs(ctx, options)
}

// ACLDelUser implements redis.Cmdable.
func (r *Redis) ACLDelUser(ctx context.Context, username string) *redis.IntCmd {
	return r.client.ACLDelUser(ctx, username)
}

// ACLList implements redis.Cmdable.
func (r *Redis) ACLList(ctx context.Context) *redis.StringSliceCmd {
	return r.client.ACLList(ctx)
}

// ACLSetUser implements redis.Cmdable.
func (r *Redis) ACLSetUser(ctx context.Context, username string, rules ...string) *redis.StatusCmd {
	return r.client.ACLSetUser(ctx, username, rules...)
}

// BFReserveWithArgs implements redis.Cmdable.
func (r *Redis) BFReserveWithArgs(ctx context.Context, key string, options *redis.BFReserveOptions) *redis.StatusCmd {
	return r.client.BFReserveWithArgs(ctx, key, options)
}

// BitFieldRO implements redis.Cmdable.
func (r *Redis) BitFieldRO(ctx context.Context, key string, values ...interface{}) *redis.IntSliceCmd {
	return r.client.BitFieldRO(ctx, key, values...)
}

// BitOpAndOr implements redis.Cmdable.
func (r *Redis) BitOpAndOr(ctx context.Context, destKey string, keys ...string) *redis.IntCmd {
	return r.client.BitOpAndOr(ctx, destKey, keys...)
}

// BitOpDiff implements redis.Cmdable.
func (r *Redis) BitOpDiff(ctx context.Context, destKey string, keys ...string) *redis.IntCmd {
	return r.client.BitOpDiff(ctx, destKey, keys...)
}

// BitOpDiff1 implements redis.Cmdable.
func (r *Redis) BitOpDiff1(ctx context.Context, destKey string, keys ...string) *redis.IntCmd {
	return r.client.BitOpDiff1(ctx, destKey, keys...)
}

// BitOpOne implements redis.Cmdable.
func (r *Redis) BitOpOne(ctx context.Context, destKey string, keys ...string) *redis.IntCmd {
	return r.client.BitOpOne(ctx, destKey, keys...)
}

// CFReserveWithArgs implements redis.Cmdable.
func (r *Redis) CFReserveWithArgs(ctx context.Context, key string, options *redis.CFReserveOptions) *redis.StatusCmd {
	return r.client.CFReserveWithArgs(ctx, key, options)
}

// ClusterMyID implements redis.Cmdable.
func (r *Redis) ClusterMyID(ctx context.Context) *redis.StringCmd {
	return r.client.ClusterMyID(ctx)
}

// FTAggregate implements redis.Cmdable.
func (r *Redis) FTAggregate(ctx context.Context, index string, query string) *redis.MapStringInterfaceCmd {
	return r.client.FTAggregate(ctx, index, query)
}

// FTAggregateWithArgs implements redis.Cmdable.
func (r *Redis) FTAggregateWithArgs(ctx context.Context, index string, query string, options *redis.FTAggregateOptions) *redis.AggregateCmd {
	return r.client.FTAggregateWithArgs(ctx, index, query, options)
}

// FTAliasAdd implements redis.Cmdable.
func (r *Redis) FTAliasAdd(ctx context.Context, index string, alias string) *redis.StatusCmd {
	return r.client.FTAliasAdd(ctx, index, alias)
}

// FTAliasDel implements redis.Cmdable.
func (r *Redis) FTAliasDel(ctx context.Context, alias string) *redis.StatusCmd {
	return r.client.FTAliasDel(ctx, alias)
}

// FTAliasUpdate implements redis.Cmdable.
func (r *Redis) FTAliasUpdate(ctx context.Context, index string, alias string) *redis.StatusCmd {
	return r.client.FTAliasUpdate(ctx, index, alias)
}

// FTAlter implements redis.Cmdable.
func (r *Redis) FTAlter(ctx context.Context, index string, skipInitialScan bool, definition []interface{}) *redis.StatusCmd {
	return r.client.FTAlter(ctx, index, skipInitialScan, definition)
}

// FTConfigGet implements redis.Cmdable.
func (r *Redis) FTConfigGet(ctx context.Context, option string) *redis.MapMapStringInterfaceCmd {
	return r.client.FTConfigGet(ctx, option)
}

// FTConfigSet implements redis.Cmdable.
func (r *Redis) FTConfigSet(ctx context.Context, option string, value interface{}) *redis.StatusCmd {
	return r.client.FTConfigSet(ctx, option, value)
}

// FTCreate implements redis.Cmdable.
func (r *Redis) FTCreate(ctx context.Context, index string, options *redis.FTCreateOptions, schema ...*redis.FieldSchema) *redis.StatusCmd {
	return r.client.FTCreate(ctx, index, options, schema...)
}

// FTCursorDel implements redis.Cmdable.
func (r *Redis) FTCursorDel(ctx context.Context, index string, cursorId int) *redis.StatusCmd {
	return r.client.FTCursorDel(ctx, index, cursorId)
}

// FTCursorRead implements redis.Cmdable.
func (r *Redis) FTCursorRead(ctx context.Context, index string, cursorId int, count int) *redis.MapStringInterfaceCmd {
	return r.client.FTCursorRead(ctx, index, cursorId, count)
}

// FTDictAdd implements redis.Cmdable.
func (r *Redis) FTDictAdd(ctx context.Context, dict string, term ...interface{}) *redis.IntCmd {
	return r.client.FTDictAdd(ctx, dict, term...)
}

// FTDictDel implements redis.Cmdable.
func (r *Redis) FTDictDel(ctx context.Context, dict string, term ...interface{}) *redis.IntCmd {
	return r.client.FTDictDel(ctx, dict, term...)
}

// FTDictDump implements redis.Cmdable.
func (r *Redis) FTDictDump(ctx context.Context, dict string) *redis.StringSliceCmd {
	return r.client.FTDictDump(ctx, dict)
}

// FTDropIndex implements redis.Cmdable.
func (r *Redis) FTDropIndex(ctx context.Context, index string) *redis.StatusCmd {
	return r.client.FTDropIndex(ctx, index)
}

// FTDropIndexWithArgs implements redis.Cmdable.
func (r *Redis) FTDropIndexWithArgs(ctx context.Context, index string, options *redis.FTDropIndexOptions) *redis.StatusCmd {
	return r.client.FTDropIndexWithArgs(ctx, index, options)
}

// FTExplain implements redis.Cmdable.
func (r *Redis) FTExplain(ctx context.Context, index string, query string) *redis.StringCmd {
	return r.client.FTExplain(ctx, index, query)
}

// FTExplainWithArgs implements redis.Cmdable.
func (r *Redis) FTExplainWithArgs(ctx context.Context, index string, query string, options *redis.FTExplainOptions) *redis.StringCmd {
	return r.client.FTExplainWithArgs(ctx, index, query, options)
}

// FTInfo implements redis.Cmdable.
func (r *Redis) FTInfo(ctx context.Context, index string) *redis.FTInfoCmd {
	return r.client.FTInfo(ctx, index)
}

// FTSearch implements redis.Cmdable.
func (r *Redis) FTSearch(ctx context.Context, index string, query string) *redis.FTSearchCmd {
	return r.client.FTSearch(ctx, index, query)
}

// FTSearchWithArgs implements redis.Cmdable.
func (r *Redis) FTSearchWithArgs(ctx context.Context, index string, query string, options *redis.FTSearchOptions) *redis.FTSearchCmd {
	return r.client.FTSearchWithArgs(ctx, index, query, options)
}

// FTSpellCheck implements redis.Cmdable.
func (r *Redis) FTSpellCheck(ctx context.Context, index string, query string) *redis.FTSpellCheckCmd {
	return r.client.FTSpellCheck(ctx, index, query)
}

// FTSpellCheckWithArgs implements redis.Cmdable.
func (r *Redis) FTSpellCheckWithArgs(ctx context.Context, index string, query string, options *redis.FTSpellCheckOptions) *redis.FTSpellCheckCmd {
	return r.client.FTSpellCheckWithArgs(ctx, index, query, options)
}

// FTSynDump implements redis.Cmdable.
func (r *Redis) FTSynDump(ctx context.Context, index string) *redis.FTSynDumpCmd {
	return r.client.FTSynDump(ctx, index)
}

// FTSynUpdate implements redis.Cmdable.
func (r *Redis) FTSynUpdate(ctx context.Context, index string, synGroupId interface{}, terms []interface{}) *redis.StatusCmd {
	return r.client.FTSynUpdate(ctx, index, synGroupId, terms)
}

// FTSynUpdateWithArgs implements redis.Cmdable.
func (r *Redis) FTSynUpdateWithArgs(ctx context.Context, index string, synGroupId interface{}, options *redis.FTSynUpdateOptions, terms []interface{}) *redis.StatusCmd {
	return r.client.FTSynUpdateWithArgs(ctx, index, synGroupId, options, terms)
}

// FTTagVals implements redis.Cmdable.
func (r *Redis) FTTagVals(ctx context.Context, index string, field string) *redis.StringSliceCmd {
	return r.client.FTTagVals(ctx, index, field)
}

// FT_List implements redis.Cmdable.
func (r *Redis) FT_List(ctx context.Context) *redis.StringSliceCmd {
	return r.client.FT_List(ctx)
}

// HExpire implements redis.Cmdable.
func (r *Redis) HExpire(ctx context.Context, key string, expiration time.Duration, fields ...string) *redis.IntSliceCmd {
	panic("unimplemented")
}

// HExpireAt implements redis.Cmdable.
func (r *Redis) HExpireAt(ctx context.Context, key string, tm time.Time, fields ...string) *redis.IntSliceCmd {
	panic("unimplemented")
}

// HExpireAtWithArgs implements redis.Cmdable.
func (r *Redis) HExpireAtWithArgs(ctx context.Context, key string, tm time.Time, expirationArgs redis.HExpireArgs, fields ...string) *redis.IntSliceCmd {
	panic("unimplemented")
}

// HExpireTime implements redis.Cmdable.
func (r *Redis) HExpireTime(ctx context.Context, key string, fields ...string) *redis.IntSliceCmd {
	panic("unimplemented")
}

// HExpireWithArgs implements redis.Cmdable.
func (r *Redis) HExpireWithArgs(ctx context.Context, key string, expiration time.Duration, expirationArgs redis.HExpireArgs, fields ...string) *redis.IntSliceCmd {
	panic("unimplemented")
}

// HGetDel implements redis.Cmdable.
func (r *Redis) HGetDel(ctx context.Context, key string, fields ...string) *redis.StringSliceCmd {
	panic("unimplemented")
}

// HGetEX implements redis.Cmdable.
func (r *Redis) HGetEX(ctx context.Context, key string, fields ...string) *redis.StringSliceCmd {
	panic("unimplemented")
}

// HGetEXWithArgs implements redis.Cmdable.
func (r *Redis) HGetEXWithArgs(ctx context.Context, key string, options *redis.HGetEXOptions, fields ...string) *redis.StringSliceCmd {
	panic("unimplemented")
}

// HPExpire implements redis.Cmdable.
func (r *Redis) HPExpire(ctx context.Context, key string, expiration time.Duration, fields ...string) *redis.IntSliceCmd {
	panic("unimplemented")
}

// HPExpireAt implements redis.Cmdable.
func (r *Redis) HPExpireAt(ctx context.Context, key string, tm time.Time, fields ...string) *redis.IntSliceCmd {
	panic("unimplemented")
}

// HPExpireAtWithArgs implements redis.Cmdable.
func (r *Redis) HPExpireAtWithArgs(ctx context.Context, key string, tm time.Time, expirationArgs redis.HExpireArgs, fields ...string) *redis.IntSliceCmd {
	panic("unimplemented")
}

// HPExpireTime implements redis.Cmdable.
func (r *Redis) HPExpireTime(ctx context.Context, key string, fields ...string) *redis.IntSliceCmd {
	panic("unimplemented")
}

// HPExpireWithArgs implements redis.Cmdable.
func (r *Redis) HPExpireWithArgs(ctx context.Context, key string, expiration time.Duration, expirationArgs redis.HExpireArgs, fields ...string) *redis.IntSliceCmd {
	panic("unimplemented")
}

// HPTTL implements redis.Cmdable.
func (r *Redis) HPTTL(ctx context.Context, key string, fields ...string) *redis.IntSliceCmd {
	panic("unimplemented")
}

// HPersist implements redis.Cmdable.
func (r *Redis) HPersist(ctx context.Context, key string, fields ...string) *redis.IntSliceCmd {
	panic("unimplemented")
}

// HScanNoValues implements redis.Cmdable.
func (r *Redis) HScanNoValues(ctx context.Context, key string, cursor uint64, match string, count int64) *redis.ScanCmd {
	panic("unimplemented")
}

// HSetEX implements redis.Cmdable.
func (r *Redis) HSetEX(ctx context.Context, key string, fieldsAndValues ...string) *redis.IntCmd {
	panic("unimplemented")
}

// HSetEXWithArgs implements redis.Cmdable.
func (r *Redis) HSetEXWithArgs(ctx context.Context, key string, options *redis.HSetEXOptions, fieldsAndValues ...string) *redis.IntCmd {
	panic("unimplemented")
}

// HStrLen implements redis.Cmdable.
func (r *Redis) HStrLen(ctx context.Context, key string, field string) *redis.IntCmd {
	panic("unimplemented")
}

// HTTL implements redis.Cmdable.
func (r *Redis) HTTL(ctx context.Context, key string, fields ...string) *redis.IntSliceCmd {
	panic("unimplemented")
}

// JSONArrAppend implements redis.Cmdable.
func (r *Redis) JSONArrAppend(ctx context.Context, key string, path string, values ...interface{}) *redis.IntSliceCmd {
	panic("unimplemented")
}

// JSONArrIndex implements redis.Cmdable.
func (r *Redis) JSONArrIndex(ctx context.Context, key string, path string, value ...interface{}) *redis.IntSliceCmd {
	panic("unimplemented")
}

// JSONArrIndexWithArgs implements redis.Cmdable.
func (r *Redis) JSONArrIndexWithArgs(ctx context.Context, key string, path string, options *redis.JSONArrIndexArgs, value ...interface{}) *redis.IntSliceCmd {
	panic("unimplemented")
}

// JSONArrInsert implements redis.Cmdable.
func (r *Redis) JSONArrInsert(ctx context.Context, key string, path string, index int64, values ...interface{}) *redis.IntSliceCmd {
	panic("unimplemented")
}

// JSONArrLen implements redis.Cmdable.
func (r *Redis) JSONArrLen(ctx context.Context, key string, path string) *redis.IntSliceCmd {
	panic("unimplemented")
}

// JSONArrPop implements redis.Cmdable.
func (r *Redis) JSONArrPop(ctx context.Context, key string, path string, index int) *redis.StringSliceCmd {
	panic("unimplemented")
}

// JSONArrTrim implements redis.Cmdable.
func (r *Redis) JSONArrTrim(ctx context.Context, key string, path string) *redis.IntSliceCmd {
	panic("unimplemented")
}

// JSONArrTrimWithArgs implements redis.Cmdable.
func (r *Redis) JSONArrTrimWithArgs(ctx context.Context, key string, path string, options *redis.JSONArrTrimArgs) *redis.IntSliceCmd {
	panic("unimplemented")
}

// JSONClear implements redis.Cmdable.
func (r *Redis) JSONClear(ctx context.Context, key string, path string) *redis.IntCmd {
	panic("unimplemented")
}

// JSONDebugMemory implements redis.Cmdable.
func (r *Redis) JSONDebugMemory(ctx context.Context, key string, path string) *redis.IntCmd {
	panic("unimplemented")
}

// JSONDel implements redis.Cmdable.
func (r *Redis) JSONDel(ctx context.Context, key string, path string) *redis.IntCmd {
	panic("unimplemented")
}

// JSONForget implements redis.Cmdable.
func (r *Redis) JSONForget(ctx context.Context, key string, path string) *redis.IntCmd {
	panic("unimplemented")
}

// JSONGet implements redis.Cmdable.
func (r *Redis) JSONGet(ctx context.Context, key string, paths ...string) *redis.JSONCmd {
	panic("unimplemented")
}

// JSONGetWithArgs implements redis.Cmdable.
func (r *Redis) JSONGetWithArgs(ctx context.Context, key string, options *redis.JSONGetArgs, paths ...string) *redis.JSONCmd {
	panic("unimplemented")
}

// JSONMGet implements redis.Cmdable.
func (r *Redis) JSONMGet(ctx context.Context, path string, keys ...string) *redis.JSONSliceCmd {
	panic("unimplemented")
}

// JSONMSet implements redis.Cmdable.
func (r *Redis) JSONMSet(ctx context.Context, params ...interface{}) *redis.StatusCmd {
	panic("unimplemented")
}

// JSONMSetArgs implements redis.Cmdable.
func (r *Redis) JSONMSetArgs(ctx context.Context, docs []redis.JSONSetArgs) *redis.StatusCmd {
	panic("unimplemented")
}

// JSONMerge implements redis.Cmdable.
func (r *Redis) JSONMerge(ctx context.Context, key string, path string, value string) *redis.StatusCmd {
	panic("unimplemented")
}

// JSONNumIncrBy implements redis.Cmdable.
func (r *Redis) JSONNumIncrBy(ctx context.Context, key string, path string, value float64) *redis.JSONCmd {
	panic("unimplemented")
}

// JSONObjKeys implements redis.Cmdable.
func (r *Redis) JSONObjKeys(ctx context.Context, key string, path string) *redis.SliceCmd {
	panic("unimplemented")
}

// JSONObjLen implements redis.Cmdable.
func (r *Redis) JSONObjLen(ctx context.Context, key string, path string) *redis.IntPointerSliceCmd {
	panic("unimplemented")
}

// JSONSet implements redis.Cmdable.
func (r *Redis) JSONSet(ctx context.Context, key string, path string, value interface{}) *redis.StatusCmd {
	panic("unimplemented")
}

// JSONSetMode implements redis.Cmdable.
func (r *Redis) JSONSetMode(ctx context.Context, key string, path string, value interface{}, mode string) *redis.StatusCmd {
	panic("unimplemented")
}

// JSONStrAppend implements redis.Cmdable.
func (r *Redis) JSONStrAppend(ctx context.Context, key string, path string, value string) *redis.IntPointerSliceCmd {
	panic("unimplemented")
}

// JSONStrLen implements redis.Cmdable.
func (r *Redis) JSONStrLen(ctx context.Context, key string, path string) *redis.IntPointerSliceCmd {
	panic("unimplemented")
}

// JSONToggle implements redis.Cmdable.
func (r *Redis) JSONToggle(ctx context.Context, key string, path string) *redis.IntPointerSliceCmd {
	panic("unimplemented")
}

// JSONType implements redis.Cmdable.
func (r *Redis) JSONType(ctx context.Context, key string, path string) *redis.JSONSliceCmd {
	panic("unimplemented")
}

// ObjectFreq implements redis.Cmdable.
func (r *Redis) ObjectFreq(ctx context.Context, key string) *redis.IntCmd {
	panic("unimplemented")
}

// TSAdd implements redis.Cmdable.
func (r *Redis) TSAdd(ctx context.Context, key string, timestamp interface{}, value float64) *redis.IntCmd {
	panic("unimplemented")
}

// TSAddWithArgs implements redis.Cmdable.
func (r *Redis) TSAddWithArgs(ctx context.Context, key string, timestamp interface{}, value float64, options *redis.TSOptions) *redis.IntCmd {
	panic("unimplemented")
}

// TSAlter implements redis.Cmdable.
func (r *Redis) TSAlter(ctx context.Context, key string, options *redis.TSAlterOptions) *redis.StatusCmd {
	panic("unimplemented")
}

// TSCreate implements redis.Cmdable.
func (r *Redis) TSCreate(ctx context.Context, key string) *redis.StatusCmd {
	panic("unimplemented")
}

// TSCreateRule implements redis.Cmdable.
func (r *Redis) TSCreateRule(ctx context.Context, sourceKey string, destKey string, aggregator redis.Aggregator, bucketDuration int) *redis.StatusCmd {
	panic("unimplemented")
}

// TSCreateRuleWithArgs implements redis.Cmdable.
func (r *Redis) TSCreateRuleWithArgs(ctx context.Context, sourceKey string, destKey string, aggregator redis.Aggregator, bucketDuration int, options *redis.TSCreateRuleOptions) *redis.StatusCmd {
	panic("unimplemented")
}

// TSCreateWithArgs implements redis.Cmdable.
func (r *Redis) TSCreateWithArgs(ctx context.Context, key string, options *redis.TSOptions) *redis.StatusCmd {
	panic("unimplemented")
}

// TSDecrBy implements redis.Cmdable.
func (r *Redis) TSDecrBy(ctx context.Context, Key string, timestamp float64) *redis.IntCmd {
	panic("unimplemented")
}

// TSDecrByWithArgs implements redis.Cmdable.
func (r *Redis) TSDecrByWithArgs(ctx context.Context, key string, timestamp float64, options *redis.TSIncrDecrOptions) *redis.IntCmd {
	panic("unimplemented")
}

// TSDel implements redis.Cmdable.
func (r *Redis) TSDel(ctx context.Context, Key string, fromTimestamp int, toTimestamp int) *redis.IntCmd {
	panic("unimplemented")
}

// TSDeleteRule implements redis.Cmdable.
func (r *Redis) TSDeleteRule(ctx context.Context, sourceKey string, destKey string) *redis.StatusCmd {
	panic("unimplemented")
}

// TSGet implements redis.Cmdable.
func (r *Redis) TSGet(ctx context.Context, key string) *redis.TSTimestampValueCmd {
	panic("unimplemented")
}

// TSGetWithArgs implements redis.Cmdable.
func (r *Redis) TSGetWithArgs(ctx context.Context, key string, options *redis.TSGetOptions) *redis.TSTimestampValueCmd {
	panic("unimplemented")
}

// TSIncrBy implements redis.Cmdable.
func (r *Redis) TSIncrBy(ctx context.Context, Key string, timestamp float64) *redis.IntCmd {
	panic("unimplemented")
}

// TSIncrByWithArgs implements redis.Cmdable.
func (r *Redis) TSIncrByWithArgs(ctx context.Context, key string, timestamp float64, options *redis.TSIncrDecrOptions) *redis.IntCmd {
	panic("unimplemented")
}

// TSInfo implements redis.Cmdable.
func (r *Redis) TSInfo(ctx context.Context, key string) *redis.MapStringInterfaceCmd {
	panic("unimplemented")
}

// TSInfoWithArgs implements redis.Cmdable.
func (r *Redis) TSInfoWithArgs(ctx context.Context, key string, options *redis.TSInfoOptions) *redis.MapStringInterfaceCmd {
	panic("unimplemented")
}

// TSMAdd implements redis.Cmdable.
func (r *Redis) TSMAdd(ctx context.Context, ktvSlices [][]interface{}) *redis.IntSliceCmd {
	panic("unimplemented")
}

// TSMGet implements redis.Cmdable.
func (r *Redis) TSMGet(ctx context.Context, filters []string) *redis.MapStringSliceInterfaceCmd {
	panic("unimplemented")
}

// TSMGetWithArgs implements redis.Cmdable.
func (r *Redis) TSMGetWithArgs(ctx context.Context, filters []string, options *redis.TSMGetOptions) *redis.MapStringSliceInterfaceCmd {
	panic("unimplemented")
}

// TSMRange implements redis.Cmdable.
func (r *Redis) TSMRange(ctx context.Context, fromTimestamp int, toTimestamp int, filterExpr []string) *redis.MapStringSliceInterfaceCmd {
	panic("unimplemented")
}

// TSMRangeWithArgs implements redis.Cmdable.
func (r *Redis) TSMRangeWithArgs(ctx context.Context, fromTimestamp int, toTimestamp int, filterExpr []string, options *redis.TSMRangeOptions) *redis.MapStringSliceInterfaceCmd {
	panic("unimplemented")
}

// TSMRevRange implements redis.Cmdable.
func (r *Redis) TSMRevRange(ctx context.Context, fromTimestamp int, toTimestamp int, filterExpr []string) *redis.MapStringSliceInterfaceCmd {
	panic("unimplemented")
}

// TSMRevRangeWithArgs implements redis.Cmdable.
func (r *Redis) TSMRevRangeWithArgs(ctx context.Context, fromTimestamp int, toTimestamp int, filterExpr []string, options *redis.TSMRevRangeOptions) *redis.MapStringSliceInterfaceCmd {
	panic("unimplemented")
}

// TSQueryIndex implements redis.Cmdable.
func (r *Redis) TSQueryIndex(ctx context.Context, filterExpr []string) *redis.StringSliceCmd {
	panic("unimplemented")
}

// TSRange implements redis.Cmdable.
func (r *Redis) TSRange(ctx context.Context, key string, fromTimestamp int, toTimestamp int) *redis.TSTimestampValueSliceCmd {
	panic("unimplemented")
}

// TSRangeWithArgs implements redis.Cmdable.
func (r *Redis) TSRangeWithArgs(ctx context.Context, key string, fromTimestamp int, toTimestamp int, options *redis.TSRangeOptions) *redis.TSTimestampValueSliceCmd {
	panic("unimplemented")
}

// TSRevRange implements redis.Cmdable.
func (r *Redis) TSRevRange(ctx context.Context, key string, fromTimestamp int, toTimestamp int) *redis.TSTimestampValueSliceCmd {
	panic("unimplemented")
}

// TSRevRangeWithArgs implements redis.Cmdable.
func (r *Redis) TSRevRangeWithArgs(ctx context.Context, key string, fromTimestamp int, toTimestamp int, options *redis.TSRevRangeOptions) *redis.TSTimestampValueSliceCmd {
	panic("unimplemented")
}

// VAdd implements redis.Cmdable.
func (r *Redis) VAdd(ctx context.Context, key string, element string, val redis.Vector) *redis.BoolCmd {
	panic("unimplemented")
}

// VAddWithArgs implements redis.Cmdable.
func (r *Redis) VAddWithArgs(ctx context.Context, key string, element string, val redis.Vector, addArgs *redis.VAddArgs) *redis.BoolCmd {
	panic("unimplemented")
}

// VCard implements redis.Cmdable.
func (r *Redis) VCard(ctx context.Context, key string) *redis.IntCmd {
	panic("unimplemented")
}

// VClearAttributes implements redis.Cmdable.
func (r *Redis) VClearAttributes(ctx context.Context, key string, element string) *redis.BoolCmd {
	panic("unimplemented")
}

// VDim implements redis.Cmdable.
func (r *Redis) VDim(ctx context.Context, key string) *redis.IntCmd {
	panic("unimplemented")
}

// VEmb implements redis.Cmdable.
func (r *Redis) VEmb(ctx context.Context, key string, element string, raw bool) *redis.SliceCmd {
	panic("unimplemented")
}

// VGetAttr implements redis.Cmdable.
func (r *Redis) VGetAttr(ctx context.Context, key string, element string) *redis.StringCmd {
	panic("unimplemented")
}

// VInfo implements redis.Cmdable.
func (r *Redis) VInfo(ctx context.Context, key string) *redis.MapStringInterfaceCmd {
	panic("unimplemented")
}

// VLinks implements redis.Cmdable.
func (r *Redis) VLinks(ctx context.Context, key string, element string) *redis.StringSliceCmd {
	panic("unimplemented")
}

// VLinksWithScores implements redis.Cmdable.
func (r *Redis) VLinksWithScores(ctx context.Context, key string, element string) *redis.VectorScoreSliceCmd {
	panic("unimplemented")
}

// VRandMember implements redis.Cmdable.
func (r *Redis) VRandMember(ctx context.Context, key string) *redis.StringCmd {
	panic("unimplemented")
}

// VRandMemberCount implements redis.Cmdable.
func (r *Redis) VRandMemberCount(ctx context.Context, key string, count int) *redis.StringSliceCmd {
	panic("unimplemented")
}

// VRem implements redis.Cmdable.
func (r *Redis) VRem(ctx context.Context, key string, element string) *redis.BoolCmd {
	panic("unimplemented")
}

// VSetAttr implements redis.Cmdable.
func (r *Redis) VSetAttr(ctx context.Context, key string, element string, attr interface{}) *redis.BoolCmd {
	panic("unimplemented")
}

// VSim implements redis.Cmdable.
func (r *Redis) VSim(ctx context.Context, key string, val redis.Vector) *redis.StringSliceCmd {
	panic("unimplemented")
}

// VSimWithArgs implements redis.Cmdable.
func (r *Redis) VSimWithArgs(ctx context.Context, key string, val redis.Vector, args *redis.VSimArgs) *redis.StringSliceCmd {
	panic("unimplemented")
}

// VSimWithArgsWithScores implements redis.Cmdable.
func (r *Redis) VSimWithArgsWithScores(ctx context.Context, key string, val redis.Vector, args *redis.VSimArgs) *redis.VectorScoreSliceCmd {
	panic("unimplemented")
}

// VSimWithScores implements redis.Cmdable.
func (r *Redis) VSimWithScores(ctx context.Context, key string, val redis.Vector) *redis.VectorScoreSliceCmd {
	panic("unimplemented")
}

// XAckDel implements redis.Cmdable.
func (r *Redis) XAckDel(ctx context.Context, stream string, group string, mode string, ids ...string) *redis.SliceCmd {
	panic("unimplemented")
}

// XDelEx implements redis.Cmdable.
func (r *Redis) XDelEx(ctx context.Context, stream string, mode string, ids ...string) *redis.SliceCmd {
	panic("unimplemented")
}

// XTrimMaxLenApproxMode implements redis.Cmdable.
func (r *Redis) XTrimMaxLenApproxMode(ctx context.Context, key string, maxLen int64, limit int64, mode string) *redis.IntCmd {
	panic("unimplemented")
}

// XTrimMaxLenMode implements redis.Cmdable.
func (r *Redis) XTrimMaxLenMode(ctx context.Context, key string, maxLen int64, mode string) *redis.IntCmd {
	panic("unimplemented")
}

// XTrimMinIDApproxMode implements redis.Cmdable.
func (r *Redis) XTrimMinIDApproxMode(ctx context.Context, key string, minID string, limit int64, mode string) *redis.IntCmd {
	panic("unimplemented")
}

// XTrimMinIDMode implements redis.Cmdable.
func (r *Redis) XTrimMinIDMode(ctx context.Context, key string, minID string, mode string) *redis.IntCmd {
	return r.client.XTrimMinIDMode(ctx, key, minID, mode)
}
