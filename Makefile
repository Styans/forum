build :
	docker build -t go-app .

run :
	docker run -p 8000:8000 go-app

	
