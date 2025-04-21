build_all:
	docker build -t stress_tester . && 	cd server && docker build -t server .

run_server:
	docker run -p 8080:8080 server

run_stress_tester:
	docker run --name stress_tester stress_tester --url=http://172.17.0.2:8080 --requests=105 --concurrency=10