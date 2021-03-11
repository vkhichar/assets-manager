# Assets manager
Asset management system for an organisation/business

# How to build, run
```
1. Install vendor dependencies
$ go mod vendor

2. Build the binary
$ go build

3. Export environment variables
Check the 'application.yml.sample' file and export the value of envs
for example:
$ export APP_PORT=9000
$ export DB_HOST=localhost

4. Create DB migrations
Create a database with the following name: 'assets_manager'
Check migrations folder for creating db tables/indexes

5. Run the server
$ ./assets-manager
```