--
-- PostgreSQL database dump
--

-- Dumped from database version 14.2
-- Dumped by pg_dump version 14.2

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

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: schema_migration; Type: TABLE; Schema: public; Owner: go
--

CREATE TABLE public.schema_migration (
    version character varying(14) NOT NULL
);


ALTER TABLE public.schema_migration OWNER TO go;

--
-- Name: stm_users; Type: TABLE; Schema: public; Owner: go
--

CREATE TABLE public.stm_users (
    id integer NOT NULL,
    first_name character varying(64) DEFAULT ''::character varying NOT NULL,
    last_name character varying(64) DEFAULT ''::character varying NOT NULL,
    email character varying(255) NOT NULL,
    password character varying(255) NOT NULL,
    access_level integer DEFAULT 1 NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE public.stm_users OWNER TO go;

--
-- Name: stm_users_id_seq; Type: SEQUENCE; Schema: public; Owner: go
--

CREATE SEQUENCE public.stm_users_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.stm_users_id_seq OWNER TO go;

--
-- Name: stm_users_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: go
--

ALTER SEQUENCE public.stm_users_id_seq OWNED BY public.stm_users.id;


--
-- Name: stm_users id; Type: DEFAULT; Schema: public; Owner: go
--

ALTER TABLE ONLY public.stm_users ALTER COLUMN id SET DEFAULT nextval('public.stm_users_id_seq'::regclass);


--
-- Name: schema_migration schema_migration_pkey; Type: CONSTRAINT; Schema: public; Owner: go
--

ALTER TABLE ONLY public.schema_migration
    ADD CONSTRAINT schema_migration_pkey PRIMARY KEY (version);


--
-- Name: stm_users stm_users_pkey; Type: CONSTRAINT; Schema: public; Owner: go
--

ALTER TABLE ONLY public.stm_users
    ADD CONSTRAINT stm_users_pkey PRIMARY KEY (id);


--
-- Name: schema_migration_version_idx; Type: INDEX; Schema: public; Owner: go
--

CREATE UNIQUE INDEX schema_migration_version_idx ON public.schema_migration USING btree (version);


--
-- PostgreSQL database dump complete
--

