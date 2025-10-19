# Kafka

![kafka](./public/f-4.png)

## Connectors

```sh
    curl -i localhost:8083
    curl -i localhost:8083/connectors
    curl -i localhost:8083/connectors/[connector-name]/status

    ## config 삭제 / 재등록
    curl -X DELETE http://localhost:8083/connectors/[connector-name]
    curl -X POST http://localhost:8083/connectors \
      -H "Content-Type: application/json" \
      -d @$(PWD)/configs//debizium/source.json
    
```

## Trouble Shooting

```sh
2025-10-19 05:38:44,735 WARN   ||  WorkerSourceTask{id=mysql-source-connector-0} failed to poll records from SourceTask. Will retry operation.   [org.apache.kafka.connect.runtime.AbstractWorkerSourceTask]
org.apache.kafka.connect.errors.RetriableException: An exception occurred in the change event producer. This connector will be restarted.
        at io.debezium.pipeline.ErrorHandler.setProducerThrowable(ErrorHandler.java:63)
        at io.debezium.pipeline.ChangeEventSourceCoordinator.lambda$start$0(ChangeEventSourceCoordinator.java:144)
        at java.base/java.util.concurrent.Executors$RunnableAdapter.call(Executors.java:515)
        at java.base/java.util.concurrent.FutureTask.run(FutureTask.java:264)
        at java.base/java.util.concurrent.ThreadPoolExecutor.runWorker(ThreadPoolExecutor.java:1128)
        at java.base/java.util.concurrent.ThreadPoolExecutor$Worker.run(ThreadPoolExecutor.java:628)
        at java.base/java.lang.Thread.run(Thread.java:829)
Caused by: io.debezium.DebeziumException: java.sql.SQLSyntaxErrorException: Access denied; you need (at least one of) the RELOAD or FLUSH_TABLES privilege(s) for this operation
        at io.debezium.pipeline.source.AbstractSnapshotChangeEventSource.execute(AbstractSnapshotChangeEventSource.java:101)
        at io.debezium.pipeline.ChangeEventSourceCoordinator.doSnapshot(ChangeEventSourceCoordinator.java:250)
        at io.debezium.pipeline.ChangeEventSourceCoordinator.doSnapshot(ChangeEventSourceCoordinator.java:234)
        at io.debezium.pipeline.ChangeEventSourceCoordinator.executeChangeEventSources(ChangeEventSourceCoordinator.java:186)
        at io.debezium.pipeline.ChangeEventSourceCoordinator.lambda$start$0(ChangeEventSourceCoordinator.java:137)
        ... 5 more
Caused by: java.sql.SQLSyntaxErrorException: Access denied; you need (at least one of) the RELOAD or FLUSH_TABLES privilege(s) for this operation
        at com.mysql.cj.jdbc.exceptions.SQLError.createSQLException(SQLError.java:121)
        at com.mysql.cj.jdbc.exceptions.SQLExceptionsMapping.translateException(SQLExceptionsMapping.java:122)
        at com.mysql.cj.jdbc.StatementImpl.executeInternal(StatementImpl.java:763)
        at com.mysql.cj.jdbc.StatementImpl.execute(StatementImpl.java:648)
        at io.debezium.jdbc.JdbcConnection.executeWithoutCommitting(JdbcConnection.java:1451)
        at io.debezium.connector.mysql.MySqlSnapshotChangeEventSource.tableLock(MySqlSnapshotChangeEventSource.java:524)
        at io.debezium.connector.mysql.MySqlSnapshotChangeEventSource.readTableStructure(MySqlSnapshotChangeEventSource.java:313)
        at io.debezium.connector.mysql.MySqlSnapshotChangeEventSource.readTableStructure(MySqlSnapshotChangeEventSource.java:59)
        at io.debezium.relational.RelationalSnapshotChangeEventSource.doExecute(RelationalSnapshotChangeEventSource.java:149)
        at io.debezium.pipeline.source.AbstractSnapshotChangeEventSource.execute(AbstractSnapshotChangeEventSource.java:92)
        ... 9 more
2025-10-19 05:38:44,736 INFO   ||  Awaiting end of restart backoff period after a retriable error   [io.debezium.connector.common.BaseSourceTask]
```

- Debizium 시, 위와 같은 에러 발생
- 권한이 재대로 부여되지 않거나, 데이터베이스 설정이 재대로 되지 않은 경우 발생

```sh
CREATE USER 'debezium'@'%' IDENTIFIED BY 'debezium1234';

# Debezium에 필요한 모든 권한 부여
GRANT SELECT ON test_db.* TO 'debezium'@'%';
GRANT RELOAD ON *.* TO 'debezium'@'%';
GRANT REPLICATION SLAVE ON *.* TO 'debezium'@'%';
GRANT REPLICATION CLIENT ON *.* TO 'debezium'@'%';
GRANT LOCK TABLES ON test_db.* TO 'debezium'@'%';
FLUSH PRIVILEGES;

# 권한 확인
SHOW GRANTS FOR 'debezium'@'%';
```

## Ref

- <a href="https://www.confluent.io/hub/"> Confluent Hub </a>
- <a href="https://debezium.io/documentation/"> Debizium Document </a>