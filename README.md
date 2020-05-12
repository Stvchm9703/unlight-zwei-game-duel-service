# unlight-zwei-account-service
 unlight-zwei account service




## Remark 

Hook-type order 
[ Before -> Proxy -> After ]

Hook-Type :
1. Before 
   - able add the effect-node

2. Proxy 
   - not effect-node exec

3. After
   - able add the effect-node



### generate doc

MD : 
```bash
protoc -I proto/*.proto --doc_out=./doc --doc_opt=markdown,service.md proto/service.proto

protoc -I proto/*.proto --doc_out=./doc --doc_opt=markdown,message.md proto/message.proto

```

UML 
```
protodot -output doc\ -config doc\protodot.conf.json -src proto\Data.proto    


```