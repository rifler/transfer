create user conn_test2 password 'aA_12345' connection limit 2;
create user conn_test5 password 'aA_12345' connection limit 5;
create user conn_test3 password 'aA_12345' connection limit 3;


CREATE TABLE public.test1(
    id         SERIAL             NOT NULL PRIMARY KEY,
    value      INT
);

CREATE TABLE public.test2(
                      id         SERIAL             NOT NULL PRIMARY KEY,
                      value      INT
);

CREATE TABLE public.test3(
                      id         SERIAL             NOT NULL PRIMARY KEY,
                      value      INT
);

CREATE TABLE public.test4(
                      id         SERIAL             NOT NULL PRIMARY KEY,
                      value      INT
);

CREATE TABLE public.test5(
                      id         SERIAL             NOT NULL PRIMARY KEY,
                      value      INT
);

CREATE TABLE public.test6(
                      id         SERIAL             NOT NULL PRIMARY KEY,
                      value      INT
);

INSERT INTO public.test1(value) SELECT generate_series(1, 1000000);
INSERT INTO public.test2(value) SELECT generate_series(1, 1000000);
INSERT INTO public.test3(value) SELECT generate_series(1, 1000000);
INSERT INTO public.test4(value) SELECT generate_series(1, 1000000);
INSERT INTO public.test5(value) SELECT generate_series(1, 1000000);
INSERT INTO public.test6(value) SELECT generate_series(1, 1000000);

GRANT ALL PRIVILEGES ON SCHEMA public TO conn_test2;
GRANT ALL PRIVILEGES ON SCHEMA public TO conn_test5;
GRANT ALL PRIVILEGES ON SCHEMA public TO conn_test3;

GRANT ALL PRIVILEGES ON ALL TABLES IN SCHEMA public TO conn_test2;
GRANT ALL PRIVILEGES ON ALL TABLES IN SCHEMA public TO conn_test5;
GRANT ALL PRIVILEGES ON ALL TABLES IN SCHEMA public TO conn_test3;