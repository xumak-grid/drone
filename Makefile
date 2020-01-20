.PHONY: deploy-service-aws deploy-aws-pod new-tag
deploy-service:
	kubectl --namespace=ehernandez create -f k8s/service.yml

deploy-pod:
	kubectl --namespace=ehernandez create -f k8s/pod.yml

remove:
	kubectl --namespace=ehernandez delete -f k8s/service.yml || true
	kubectl --namespace=ehernandez delete -f k8s/pod.yml || true

new-tag:
	git tag -a v"$(VERSION)" -m "version v$(VERSION)"
	git push --tags
