CREATE TABLE `menu` (
  `id` integer PRIMARY KEY AUTO_INCREMENT,
  `menu_categories_id` integer,
  `name` varchar(255),
  `description` varchar(255),
  `price` decimal(10, 2),
  `created_at` timestamp,
  `updated_at` timestamp
);
