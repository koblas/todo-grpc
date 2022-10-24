
all:
	kubectl apply -f dev-variables-env-configmap.yaml
	kubectl apply -f backend-networkpolicy.yaml
	kubectl apply -f dmz-networkpolicy.yaml
	kubectl apply -f frontend-networkpolicy.yaml
	#
	kubectl apply -f redis-deployment.yaml
	kubectl apply -f redis-service.yaml
	#
	kubectl apply -f publicapi-auth-deployment.yaml
	kubectl apply -f publicapi-auth-service.yaml
	kubectl apply -f publicapi-todo-deployment.yaml
	kubectl apply -f publicapi-todo-service.yaml
	kubectl apply -f publicapi-websocket-deployment.yaml
	kubectl apply -f publicapi-websocket-service.yaml
	#
	kubectl apply -f core-oauth-user-deployment.yaml
	kubectl apply -f core-oauth-user-service.yaml
	kubectl apply -f core-user-deployment.yaml
	kubectl apply -f core-user-service.yaml
	kubectl apply -f core-todo-deployment.yaml
	kubectl apply -f core-todo-service.yaml
	#
	kubectl apply -f core-send-email-deployment.yaml
	kubectl apply -f core-workers-deployment.yaml
	kubectl apply -f core-websocket-todo-deployment.yaml
	#
	kubectl apply -f client-deployment.yaml
	kubectl apply -f client-service.yaml
	#
	kubectl apply -f ingress.yaml
