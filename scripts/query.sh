
echo "Query AttachFile"
echo
AttachFile=$(curl -s -X GET \
  http://localhost:8080/api/v1/attachfile/info/8d1f65547e284dadb4e42372fc879985)
echo "$AttachFile"
echo
echo

echo "Query AttachFile History"
echo
AttachFile=$(curl -s -X GET \
  http://localhost:8080/api/v1/attachfile/history/8d1f65547e284dadb4e42372fc879985)
echo "$AttachFile"
echo
echo

echo "Query Org Obj"
echo
Org=$(curl -s -X GET \
  http://localhost:8080/api/v1/org/obj/199217348701324789)
echo "$Org"
echo
echo

echo "Query Org"
echo
Org=$(curl -s -X GET \
  http://localhost:8080/api/v1/org/info/199217348701324789)
echo "$Org"
echo
echo

echo "Query Org History"
echo
Org=$(curl -s -X GET \
  http://localhost:8080/api/v1/org/history/199217348701324789)
echo "$Org"
echo
echo