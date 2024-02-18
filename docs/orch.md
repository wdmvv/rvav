Orchestrator internals & endpoints

# Internals
It is described on picture, but this one is more informative:<br>
1. It reads config. Config has many fields:<br>
```js
{
    "plus": int, 
    "minus": int,
    "mul": int,
    "div": int, // operation timeouts, in ms
    "timeout": int, // total expression calculation timeout 
    "agent_port": int, // on which port agent is running
    // DB config
    "usedb": bool // whether you want to use db or not
    "user": string,
    "password": string,
    "dbname": string,
    "tablename": string
}
```
2. It connects to DB and creates table. Table has following schema:<br>
```sql
  expr VARCHAR(255) PRIMARY KEY
  result REAL
```
Where expr is expression, result it what expression evaluates to<br>
3. It starts server. There are many handlers on it but ill cover eval loop only, which is on /addexpr<br>
4. Get /addexpr request<br>
5. Converts expression into postfix<br>
6. Checks if expression is in db already, if yes then returns it, if no loop continues<br>
7. Calculates postfix by sending /eval request to the agent on each operation<br>
8. Writes to db if no error occurred<br>
9. Returns result to who send it<br>

# Endpoints
## /status
Pingpong. No req. body<br>
Response:
```js
{
  "msg": string //that it is up
}
```
## /timeouts
Tells all current operation timeouts. No req. body<br>
Response:
```js
{
  "plus": int,
  "minus": int,
  "mul": int,
  "div": int
}
```
## /chtime
Changes timeout of one operation(sign)<br>
Req. body:
```js
{
  "sign": string, // can be either one of "+-*/" or ["plus", "minus", "mul", "div"]
  "ms": int //timeout, ms
}
```
Response:
```js
{
  "errmsg": string //if error
}
```
## /addexpr
Calculates operation. Supports 4 operations and brackets. No floats, ints only.<br>
Req. body:
```js
{
  "expr": string, //expression to calculate
  "id": string, // req. id
}
```
Response:
```js
{
  "result": float64, //calculation result
  "errmsg": string //err message if any
}
```
## /jobs
Shows all expressions that were ever added. No req. body<br>
Response:
```js
{
    "running": {
        <id1>:{
            "expr": string,
            "start": string,
            "end": string,
        },
        <id2>:{
            "expr": string,
            "start": string,
            "end": string,
        } // and so on, completed and failed have the same structure
    },
    "completed":{}
    "failed":{}
}
```

# Structure 
/app - endpoint bindings<br>
/calc/ - for calculating expression, i.e converting to postfix, checking db and requesting agent<br>
/config/ - config file & file for reading and parsing it<br>
/db/ - for connection to db & various operations on connection struct<br>
/handlers/ - handlers, *_h.go is a handler for that endpoint<br>
/logging/ - small thingy for printing logs<br>

