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
	  docker.withRegistry("https://index.docker.io/v1/", "36c75af7-2bbc-4897-a668-469f126da32b") {
	    clientImage.push()		  
	    clientImage.push("latest")		  
	    serverImage.push()
	    serverImage.push("latest")
	  }
	}
      }
    }

    stage("Deploy to staging") {
      steps {
	sh "helm upgrade --install staging -n staging -f ./deployment/k8s/base.yaml -f ./deployment/k8s/values-staging.yaml ./deployment/k8s/app --set crud.image.tag=${GIT_COMMIT} --set gateway.image.tag=${GIT_COMMIT}"
      }
    }
  }
}