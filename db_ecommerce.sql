-- phpMyAdmin SQL Dump
-- version 5.2.0
-- https://www.phpmyadmin.net/
--
-- Host: localhost
-- Generation Time: Sep 07, 2022 at 07:13 PM
-- Server version: 10.4.21-MariaDB
-- PHP Version: 7.4.29

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `db_ecommerce`
--

-- --------------------------------------------------------

--
-- Table structure for table `categories`
--

CREATE TABLE `categories` (
  `id` int(11) NOT NULL,
  `name` varchar(255) NOT NULL,
  `image` varchar(255) NOT NULL,
  `created_at` varchar(64) NOT NULL,
  `updated_at` varchar(64) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `categories`
--

INSERT INTO `categories` (`id`, `name`, `image`, `created_at`, `updated_at`) VALUES
(1, 'Fashion Pria', '/assets/images/categories/fashion_pria.png', '2022-08-29 16:35:24.446988 +0700 WIB', '2022-08-29 16:35:24.446994 +0700 WIB');

-- --------------------------------------------------------

--
-- Table structure for table `products_draft`
--

CREATE TABLE `products_draft` (
  `id` int(11) NOT NULL,
  `name` varchar(255) NOT NULL,
  `condition` varchar(16) NOT NULL,
  `price` int(32) NOT NULL,
  `discount` int(32) NOT NULL,
  `discount_type` int(1) NOT NULL,
  `weight` int(32) NOT NULL,
  `id_category` int(8) NOT NULL,
  `description` text NOT NULL,
  `minimum` int(11) NOT NULL,
  `status` int(1) NOT NULL,
  `created_at` varchar(64) NOT NULL,
  `updated_at` varchar(64) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- --------------------------------------------------------

--
-- Table structure for table `products_image_draft`
--

CREATE TABLE `products_image_draft` (
  `id` int(11) NOT NULL,
  `id_product` int(11) NOT NULL,
  `image` varchar(255) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- --------------------------------------------------------

--
-- Table structure for table `products_similiar_drafts`
--

CREATE TABLE `products_similiar_drafts` (
  `id` int(11) NOT NULL,
  `id_product` int(11) NOT NULL,
  `title` varchar(255) NOT NULL,
  `description` varchar(255) NOT NULL,
  `link` varchar(255) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- --------------------------------------------------------

--
-- Table structure for table `promos_slider`
--

CREATE TABLE `promos_slider` (
  `id` int(11) NOT NULL,
  `name` varchar(255) NOT NULL,
  `description` varchar(255) NOT NULL,
  `image` varchar(255) NOT NULL,
  `kode_promo` varchar(255) NOT NULL,
  `expired` int(64) NOT NULL,
  `status` int(1) NOT NULL,
  `created_at` varchar(64) NOT NULL,
  `updated_at` varchar(64) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `promos_slider`
--

INSERT INTO `promos_slider` (`id`, `name`, `description`, `image`, `kode_promo`, `expired`, `status`, `created_at`, `updated_at`) VALUES
(1, 'Semarak Kemerdekaan', 'Promo akan dimulai pada tanggal 10 Agustus - 10 September 2022', '/assets/images/promos/semarak_kemerdekaan.png', 'KEMERDEKAAN', 3800, 1, '2022-08-29 16:32:53.293513 +0700 WIB', '2022-08-29 16:32:53.293519 +0700 WIB');

-- --------------------------------------------------------

--
-- Table structure for table `users`
--

CREATE TABLE `users` (
  `id` int(11) NOT NULL,
  `username` varchar(255) NOT NULL,
  `email` varchar(255) NOT NULL,
  `password` varchar(255) NOT NULL,
  `phone_number` varchar(16) NOT NULL,
  `signup_with` int(2) NOT NULL,
  `role` int(1) NOT NULL,
  `created_at` varchar(64) NOT NULL,
  `updated_at` varchar(64) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `users`
--

INSERT INTO `users` (`id`, `username`, `email`, `password`, `phone_number`, `signup_with`, `role`, `created_at`, `updated_at`) VALUES
(1, 'riyan', 'riyan@gmail.com', '$2a$14$.5/xSocI00B83rpL/G3vDu2VA.If0H2FV6xYZjqcdjvs.YljL3kPa', '+6281381161992', 1, 2, '2022-08-27 20:46:46.662026 +0700 WIB', '2022-08-27 20:46:46.662153 +0700 WIB'),
(2, 'superuser', 'superuser@admin.com', '$2a$14$B0hPfhi3lshXe4GTqLo1MeAzzxlfU2OaObgzuG9sGqLQNl/jsyTHm', '+6281381161992', 1, 1, '2022-08-28 00:31:03.945684 +0700 WIB', '2022-08-28 00:31:03.945688 +0700 WIB');

-- --------------------------------------------------------

--
-- Table structure for table `users_address`
--

CREATE TABLE `users_address` (
  `id` int(11) NOT NULL,
  `id_user` int(11) NOT NULL,
  `latitude` double NOT NULL,
  `longitude` double NOT NULL,
  `country` varchar(255) NOT NULL,
  `province` varchar(255) NOT NULL,
  `city` varchar(255) NOT NULL,
  `district` varchar(255) NOT NULL,
  `postal_code` int(6) NOT NULL,
  `address` varchar(255) NOT NULL,
  `recipient` varchar(255) NOT NULL,
  `address_name` varchar(255) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- --------------------------------------------------------

--
-- Table structure for table `users_token`
--

CREATE TABLE `users_token` (
  `id` int(11) NOT NULL,
  `id_user` int(11) NOT NULL,
  `fcm_token` varchar(255) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Indexes for dumped tables
--

--
-- Indexes for table `categories`
--
ALTER TABLE `categories`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `products_draft`
--
ALTER TABLE `products_draft`
  ADD PRIMARY KEY (`id`),
  ADD KEY `category_foreign` (`id_category`);

--
-- Indexes for table `products_image_draft`
--
ALTER TABLE `products_image_draft`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `promos_slider`
--
ALTER TABLE `promos_slider`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `users`
--
ALTER TABLE `users`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `users_address`
--
ALTER TABLE `users_address`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `users_token`
--
ALTER TABLE `users_token`
  ADD PRIMARY KEY (`id`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `categories`
--
ALTER TABLE `categories`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=2;

--
-- AUTO_INCREMENT for table `products_draft`
--
ALTER TABLE `products_draft`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `products_image_draft`
--
ALTER TABLE `products_image_draft`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `promos_slider`
--
ALTER TABLE `promos_slider`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=2;

--
-- AUTO_INCREMENT for table `users`
--
ALTER TABLE `users`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=3;

--
-- AUTO_INCREMENT for table `users_address`
--
ALTER TABLE `users_address`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `users_token`
--
ALTER TABLE `users_token`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
