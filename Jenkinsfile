pipeline {
  agent any

  environment {
    IMAGE_NAME = "namduong0606/golang-portal"
    IMAGE_TAG  = "build-${BUILD_NUMBER}"
    FULL_IMAGE = "${IMAGE_NAME}:${IMAGE_TAG}"
    NAMESPACE  = "default"
  }

  stages {

    stage('Checkout') {
      steps {
        checkout scm
      }
    }

    stage('Docker Login') {
      steps {
        withCredentials([
          usernamePassword(
            credentialsId: 'dockerhub',
            usernameVariable: 'DOCKER_USERNAME',
            passwordVariable: 'DOCKER_PASSWORD'
          )
        ]) {
          bat '''
            echo ==== DOCKER LOGIN ====
            echo %DOCKER_PASSWORD% | docker login -u %DOCKER_USERNAME% --password-stdin
          '''
        }
      }
    }

    stage('Build Image') {
      steps {
        bat '''
          echo ==== BUILD IMAGE ====
          docker build -t %FULL_IMAGE% .
        '''
      }
    }

    stage('Push Image') {
      steps {
        bat '''
          echo ==== PUSH IMAGE ====
          docker push %FULL_IMAGE%
        '''
      }
    }

    stage('Deploy') {
      steps {
        withCredentials([
          file(
            credentialsId: 'kubeconfig',
            variable: 'KUBECONFIG'
          )
        ]) {
          bat '''
            echo ==== DEPLOY ====
            kubectl --kubeconfig="%KUBECONFIG%" set image deployment/golang-portal ^
              golang-portal=%FULL_IMAGE% -n %NAMESPACE%
            kubectl --kubeconfig="%KUBECONFIG%" rollout status deployment/golang-portal -n %NAMESPACE%
          '''
        }
      }
    }
  }

  post {
    success {
      echo "✅ Deployed image: ${FULL_IMAGE}"
    }
  }
}
