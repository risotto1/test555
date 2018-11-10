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
	sh "ls -la"
	sh "go test ./... -mod=vendor"
      }
    }
    stage("Build") {
      steps {
	echo "starting build"
	sh "ls"
	script {
	  def clientImage = docker.build("risla8/gateway:${GIT_COMMIT}", ". -f ./deployment/docker/Dockerfile.client")
	  clientImage.push()		  
	}
	// sh """
        //   docker build -t risla8/gateway . -f deployment/docker/Dockerfile.client
        //   docker tag risla8/gateway risla8/gateway:${GIT_COMMIT}
        //   docker push risla8/gateway:${GIT_COMMIT}
	// """
      }
    }
  }
}