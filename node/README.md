## Laravel
```
docker compose up 
```
Visit: http://localhost/api/message?text=HelloGDGCloud

On terminal
```
docker exec -it laravel-laravel.test-1 bash
php artisan queue:work --queue=test-sub
```


## Node.js
export GOOGLE_APPLICATION_CREDENTIALS="path/to/your/service-account-key.json"
node index.js publish "Hello, GDG Cloud Rajkot"
node index.js listen