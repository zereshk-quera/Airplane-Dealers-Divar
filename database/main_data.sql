--
-- PostgreSQL database dump
--

-- Dumped from database version 14.8 (Ubuntu 14.8-0ubuntu0.22.04.1)
-- Dumped by pg_dump version 14.8 (Ubuntu 14.8-0ubuntu0.22.04.1)

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

INSERT INTO public.categories (id, name) VALUES (1, 'small-passenger');
INSERT INTO public.categories (id, name) VALUES (2, 'big-passenger');

INSERT INTO public.users (id, username, password, role, token, is_active) VALUES (1, 'matin', '$2a$10$c5NLMAH6NEyN.R/YJ5V7MuD5YXeR05ClP42vh/YkuGH4k40Zvhx7G', 'Matin', 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhZG1pbiI6ZmFsc2UsImV4cCI6MTY4OTA4NzI4NywiaWQiOjR9.4dw7tG9ZQoEUlRLzaJYi352awWPU8sjWfYCVPUqqAVE', true);