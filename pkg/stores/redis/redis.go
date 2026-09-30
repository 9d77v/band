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

// ACLUsers implements [redis.Cmdable].
func (r *Redis) ACLUsers(ctx context.Context) *redis.StringSliceCmd {
	return r.client.ACLUsers(ctx)
}

// ACLWhoAmI implements [redis.Cmdable].
func (r *Redis) ACLWhoAmI(ctx context.Context) *redis.StringCmd {
	return r.client.ACLWhoAmI(ctx)
}

// ARCount implements [redis.Cmdable].
func (r *Redis) ARCount(ctx context.Context, key string) *redis.UintCmd {
	return r.client.ARCount(ctx, key)
}

// ARDel implements [redis.Cmdable].
func (r *Redis) ARDel(ctx context.Context, key string, indexes ...uint64) *redis.IntCmd {
	return r.client.ARDel(ctx, key, indexes...)
}

// ARDelRange implements [redis.Cmdable].
func (r *Redis) ARDelRange(ctx context.Context, key string, ranges ...redis.ARRange) *redis.UintCmd {
	return r.client.ARDelRange(ctx, key, ranges...)
}

// ARGet implements [redis.Cmdable].
func (r *Redis) ARGet(ctx context.Context, key string, index uint64) *redis.StringCmd {
	return r.client.ARGet(ctx, key, index)
}

// ARGetRange implements [redis.Cmdable].
func (r *Redis) ARGetRange(ctx context.Context, key string, start uint64, end uint64) *redis.SliceCmd {
	return r.client.ARGetRange(ctx, key, start, end)
}

// ARGrep implements [redis.Cmdable].
func (r *Redis) ARGrep(ctx context.Context, key string, start string, end string, args *redis.ARGrepArgs) *redis.UintSliceCmd {
	return r.client.ARGrep(ctx, key, start, end, args)
}

// ARGrepWithValues implements [redis.Cmdable].
func (r *Redis) ARGrepWithValues(ctx context.Context, key string, start string, end string, args *redis.ARGrepArgs) *redis.AREntrySliceCmd {
	return r.client.ARGrepWithValues(ctx, key, start, end, args)
}

// ARInfo implements [redis.Cmdable].
func (r *Redis) ARInfo(ctx context.Context, key string) *redis.MapStringInterfaceCmd {
	return r.client.ARInfo(ctx, key)
}

// ARInfoFull implements [redis.Cmdable].
func (r *Redis) ARInfoFull(ctx context.Context, key string) *redis.MapStringInterfaceCmd {
	return r.client.ARInfoFull(ctx, key)
}

// ARInsert implements [redis.Cmdable].
func (r *Redis) ARInsert(ctx context.Context, key string, values ...string) *redis.UintCmd {
	return r.client.ARInsert(ctx, key, values...)
}

// ARLastItems implements [redis.Cmdable].
func (r *Redis) ARLastItems(ctx context.Context, key string, count uint64, rev bool) *redis.SliceCmd {
	return r.client.ARLastItems(ctx, key, count, rev)
}

// ARLen implements [redis.Cmdable].
func (r *Redis) ARLen(ctx context.Context, key string) *redis.UintCmd {
	return r.client.ARLen(ctx, key)
}

// ARMGet implements [redis.Cmdable].
func (r *Redis) ARMGet(ctx context.Context, key string, indexes ...uint64) *redis.SliceCmd {
	return r.client.ARMGet(ctx, key, indexes...)
}

// ARMSet implements [redis.Cmdable].
func (r *Redis) ARMSet(ctx context.Context, key string, members ...redis.AREntry) *redis.IntCmd {
	return r.client.ARMSet(ctx, key, members...)
}

// ARNext implements [redis.Cmdable].
func (r *Redis) ARNext(ctx context.Context, key string) *redis.UintCmd {
	return r.client.ARNext(ctx, key)
}

// AROpAnd implements [redis.Cmdable].
func (r *Redis) AROpAnd(ctx context.Context, key string, start uint64, end uint64) *redis.IntCmd {
	return r.client.AROpAnd(ctx, key, start, end)
}

// AROpMatch implements [redis.Cmdable].
func (r *Redis) AROpMatch(ctx context.Context, key string, start uint64, end uint64, value string) *redis.IntCmd {
	return r.client.AROpMatch(ctx, key, start, end, value)
}

// AROpMax implements [redis.Cmdable].
func (r *Redis) AROpMax(ctx context.Context, key string, start uint64, end uint64) *redis.StringCmd {
	return r.client.AROpMax(ctx, key, start, end)
}

// AROpMin implements [redis.Cmdable].
func (r *Redis) AROpMin(ctx context.Context, key string, start uint64, end uint64) *redis.StringCmd {
	return r.client.AROpMin(ctx, key, start, end)
}

// AROpOr implements [redis.Cmdable].
func (r *Redis) AROpOr(ctx context.Context, key string, start uint64, end uint64) *redis.IntCmd {
	return r.client.AROpOr(ctx, key, start, end)
}

// AROpSum implements [redis.Cmdable].
func (r *Redis) AROpSum(ctx context.Context, key string, start uint64, end uint64) *redis.StringCmd {
	return r.client.AROpSum(ctx, key, start, end)
}

// AROpUsed implements [redis.Cmdable].
func (r *Redis) AROpUsed(ctx context.Context, key string, start uint64, end uint64) *redis.IntCmd {
	return r.client.AROpUsed(ctx, key, start, end)
}

// AROpXor implements [redis.Cmdable].
func (r *Redis) AROpXor(ctx context.Context, key string, start uint64, end uint64) *redis.IntCmd {
	return r.client.AROpXor(ctx, key, start, end)
}

// ARRing implements [redis.Cmdable].
func (r *Redis) ARRing(ctx context.Context, key string, size uint64, values ...string) *redis.UintCmd {
	return r.client.ARRing(ctx, key, size, values...)
}

// ARScan implements [redis.Cmdable].
func (r *Redis) ARScan(ctx context.Context, key string, start uint64, end uint64, args *redis.ARScanArgs) *redis.AREntrySliceCmd {
	return r.client.ARScan(ctx, key, start, end, args)
}

// ARSeek implements [redis.Cmdable].
func (r *Redis) ARSeek(ctx context.Context, key string, index uint64) *redis.IntCmd {
	return r.client.ARSeek(ctx, key, index)
}

// ARSet implements [redis.Cmdable].
func (r *Redis) ARSet(ctx context.Context, key string, index uint64, values ...string) *redis.IntCmd {
	return r.client.ARSet(ctx, key, index, values...)
}

// BLMoveM implements [redis.Cmdable].
func (r *Redis) BLMoveM(ctx context.Context, source string, destination string, srcpos string, destpos string, timeout time.Duration, args redis.LMoveMArgs) *redis.StringSliceCmd {
	return r.client.BLMoveM(ctx, source, destination, srcpos, destpos, timeout, args)
}

// ClientMaintNotifications implements [redis.Cmdable].
func (r *Redis) ClientMaintNotifications(ctx context.Context, enabled bool, endpointType string) *redis.StatusCmd {
	return r.client.ClientMaintNotifications(ctx, enabled, endpointType)
}

// ClientTracking implements [redis.Cmdable].
func (r *Redis) ClientTracking(ctx context.Context, on bool, opt *redis.ClientTrackingOptions) *redis.StatusCmd {
	return r.client.ClientTracking(ctx, on, opt)
}

