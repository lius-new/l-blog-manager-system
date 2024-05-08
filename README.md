### 博客服务

`go-zero`重构博客后台

#### TOKEN被浏览器正常解析

当后台生成token后放到一些在线解析token站点可以解析出正常内容出来, 通常token分为几个部分组成(头/载荷/签名).

这些在线解析并不关注签名所以还是能被解析出来token中的信息，所以在token不应该存放关键信息。事实上通过一些算法对下发的token再加密然后这些在线平台就不能正常解析了，大不了后台再解析前解密。

#### 统一响应

[https://go-zero.dev/docs/tutorials/customization/template](https://go-zero.dev/docs/tutorials/customization/template)