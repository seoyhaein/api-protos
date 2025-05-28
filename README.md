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
- F1 누르면 buf lint 실행.

- check https://buf.build/seoyhaein

### TODO
- 일단 마무리 하고, 이제 실제적인 코딩 작업에 들어가자.  
- 에러 체크와 만들어진 pb.go 파일들 확인하면서 마무리 하자.  