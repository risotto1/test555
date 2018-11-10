pipeline {
  agent any
  options {
    timestamps()
  }
  
  stages {
    stage("Tests") {
      agent {
	docker { image "golang:alpine" }
      }
      steps {
	sh "go test ./..."
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