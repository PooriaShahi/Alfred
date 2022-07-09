package templates

func MavenSettings(mavenId, mavenUrl, username, password string) string {
	tmpl := `<?xml version="1.0" encoding="UTF-8"?>
	<settings xmlns="http://maven.apache.org/SETTINGS/1.0.0"
		 xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
		 xsi:schemaLocation="http://maven.apache.org/SETTINGS/1.0.0 http://maven.apache.org/xsd/settings-1.0.0.xsd">
	
	  <mirrors>
		<mirror>
		  <id>` + mavenId + `</id>
		  <mirrorOf>*</mirrorOf>
		  <url>` + mavenUrl + `/url>
		</mirror>
	  </mirrors>
	
	   <servers>
		 <server>
		   <id>` + mavenId + `</id>
		   <username>` + username + `</username>
		   <password>` + password + `</password>
		 </server>
	   </servers>
	</settings>`
	return tmpl
}
