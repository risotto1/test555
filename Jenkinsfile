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
	GOOS = "linux"
	CGO_ENABLED = "0"
      }
      steps {
	sh "ls -la"
	sh "mkdir $WORKSPACE/.cache"
	sh "go test ./... -mod=vendor"
      }
    }
    stage("Build") {
      steps {
	sh """
          docker build -t risla8/gateway . -f deployments/docker/Dockerfile.client
          docker tag risla8/gateway risla8/gateway:${GIT_COMMIT}
          docker push risla8/gateway:${GIT_COMMIT}
	"""
      }
    }
  }
}