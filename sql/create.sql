CREATE TABLE `t_dynamic_form_config` (
   `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
   `name` varchar(32) NOT NULL COMMENT '输入框名',
   `input_type` tinyint(2) NOT NULL DEFAULT '0' COMMENT '用户输入类型，0为普通短文本，具体见常量定义',
   `check_type` tinyint(2) NOT NULL DEFAULT '0' COMMENT '校验用户输入类型，0为不校验，具体见常量定义',
   `size_range` varchar(64) NOT NULL DEFAULT '' COMMENT '大小范围，文本或数字，数据存json，可选字段仅有min和max',
   `sub_option` varchar(2048) NOT NULL DEFAULT '' COMMENT '子选项列表，数据存json，[[id=>int,name=>string]…]',
   `hint` varchar(128) NOT NULL DEFAULT '' COMMENT '输入框背景提示',
   `err_msg` varchar(256) NOT NULL DEFAULT '' COMMENT '用户输入出错时的文案提示',
   `create_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
   PRIMARY KEY (`id`)
 ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='动态表单数据配置'

 CREATE TABLE `t_dynamic_form_issued` (
   `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
   `business_id` tinyint(4) NOT NULL COMMENT '业务id，各个业务需保持唯一性',
   `module_id` tinyint(2) NOT NULL COMMENT '该输入框所在模块id，用于划分同一业务下不同模块',
   `showcase_id` tinyint(2) NOT NULL DEFAULT '0' COMMENT '同一模块下根据不同业务的展示方案',
   `form_id` int(11) NOT NULL COMMENT '对应表单数据配置表id',
   `parent_id` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '父级输入框id，0表示为顶级输入框',
   `is_required` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否必填:0非必填/1必填',
   `sort_index` tinyint(2) NOT NULL COMMENT '该输入框排序index',
   `tag` varchar(32) NOT NULL DEFAULT '' COMMENT '对于一些有特殊处理的输入框的标签',
   `table_name` varchar(32) NOT NULL COMMENT '对应提交表单存储数据表名',
   `field_name` varchar(32) NOT NULL COMMENT '对应提交表单存储数据表的字段名',
   `status` tinyint(1) NOT NULL DEFAULT '0' COMMENT '状态:0正常/1停用',
   `create_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
   PRIMARY KEY (`id`),
   UNIQUE KEY `idx_bid_mid_sid_si` (`business_id`,`module_id`,`showcase_id`,`sort_index`)
 ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='动态表单业务下发配置'
