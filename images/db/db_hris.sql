-- phpMyAdmin SQL Dump
-- version 5.2.0
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Generation Time: Dec 29, 2022 at 01:58 AM
-- Server version: 10.4.24-MariaDB
-- PHP Version: 8.1.6

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `db_hris`
--

-- --------------------------------------------------------

--
-- Table structure for table `companies`
--

CREATE TABLE `companies` (
  `id` int(11) NOT NULL,
  `name` varchar(256) NOT NULL,
  `secret_key` text NOT NULL,
  `created_at` varchar(64) NOT NULL,
  `updated_at` varchar(64) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `companies`
--

INSERT INTO `companies` (`id`, `name`, `secret_key`, `created_at`, `updated_at`) VALUES
(1, 'Arkademi', 'FXpfuQySAfJP0yE1zVyX5HDY5n01Di1LLFLOGy7mmbwXC45JcRXNVjW41g4EwJ7T', '2022-11-26T23:43:43+07:00', '2022-11-26T23:43:43+07:00'),
(2, 'PT Senang Sentosa', 'MqHZlAI6UwfQ3xJybhxtNhojSqYJjlex0r3464lF2S3pQnn0sNoxyP6LANM1U4hN', '2022-11-28T23:12:26+07:00', '2022-11-28T23:12:26+07:00');

-- --------------------------------------------------------

--
-- Table structure for table `menu`
--

CREATE TABLE `menu` (
  `id` int(11) NOT NULL,
  `title` varchar(255) NOT NULL,
  `image` varchar(255) NOT NULL,
  `status` int(11) NOT NULL,
  `route` varchar(64) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- --------------------------------------------------------

--
-- Table structure for table `users`
--

CREATE TABLE `users` (
  `id` int(11) NOT NULL,
  `employee_id` varchar(64) NOT NULL,
  `username` varchar(255) NOT NULL,
  `email` varchar(255) NOT NULL,
  `password` varchar(255) NOT NULL,
  `phone_number` varchar(16) NOT NULL,
  `role` int(1) NOT NULL,
  `company_secret_key` varchar(64) NOT NULL,
  `created_at` varchar(64) NOT NULL,
  `updated_at` varchar(64) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `users`
--

INSERT INTO `users` (`id`, `employee_id`, `username`, `email`, `password`, `phone_number`, `role`, `company_secret_key`, `created_at`, `updated_at`) VALUES
(1, '2022.02.02', 'superadmin', 'super@admin.com', '$2a$14$/Y3KP/J/cj/ZZN1S6IVhzOOhbSkczaXhGKibo8/zG5M8tPXsMQHGO', '-', 1, 'MqHZlAI6UwfQ3xJybhxtNhojSqYJjlex0r3464lF2S3pQnn0sNoxyP6LANM1U4hN', '2022-11-28T23:15:59+07:00', '2022-11-28T23:15:59+07:00');

-- --------------------------------------------------------

--
-- Table structure for table `user_absent`
--

CREATE TABLE `user_absent` (
  `id` int(11) NOT NULL,
  `id_user` int(11) NOT NULL,
  `id_employe` varchar(256) NOT NULL,
  `tanggal` varchar(64) NOT NULL,
  `latitude` double NOT NULL,
  `longitude` double NOT NULL,
  `catatan` text NOT NULL,
  `tipe` int(1) NOT NULL,
  `photo` text NOT NULL,
  `created_at` varchar(64) NOT NULL,
  `updated_at` varchar(64) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- --------------------------------------------------------

--
-- Table structure for table `user_company`
--

CREATE TABLE `user_company` (
  `id` int(11) NOT NULL,
  `id_user` int(11) NOT NULL,
  `id_company` int(11) NOT NULL,
  `status` int(1) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- --------------------------------------------------------

--
-- Table structure for table `user_info_pekerjaan`
--

CREATE TABLE `user_info_pekerjaan` (
  `id` int(11) NOT NULL,
  `employee_id` int(11) NOT NULL,
  `nama_perusahaan` int(11) NOT NULL,
  `cabang` int(11) NOT NULL,
  `divisi` int(11) NOT NULL,
  `posisi` int(11) NOT NULL,
  `level_pekerjaan` int(11) NOT NULL,
  `status_pekerjaan` int(11) NOT NULL,
  `tanggal_bergabung` int(11) NOT NULL,
  `tanggal_berakhir` int(11) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- --------------------------------------------------------

--
-- Table structure for table `user_info_personal`
--

CREATE TABLE `user_info_personal` (
  `id` int(11) NOT NULL,
  `id_user` int(11) NOT NULL,
  `nama_depan` varchar(256) NOT NULL,
  `nama_belakang` varchar(256) NOT NULL,
  `jenis_kelamin` varchar(16) NOT NULL,
  `tempat_lahir` varchar(256) NOT NULL,
  `tanggal_lahir` varchar(64) NOT NULL,
  `no_hp` varchar(16) NOT NULL,
  `telepon` varchar(16) NOT NULL,
  `status_pernikahan` varchar(128) NOT NULL,
  `agama` varchar(32) NOT NULL,
  `nomor_id` varchar(32) NOT NULL,
  `tipe_id` varchar(32) NOT NULL,
  `tanggal_kadaluarsa` varchar(32) NOT NULL,
  `alamat_ktp` text NOT NULL,
  `alamat_domisili` text NOT NULL,
  `golongan_darah` varchar(8) NOT NULL,
  `created_at` varchar(64) NOT NULL,
  `updated_at` varchar(644) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- --------------------------------------------------------

--
-- Table structure for table `user_payroll`
--

CREATE TABLE `user_payroll` (
  `id` int(11) NOT NULL,
  `id_user` int(32) NOT NULL,
  `name` varchar(256) NOT NULL,
  `amount` int(64) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Indexes for dumped tables
--

--
-- Indexes for table `companies`
--
ALTER TABLE `companies`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `menu`
--
ALTER TABLE `menu`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `users`
--
ALTER TABLE `users`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `user_absent`
--
ALTER TABLE `user_absent`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `user_company`
--
ALTER TABLE `user_company`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `user_info_pekerjaan`
--
ALTER TABLE `user_info_pekerjaan`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `user_info_personal`
--
ALTER TABLE `user_info_personal`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `user_payroll`
--
ALTER TABLE `user_payroll`
  ADD PRIMARY KEY (`id`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `companies`
--
ALTER TABLE `companies`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=3;

--
-- AUTO_INCREMENT for table `menu`
--
ALTER TABLE `menu`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `users`
--
ALTER TABLE `users`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=2;

--
-- AUTO_INCREMENT for table `user_absent`
--
ALTER TABLE `user_absent`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `user_company`
--
ALTER TABLE `user_company`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `user_info_pekerjaan`
--
ALTER TABLE `user_info_pekerjaan`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `user_info_personal`
--
ALTER TABLE `user_info_personal`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=8;

--
-- AUTO_INCREMENT for table `user_payroll`
--
ALTER TABLE `user_payroll`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
