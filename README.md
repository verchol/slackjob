# Helm3 hooks example that sends slack message on post-install and post-upgrade hooks

Example of how to use Helm3 hookst to execte customer job 
In our case it's a simple slack hook but it can be Codefresh  pipleine or anything else.

In this repo there is a simple implementation of slack notify component that recieves arguments of 
 SLACK_CHANNEL - env variable full url for shared channel to send a messgage
 msg flag - that help to override message
 
 
 In ./demochart/template/test there is a definition of job , it uses already build verchol/slackjob public images built from 
 this repo (see below)
 
 to execute the demo 
 ```
 git clone https://github.com/verchol/slackjob.git && cd ./slackjob
 
 helm3 install  --debug d2 ./demochart/ --set Slack.SLACK_CHANNEL=$SLACK_CHANNEL --set Slack.SLACK_MSG="notificaiton message"
 
 ```
 
 
 ```
{{define "defaultmessage"}}
  "hello from" -  {{template "demochart.fullname" .}}
{{end}}
apiVersion: batch/v1
kind: Job
metadata:
  name: "{{ include "demochart.fullname" . }}-{{ uuidv4 }}-job"
  annotations:
    "helm.sh/hook": "post-install, post-upgrade"
    "helm.sh/hook-delete-policy": "hook-failed, hook-succeded" 
 
spec:
  backoffLimit: 6
  completions: 1
  parallelism: 1
  template:
    spec:
      containers:
      - image: verchol/slackjob
        imagePullPolicy: Always
        env:
        - name: SLACK_CHANNEL
          value: {{.Values.Slack.SLACK_CHANNEL}}
        name: j1
        args: ["--msg", {{ default (include "defaultmessage" .) .Values.Slack.SLACK_MSG | quote }}]
      restartPolicy: Never
      
  ```
  
  
 
