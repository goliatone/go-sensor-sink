

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
DROP VIEW IF EXISTS dht_readings_5m CASCADE;
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
GROUP BY bucket, hardware;
```

```sql
DROP VIEW IF EXISTS dht_readings_1h CASCADE;
CREATE VIEW dht_readings_1h
WITH (
    timescaledb.continuous,
    timescaledb.ignore_invalidation_older_than = '5d',
    timescaledb.refresh_lag = '-30m',
    timescaledb.refresh_interval = '1h'
)
AS 
    SELECT
        time_bucket('1h', time) as bucket, 
        hardware,
        AVG(humidity) as humidity_avg,
        MAX(humidity) as humidity_max,
        MIN(humidity) as humidity_min,
        AVG(temperature) as temperature_avg,
        MAX(temperature) as temperature_max,
        MIN(temperature) as temperature_min
    FROM dht_readings
    GROUP BY bucket, hardware;
```

```sql
DROP VIEW IF EXISTS dht_readings_1d CASCADE;
CREATE VIEW dht_readings_1d
WITH (
    timescaledb.continuous,
    timescaledb.ignore_invalidation_older_than = '31d',
    timescaledb.refresh_lag = '-30m',
    timescaledb.refresh_interval = '1d'
)
AS SELECT
  time_bucket('1d', time) as bucket, 
  hardware,
  AVG(humidity) as humidity_avg,
  MAX(humidity) as humidity_max,
  MIN(humidity) as humidity_min,
  AVG(temperature) as temperature_avg,
  MAX(temperature) as temperature_max,
  MIN(temperature) as temperature_min
FROM dht_readings
GROUP BY bucket, hardware;
```

```sql
DROP VIEW IF EXISTS dht_readings_30d CASCADE;
CREATE VIEW dht_readings_30d
WITH (
    timescaledb.continuous,
    timescaledb.ignore_invalidation_older_than = '31d',
    timescaledb.refresh_lag = '-30m',
    timescaledb.refresh_interval = '1d'
)
AS 
    SELECT
        time_bucket('30 days', time) as bucket, 
        hardware,
        AVG(humidity) as humidity_avg,
        MAX(humidity) as humidity_max,
        MIN(humidity) as humidity_min,
        AVG(temperature) as temperature_avg,
        MAX(temperature) as temperature_max,
        MIN(temperature) as temperature_min
    FROM dht_readings
    GROUP BY bucket, hardware;
```

NOTE: If you get an error similar to this 
>pq: relation "dht22_readings" does not exist
It's because we named our table `dht_readings` and our struct `DHT22Reading` and gorm will automatically generate a table name `dht22_readings` for it.



## MQTT

Configuration options:

- `SetCleanSession`: SetCleanSession will set the "clean session" flag in the connect message when this client connects to an MQTT broker. By setting this flag, you are indicating that no messages saved by the broker for this client should be delivered. Any messages that were going to be sent by this client before disconnecting previously but didn't will not be sent upon connecting to the broker. 


Resources:

- [Paho Go - MQTT Client Library Encyclopedia](https://www.hivemq.com/blog/mqtt-client-library-encyclopedia-golang/)
