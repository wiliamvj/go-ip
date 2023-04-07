## Go IP Search
This is a command-line application written in Go that allows you to search for IP addresses and servers on the web by entering a hostname.

Usage
To use the application, you must have Go installed on your system. You can run the following command to download and install the application:

go
Copy code
go get github.com/wiliamvj/go-ip-search
Once installed, you can run the application using the following command:

css
Copy code
go-ip-search [COMMAND] [FLAGS]
Where [COMMAND] is one of the following:

ip: Search for IP addresses on the web
server: Search for servers on the web
And [FLAGS] is an optional list of flags that can be passed to the command. The available flags are:

--host: The hostname to search for. Defaults to github.com.br
Examples
To search for IP addresses using the default hostname, run the following command:

go
Copy code
go-ip-search ip
To search for servers using a custom hostname, run the following command:

sql
Copy code
go-ip-search server --host google.com
Contributing
If you would like to contribute to this project, please feel free to open an issue or submit a pull request. Your feedback and contributions are greatly appreciated!
