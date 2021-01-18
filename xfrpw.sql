/*
 Navicat Premium Data Transfer

 Source Server         : localhost
 Source Server Type    : MySQL
 Source Server Version : 50527
 Source Host           : localhost:3306
 Source Schema         : xfrpw

 Target Server Type    : MySQL
 Target Server Version : 50527
 File Encoding         : 65001

 Date: 26/12/2020 23:56:57
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for ccit_admin
-- ----------------------------
DROP TABLE IF EXISTS `ccit_admin`;
CREATE TABLE `ccit_admin`  (
  `admin_id` varchar(20) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '1' COMMENT '用户id',
  `admin_account` varchar(30) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '登录账户',
  `admin_password` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '密码',
  `register_time` datetime NULL DEFAULT NULL COMMENT '注册日期',
  `admin_permission` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '用户权限',
  PRIMARY KEY (`admin_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Compact;

-- ----------------------------
-- Records of ccit_admin
-- ----------------------------
INSERT INTO `ccit_admin` VALUES ('1', 'admin', '1', '2019-05-18 17:43:54', NULL);

-- ----------------------------
-- Table structure for ccit_classify
-- ----------------------------
DROP TABLE IF EXISTS `ccit_classify`;
CREATE TABLE `ccit_classify`  (
  `classify_id` varchar(18) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '方向分类id',
  `add_date` date NOT NULL COMMENT '添加时间',
  `classify_name` varchar(30) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '分类名称',
  `indexes` varchar(20) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '索引标签',
  `num` int(5) NOT NULL DEFAULT 0 COMMENT '旗下的课程数量',
  `classify_path` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '这个专业方向内课程路径',
  PRIMARY KEY (`classify_id`) USING BTREE,
  INDEX `classify_id`(`classify_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Compact;

-- ----------------------------
-- Records of ccit_classify
-- ----------------------------
INSERT INTO `ccit_classify` VALUES ('156077761260371130', '2019-06-17', 'Java', 'java', 4, '156077761260371130');
INSERT INTO `ccit_classify` VALUES ('156077800709779450', '2019-06-17', 'PHP', 'php', 0, '156077800709779450');
INSERT INTO `ccit_classify` VALUES ('156077811375324030', '2019-06-17', 'SpringBoot', 'springboot', 0, '156077811375324030');
INSERT INTO `ccit_classify` VALUES ('156077813082722890', '2019-06-17', 'Python', 'py', 0, '156077813082722890');
INSERT INTO `ccit_classify` VALUES ('156078318122944900', '2019-06-17', 'HTML', 'html', 2, '156078318122944900');
INSERT INTO `ccit_classify` VALUES ('156078321216471040', '2019-06-17', '大数据', 'dsj', 2, '156078321216471040');
INSERT INTO `ccit_classify` VALUES ('156078324149150720', '2019-06-17', 'IOS', 'ios', 1, '156078324149150720');
INSERT INTO `ccit_classify` VALUES ('156085224007781910', '2019-06-18', 'MySQL', 'mysql', 1, '156085224007781910');
INSERT INTO `ccit_classify` VALUES ('156085224007781911', '2019-07-17', '人工智能', 'rgzn', 1, '156085224007781911');
INSERT INTO `ccit_classify` VALUES ('156085224007781912', '2019-06-24', 'Android', 'android', 13, '156085224007781912');
INSERT INTO `ccit_classify` VALUES ('156085224007781913', '2019-09-12', 'CSS', 'css', 2, '156085224007781913');
INSERT INTO `ccit_classify` VALUES ('647421722570850950', '2019-11-07', 'Golang', 'golang', 3, '');
INSERT INTO `ccit_classify` VALUES ('717692328245959885', '2019-11-05', '测试1', 'cs1', 5, '717692328245959885');

-- ----------------------------
-- Table structure for ccit_course
-- ----------------------------
DROP TABLE IF EXISTS `ccit_course`;
CREATE TABLE `ccit_course`  (
  `courset_id` varchar(18) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '课程id',
  `classify_id` varchar(18) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '分类id',
  `course_name` varchar(30) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '课程名称',
  `course_info` text CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '课程信息',
  `cover_img` text CHARACTER SET utf8 COLLATE utf8_general_ci NULL COMMENT '封面',
  `putaway_time` date NULL DEFAULT NULL COMMENT '上架时间',
  `label` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '标签',
  `is_permit` int(1) NULL DEFAULT NULL COMMENT '是否上架',
  `course_path` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '课程路径',
  PRIMARY KEY (`courset_id`) USING BTREE,
  INDEX `classify_id`(`classify_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Compact;

-- ----------------------------
-- Records of ccit_course
-- ----------------------------
INSERT INTO `ccit_course` VALUES ('676651078763931951', '717692328245959885', '测试1', '测试1', 'static/upload/2019-11-23/sa264sd454112dvbss.jpg', '2019-11-23', '', 0, 'static/upvideo/717692328245959885/676651078763931951');

-- ----------------------------
-- Table structure for ccit_knobble
-- ----------------------------
DROP TABLE IF EXISTS `ccit_knobble`;
CREATE TABLE `ccit_knobble`  (
  `knobble_id` varchar(20) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '小节id',
  `courset_id` varchar(18) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '224680481405302848' COMMENT '课程id',
  `section_id` varchar(20) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '章节id',
  `knobble_name` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '小节标题',
  `video_address` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '小节视频地址',
  `section_part` varchar(5) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '章节标号',
  PRIMARY KEY (`knobble_id`) USING BTREE,
  INDEX `section_id`(`section_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Compact;

-- ----------------------------
-- Table structure for ccit_section
-- ----------------------------
DROP TABLE IF EXISTS `ccit_section`;
CREATE TABLE `ccit_section`  (
  `courset_id` varchar(18) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '课程id',
  `section_id` varchar(18) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '章节id',
  `section_name` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '章节名称',
  `section_info` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '章节介绍',
  `section_part` varchar(5) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '章节标号',
  `putaway_time` date NULL DEFAULT NULL COMMENT '添加时间',
  PRIMARY KEY (`section_id`) USING BTREE,
  INDEX `courset_id`(`courset_id`) USING BTREE,
  INDEX `section_id`(`section_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Compact;

-- ----------------------------
-- Table structure for ccit_user
-- ----------------------------
DROP TABLE IF EXISTS `ccit_user`;
CREATE TABLE `ccit_user`  (
  `student_number` varchar(11) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '学号',
  `student_password` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '密码',
  `student_name` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '姓名',
  `student_phone` varchar(11) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '手机',
  `student_activity` int(1) NULL DEFAULT NULL COMMENT '是否激活',
  PRIMARY KEY (`student_number`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Compact;

-- ----------------------------
-- Records of ccit_user
-- ----------------------------
INSERT INTO `ccit_user` VALUES ('16091231043', '123456', '张帅威', '13151289178', 1);
INSERT INTO `ccit_user` VALUES ('16091231044', '123457', '张帅威', '13151289179', 1);
INSERT INTO `ccit_user` VALUES ('16091231045', '123458', '张帅威', '13151289180', 1);
INSERT INTO `ccit_user` VALUES ('16091231046', '123459', '张帅威', '13151289181', 1);
INSERT INTO `ccit_user` VALUES ('16091231047', '123460', '张帅威', '13151289182', 1);
INSERT INTO `ccit_user` VALUES ('16091231048', '123461', '张帅威', '13151289183', 1);
INSERT INTO `ccit_user` VALUES ('16091231049', '123462', '张帅威', '13151289184', 1);
INSERT INTO `ccit_user` VALUES ('16091231050', '123463', '张帅威', '13151289185', 1);
INSERT INTO `ccit_user` VALUES ('16091231051', '123464', '张帅威', '13151289186', 1);
INSERT INTO `ccit_user` VALUES ('16091231052', '123465', '张帅威', '13151289187', 1);

SET FOREIGN_KEY_CHECKS = 1;
