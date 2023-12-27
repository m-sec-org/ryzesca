/*
 Navicat Premium Data Transfer

 Source Server         : 本机mysql
 Source Server Type    : MySQL
 Source Server Version : 80031
 Source Host           : localhost:3306
 Source Schema         : themis

 Target Server Type    : MySQL
 Target Server Version : 80031
 File Encoding         : 65001

 Date: 15/08/2023 09:46:38
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for cnnvd_metric
-- ----------------------------
DROP TABLE IF EXISTS `cnnvd_metric`;
CREATE TABLE `cnnvd_metric`  (
  `id` int(0) NOT NULL,
  `nssvd_id` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `cnnvd_id` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `impact_score` double NULL DEFAULT NULL,
  `access_vector` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `cnnvdmetric_cnnvd_id`(`cnnvd_id`) USING BTREE,
  INDEX `cnnvdmetric_nssvd_id`(`nssvd_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for cnvd_metric
-- ----------------------------
DROP TABLE IF EXISTS `cnvd_metric`;
CREATE TABLE `cnvd_metric`  (
  `id` int(0) NOT NULL,
  `nssvd_id` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `cnvd_id` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `impact_score` double NULL DEFAULT NULL,
  `access_vector` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL,
  `patch_name` varchar(2000) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL,
  `patch_description` text CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `cnvdmetric_cnvd_id`(`cnvd_id`) USING BTREE,
  INDEX `cnvdmetric_nssvd_id`(`nssvd_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for cve_infos
-- ----------------------------
DROP TABLE IF EXISTS `cve_infos`;
CREATE TABLE `cve_infos`  (
  `nssvd_id` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `cve_id` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL,
  `cnnvd_id` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL,
  `cnnvd_name` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL,
  `cnvd_id` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL,
  `cnvd_name` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL,
  `is_metric2` int(0) NOT NULL,
  `is_metric3` int(0) NOT NULL,
  `is_cnnvd_metric` int(0) NOT NULL,
  `is_cnvd_metric` int(0) NOT NULL,
  `is_rela` int(0) NOT NULL,
  `description_en` text CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL,
  `description_zh` text CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL,
  `solution_en` text CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL,
  `solution_zh` text CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL,
  `published_date` datetime(0) NULL DEFAULT NULL,
  `last_modified_date` datetime(0) NULL DEFAULT NULL,
  PRIMARY KEY (`nssvd_id`) USING BTREE,
  INDEX `cveinfos_cnnvd_id`(`cnnvd_id`) USING BTREE,
  INDEX `cveinfos_cnvd_id`(`cnvd_id`) USING BTREE,
  INDEX `cveinfos_cve_id`(`cve_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for cve_metric2
-- ----------------------------
DROP TABLE IF EXISTS `cve_metric2`;
CREATE TABLE `cve_metric2`  (
  `id` int(0) NOT NULL,
  `nssvd_id` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `severity` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL,
  `exploitability_score` double NULL DEFAULT NULL,
  `impact_score` double NULL DEFAULT NULL,
  `ac_insuf_info` int(0) NULL DEFAULT NULL,
  `obtain_all_privilege` int(0) NULL DEFAULT NULL,
  `obtain_user_privilege` int(0) NULL DEFAULT NULL,
  `obtain_other_privilege` int(0) NULL DEFAULT NULL,
  `user_interaction_required` int(0) NULL DEFAULT NULL,
  `version` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL,
  `vector_string` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL,
  `access_vector` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL,
  `access_complexity` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL,
  `authentication` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL,
  `confidentiality_impact` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL,
  `integrity_impact` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL,
  `availability_impact` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL,
  `base_score` double NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `cvemetric2_nssvd_id`(`nssvd_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for cve_metric3
-- ----------------------------
DROP TABLE IF EXISTS `cve_metric3`;
CREATE TABLE `cve_metric3`  (
  `id` int(0) NOT NULL,
  `nssvd_id` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `exploitability_score` double NULL DEFAULT NULL,
  `impact_score` double NULL DEFAULT NULL,
  `version` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL,
  `vector_string` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL,
  `attack_vector` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL,
  `attack_complexity` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL,
  `privileges_required` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL,
  `user_interaction` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL,
  `scope` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL,
  `confidentiality_impact` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL,
  `integrity_impact` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL,
  `availability_impact` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL,
  `base_score` double NULL DEFAULT NULL,
  `base_severity` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `cvemetric3_nssvd_id`(`nssvd_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for cve_references
-- ----------------------------
DROP TABLE IF EXISTS `cve_references`;
CREATE TABLE `cve_references`  (
  `id` int(0) NOT NULL,
  `nssvd_id` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `refsource` varchar(1024) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL,
  `name` text CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL,
  `url` text CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL,
  `tags` text CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `cvereferences_nssvd_id`(`nssvd_id`) USING BTREE,
  CONSTRAINT `cve_references_ibfk_1` FOREIGN KEY (`nssvd_id`) REFERENCES `cve_infos` (`nssvd_id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for cve_rela
-- ----------------------------
DROP TABLE IF EXISTS `cve_rela`;
CREATE TABLE `cve_rela`  (
  `nssvd_id` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `cwe_id` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `cwe_name` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL,
  PRIMARY KEY (`nssvd_id`, `cwe_id`) USING BTREE,
  INDEX `cverela_cwe_id`(`cwe_id`) USING BTREE,
  INDEX `cverela_nssvd_id`(`nssvd_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for cve_software
-- ----------------------------
DROP TABLE IF EXISTS `cve_software`;
CREATE TABLE `cve_software`  (
  `id` int(0) NOT NULL,
  `nssvd_id` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `vendor` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `product` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `version` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `update_version` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `cpe` varchar(1024) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `version_end_excluding` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL,
  `version_end_including` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL,
  `version_start_excluding` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL,
  `version_start_including` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL,
  `vulnerable` int(0) NOT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `cvesoftware_nssvd_id`(`nssvd_id`) USING BTREE,
  INDEX `cvesoftware_product`(`product`) USING BTREE,
  INDEX `cvesoftware_update_version`(`update_version`) USING BTREE,
  INDEX `cvesoftware_vendor`(`vendor`) USING BTREE,
  INDEX `cvesoftware_version`(`version`) USING BTREE,
  CONSTRAINT `cve_software_ibfk_1` FOREIGN KEY (`nssvd_id`) REFERENCES `cve_infos` (`nssvd_id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci ROW_FORMAT = Dynamic;
