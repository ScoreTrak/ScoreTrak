-- Create "checks" table
CREATE TABLE `checks` (`id` integer NOT NULL PRIMARY KEY AUTOINCREMENT, `log` text NOT NULL, `error` text NOT NULL, `passed` bool NOT NULL, `service_checks` integer NULL, `service_properties` integer NULL, CONSTRAINT `checks_services_checks` FOREIGN KEY (`service_checks`) REFERENCES `services` (`id`) ON DELETE SET NULL, CONSTRAINT `checks_services_properties` FOREIGN KEY (`service_properties`) REFERENCES `services` (`id`) ON DELETE SET NULL);
-- Create "competitions" table
CREATE TABLE `competitions` (`id` integer NOT NULL PRIMARY KEY AUTOINCREMENT, `name` text NOT NULL, `started_at` datetime NOT NULL, `finished_at` datetime NOT NULL);
-- Create "competition_settings" table
CREATE TABLE `competition_settings` (`id` integer NOT NULL PRIMARY KEY AUTOINCREMENT);
-- Create "hosts" table
CREATE TABLE `hosts` (`id` integer NOT NULL PRIMARY KEY AUTOINCREMENT, `address` text NOT NULL, `address_list_range` text NOT NULL, `team_hosts` integer NULL, CONSTRAINT `hosts_teams_hosts` FOREIGN KEY (`team_hosts`) REFERENCES `teams` (`id`) ON DELETE SET NULL);
-- Create "host_groups" table
CREATE TABLE `host_groups` (`id` integer NOT NULL PRIMARY KEY AUTOINCREMENT);
-- Create "properties" table
CREATE TABLE `properties` (`id` integer NOT NULL PRIMARY KEY AUTOINCREMENT);
-- Create "rounds" table
CREATE TABLE `rounds` (`id` integer NOT NULL PRIMARY KEY AUTOINCREMENT, `create_time` datetime NOT NULL, `update_time` datetime NOT NULL, `round_number` integer NOT NULL, `note` text NOT NULL, `err` text NOT NULL, `started_at` datetime NOT NULL, `finished_at` datetime NOT NULL);
-- Create "services" table
CREATE TABLE `services` (`id` integer NOT NULL PRIMARY KEY AUTOINCREMENT, `name` text NOT NULL, `display_name` text NOT NULL, `weight` integer NOT NULL, `point_boost` integer NOT NULL, `round_units` integer NOT NULL, `reound_delay` integer NOT NULL, `host_services` integer NULL, CONSTRAINT `services_hosts_services` FOREIGN KEY (`host_services`) REFERENCES `hosts` (`id`) ON DELETE SET NULL);
-- Create "teams" table
CREATE TABLE `teams` (`id` integer NOT NULL PRIMARY KEY AUTOINCREMENT, `create_time` datetime NOT NULL, `update_time` datetime NOT NULL, `pause` bool NOT NULL, `hidden` bool NOT NULL, `name` text NOT NULL, `index` integer NOT NULL, `competition_teams` integer NULL, CONSTRAINT `teams_competitions_teams` FOREIGN KEY (`competition_teams`) REFERENCES `competitions` (`id`) ON DELETE SET NULL);
-- Create "users" table
CREATE TABLE `users` (`id` integer NOT NULL PRIMARY KEY AUTOINCREMENT, `name` text NOT NULL);
-- Create "team_users" table
CREATE TABLE `team_users` (`team_id` integer NOT NULL, `user_id` integer NOT NULL, PRIMARY KEY (`team_id`, `user_id`), CONSTRAINT `team_users_team_id` FOREIGN KEY (`team_id`) REFERENCES `teams` (`id`) ON DELETE CASCADE, CONSTRAINT `team_users_user_id` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE);
