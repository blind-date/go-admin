ALTER TABLE `sys_user` ADD COLUMN `real_name` varchar(128) DEFAULT NULL COMMENT '真实姓名' AFTER `nick_name`;
ALTER TABLE `sys_user` ADD COLUMN `wx_id` varchar(128) DEFAULT NULL COMMENT '微信号' AFTER `real_name`;
ALTER TABLE `sys_user` ADD COLUMN `wx_image` varchar(255) DEFAULT NULL COMMENT '微信二维码' AFTER `wx_id`;


DROP TABLE IF EXISTS `case`;
CREATE TABLE `case` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '编码',
  `mr` varchar(255) DEFAULT NULL COMMENT '男士称呼',
  `mr_id` bigint DEFAULT NULL COMMENT '男士id',
  `miss` varchar(255) DEFAULT NULL COMMENT '女士称呼',
  `miss_id` bigint DEFAULT NULL COMMENT '女士id',
  `sort` int DEFAULT NULL COMMENT '排序',
  `wish` varchar(255) DEFAULT NULL COMMENT '贺词',
  `status` int DEFAULT NULL COMMENT '状态',
  `marry_time` timestamp NULL DEFAULT NULL COMMENT '结婚日期',
  `img` varchar(255) DEFAULT NULL COMMENT '结婚照',
  `dept_id` bigint DEFAULT NULL COMMENT '城市ID',
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  `create_by` int(11) unsigned DEFAULT NULL,
  `update_by` int(11) unsigned DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_case_deleted_at` (`deleted_at`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COMMENT='案例';


DROP TABLE IF EXISTS `article_category`;
CREATE TABLE `article_category` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '主键编码',
  `name` varchar(255) DEFAULT NULL COMMENT '名称',
  `status` int DEFAULT NULL COMMENT '状态',
  `remark` varchar(255) DEFAULT NULL COMMENT '备注',
  `create_by` bigint DEFAULT NULL COMMENT '创建者',
  `update_by` bigint DEFAULT NULL COMMENT '更新者',
  `created_at` datetime(3) DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime(3) DEFAULT NULL COMMENT '最后更新时间',
  `deleted_at` datetime(3) DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`),
  KEY `idx_article_category_create_by` (`create_by`),
  KEY `idx_article_category_update_by` (`update_by`),
  KEY `idx_article_category_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='文章分类';

DROP TABLE IF EXISTS `article`;
CREATE TABLE `article` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '主键编码',
  `cate_id` int DEFAULT NULL COMMENT '分类id',
  `name` varchar(255) DEFAULT NULL COMMENT '名称',
  `status` int DEFAULT NULL COMMENT '状态',
  `img` varchar(255) DEFAULT NULL COMMENT '图片',
  `content` text COMMENT '内容',
  `remark` varchar(255) DEFAULT NULL COMMENT '备注',
  `jump_url` varchar(255) DEFAULT NULL COMMENT '跳转地址',
  `click_count` int DEFAULT NULL COMMENT '点击数',
  `virtual_click_count` int DEFAULT NULL COMMENT '虚拟点击数',
  `sort` int DEFAULT NULL COMMENT '排序',
  `dept_id` bigint DEFAULT NULL COMMENT '城市ID',
  `create_by` bigint DEFAULT NULL COMMENT '创建者',
  `update_by` bigint DEFAULT NULL COMMENT '更新者',
  `created_at` datetime(3) DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime(3) DEFAULT NULL COMMENT '最后更新时间',
  `deleted_at` datetime(3) DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`),
  KEY `idx_article_create_by` (`create_by`),
  KEY `idx_article_update_by` (`update_by`),
  KEY `idx_article_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='文章';


DROP TABLE IF EXISTS `article_comment`;
CREATE TABLE `article_comment` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '主键编码',
  `article_id` int DEFAULT NULL COMMENT '文章id',
  `member_id` bigint DEFAULT NULL COMMENT '会员id',
  `content` text COMMENT '内容',
  `status` int DEFAULT NULL COMMENT '状态',
  `create_by` bigint DEFAULT NULL COMMENT '创建者',
  `update_by` bigint DEFAULT NULL COMMENT '更新者',
  `created_at` datetime(3) DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime(3) DEFAULT NULL COMMENT '最后更新时间',
  `deleted_at` datetime(3) DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`),
  KEY `idx_article_comment_create_by` (`create_by`),
  KEY `idx_article_comment_update_by` (`update_by`),
  KEY `idx_article_comment_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='文章评论';

DROP TABLE IF EXISTS `member`;
CREATE TABLE `member` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '编码',
  `dept_id` bigint DEFAULT NULL COMMENT '城市ID',
  `level` bigint DEFAULT 1 COMMENT '会员等级',
  `level_deadline` datetime DEFAULT NULL COMMENT '会员到期时间',
  `password` varchar(128) DEFAULT NULL COMMENT '密码',
  `sex` varchar(255) DEFAULT NULL COMMENT '性别',
  `nick_name` varchar(128) DEFAULT NULL COMMENT '昵称',
  `real_name` varchar(128) DEFAULT NULL COMMENT '真实姓名',
  `card_id` varchar(128) DEFAULT NULL COMMENT '身份证ID',
  `card_img_obverse` varchar(255) DEFAULT NULL COMMENT '身份证正面',
  `card_img_back` varchar(255) DEFAULT NULL COMMENT '身份证背面',
  `card_img_hand` varchar(255) DEFAULT NULL COMMENT '手持身份证',
  `card_audit_status`varchar(20) DEFAULT NULL COMMENT '身份验证通过',
  `school_roll` varchar(255) DEFAULT NULL COMMENT '学籍PDF_URL',
  `school_education` varchar(255) DEFAULT NULL COMMENT '学历PDF_URL',
  `school_audit_status`varchar(20) DEFAULT NULL COMMENT '学历审核',
  `blind_date_type` varchar(20) DEFAULT NULL COMMENT '相亲类型',
  `matchmaker` bigint DEFAULT NULL COMMENT '红娘ID',
  `wx_id` varchar(128) DEFAULT NULL COMMENT '微信号',
  `open_id` varchar(128) DEFAULT NULL COMMENT 'openId',
  `phone` varchar(25) DEFAULT NULL COMMENT '手机号',
  `channel` varchar(255) DEFAULT NULL COMMENT '注册渠道',
  `other_contact` varchar(255) DEFAULT NULL COMMENT '其他联系方式',
  `activity_show` varchar(4) DEFAULT NULL COMMENT '活动现场展示',
  `head_img` varchar(255) DEFAULT NULL COMMENT '头像',
  `head_img_show` varchar(4) DEFAULT NULL COMMENT '头像公开',
  `album_show` varchar(4) DEFAULT NULL COMMENT '相册公开',
  `has_album` varchar(4) DEFAULT '0' COMMENT '是否有相册',
  `birth_date` datetime NULL DEFAULT NULL COMMENT '出生日期',
  `household` varchar(255) DEFAULT NULL  COMMENT '户籍编码',
  `household_name` varchar(255) DEFAULT NULL  COMMENT '户籍',
  `address` varchar(255) DEFAULT NULL  COMMENT '现居住地',
  `nation` varchar(20) DEFAULT NULL COMMENT '民族',
  `education` varchar(20) DEFAULT NULL COMMENT '学历',
  `blood` varchar(20) DEFAULT NULL COMMENT '血型',
  `health` varchar(255) DEFAULT NULL COMMENT '健康状况',
  `company` varchar(255) DEFAULT NULL COMMENT '工作单位',
  `organization_nature` varchar(20) DEFAULT NULL COMMENT '单位性质',
  `industry` varchar(20) DEFAULT NULL COMMENT '所属行业',
  `post` varchar(20) DEFAULT NULL COMMENT '职务/岗位',
  `work_status` varchar(20) DEFAULT NULL COMMENT '工作状态',
  `house_situation` varchar(20) DEFAULT NULL COMMENT '购房情况',
  `car_situation` varchar(20) DEFAULT NULL COMMENT '购车情况',
  `height` int DEFAULT NULL COMMENT '身高',
  `weight` int DEFAULT NULL COMMENT '体重',
  `nature` varchar(100) DEFAULT NULL COMMENT '性格',
  `hobby` varchar(100) DEFAULT NULL COMMENT '爱好',
  `shape` varchar(20) DEFAULT NULL COMMENT '体型',
  `major` varchar(20) DEFAULT NULL COMMENT '专业',
  `lifestyle` varchar(20) DEFAULT NULL COMMENT '生活作息',
  `appearance` varchar(20) DEFAULT NULL COMMENT '外貌自评',
  `faith` varchar(20) DEFAULT NULL COMMENT '宗教信仰',
  `marry_situation` varchar(20) DEFAULT NULL COMMENT '婚史',
  `blind_date_way` varchar(20) DEFAULT NULL COMMENT '偏爱约会方式',
  `smoke` varchar(20) DEFAULT NULL COMMENT '吸烟',
  `drink` varchar(20) DEFAULT NULL COMMENT '喝酒',
  `marry_time` varchar(20) DEFAULT NULL COMMENT '结婚时间',
  `need_child` varchar(20) DEFAULT NULL COMMENT '是否想要小孩',
  `wedding_style` varchar(20) DEFAULT NULL COMMENT '期待婚礼形式',
  `another_focus` varchar(20) DEFAULT NULL COMMENT '希望对方看重',
  `cook_situation` varchar(20) DEFAULT NULL COMMENT '厨艺状况',
  `house_work`varchar(20) DEFAULT NULL COMMENT '家务分工',
  `live_with_parent` varchar(20) DEFAULT NULL COMMENT '是否愿与对方父母同住',
  `divorce_certificate_img` varchar(255) DEFAULT NULL COMMENT '离婚证',
  `salary` varchar(20) DEFAULT NULL COMMENT '年薪/收入',
  `deposit_money`varchar(20) DEFAULT NULL COMMENT '存款',
  `child_rank`varchar(20) DEFAULT NULL COMMENT '家中排行',
  `family` varchar(100) DEFAULT NULL COMMENT '家庭成员',
  `language` varchar(100) DEFAULT NULL COMMENT '语言',
  `introduction` varchar(255) DEFAULT NULL COMMENT '自我简介',
  `parent_situation` varchar(20) DEFAULT NULL COMMENT '父母情况',
  `father_work` varchar(20) DEFAULT NULL COMMENT '父亲工作',
  `mother_work` varchar(20) DEFAULT NULL COMMENT '母亲工作',
  `parent_money` varchar(20) DEFAULT NULL COMMENT '父母经济',
  `parent_medical` varchar(20) DEFAULT NULL COMMENT '父母社保',
  `view_count` int DEFAULT NULL COMMENT '浏览数/热度',
  `integrity` int DEFAULT NULL COMMENT '资料完整度',
  `red_line_count` int DEFAULT NULL COMMENT '红线数量',
  `status` varchar(20) DEFAULT NULL COMMENT '状态',
  `recent_login` timestamp DEFAULT NULL COMMENT '最近登录时间',
  `longitude` decimal(10,3) NOT NULL COMMENT '经度',
  `latitude` decimal(10,3) NOT NULL COMMENT '维度',
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  `create_by` int(11) unsigned DEFAULT NULL,
  `update_by` int(11) unsigned DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_member_create_by` (`create_by`),
  KEY `idx_member_update_by` (`update_by`),
  KEY `idx_member_created_at` (`created_at`),
  KEY `idx_member_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COMMENT='会员基础信息';


CREATE TABLE `member_spouse` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '编码',
  `member_id` bigint DEFAULT NULL COMMENT '会员id',
  `refuse_industries` varchar(50) DEFAULT NULL COMMENT '不接受的择偶职业',
  `height_from` int DEFAULT NULL COMMENT '择偶最低身高要求',
  `height_to` int DEFAULT NULL COMMENT '择偶最高身高要求',
  `weight_from` int DEFAULT NULL COMMENT '择偶最低体重要求',
  `weight_to` int DEFAULT NULL COMMENT '择偶最高体重要求',
  `appearance` varchar(20) DEFAULT NULL COMMENT '择偶外貌要求',
  `nature` varchar(100) DEFAULT NULL COMMENT '择偶性格要求',
  `post` varchar(20) DEFAULT NULL COMMENT '职务/岗位',
  `age_from` int DEFAULT NULL COMMENT '择偶年龄最小要求',
  `age_to` int DEFAULT NULL COMMENT '择偶年龄最大要求',
  `marry_situation` varchar(20) DEFAULT NULL COMMENT '择偶婚史要求',
  `education` varchar(20) DEFAULT NULL COMMENT '择偶学历',
  `salary` varchar(20) DEFAULT NULL COMMENT '择偶年薪',
  `company` varchar(255) DEFAULT NULL COMMENT '择偶工作单位要求',
  `face_style` varchar(20) DEFAULT NULL COMMENT '择偶见面方式',
  `focus_requirement` varchar(100) DEFAULT NULL COMMENT '对择偶的关注点',
  `focus_shortcoming` varchar(100) DEFAULT NULL COMMENT '对择偶不能接受的缺点',
  `house_requirement` varchar(100) DEFAULT NULL COMMENT '对择偶的住房要求',
  `car_requirement` varchar(100) DEFAULT NULL COMMENT '对择偶的购车要求',
  `has_child_requirement` varchar(20) DEFAULT NULL COMMENT '对择偶是否有小孩要求',
  `area_requirement` varchar(20) DEFAULT NULL COMMENT '对择偶地域要求',
  `remark` varchar(255) DEFAULT NULL COMMENT '对择偶要求的备注',
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  `create_by` int(11) unsigned DEFAULT NULL,
  `update_by` int(11) unsigned DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_member_spouse_create_by` (`create_by`),
  KEY `idx_member_spouse_update_by` (`update_by`),
  KEY `idx_member_spouse_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COMMENT='会员择偶要求';


DROP TABLE IF EXISTS `member_photo_album`;
CREATE TABLE `member_photo_album` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `member_id` bigint DEFAULT NULL COMMENT '会员id',
  `img` varchar(255) DEFAULT NULL COMMENT '照片',
  `img_show` varchar(4) DEFAULT NULL COMMENT '是否公开',
  `sort` int DEFAULT NULL COMMENT '排序',
  `status` varchar(20) DEFAULT NULL COMMENT '审核状态',
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  `create_by` int(11) unsigned DEFAULT NULL,
  `update_by` int(11) unsigned DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idxmember_photo_album_create_by` (`create_by`),
  KEY `idx_member_photo_album_update_by` (`update_by`),
  KEY `idx_member_photo_album_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COMMENT='会员相册';


DROP TABLE IF EXISTS `member_level`;
CREATE TABLE `member_level` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `name` varchar(50) DEFAULT NULL COMMENT '名称',
  `icon_big` varchar(255) DEFAULT NULL COMMENT '大图标',
  `icon_middle` varchar(255) DEFAULT NULL COMMENT '中图标',
  `icon_small` varchar(255) DEFAULT NULL COMMENT '小图标',
  `focus_count` int DEFAULT NULL COMMENT '关注数量',
  `price` decimal(10,2) DEFAULT NULL COMMENT '价格',
  `price_name` varchar(50) DEFAULT NULL COMMENT '价格名称',
  `days` int DEFAULT NULL COMMENT '会员等级有效期天数',
  `rights`  varchar(255) DEFAULT NULL COMMENT '权益',
  `weight` int DEFAULT NULL COMMENT '权重',
  `status` int DEFAULT NULL COMMENT '状态',
  `remark` varchar(255) DEFAULT NULL COMMENT '备注',
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  `create_by` int(11) unsigned DEFAULT NULL,
  `update_by` int(11) unsigned DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_member_level_create_by` (`create_by`),
  KEY `idx_member_level_update_by` (`update_by`),
  KEY `idx_member_level_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COMMENT='会员级别';

INSERT INTO `member_level` (`id`, `name`,`focus_count`, `status`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`, `deleted_at`) VALUES (1, '普通会员', 10,1, '普通会员',1,1, '2021-07-05 08:52:06.067', '2021-07-05 08:52:06.067', NULL);

