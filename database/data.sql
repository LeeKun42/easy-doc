/*
 Navicat MySQL Data Transfer

 Source Server         : local
 Source Server Type    : MySQL
 Source Server Version : 80033
 Source Host           : 127.0.0.1:3309
 Source Schema         : doc

 Target Server Type    : MySQL
 Target Server Version : 80033
 File Encoding         : 65001

 Date: 07/09/2023 15:31:22
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for project_apis
-- ----------------------------
DROP TABLE IF EXISTS `project_apis`;
CREATE TABLE `project_apis`  (
                                 `id` int(0) UNSIGNED NOT NULL AUTO_INCREMENT,
                                 `project_id` int(0) UNSIGNED NOT NULL COMMENT '所属项目id',
                                 `directory_id` int(0) UNSIGNED NOT NULL COMMENT '所属目录id',
                                 `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '接口名称',
                                 `path` varchar(512) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '接口地址',
                                 `method` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '请求类型 GET POST ....',
                                 `request_headers` text CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL COMMENT '请求头',
                                 `request_query` text CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL COMMENT 'query参数',
                                 `request_path` text CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL COMMENT 'path参数',
                                 `request_body` text CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL COMMENT '请求体',
                                 `response_body` text CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL COMMENT '响应参数',
                                 `desc` text CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL COMMENT '接口说明',
                                 `seq` int(0) UNSIGNED NOT NULL DEFAULT 10 COMMENT '排序号从小到大排序',
                                 `created_at` timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP(0),
                                 `updated_at` timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP(0) ON UPDATE CURRENT_TIMESTAMP(0),
                                 PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for project_directory
-- ----------------------------
DROP TABLE IF EXISTS `project_directory`;
CREATE TABLE `project_directory`  (
                                      `id` int(0) UNSIGNED NOT NULL AUTO_INCREMENT,
                                      `parent_id` int(0) UNSIGNED NOT NULL DEFAULT 0 COMMENT '父级目录id',
                                      `project_id` int(0) UNSIGNED NOT NULL COMMENT '所属项目id',
                                      `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL COMMENT '目录名称',
                                      `desc` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL COMMENT '目录说明',
                                      `seq` int(0) UNSIGNED NOT NULL COMMENT '排序号从小到大排序',
                                      `created_at` timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP(0),
                                      `updated_at` timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP(0) ON UPDATE CURRENT_TIMESTAMP(0),
                                      PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci COMMENT = '项目目录表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for project_users
-- ----------------------------
DROP TABLE IF EXISTS `project_users`;
CREATE TABLE `project_users`  (
                                  `project_id` int(0) UNSIGNED NOT NULL,
                                  `user_id` int(0) UNSIGNED NOT NULL,
                                  PRIMARY KEY (`project_id`, `user_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci COMMENT = '项目成员表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for projects
-- ----------------------------
DROP TABLE IF EXISTS `projects`;
CREATE TABLE `projects`  (
                             `id` int(0) UNSIGNED NOT NULL AUTO_INCREMENT,
                             `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '项目名称',
                             `owner_user_id` int(0) UNSIGNED NOT NULL COMMENT '项目所有者/创建人',
                             `created_at` timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP(0),
                             `updated_at` timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP(0) ON UPDATE CURRENT_TIMESTAMP(0),
                             PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci COMMENT = '项目信息表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for users
-- ----------------------------
DROP TABLE IF EXISTS `users`;
CREATE TABLE `users`  (
                          `id` int(0) UNSIGNED NOT NULL AUTO_INCREMENT,
                          `account` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '账号',
                          `passwd` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '密码',
                          `nick_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '姓名/昵称',
                          `status` tinyint(0) UNSIGNED NOT NULL COMMENT '状态：1：正常  0：禁用',
                          `created_at` timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP(0),
                          `updated_at` timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP(0) ON UPDATE CURRENT_TIMESTAMP(0),
                          PRIMARY KEY (`id`) USING BTREE,
                          UNIQUE INDEX `account`(`account`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

SET FOREIGN_KEY_CHECKS = 1;
