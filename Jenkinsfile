pipeline {
    environment {
        goPath="/root/go"
        PROJECT = "demo-tools"
        BUILDBRANCH = "master"
        REGISTRY = "registry"
        REPO = "repo"
    }
    agent any
    /*agent {
        label 'docker'
    }*/
    stages {
        stage('Verify packages and dependencies') {
            steps {
                sh '''
                echo "<RAJ> [Entering] Verify packages and dependencies"
                pwd
                echo $$
                '''
                sh '''
                echo "docker path:`type docker`"
                echo "git path:`type git`"
                echo "id:`id`"
                exit 0
                '''
                sh '''
                echo "<RAJ> [Leaving] Verify packages and dependencies"
                pwd
                echo $$
                '''
            }
        }
        stage('Build Golang'){
            steps{
                sh '''
                    echo "[RAJ] [Entering] Build Golang"
                    echo "Install and build Golang"
                    pwd
                    type docker
                    mkdir .tmp
                    cd .tmp
                    curl -O https://storage.googleapis.com/golang/go1.11.2.linux-amd64.tar.gz >> /dev/null 2>&1
                    tar -xvf go1.11.2.linux-amd64.tar.gz >> /dev/null 2>&1
                    rm -rf /usr/local/go
                    mv go /usr/local
                    ls -lrt /usr/local/go
                    rm go1.11.2.linux-amd64.tar.gz
                    cd ../
                    rmdir .tmp
                    export PATH=$PATH:/usr/local/go/bin
                    cd /root/
                    mkdir -p go/src go/bin go/pkg
                    export GOPATH=/root/go
                    cd $GOPATH/src
                    type go
                    go version
                    go env
                    echo $GOPATH
                    pwd
                    echo "[RAJ] [Leaving] Build Golang"
                '''
           }
        }
        stage('Checkout Source Code') {
            steps {
                sh '''
                echo "<RAJ> [Entering] Checkout Source Code"
                pwd
                ls -lrt
                echo $$
                '''
                // https://rajesh81sekaran@bitbucket.org/rajesh81sekaran/demo-tools.git
                git branch: "${BUILDBRANCH}", credentialsId: 'raj-bitbucket-credentials', url: 'https://rajesh81sekaran@bitbucket.org/rajesh81sekaran/demo-tools.git'
                script {
                    commit = sh(returnStdout: true, script: "git log -n 1 --pretty=format:'%h'").trim()
                    tag = "DEV-SNAPSHOT-${commit}"
                    imageName = "${REGISTRY}/${PROJECT}/${REPO}:${tag}"
                }
                sh "echo Image:${imageName}"
                sh '''
                echo "<RAJ> [Leaving] Checkout Source Code"
                pwd
                ls -lrt
                echo $$
                '''
            }
        }
        /*stage('build-demo-tools') {
            steps {
                sh '''
                docker build -t "demo-tools:1.0.0" --file ./bld .
                '''
            }
        }
        stage('run-demo-tools') {
            steps {
                sh '''
                docker run -v $PWD/dst:/go/src/demo-tools/dst demo-tools:1.0.0
                '''
            }
        }
        stage('build-release-tools') {
            steps {
                sh '''
                docker build -t "release-tools:1.0.0" --file ./rel .
                '''
            }
        }
        stage('run-release-tools') {
            steps {
                sh '''
                docker run release-tools:1.0.0
                '''
            }
        }
        stage('verify docker env') {
            steps {
                sh '''
                docker images
                docker ps -a
                '''
            }
        }
        stage('many scripts') {
            steps {
                sh '''
                echo "many scripts Line 1"
                echo "many scripts Line 2"
                echo "many scripts Line 3"
                type docker
                exit 0
                '''
            }
        }*/
    }
    post {
        always {
            sh '''
            echo "<RAJ> [Entering] post"
            pwd
            echo $$
            '''
            echo '[post] This will always run'
            echo '[post] Before cleanWs()'
            //cleanWs()
            echo '[post] After cleanWs()'
            sh '''
            echo "<RAJ> [Leaving] post"
            pwd
            echo $$
            '''
        }
        success {
            echo '[post] Successful'
        }
        failure {
            echo '[post] Failed'
        }
        unstable {
            echo '[post] Unstable'
        }
        changed {
            echo '[post] Pipeline changed'
        }
    }
}