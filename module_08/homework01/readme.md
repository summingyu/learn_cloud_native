现在你对 Kubernetes 的控制面板的工作机制是否有了深入的了解呢？
是否对如何构建一个优雅的云上应用有了深刻的认识，那么接下来用最近学过的知识把你之前编写的 http 以优雅的方式部署起来吧，你可能需要审视之前代码是否能满足优雅上云的需求。
作业要求：编写 Kubernetes 部署脚本将 httpserver 部署到 Kubernetes 集群，以下是你可以思考的维度。

- 优雅启动 -- readinessProbe
- 优雅终止 -- lifecycle
- 资源需求和 QoS 保证 --resources,
- 探活 -- livenessProbe
- 日常运维需求，日志等级
- 配置和代码分离 --创建config_map 记录 VERSION环境变量
[strong_begin] 提交地址： https://jinshuju.net/f/rJC4DG
截止日期：2022 年 4 月 10 日 23:59
