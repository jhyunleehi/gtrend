# 개발환경 설정
### 1. vscode 설치 

* vs code 다운로드

[https://code.visualstudio.com/download](https://code.visualstudio.com/download)


### 2. go lang 설치 
* go 언어

>GO 프로그래밍 언어는 2007년 구글에서 개발을 시작하여 2012년 GO 버젼 1.0을 완성하였다. GO는 이후 계속 향상된 버젼을 내 놓았으며 2022년 초에는 버젼 1.18 에 이르렀다.

> GO 공식 사이트: https://go.dev/ (혹은 https://golang.org)

>흔히 golang 이라고도 불리우는 Go 프로그래밍 언어는 구글의 V8 Javascript 엔진 개발에 참여했던 Robert Griesemer, Bell Labs에서 유닉스 개발에 참여했던 Rob Pike, 그리고 역시 Bell Labs에서 유닉스 개발했으며 C 언어의 전신인 B 언어를 개발했던 Ken Thompson이 함께 개발하였다.

>Go는 전통적인 컴파일, 링크 모델을 따르는 범용 프로그래밍 언어이다. Go는 일차적으로 시스템 프로그래밍을 위해 개발되었으며, C++, Java, Python의 장점들을 뽑아 만들어졌다. C++와 같이 Go는 컴파일러를 통해 컴파일되며, 정적 타입 (Statically Typed)의 언어이다. 또한 Java와 같이 Go는 Garbage Collection 기능을 제공한다. Go는 단순하고 간결한 프로그래밍 언어를 지향하였는데, Java의 절반에 해당하는 25개의 키워드만으로 프로그래밍이 가능하게 하였다. 마지막으로 Go의 큰 특징으로 Go는 Communicating Sequential Processes (CSP) 스타일의 Concurrent 프로그래밍을 지원한다.


#### go 다운로드
[](https://go.dev/dl/go1.18.3.windows-amd64.msi)
##### go 설치 위치
````
c:\Go
````

##### 환경 설정 

* 환경 변수 설정 

```sh
D:\Code\lk>set
GOPATH=C:\Gocode
GOROOT=C:\Go
```
Path 설정에  `C:\Gocode\bin;C:\Go\bin` 추가 





## go 설치

#### go package 설치 

```sh
C:\Gocode\src\gtrend>go mod init 
C:\Gocode\src\gtrend>go get 
go: downloading github.com/go-echarts/go-echarts/v2 v2.2.4
go: downloading github.com/go-echarts/go-echarts v1.0.0
go get: added github.com/go-echarts/go-echarts/v2 v2.2.4
C:\Gocode\src\gtrend>go mod init
go: C:\Gocode\src\gtrend\go.mod already exists
C:\Gocode\src\gtrend>go mod vendor
```


#### git 설정

* git clone
```sh
C:\Gocode\src\gtrend>git clone https://github.com/jhyunleehi/gtrend.git 
Cloning into 'gtrend'...
remote: Enumerating objects: 72, done.
remote: Counting objects: 100% (72/72), done.
remote: Compressing objects: 100% (44/44), done.
remote: Total 72 (delta 20), reused 61 (delta 20), pack-reused 0
Unpacking objects: 100% (72/72), done.
```


* git 초기화 설정이 필요할 경우 
````
C:\Gocode\src\gtrend>git add .
C:\Gocode\src\gtrend>git commit -m "first commit"
C:\Gocode\src\gtrend>git push  -u --force origin master
C:\Gocode\src\gtrend>git log
C:\Gocode\src\gtrend>git  remote  add  origin https://github.com/jhyunleehi/gtrend.git 
C:\Gocode\src\gtrend>git remote -v
origin  https://github.com/jhyunleehi/gtrend.git (fetch)
origin  https://github.com/jhyunleehi/gtrend.git (push)
```

* git checkout branch, push
```
C:\Gocode\src\gtrend>git checkout -b develop
C:\Gocode\src\gtrend>git push --set-upstream origin develop
Total 0 (delta 0), reused 0 (delta 0)
remote: 
remote: Create a pull request for 'develop' on GitHub by visiting:
remote:      https://github.com/jhyunleehi/gtrend/pull/new/develop
remote:
To https://github.com/jhyunleehi/gtrend.git
 * [new branch]      develop -> develop
Branch 'develop' set up to track remote branch 'develop' from 'origin'.
```
