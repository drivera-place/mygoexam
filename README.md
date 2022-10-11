# Lector OCR

This program receives a file with digits inputs and returns parsed and validate the input as a string. 

e.g.

Input:
```
                           
|_||_||_||_||_||_||_||_||_|
  |  |  |  |  |  |  |  |  |
```
Output:
```
444444444 OK
```

### Requirements:
- Go 1.17+

### Execution:
``` go run main.go ```

### Test:

Server listens at `http://localhost:5050/upload`

Send a POST request with the input file, you can test with curl:
```
curl -X POST http://localhost:5050/upload \
-F "file=@./entradas.txt" \
-H "Content-Type: multipart/form-data"
```

### Input file:
[entradas.txt](./entradas.txt)

