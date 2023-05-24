create:
	docker build -t anchik . 
run:
	docker run -d -p 8080:8080 --rm --name anuarsabitovich anchik

start:
	docker start anuarsabitovich
	
stop: 
	docker stop anuarsabitovich

prune:
	docker container prune 
