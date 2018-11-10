pipeline {
  agent any

  options {
    timestamps()
  }

  environment {
    GOCACHE = "$WORKSPACE"
  }

  stages {
    stage("Tests") {
      agent {
	docker { image "golang:alpine" }
      }
      environment {
	CGO_ENABLED = "0"
      }
      steps {
	sh "go test ./... -mod=vendor"
      }
    }
    stage("Build") {
      steps {
	script {
	  def clientImage = docker.build("risla8/client:${GIT_COMMIT}", ". -f ./deployment/docker/Dockerfile.client")
	  def serverImage = docker.build("risla8/server:${GIT_COMMIT}", ". -f ./deployment/docker/Dockerfile.server")
	  docker.withRegistry("https://index.docker.io/v1/", "9f00e117-89d7-4ec6-afb9-d4c415878fa2") {
	    clientImage.push()		  
	    clientImage.push("latest")		  
	    serverImage.push()
	    serverImage.push("latest")
	  }
	}
      }
    }
  }
}