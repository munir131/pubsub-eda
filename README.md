# Pub/Sub Event-Driven Architecture Demo (Laravel + Node.js)

This repository demonstrates a **minimal, production-oriented Event-Driven Architecture (EDA)** using **Google Cloud Pub/Sub** with:

* **Laravel** as the event producer & consumer
* **munir131/laravel-pubsub-queue** as the Pub/Sub queue driver
* **Node.js** as an independent subscriber example

The goal of this repository is to provide a **clean reference implementation** for building fully decoupled, event-driven microservices using Pub/Sub.

---

## 🚀 Features

* ✔️ Laravel 10 app wired with Google Pub/Sub
* ✔️ Queue driver provided by **munir131/laravel-pubsub-queue**
* ✔️ Example endpoint that dispatches messages (`/send`)
* ✔️ Dedicated Pub/Sub queue connection (`pubsub`) 
* ✔️ Example job: `ProcessMessage` logs the received message 
* ✔️ Node.js minimal subscriber module
* ✔️ Docker-ready through Laravel Sail

---

## 📁 Repository Structure

```
munir131-pubsub-eda/
├── laravel/    # Laravel PubSub producer + consumer
│   ├── app/Jobs/ProcessMessage.php
│   ├── app/Http/Controllers/MessageController.php
│   ├── config/queue.php
│   └── ...
└── node/       # Node.js subscriber example
    ├── index.js
    └── package.json
```

Key Laravel files referenced:

* **queue.php** defines Pub/Sub connection configuration:

  ```php
  'pubsub' => [
      'driver' => 'pubsub',
      'queue' => 'test',
      'project_id' => env('PUBSUB_PROJECT_ID', 'pubsub-eda'),
      'keyFilePath' => env('PUBSUB_KEY_FILE', storage_path('pubsub-eda.json')),
      'subscribers' => [
          'test-sub' => 'test'
      ]
  ]
  ```



---

# 🧰 Requirements

* PHP 8.1+
* Composer
* Node.js 18+
* Docker (optional but recommended)
* Google Cloud Pub/Sub project + service account key

---

# ⚙️ Setup Instructions

## 1️⃣ Clone the repository

```bash
git clone https://github.com/munir131/pubsub-eda.git
cd pubsub-eda
```

---

# 📌 Laravel Setup

## 2️⃣ Install dependencies

```bash
cd laravel
composer install
```

## 3️⃣ Configure Pub/Sub

Copy environment file:

```bash
cp .env.example .env
```

Add:

```
PUBSUB_PROJECT_ID=pubsub-eda
PUBSUB_KEY_FILE=storage/pubsub-eda.json
QUEUE_CONNECTION=pubsub
```

Place your **Google service account JSON** at:

```
laravel/storage/pubsub-eda.json
```

## 4️⃣ Run Laravel

Using Sail (recommended):

```bash
./vendor/bin/sail up
```

Or locally:

```bash
php artisan serve
```

---

# 🚦 How to Test Pub/Sub Flow

## ▶️ 1. Trigger a Pub/Sub message

```
GET http://localhost/send?text=Hello-PubSub
```

This executes `send()` in `MessageController` which dispatches the job to Pub/Sub.


---

## ▶️ 2. Laravel worker listens and processes messages

```bash
php artisan queue:work pubsub
```

You should see:

```
Processing: App\Jobs\ProcessMessage
```

The job logs the message:

```php
Log::info($this->message);
```

Laravel logs contain the message payload.

---

# 🟦 Node.js Subscriber Example

A simple subscriber is located under `node/`.

## Install node dependencies

```bash
cd node
npm install
```

## Run subscriber

```bash
node index.js
```

This script:

✔ Connects to Google Pub/Sub
✔ Subscribes to a topic
✔ Prints incoming messages

Use this to simulate non-Laravel microservices in your event-driven environment.

---

# 🧩 Understanding the Flow

```
Client → Laravel Controller → Pub/Sub Topic → Queue Worker → ProcessMessage Job
                                            ↘
                                             ↘ Node Subscriber
```

Laravel produces events →
Google Pub/Sub delivers messages →
Laravel workers & Node.js services consume them independently.

---

# 🧪 Testing Scenarios

| Scenario                        | What Happens                                       |
| ------------------------------- | -------------------------------------------------- |
| Hit `/send` without query param | Sends default message `"This is default message"`  |
| Node subscriber running         | Both Laravel worker & Node receive messages        |
| Worker stopped                  | Messages stay in Pub/Sub until pulled              |

---

# 📦 Packages Used

* **munir131/laravel-pubsub-queue** — custom Pub/Sub driver
* **Laravel 10** — application framework
* **Google Cloud Pub/Sub SDK** via queue driver
* **Node.js PubSub client** for subscriber example

---

# 🤝 Contributing

PRs are welcome!

If you'd like to improve the architecture examples (fan-out patterns, dead-letter queues, saga orchestration), feel free to open an issue.

---

# 📝 License

This project is open-sourced under the MIT license.
