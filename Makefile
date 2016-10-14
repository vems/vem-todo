.PHONY: deploy destroy

deploy:
	kubectl create -f kubernetes/todo/todo-deployment.yaml --validate=false
	kubectl create -f kubernetes/todo/todo-service.yaml --validate=false
	kubectl create -f kubernetes/web/web-deployment.yaml --validate=false
	kubectl create -f kubernetes/web/web-service.yaml --validate=false
	kubectl create -f kubernetes/api/api-deployment.yaml --validate=false
	kubectl create -f kubernetes/api/api-service.yaml --validate=false
	kubectl create -f kubernetes/proxy/proxy-deployment.yaml --validate=false
	kubectl create -f kubernetes/proxy/proxy-service.yaml --validate=false

destroy:
	kubectl delete services todo
	kubectl delete services web
	kubectl delete services api
	kubectl delete services proxy
	kubectl delete deployments todo
	kubectl delete deployments web
	kubectl delete deployments api
	kubectl delete deployments proxy
