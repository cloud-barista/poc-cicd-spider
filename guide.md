(작성중) 클라우드바리스타 CI/CD 가이드 입니다.

비고: 본 문서는 배포후에 참조(링크)하시기 바랍니다. 문서의 최종 위치가 변경될 수 있습니다.

# Cloud-Barista CI/CD 구축 가이드

## [목 차]

1. [개요](#개요)
2. [GitHub 환경 설정](#GitHub-환경-설정)
3. [Workflow 파일 구성](#Workflow-파일-구성)
4. [Dashboard 구축](#Dashboard-구축)
5. [Unit Test 파일 구성](#Unit-Test-파일-구성)
6. [CI Workflow 수정](#CI-Workflow-구축)
7. [CD Workflow 수정](#CD-Workflow-구축)
8. [REPORTS Workflow 수정](#REPORTS-Workflow-구축)
9. [Workflow Job 추가](#Workflow-Job-추가)
10. [Unit Test 시나리오 추가](#Unit-Test-시나리오-추가)

## [개요]

Cloud-Barista 시스템은 다양한 프레임워크로 구성되어 있고, 프레임워크별로 개발 / 운영하고 있다. 기존 프레임워크 또는 새로운 프레임워크에서 CI/CD를 적용할 수 있도록 표준화된 방법을 제시하고, 실제로 구현할 수 있도록 가이드를 제공한다.  
본 가이드에서는 GitHub Actions 를 이용한 Cloud-Barista CI/CD 구축 과정을 [poc-cicd-spider](https://github.com/cloud-barista/poc-cicd-spider/tree/master/.github/workflows) 를 기준으로 [CB-SPIDER](https://github.com/cloud-barista/cb-spider) 에 적용하는 방법을 상세히 소개한다.

## [GitHub 환경 설정]

### (1) Secrets 변수 설정

### (2) Self-hosted Runner 설정

## [Workflow 파일 구성]

### (1) .github 폴더 생성

### (2) .golangci.yaml 파일 복사

### (3) workflows 폴더 생성

### (4) cb-ci-actions.yaml 파일 복사

### (5) cb-cd-actions.yaml 파일 복사

### (6) cb-report.yaml 파일 복사

## [Dashboard 구축]

### (1) Dashboard Repository 생성

### (2) GitHub Pages 설정

### (3) Dashboard App 복사

## [Unit Test 파일 구성]

### (1) Unit Test 시나리오 복사

### (2) Unit Test 실행 방법

### (3) Backend 서버 버전 변경 방법

## [CI Workflow 수정]

### (1) skip_tags 변수 활용

### (2) Lint 수정

### (3) Converage 실행 환경 수정

### (4) go build 버전 추가

## [CD Workflow 수정]

### (1) 이미지 이름 변경

## [REPORTS Workflow 수정]

### (1) Dashboard Repository 변경

## [Workflow Job 추가]

### (1) Echo Job 생성

### (2) Output 폴더 생성 Step 작성

### (3) Echo Step 작성

### (4) Result 를 Artifact 로 업로드

### (5) Artifact 에서 Result 다운로드

### (6) Result 파일 정보 로딩

### (7) Summary Result 통합

### (8) Dashboard Json 통합

### (9) Dashboard App 수정

## [Unit Test 시나리오 추가]

### (1) Rest 시나리오 파일 생성

### (2) Rest 시나리오 환경 구성

### (3) TestCase 생성
