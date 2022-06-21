--
-- PostgreSQL database dump
--

-- Dumped from database version 14.1
-- Dumped by pg_dump version 14.1

-- Started on 2022-06-21 16:36:57

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

--
-- TOC entry 5 (class 2615 OID 24661)
-- Name: desafiotecnicoprincipal; Type: SCHEMA; Schema: -; Owner: postgres
--

CREATE SCHEMA desafiotecnicoprincipal;


ALTER SCHEMA desafiotecnicoprincipal OWNER TO postgres;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- TOC entry 211 (class 1259 OID 24663)
-- Name: accounts; Type: TABLE; Schema: desafiotecnicoprincipal; Owner: postgres
--

CREATE TABLE desafiotecnicoprincipal.accounts (
    id bigint NOT NULL,
    name text NOT NULL,
    cpf text NOT NULL,
    secret text NOT NULL,
    balance bigint NOT NULL,
    created_at bigint NOT NULL
);


ALTER TABLE desafiotecnicoprincipal.accounts OWNER TO postgres;

--
-- TOC entry 210 (class 1259 OID 24662)
-- Name: accounts_id_seq; Type: SEQUENCE; Schema: desafiotecnicoprincipal; Owner: postgres
--

CREATE SEQUENCE desafiotecnicoprincipal.accounts_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE desafiotecnicoprincipal.accounts_id_seq OWNER TO postgres;

--
-- TOC entry 3331 (class 0 OID 0)
-- Dependencies: 210
-- Name: accounts_id_seq; Type: SEQUENCE OWNED BY; Schema: desafiotecnicoprincipal; Owner: postgres
--

ALTER SEQUENCE desafiotecnicoprincipal.accounts_id_seq OWNED BY desafiotecnicoprincipal.accounts.id;


--
-- TOC entry 213 (class 1259 OID 24677)
-- Name: transfers; Type: TABLE; Schema: desafiotecnicoprincipal; Owner: postgres
--

CREATE TABLE desafiotecnicoprincipal.transfers (
    id bigint NOT NULL,
    account_origin_id bigint NOT NULL,
    account_destination_id bigint NOT NULL,
    amount bigint NOT NULL,
    created_at bigint NOT NULL,
    CONSTRAINT account_destination_id_min CHECK ((account_destination_id > 0)),
    CONSTRAINT account_origin_id_min CHECK ((account_origin_id > 0)),
    CONSTRAINT accounts_transfers_check CHECK ((account_origin_id <> account_destination_id)),
    CONSTRAINT amount_min CHECK ((amount > 0)),
    CONSTRAINT created_at_min CHECK ((created_at > 0))
);


ALTER TABLE desafiotecnicoprincipal.transfers OWNER TO postgres;

--
-- TOC entry 212 (class 1259 OID 24676)
-- Name: transfers_id_seq; Type: SEQUENCE; Schema: desafiotecnicoprincipal; Owner: postgres
--

CREATE SEQUENCE desafiotecnicoprincipal.transfers_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE desafiotecnicoprincipal.transfers_id_seq OWNER TO postgres;

--
-- TOC entry 3332 (class 0 OID 0)
-- Dependencies: 212
-- Name: transfers_id_seq; Type: SEQUENCE OWNED BY; Schema: desafiotecnicoprincipal; Owner: postgres
--

ALTER SEQUENCE desafiotecnicoprincipal.transfers_id_seq OWNED BY desafiotecnicoprincipal.transfers.id;


--
-- TOC entry 3170 (class 2604 OID 24666)
-- Name: accounts id; Type: DEFAULT; Schema: desafiotecnicoprincipal; Owner: postgres
--

ALTER TABLE ONLY desafiotecnicoprincipal.accounts ALTER COLUMN id SET DEFAULT nextval('desafiotecnicoprincipal.accounts_id_seq'::regclass);


--
-- TOC entry 3175 (class 2604 OID 24680)
-- Name: transfers id; Type: DEFAULT; Schema: desafiotecnicoprincipal; Owner: postgres
--

ALTER TABLE ONLY desafiotecnicoprincipal.transfers ALTER COLUMN id SET DEFAULT nextval('desafiotecnicoprincipal.transfers_id_seq'::regclass);


--
-- TOC entry 3182 (class 2606 OID 24702)
-- Name: accounts accounts_cpf_key; Type: CONSTRAINT; Schema: desafiotecnicoprincipal; Owner: postgres
--

ALTER TABLE ONLY desafiotecnicoprincipal.accounts
    ADD CONSTRAINT accounts_cpf_key UNIQUE (cpf);


--
-- TOC entry 3184 (class 2606 OID 24670)
-- Name: accounts accounts_pkey; Type: CONSTRAINT; Schema: desafiotecnicoprincipal; Owner: postgres
--

ALTER TABLE ONLY desafiotecnicoprincipal.accounts
    ADD CONSTRAINT accounts_pkey PRIMARY KEY (id);


--
-- TOC entry 3171 (class 2606 OID 24709)
-- Name: accounts created_at_min; Type: CHECK CONSTRAINT; Schema: desafiotecnicoprincipal; Owner: postgres
--

ALTER TABLE desafiotecnicoprincipal.accounts
    ADD CONSTRAINT created_at_min CHECK ((created_at > 0)) NOT VALID;


--
-- TOC entry 3172 (class 2606 OID 24711)
-- Name: accounts go_cpf_nil; Type: CHECK CONSTRAINT; Schema: desafiotecnicoprincipal; Owner: postgres
--

ALTER TABLE desafiotecnicoprincipal.accounts
    ADD CONSTRAINT go_cpf_nil CHECK ((cpf <> ''::text)) NOT VALID;


--
-- TOC entry 3173 (class 2606 OID 24710)
-- Name: accounts go_name_nil; Type: CHECK CONSTRAINT; Schema: desafiotecnicoprincipal; Owner: postgres
--

ALTER TABLE desafiotecnicoprincipal.accounts
    ADD CONSTRAINT go_name_nil CHECK ((name <> ''::text)) NOT VALID;


--
-- TOC entry 3174 (class 2606 OID 24712)
-- Name: accounts go_secret_nil; Type: CHECK CONSTRAINT; Schema: desafiotecnicoprincipal; Owner: postgres
--

ALTER TABLE desafiotecnicoprincipal.accounts
    ADD CONSTRAINT go_secret_nil CHECK ((secret <> ''::text)) NOT VALID;


--
-- TOC entry 3186 (class 2606 OID 24682)
-- Name: transfers transfers_pkey; Type: CONSTRAINT; Schema: desafiotecnicoprincipal; Owner: postgres
--

ALTER TABLE ONLY desafiotecnicoprincipal.transfers
    ADD CONSTRAINT transfers_pkey PRIMARY KEY (id);


-- Completed on 2022-06-21 16:36:57

--
-- PostgreSQL database dump complete
--

