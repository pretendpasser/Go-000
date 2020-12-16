学习笔记

此为之前仿写的小项目，有几个问题。

1.此项目的目录结构是否符合标准？每一层的实现是否合理

2.此项目是否可以将middleware serializer server service全部放在internal文件夹里作为对外不可见的包？

3.是否通过grpc来改装api层，能否提供一个思路？

通过使用gin框架，先是通过server层的路由去调api层的接口并通过api层调用service层的具体功能实现。
