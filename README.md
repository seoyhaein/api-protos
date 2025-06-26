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
~~- dev container 라서, ~/.bashrc 에서 BUF_TOKEN 설정해놓아도, 지속되지 않는다. 이거 버그인데, 시간날때 처리하자.~~    
- git commit 과 push 문제, buf 파일 생성 폴더 문제까지 해결했다. 다만 이게 실수가 나올 수 있으니, 문서화 해놓아야 한다. 중요.  
- service code 넣어 두었는데 이것도 태깅 해야 해서 일단 커밋 하지 않고 푸쉬도 않했다. 커밋 버전 생각해서 진행한다.  (중요)

### tagging 방법 (문서화 해야 함. 일단 잊어버리 않기 위해 작성함. 이거 엉키면 고생함.)
- 모노리포 인데, grpc 서비스 별로 분리해서 좀 복잡함.  
- commit 후  
- git tag gen/go/datablock/ichthys/v1.0.0  
- git push origin gen/go/datablock/ichthys/v1.0.0  
- 해당 서비스 및 코드 가져올 때,   
- go get github.com/seoyhaein/api-protos/gen/go/datablock/ichthys@v1.0.0  
 
- 해당 서비스 내용의 변경사항이 있으면 commit 후  
- git tag gen/go/syncfolders/ichthys/v1.0.0  
- git push origin gen/go/syncfolders/ichthys/v1.0.0  
- 또는 git push --tags (모든 tags push)  
- go get github.com/seoyhaein/api-protos/gen/go/syncfolders/ichthys@v1.0.0  

- 외부 공개 아닌 부분들은 즉, 일반적인 변경들은 그냥 커밋해서 푸시하면 됨.  
- https://g.co/gemini/share/81566349de9a 대화 내용 정리 해야함. 

### go mod 적용하는 방법, 반드시 이렇게 해야 하고 잘못하면 꼬임. 약간 트릭을 씀.
- 해당 폴더로 들어가서 "gen/go/tool/ichthys" <- (여기는 pb.go 파일들이 생기는 폴더임) 여기에서 go mod init 을 해줘야 하는데, 모듈이름을 넣어주어야 함.  
- "go mod init github.com/seoyhaein/api-protos/gen/go/tool/ichthys" <- 이런 식으로 해줘야 함. 그럼 go mod 모듈이 디렉토리에 맞게 낳옴.   
- go.mod 파일을 보면, module github.com/seoyhaein/api-protos/gen/go/tool/ichthys <- 이런 식으로 되어 있음. 이게 정상임.  
- 여기서 또 중요한데, 이제 해당 프로젝트에 깃 테그를 달아줘야 함.  
- 루트로 옮겨가서 즉 .git 이 있는 폴더로 이동해서. 커밋하고 태그를 달아줘야 하는데 태그 이름을 잘 해줘야 함.  