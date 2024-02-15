# rvav
Has 2 entities: agent and orchestrator. Orchestrator splits equations, agent calculates. By default orchestrator launches on :8080, agent on :8081

ToC:
<ul>
<li>[Installation](#installation)</li>
<li>[Endpoints info](#endpoints-wip)</li>
<li>[Internals](#how-does-this-work)</li>
</ul>

# Installation
```sh
git clone https://github.com/wdmvv/rvav
cd rvav
# building agent
cd agent
go get && go build
# building orchestrator
cd ../orchestrator
go get && go build
# you'll get 2 binaries - agent and orchestrator
```

# Endpoints (WIP)
All requests have json body & response<br>
## Agent:<br>
### /eval
Evalutes expression based on given info & has timeout<br>
Request:<br>
```js
{
    "op1": float64, // - first operand
    "op2": float64, // - second operand
    "sign": str, // - sign, can be one of "+-*/"
    "timeout": int, // - operation execution timeout
}
```
Response:<br>
```js
{
    "result": float64, // - operation result,
    "errmsg": string, // - error message if any
}
```
### /status
Checks if agent is up, basically a pingpong<br>
Request: None<br>
Response:<br>
```js
{
    "msg": string // just says "agent is running!"
}
```

# "How does this work?"
If I did not change anything these are orchestrator endpoints, visualized 
![image](./images/orchestrator.png)
Eval loop, i.e how does it calculate stuff
![image](./images/eval.png)
And agent
![image](./images/agent.png)