File that describes how agent works & its endpoints<br>
Note: even though some endpoints should've been POST, all are GET
# Internals
Basically, there are 2 handlers - status and eval. Status is just simple pingpong, eval calculates expression<br>
Calculation is simple too, it tries to acquire semaphore to limit number of running operations, then calculates and imitates job by sleeping set time

# Endpoints
## /status
No req. body is expected<br>
Response:
```js
{
  "msg": str // string that says that it is up
}
```
## /eval
Req. body:
```js
{
  "op1": float64,
  "op2": float64, // two parameters to work with
  "sign": string, // operation sign, must be one of "+-*/"
  "timeout": int // operation timeout, i.e how for how long job is imitated
}
```
Response:
```js
{
  "result": float64, //op result, 0 if error
  "errmsg": string //error
}
```

