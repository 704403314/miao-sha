# miaosha
模拟秒杀项目：
为了下单不扣减真实库存，支付完成才扣库存，本项目库存有两种，一个是真实库存，一个是预库存
1. 预先向redis队列里添加库存(不是真实库存，为了支付时预扣库存使用)
2. 获取真实库存，大于0，可以进入下订单页面
3. 下单页面 点击支付，判断预库存队列是否为空，不为空，从预库存队列删除一个数据，不报错，则进入支付页面
4. 进入支付页面，取消订单 会向预库存队列添加一个数据
5. 支付成功  扣减真实库存