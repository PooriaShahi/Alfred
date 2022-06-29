package templates

func JavaSpringDockerfile(jarFileName string) string {
	tmpl := `FROM maven:3.6.3-openjdk-8 as build
WORKDIR /develop
COPY . .
ADD settings.xml /root/.m2/settings.xml
RUN mvn package -DskipTests

# production environment
FROM store/oracle/serverjre:1.8.0_241-b07
WORKDIR /app
COPY --from=build /develop/target/` + jarFileName + ` /app
CMD ["java", "-jar", "/app/` + jarFileName + `]`
	return tmpl
}

func JavaSpringDockerCompose(projectName, desiredPort, appPort string) string {
	tmpl := `version: '3.8'

	services:
	  ` + projectName + `:
		container_name: ` + projectName + `
		image: ` + projectName + `:${TAG_NAME}
		env_file:
		  - variables.env
		ports:
		  - "127.0.0.1:` + desiredPort + `:` + appPort + `
		hostname: '` + projectName + `'
		networks:
		  - ${DEPLOY_NAME}
		ulimits:
		  nproc: 65535
		  nofile: 
			soft: 65535
			hard: 65535
	
	networks:
	  dev:
		name: dev
	  stage:
		name: stage
	  prod:
		name: prod`
	return tmpl
}