// ClientTrackingOff implements [redis.Cmdable].
func (r *Redis) ClientTrackingOff(ctx context.Context) *redis.StatusCmd {
	return r.client.ClientTrackingOff(ctx)
}

// ClientTrackingOn implements [redis.Cmdable].
func (r *Redis) ClientTrackingOn(ctx context.Context, opt *redis.ClientTrackingOptions) *redis.StatusCmd {
	return r.client.ClientTrackingOn(ctx, opt)
}

// DelExArgs implements [redis.Cmdable].
func (r *Redis) DelExArgs(ctx context.Context, key string, a redis.DelExArgs) *redis.IntCmd {
	return r.client.DelExArgs(ctx, key, a)
}

// Digest implements [redis.Cmdable].
func (r *Redis) Digest(ctx context.Context, key string) *redis.DigestCmd {
	return r.client.Digest(ctx, key)
}

// FTAliasList implements [redis.Cmdable].
func (r *Redis) FTAliasList(ctx context.Context, index string) *redis.StringSliceCmd {
	return r.client.FTAliasList(ctx, index)
}

// FTHybrid implements [redis.Cmdable].
func (r *Redis) FTHybrid(ctx context.Context, index string, searchExpr string, vectorField string, vectorData redis.Vector) *redis.FTHybridCmd {
	return r.client.FTHybrid(ctx, index, searchExpr, vectorField, vectorData)
}

// FTHybridWithArgs implements [redis.Cmdable].
func (r *Redis) FTHybridWithArgs(ctx context.Context, index string, options *redis.FTHybridOptions) *redis.FTHybridCmd {
	return r.client.FTHybridWithArgs(ctx, index, options)
}

// GetToBuffer implements [redis.Cmdable].
func (r *Redis) GetToBuffer(ctx context.Context, key string, buf []byte) *redis.ZeroCopyStringCmd {
	return r.client.GetToBuffer(ctx, key, buf)
}

// HImportDiscard implements [redis.Cmdable].
func (r *Redis) HImportDiscard(ctx context.Context, fieldsetName string) *redis.IntCmd {
	return r.client.HImportDiscard(ctx, fieldsetName)
}

// HImportDiscardAll implements [redis.Cmdable].
func (r *Redis) HImportDiscardAll(ctx context.Context) *redis.IntCmd {
	return r.client.HImportDiscardAll(ctx)
}

// HImportPrepare implements [redis.Cmdable].
func (r *Redis) HImportPrepare(ctx context.Context, fieldsetName string, fields ...string) *redis.StatusCmd {
	return r.client.HImportPrepare(ctx, fieldsetName, fields...)
}

// HImportSet implements [redis.Cmdable].
func (r *Redis) HImportSet(ctx context.Context, key string, fieldsetName string, values ...interface{}) *redis.StatusCmd {
	return r.client.HImportSet(ctx, key, fieldsetName, values...)
}

// IncrEXFloat implements [redis.Cmdable].
func (r *Redis) IncrEXFloat(ctx context.Context, key string, args redis.IncrEXFloatArgs) *redis.IncrEXFloatCmd {
	return r.client.IncrEXFloat(ctx, key, args)
}

// IncrEXInt implements [redis.Cmdable].
func (r *Redis) IncrEXInt(ctx context.Context, key string, args redis.IncrEXIntArgs) *redis.IncrEXIntCmd {
	return r.client.IncrEXInt(ctx, key, args)
}

// InfoMap implements [redis.Cmdable].
func (r *Redis) InfoMap(ctx context.Context, section ...string) *redis.InfoCmd {
	return r.client.InfoMap(ctx, section...)
}

// JSONSetWithArgs implements [redis.Cmdable].
func (r *Redis) JSONSetWithArgs(ctx context.Context, key string, path string, value interface{}, options *redis.JSONSetArgsOptions) *redis.StatusCmd {
	return r.client.JSONSetWithArgs(ctx, key, path, value, options)
}

// LMoveM implements [redis.Cmdable].
func (r *Redis) LMoveM(ctx context.Context, source string, destination string, srcpos string, destpos string, args redis.LMoveMArgs) *redis.StringSliceCmd {
	return r.client.LMoveM(ctx, source, destination, srcpos, destpos, args)
}

// Latency implements [redis.Cmdable].
func (r *Redis) Latency(ctx context.Context) *redis.LatencyCmd {
	return r.client.Latency(ctx)
}

// LatencyReset implements [redis.Cmdable].
func (r *Redis) LatencyReset(ctx context.Context, events ...interface{}) *redis.StatusCmd {
	return r.client.LatencyReset(ctx, events...)
}

// MSetEX implements [redis.Cmdable].
func (r *Redis) MSetEX(ctx context.Context, args redis.MSetEXArgs, values ...interface{}) *redis.IntCmd {
	return r.client.MSetEX(ctx, args, values...)
}

// ReplicaOf implements [redis.Cmdable].
func (r *Redis) ReplicaOf(ctx context.Context, host string, port string) *redis.StatusCmd {
	return r.client.ReplicaOf(ctx, host, port)
}

// SDiffCard implements [redis.Cmdable].
func (r *Redis) SDiffCard(ctx context.Context, opts *redis.SDiffCardOptions, keys ...string) *redis.IntCmd {
	return r.client.SDiffCard(ctx, opts, keys...)
}

// SUnionCard implements [redis.Cmdable].
func (r *Redis) SUnionCard(ctx context.Context, opts *redis.SUnionCardOptions, keys ...string) *redis.IntCmd {
	return r.client.SUnionCard(ctx, opts, keys...)
}

// SetFromBuffer implements [redis.Cmdable].
func (r *Redis) SetFromBuffer(ctx context.Context, key string, buf []byte) *redis.StatusCmd {
	return r.client.SetFromBuffer(ctx, key, buf)
}

// SetIFDEQ implements [redis.Cmdable].
func (r *Redis) SetIFDEQ(ctx context.Context, key string, value interface{}, matchDigest uint64, expiration time.Duration) *redis.StatusCmd {
	return r.client.SetIFDEQ(ctx, key, value, matchDigest, expiration)
}

// SetIFDEQGet implements [redis.Cmdable].
func (r *Redis) SetIFDEQGet(ctx context.Context, key string, value interface{}, matchDigest uint64, expiration time.Duration) *redis.StringCmd {
	return r.client.SetIFDEQGet(ctx, key, value, matchDigest, expiration)
}

// SetIFDNE implements [redis.Cmdable].
func (r *Redis) SetIFDNE(ctx context.Context, key string, value interface{}, matchDigest uint64, expiration time.Duration) *redis.StatusCmd {
	return r.client.SetIFDNE(ctx, key, value, matchDigest, expiration)
}

// SetIFDNEGet implements [redis.Cmdable].
func (r *Redis) SetIFDNEGet(ctx context.Context, key string, value interface{}, matchDigest uint64, expiration time.Duration) *redis.StringCmd {
	return r.client.SetIFDNEGet(ctx, key, value, matchDigest, expiration)
}

// SetIFEQ implements [redis.Cmdable].
func (r *Redis) SetIFEQ(ctx context.Context, key string, value interface{}, matchValue interface{}, expiration time.Duration) *redis.StatusCmd {
	return r.client.SetIFEQ(ctx, key, value, matchValue, expiration)
}

// SetIFEQGet implements [redis.Cmdable].
func (r *Redis) SetIFEQGet(ctx context.Context, key string, value interface{}, matchValue interface{}, expiration time.Duration) *redis.StringCmd {
	return r.client.SetIFEQGet(ctx, key, value, matchValue, expiration)
}

