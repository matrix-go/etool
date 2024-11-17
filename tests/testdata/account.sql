CREATE TABLE `t_account` (
 `id` bigint NOT NULL AUTO_INCREMENT COMMENT '主键',
 `uid` bigint NOT NULL COMMENT '用户关联id',
 `balance` decimal(19,4) NOT NULL COMMENT '账户余额',
 `ctime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
 `utime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间',
 PRIMARY KEY (`id`) COMMENT '主键自增'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci