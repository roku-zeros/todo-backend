# Архитектура
├── docker-compose.yaml           # Конфигурация Docker Compose  
├── lib                           # Библиотеки  
│   ├── config_utils              # Утилиты для работы с конфигурацией  
│   └── logger                    # Логирование  
└── services                      # Сервисы приложения  
    └── task                      # Сервис управления задачами  
        ├── Dockerfile            # Файл Docker для сервиса задач  
        ├── README.md             # Документация сервиса задач  
        ├── cmd                   # Команды для запуска приложения  
        │   └── main.go           # Главный файл приложения    
        ├── docs                  # Документация API  
        │   └── api               # API документация  
        │       └── v1            # Версия 1 API  
        │           └── swagger.yaml  # Swagger документация API  
        ├── internal              # Внутренние пакеты приложения  
        │   ├── app                # Логика приложения  
        │   ├── config             # Конфигурация приложения    
        │   ├── models             # Модели данных  
        │   ├── providers          #  Бизнес Слой
        │   ├── repository         # Слой с БД   
        │   └── server             # Серверный слой
        ├── migrations              # Миграции базы данных  
        │   └── postgres            # Миграции для PostgreSQL  
        └── test                   # Тесты приложения  
            └── test.sh            # Скрипт для запуска интеграционных тестов  
# Запуск
docker-compose up -d
# Запуск тестов
Понадобится jq 
### linux
sudo apt-get update
sudo apt-get install jq
### macos
brew install jq
### Windows
choco install jq


cd services/task/test  
./test.sh
# Swagger
[http://localhost:8081](http://localhost:8081) (после запуска приложения)
