# api-protos
- module 확인

```
buf config ls-modules --format path
```

- buf build push

```
# 기본 라벨(v1)로 푸시할 때 buf build 에 push 할때,
./push_modules.sh

# 다른 라벨(예: v2)을 사용하고 buf build 에 push 할때,
./push_modules.sh v2

# 이후 기본 라벨에서 업데이트 할때는 
buf push

```


- check https://buf.build/seoyhaein