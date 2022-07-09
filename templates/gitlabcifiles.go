package templates

import "strings"

func JavaSpringGitlabCi(gitlabProject, prefix string) string {
	strings.ToLower(gitlabProject)
	tmpl := `stages:
	- build
	- deploy

  build-` + strings.ToLower(gitlabProject) + `-dev:
	stage: build
	script:
	  - docker build -f Dockerfile -t ` + strings.ToLower(gitlabProject) + `:${CI_PIPELINE_ID} .
	only:
	  - dev
	tags:
	  - ` + strings.ToLower(gitlabProject) + `-dev
  
  deploy-` + strings.ToLower(gitlabProject) + `-dev:
	stage: deploy
	environment: DEV
	script:
	  - echo "DEPLOY_NAME=dev" >> .env
	  - echo "TAG_NAME=${CI_PIPELINE_ID}" >> .env
	  - env | grep '` + prefix + `_' | sed 's/` + prefix + `_//' >> variables.env
	  - cat .env
	  - docker-compose up -d
	only:
	  - dev
	tags:
	  - ` + strings.ToLower(gitlabProject) + `-dev
  
  build-` + strings.ToLower(gitlabProject) + `-stage:
	stage: build
	script:
	  - echo ${CI_PIPELINE_ID}
	  - docker build -f Dockerfile -t ` + strings.ToLower(gitlabProject) + `:${CI_PIPELINE_ID} .
	only:
	  - stage
	tags:
	  - ` + strings.ToLower(gitlabProject) + `-stage
  
  deploy-` + strings.ToLower(gitlabProject) + `-stage:
	stage: deploy
	environment: STAGE
	script:
	  - echo "REPLICAS=1" >> .env
	  - echo "DEPLOY_NAME=stage" >> .env
	  - echo "TAG_NAME=${CI_PIPELINE_ID}" >> .env
	  - env | grep '` + prefix + `_' | sed 's/` + prefix + `_//' >> variables.env
	  - cat .env
	  - docker-compose up -d
	only:
	  - stage
	tags:
	  - ` + strings.ToLower(gitlabProject) + `-stage
  
  
  build-` + strings.ToLower(gitlabProject) + `-prod:
	stage: build
	script:
	  - echo ${CI_PIPELINE_ID}
	  - docker build -f Dockerfile -t ` + strings.ToLower(gitlabProject) + `:${CI_PIPELINE_ID} .
	only:
	  - master
	tags:
	  - ` + strings.ToLower(gitlabProject) + `-prod
  
  
  deploy-` + strings.ToLower(gitlabProject) + `-prod:
	environment: PROD
	stage: deploy
	script:
	  - echo "REPLICAS=1" >> .env
	  - echo "DEPLOY_NAME=prod" >> .env
	  - echo "TAG_NAME=${CI_PIPELINE_ID}" >> .env
	  - env | grep '` + prefix + `_' | sed 's/` + prefix + `_//' >> variables.env
	  - cat .env
	  - docker-compose up -d
	only:
	  - master
	tags:
	  - ` + strings.ToLower(gitlabProject) + `-prod`
	return tmpl
}
