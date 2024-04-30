# genius-ecommerce

A simple e-commerce project for practicing k8s!

## Technical requirements

### Front-end

1. React/Next.js
2. Responsive design
3. Mobx

### Back-end

1. Go
2. Gin
3. Golang validator V10 and default Go Error handling

### Database

1. PostgreSQL
2. Separated databases for each of microsservice
3. Minimal constraints for each database

### Kubernetes

1. K8s manifests for each microsservice
2. Use of Helm charts to manage deployments
3. Sensitive data requires using secrets

### Observability

1. Uses Prometheus and Grafana
2. Alert configuration for critical metrics
3. Centralizing logs with Fluentd

### Security

1. RBAC for k8s resources
2. Network policies for network communication
3. Use HTTPS

### Pipeline

1. Use github actions, online or offline
2. Test and build automation
