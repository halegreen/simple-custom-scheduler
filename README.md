# simple-custom-scheduler
基于 k8s Scheduler Framework 开发的自定义调度器

部署：
```
$ kubectl apply -f deploy/sample-scheduler.yml
```

验证：
```
$ kubectl apply -f deploy/application.yml
```