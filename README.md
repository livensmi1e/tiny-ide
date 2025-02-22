Online IDE and Code Execution

This is a self-learning project.

Based on wonderful research paper from Judge0

-   [Robust and Scalable Online Code Execution System](https://www.researchgate.net/publication/346751837_Robust_and_Scalable_Online_Code_Execution_System)

Reference repositories

-   [cpinitiative/ide](https://github.com/cpinitiative/ide)
-   [judge0/judge0](https://github.com/judge0/judge0)
-   [judge0/ide](https://github.com/judge0/ide)
-   [criyle/go-judge](https://github.com/criyle/go-judge)
-   [seymuromarov/go-sandbox](https://github.com/seymuromarov/go-sandbox/)
-   [Narasimha1997/gopg](https://github.com/Narasimha1997/gopg)

Approaches & further development

-   [x] Use pure Docker container for sandbox
-   [ ] Use gVisor: [gvisor.dev](https://gvisor.dev/)
-   [ ] Use isolate

Design

-   Executor server is sandbox for executing code
-   Workers communicate with executor servers via gRPC
-   Use nginx as gRPC load balancer - a.k.a server side load balancer

Development

-   Make `.env.development` file based on `.env.example`
