# travelrely-server

ticketchain server

## health

curl http://115.28.151.227:8080/health

## 用户管理

- api/v1/user/register 用户注册

curl -H "Content-Type:application/json" -X POST --data '{"cmd":"user_register","userID":"eric"}' http://115.28.151.227:8080/api/v1/user/register

- api/v1/user/unregister  注销用户

curl -H "Content-Type:application/json" -X POST --data '{"cmd":"user_unregister",
 "bcuser_id":"eric",
 "bcuser_key":"private"}' http://115.28.151.227:8080/api/v1/user/unregister

## msisdn

- POST api/v1/msisdn/subscribe 签约

curl -H "Content-Type:application/json" -X POST --data '{"cmd":"subscribe","bcuser_id":"bcuser_id","bcuser_key":"bcuser_key","msisdn":"13810167616","eki2":"eki2"}' http://115.28.151.227:8080/api/v1/msisdn/subscribe

- POST api/v1/msisdn/unsubscribe 解约

curl -H "Content-Type:application/json" -X POST --data '{"cmd":"user_unregister","bcuser_id":"bcuser_id","bcuser_key":"bcuser_key","msisdn":"13810167616","asset_id":"uuid1234"}' http://115.28.151.227:8080/api/v1/msisdn/unsubscribe

- api/v1/msisdn/state/:msisdn 查询号码签约信息

curl http://115.28.151.227:8080/api/v1/msisdn/state/13810167616

- api/v1/msisdn/history/:msisdn 查询号码签约历史信息

curl http://115.28.151.227:8080/api/v1/msisdn/history/13810167616

- api/v1/asset/info/:asset 查询号码签约历史信息

curl http://115.28.151.227:8080/api/v1/asset/info/uuid1234

## transfer

- /transfer/query 查询

curl -H "Content-Type:application/json" -X POST --data '{"cmd":"transt_list","tm":1527916756,"start":0,"max_no":10}' http://115.28.151.227:8080/api/v1/transfer/query

## block

- /channels/:channel/height 高度

curl http://115.28.151.227:8080/api/v1/channels/mychannel/height