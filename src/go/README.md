
lambda -- main entry points for lambda functions
compose -- main entry points for docker-compose style (aka Kubernettes)


General structure:

```
  cmd/              -- where all main functions reside
     lambda/        -- lambda mains
       core/        -- direct data access
       publicapi/   -- API access that calls backend services
       websocket/   -- eventbus: translate events to websocket payloads
       workers/     -- eventbus: further processing
       trigger/     -- translate AWS events to eventbus
     compose/ -- kubernetes (standalone) mains
       core/        -- direct data access
       publicapi/   -- API access that calls backend services
       websocket/   -- eventbus: translate events to websocket payloads
       workers/     -- eventbus: further processing
```


