# rvav
Has 2 entities: agent and orchestrator. Orchestrator splits equations, agent calculates. By default orchestrator launches on :8080, agent on :8081

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
# if you want gui-ish thing too
cd ../front && go build
# will give you front binary, launch it along others and open in browser
```
You'll also need psql running. Just set it up and make it work.<br>
After setting everything up, fill in /config/orchestrator.json with your data on everything<br>
To launch:
```sh
# first you need agent running, is on port 8001
./agent/agent
# and then orchestrator, is on port 8000
./orchestrator/orchestrator
# and frontend, is on port 7999
./front/front
```

# "How does this work?"
![image](./images/total.png)

Close-ups:<br>
Orchestrator loop
![image](./images/orchestrator.png)
Eval loop, i.e how does it calculate stuff
![image](./images/eval.png)
And agent
![image](./images/agent.png)

More detailed info on each:<br>
[For agent](/docs/agent.md)<br>
[For orchestrator](/docs/orch.md)<br>

