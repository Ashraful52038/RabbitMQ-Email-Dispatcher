📧 Email Dispatcher System
Multi-App Email Dispatch System with RabbitMQ & Mailpit

https://img.shields.io/badge/version-1.0.0-blue
https://img.shields.io/badge/Go-1.21+-00ADD8
https://img.shields.io/badge/RabbitMQ-3.13-FF6600
https://img.shields.io/badge/Docker-Compose-2496ED
📋 Project Overview

This project is an email dispatcher system that collects emails from 3 different applications, queues them in RabbitMQ, and dispatches them to Mailpit (test mail server) using 2 concurrent workers.
🏗️ Architecture Diagram
🔄 Workflow Diagram
✨ Features

    ✅ 3 Independent Publisher Apps - Each sends 5 emails

    ✅ RabbitMQ Message Queue - Reliable message delivery

    ✅ 2 Concurrent Workers - Parallel processing as per Phase 2

    ✅ Mailpit Integration - Test mail server with Web UI

    ✅ Delay Mechanisms - 1.5s publish, 2s consume delay

    ✅ Docker Compose - Easy service management

    ✅ Makefile - Command shortcuts

🛠️ Technology Stack

    Backend Language: Go 1.21+

    Message Broker: RabbitMQ 3.13

    Test Mail Server: Mailpit

    Container: Docker & Docker Compose

    Protocol: AMQP, HTTP

📁 Project Structure
text

email-dispatcher/
│
├── 📄 docker-compose.yml      # RabbitMQ & Mailpit service config
├── 📄 Makefile                 # Command shortcuts
├── 📄 go.mod                   # Go module dependencies
├── 📄 go.sum                   # Go checksums
│
├── 📂 publisher/               # Publisher application
│   └── 📄 main.go              # Simulates 3 apps
│
├── 📂 consumer/                # Consumer application
│   ├── 📄 main.go              # Consumer with 2 workers
│   └── 📂 mailpit/
│       └── 📄 client.go        # Mailpit API client
│
└── 📂 shared/                   # Shared code
    └── 📂 models/
        └── 📄 email.go          # Common email model

🚀 Installation Guide
Prerequisites

    Docker & Docker Compose

    Go 1.21 or higher

    Make (optional)

Step by Step Installation
1. Clone the Project
bash

git clone https://github.com/yourusername/email-dispatcher.git
cd email-dispatcher

2. Install Dependencies
bash

go mod download

3. Start Docker Services
bash

make up
# OR
docker compose up -d

4. Check Services
bash

docker compose ps

🎮 How to Use
1. Start Consumer (Terminal 1)
bash

make consumer
# OR
go run consumer/main.go

2. Start Publisher (Terminal 2)
bash

make publisher
# OR
go run publisher/main.go

3. Access UIs
Service	URL	Credentials
RabbitMQ Management	http://localhost:15673	guest/guest
Mailpit UI	http://localhost:8025	None
📊 Ports & Services
Service	Port	Usage
RabbitMQ AMQP	5673	Go app connection
RabbitMQ Management	15673	Web UI
Mailpit SMTP	1025	Email sending
Mailpit HTTP	8025	Web UI & API
💡 Mechanism Explanation
Publisher Flow
Consumer Flow
🧪 Testing
Mailpit API Test
bash

# Send test email
curl -X POST http://localhost:8025/api/v1/send \
  -H "Content-Type: application/json" \
  -d '{
    "To": [{"Email": "test@example.com"}],
    "From": {"Email": "sender@example.com"},
    "Subject": "Test",
    "Text": "Hello"
  }'

# View messages
curl -s http://localhost:8025/api/v1/messages | python3 -m json.tool

RabbitMQ API Test
bash

# View queues
curl -s -u guest:guest http://localhost:15673/api/queues | python3 -m json.tool

# View consumers
curl -s -u guest:guest http://localhost:15673/api/consumers | python3 -m json.tool

🔧 Troubleshooting
Common Issues & Solutions
Issue	Cause	Solution
Port already in use	Another program using port	Kill process with sudo lsof -i :5673
Container not starting	Docker image issue	docker compose down -v && docker compose up -d
Connection refused	Service not ready	Wait with sleep 10
401 Unauthorized	Wrong credentials	Use guest/guest
Diagnostic Commands
bash

# Check all containers
docker compose ps

# View RabbitMQ logs
docker compose logs rabbitmq

# View Mailpit logs
docker compose logs mailpit

# Check queue
docker compose exec rabbitmq rabbitmqctl list_queues

📈 Performance

    Total Emails: 15 (3 apps × 5 emails)

    Publish Rate: 3 emails every 1.5 seconds

    Consume Rate: 2 parallel workers

    Processing Time: 2s delay + 0.5s processing

🛡️ License

This project is licensed under the MIT License.
👥 Contributing

    Fork the project

    Create feature branch (git checkout -b feature/AmazingFeature)

    Commit changes (git commit -m 'Add some AmazingFeature')

    Push to branch (git push origin feature/AmazingFeature)

    Open a Pull Request

📞 Contact

    Project Link: https://github.com/yourusername/email-dispatcher

    RabbitMQ Docs: https://www.rabbitmq.com/documentation.html

    Mailpit Docs: https://mailpit.axllent.org/docs/

🎯 Application Flow (Code Level)
🏁 Quick Start Commands
bash

# 1. Clone project
git clone <repo-url>
cd email-dispatcher

# 2. Start services
make up

# 3. Start consumer (Terminal 1)
make consumer

# 4. Start publisher (Terminal 2)
make publisher

# 5. Open UIs
open http://localhost:15673
open http://localhost:8025

# 6. Shutdown
make down

Project Complete 🎉. Feel free to open issues if you have any questions.
