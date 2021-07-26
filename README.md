# owc-insight
openwechat insight 

## Features

- Using the industry best practices, e.g., machine-readable&analyzable logging.
- Trying to be as considerate/user-friendly as possible. E.g., display QR Code right in terminal, and even in small 80x25 size windows:  
  ![image](https://user-images.githubusercontent.com/422244/126873409-6cd02e31-cf96-4a60-b60d-cbe90eb31725.png)
- Make every possible effort to be robust and fault-tolerance. E.g., trying to recover from any _temporary_ failures:  
  ![image](https://user-images.githubusercontent.com/422244/126901127-c1669681-269e-4142-b60f-0f895ba63a0c.png)
- Using every possible attempt to stay alive & connected, while making the approach discreet (undetectable from the server end). E.g., stay-alive in debug mode:  
  ![image](https://user-images.githubusercontent.com/422244/126901864-b25ab53d-b9dd-4241-87c6-37e08c6efdb1.png)

## Execution

Fault-tolerance execution loop, until no longer able to do hot relogin:

```sh
while :; do OWCI_LOG=2 owc-insight ; [ $? -eq 9 ] && break; sleep 5m; done; rm storage.json; echo "Down at `date`", send Down Alert
```
![image](https://user-images.githubusercontent.com/422244/126921429-d57c7853-bddd-46ca-8803-5dc240271467.png)

This has been running for the whole day and is still going strong. 
