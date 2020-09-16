### Create VPC Serverless connector

```
gcloud services enable vpcaccess.googleapis.com
```

```
gcloud compute networks vpc-access connectors create [CONNECTOR_NAME] \
--network [VPC_NETWORK] \
--region [REGION] \
--range [IP_RANGE]
```


```
gcloud functions deploy flusher \
--runtime go113 \
--trigger-topic [my-topic]  \
--region [REGION] \
--vpc-connector projects/[PROJECT_ID]/locations/[REGION]/connectors/[CONNECTOR_NAME] \
--set-env-vars REDISHOST=[REDIS_IP],REDISPORT=[REDIS_PORT]
```


### Call the function
```
gcloud functions call publish --data '{"topic":"[MY_TOPIC]","message":"flushall"}'
```

OR

publish to pubsub

