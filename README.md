# Action logs service
==========================

## Save some log data

```
# POST /create
{
	"url": "http://vvv.vvv",
	"type": "myTypeName",
	"targetId": "someTargetIdValue",
	"userId": "userIdValue",
	"guestId": "guestUserIdValue",
	"createdAt": "timestamp"
}
```

## Get all logs

```
# GET /get
[
	{
		"id": "idValue1",
		"url": "http://vvv.vvv",
		"type": "myTypeName",
		"targetId": "someTargetIdValue",
		"userId": "userIdValue",
		"guestId": "guestUserIdValue",
		"createdAt": "timestamp"
	},
	...
	{
		"id": "anotherIdValue2",
		"url": "http://vvv.vvv",
		"type": "myTypeName",
		"targetId": "someTargetIdValue",
		"userId": "userIdValue",
		"guestId": "guestUserIdValue",
		"createdAt": "timestamp"
	}
]
```