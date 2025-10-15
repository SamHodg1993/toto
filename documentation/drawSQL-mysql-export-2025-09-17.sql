CREATE TABLE `todos`(
    `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
    `title` VARCHAR(255) NOT NULL,
    `description` TEXT NULL,
    `project_id` BIGINT NOT NULL,
    `completed` BOOLEAN NOT NULL,
    `completed_at` DATETIME NULL,
    `created_at` DATETIME NOT NULL,
    `updated_at` DATETIME NOT NULL,
    `jira_ticket_id` BIGINT NOT NULL
);
CREATE TABLE `projects`(
    `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
    `title` VARCHAR(255) NOT NULL,
    `description` TEXT NULL,
    `archived` BOOLEAN NOT NULL,
    `filepath` VARCHAR(255) NOT NULL COMMENT 'This is handled by the code, but the default is the current working directory',
    `created_at` DATETIME NOT NULL,
    `updated_at` DATETIME NOT NULL
);
CREATE TABLE `jira_tickets`(
    `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
    `jira_key` VARCHAR(50) NOT NULL,
    `title` VARCHAR(500) NOT NULL,
    `status` VARCHAR(50) NOT NULL,
    `project_key` VARCHAR(50) NULL,
    `issue_type` VARCHAR(50) NULL,
    `url` VARCHAR(500) NULL,
    `last_synced_at` DATETIME NULL,
    `created_at` DATETIME NOT NULL,
    `updated_at` DATETIME NOT NULL
);
ALTER TABLE
    `todos` ADD CONSTRAINT `todos_jira_ticket_id_foreign` FOREIGN KEY(`jira_ticket_id`) REFERENCES `jira_tickets`(`id`);
ALTER TABLE
    `todos` ADD CONSTRAINT `todos_project_id_foreign` FOREIGN KEY(`project_id`) REFERENCES `projects`(`id`);