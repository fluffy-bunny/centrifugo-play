# CLI

Centrifugo pub sub testing.

## Usage

Run docker with centrifugo pro and a mock oauth2 server.

```bash
docker-compose -f .\docker-compose-pro.yml pull
docker-compose -f .\docker-compose-pro.yml up
```

## OAuth2 Server

you can create your own clients with claims by modifying the [config file](../../configs/mockoauth2/clients.json).  

The client:```centrifugo-connector-sa``` is used to excercise the ```caps``` claim functionality.  

## Build the CLI (windows)
  
```powershell
go build -o centrifugo-cli.exe cmd/cli/main.go
```

## Using CAPS tokens

Example [jwt](https://jwt.io/#debugger-io?token=eyJhbGciOiJFUzI1NiIsImtpZCI6IjBiMmNkMmU1NGM5MjRjZTg5ZjAxMGYyNDI4NjIzNjdkIiwidHlwIjoiSldUIn0.eyJjYXBzIjpbeyJhbGxvdyI6WyJwdWIiLCJzdWIiLCJwcnMiLCJoc3QiXSwiY2hhbm5lbHMiOlsiY29ubmVjdG9yOioiXSwibWF0Y2giOiJ3aWxkY2FyZCJ9XSwiY2xpZW50X2lkIjoiY2VudHJpZnVnby1jb25uZWN0b3Itc2EiLCJleHAiOjE3MTQ2NjIxNDgsImlhdCI6MTcxNDY1ODU0OCwiaXNzIjoiaHR0cDovL2xvY2FsaG9zdDo5ODAyIiwic3ViIjoiY2VudHJpZnVnby1jb25uZWN0b3Itc2EifQ.A4iM7qhLkTvVIoTeKz0-NZBuPjqV7lEbLkEw902TUzgTCnspygUP7l5stQpEY_0Ma6qM4G8CizaiAdBcs2-HBQ) is used for both publishing and subscribing to the channel ```connector:foobar```.  

The following will publish and subscribe to the channel ```connector:foobar``` using the default client:```centrifugo-connector-sa```  

### Publishing a message

```powershell
.\centrifugo-cli.exe publish --message='{"a":"b"}'
```

```console
7:23AM INF cmd\cli\root\publish\publish.go:37 > got token token={"access_token":"eyJhbGciOiJFUzI1NiIsImtpZCI6IjBiMmNkMmU1NGM5MjRjZTg5ZjAxMGYyNDI4NjIzNjdkIiwidHlwIjoiSldUIn0.eyJjYXBzIjpbeyJhbGxvdyI6WyJwdWIiLCJzdWIiLCJwcnMiLCJoc3QiXSwiY2hhbm5lbHMiOlsiY29ubmVjdG9yOioiXSwibWF0Y2giOiJ3aWxkY2FyZCJ9XSwiY2xpZW50X2lkIjoiY2VudHJpZnVnby1jb25uZWN0b3Itc2EiLCJleHAiOjE3MTQ2NjM0MjcsImlhdCI6MTcxNDY1OTgyNywiaXNzIjoiaHR0cDovL2xvY2FsaG9zdDo5ODAyIiwic3ViIjoiY2VudHJpZnVnby1jb25uZWN0b3Itc2EifQ._Uy2HqQfMLg2EWwFo2THBIUgaF__-7a7jldjqUndnznusE5Kg4Hjl5_SUx6GNHq0gTJ_zNzYJheKbl48LHIbdg","expiry":"2024-05-02T08:23:47.2499693-07:00","token_type":"Bearer"}
7:23AM INF cmd\cli\root\publish\publish.go:54 > OnConnecting event={"Code":0,"Reason":"connect called"}
7:23AM INF cmd\cli\root\publish\publish.go:57 > OnConnected event={"ClientID":"18a83098-7350-41ad-82e8-43a2b2501110","Data":null,"Version":"5.3.2"}
7:23AM INF cmd\cli\root\publish\publish.go:102 > published message
```

### Subscribing to a channel

```powershell
.\centrifugo-cli.exe subscribe  
```

```bash
7:25AM INF cmd\cli\root\subscribe\subscribe.go:36 > got token channel=connector:foobar token={"access_token":"eyJhbGciOiJFUzI1NiIsImtpZCI6IjBiMmNkMmU1NGM5MjRjZTg5ZjAxMGYyNDI4NjIzNjdkIiwidHlwIjoiSldUIn0.eyJjYXBzIjpbeyJhbGxvdyI6WyJwdWIiLCJzdWIiLCJwcnMiLCJoc3QiXSwiY2hhbm5lbHMiOlsiY29ubmVjdG9yOioiXSwibWF0Y2giOiJ3aWxkY2FyZCJ9XSwiY2xpZW50X2lkIjoiY2VudHJpZnVnby1jb25uZWN0b3Itc2EiLCJleHAiOjE3MTQ2NjM1MjYsImlhdCI6MTcxNDY1OTkyNiwiaXNzIjoiaHR0cDovL2xvY2FsaG9zdDo5ODAyIiwic3ViIjoiY2VudHJpZnVnby1jb25uZWN0b3Itc2EifQ.eMeTa2MY0b9p6ln_dzIrHJbnrdtesr4Fe44BdGzHMpOOXlq76zl3qUMfVMK7lZ87RPVtJhdz5-eDXmodxyFH8A","expiry":"2024-05-02T08:25:26.1015441-07:00","token_type":"Bearer"}
7:25AM INF cmd\cli\root\subscribe\subscribe.go:54 > OnConnecting channel=connector:foobar context=client event={"Code":0,"Reason":"connect called"}
7:25AM INF cmd\cli\root\subscribe\subscribe.go:117 > published message channel=connector:foobar
7:25AM INF cmd\cli\root\subscribe\subscribe.go:123 > sub.Subscribe channel=connector:foobar context=subscribe
7:25AM INF cmd\cli\root\subscribe\subscribe.go:57 > OnConnected channel=connector:foobar context=client event={"ClientID":"4d25aa4e-2671-4a2b-9988-261667e44547","Data":null,"Version":"5.3.2"}
7:25AM INF cmd\cli\root\subscribe\subscribe.go:105 > OnSubscribed channel=connector:foobar context=subscribe event={"Data":null,"Positioned":true,"Recoverable":true,"Recovered":false,"StreamPosition":{"Epoch":"VUix","Offset":13},"WasRecovering":false}
7:25AM INF cmd\cli\root\subscribe\subscribe.go:114 > OnPublication channel=connector:foobar context=subscribe event={"Data":"eyJpbnB1dCI6IntcImFcIjpcImJcIn0ifQ==","Info":{"ChanInfo":null,"Client":"d58e294c-aa19-484d-b407-f30b4a5ba510","ConnInfo":null,"User":"centrifugo-connector-sa"},"Offset":14,"Tags":null}
```

## Using CAPS tokens and explicit subscribe token

In this example we will publish using the ```caps``` claim and subscribe using the ```channels``` claim.`
The client:```connector-foobar``` is used to excercise connececting directly to a channel when the token has a ```channels``` claim.  

### Publishing a message

```powershell
.\centrifugo-cli.exe publish --message='{"a":"b"}'
```

### Subscribing to a channel by calling sub.Subscribe()

This is the example [jwt](https://jwt.io/#debugger-io?token=eyJhbGciOiJFUzI1NiIsImtpZCI6IjBiMmNkMmU1NGM5MjRjZTg5ZjAxMGYyNDI4NjIzNjdkIiwidHlwIjoiSldUIn0.eyJjaGFubmVscyI6WyJjb25uZWN0b3I6Zm9vYmFyIl0sImNsaWVudF9pZCI6ImNvbm5lY3Rvci1mb29iYXIiLCJleHAiOjE3MTQ2NjczNTksImlhdCI6MTcxNDY2Mzc1OSwiaXNzIjoiaHR0cDovL2xvY2FsaG9zdDo5ODAyIiwic3ViIjoiY29ubmVjdG9yLWZvb2JhciJ9.bhSf4sh3K6ntpIAxLNlC_mGVYZ-aSmKm6K966PCu1IQ17r4idivUIPYx_faqqnOO6RTHedRtzffIdtAa8jnBnQ) that is used to subscribe to the channel ```connector:foobar```.  

```powershell
.\centrifugo-cli.exe --oauth2-client-id=connector-foobar subscribe
```

This produces an error and doesn't get any OnPublication events.  

```console  
8:18AM INF cmd\cli\root\subscribe\subscribe.go:36 > got token channel=connector:foobar token={"access_token":"eyJhbGciOiJFUzI1NiIsImtpZCI6IjBiMmNkMmU1NGM5MjRjZTg5ZjAxMGYyNDI4NjIzNjdkIiwidHlwIjoiSldUIn0.eyJjaGFubmVscyI6WyJjb25uZWN0b3I6Zm9vYmFyIl0sImNsaWVudF9pZCI6ImNvbm5lY3Rvci1mb29iYXIiLCJleHAiOjE3MTQ2NjY3MzcsImlhdCI6MTcxNDY2MzEzNywiaXNzIjoiaHR0cDovL2xvY2FsaG9zdDo5ODAyIiwic3ViIjoiY29ubmVjdG9yLWZvb2JhciJ9.TDFYXxEHDnUCbxSyU7XpwnTW-jsir8wjd8B8qdf3MKE0Z3LBaCTBNnVVJgWRp_br5bclea4MMBf-f1dVDrpmng","expiry":"2024-05-02T09:18:57.8345396-07:00","token_type":"Bearer"}
8:18AM INF cmd\cli\root\subscribe\subscribe.go:54 > OnConnecting channel=connector:foobar context=client event={"Code":0,"Reason":"connect called"}
8:18AM INF cmd\cli\root\subscribe\subscribe.go:117 > published message channel=connector:foobar
8:18AM INF cmd\cli\root\subscribe\subscribe.go:123 > sub.Subscribe channel=connector:foobar context=subscribe
8:18AM INF cmd\cli\root\subscribe\subscribe.go:57 > OnConnected channel=connector:foobar context=client event={"ClientID":"38bd9807-22b9-4879-81b9-be86bde1c94e","Data":null,"Version":"5.3.2"}
8:18AM INF cmd\cli\root\subscribe\subscribe.go:70 > OnSubscribed channel=connector:foobar context=client event={"Channel":"connector:foobar","Data":null,"Positioned":true,"Recoverable":true,"Recovered":false,"StreamPosition":{"Epoch":"VUix","Offset":14},"WasRecovering":false}
8:18AM ERR cmd\cli\root\subscribe\subscribe.go:108 > OnError channel=connector:foobar context=subscribe event={"Error":{"Err":{"Code":105,"Message":"already subscribed","Temporary":false}}}
8:18AM INF cmd\cli\root\subscribe\subscribe.go:111 > OnUnsubscribed channel=connector:foobar context=subscribe event={"Code":105,"Reason":"already subscribed"}
```

### Subscribing to a channel by NOT calling sub.Subscribe()

This one we don't call;

```go
err = sub.Subscribe()
```

It still doesn't get the OnPublication events.  

```powershell
.\centrifugo-cli.exe --oauth2-client-id=connector-foobar subscribe --with-channel-token
```

```console
8:29AM INF cmd\cli\root\subscribe\subscribe.go:36 > got token channel=connector:foobar token={"access_token":"eyJhbGciOiJFUzI1NiIsImtpZCI6IjBiMmNkMmU1NGM5MjRjZTg5ZjAxMGYyNDI4NjIzNjdkIiwidHlwIjoiSldUIn0.eyJjaGFubmVscyI6WyJjb25uZWN0b3I6Zm9vYmFyIl0sImNsaWVudF9pZCI6ImNvbm5lY3Rvci1mb29iYXIiLCJleHAiOjE3MTQ2NjczNDAsImlhdCI6MTcxNDY2Mzc0MCwiaXNzIjoiaHR0cDovL2xvY2FsaG9zdDo5ODAyIiwic3ViIjoiY29ubmVjdG9yLWZvb2JhciJ9.tGqSCvmqx4_JNjtcKZ1gviJkzPDOjGMOeJ5bQjICwRFTJ68OxVpaUVu9TzziZpFB_zhjpIHpXOH4h7fZI674Og","expiry":"2024-05-02T09:29:00.9399201-07:00","token_type":"Bearer"}
8:29AM INF cmd\cli\root\subscribe\subscribe.go:54 > OnConnecting channel=connector:foobar context=client event={"Code":0,"Reason":"connect called"}
8:29AM INF cmd\cli\root\subscribe\subscribe.go:123 > sub.Subscribe channel=connector:foobar context=subscribe
8:29AM INF cmd\cli\root\subscribe\subscribe.go:57 > OnConnected channel=connector:foobar context=client event={"ClientID":"0405ea0b-3d5c-4f7e-8fd3-60429243b0ff","Data":null,"Version":"5.3.2"}
8:29AM INF cmd\cli\root\subscribe\subscribe.go:70 > OnSubscribed channel=connector:foobar context=client event={"Channel":"connector:foobar","Data":null,"Positioned":true,"Recoverable":true,"Recovered":false,"StreamPosition":{"Epoch":"VUix","Offset":15},"WasRecovering":false}
```
