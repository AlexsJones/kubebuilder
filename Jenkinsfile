def projectName = "kubebuilder"
node('k8s')  {
  stage("Fetch") {
    git branch: "${env.BRANCH_NAME}", credentialsId: '54213d77-d44e-49ce-a635-4a2bef7fdf71', url: "git@github.com:AlexsJones/"+ projectName + ".git"
    version = sh returnStdout: true, script: 'git rev-parse --short HEAD | tr -d "\n"'
    version = version + "-${env.BUILD_NUMBER}"
    slackSend channel: "#dev-null", color: '#ffff00', message: "Building preview ${env.JOB_NAME} $version (<${env.BUILD_URL}|Open>)"
  }
  stage("Build") {

    withCredentials([file(credentialsId: 'PREVIEW', variable: 'PREVIEW')]) {
      sh '/opt/google-cloud-sdk/bin/gcloud auth activate-service-account --key-file $PREVIEW'
    }
    ret = sh returnStatus: true, script: "/opt/google-cloud-sdk/bin/gcloud docker -- build -t " + projectName + ":$version ."
    if(ret != 0) {
      slackSend channel: "#dev-null", color: '#E11B1B', message: "Docker build in preview failed! ${env.JOB_NAME} $version (<${env.BUILD_URL}|Open>)"
      return
    }
  }
}
