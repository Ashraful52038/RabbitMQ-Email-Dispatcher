## 📧 Multi-App Email Dispatch System with RabbitMQ & Mailpit

![Version](https://img.shields.io/badge/version-1.0.0-blue)
![Go](https://img.shields.io/badge/Go-1.21+-00ADD8)
![RabbitMQ](https://img.shields.io/badge/RabbitMQ-3.13-FF6600)
![Docker Compose](https://img.shields.io/badge/Docker-Compose-2496ED)

**📋 Project Overview**  
This project is an email dispatcher system that collects emails from 3 different applications, queues them in RabbitMQ, and dispatches them to Mailpit (test mail server) using 2 concurrent workers.  

➡️ In short: **A RabbitMQ‑powered pipeline that reliably delivers emails to Mailpit with concurrent worker processing.**


**🏗️ Architecture Diagram**
graph TD
    subgraph "Publisher Apps"
        A1[App 1<br/>5 emails]
        A2[App 2<br/>5 emails]
        A3[App 3<br/>5 emails]
    end

    subgraph "Message Broker"
        RMQ[RabbitMQ<br/>email_queue]
    end

    subgraph "Consumer Workers"
        W1[Worker 1]
        W2[Worker 2]
    end

    subgraph "Mail Server"
        MP[Mailpit<br/>Test Mail Server]
    end

    A1 -->|1.5s interval| RMQ
    A2 -->|1.5s interval| RMQ
    A3 -->|1.5s interval| RMQ
    
    RMQ -->|2s delay| W1
    RMQ -->|2s delay| W2
    
    W1 --> MP
    W2 --> MP
    
    MP -->|UI| B1[Browser<br/>localhost:8025]
    RMQ -->|Management UI| B2[Browser<br/>localhost:15673]


   
    
**🔄 Workflow Diagram**

sequenceDiagram
    participant App1 as App 1
    participant App2 as App 2
    participant App3 as App 3
    participant RMQ as RabbitMQ
    participant W1 as Worker 1
    participant W2 as Worker 2
    participant MP as Mailpit

    loop Each App (5 iterations)
        App1->>RMQ: Send Email (1.5s interval)
        App2->>RMQ: Send Email (1.5s interval)
        App3->>RMQ: Send Email (1.5s interval)
    end

    RMQ-->>W1: Deliver Message
    RMQ-->>W2: Deliver Message
    
    Note over W1,W2: 2 seconds processing delay
    
    W1->>MP: Dispatch Email
    W2->>MP: Dispatch Email
    
    MP-->>Browser: Show in UI (port 8025)
    RMQ-->>Browser: Management UI (port 15673)
    
**✨ Features**

    ✅ 3 Independent Publisher Apps - Each sends 5 emails

    ✅ RabbitMQ Message Queue - Reliable message delivery

    ✅ 2 Concurrent Workers - Parallel processing as per Phase 2

    ✅ Mailpit Integration - Test mail server with Web UI

    ✅ Delay Mechanisms - 1.5s publish, 2s consume delay

    ✅ Docker Compose - Easy service management

    ✅ Makefile - Command shortcuts

**🛠️ Technology Stack**

    Backend Language: Go 1.21+

    Message Broker: RabbitMQ 3.13

    Test Mail Server: Mailpit

    Container: Docker & Docker Compose

    Protocol: AMQP, HTTP

📁 Project Structure
graph TD
    A[📂 email-dispatcher]

    A --> B[docker-compose.yml<br/>RabbitMQ & Mailpit service config]
    A --> C[Makefile<br/>Command shortcuts]
    A --> D[go.mod<br/>Go module dependencies]
    A --> E[go.sum<br/>Go checksums]

    A --> F[📂 publisher]
    F --> F1[main.go<br/>Simulates 3 apps]

    A --> G[📂 consumer]
    G --> G1[main.go<br/>Consumer with 2 workers]
    G --> G2[📂 mailpit]
    G2 --> G3[client.go<br/>Mailpit API client]

    A --> H[📂 shared]
    H --> H1[📂 models]
    H1 --> H2[email.go<br/>Common email model]

**🚀 Installation Guide**

Prerequisites

    Docker & Docker Compose

    Go 1.21 or higher

    Make (optional)

Step by Step Installation
1. Clone the Project
        
        git clone https://github.com/yourusername/email-dispatcher.git
        cd email-dispatcher

2. Install Dependencies
        
        go mod download

3. Start Docker Services
        
        make up
        # OR
        docker compose up -d

4. Check Services
        
        docker compose ps

🎮 How to Use
1. Start Consumer (Terminal 1)
        
        make consumer
        # OR
        go run consumer/main.go

2. Start Publisher (Terminal 2)
        bash
        
        make publisher
        # OR
        go run publisher/main.go

3. Access UIs
     | Service               | URL                       | Credentials   |
|-----------------------|---------------------------|---------------|
| RabbitMQ Management   | http://localhost:15673    | guest/guest   |
| Mailpit UI            | http://localhost:8025     | None          |

**📊 Ports & Services**
| Service              | Port  | Usage                  |
|----------------------|-------|------------------------|
| RabbitMQ AMQP        | 5673  | Go app connection      |
| RabbitMQ Management  | 15673 | Web UI                 |
| Mailpit SMTP         | 1025  | Email sending          |
| Mailpit HTTP         | 8025  | Web UI & API           |


**💡 Mechanism Explanation**
Publisher Flow
flowchart LR
    Start([Start]) --> Loop{Loop 5 times}
    Loop -->|Yes| Create[Create Email]
    Create --> Send[Send to RabbitMQ]
    Send --> Wait[Wait 1.5s]
    Wait --> Loop
    Loop -->|No| Stop([Stop])
    
Consumer Flow
flowchart LR
    Start([Start]) --> Workers{2 Workers}
    Workers --> W1[Worker 1]
    Workers --> W2[Worker 2]
    
    W1 --> Get[Get from Queue]
    W2 --> Get
    
    Get --> Delay[Wait 2s]
    Delay --> Process[Process Email]
    Process --> Mailpit[Send to Mailpit]
    Mailpit --> Ack[Acknowledge]
    Ack --> Get
    
**🧪 Testing**

Mailpit API Test
        
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
        
        # View queues
        curl -s -u guest:guest http://localhost:15673/api/queues | python3 -m json.tool
        
        # View consumers
        curl -s -u guest:guest http://localhost:15673/api/consumers | python3 -m json.tool

**🔧 Troubleshooting**
Common Issues & Solutions
| Issue                | Cause                          | Solution                                      |
|-----------------------|--------------------------------|-----------------------------------------------|
| Port already in use   | Another program using port     | Kill process with `sudo lsof -i :5673`        |
| Container not starting| Docker image issue             | `docker compose down -v && docker compose up -d` |
| Connection refused    | Service not ready              | Wait with `sleep 10`                          |
| 401 Unauthorized      | Wrong credentials              | Use `guest/guest`                             |

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

**📈 Performance**

    Total Emails: 15 (3 apps × 5 emails)

    Publish Rate: 3 emails every 1.5 seconds

    Consume Rate: 2 parallel workers

    Processing Time: 2s delay + 0.5s processing

**👥 Contributing**

    Fork the project

    Create feature branch (git checkout -b feature/AmazingFeature)

    Commit changes (git commit -m 'Add some AmazingFeature')

    Push to branch (git push origin feature/AmazingFeature)

    Open a Pull Request

**📞 Contact**

**🔗 Useful Links**

- **Project Link:** [RabbitMQ Email Dispatcher](https://github.com/Ashraful52038/RabbitMQ-Email-Dispatcher)  
- **RabbitMQ Docs:** [https://www.rabbitmq.com/documentation.html](https://www.rabbitmq.com/documentation.html)  
- **Mailpit Docs:** [https://mailpit.axllent.org/docs/](https://mailpit.axllent.org/docs/)

**🎯 Application Flow (Code Level)**
classDiagram
    class EmailEvent {
        +string AppID
        +string EmailID
        +string Recipient
        +string Subject
        +string Body
        +int64 Timestamp
    }
    
    class Publisher {
        +publishEmails()
    }
    
    class Consumer {
        +worker()
    }
    
    class MailpitClient {
        +SendEmail()
        +SendEmailViaAPI()
    }
    
    Publisher --> EmailEvent
    Consumer --> EmailEvent
    Consumer --> MailpitClient
    MailpitClient --> EmailEvent

**🏁 Quick Start Commands**
        
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
