// Jenkinsfile - Manual CI/CD Pipeline for golang-portal
// Job type: Pipeline (NOT Freestyle)
// This requires:
// - Jenkins agent with docker, kubectl, bash
// - Credential IDs: docker-hub-creds (username+password), kubeconfig (file)
// - Environment variables or parameters: TAG (default: build-${BUILD_NUMBER})

pipeline {
  agent any

  parameters {
    string(name: 'TAG', defaultValue: '', description: 'Image tag (default: build-BUILD_NUMBER)')
    string(name: 'NAMESPACE', defaultValue: 'default', description: 'K8s namespace')
  }

  environment {
    IMAGE_NAME = "namduong0606/golang-portal"
    IMAGE_TAG = "${params.TAG ?: 'build-' + env.BUILD_NUMBER}"
    FULL_IMAGE = "${IMAGE_NAME}:${IMAGE_TAG}"
    DOCKER_CRED_ID = "docker-hub-creds"
    KUBECONFIG_CRED_ID = "kubeconfig"
  }

  stages {
    stage('Checkout') {
      steps {
        echo "🔄 Checking out from GitHub..."
        checkout scm
      }
    }

    stage('Check Environment') {
      steps {
        echo "🔍 Verifying tools..."
        bat '''
          echo ==== TOOLS CHECK ====
          where docker
          where kubectl
          where git
          docker version --format "Docker: {{.Client.Version}}"
        '''
      }
    }

    stage('Build Image') {
      steps {
        echo "🔨 Building image: ${FULL_IMAGE}"
        bat '''
          setlocal enabledelayedexpansion
          echo ==== BUILD IMAGE ====
          docker build -t %FULL_IMAGE% .
        '''
      }
    }

    stage('Push Image') {
      steps {
        echo "📤 Pushing image: ${FULL_IMAGE}"
        withCredentials([usernamePassword(credentialsId: params.DOCKER_CREDENTIALS, usernameVariable: 'DOCKER_USERNAME', passwordVariable: 'DOCKER_PASSWORD')]) {
          bat '''
            setlocal enabledelayedexpansion
            echo ==== DOCKER LOGIN ====
            echo %DOCKER_PASSWORD% | docker login -u %DOCKER_USERNAME% --password-stdin
            echo ==== PUSH IMAGE ====
            docker push %FULL_IMAGE%
          '''
        }
      }
    }

    stage('Deploy') {
      steps {
        echo "🚀 Deploying to namespace: ${params.NAMESPACE}"
        withCredentials([file(credentialsId: 'kubeconfig', variable: 'KUBECONFIG_FILE')]) {
          bat '''
            setlocal enabledelayedexpansion
            echo ==== DEPLOY ====
            set KUBECONFIG=%KUBECONFIG_FILE%
            kubectl set image deployment/golang-portal golang-portal=%FULL_IMAGE% -n %NAMESPACE% || (
              echo set-image failed, trying rollout restart
              kubectl rollout restart deployment/golang-portal -n %NAMESPACE%
            )
            echo ==== ROLLOUT STATUS ====
            kubectl rollout status deployment/golang-portal -n %NAMESPACE%
            echo ==== POD STATUS ====
            kubectl get pods -l app=golang-portal -n %NAMESPACE% -o wide
          '''
        }
      }
    }
  }

  post {
    success {
      echo "✅ Pipeline succeeded! Image ${FULL_IMAGE} deployed to namespace: ${params.NAMESPACE}"
    }
    failure {
      echo "❌ Pipeline failed. Check logs above."
    }
  }
}
