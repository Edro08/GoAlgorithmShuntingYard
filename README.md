# Edro08/GoAlgorithmShuntingYard

### Endpoints
#### Health
```shell
curl --request GET \
  --url http://localhost:90/health
```

#### EnvExpMath
```shell
curl --request POST \
  --url http://localhost:90/GoAlgorithmShuntingYard/v1/envExpMath \
  --header 'Content-Type: application/json' \
  --data '{
	"infix": "3 + 4 * 2 / 1 - 5"
}'
``` 