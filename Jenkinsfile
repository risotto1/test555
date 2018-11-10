pipeline {
  agent any

  options {
    timestamps()
  }

  environment {
    GOCACHE = "$WORKSPACE"
  }

  stages {
    stage("Test") {
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
	  docker.withRegistry("https://index.docker.io/v1/", "6f769d37-4183-46c8-80ad-c30bdcab6f02") {
	    clientImage.push()		  
	    clientImage.push("latest")		  
	    serverImage.push()
	    serverImage.push("latest")
	  }
	}
      }
    }

    stage("Deploy staging") {
      steps {
	sh "helm upgrade --install staging -f ./deployment/k8s/base.yaml -f ./deployment/k8s/values-staging.yaml ./deployment/k8s/app --set crud.image.tag=${GIT_COMMIT} --set gateway.image.tag=${GIT_COMMIT}"
      }
    }

    stage("Deploy production") {
      steps {
	milestone label: "Deploy to production?", ordinal: Integer.parseInt(env.BUILD_ID)
	echo "mm"
      }
    }
  }
}