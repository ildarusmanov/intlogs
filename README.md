Action logs service
==========================
[![Build Status](https://travis-ci.org/ildarusmanov/intlogs.svg?branch=master)](https://travis-ci.org/ildarusmanov/intlogs)

## Save some log data

```
# POST /v1/create
# Auth-Token: {token}

{
   "ActionName":"action name",
   "ActionTarget":"action target",
   "ActionTargetId": "123",
   "ActionCost":0,
   "UserId":"dfg",
   "GuestUserId":"user id",
   "Url":"related url",
   "CreatedAt":12323123, // timestamp
   "Params":{
      "param1":"value1",
      "param2":"value2",
   }
}

```

## Get all logs

```
# GET /v1/get?page=0
# Auth-Token: {token}
[
    {
        "Id": "599588207f1d2e7ca30541b7",
        "ActionName": "asdasd",
        "ActionTarget": "",
        "ActionTargetId": "",
        "ActionCost": 0,
        "UserId": "",
        "GuestUserId": "",
        "Url": "",
        "CreatedAt": 0,
        "Params": null
    },
    // ...
    {
        "Id": "599ab4a67f1d2e7ca305490a",
        "ActionName": "sfdsdfq",
        "ActionTarget": "sdfsf",
        "ActionTargetId": "123",
        "ActionCost": 0,
        "UserId": "dfg",
        "GuestUserId": "sfsdfs",
        "Url": "sdfsdf",
        "CreatedAt": 100,
        "Params": {
            "asdasd": "asdad",
            "asdfgd": "afgdfgs"
        }
    }
]
```

### Filters
You can use query variables to filter logs list.
Available filters:
```
userId e.g. userId=1
userIds e.g. userIds=1,2,3

name e.g. name=abc
names e.g. names=abc,def

targetId e.g. targetId=111
targetIds e.g. targetIds=111,222

target e.g. target=abc
targets e.g. target=abc,def

cost e.g. cost=123
costFrom e.g. cost=1
costTo e.g. cost=100

createdFrom e.g. createdFrom=1451610610
createdTo e.g. createdTo=1530168403
```

# Setup
## Clone repository
```
git clone [repo]
```

## Run docker container
```
cd intlogs
sudo docker build -t intlogs .
// prod
sudo docker run -d -p 10.90.137.73:8003:8003 --network host intlogs
// or dev
sudo docker run -p 8003:8003 --network host intlogs
// list containers
sudo docker ps
```

## Run tests
```
cd {project_directory}
go test ./models ./controllers ./user
```