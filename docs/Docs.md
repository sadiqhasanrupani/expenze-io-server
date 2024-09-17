## Expenze IO Server Installation Process

---

### Clone the Repository

---

```zsh
git clone https://github.com/boogySquad/expenzeIo-server.git
```

### Install dependencies

---

```zsh
go mod tidy
```


### Configure .env file
1. Create a `.env` file in your repository by going to the directory by `cd expenzeIo-server` and then write this command to create `.env` file
	```bash
	touch .env
	```
2. then paste the PostgreSQL Conn string
	```env
	PG_CONNSTR = "postgres://username:password@hostname:port/expenze-io?sslmode=disable"

	PG_CONNSTR = "postgres://postgres:newpassword@localhost:5432/expenze-io?sslmode=disable"

	COMPANY_NAME = "company_name"
	COMPANY_EMAIL = "company_email"
	
	APP_PASS = "app_password"
	SMTP_EMAIL = "smtp email service"	
	```


### To methods to run the project

1. Dev Mode
	1. navigate to the project
		```bash
		cd expenzeIo-server
		```
	1. Download air in your  device
		```bash
		go install github.com/air-verse/air@latest
		```
	3. It will create `.air.toml` file in your repo, and you need to change these line inside that file
		```toml
		bin = "./tmp/main"
		cmd = "go build -o ./tmp/main ./app/main.go"
		```
	4. Now that you are successfully install the air cli, then you just need to run this command
		```bash
			air
		```
	5. It will show this type of logs,
		```bash
		air

		  __    _   ___
		 / /\  | | | |_)
		/_/--\ |_| |_| \_ v1.52.3, built with Go go1.23.1
		
		watching .
		watching app
		watching db
		watching db/migrations
		watching db/seeders
		watching docs
		watching internal
		watching internal/config
		watching internal/controllers
		watching internal/handlers
		watching internal/middlewares
		watching internal/models
		watching internal/repositories
		watching internal/routes
		watching internal/services
		watching internal/tables
		watching pkg
		watching routes
		!exclude tmp
		building...
		running...
		2024/09/11 18:23:38 Users table created successfully
		2024/09/11 18:23:38 Otps table created successfully.
		Database is up and running
		[GIN-debug] [WARNING] Creating an Engine instance with the Logger and Recovery middleware already attached.
		
		[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
		 - using env:   export GIN_MODE=release
		 - using code:  gin.SetMode(gin.ReleaseMode)
		
		[GIN-debug] GET    /                         --> expenze-io.com/routes.RegisterRoutes.func1 (3 handlers)
		[GIN-debug] POST   /api/v1/auth/register     --> expenze-io.com/internal/controllers.RegisterHandler (3 handlers)
		[GIN-debug] POST   /api/v1/auth/login        --> expenze-io.com/internal/controllers.LoginHandler (3 handlers)
		[GIN-debug] [WARNING] You trusted all proxies, this is NOT safe. We recommend you to set a value.
		Please check https://pkg.go.dev/github.com/gin-gonic/gin#readme-don-t-trust-all-proxies for details.
		[GIN-debug] Listening and serving HTTP on :8080

		```

2. Second way is run `go run ./app/main.go` and that's it