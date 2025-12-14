pipeline {
  agent any

  parameters {
    string(name: 'TAG', defaultValue: 'latest', description: 'Docker image tag (e.g., ci-123, v1.0.0)')
    booleanParam(name: 'MINIKUBE', defaultValue: false, description: 'Build inside Minikube daemon (no push to registry)')
    string(name: 'NAMESPACE', defaultValue: 'default', description: 'Kubernetes namespace for deployment')
    booleanParam(name: 'SKIP_BUILD', defaultValue: false, description: 'Skip build, deploy only')
  }

  environment {
    IMAGE = "namduong0606/golang-portal"
    FULL_IMAGE = "${IMAGE}:${params.TAG}"
  }

  stages {
    stage('Checkout') {
      steps {
        echo "Checking out from GitHub..."
        checkout scm
      }
    }

    stage('Check Environment') {
      steps {
        echo "Verifying required tools..."
        bat '''
          echo ==== CHECK ENV ====
          where make
          where bash
          where docker
          where kubectl
          git --version
        '''
      }
    }

    stage('Build & Push') {
      when {
        expression { !params.SKIP_BUILD }
      }
      steps {
        echo "Building and pushing image: ${FULL_IMAGE}"
        withCredentials([usernamePassword(credentialsId: 'docker-hub-creds', usernameVariable: 'DOCKER_USERNAME', passwordVariable: 'DOCKER_PASSWORD')]) {
          bat '''
            echo %DOCKER_PASSWORD% | docker login -u %DOCKER_USERNAME% --password-stdin
          '''

          script {
            // Gọi bash với file trực tiếp (không cần chmod)
            if (params.MINIKUBE) {
              bat "\"C:\\Program Files\\Git\\bin\\bash.exe\" ./run.sh --tag ${params.TAG} --minikube --namespace ${params.NAMESPACE}"
            } else {
              bat "\"C:\\Program Files\\Git\\bin\\bash.exe\" ./run.sh --tag ${params.TAG} --namespace ${params.NAMESPACE}"
            }
          }
        }
      }
    }

    stage('Deploy') {
      steps {
        echo "Deploying to namespace: ${params.NAMESPACE}"
        withCredentials([file(credentialsId: 'kubeconfig', variable: 'KUBECONFIG_PATH')]) {
          bat '''
            setlocal enabledelayedexpansion
            kubectl --kubeconfig="%KUBECONFIG_PATH%" set image deployment/golang-portal golang-portal=%FULL_IMAGE% -n %NAMESPACE% || (
              echo set-image failed, trying rollout restart
              kubectl --kubeconfig="%KUBECONFIG_PATH%" rollout restart deployment/golang-portal -n %NAMESPACE%
            )
            kubectl --kubeconfig="%KUBECONFIG_PATH%" rollout status deployment/golang-portal -n %NAMESPACE%
            kubectl --kubeconfig="%KUBECONFIG_PATH%" get pods -l app=golang-portal -n %NAMESPACE% -o wide
          '''
        }
      }
    }
  }

  post {
    always {
      echo "Pipeline execution completed."
    }
    failure {
      echo "Pipeline failed! Check console output above."
    }
    success {
      echo "Pipeline succeeded! Deployment should be rolling out."
    }
  }
}
