pipeline {
  agent any

  parameters {
    string(name: 'TAG', defaultValue: 'latest', description: 'Docker image tag')
    booleanParam(name: 'MINIKUBE', defaultValue: false, description: 'Use Minikube Docker daemon')
    booleanParam(name: 'SKIP_BUILD', defaultValue: false, description: 'Skip build & push')
    string(name: 'NAMESPACE', defaultValue: 'default', description: 'K8s namespace')
    string(name: 'BRANCH', defaultValue: 'main', description: 'Git branch')
  }

  environment {
    IMAGE = "namduong0606/golang-portal"
    FULL_IMAGE = "${IMAGE}:${params.TAG}"
  }

  stages {

    /* ================= CHECKOUT ================= */
    stage('Checkout') {
      steps {
        echo "Checkout branch: ${params.BRANCH}"
        checkout([
          $class: 'GitSCM',
          branches: [[name: "*/${params.BRANCH}"]],
          userRemoteConfigs: [[url: 'https://github.com/your-org/your-repo.git']]
        ])
      }
    }

    /* ================= ENV CHECK ================= */
    stage('Check Environment') {
      steps {
        bat '''
          echo ==== CHECK ENV ====
          where docker
          where kubectl
          where git
          docker version
          kubectl version --client
        '''
      }
    }

    /* ================= DOCKER LOGIN ================= */
    stage('Docker Login') {
      when {
        allOf {
          expression { !params.MINIKUBE }
          expression { !params.SKIP_BUILD }
        }
      }
      steps {
        withCredentials([usernamePassword(
          credentialsId: 'dockerhub-creds',
          usernameVariable: 'DOCKER_USERNAME',
          passwordVariable: 'DOCKER_PASSWORD'
        )]) {
          bat 'echo %DOCKER_PASSWORD% | docker login -u %DOCKER_USERNAME% --password-stdin'
        }
      }
    }

    /* ================= BUILD IMAGE ================= */
    stage('Build Image') {
      when {
        expression { !params.SKIP_BUILD }
      }
      steps {
        bat '''
          IF "%MINIKUBE%"=="true" (
            echo Using Minikube docker-env
            FOR /f "tokens=*" %%i IN ('minikube docker-env --shell cmd') DO %%i
          )

          docker build -t %FULL_IMAGE% .
        '''
      }
    }

    /* ================= PUSH IMAGE ================= */
    stage('Push Image') {
      when {
        allOf {
          expression { !params.MINIKUBE }
          expression { !params.SKIP_BUILD }
        }
      }
      steps {
        bat 'docker push %FULL_IMAGE%'
      }
    }

    /* ================= DEPLOY ================= */
    stage('Deploy') {
      steps {
        withCredentials([file(credentialsId: 'kubeconfig', variable: 'KUBECONFIG_PATH')]) {
          bat '''
            kubectl --kubeconfig="%KUBECONFIG_PATH%" \
              set image deployment/golang-portal golang-portal=%FULL_IMAGE% -n %NAMESPACE% || (
              kubectl --kubeconfig="%KUBECONFIG_PATH%" \
                rollout restart deployment/golang-portal -n %NAMESPACE%
            )

            kubectl --kubeconfig="%KUBECONFIG_PATH%" \
              rollout status deployment/golang-portal -n %NAMESPACE%
          '''
        }
      }
    }
  }

  post {
    success {
      echo "✅ Deploy success: ${FULL_IMAGE}"
    }
    failure {
      echo "❌ Pipeline failed"
    }
  }
}
