## User

### CreateUser
echo "POST create user"
echo
USER=$(curl -s -X POST \
  http://localhost:8080/api/v1/users/create \
  --data "userID=eric&passwd=a11111&name=Eric&email=ministor@126.com")
echo "$USER"
echo
echo

### Login 
echo "POST login"
echo
TOKEN=$(curl -s -X POST \
  http://localhost:8080/api/v1/login \
  --data "userID=eric&passwd=0fb6c6c0b7621fb7bd6ff1e6fb656bc746e2254a4f671dee25c0ce3ddd9ccf3e")
echo "$TOKEN"
echo
echo

### Auth
echo "POST auth"
echo
Auth=$(curl -s -X POST \
  http://localhost:8080/api/v1/auth \
  --data "token=4a16bb38f8d5a9441d90edf462c5b8")
echo "$Auth"
echo
echo

### Logout
echo "POST logout"
echo
Logout=$(curl -s -X POST \
  http://localhost:8080/api/v1/logout \
  --data "token=a4158dd9293e98d6398642857a395f")
echo "$Logout"
echo
echo

### UserInfo
echo "GET /users/info"
echo
UserInfo=$(curl -s -X GET \
  http://localhost:8080/api/v1/users/info?token=f1c751ec357075d117a10480369be4)
echo "$UserInfo"
echo
echo

### UpdateUser
echo "POST UpdateUser"
echo
UpdateUser=$(curl -s -X PUT \
  http://localhost:8080/api/v1/users/update \
  --data "token=4a16bb38f8d5a9441d90edf462c5b8&name=Ministor&email=ministor@126.com")
echo "$UpdateUser"
echo
echo

### UpdateUserPasswd
echo "POST UpdateUserPasswd"
echo
UpdateUserPasswd=$(curl -s -X PUT \
  http://localhost:8080/api/v1/users/updatepasswd \
  --data "token=4a16bb38f8d5a9441d90edf462c5b8&expasswd=9cf93debe36288625002f7476dad6459ec23a1dd5c8db619f59dfad747b0e31c&newpasswd=0fb6c6c0b7621fb7bd6ff1e6fb656bc746e2254a4f671dee25c0ce3ddd9ccf3e")
echo "$UpdateUserPasswd"
echo
echo

## Create AttachFile post

echo "POST Create AttachFile"
echo
AttachFile=$(curl -s -X POST \
  http://localhost:8080/api/v1/attachfile/create \
  -H "content-type: application/json" \
  -d '{
	"name":"传奇",
	"desc":"传奇人物",
  "path":"http://fifu.com/upload/1.pdf",
  "hash":"1484e17809cdec21e9a2582487602720"
}')
echo "$AttachFile"
echo
echo

return txid:6f0550c2c9eacdcb2395f35b969f6f59e7f42192147aa11a70e09eaef978b3b5


UUID:a82e8761606c4bb79804befa126a49a3

## contract post

echo "POST create contract"
echo
CONTRACT=$(curl -s -X POST \
  http://localhost:8080/api/v1/contracts/create \
  -H "content-type: application/json" \
  -d '{
	"uscc":"1001",
	"orderSN":"SN20180110002",
  "amount":1025
}')
echo "$CONTRACT"
echo
echo

