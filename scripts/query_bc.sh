echo "query Block by blockNumber"
echo
Block=$(curl -s -X GET \
  http://localhost:8080/api/v1/blockchain/queryblock/2)
echo "$Block"
echo
echo

echo "query Transaction by TransactionID"
echo
Transaction=$(curl -s -X GET \
  http://localhost:8080/api/v1/blockchain/querytrx/6f0550c2c9eacdcb2395f35b969f6f59e7f42192147aa11a70e09eaef978b3b5)
echo "$Transaction"
echo
echo

echo "Query ChainInfo"
echo
ChainInfo=$(curl -s -X GET \
  http://localhost:8080/api/v1/blockchain/querychainInfo)
echo "$ChainInfo"
echo
echo

echo "Query Installed"
echo
Installed=$(curl -s -X GET \
  http://localhost:8080/api/v1/blockchain/queryinstalled)
echo "$Installed"
echo
echo

echo "Query Instantiated"
echo
Instantiated=$(curl -s -X GET \
  http://localhost:8080/api/v1/blockchain/queryinstantiated)
echo "$Instantiated"
echo
echo

echo "Query Channels"
echo
Channels=$(curl -s -X GET \
  http://localhost:8080/api/v1/blockchain/querychannels)
echo "$Channels"
echo
echo
