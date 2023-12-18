CREATE TABLE `users` (
  `id` int AUTO_INCREMENT PRIMARY KEY,
  `username` varchar(255) NOT NULL,
  `email` varchar(255) NOT NULL,
  `password` varchar(255) NOT NULL,
  `gender` enum('mele','femele','other') NOT NULL DEFAULT 'other',
  `birthday` varchar(10) DEFAULT NULL,
  `points` int NOT NULL DEFAULT 0,
  `created_at` timestamp NOT NULL,
  `updated_at` timestamp DEFAULT NULL
);