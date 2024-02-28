CREATE TABLE analytics (
  Timestamp DateTime64(9) CODEC (Delta, ZSTD(1)),
  provider String CODEC (ZSTD(1)),
  model String CODEC (ZSTD(1)),
  latency String CODEC (ZSTD(1)),
  prompt_tokens UInt32 CODEC (ZSTD(1)),
  completion_tokens UInt32 CODEC (ZSTD(1)),
  total_tokens UInt32 CODEC (ZSTD(1)),
  INDEX idx_provider provider TYPE bloom_filter GRANULARITY 1,
) ENGINE = MergeTree PARTITION BY toDate(Timestamp)
ORDER BY
  toUnixTimestamp(Timestamp) TTL toDateTime(Timestamp) + toIntervalDay(30) SETTINGS index_granularity = 8192,
  ttl_only_drop_parts = 1;