default:
	@echo "To build all: make build\nTo build docker images: make docker-build\nTo run all: make run\nTo run docker images: make docker-run"
build-leads:
	cd backend/leadsvc && go build .
build-contacts:
	cd backend/contactsvc && go build .
build-company:
	cd backend/companysvc && go build .
build-opportunity:
	cd backend/opportunitysvc && go build .
build-query:
	cd backend/querysvc && dotnet build .
build-processor:
	cd backend/processorsvc && dotnet build .
build: build-leads build-contacts build-company build-opportunity build-query build-processor
	@echo "Building all services..."
docker-leads:
	cd backend/leadsvc && docker build -t am8850/contoso-crm-leads .
docker-build: docker-leads
	@echo "Building docker images..."
