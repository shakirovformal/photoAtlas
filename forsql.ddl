-- Table: public.geo_data

-- DROP TABLE IF EXISTS public.geo_data;

CREATE TABLE IF NOT EXISTS public.geo_data
(
    geo_data_id integer NOT NULL DEFAULT nextval('geo_data_geo_data_id_seq'::regclass),
    langitude double precision,
    latitude double precision,
    address_from_yandex text COLLATE pg_catalog."default",
    CONSTRAINT geo_data_pkey PRIMARY KEY (geo_data_id)
)

TABLESPACE pg_default;

ALTER TABLE IF EXISTS public.geo_data
    OWNER to postgres;


CREATE TABLE IF NOT EXISTS public.images
(
    unique_id uuid NOT NULL DEFAULT gen_random_uuid(),
    name_file character varying(100) COLLATE pg_catalog."default" NOT NULL,
    geo_data_id integer,
    CONSTRAINT images_pkey PRIMARY KEY (unique_id),
    CONSTRAINT fk_images_ref_geo_data FOREIGN KEY (geo_data_id)
        REFERENCES public.geo_data (geo_data_id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION
)

TABLESPACE pg_default;

ALTER TABLE IF EXISTS public.images
    OWNER to postgres;	