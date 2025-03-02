mock:
	mockgen -source=controller/category_controller.go -destination=controller/mocks/category_controller_mock.go -package=mocks
	mockgen -source=repository/category_repository.go -destination=repository/mocks/category_repository_mock.go -package=mocks
	mockgen -source=service/category_service.go -destination=service/mocks/category_service_mock.go -package=mocks
	mockgen -source=repository/product_repository.go -destination=repository/mocks/product_repository_mock.go -package=mocks
	mockgen -source=service/product_service.go -destination=service/mocks/product_service_mock.go -package=mocks
	mockgen -source=controller/product_controller.go -destination=controller/mocks/product_controller_mock.go -package=mocks
