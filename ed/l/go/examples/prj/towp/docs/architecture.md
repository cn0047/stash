Architecture
-

Current high level architectural overview:

Regular REST API request:

````mermaid
flowchart TD

client[Client]
cloudLB[Cloud Load Balancing]
cloudRun[Cloud Run]
pubSub[PubSub]
spanner[Cloud Spanner]

client -- Request via Identity-Aware Proxy --> cloudLB
cloudLB --> cloudRun
cloudRun --> pubSub
cloudRun --> spanner
````

Internal scheduled event:

````mermaid
flowchart TD
  
cloudScheduler[Cloud Scheduler]
cloudLB[Cloud Load Balancing]
cloudRun[Cloud Run]
pubSub[PubSub]
spanner[Cloud Spanner]

cloudScheduler -- Request via Identity-Aware Proxy --> cloudLB
cloudLB --> cloudRun
cloudRun --> pubSub
cloudRun --> spanner
````
