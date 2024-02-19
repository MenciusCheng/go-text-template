# Generator API V2 版

## Generator 生成器

定义 解析器 和 模版 部件。使用 Data map[string]interface{} 存储数据，类似**JSON**操作方式。

生成路径如下：

文本 => 解析器(Parser) => Data(Json Map) + 模板 => 执行器(Executor) => 结果

## Generator Stream 流式生成器

通过流式去定义生成器的每个操作流程。使用 Data [][]string 存储数据，类似**集合**操作方式。

流程主要分为三类：读取、转化和输出。细分每个操作流程的方法。
