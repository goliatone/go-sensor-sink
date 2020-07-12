

## Development 


### TimescaleDB

To run a local copy of TimescaleDB using [docker](https://github.com/timescale/timescaledb-docker):

```bash
$ docker run -d --name timescaledb \
  -p 8432:5432 \
  -e POSTGRES_USER=root \
  -e POSTGRES_PASSWORD=password \
  -e POSTGRES_DB=sensor_sink \
  -e TIMESCALEDB_TELEMETRY=off \
  timescale/timescaledb:latest-pg12
```

To get you started you can find a list of tutorials in the [TimescaleDB documentation site](https://docs.timescale.com/latest/tutorials).


```sql
CREATE TABLE IF NOT EXISTS dht_readings (
    time TIMESTAMPTZ NOT NULL,
    hardware TEXT,
    temperature FLOAT NOT NULL,
    humidity FLOAT NOT NULL,
    PRIMARY KEY(time, hardware)
)

SELECT create_hypertable('dht_readings', 'time', if_not_exists => TRUE);
```


```sql
CREATE VIEW dht_readings_5m 
WITH (
    timescaledb.continuous,
    timescaledb.ignore_invalidation_older_than = '5d',
    timescaledb.refresh_lag = '-30m',
    timescaledb.refresh_interval = '5m'
)
AS SELECT
  time_bucket('5m', time) as bucket, 
  hardware,
  AVG(humidity) as humidity_avg,
  MAX(humidity) as humidity_max,
  MIN(humidity) as humidity_min,
  AVG(temperature) as temperature_avg,
  MAX(temperature) as temperature_max,
  MIN(temperature) as temperature_min
FROM dht_readings
GROUP BY bucket, hardware, temperature, humidity;
```

NOTE: If you get an error similar to this 
>pq: relation "dht22_readings" does not exist
It's because we named our table `dht_readings` and our struct `DHT22Reading` and gorm will automatically generate a table name `dht22_readings` for it.