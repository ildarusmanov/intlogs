Action logs service
==========================

## Save some log data

```
# POST /v1/create?token=super-token
{
   "ActionName":"action name",
   "ActionTarget":"action target",
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
# GET /v1/get?token=super-token&page=0
[
    {
        "Id": "599588207f1d2e7ca30541b7",
        "ActionName": "asdasd",
        "ActionTarget": "",
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

Run docker container
```
cd intlogs
sudo docker build -t intlogs .
// prod
sudo docker run -d -p 10.90.137.73:8000:8000 --network host intlogs
// or dev
sudo docker run -p 8000:8000 --network host intlogs
// list containers
sudo docker ps
```

```

## Run tests
```
cd {project_directory}
go test ./models ./controllers ./user
```