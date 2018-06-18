package main

var (
	simpleConfig = MavenConfig{
		RepoUrl:    "https://repo.maven.apache.org/maven2/",
		GroupId:    "com.group.my",
		ArtifactId: "my-artifact",
		Version:    "1.0.0",
		Extension:  "jar",
	}
	classifierConfig = MavenConfig{
		RepoUrl:    "https://repo.maven.apache.org/maven2/",
		GroupId:    "com.group.my",
		ArtifactId: "my-artifact",
		Version:    "1.0.0",
		Classifier: "javadoc",
		Extension:  "jar",
	}
	zipConfig = MavenConfig{
		RepoUrl:    "https://repo.maven.apache.org/maven2/",
		GroupId:    "com.group.my",
		ArtifactId: "my-artifact",
		Version:    "1.0.0",
		Extension:  "zip",
	}
	allMavenKeysConfig = MavenConfig{
		RepoUrl:      "https://repo.maven.apache.org/maven2/",
		GroupId:      "com.group.my",
		ArtifactId:   "my-artifact",
		Version:      "1.0.0",
		Classifier:   "source",
		Extension:    "zip",
		RepoUsername: "user",
		RepoPassword: "password",
	}
)
