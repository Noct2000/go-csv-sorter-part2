## task 2
### for Development of microservice systems in the Golang language
“CSV Sorter” is a CLI application that allows sorting of its input presented as CSV-text.
Technical details
### Required features:
1. The application runs as a CLI application.
2. It reads STDIN line by line. The end of the input is an empty line.
3. Each line is a list of comma-separated values (CSV). Each value is considered as a piece of text. The
number of values is the same in each line.
4. The application sorts all lines alphabetically by the first value in each line.
5. The application prints the result immediately, when the user ends to enter input text (presses <Enter> at a new line).
6. The application supports options:
Option usage
         -i file-name | Use a file with the name file-name as an input.
         -o file-name | Use a file with the name file-name as an output.
         -h           | The first line is a header that must be ignored during sorting but included in the output.
         -f N         | Sort input lines by value number N.
         -r           | Sort input lines in reverse order.
         -d dir-name  | That specifies a directory where it must read input files from
 ### How to use
 1. clone repository
 2. run 
 ```shell
 cd go-csv-sorter-part2
 go run main.go
 ```
