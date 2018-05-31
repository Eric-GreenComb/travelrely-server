echo "POST create order"
echo
ORDER=$(curl -s -X POST \
  http://localhost:8080/api/v1/orders/create \
  -H "content-type: application/json" \
  -d '{
	"orgID":1001,
	"orderSN":"SN20180110002",
    "amount":1025
}')
echo "$ORDER"
echo
echo

# Create AttachFile post
echo "POST Create AttachFile"
echo
AttachFile=$(curl -s -X POST \
  http://localhost:8080/api/v1/attachfile/create \
  -H "content-type: application/json" \
  -d '{
	"name":"传奇人物",
	"desc":"传奇人物-华盛顿",
  "path":"http://fifu.com/upload/1.pdf",
  "hash":"1484e17809cdec21e9a2582487602720"
}')
echo "$AttachFile"
echo
echo

## Update AttachFile post
echo "POST Update AttachFile"
echo
AttachFile=$(curl -s -X POST \
  http://localhost:8080/api/v1/attachfile/update/8d1f65547e284dadb4e42372fc879985 \
  -H "content-type: application/json" \
  -d '{
    "uuid":"2a5b15ec08104e56bcf41e5845d2ab1f",
	"name":"传奇1",
	"desc":"传奇人物",
  "path":"http://fifu.com/upload/1.pdf",
  "hash":"1484e17809cdec21e9a2582487602720"
}')
echo "$AttachFile"
echo
echo

# Create Org post
echo "POST Create Org"
echo
Org=$(curl -s -X POST \
  http://localhost:8080/api/v1/org/create \
  -H "content-type: application/json" \
  -d '{
  "uscc":"199217348701324789",  
	"name":"绿蜂巢",
	"licence":"124123",
  "orgCode":"2222222222",
  "taxregCode":"8678678"
}')
echo "$Org"
echo
echo

## Update Org post
echo "POST Update Org"
echo
Org=$(curl -s -X POST \
  http://localhost:8080/api/v1/org/update/199217348701324789 \
  -H "content-type: application/json" \
  -d '{
    "uscc":"199217348701324789",
	"name":"绿蜂巢1",
	"licence":"124123",
  "orgCode":"2222222222",
  "taxregCode":"8678678"
}')
echo "$Org"
echo
echo