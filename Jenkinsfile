@Library('jenkins-library')

def pipeline = new org.docker.AppPipeline(steps: this,
    dockerImageName:        'iroha/matterbridge',
    dockerRegistryCred:     'bot-soramitsu-rw',
    dockerImageTags:        ['master': 'latest', 'origin/master': 'latest'])
pipeline.runPipeline()
