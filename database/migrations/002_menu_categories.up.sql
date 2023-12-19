CREATE TABLE `menu_categories` (
  `id` integer AUTO_INCREMENT PRIMARY KEY,
  `name` varchar(255) NOT NULL,
  `created_at` timestamp NOT NULL,
  `updated_at` timestamp DEFAULT NULL
);