# ssh-load-tester

[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

CLI tool for load testing ssh logins.

This should only be used in a controlled, ethical environment (like your own infrastructure) and with express permissions to avoid any potential damage or violations of terms of service.

## Install

```bash
go install github.com/adegoodyer/ssh-load-tester@latest
```

## Usage
```bash
ssh-load-tester
Error: -h flag (host) is required.
Usage of ssh-load-tester:
  -P string
        password for SSH login (default "1nV@l!dP@ss")
  -U string
        username for SSH login (default "ssh-load-tester")
  -d duration
        test duration (default 10s)
  -h string
        host IP or hostname (required)
  -p string
        SSH port (default "22")
  -r int
        rate of attempts per second (default 5)
```

## Samples
```bash
# specify host
ssh-load-tester -h localhost

# specify host and port
ssh-load-tester -h localhost -p 2022

# specify duration and logins-per-second
ssh-load-tester -h localhost -d 5s -r 10

# specify username and password
ssh-load-tester -h localhost -U bob -P p@ssw0rd!
```

## Sample Output
```bash
ssh-load-tester -h localhost -d 2s
----- Test Summary -----
Host: localhost
Port: 22
Username: ssh-load-tester
Test Duration: 2s
Rate of Attempts: 5 per second
Interval between Attempts: 200ms
------------------------
12:09:07.963 Starting attempt 1 with interval 200ms...
12:09:07.966 Attempt 1 failed: dial tcp [::1]:22: connect: connection refused
12:09:08.164 Starting attempt 2 with interval 200ms...
12:09:08.166 Attempt 2 failed: dial tcp [::1]:22: connect: connection refused
12:09:08.366 Starting attempt 3 with interval 200ms...
12:09:08.367 Attempt 3 failed: dial tcp [::1]:22: connect: connection refused
12:09:08.567 Starting attempt 4 with interval 200ms...
12:09:08.568 Attempt 4 failed: dial tcp [::1]:22: connect: connection refused
12:09:08.767 Starting attempt 5 with interval 200ms...
12:09:08.768 Attempt 5 failed: dial tcp [::1]:22: connect: connection refused
12:09:08.967 Starting attempt 6 with interval 200ms...
12:09:08.969 Attempt 6 failed: dial tcp [::1]:22: connect: connection refused
12:09:09.168 Starting attempt 7 with interval 200ms...
12:09:09.169 Attempt 7 failed: dial tcp [::1]:22: connect: connection refused
12:09:09.370 Starting attempt 8 with interval 200ms...
12:09:09.371 Attempt 8 failed: dial tcp [::1]:22: connect: connection refused
12:09:09.571 Starting attempt 9 with interval 200ms...
12:09:09.572 Attempt 9 failed: dial tcp [::1]:22: connect: connection refused
12:09:09.772 Starting attempt 10 with interval 200ms...
12:09:09.773 Attempt 10 failed: dial tcp [::1]:22: connect: connection refused
12:09:09.973 ----- Test Completed -----
Total Attempts Made: 10
All login attempts finished.
```

## Tags

- `latest`: Most recent stable build
- `x.y.z`: Specific version builds (e.g., `2.7.5`)
- `x.y`: Minor version builds (e.g., `2.7`)

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
