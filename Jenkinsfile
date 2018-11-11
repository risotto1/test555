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

    stage("Build images") {
      steps {
	script {
	  def clientImage = docker.build("risla8/client:${GIT_COMMIT}", ". -f ./deployment/docker/Dockerfile.client")
	  def serverImage = docker.build("risla8/server:${GIT_COMMIT}", ". -f ./deployment/docker/Dockerfile.server")
	}
      }
    }

    stage("Push images to hub") {
      steps {
	script {
	  docker.withRegistry("https://index.docker.io/v1/", "6f769d37-4183-46c8-80ad-c30bdcab6f02") {
	    clientImage.push()		  
	    clientImage.push("latest")		  
	    serverImage.push()
	    serverImage.push("latest")
	  }
	}
      }
    }

    stage("Deploying to staging") {
      environment {
	BASE = "./deployment/k8s/base.yaml"
	STAGING = "./deployment/k8s/values-staging.yaml"
	CHART = "./deployment/k8s/app"
      }
      steps {
	sh "helm upgrade --install staging -n staging -f ${BASE} -f ${STAGING} ${CHART} --set crud.image.tag=${GIT_COMMIT} --set gateway.image.tag=${GIT_COMMIT}"
      }
    }

    stage("Deploy to production input") {
      steps {
	milestone(1)
	input "Deploy to production?"
	milestone(2)
      }
    }

    stage("Deploying to production") {
      steps {
	echo "done"
      }
    }
  }
}