// SetIFNE implements [redis.Cmdable].
func (r *Redis) SetIFNE(ctx context.Context, key string, value interface{}, matchValue interface{}, expiration time.Duration) *redis.StatusCmd {
	return r.client.SetIFNE(ctx, key, value, matchValue, expiration)
}

// SetIFNEGet implements [redis.Cmdable].
func (r *Redis) SetIFNEGet(ctx context.Context, key string, value interface{}, matchValue interface{}, expiration time.Duration) *redis.StringCmd {
	return r.client.SetIFNEGet(ctx, key, value, matchValue, expiration)
}

// SlowLogLen implements [redis.Cmdable].
func (r *Redis) SlowLogLen(ctx context.Context) *redis.IntCmd {
	return r.client.SlowLogLen(ctx)
}

// SlowLogReset implements [redis.Cmdable].
func (r *Redis) SlowLogReset(ctx context.Context) *redis.StatusCmd {
	return r.client.SlowLogReset(ctx)
}

// TSNRange implements [redis.Cmdable].
func (r *Redis) TSNRange(ctx context.Context, keys []string, fromTimestamp interface{}, toTimestamp interface{}) *redis.TSNRangePivotRowSliceCmd {
	return r.client.TSNRange(ctx, keys, fromTimestamp, toTimestamp)
}

// TSNRangeWithArgs implements [redis.Cmdable].
func (r *Redis) TSNRangeWithArgs(ctx context.Context, keys []string, fromTimestamp interface{}, toTimestamp interface{}, options *redis.TSNRangeOptions) *redis.TSNRangePivotRowSliceCmd {
	return r.client.TSNRangeWithArgs(ctx, keys, fromTimestamp, toTimestamp, options)
}

// TSNRevRange implements [redis.Cmdable].
func (r *Redis) TSNRevRange(ctx context.Context, keys []string, fromTimestamp interface{}, toTimestamp interface{}) *redis.TSNRangePivotRowSliceCmd {
	return r.client.TSNRevRange(ctx, keys, fromTimestamp, toTimestamp)
}

// TSNRevRangeWithArgs implements [redis.Cmdable].
func (r *Redis) TSNRevRangeWithArgs(ctx context.Context, keys []string, fromTimestamp interface{}, toTimestamp interface{}, options *redis.TSNRevRangeOptions) *redis.TSNRangePivotRowSliceCmd {
	return r.client.TSNRevRangeWithArgs(ctx, keys, fromTimestamp, toTimestamp, options)
}

// TSQueryLabelValues implements [redis.Cmdable].
func (r *Redis) TSQueryLabelValues(ctx context.Context, label string, filterExpr []string) *redis.StringSliceCmd {
	return r.client.TSQueryLabelValues(ctx, label, filterExpr)
}

// TSQueryLabels implements [redis.Cmdable].
func (r *Redis) TSQueryLabels(ctx context.Context, filterExpr []string) *redis.StringSliceCmd {
	return r.client.TSQueryLabels(ctx, filterExpr)
}

// TSRead implements [redis.Cmdable].
func (r *Redis) TSRead(ctx context.Context, key string, timestamp interface{}) *redis.TSTimestampValueSliceCmd {
	return r.client.TSRead(ctx, key, timestamp)
}

// TSReadWithArgs implements [redis.Cmdable].
func (r *Redis) TSReadWithArgs(ctx context.Context, key string, timestamp interface{}, options *redis.TSReadOptions) *redis.TSTimestampValueSliceCmd {
	return r.client.TSReadWithArgs(ctx, key, timestamp, options)
}

// VIsMember implements [redis.Cmdable].
func (r *Redis) VIsMember(ctx context.Context, key string, element string) *redis.BoolCmd {
	return r.client.VIsMember(ctx, key, element)
}

// VRange implements [redis.Cmdable].
func (r *Redis) VRange(ctx context.Context, key string, start string, end string, count int64) *redis.StringSliceCmd {
	return r.client.VRange(ctx, key, start, end, count)
}

// VSimWithArgsWithAttribs implements [redis.Cmdable].
func (r *Redis) VSimWithArgsWithAttribs(ctx context.Context, key string, val redis.Vector, args *redis.VSimArgs) *redis.VectorAttribSliceCmd {
	return r.client.VSimWithArgsWithAttribs(ctx, key, val, args)
}

// VSimWithArgsWithScoresWithAttribs implements [redis.Cmdable].
func (r *Redis) VSimWithArgsWithScoresWithAttribs(ctx context.Context, key string, val redis.Vector, args *redis.VSimArgs) *redis.VectorScoreAttribSliceCmd {
	return r.client.VSimWithArgsWithScoresWithAttribs(ctx, key, val, args)
}

// XAutoClaimWithDeleted implements [redis.Cmdable].
func (r *Redis) XAutoClaimWithDeleted(ctx context.Context, a *redis.XAutoClaimArgs) *redis.XAutoClaimWithDeletedCmd {
	return r.client.XAutoClaimWithDeleted(ctx, a)
}

// XCfgSet implements [redis.Cmdable].
func (r *Redis) XCfgSet(ctx context.Context, a *redis.XCfgSetArgs) *redis.StatusCmd {
	return r.client.XCfgSet(ctx, a)
}

