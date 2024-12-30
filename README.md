Here’s a sample `README.md` file for your project:

```markdown
# sv-cli

`sv-cli` is a cross-platform command-line tool that provides system information and checks the status of services, databases, and Docker containers. It works on both Linux and Windows.

## Features

- Retrieve RAM and disk usage statistics.
- Check if a service is running (supports Linux and Windows).
- Check the status of Docker containers.
- Test database connectivity for MySQL, PostgreSQL, and MSSQL.

## Prerequisites

- **Go**: Ensure you have Go (1.20 or later) installed on your system. You can download it from [Go Downloads](https://go.dev/dl/).

## Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/yourusername/sv-cli.git
   cd sv-cli
   ```

2. Install dependencies:
   ```bash
   go mod tidy
   ```

## Building the Application

### For Linux

Run the following command:
```bash
go build -o sv
```

This will create an executable named `sv` in the current directory.

### For Windows

Run the following command:
```bash
go build -o sv.exe
```

This will create an executable named `sv.exe` in the current directory.

## Usage

`sv-cli` uses subcommands to perform various operations. Each subcommand has specific options and flags.

### Common Flag

- `--format`: Specifies the output format (`table` or `json`). Default is `table`.

### Commands

#### RAM Usage
Retrieve RAM usage:
```bash
./sv ram --format=json
```

#### Disk Usage
Retrieve disk usage:
```bash
./sv disk --format=table
```

#### Service Status
Check if a service is running:
```bash
./sv service [service_name] --format=json
```
Example:
```bash
./sv service nginx --format=table
```

#### Docker Container Status
Check if a Docker container is running:
```bash
./sv docker [container_name] --format=table
```
Example:
```bash
./sv docker my-container --format=json
```

#### Database Connectivity
Check database connectivity using a DSN (Data Source Name).

##### MySQL
```bash
./sv db mysql --dsn="user:password@tcp(localhost:3306)/dbname" --format=json
```

##### PostgreSQL
```bash
./sv db postgres --dsn="postgres://user:password@localhost:5432/dbname?sslmode=disable" --format=table
```

##### MSSQL
```bash
./sv db mssql --dsn="sqlserver://user:password@localhost:1433?database=dbname" --format=json
```

## Example Outputs

### JSON Output
```json
{
  "database": "MySQL",
  "status": "success",
  "message": "Connection successful"
}
```

### Table Output
```
+----------+---------+---------------------+
| DATABASE | STATUS  |       MESSAGE       |
+----------+---------+---------------------+
| MySQL    | success | Connection successful |
+----------+---------+---------------------+
```

## Development

To run the application without building:
```bash
go run main.go [command]
```

Example:
```bash
go run main.go ram --format=json
```

## Contributing

Feel free to fork this repository and submit pull requests. Issues and feature requests are welcome!

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
```

---

### How to Use
- Replace `yourusername` in the `git clone` command with your GitHub username if you host the repository on GitHub.
- Update the `LICENSE` section if you use a different license.

Let me know if you’d like to customize this further!