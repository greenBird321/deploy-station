-- phpMyAdmin SQL Dump
-- version 4.7.8
-- https://www.phpmyadmin.net/
--
-- Host: localhost
-- Generation Time: 2018-10-26 10:17:59
-- 服务器版本： 5.7.21
-- PHP Version: 7.1.13

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
SET AUTOCOMMIT = 0;
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `deploy_station`
--

-- --------------------------------------------------------

--
-- 表的结构 `deploy_logs`
--

CREATE TABLE `deploy_logs` (
  `id` int(11) NOT NULL,
  `item_id` int(11) DEFAULT '0',
  `node_id` int(11) DEFAULT '0',
  `version` float DEFAULT NULL,
  `ctime` int(11) DEFAULT '0',
  `status` enum('process','success','failure') DEFAULT NULL,
  `message` varchar(2000) DEFAULT ''
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- --------------------------------------------------------

--
-- 表的结构 `items`
--

CREATE TABLE `items` (
  `id` int(11) NOT NULL,
  `name` varchar(128) DEFAULT '',
  `repo_url` varchar(512) DEFAULT '',
  `repo_type` enum('git','svn') DEFAULT 'git',
  `repo_private_key` varchar(2000) DEFAULT '',
  `remark` varchar(128) DEFAULT '',
  `notify` varchar(512) DEFAULT '',
  `ctime` int(11) DEFAULT '0',
  `mtime` int(11) DEFAULT '0'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- --------------------------------------------------------

--
-- 表的结构 `items_nodes`
--

CREATE TABLE `items_nodes` (
  `id` int(11) NOT NULL,
  `item_id` int(11) DEFAULT NULL,
  `node_id` int(11) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- --------------------------------------------------------

--
-- 表的结构 `nodes`
--

CREATE TABLE `nodes` (
  `id` int(11) NOT NULL,
  `ip` varchar(128) DEFAULT '',
  `ip_intranet` varchar(128) NOT NULL DEFAULT '',
  `ssh_port` int(11) DEFAULT '0',
  `os` enum('centos','ubuntu','freebsd','debian') DEFAULT NULL,
  `os_version` float NOT NULL,
  `region` varchar(64) NOT NULL DEFAULT '',
  `remark` varchar(255) DEFAULT '',
  `ctime` int(11) DEFAULT '0',
  `mtime` int(11) DEFAULT '0'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Indexes for dumped tables
--

--
-- Indexes for table `deploy_logs`
--
ALTER TABLE `deploy_logs`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `items`
--
ALTER TABLE `items`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `items_name_uindex` (`name`);

--
-- Indexes for table `items_nodes`
--
ALTER TABLE `items_nodes`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `nodes`
--
ALTER TABLE `nodes`
  ADD PRIMARY KEY (`id`);

--
-- 在导出的表使用AUTO_INCREMENT
--

--
-- 使用表AUTO_INCREMENT `deploy_logs`
--
ALTER TABLE `deploy_logs`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT;

--
-- 使用表AUTO_INCREMENT `items`
--
ALTER TABLE `items`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT;

--
-- 使用表AUTO_INCREMENT `items_nodes`
--
ALTER TABLE `items_nodes`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT;

--
-- 使用表AUTO_INCREMENT `nodes`
--
ALTER TABLE `nodes`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
