# 此层的作用
- 这一层定义了持久化领域服务和领域事件相关实体的接口。具体实现在基础层中实现。
- 这么做是让领域层独立，不依赖于基础层的实现，而让基础层为领域层服务。