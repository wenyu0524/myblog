-- PostgreSQL 初始化脚本：创建 user_db 和 blog_db 两个数据库及其表结构
-- 此脚本由 Docker PostgreSQL 镜像在首次启动时自动执行

-- =========================
-- 创建用户数据库
-- =========================
CREATE DATABASE user_db;

\c user_db;

CREATE TABLE users (
    id         BIGSERIAL PRIMARY KEY,
    username   VARCHAR(64)  NOT NULL UNIQUE,
    password   VARCHAR(255) NOT NULL,
    email      VARCHAR(128) NOT NULL,
    created_at TIMESTAMPTZ  NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ  NOT NULL DEFAULT NOW()
);

-- =========================
-- 创建博客数据库
-- =========================
CREATE DATABASE blog_db;

\c blog_db;

CREATE TABLE posts (
    id         BIGSERIAL PRIMARY KEY,
    user_id    BIGINT       NOT NULL,
    title      VARCHAR(255) NOT NULL,
    content    TEXT         NOT NULL,
    created_at TIMESTAMPTZ  NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ  NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_posts_user_id ON posts (user_id);