// XNack implements [redis.Cmdable].
func (r *Redis) XNack(ctx context.Context, a *redis.XNackArgs) *redis.IntCmd {
	return r.client.XNack(ctx, a)
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
func (r *Redis) BFAdd(ctx context.Context, key string, element any) *redis.BoolCmd {
	return r.client.BFAdd(ctx, key, element)
}

// BFCard implements redis.Cmdable.
func (r *Redis) BFCard(ctx context.Context, key string) *redis.IntCmd {
	return r.client.BFCard(ctx, key)
}

// BFExists implements redis.Cmdable.
func (r *Redis) BFExists(ctx context.Context, key string, element any) *redis.BoolCmd {
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
func (r *Redis) BFInsert(ctx context.Context, key string, options *redis.BFInsertOptions, elements ...any) *redis.BoolSliceCmd {
	return r.client.BFInsert(ctx, key, options, elements...)
}

// BFLoadChunk implements redis.Cmdable.
func (r *Redis) BFLoadChunk(ctx context.Context, key string, iterator int64, data any) *redis.StatusCmd {
	return r.client.BFLoadChunk(ctx, key, iterator, data)
}

// BFMAdd implements redis.Cmdable.
func (r *Redis) BFMAdd(ctx context.Context, key string, elements ...any) *redis.BoolSliceCmd {
	return r.client.BFMAdd(ctx, key, elements...)
}

// BFMExists implements redis.Cmdable.
func (r *Redis) BFMExists(ctx context.Context, key string, elements ...any) *redis.BoolSliceCmd {
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
func (r *Redis) CFAdd(ctx context.Context, key string, element any) *redis.BoolCmd {
	return r.client.CFAdd(ctx, key, element)
}

// CFAddNX implements redis.Cmdable.
func (r *Redis) CFAddNX(ctx context.Context, key string, element any) *redis.BoolCmd {
	return r.client.CFAddNX(ctx, key, element)
}

// CFCount implements redis.Cmdable.
func (r *Redis) CFCount(ctx context.Context, key string, element any) *redis.IntCmd {
	return r.client.CFCount(ctx, key, element)
}

// CFDel implements redis.Cmdable.
func (r *Redis) CFDel(ctx context.Context, key string, element any) *redis.BoolCmd {
	return r.client.CFDel(ctx, key, element)
}

// CFExists implements redis.Cmdable.
func (r *Redis) CFExists(ctx context.Context, key string, element any) *redis.BoolCmd {
	return r.client.CFExists(ctx, key, element)
}

// CFInfo implements redis.Cmdable.
func (r *Redis) CFInfo(ctx context.Context, key string) *redis.CFInfoCmd {
	return r.client.CFInfo(ctx, key)
}

// CFInsert implements redis.Cmdable.
func (r *Redis) CFInsert(ctx context.Context, key string, options *redis.CFInsertOptions, elements ...any) *redis.BoolSliceCmd {
	return r.client.CFInsert(ctx, key, options, elements...)
}

// CFInsertNX implements redis.Cmdable.
func (r *Redis) CFInsertNX(ctx context.Context, key string, options *redis.CFInsertOptions, elements ...any) *redis.IntSliceCmd {
	return r.client.CFInsertNX(ctx, key, options, elements...)
}

// CFLoadChunk implements redis.Cmdable.
func (r *Redis) CFLoadChunk(ctx context.Context, key string, iterator int64, data any) *redis.StatusCmd {
	return r.client.CFLoadChunk(ctx, key, iterator, data)
}

// CFMExists implements redis.Cmdable.
func (r *Redis) CFMExists(ctx context.Context, key string, elements ...any) *redis.BoolSliceCmd {
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
func (r *Redis) CMSIncrBy(ctx context.Context, key string, elements ...any) *redis.IntSliceCmd {
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
func (r *Redis) CMSQuery(ctx context.Context, key string, elements ...any) *redis.IntSliceCmd {
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
func (r *Redis) TopKAdd(ctx context.Context, key string, elements ...any) *redis.StringSliceCmd {
	return r.client.TopKAdd(ctx, key, elements...)
}

// TopKCount implements redis.Cmdable.
func (r *Redis) TopKCount(ctx context.Context, key string, elements ...any) *redis.IntSliceCmd {
	return r.client.TopKCount(ctx, key, elements...)
}

// TopKIncrBy implements redis.Cmdable.
func (r *Redis) TopKIncrBy(ctx context.Context, key string, elements ...any) *redis.StringSliceCmd {
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
func (r *Redis) TopKQuery(ctx context.Context, key string, elements ...any) *redis.BoolSliceCmd {
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
func (r *Redis) ACLDryRun(ctx context.Context, username string, command ...any) *redis.StringCmd {
	return r.client.ACLDryRun(ctx, username, command...)
}

// ACLGenPass implements redis.Cmdable
func (r *Redis) ACLGenPass(ctx context.Context, bit int) *redis.StringCmd {
	return r.client.ACLGenPass(ctx, bit)
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
func (r *Redis) CommandGetKeys(ctx context.Context, commands ...any) *redis.StringSliceCmd {
	return r.client.CommandGetKeys(ctx, commands...)
}

// CommandGetKeysAndFlags implements redis.Cmdable
func (r *Redis) CommandGetKeysAndFlags(ctx context.Context, commands ...any) *redis.KeyFlagsCmd {
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
func (r *Redis) FCall(ctx context.Context, function string, keys []string, args ...any) *redis.Cmd {
	return r.client.FCall(ctx, function, keys, args...)
}

// FCallRO implements redis.Cmdable
func (r *Redis) FCallRO(ctx context.Context, function string, keys []string, args ...any) *redis.Cmd {
	return r.client.FCallRO(ctx, function, keys, args...)
}

// FCallRo implements redis.Cmdable
func (r *Redis) FCallRo(ctx context.Context, function string, keys []string, args ...any) *redis.Cmd {
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
func (r *Redis) EvalRO(ctx context.Context, script string, keys []string, args ...any) *redis.Cmd {
	return r.client.EvalRO(ctx, script, keys, args...)
}

// EvalShaRO implements redis.Cmdable
func (r *Redis) EvalShaRO(ctx context.Context, sha1 string, keys []string, args ...any) *redis.Cmd {
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
func (r *Redis) SPublish(ctx context.Context, channel string, message any) *redis.IntCmd {
	return r.client.SPublish(ctx, channel, message)
}

// SetEx implements redis.Cmdable
func (r *Redis) SetEx(ctx context.Context, key string, value any, expiration time.Duration) *redis.StatusCmd {
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
	return r.client.BLMove(ctx, source, destination, srcpos, destpos, timeout)
}

// BitField implements redis.Cmdable
func (r *Redis) BitField(ctx context.Context, key string, args ...any) *redis.IntSliceCmd {
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
	return r.client.GeoRadiusStore(ctx, key, longitude, latitude, query)
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
func (r *Redis) SMIsMember(ctx context.Context, key string, members ...any) *redis.BoolSliceCmd {
	return r.client.SMIsMember(ctx, key, members...)
}

// ScanType implements redis.Cmdable
func (r *Redis) ScanType(ctx context.Context, cursor uint64, match string, count int64, keyType string) *redis.ScanCmd {
	return r.client.ScanType(ctx, cursor, match, count, keyType)
}

// SetArgs implements redis.Cmdable
func (r *Redis) SetArgs(ctx context.Context, key string, value any, a redis.SetArgs) *redis.StatusCmd {
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

func (r *Redis) Echo(ctx context.Context, message any) *redis.StringCmd {
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

func (r *Redis) GetSet(ctx context.Context, key string, value any) *redis.StringCmd {
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

func (r *Redis) MSet(ctx context.Context, pairs ...any) *redis.StatusCmd {
	return r.client.MSet(ctx, pairs...)
}

func (r *Redis) MSetNX(ctx context.Context, pairs ...any) *redis.BoolCmd {
	return r.client.MSetNX(ctx, pairs...)
}

func (r *Redis) Set(ctx context.Context, key string, value any, expiration time.Duration) *redis.StatusCmd {
	return r.client.Set(ctx, key, value, expiration)
}

func (r *Redis) SetBit(ctx context.Context, key string, offset int64, value int) *redis.IntCmd {
	return r.client.SetBit(ctx, key, offset, value)
}

func (r *Redis) SetNX(ctx context.Context, key string, value any, expiration time.Duration) *redis.BoolCmd {
	return r.client.SetNX(ctx, key, value, expiration)
}

func (r *Redis) SetXX(ctx context.Context, key string, value any, expiration time.Duration) *redis.BoolCmd {
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

func (r *Redis) HMSet(ctx context.Context, key string, values ...any) *redis.BoolCmd {
	return r.client.HMSet(ctx, key, values...)
}

func (r *Redis) HSet(ctx context.Context, key string, values ...any) *redis.IntCmd {
	return r.client.HSet(ctx, key, values...)
}

func (r *Redis) HSetNX(ctx context.Context, key, field string, value any) *redis.BoolCmd {
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

func (r *Redis) LInsert(ctx context.Context, key, op string, pivot, value any) *redis.IntCmd {
	return r.client.LInsert(ctx, key, op, pivot, value)
}

func (r *Redis) LInsertBefore(ctx context.Context, key string, pivot, value any) *redis.IntCmd {
	return r.client.LInsertBefore(ctx, key, pivot, value)
}

func (r *Redis) LInsertAfter(ctx context.Context, key string, pivot, value any) *redis.IntCmd {
	return r.client.LInsertAfter(ctx, key, pivot, value)
}

func (r *Redis) LLen(ctx context.Context, key string) *redis.IntCmd {
	return r.client.LLen(ctx, key)
}

func (r *Redis) LPop(ctx context.Context, key string) *redis.StringCmd {
	return r.client.LPop(ctx, key)
}

func (r *Redis) LPush(ctx context.Context, key string, values ...any) *redis.IntCmd {
	return r.client.LPush(ctx, key, values...)
}

func (r *Redis) LPushX(ctx context.Context, key string, values ...any) *redis.IntCmd {
	return r.client.LPushX(ctx, key, values...)
}

func (r *Redis) LRange(ctx context.Context, key string, start, stop int64) *redis.StringSliceCmd {
	return r.client.LRange(ctx, key, start, stop)
}

func (r *Redis) LRem(ctx context.Context, key string, count int64, value any) *redis.IntCmd {
	return r.client.LRem(ctx, key, count, value)
}

func (r *Redis) LSet(ctx context.Context, key string, index int64, value any) *redis.StatusCmd {
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

func (r *Redis) RPush(ctx context.Context, key string, values ...any) *redis.IntCmd {
	return r.client.RPush(ctx, key, values...)
}

func (r *Redis) RPushX(ctx context.Context, key string, values ...any) *redis.IntCmd {
	return r.client.RPushX(ctx, key, values...)
}

func (r *Redis) SAdd(ctx context.Context, key string, members ...any) *redis.IntCmd {
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

func (r *Redis) SIsMember(ctx context.Context, key string, member any) *redis.BoolCmd {
	return r.client.SIsMember(ctx, key, member)
}

func (r *Redis) SMembers(ctx context.Context, key string) *redis.StringSliceCmd {
	return r.client.SMembers(ctx, key)
}

func (r *Redis) SMembersMap(ctx context.Context, key string) *redis.StringStructMapCmd {
	return r.client.SMembersMap(ctx, key)
}

func (r *Redis) SMove(ctx context.Context, source, destination string, member any) *redis.BoolCmd {
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

func (r *Redis) SRem(ctx context.Context, key string, members ...any) *redis.IntCmd {
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

func (r *Redis) ZRem(ctx context.Context, key string, members ...any) *redis.IntCmd {
	return r.client.ZRem(ctx, key, members...)
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

func (r *Redis) PFAdd(ctx context.Context, key string, els ...any) *redis.IntCmd {
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

func (r *Redis) Eval(ctx context.Context, script string, keys []string, args ...any) *redis.Cmd {
	return r.client.Eval(ctx, script, keys, args...)
}

func (r *Redis) EvalSha(ctx context.Context, sha1 string, keys []string, args ...any) *redis.Cmd {
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

func (r *Redis) Publish(ctx context.Context, channel string, message any) *redis.IntCmd {
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
func (r *Redis) BitFieldRO(ctx context.Context, key string, values ...any) *redis.IntSliceCmd {
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
func (r *Redis) FTAlter(ctx context.Context, index string, skipInitialScan bool, definition []any) *redis.StatusCmd {
	return r.client.FTAlter(ctx, index, skipInitialScan, definition)
}

// FTConfigGet implements redis.Cmdable.
func (r *Redis) FTConfigGet(ctx context.Context, option string) *redis.MapMapStringInterfaceCmd {
	return r.client.FTConfigGet(ctx, option)
}

// FTConfigSet implements redis.Cmdable.
func (r *Redis) FTConfigSet(ctx context.Context, option string, value any) *redis.StatusCmd {
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
func (r *Redis) FTDictAdd(ctx context.Context, dict string, term ...any) *redis.IntCmd {
	return r.client.FTDictAdd(ctx, dict, term...)
}

// FTDictDel implements redis.Cmdable.
func (r *Redis) FTDictDel(ctx context.Context, dict string, term ...any) *redis.IntCmd {
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
func (r *Redis) FTSynUpdate(ctx context.Context, index string, synGroupId any, terms []any) *redis.StatusCmd {
	return r.client.FTSynUpdate(ctx, index, synGroupId, terms)
}

// FTSynUpdateWithArgs implements redis.Cmdable.
func (r *Redis) FTSynUpdateWithArgs(ctx context.Context, index string, synGroupId any, options *redis.FTSynUpdateOptions, terms []any) *redis.StatusCmd {
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
	return r.client.HExpire(ctx, key, expiration, fields...)
}

// HExpireAt implements redis.Cmdable.
func (r *Redis) HExpireAt(ctx context.Context, key string, tm time.Time, fields ...string) *redis.IntSliceCmd {
	return r.client.HExpireAt(ctx, key, tm, fields...)
}

// HExpireAtWithArgs implements redis.Cmdable.
func (r *Redis) HExpireAtWithArgs(ctx context.Context, key string, tm time.Time, expirationArgs redis.HExpireArgs, fields ...string) *redis.IntSliceCmd {
	return r.client.HExpireAtWithArgs(ctx, key, tm, expirationArgs, fields...)
}

// HExpireTime implements redis.Cmdable.
func (r *Redis) HExpireTime(ctx context.Context, key string, fields ...string) *redis.IntSliceCmd {
	return r.client.HExpireTime(ctx, key, fields...)
}

// HExpireWithArgs implements redis.Cmdable.
func (r *Redis) HExpireWithArgs(ctx context.Context, key string, expiration time.Duration, expirationArgs redis.HExpireArgs, fields ...string) *redis.IntSliceCmd {
	return r.client.HExpireWithArgs(ctx, key, expiration, expirationArgs, fields...)
}

// HGetDel implements redis.Cmdable.
func (r *Redis) HGetDel(ctx context.Context, key string, fields ...string) *redis.StringSliceCmd {
	return r.client.HGetDel(ctx, key, fields...)
}

// HGetEX implements redis.Cmdable.
func (r *Redis) HGetEX(ctx context.Context, key string, fields ...string) *redis.StringSliceCmd {
	return r.client.HGetEX(ctx, key, fields...)
}

// HGetEXWithArgs implements redis.Cmdable.
func (r *Redis) HGetEXWithArgs(ctx context.Context, key string, options *redis.HGetEXOptions, fields ...string) *redis.StringSliceCmd {
	return r.client.HGetEXWithArgs(ctx, key, options, fields...)
}

// HPExpire implements redis.Cmdable.
func (r *Redis) HPExpire(ctx context.Context, key string, expiration time.Duration, fields ...string) *redis.IntSliceCmd {
	return r.client.HPExpire(ctx, key, expiration, fields...)
}

// HPExpireAt implements redis.Cmdable.
func (r *Redis) HPExpireAt(ctx context.Context, key string, tm time.Time, fields ...string) *redis.IntSliceCmd {
	return r.client.HPExpireAt(ctx, key, tm, fields...)
}

// HPExpireAtWithArgs implements redis.Cmdable.
func (r *Redis) HPExpireAtWithArgs(ctx context.Context, key string, tm time.Time, expirationArgs redis.HExpireArgs, fields ...string) *redis.IntSliceCmd {
	return r.client.HPExpireAtWithArgs(ctx, key, tm, expirationArgs, fields...)
}

// HPExpireTime implements redis.Cmdable.
func (r *Redis) HPExpireTime(ctx context.Context, key string, fields ...string) *redis.IntSliceCmd {
	return r.client.HPExpireTime(ctx, key, fields...)
}

// HPExpireWithArgs implements redis.Cmdable.
func (r *Redis) HPExpireWithArgs(ctx context.Context, key string, expiration time.Duration, expirationArgs redis.HExpireArgs, fields ...string) *redis.IntSliceCmd {
	return r.client.HPExpireWithArgs(ctx, key, expiration, expirationArgs, fields...)
}

// HPTTL implements redis.Cmdable.
func (r *Redis) HPTTL(ctx context.Context, key string, fields ...string) *redis.IntSliceCmd {
	return r.client.HPTTL(ctx, key, fields...)
}

// HPersist implements redis.Cmdable.
func (r *Redis) HPersist(ctx context.Context, key string, fields ...string) *redis.IntSliceCmd {
	return r.client.HPersist(ctx, key, fields...)
}

// HScanNoValues implements redis.Cmdable.
func (r *Redis) HScanNoValues(ctx context.Context, key string, cursor uint64, match string, count int64) *redis.ScanCmd {
	return r.client.HScanNoValues(ctx, key, cursor, match, count)
}

// HSetEX implements redis.Cmdable.
func (r *Redis) HSetEX(ctx context.Context, key string, fieldsAndValues ...string) *redis.IntCmd {
	return r.client.HSetEX(ctx, key, fieldsAndValues...)
}

// HSetEXWithArgs implements redis.Cmdable.
func (r *Redis) HSetEXWithArgs(ctx context.Context, key string, options *redis.HSetEXOptions, fieldsAndValues ...string) *redis.IntCmd {
	return r.client.HSetEXWithArgs(ctx, key, options, fieldsAndValues...)
}

// HStrLen implements redis.Cmdable.
func (r *Redis) HStrLen(ctx context.Context, key string, field string) *redis.IntCmd {
	return r.client.HStrLen(ctx, key, field)
}

// HTTL implements redis.Cmdable.
func (r *Redis) HTTL(ctx context.Context, key string, fields ...string) *redis.IntSliceCmd {
	return r.client.HTTL(ctx, key, fields...)
}

// JSONArrAppend implements redis.Cmdable.
func (r *Redis) JSONArrAppend(ctx context.Context, key string, path string, values ...any) *redis.IntSliceCmd {
	return r.client.JSONArrAppend(ctx, key, path, values...)
}

// JSONArrIndex implements redis.Cmdable.
func (r *Redis) JSONArrIndex(ctx context.Context, key string, path string, value ...any) *redis.IntSliceCmd {
	return r.client.JSONArrIndex(ctx, key, path, value...)
}

// JSONArrIndexWithArgs implements redis.Cmdable.
func (r *Redis) JSONArrIndexWithArgs(ctx context.Context, key string, path string, options *redis.JSONArrIndexArgs, value ...any) *redis.IntSliceCmd {
	return r.client.JSONArrIndexWithArgs(ctx, key, path, options, value...)
}

// JSONArrInsert implements redis.Cmdable.
func (r *Redis) JSONArrInsert(ctx context.Context, key string, path string, index int64, values ...any) *redis.IntSliceCmd {
	return r.client.JSONArrInsert(ctx, key, path, index, values...)
}

// JSONArrLen implements redis.Cmdable.
func (r *Redis) JSONArrLen(ctx context.Context, key string, path string) *redis.IntSliceCmd {
	return r.client.JSONArrLen(ctx, key, path)
}

// JSONArrPop implements redis.Cmdable.
func (r *Redis) JSONArrPop(ctx context.Context, key string, path string, index int) *redis.StringSliceCmd {
	return r.client.JSONArrPop(ctx, key, path, index)
}

// JSONArrTrim implements redis.Cmdable.
func (r *Redis) JSONArrTrim(ctx context.Context, key string, path string) *redis.IntSliceCmd {
	return r.client.JSONArrTrim(ctx, key, path)
}

// JSONArrTrimWithArgs implements redis.Cmdable.
func (r *Redis) JSONArrTrimWithArgs(ctx context.Context, key string, path string, options *redis.JSONArrTrimArgs) *redis.IntSliceCmd {
	return r.client.JSONArrTrimWithArgs(ctx, key, path, options)
}

// JSONClear implements redis.Cmdable.
func (r *Redis) JSONClear(ctx context.Context, key string, path string) *redis.IntCmd {
	return r.client.JSONClear(ctx, key, path)
}

// JSONDebugMemory implements redis.Cmdable.
func (r *Redis) JSONDebugMemory(ctx context.Context, key string, path string) *redis.IntCmd {
	return r.client.JSONDebugMemory(ctx, key, path)
}

// JSONDel implements redis.Cmdable.
func (r *Redis) JSONDel(ctx context.Context, key string, path string) *redis.IntCmd {
	return r.client.JSONDel(ctx, key, path)
}

// JSONForget implements redis.Cmdable.
func (r *Redis) JSONForget(ctx context.Context, key string, path string) *redis.IntCmd {
	return r.client.JSONForget(ctx, key, path)
}

// JSONGet implements redis.Cmdable.
func (r *Redis) JSONGet(ctx context.Context, key string, paths ...string) *redis.JSONCmd {
	return r.client.JSONGet(ctx, key, paths...)
}

// JSONGetWithArgs implements redis.Cmdable.
func (r *Redis) JSONGetWithArgs(ctx context.Context, key string, options *redis.JSONGetArgs, paths ...string) *redis.JSONCmd {
	return r.client.JSONGetWithArgs(ctx, key, options, paths...)
}

// JSONMGet implements redis.Cmdable.
func (r *Redis) JSONMGet(ctx context.Context, path string, keys ...string) *redis.JSONSliceCmd {
	return r.client.JSONMGet(ctx, path, keys...)
}

// JSONMSet implements redis.Cmdable.
func (r *Redis) JSONMSet(ctx context.Context, params ...any) *redis.StatusCmd {
	return r.client.JSONMSet(ctx, params...)
}

// JSONMSetArgs implements redis.Cmdable.
func (r *Redis) JSONMSetArgs(ctx context.Context, docs []redis.JSONSetArgs) *redis.StatusCmd {
	return r.client.JSONMSetArgs(ctx, docs)
}

// JSONMerge implements redis.Cmdable.
func (r *Redis) JSONMerge(ctx context.Context, key string, path string, value string) *redis.StatusCmd {
	return r.client.JSONMerge(ctx, key, path, value)
}

// JSONNumIncrBy implements redis.Cmdable.
func (r *Redis) JSONNumIncrBy(ctx context.Context, key string, path string, value float64) *redis.JSONCmd {
	return r.client.JSONNumIncrBy(ctx, key, path, value)
}

// JSONObjKeys implements redis.Cmdable.
func (r *Redis) JSONObjKeys(ctx context.Context, key string, path string) *redis.SliceCmd {
	return r.client.JSONObjKeys(ctx, key, path)
}

// JSONObjLen implements redis.Cmdable.
func (r *Redis) JSONObjLen(ctx context.Context, key string, path string) *redis.IntPointerSliceCmd {
	return r.client.JSONObjLen(ctx, key, path)
}

// JSONSet implements redis.Cmdable.
func (r *Redis) JSONSet(ctx context.Context, key string, path string, value any) *redis.StatusCmd {
	return r.client.JSONSet(ctx, key, path, value)
}

// JSONSetMode implements redis.Cmdable.
func (r *Redis) JSONSetMode(ctx context.Context, key string, path string, value any, mode string) *redis.StatusCmd {
	return r.client.JSONSetMode(ctx, key, path, value, mode)
}

// JSONStrAppend implements redis.Cmdable.
func (r *Redis) JSONStrAppend(ctx context.Context, key string, path string, value string) *redis.IntPointerSliceCmd {
	return r.client.JSONStrAppend(ctx, key, path, value)
}

// JSONStrLen implements redis.Cmdable.
func (r *Redis) JSONStrLen(ctx context.Context, key string, path string) *redis.IntPointerSliceCmd {
	return r.client.JSONStrLen(ctx, key, path)
}

// JSONToggle implements redis.Cmdable.
func (r *Redis) JSONToggle(ctx context.Context, key string, path string) *redis.IntPointerSliceCmd {
	return r.client.JSONToggle(ctx, key, path)
}

// JSONType implements redis.Cmdable.
func (r *Redis) JSONType(ctx context.Context, key string, path string) *redis.JSONSliceCmd {
	return r.client.JSONType(ctx, key, path)
}

// ObjectFreq implements redis.Cmdable.
func (r *Redis) ObjectFreq(ctx context.Context, key string) *redis.IntCmd {
	return r.client.ObjectFreq(ctx, key)
}

// TSAdd implements redis.Cmdable.
func (r *Redis) TSAdd(ctx context.Context, key string, timestamp any, value float64) *redis.IntCmd {
	return r.client.TSAdd(ctx, key, timestamp, value)
}

// TSAddWithArgs implements redis.Cmdable.
func (r *Redis) TSAddWithArgs(ctx context.Context, key string, timestamp any, value float64, options *redis.TSOptions) *redis.IntCmd {
	return r.client.TSAddWithArgs(ctx, key, timestamp, value, options)
}

// TSAlter implements redis.Cmdable.
func (r *Redis) TSAlter(ctx context.Context, key string, options *redis.TSAlterOptions) *redis.StatusCmd {
	return r.client.TSAlter(ctx, key, options)
}

// TSCreate implements redis.Cmdable.
func (r *Redis) TSCreate(ctx context.Context, key string) *redis.StatusCmd {
	return r.client.TSCreate(ctx, key)
}

// TSCreateRule implements redis.Cmdable.
func (r *Redis) TSCreateRule(ctx context.Context, sourceKey string, destKey string, aggregator redis.Aggregator, bucketDuration int) *redis.StatusCmd {
	return r.client.TSCreateRule(ctx, sourceKey, destKey, aggregator, bucketDuration)
}

// TSCreateRuleWithArgs implements redis.Cmdable.
func (r *Redis) TSCreateRuleWithArgs(ctx context.Context, sourceKey string, destKey string, aggregator redis.Aggregator, bucketDuration int, options *redis.TSCreateRuleOptions) *redis.StatusCmd {
	return r.client.TSCreateRuleWithArgs(ctx, sourceKey, destKey, aggregator, bucketDuration, options)
}

// TSCreateWithArgs implements redis.Cmdable.
func (r *Redis) TSCreateWithArgs(ctx context.Context, key string, options *redis.TSOptions) *redis.StatusCmd {
	return r.client.TSCreateWithArgs(ctx, key, options)
}

// TSDecrBy implements redis.Cmdable.
func (r *Redis) TSDecrBy(ctx context.Context, Key string, timestamp float64) *redis.IntCmd {
	return r.client.TSDecrBy(ctx, Key, timestamp)
}

// TSDecrByWithArgs implements redis.Cmdable.
func (r *Redis) TSDecrByWithArgs(ctx context.Context, key string, timestamp float64, options *redis.TSIncrDecrOptions) *redis.IntCmd {
	return r.client.TSDecrByWithArgs(ctx, key, timestamp, options)
}

// TSDel implements redis.Cmdable.
func (r *Redis) TSDel(ctx context.Context, Key string, fromTimestamp int, toTimestamp int) *redis.IntCmd {
	return r.client.TSDel(ctx, Key, fromTimestamp, toTimestamp)
}

// TSDeleteRule implements redis.Cmdable.
func (r *Redis) TSDeleteRule(ctx context.Context, sourceKey string, destKey string) *redis.StatusCmd {
	return r.client.TSDeleteRule(ctx, sourceKey, destKey)
}

// TSGet implements redis.Cmdable.
func (r *Redis) TSGet(ctx context.Context, key string) *redis.TSTimestampValueCmd {
	return r.client.TSGet(ctx, key)
}

// TSGetWithArgs implements redis.Cmdable.
func (r *Redis) TSGetWithArgs(ctx context.Context, key string, options *redis.TSGetOptions) *redis.TSTimestampValueCmd {
	return r.client.TSGetWithArgs(ctx, key, options)
}

// TSIncrBy implements redis.Cmdable.
func (r *Redis) TSIncrBy(ctx context.Context, Key string, timestamp float64) *redis.IntCmd {
	return r.client.TSIncrBy(ctx, Key, timestamp)
}

// TSIncrByWithArgs implements redis.Cmdable.
func (r *Redis) TSIncrByWithArgs(ctx context.Context, key string, timestamp float64, options *redis.TSIncrDecrOptions) *redis.IntCmd {
	return r.client.TSIncrByWithArgs(ctx, key, timestamp, options)
}

// TSInfo implements redis.Cmdable.
func (r *Redis) TSInfo(ctx context.Context, key string) *redis.MapStringInterfaceCmd {
	return r.client.TSInfo(ctx, key)
}

// TSInfoWithArgs implements redis.Cmdable.
func (r *Redis) TSInfoWithArgs(ctx context.Context, key string, options *redis.TSInfoOptions) *redis.MapStringInterfaceCmd {
	return r.client.TSInfoWithArgs(ctx, key, options)
}

// TSMAdd implements redis.Cmdable.
func (r *Redis) TSMAdd(ctx context.Context, ktvSlices [][]any) *redis.IntSliceCmd {
	return r.client.TSMAdd(ctx, ktvSlices)
}

// TSMGet implements redis.Cmdable.
func (r *Redis) TSMGet(ctx context.Context, filters []string) *redis.MapStringSliceInterfaceCmd {
	return r.client.TSMGet(ctx, filters)
}

// TSMGetWithArgs implements redis.Cmdable.
func (r *Redis) TSMGetWithArgs(ctx context.Context, filters []string, options *redis.TSMGetOptions) *redis.MapStringSliceInterfaceCmd {
	return r.client.TSMGetWithArgs(ctx, filters, options)
}

// TSMRange implements redis.Cmdable.
func (r *Redis) TSMRange(ctx context.Context, fromTimestamp int, toTimestamp int, filterExpr []string) *redis.MapStringSliceInterfaceCmd {
	return r.client.TSMRange(ctx, fromTimestamp, toTimestamp, filterExpr)
}

// TSMRangeWithArgs implements redis.Cmdable.
func (r *Redis) TSMRangeWithArgs(ctx context.Context, fromTimestamp int, toTimestamp int, filterExpr []string, options *redis.TSMRangeOptions) *redis.MapStringSliceInterfaceCmd {
	return r.client.TSMRangeWithArgs(ctx, fromTimestamp, toTimestamp, filterExpr, options)
}

// TSMRevRange implements redis.Cmdable.
func (r *Redis) TSMRevRange(ctx context.Context, fromTimestamp int, toTimestamp int, filterExpr []string) *redis.MapStringSliceInterfaceCmd {
	return r.client.TSMRevRange(ctx, fromTimestamp, toTimestamp, filterExpr)
}

// TSMRevRangeWithArgs implements redis.Cmdable.
func (r *Redis) TSMRevRangeWithArgs(ctx context.Context, fromTimestamp int, toTimestamp int, filterExpr []string, options *redis.TSMRevRangeOptions) *redis.MapStringSliceInterfaceCmd {
	return r.client.TSMRevRangeWithArgs(ctx, fromTimestamp, toTimestamp, filterExpr, options)
}

// TSQueryIndex implements redis.Cmdable.
func (r *Redis) TSQueryIndex(ctx context.Context, filterExpr []string) *redis.StringSliceCmd {
	return r.client.TSQueryIndex(ctx, filterExpr)
}

// TSRange implements redis.Cmdable.
func (r *Redis) TSRange(ctx context.Context, key string, fromTimestamp int, toTimestamp int) *redis.TSTimestampValueSliceCmd {
	return r.client.TSRange(ctx, key, fromTimestamp, toTimestamp)
}

// TSRangeWithArgs implements redis.Cmdable.
func (r *Redis) TSRangeWithArgs(ctx context.Context, key string, fromTimestamp int, toTimestamp int, options *redis.TSRangeOptions) *redis.TSTimestampValueSliceCmd {
	return r.client.TSRangeWithArgs(ctx, key, fromTimestamp, toTimestamp, options)
}

// TSRevRange implements redis.Cmdable.
func (r *Redis) TSRevRange(ctx context.Context, key string, fromTimestamp int, toTimestamp int) *redis.TSTimestampValueSliceCmd {
	return r.client.TSRevRange(ctx, key, fromTimestamp, toTimestamp)
}

// TSRevRangeWithArgs implements redis.Cmdable.
func (r *Redis) TSRevRangeWithArgs(ctx context.Context, key string, fromTimestamp int, toTimestamp int, options *redis.TSRevRangeOptions) *redis.TSTimestampValueSliceCmd {
	return r.client.TSRevRangeWithArgs(ctx, key, fromTimestamp, toTimestamp, options)
}

// VAdd implements redis.Cmdable.
func (r *Redis) VAdd(ctx context.Context, key string, element string, val redis.Vector) *redis.BoolCmd {
	return r.client.VAdd(ctx, key, element, val)
}

// VAddWithArgs implements redis.Cmdable.
func (r *Redis) VAddWithArgs(ctx context.Context, key string, element string, val redis.Vector, addArgs *redis.VAddArgs) *redis.BoolCmd {
	return r.client.VAddWithArgs(ctx, key, element, val, addArgs)
}

// VCard implements redis.Cmdable.
func (r *Redis) VCard(ctx context.Context, key string) *redis.IntCmd {
	return r.client.VCard(ctx, key)
}

// VClearAttributes implements redis.Cmdable.
func (r *Redis) VClearAttributes(ctx context.Context, key string, element string) *redis.BoolCmd {
	return r.client.VClearAttributes(ctx, key, element)
}

// VDim implements redis.Cmdable.
func (r *Redis) VDim(ctx context.Context, key string) *redis.IntCmd {
	return r.client.VDim(ctx, key)
}

// VEmb implements redis.Cmdable.
func (r *Redis) VEmb(ctx context.Context, key string, element string, raw bool) *redis.SliceCmd {
	return r.client.VEmb(ctx, key, element, raw)
}

// VGetAttr implements redis.Cmdable.
func (r *Redis) VGetAttr(ctx context.Context, key string, element string) *redis.StringCmd {
	return r.client.VGetAttr(ctx, key, element)
}

// VInfo implements redis.Cmdable.
func (r *Redis) VInfo(ctx context.Context, key string) *redis.MapStringInterfaceCmd {
	return r.client.VInfo(ctx, key)
}

// VLinks implements redis.Cmdable.
func (r *Redis) VLinks(ctx context.Context, key string, element string) *redis.StringSliceSliceCmd {
	return r.client.VLinks(ctx, key, element)
}

// VLinksWithScores implements redis.Cmdable.
func (r *Redis) VLinksWithScores(ctx context.Context, key string, element string) *redis.VectorScoreSliceSliceCmd {
	return r.client.VLinksWithScores(ctx, key, element)
}

// VRandMember implements redis.Cmdable.
func (r *Redis) VRandMember(ctx context.Context, key string) *redis.StringCmd {
	return r.client.VRandMember(ctx, key)
}

// VRandMemberCount implements redis.Cmdable.
func (r *Redis) VRandMemberCount(ctx context.Context, key string, count int) *redis.StringSliceCmd {
	return r.client.VRandMemberCount(ctx, key, count)
}

// VRem implements redis.Cmdable.
func (r *Redis) VRem(ctx context.Context, key string, element string) *redis.BoolCmd {
	return r.client.VRem(ctx, key, element)
}

// VSetAttr implements redis.Cmdable.
func (r *Redis) VSetAttr(ctx context.Context, key string, element string, attr any) *redis.BoolCmd {
	return r.client.VSetAttr(ctx, key, element, attr)
}

// VSim implements redis.Cmdable.
func (r *Redis) VSim(ctx context.Context, key string, val redis.Vector) *redis.StringSliceCmd {
	return r.client.VSim(ctx, key, val)
}

// VSimWithArgs implements redis.Cmdable.
func (r *Redis) VSimWithArgs(ctx context.Context, key string, val redis.Vector, args *redis.VSimArgs) *redis.StringSliceCmd {
	return r.client.VSimWithArgs(ctx, key, val, args)
}

// VSimWithArgsWithScores implements redis.Cmdable.
func (r *Redis) VSimWithArgsWithScores(ctx context.Context, key string, val redis.Vector, args *redis.VSimArgs) *redis.VectorScoreSliceCmd {
	return r.client.VSimWithArgsWithScores(ctx, key, val, args)
}

// VSimWithScores implements redis.Cmdable.
func (r *Redis) VSimWithScores(ctx context.Context, key string, val redis.Vector) *redis.VectorScoreSliceCmd {
	return r.client.VSimWithScores(ctx, key, val)
}

// XAckDel implements redis.Cmdable.
func (r *Redis) XAckDel(ctx context.Context, stream string, group string, mode string, ids ...string) *redis.SliceCmd {
	return r.client.XAckDel(ctx, stream, group, mode, ids...)
}

// XDelEx implements redis.Cmdable.
func (r *Redis) XDelEx(ctx context.Context, stream string, mode string, ids ...string) *redis.SliceCmd {
	return r.client.XDelEx(ctx, stream, mode, ids...)
}

// XTrimMaxLenApproxMode implements redis.Cmdable.
func (r *Redis) XTrimMaxLenApproxMode(ctx context.Context, key string, maxLen int64, limit int64, mode string) *redis.IntCmd {
	return r.client.XTrimMaxLenApproxMode(ctx, key, maxLen, limit, mode)
}

// XTrimMaxLenMode implements redis.Cmdable.
func (r *Redis) XTrimMaxLenMode(ctx context.Context, key string, maxLen int64, mode string) *redis.IntCmd {
	return r.client.XTrimMaxLenMode(ctx, key, maxLen, mode)
}

// XTrimMinIDApproxMode implements redis.Cmdable.
func (r *Redis) XTrimMinIDApproxMode(ctx context.Context, key string, minID string, limit int64, mode string) *redis.IntCmd {
	return r.client.XTrimMinIDApproxMode(ctx, key, minID, limit, mode)
}

// XTrimMinIDMode implements redis.Cmdable.
func (r *Redis) XTrimMinIDMode(ctx context.Context, key string, minID string, mode string) *redis.IntCmd {
	return r.client.XTrimMinIDMode(ctx, key, minID, mode)
}