DROP TABLE IF EXISTS `member_level_right`;
CREATE TABLE `member_level_right` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `name` varchar(50) DEFAULT NULL COMMENT '名称',
  `icon` varchar(255) DEFAULT NULL COMMENT '图标',
  `brief` varchar(100) DEFAULT NULL COMMENT '简介，100字',
  `detail` varchar(2000) DEFAULT NULL COMMENT '详细，2000字',
  `status` int DEFAULT NULL COMMENT '状态',
  `remark` varchar(255) DEFAULT NULL COMMENT '备注',
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  `create_by` int(11) unsigned DEFAULT NULL,
  `update_by` int(11) unsigned DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_member_level_create_by` (`create_by`),
  KEY `idx_member_level_update_by` (`update_by`),
  KEY `idx_member_level_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COMMENT='会员权益';

-- DROP TABLE IF EXISTS `member_level_price`;
-- CREATE TABLE `member_level_price` (
--   `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
--   `name` varchar(50) DEFAULT NULL COMMENT '名称',
--   `price` decimal(10,2) DEFAULT NULL COMMENT '价格',
--   `days`  int DEFAULT NULL COMMENT '有效期，单位天数',
--   `status` int DEFAULT NULL COMMENT '状态',
--   `remark` varchar(255) DEFAULT NULL COMMENT '备注',
--   `created_at` timestamp NULL DEFAULT NULL,
--   `updated_at` timestamp NULL DEFAULT NULL,
--   `deleted_at` timestamp NULL DEFAULT NULL,
--   `create_by` int(11) unsigned DEFAULT NULL,
--   `update_by` int(11) unsigned DEFAULT NULL,
--   PRIMARY KEY (`id`),
--   KEY `idx_member_level_price_create_by` (`create_by`),
--   KEY `idx_member_level_price_update_by` (`update_by`),
--   KEY `idx_member_level_price_deleted_at` (`deleted_at`)
-- ) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COMMENT='会员级别价格';


DROP TABLE IF EXISTS `member_level_order`;
CREATE TABLE `member_level_order` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `member_id` bigint DEFAULT NULL COMMENT '会员ID',
  `level_id` bigint DEFAULT NULL COMMENT '会员级别ID',
  `order_id` varchar(255)  DEFAULT NULL COMMENT '订单ID',
  `price` decimal(10,2) DEFAULT NULL COMMENT '红线支付价格',
  `days` bigint DEFAULT NULL COMMENT '会员等级有效期天数',
  `status` int DEFAULT NULL COMMENT '订单状态',
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  `create_by` int(11) unsigned DEFAULT NULL,
  `update_by` int(11) unsigned DEFAULT NULL,  
  PRIMARY KEY (`id`),
  KEY `idx_member_level_order_create_by` (`create_by`),
  KEY `idx_member_level_order_update_by` (`update_by`),
  KEY `idx_member_level_order_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COMMENT='会员级别订单';


DROP TABLE IF EXISTS `wechat_callback_log`;
CREATE TABLE `wechat_callback_log` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `order_id` varchar(255)  DEFAULT NULL COMMENT '订单ID',
  `order_type` int DEFAULT NULL COMMENT '订单类型',
  `content` varchar(2048) DEFAULT NULL COMMENT 'content',
  `notify` varchar(2048)  DEFAULT NULL COMMENT 'notify',
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  `create_by` int(11) unsigned DEFAULT NULL,
  `update_by` int(11) unsigned DEFAULT NULL,  
  PRIMARY KEY (`id`),
  KEY `idx_wechat_callback_log_create_by` (`create_by`),
  KEY `idx_wechat_callback_log_update_by` (`update_by`),
  KEY `idx_wechat_callback_log_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COMMENT='微信回调';


DROP TABLE IF EXISTS `member_browser`;
CREATE TABLE `member_browser` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `browse_id` bigint DEFAULT NULL COMMENT '查看者会员ID',
  `browsed_id` bigint DEFAULT NULL COMMENT '被查看者会员ID',
  `count` int DEFAULT NULL COMMENT '浏览次数',
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_member_browse_browse_id` (`browse_id`),
  KEY `idx_member_browse_browsed_id` (`browsed_id`),
  KEY `idx_member_browse_created_at` (`created_at`),
  KEY `idx_member_browse_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COMMENT='会员浏览记录';

DROP TABLE IF EXISTS `member_focus`;
CREATE TABLE `member_focus` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `focus_id` bigint DEFAULT NULL COMMENT '关注者会员ID',
  `focused_id` bigint DEFAULT NULL COMMENT '被关注者会员ID',
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_member_focus_focus_id` (`focus_id`),
  KEY `idx_member_focus_focused_id` (`focused_id`),
  KEY `idx_member_focus_created_at` (`created_at`),
  KEY `idx_member_focus_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COMMENT='会员关注记录';


DROP TABLE IF EXISTS `red_line_record`;
CREATE TABLE `red_line_record` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `from_id` bigint DEFAULT NULL COMMENT '发送红线请求会员ID',
  `to_id` bigint DEFAULT NULL COMMENT '接受红线请求ID',
  `greeting` varchar(20) DEFAULT NULL COMMENT '问候语',
  `answer_accept` varchar(20) DEFAULT NULL COMMENT '回答-接受字典',
  `answer_accept_remark` varchar(255) DEFAULT NULL COMMENT '回答-接受内容备注',
  `answer_reject` varchar(20) DEFAULT NULL COMMENT '回答-拒绝字典',
  `status` varchar(20) DEFAULT NULL COMMENT '红线状态', 
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_red_line_record_from_id` (`from_id`),
  KEY `idx_red_line_record_to_id` (`to_id`),
  KEY `idx_red_line_record_created_at` (`created_at`),
  KEY `idx_red_line_record_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COMMENT='红线记录表';


DROP TABLE IF EXISTS `advertisement`;
CREATE TABLE `advertisement` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '主键编码',
  `name` varchar(255) DEFAULT NULL COMMENT '标题',
  `adv_status` int DEFAULT NULL COMMENT '广告状态',
  `img` varchar(255) DEFAULT NULL COMMENT '图片',
  `remark` varchar(255) DEFAULT NULL COMMENT '备注',
  `jump_url` varchar(255) DEFAULT NULL COMMENT '跳转地址',
  `expire_time` timestamp NULL DEFAULT NULL COMMENT '有效日期',
  `sort` int DEFAULT NULL COMMENT '排序',
  `dept_id` bigint DEFAULT NULL COMMENT '城市ID',
  `tag_status` int DEFAULT NULL COMMENT '广告标识状态',
  `location` varchar(20) DEFAULT NULL COMMENT '广告位置如官网banner等',
  `create_by` bigint DEFAULT NULL COMMENT '创建者',
  `update_by` bigint DEFAULT NULL COMMENT '更新者',
  `created_at` datetime(3) DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime(3) DEFAULT NULL COMMENT '最后更新时间',
  `deleted_at` datetime(3) DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`),
  KEY `idx_advertisement_create_by` (`create_by`),
  KEY `idx_advertisement_update_by` (`update_by`),
  KEY `idx_advertisement_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci  COMMENT='广告';


DROP TABLE IF EXISTS `activity`;
CREATE TABLE `activity` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '主键编码',
  `dept_id` bigint DEFAULT NULL COMMENT '城市ID',
  `name` varchar(255) DEFAULT NULL COMMENT '标题',
  `img` varchar(255) DEFAULT NULL COMMENT '封面',
  `act_begin_time` timestamp NULL DEFAULT NULL COMMENT '活动开始日期',
  `act_end_time` timestamp NULL DEFAULT NULL COMMENT '活动截止日期',
  `apply_begin_time` timestamp NULL DEFAULT NULL COMMENT '报名开始日期',
  `apply_end_time` timestamp NULL DEFAULT NULL COMMENT '报名截止日期',
  `male_count` int DEFAULT NULL COMMENT '报名男生数量',
  `female_count` int DEFAULT NULL COMMENT '报名女生数量',
  `list_visible` varchar(4) DEFAULT NULL COMMENT '名单可见',
  `count_visible` varchar(4) DEFAULT NULL COMMENT '报名人数可见',
  `location` varchar(255) DEFAULT NULL COMMENT '地点',
  `geo` varchar(255) DEFAULT NULL COMMENT '地点坐标xy,空格区分',
  `price` decimal(8,2) DEFAULT NULL COMMENT '价格/人次',
  `content` text DEFAULT NULL COMMENT '内容',
  `remark` varchar(255) DEFAULT NULL COMMENT '备注',
  `status` varchar(4) DEFAULT NULL COMMENT '状态',
  `create_by` bigint DEFAULT NULL COMMENT '创建者',
  `update_by` bigint DEFAULT NULL COMMENT '更新者',
  `created_at` datetime(3) DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime(3) DEFAULT NULL COMMENT '最后更新时间',
  `deleted_at` datetime(3) DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`),
  KEY `idx_activity_create_by` (`create_by`),
  KEY `idx_activity_update_by` (`update_by`),
  KEY `idx_activity_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci  COMMENT='活动';


DROP TABLE IF EXISTS `activity_apply`;
CREATE TABLE `activity_apply` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '主键编码',
  `act_id` bigint DEFAULT NULL COMMENT '活动ID',
  `sex` varchar(255) DEFAULT NULL COMMENT '性别',
  `name` varchar(255) DEFAULT NULL COMMENT '姓名',
  `status` varchar(4) DEFAULT NULL COMMENT '状态',
  `create_by` bigint DEFAULT NULL COMMENT '创建者',
  `update_by` bigint DEFAULT NULL COMMENT '更新者',
  `created_at` datetime(3) DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime(3) DEFAULT NULL COMMENT '最后更新时间',
  `deleted_at` datetime(3) DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`),
  KEY `idx_activity_apply_create_by` (`create_by`),
  KEY `idx_activity_apply_update_by` (`update_by`),
  KEY `idx_activity_apply_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci  COMMENT='活动报名';


DROP TABLE IF EXISTS `recommend_position`;
CREATE TABLE `recommend_position` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '主键编码',
  `dept_id` bigint DEFAULT NULL COMMENT '活动ID',
  `member_id` bigint DEFAULT NULL COMMENT '会员ID',
  `remark` varchar(255) DEFAULT NULL COMMENT '备注',
  `status` varchar(4) DEFAULT NULL COMMENT '状态',
  `create_by` bigint DEFAULT NULL COMMENT '创建者',
  `update_by` bigint DEFAULT NULL COMMENT '更新者',
  `created_at` datetime(3) DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime(3) DEFAULT NULL COMMENT '最后更新时间',
  `deleted_at` datetime(3) DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`),
  KEY `idx_recommend_position_create_by` (`create_by`),
  KEY `idx_recommend_position_update_by` (`update_by`),
  KEY `idx_recommend_position_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci  COMMENT='推荐位';


DROP TABLE IF EXISTS `product_red_line`;
CREATE TABLE `product_red_line` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '主键编码',
  `name` varchar(255) DEFAULT NULL COMMENT '商品名称',
  `count` int DEFAULT NULL COMMENT '红线数量',
  `price` decimal(10,2) DEFAULT NULL COMMENT '红线价格',
  `first_price` decimal(10,2) DEFAULT NULL COMMENT '首次购买价格',
  `sort` int DEFAULT NULL COMMENT '排序',
  `remark` varchar(255) DEFAULT NULL COMMENT '备注',
  `status` varchar(4) DEFAULT NULL COMMENT '状态',
  `create_by` bigint DEFAULT NULL COMMENT '创建者',
  `update_by` bigint DEFAULT NULL COMMENT '更新者',
  `created_at` datetime(3) DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime(3) DEFAULT NULL COMMENT '最后更新时间',
  `deleted_at` datetime(3) DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`),
  KEY `idx_product_red_line_create_by` (`create_by`),
  KEY `idx_product_red_line_update_by` (`update_by`),
  KEY `idx_product_red_line_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci  COMMENT='红线商品表';

DROP TABLE IF EXISTS `product_red_line_order`;
CREATE TABLE `product_red_line_order` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '主键编码',
  `member_id` bigint DEFAULT NULL COMMENT '会员ID',
  `product_id` bigint DEFAULT NULL COMMENT '商品ID',
  `order_id` varchar(255)  DEFAULT NULL COMMENT '订单ID',
  `count` int DEFAULT NULL COMMENT '购买红线数量',
  `price` decimal(10,2) DEFAULT NULL COMMENT '红线支付价格',
  `status` int DEFAULT NULL COMMENT '订单状态',
  `created_at` datetime(3) DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime(3) DEFAULT NULL COMMENT '最后更新时间',
  `deleted_at` datetime(3) DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`),
  KEY `idx_product_red_line_created_at` (`created_at`),
  KEY `idx_product_red_line_updated_at` (`updated_at`),
  KEY `idx_product_red_line_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci  COMMENT='红线商品表';

DROP TABLE IF EXISTS `report`;
CREATE TABLE `report` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '主键编码',
  `report_id` bigint DEFAULT NULL COMMENT '举报人ID',
  `reported_id` bigint DEFAULT NULL COMMENT '被举报人ID',
  `report_type` varchar(20) DEFAULT NULL COMMENT '举报类型',
  `content` varchar(255) DEFAULT NULL COMMENT '举报内容',
  `source` int DEFAULT NULL COMMENT '来源 小程序/h5等',
  `ip` varchar(255) DEFAULT NULL COMMENT '举报人IP',
  `status` int DEFAULT NULL COMMENT '处理状态，1已处理',
  `user_id` bigint DEFAULT NULL COMMENT '处理人ID,对应后台sys_user表',
  `created_at` datetime(3) DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime(3) DEFAULT NULL COMMENT '最后更新时间',
  `deleted_at` datetime(3) DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`),
  KEY `idx_report_created_at` (`created_at`),
  KEY `idx_report_updated_at` (`updated_at`),
  KEY `idx_report_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci  COMMENT='举报表';