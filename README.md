# Elastiflow to NetMeta utility

Hacky utility to migrate Elastiflow data (dumped with elasticdump) to NetMeta.

Reads raw JSON for each day and inserts it into Clickhouse.

Elasticdump command example:
```
elasticdump --input https://localhost:9200/elastiflow-flow-codex-2.2 --output=/tmp/2023-11-01.json --type=data --searchBody='{"docvalue_fields": [{"field": "flow.collect.timestamp", "format": "date_time"}, {"field": "@timestamp", "format": "date_time"}],"query": {"bool": {"filter": [{"range": {"@timestamp": {"gte": "2023-11-01T00:00:00.000Z", "lte": "2023-11-01T23:59:59.999Z", "format": "strict_date_optional_time"}}}]}}}' --limit 5000
```
