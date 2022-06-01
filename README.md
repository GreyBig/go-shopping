# go-shopping

# 准备数据库和用户表

## 创建数据库

```sql
create database shopping_db;
```

## 准备用户表

```sql
CREATE TABLE `user` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '用户id（主键）',
  `username` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '用户名称',
  `password` varchar(80) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '用户密码',
  `status` tinyint(20) NOT NULL DEFAULT '1' COMMENT '用户状态，1表示正常，0表示暂停',
  `created` char(50) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '创建时间',
  `updated` char(50) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE
) ;
```





## 测试验证码

在浏览器中输入如下地址：http://localhost:8000/web/captcha，输出内容如下：

```json
{"code":200,"message":"操作成功","data":{"captchaId":"sxXVGXG6rcunegcvUXVu","captchaImg":"data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAHgAAAAoCAMAAAACNM4XAAAA81BMVEUAAABpWhWThD91ZiGjlE+UhUCzpF+2p2JtXhl6ayaNfjlZSgW+r2p1ZiGun1qfkEu9rmnIuXS3qGOwoVxkVRCPgDvSw37czYhYSQSVhkGRgj20pWB1ZiGPgDuik06+r2pxYh2un1q4qWRWRwKpmlXMvXjh0o2VhkGrnFdzZB9kVRBsXRjUxYCFdjGlllFURQCFdjGej0qWh0K3qGNmVxLHuHOOfzray4ZtXhmNfjm+r2rEtXCvoFu7rGfXyIPEtXBdTgnSw37TxH9tXhnp2pXQwXyml1J0ZSDk1ZDXyIN7bCfczYjez4qun1rEtXDNvnmrnFfklloJAAAAAXRSTlMAQObYZgAAA0NJREFUeJysmM1uwjgQx2dArcSBS6GHBsoFocASqYSWpgghgqA0Egrv/zir+CMef2HvtnPAduzMb/5hbMcBl+1kZeTsVlaJsg6Mi7TdTpBHIxe5p7gVJ9f1X5FF+TH6MHrOAL1eD2aSLErJ7YQ8b9nvTzACi3tuyDCbzZzDOx2DvGC/iKzYAGy3Dfnnh5Jf/wmGwRU35uZaiheLhQJvNhuX4tdXRX6ICSHKFgyrFLuMcB9+Qe62NZSlBIft8RfcbtegNRVEKGK4j5z8FBx5UlXuHpniFRKVa2w6iyKKzLlPIfLpdJJUZI8WWW21Wqlnu16vIVaxtEjF5l+IuKLNNagJZYxzO41NCG574s+6s/JA3BejUlGO2u/3bS5ZWqqq8jhDF98a+qk3JwSAuG/TiRU6qPI9V4QIiZ+fGnkymfBKzudOywXn1JXtow/vtjE4FHPveZ4jEtcosBR9aHuPx2Msk3HHYztMqcrgSiQRfTgc2gdqKg6R3/T2G9KVCaAkYNRCqIViPuBkx2+0twbojZKRNps7y7IkC7OqIHlDwD5iu+B4yXzbpGRePAuXhMtuLEl+6z21TLZ+vyE7Eli/tLX6Gff52RyIiDcbR2st2eVxYIt2kx0L5O12Qy0QOwKv58FgEIaC2uJoyiLeEHFMhjhoXvTgfje9naVtRrMaEOl8w6UN5jke8OzrA6UkyzIqjCi+Ai6XSwN7BWhXt/+6JRlZm3HXdbtWSMT1CrA03LOLQYQZUccYTCNoJ6m66erycfWoMdEat9PRnps2W7D2RPu/rKuDOua0JZUeumL12EuIy94S+XonUnHvIiOyoww49X6zX+2l/+XlDrkE8l6MkCIM+cuFBlabd8+zj39/N2TjmHOPW5a0mabpcDgEQ7FYOxAxh9Q3QRyK71qpN1Ngik0TsDzP0zTyzezPjNHyJjKHvRtt/owTAJiLK1k8aWpf8n0reH/XyAnPqiRJYD6fs6U5y6LJ0+lUncEF1/mtAEzFSZIQxXwzOvu5cuu8tGT11aEl3w2VfVLYCKCygTzbO63o9zn5crmQy5L7dZcouLudOIg77Ow5gRdFYSkm9vUVRQb/Qdx7Ag+c62K4QXsA+DcAAP//avVCZeS6TCYAAAAASUVORK5CYII="}}

```

## 测试用户登录

这里我们先把验证码逻辑注释掉，即修改登录controller，注释掉如下代码，并重新启动项目：

```go
func WebUserLogin(c *gin.Context) {
	var param models.AdminWebUserLoginVO
	if err := c.ShouldBind(&param); err != nil {
		response.Failed("请求参数无效", c)
		return
	}

	// 检查验证码
/* 	if !common.VerifyCaptcha(param.CaptchaId, param.CaptchaValue) {
		response.Failed("验证码错误", c)
		return
	} */
	// 生成token
	uid := user.Login(param)
	if uid > 0 {
		token, _ := common.GenerateToken(param.Username)
		userInfo := models.AdminWebUserInfo{
			Uid:   uid,
			Token: token,
		}
		response.Success("登录成功", userInfo, c)
		return
	}
	response.Failed("用户名或密码错误", c)
}
```

## 使用postman测试

#### 在数据库user表添加一条数据

```sql
insert into user(username, password)values('admin','123')
```

#### 测试json

```json
{
    "username":"admin",
    "password":"123",
    "captchaId":"111",
    "captchaValue":"1111"
}
```
