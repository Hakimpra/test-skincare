-- phpMyAdmin SQL Dump
-- version 5.0.2
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Generation Time: Aug 20, 2020 at 06:16 PM
-- Server version: 8.0.19
-- PHP Version: 7.4.4

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `test-skincare2`
--

-- --------------------------------------------------------

--
-- Table structure for table `pesantreatment`
--

CREATE TABLE `pesantreatment` (
  `id` int NOT NULL,
  `user_id` varchar(7) NOT NULL,
  `treatment_id` varchar(7) NOT NULL,
  `total_bayar` int NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

--
-- Dumping data for table `pesantreatment`
--

INSERT INTO `pesantreatment` (`id`, `user_id`, `treatment_id`, `total_bayar`) VALUES
(3, '1', '1', 120000);

-- --------------------------------------------------------

--
-- Table structure for table `treatment`
--

CREATE TABLE `treatment` (
  `id` int NOT NULL,
  `title` varchar(50) NOT NULL DEFAULT '0',
  `harga` int NOT NULL DEFAULT '0',
  `keterangan` varchar(250) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '0'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

--
-- Dumping data for table `treatment`
--

INSERT INTO `treatment` (`id`, `title`, `harga`, `keterangan`) VALUES
(1, 'ACNE CLEAR TREATMENT', 120000, 'Perawatan kulit dengan IPL yang di padukan dengan serum acne yang efektif membunuh bakteri penyebab jerawat, mengatasi segera jerawat meradang, mengeringkan jerawat dan membuat jerawat tidak mudah kambuh.'),
(2, 'BREAST TREATMENT', 300000, 'Treatment untuk mengencangkan payudara dengan IPL panjang gelombang 750 – 1200nm.'),
(3, 'coba edit 222', 140000, 'keteranan coba edit'),
(6, 'coba', 140000, 'keteranan coba');

-- --------------------------------------------------------

--
-- Table structure for table `users`
--

CREATE TABLE `users` (
  `id` int NOT NULL,
  `username` varchar(50) NOT NULL,
  `password` varchar(120) NOT NULL,
  `email` varchar(200) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

--
-- Dumping data for table `users`
--

INSERT INTO `users` (`id`, `username`, `password`, `email`) VALUES
(1, 'HakimPra2', '$2a$10$1T1sxdGoStVG9Kikedz0lu4qCulcj/h2daQxrGu.MJZDtuxqfFIOy', 'hakimpra@kelegan.nett');

--
-- Indexes for dumped tables
--

--
-- Indexes for table `pesantreatment`
--
ALTER TABLE `pesantreatment`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `treatment`
--
ALTER TABLE `treatment`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `users`
--
ALTER TABLE `users`
  ADD PRIMARY KEY (`id`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `pesantreatment`
--
ALTER TABLE `pesantreatment`
  MODIFY `id` int NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=4;

--
-- AUTO_INCREMENT for table `treatment`
--
ALTER TABLE `treatment`
  MODIFY `id` int NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=7;

--
-- AUTO_INCREMENT for table `users`
--
ALTER TABLE `users`
  MODIFY `id` int NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=2;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
