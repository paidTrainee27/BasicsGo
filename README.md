# Basic GO codes.

# Basic project structure is:

├── cmd/
│   └── main.go //启动函数
├── etc
│   └── dev_conf.yaml              // The configuration file
├── global
│   └── global.go //Global variable references, such as databases, kafka, etc
├── internal/
│       └── service/
│           └── xxx_service.go //Business logic processing classes
│           └── xxx_service_test.go 
│       └── model/
│           └── xxx_info.go//The structure 
│       └── api/
│           └── xxx_api.go//The interface to the route is implemented
│       └── router/
│           └── router.go
│       └── pkg/
│           └── datetool//Time tool class
│           └── jsontool//Json utility class