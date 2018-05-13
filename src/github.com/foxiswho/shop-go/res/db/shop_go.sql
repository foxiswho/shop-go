/*
 Navicat Premium Data Transfer

 Source Server         : localhost
 Source Server Type    : MySQL
 Source Server Version : 50721
 Source Host           : localhost:3306
 Source Schema         : fox

 Target Server Type    : MySQL
 Target Server Version : 50721
 File Encoding         : 65001

 Date: 13/05/2018 16:09:16
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for admin
-- ----------------------------
DROP TABLE IF EXISTS `admin`;
CREATE TABLE `admin` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `username` char(30) DEFAULT NULL COMMENT '用户名',
  `password` char(32) DEFAULT NULL COMMENT '密码',
  `mail` varchar(80) DEFAULT NULL COMMENT '邮箱',
  `salt` varchar(10) DEFAULT NULL COMMENT '干扰码',
  `gmt_create` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `gmt_modified` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `ip` char(15) DEFAULT NULL COMMENT '添加IP',
  `job_no` varchar(15) DEFAULT NULL COMMENT '工号',
  `nick_name` varchar(50) DEFAULT NULL COMMENT '昵称',
  `true_name` varchar(50) DEFAULT NULL COMMENT '真实姓名',
  `qq` varchar(50) DEFAULT NULL COMMENT 'qq',
  `phone` varchar(50) DEFAULT NULL COMMENT '电话',
  `mobile` varchar(20) DEFAULT NULL COMMENT '手机',
  `name` varchar(255) DEFAULT NULL COMMENT '显示名称',
  `is_del` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '删除0否1是',
  `department_id` int(11) NOT NULL DEFAULT '0' COMMENT '部门id',
  `team_id` int(11) NOT NULL COMMENT '团队ID',
  `master_id` int(11) NOT NULL COMMENT '师傅id',
  `leader_id` int(11) NOT NULL COMMENT '领导id',
  `post_id` int(11) NOT NULL COMMENT '职务id',
  `role_id` int(11) NOT NULL COMMENT '角色id(主)',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `username` (`username`),
  KEY `is_del` (`is_del`),
  KEY `role_id` (`role_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ROW_FORMAT=COMPACT COMMENT='管理员表';

-- ----------------------------
-- Table structure for admin_menu
-- ----------------------------
DROP TABLE IF EXISTS `admin_menu`;
CREATE TABLE `admin_menu` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `name` char(100) DEFAULT NULL COMMENT '名称',
  `parent_id` int(11) NOT NULL DEFAULT '0' COMMENT '上级菜单',
  `s` char(60) DEFAULT NULL COMMENT '模块/控制器/动作',
  `data` char(100) DEFAULT NULL COMMENT '其他参数',
  `sort` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '排序',
  `remark` varchar(255) DEFAULT NULL COMMENT '备注',
  `type` char(32) NOT NULL DEFAULT 'url' COMMENT '类别url菜单function独立功能user用户独有',
  `level` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '级别',
  `level1_id` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '1级栏目ID',
  `md5` char(32) DEFAULT NULL COMMENT 's的md5值',
  `is_show` tinyint(1) NOT NULL DEFAULT '1' COMMENT '显示隐藏;1显示;0隐藏',
  `is_unique` tinyint(1) NOT NULL DEFAULT '0' COMMENT '用户独有此功能1是0否',
  PRIMARY KEY (`id`),
  KEY `sort` (`sort`),
  KEY `parent_id` (`parent_id`),
  KEY `s` (`s`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ROW_FORMAT=COMPACT COMMENT='菜单';

-- ----------------------------
-- Table structure for admin_role_access
-- ----------------------------
DROP TABLE IF EXISTS `admin_role_access`;
CREATE TABLE `admin_role_access` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `aid` int(11) NOT NULL DEFAULT '0' COMMENT '管理员ID',
  `role_id` int(11) NOT NULL DEFAULT '0' COMMENT '角色ID',
  `is_default` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否默认',
  PRIMARY KEY (`id`),
  UNIQUE KEY `aid_role_id` (`aid`,`role_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ROW_FORMAT=COMPACT COMMENT='管理员与角色的关系、一个管理员可以有多个角色';

-- ----------------------------
-- Table structure for admin_role_priv
-- ----------------------------
DROP TABLE IF EXISTS `admin_role_priv`;
CREATE TABLE `admin_role_priv` (
  `id` int(10) NOT NULL AUTO_INCREMENT,
  `role_id` smallint(3) unsigned NOT NULL DEFAULT '0' COMMENT '角色ID',
  `s` char(100) DEFAULT NULL COMMENT '模块/控制器/动作',
  `data` char(50) DEFAULT NULL COMMENT '其他参数',
  `aid` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '管理员ID',
  `menu_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '菜单ID',
  `type` char(32) NOT NULL DEFAULT 'url' COMMENT '类别url菜单function独立功能user用户独有',
  PRIMARY KEY (`id`),
  KEY `role_id` (`role_id`),
  KEY `role_id_2` (`role_id`,`s`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ROW_FORMAT=COMPACT COMMENT='角色权限表';

-- ----------------------------
-- Table structure for admin_status
-- ----------------------------
DROP TABLE IF EXISTS `admin_status`;
CREATE TABLE `admin_status` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `login_time` timestamp NULL DEFAULT NULL COMMENT '登录时间',
  `login_ip` char(20) DEFAULT NULL COMMENT 'IP',
  `login` int(11) NOT NULL DEFAULT '0' COMMENT '登录次数',
  `aid_add` int(11) NOT NULL DEFAULT '0' COMMENT '添加人',
  `aid_update` int(11) NOT NULL DEFAULT '0' COMMENT '更新人',
  `gmt_modified` timestamp NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `remark` varchar(255) DEFAULT NULL COMMENT '备注',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ROW_FORMAT=COMPACT COMMENT='状态';

-- ----------------------------
-- Table structure for app
-- ----------------------------
DROP TABLE IF EXISTS `app`;
CREATE TABLE `app` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `type_id` int(11) NOT NULL DEFAULT '0' COMMENT 'app_id,来源type表',
  `name` varchar(100) DEFAULT NULL COMMENT '名称',
  `mark` char(32) DEFAULT NULL COMMENT '标志',
  `setting` varchar(5000) DEFAULT NULL COMMENT '扩展参数',
  `remark` varchar(255) DEFAULT NULL COMMENT '备注',
  `is_del` int(11) NOT NULL DEFAULT '0' COMMENT '是否删除0否1是',
  `gmt_create` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '添加时间',
  `gmt_modified` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `type_id` (`type_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='应用';

-- ----------------------------
-- Table structure for area
-- ----------------------------
DROP TABLE IF EXISTS `area`;
CREATE TABLE `area` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` char(50) DEFAULT '' COMMENT '名称',
  `name_en` varchar(100) DEFAULT '' COMMENT '英文名称',
  `parent_id` int(11) DEFAULT '0' COMMENT '上级栏目ID',
  `type` tinyint(4) DEFAULT '0' COMMENT '类别;0默认;',
  `name_traditional` varchar(50) DEFAULT '' COMMENT '繁体名称',
  `sort` int(11) DEFAULT '0' COMMENT '排序',
  PRIMARY KEY (`id`),
  KEY `parent_id` (`parent_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ROW_FORMAT=COMPACT COMMENT='地区表';

-- ----------------------------
-- Table structure for area_ext
-- ----------------------------
DROP TABLE IF EXISTS `area_ext`;
CREATE TABLE `area_ext` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `area_id` int(11) DEFAULT '0' COMMENT 'ID',
  `name` char(50) DEFAULT '' COMMENT '名称',
  `name_en` varchar(100) DEFAULT '' COMMENT '英文名称',
  `parent_id` int(11) DEFAULT '0' COMMENT '上级栏目ID',
  `type` tinyint(4) DEFAULT '0' COMMENT '类别;0默认;1又名;2;3属于;11已合并到;12已更名为',
  `name_traditional` varchar(50) DEFAULT '' COMMENT '繁体名称',
  `sort` int(11) DEFAULT '0' COMMENT '排序',
  `type_name` varchar(50) DEFAULT '' COMMENT '类别名称',
  `other_name` varchar(50) DEFAULT '' COMMENT '根据类别名称填写',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `id` (`area_id`,`parent_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ROW_FORMAT=COMPACT COMMENT='地区扩展表';

-- ----------------------------
-- Table structure for attachment
-- ----------------------------
DROP TABLE IF EXISTS `attachment`;
CREATE TABLE `attachment` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '附件ID',
  `module` char(32) DEFAULT NULL COMMENT '模块',
  `mark` char(60) DEFAULT NULL COMMENT '标记标志',
  `type_id` int(5) unsigned NOT NULL DEFAULT '0' COMMENT '类别ID',
  `name` char(50) DEFAULT NULL COMMENT '保存的文件名称',
  `name_original` varchar(255) DEFAULT NULL COMMENT '原文件名',
  `path` char(200) DEFAULT NULL COMMENT '文件路径',
  `size` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '文件大小',
  `ext` char(10) DEFAULT NULL COMMENT '文件后缀',
  `is_image` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '是否图片1是0否',
  `is_thumb` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '是否缩略图1是0否',
  `downloads` int(8) unsigned NOT NULL DEFAULT '0' COMMENT '下载次数',
  `gmt_create` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '添加时间上传时间',
  `ip` char(15) DEFAULT NULL COMMENT '上传IP',
  `status` tinyint(2) unsigned NOT NULL DEFAULT '0' COMMENT '状态99正常;',
  `md5` char(32) DEFAULT NULL COMMENT 'md5',
  `sha1` char(40) DEFAULT NULL COMMENT 'sha1',
  `from_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '所属ID',
  `aid` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '后台管理员ID',
  `uid` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '前台用户ID',
  `is_show` tinyint(1) unsigned NOT NULL DEFAULT '1' COMMENT '是否显示1是0否',
  `http` varchar(100) DEFAULT NULL COMMENT '图片http地址',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `md5` (`md5`),
  KEY `module` (`module`),
  KEY `mark` (`mark`),
  KEY `id` (`from_id`),
  KEY `status` (`status`),
  KEY `aid` (`aid`),
  KEY `uid` (`uid`),
  KEY `is_show` (`is_show`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ROW_FORMAT=COMPACT COMMENT='附件表';

-- ----------------------------
-- Table structure for connect
-- ----------------------------
DROP TABLE IF EXISTS `connect`;
CREATE TABLE `connect` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `type_id` int(11) NOT NULL DEFAULT '0' COMMENT '类别id',
  `uid` int(11) NOT NULL DEFAULT '0' COMMENT '用户id',
  `open_id` char(80) DEFAULT NULL COMMENT '对应唯一开放id',
  `token` varchar(80) DEFAULT NULL COMMENT '开放密钥',
  `type` int(11) NOT NULL DEFAULT '1' COMMENT '登录类型1腾讯QQ2新浪微博',
  `type_login` int(11) NOT NULL DEFAULT '0' COMMENT '登录模块;302前台还是后台301',
  `gmt_create` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `extend` varchar(5000) DEFAULT '' COMMENT '扩展参数',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `openid` (`open_id`),
  KEY `uid` (`uid`) USING BTREE,
  KEY `type_id` (`type_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ROW_FORMAT=COMPACT COMMENT='快捷登陆/qq';

-- ----------------------------
-- Table structure for log
-- ----------------------------
DROP TABLE IF EXISTS `log`;
CREATE TABLE `log` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `from_id` int(11) NOT NULL DEFAULT '0' COMMENT 'id',
  `aid` int(11) NOT NULL DEFAULT '0' COMMENT '管理员ID',
  `uid` int(11) NOT NULL DEFAULT '0' COMMENT '用户id',
  `gmt_create` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `mark` char(32) DEFAULT NULL COMMENT '标志自定义标志',
  `data` text COMMENT '其他内容',
  `no` char(50) DEFAULT NULL COMMENT '单号',
  `type_login` int(11) NOT NULL DEFAULT '0' COMMENT '登录方式;302前台还是后台301',
  `type_client` int(11) NOT NULL DEFAULT '0' COMMENT '登录客户端类别;321电脑;322安卓;323IOS;324手机网页;325其他',
  `ip` char(20) DEFAULT NULL COMMENT 'IP',
  `msg` varchar(255) DEFAULT NULL COMMENT '自定义说明',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `type_login` (`type_login`),
  KEY `type_client` (`type_client`),
  KEY `uid` (`uid`),
  KEY `aid` (`aid`),
  KEY `id` (`from_id`),
  KEY `no` (`no`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ROW_FORMAT=COMPACT COMMENT='日志表';

-- ----------------------------
-- Table structure for news
-- ----------------------------
DROP TABLE IF EXISTS `news`;
CREATE TABLE `news` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `aid` int(11) NOT NULL DEFAULT '0' COMMENT '管理员AID',
  `is_del` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否删除1是0否',
  `is_open` tinyint(1) NOT NULL DEFAULT '1' COMMENT '启用1是0否',
  `status` int(11) NOT NULL DEFAULT '0' COMMENT '状态',
  `gmt_system` timestamp NULL DEFAULT NULL COMMENT '创建时间,系统时间不可修改',
  `gmt_modified` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `gmt_create` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '添加时间,可修改',
  `title` varchar(255) DEFAULT NULL COMMENT '标题',
  `author` varchar(255) DEFAULT NULL COMMENT '作者',
  `url` varchar(255) DEFAULT NULL COMMENT '网址',
  `url_source` varchar(255) DEFAULT NULL COMMENT '来源地址(转载)',
  `url_rewrite` char(150) DEFAULT NULL COMMENT '自定义伪静态Url',
  `description` varchar(255) DEFAULT NULL COMMENT '摘要',
  `content` text COMMENT '内容',
  `type` int(11) NOT NULL DEFAULT '0' COMMENT '类型0文章10001博客栏目',
  `module_id` int(10) NOT NULL DEFAULT '0' COMMENT '模块10019技术10018生活',
  `source_id` int(11) NOT NULL DEFAULT '0' COMMENT '来源:后台，接口，其他',
  `type_id` int(11) NOT NULL DEFAULT '0' COMMENT '类别ID，原创，转载，翻译',
  `cat_id` int(11) NOT NULL DEFAULT '0' COMMENT '分类ID，栏目',
  `tag` varchar(255) DEFAULT NULL COMMENT '标签',
  `thumb` varchar(255) DEFAULT NULL COMMENT '缩略图',
  `is_relevant` tinyint(1) NOT NULL DEFAULT '0' COMMENT '相关文章1是0否',
  `is_jump` tinyint(1) NOT NULL DEFAULT '0' COMMENT '跳转1是0否',
  `is_comment` tinyint(1) NOT NULL DEFAULT '1' COMMENT '允许评论1是0否',
  `is_read` int(11) NOT NULL DEFAULT '10014' COMMENT '是否阅读10014未看10015在看10016已看',
  `sort` int(11) NOT NULL DEFAULT '0' COMMENT '排序',
  `remark` varchar(255) DEFAULT NULL COMMENT '备注',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `is_del` (`is_del`,`is_open`,`status`,`type_id`,`cat_id`,`sort`),
  KEY `url_rewrite` (`url_rewrite`),
  KEY `type` (`type`),
  KEY `module_id` (`module_id`) USING BTREE,
  KEY `source_id` (`source_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ROW_FORMAT=COMPACT COMMENT='博客内容';

-- ----------------------------
-- Table structure for news_statistics
-- ----------------------------
DROP TABLE IF EXISTS `news_statistics`;
CREATE TABLE `news_statistics` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `news_id` int(11) NOT NULL DEFAULT '0' COMMENT '文章ID',
  `comment` int(11) NOT NULL DEFAULT '0' COMMENT '评论人数',
  `read` int(11) NOT NULL DEFAULT '0' COMMENT '阅读人数',
  `seo_title` varchar(255) DEFAULT NULL COMMENT 'SEO标题',
  `seo_description` varchar(255) DEFAULT NULL COMMENT 'SEO摘要',
  `seo_keyword` varchar(255) DEFAULT NULL COMMENT 'SEO关键词',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `blog_id` (`news_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ROW_FORMAT=COMPACT COMMENT='博客统计';

-- ----------------------------
-- Table structure for news_sync_mapping
-- ----------------------------
DROP TABLE IF EXISTS `news_sync_mapping`;
CREATE TABLE `news_sync_mapping` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `news_id` int(11) NOT NULL DEFAULT '0' COMMENT '本站blog的id',
  `type_id` int(11) NOT NULL DEFAULT '0' COMMENT '类别id',
  `to_id` varchar(64) DEFAULT NULL COMMENT 'csdn的id',
  `gmt_modified` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '最后一次更新时间',
  `gmt_create` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '插入时间',
  `mark` char(32) DEFAULT NULL COMMENT '标志',
  `is_sync` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否同步过',
  `extend` varchar(5000) DEFAULT NULL COMMENT '扩展参数',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='本站blog_id 与其他同步站点的id关系';

-- ----------------------------
-- Table structure for news_sync_queue
-- ----------------------------
DROP TABLE IF EXISTS `news_sync_queue`;
CREATE TABLE `news_sync_queue` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `news_id` int(11) NOT NULL DEFAULT '0' COMMENT '本站博客id',
  `type_id` int(11) NOT NULL DEFAULT '0' COMMENT '类型',
  `status` tinyint(3) NOT NULL DEFAULT '0' COMMENT '状态：0:待运行 10:失败 99:成功',
  `gmt_modified` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '最后一次更新时间',
  `gmt_create` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '插入时间',
  `msg` varchar(255) DEFAULT NULL COMMENT '内容',
  `map_id` int(11) NOT NULL DEFAULT '0' COMMENT '同步ID',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='博客同步队列';

-- ----------------------------
-- Table structure for news_tag
-- ----------------------------
DROP TABLE IF EXISTS `news_tag`;
CREATE TABLE `news_tag` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` char(100) DEFAULT NULL COMMENT '名称',
  `gmt_create` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '添加时间',
  `aid` int(11) NOT NULL DEFAULT '0' COMMENT '管理员ID',
  `news_id` int(11) NOT NULL DEFAULT '0' COMMENT '文章ID',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ROW_FORMAT=COMPACT COMMENT='博客标签';

-- ----------------------------
-- Table structure for session
-- ----------------------------
DROP TABLE IF EXISTS `session`;
CREATE TABLE `session` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `uid` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '用户UID',
  `ip` char(15) DEFAULT NULL COMMENT 'IP',
  `error_count` tinyint(1) NOT NULL DEFAULT '0' COMMENT '密码输入错误次数',
  `app_id` int(11) NOT NULL DEFAULT '0' COMMENT '登录应用',
  `md5` char(32) DEFAULT NULL COMMENT 'md5',
  `type_login` int(11) NOT NULL DEFAULT '0' COMMENT '登录方式;302前台还是后台301',
  `type_client` int(11) NOT NULL DEFAULT '0' COMMENT '登录客户端类别;321电脑;322安卓;323IOS;324手机网页;325其他',
  `gmt_create` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '登录时间',
  `gmt_modified` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `uid` (`uid`,`type_login`,`type_client`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ROW_FORMAT=COMPACT COMMENT='SESSION';

-- ----------------------------
-- Table structure for tag
-- ----------------------------
DROP TABLE IF EXISTS `tag`;
CREATE TABLE `tag` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` char(50) DEFAULT NULL COMMENT '名称',
  `gmt_create` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '添加时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ROW_FORMAT=COMPACT COMMENT='标签';

-- ----------------------------
-- Table structure for template
-- ----------------------------
DROP TABLE IF EXISTS `template`;
CREATE TABLE `template` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '模板ID',
  `name` varchar(80) DEFAULT NULL COMMENT '模板名称(中文)',
  `mark` varchar(80) DEFAULT NULL COMMENT '模板名称标志(英文)（调用时使用）',
  `title` varchar(255) DEFAULT NULL COMMENT '邮件标题',
  `type` tinyint(1) NOT NULL DEFAULT '0' COMMENT '模板类型1短信模板2邮箱模板',
  `use` tinyint(2) unsigned NOT NULL DEFAULT '0' COMMENT '用途',
  `content` text,
  `remark` varchar(1024) DEFAULT NULL COMMENT '备注',
  `gmt_create` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `gmt_modified` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `code_num` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '验证码位数',
  `aid` int(11) NOT NULL DEFAULT '0' COMMENT '添加人',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ROW_FORMAT=COMPACT COMMENT='会员表';

-- ----------------------------
-- Table structure for type
-- ----------------------------
DROP TABLE IF EXISTS `type`;
CREATE TABLE `type` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` char(100) DEFAULT NULL COMMENT '名称',
  `name_en` char(100) DEFAULT NULL COMMENT '名称',
  `code` char(32) DEFAULT NULL COMMENT '代码',
  `mark` char(32) DEFAULT NULL COMMENT '标志',
  `type_id` int(11) NOT NULL DEFAULT '0' COMMENT '所属类别ID',
  `parent_id` int(11) NOT NULL DEFAULT '0' COMMENT '上级ID、属于/上级ID',
  `value` int(10) NOT NULL DEFAULT '0' COMMENT '值',
  `content` varchar(255) DEFAULT NULL COMMENT '内容',
  `is_del` int(11) NOT NULL DEFAULT '0' COMMENT '是否删除0否1是',
  `sort` int(11) NOT NULL DEFAULT '0' COMMENT '排序',
  `remark` varchar(255) DEFAULT NULL COMMENT '备注',
  `gmt_create` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '添加时间',
  `aid` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '添加人',
  `module` char(50) DEFAULT NULL COMMENT '模块',
  `setting` varchar(255) DEFAULT NULL COMMENT '扩展参数',
  `is_default` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否默认',
  `is_child` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否有子类1是0否',
  `is_system` tinyint(1) NOT NULL DEFAULT '0' COMMENT '系统参数禁止修改',
  `is_show` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否显示在配置页面上',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `is_del` (`is_del`),
  KEY `parent_id` (`parent_id`),
  KEY `sort` (`sort`),
  KEY `mark` (`mark`),
  KEY `type_id` (`type_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ROW_FORMAT=COMPACT COMMENT='类别';

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `mobile` char(11) CHARACTER SET utf8 COLLATE utf8_bin DEFAULT NULL,
  `username` char(30) DEFAULT NULL COMMENT '用户名',
  `mail` char(32) DEFAULT NULL COMMENT '邮箱',
  `password` char(32) DEFAULT NULL COMMENT '密码',
  `salt` char(6) DEFAULT NULL COMMENT '干扰码',
  `reg_ip` char(15) DEFAULT NULL COMMENT '注册IP',
  `reg_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '注册时间',
  `is_del` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '状态0正常1删除',
  `group_id` int(11) unsigned NOT NULL DEFAULT '410' COMMENT '用户组ID',
  `true_name` varchar(32) DEFAULT NULL COMMENT '真实姓名',
  `name` varchar(100) DEFAULT NULL COMMENT '店铺名称',
  `gmt_create` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '添加时间',
  `gmt_modified` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `mobile` (`mobile`),
  KEY `is_del` (`is_del`),
  KEY `username` (`username`),
  KEY `email` (`mail`),
  KEY `group_id` (`group_id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 ROW_FORMAT=COMPACT COMMENT='用户表';

-- ----------------------------
-- Records of user
-- ----------------------------

-- ----------------------------
-- Table structure for user_group_ext
-- ----------------------------
DROP TABLE IF EXISTS `user_group_ext`;
CREATE TABLE `user_group_ext` (
  `group_id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '用户ID',
  PRIMARY KEY (`group_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ROW_FORMAT=COMPACT COMMENT='会员用户组扩展';

-- ----------------------------
-- Table structure for user_profile
-- ----------------------------
DROP TABLE IF EXISTS `user_profile`;
CREATE TABLE `user_profile` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `sex` tinyint(1) unsigned NOT NULL DEFAULT '1' COMMENT '性别1男2女3中性0保密',
  `job` varchar(50) DEFAULT NULL COMMENT '担任职务',
  `qq` varchar(20) DEFAULT NULL,
  `phone` varchar(20) DEFAULT NULL COMMENT '电话',
  `county` int(11) unsigned NOT NULL DEFAULT '1' COMMENT '国家',
  `province` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '省',
  `city` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '市',
  `district` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '区',
  `address` varchar(255) DEFAULT NULL COMMENT '地址',
  `wechat` varchar(20) DEFAULT NULL COMMENT '微信',
  `remark_admin` text COMMENT '客服备注',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ROW_FORMAT=COMPACT COMMENT='用户其他介绍';

-- ----------------------------
-- Table structure for user_status
-- ----------------------------
DROP TABLE IF EXISTS `user_status`;
CREATE TABLE `user_status` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `reg_ip` char(15) DEFAULT NULL COMMENT '注册IP',
  `reg_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '注册时间',
  `reg_type` int(11) NOT NULL DEFAULT '0' COMMENT '注册方式',
  `reg_app_id` int(11) NOT NULL DEFAULT '1' COMMENT '注册来源',
  `last_login_ip` char(15) DEFAULT NULL COMMENT '最后登录IP',
  `last_login_time` timestamp NULL DEFAULT NULL COMMENT '最后登录时间',
  `last_login_app_id` int(11) NOT NULL DEFAULT '0' COMMENT '最后登录app_id',
  `login` smallint(5) NOT NULL DEFAULT '0' COMMENT '登录次数',
  `is_mobile` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '手机号是否已验证1已验证0未验证',
  `is_email` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '邮箱是否已验证1已验证0未验证',
  `aid_add` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '客服AID',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ROW_FORMAT=COMPACT COMMENT='用户状态';

SET FOREIGN_KEY_CHECKS = 1;
