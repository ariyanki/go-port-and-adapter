-- phpMyAdmin SQL Dump
-- version 5.2.1
-- https://www.phpmyadmin.net/
--
-- Host: localhost
-- Generation Time: Oct 03, 2023 at 06:21 AM
-- Server version: 11.1.2-MariaDB
-- PHP Version: 8.2.10

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `template_user_management`
--

-- --------------------------------------------------------

--
-- Table structure for table `app_permissions`
--

CREATE TABLE `app_permissions` (
  `id` int(11) NOT NULL,
  `resource` varchar(255) NOT NULL,
  `actions` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL CHECK (json_valid(`actions`))
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3 COLLATE=utf8mb3_general_ci;

-- --------------------------------------------------------

--
-- Table structure for table `app_roles`
--

CREATE TABLE `app_roles` (
  `id` int(10) NOT NULL,
  `role_name` varchar(255) NOT NULL,
  `status` int(11) DEFAULT NULL,
  `created_by` int(11) DEFAULT NULL,
  `updated_by` int(11) DEFAULT NULL,
  `created_at` timestamp(6) NOT NULL DEFAULT current_timestamp(6),
  `updated_at` timestamp(6) NOT NULL DEFAULT current_timestamp(6),
  `deleted_at` timestamp NULL DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3 COLLATE=utf8mb3_general_ci;

--
-- Dumping data for table `app_roles`
--

INSERT INTO `app_roles` (`id`, `role_name`, `status`, `created_by`, `updated_by`, `created_at`, `updated_at`, `deleted_at`) VALUES
(1, 'superadmin', 1, 1, 1, '2020-07-19 00:30:15.463891', '2020-07-19 00:30:15.463877', NULL),
(2, 'admin', 1, 1, 1, '2020-07-19 00:30:15.470720', '2020-07-19 00:30:15.470704', NULL),
(3, 'user', 1, 1, 1, '2020-07-19 00:30:15.476482', '2020-07-19 00:30:15.476470', NULL);

-- --------------------------------------------------------

--
-- Table structure for table `app_role_permissions`
--

CREATE TABLE `app_role_permissions` (
  `id` int(10) NOT NULL,
  `role_id` int(11) NOT NULL,
  `permission_id` int(11) NOT NULL,
  `status` int(11) DEFAULT NULL,
  `created_by` int(11) DEFAULT NULL,
  `updated_by` int(11) DEFAULT NULL,
  `created_at` timestamp(6) NOT NULL DEFAULT current_timestamp(6),
  `updated_at` timestamp(6) NOT NULL DEFAULT current_timestamp(6),
  `deleted_at` timestamp NULL DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3 COLLATE=utf8mb3_general_ci;

-- --------------------------------------------------------

--
-- Table structure for table `app_users`
--

CREATE TABLE `app_users` (
  `id` bigint(10) NOT NULL,
  `facebook_login_id` varchar(255) DEFAULT NULL,
  `google_login_id` varchar(255) DEFAULT NULL,
  `username` varchar(255) NOT NULL,
  `password` varchar(500) NOT NULL,
  `password_salt` varchar(500) DEFAULT NULL,
  `fullname` varchar(255) NOT NULL,
  `phonenumber` varchar(255) NOT NULL,
  `email` varchar(255) NOT NULL,
  `longitude` varchar(255) DEFAULT NULL,
  `latitude` varchar(255) DEFAULT NULL,
  `photo_filename` varchar(255) DEFAULT NULL,
  `is_loggedin` int(11) NOT NULL DEFAULT 0,
  `login_attempt` int(11) NOT NULL DEFAULT 0,
  `last_loggedin_at` timestamp(6) NULL DEFAULT NULL,
  `access_source` varchar(255) DEFAULT NULL,
  `last_access_at` timestamp(6) NULL DEFAULT NULL,
  `status` int(11) NOT NULL DEFAULT 0,
  `created_by` int(11) DEFAULT NULL,
  `updated_by` int(11) DEFAULT NULL,
  `created_at` timestamp(6) NOT NULL DEFAULT current_timestamp(6),
  `updated_at` timestamp(6) NOT NULL DEFAULT current_timestamp(6),
  `deleted_at` timestamp NULL DEFAULT NULL,
  `google_login_data` varchar(10000) DEFAULT NULL,
  `facebook_login_data` varchar(10000) DEFAULT NULL,
  `birthdate` date DEFAULT NULL,
  `gender` varchar(10) DEFAULT NULL,
  `fcm_token` varchar(1000) DEFAULT '',
  `city` varchar(255) DEFAULT NULL,
  `address` text DEFAULT NULL,
  `merchant_id` int(11) NOT NULL DEFAULT 0,
  `apple_login_id` varchar(255) DEFAULT NULL,
  `apple_login_data` text DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1 COLLATE=latin1_swedish_ci;

--
-- Dumping data for table `app_users`
--

INSERT INTO `app_users` (`id`, `facebook_login_id`, `google_login_id`, `username`, `password`, `password_salt`, `fullname`, `phonenumber`, `email`, `longitude`, `latitude`, `photo_filename`, `is_loggedin`, `login_attempt`, `last_loggedin_at`, `access_source`, `last_access_at`, `status`, `created_by`, `updated_by`, `created_at`, `updated_at`, `deleted_at`, `google_login_data`, `facebook_login_data`, `birthdate`, `gender`, `fcm_token`, `city`, `address`, `merchant_id`, `apple_login_id`, `apple_login_data`) VALUES
(1, NULL, NULL, 'admin@admin.com', '26fce84bdf52dced60502b2f126140d3edfbd5b0d7c6586bb79b5cb45680debf', 'e02e8140-06a1-45ce-9240-b610ee58add4', 'Administrator', '0824324234', 'admin@admin.com', NULL, NULL, 'admin.png', 1, 1, '2020-11-19 14:13:56.000000', NULL, NULL, 1, 1, 1, '2019-11-20 17:45:48.337661', '2020-11-19 14:13:56.923115', NULL, NULL, NULL, '1992-12-31', 'Men', 'fjg11e1OaWg:APA91bEyO5n1Orm8eXexOWnPoEdUnBbod40fqeUoplGmOPLIygqNnvjmmzjuRRX0YxolYmtwxP4LDaEw2B9IHs9yQZqaldyzYVA9KfFIiy-_e-TWkulFKP2SUTJ4_-3uqd4A22CV5v6D', NULL, NULL, 0, NULL, NULL);

-- --------------------------------------------------------

--
-- Table structure for table `app_user_roles`
--

CREATE TABLE `app_user_roles` (
  `id` bigint(10) NOT NULL,
  `user_id` int(11) NOT NULL,
  `role_id` int(11) NOT NULL,
  `status` int(11) DEFAULT NULL,
  `created_by` int(11) DEFAULT NULL,
  `updated_by` int(11) DEFAULT NULL,
  `created_at` timestamp(6) NOT NULL DEFAULT current_timestamp(6),
  `updated_at` timestamp(6) NOT NULL DEFAULT current_timestamp(6),
  `deleted_at` timestamp NULL DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3 COLLATE=utf8mb3_general_ci;

--
-- Dumping data for table `app_user_roles`
--

INSERT INTO `app_user_roles` (`id`, `user_id`, `role_id`, `status`, `created_by`, `updated_by`, `created_at`, `updated_at`, `deleted_at`) VALUES
(1, 1, 1, 1, 1, 1, '2023-09-10 07:31:46.603980', '2023-09-10 07:31:46.603980', NULL);

--
-- Indexes for dumped tables
--

--
-- Indexes for table `app_permissions`
--
ALTER TABLE `app_permissions`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `app_roles`
--
ALTER TABLE `app_roles`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `app_role_permissions`
--
ALTER TABLE `app_role_permissions`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `app_users`
--
ALTER TABLE `app_users`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `app_user_roles`
--
ALTER TABLE `app_user_roles`
  ADD PRIMARY KEY (`id`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `app_permissions`
--
ALTER TABLE `app_permissions`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `app_roles`
--
ALTER TABLE `app_roles`
  MODIFY `id` int(10) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=4;

--
-- AUTO_INCREMENT for table `app_role_permissions`
--
ALTER TABLE `app_role_permissions`
  MODIFY `id` int(10) NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `app_users`
--
ALTER TABLE `app_users`
  MODIFY `id` bigint(10) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=2;

--
-- AUTO_INCREMENT for table `app_user_roles`
--
ALTER TABLE `app_user_roles`
  MODIFY `id` bigint(10) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=2;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
