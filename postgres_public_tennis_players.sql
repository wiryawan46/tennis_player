create table tennis_players
(
    id          serial not null
        constraint tennis_players_pk
            primary key,
    player_name text
);

alter table tennis_players
    owner to postgres;

create index tennis_players_id_index
    on tennis_players (id);

INSERT INTO public.tennis_players (id, player_name) VALUES (1, 'John');
INSERT INTO public.tennis_players (id, player_name) VALUES (2, 'Chris');
INSERT INTO public.tennis_players (id, player_name) VALUES (3, 'Niki');
INSERT INTO public.tennis_players (id, player_name) VALUES (4, 'Herald');
INSERT INTO public.tennis_players (id, player_name) VALUES (5, 'Chris');
INSERT INTO public.tennis_players (id, player_name) VALUES (6, 'Mike');
INSERT INTO public.tennis_players (id, player_name) VALUES (7, 'Richard');
INSERT INTO public.tennis_players (id, player_name) VALUES (8, 'Hana');
INSERT INTO public.tennis_players (id, player_name) VALUES (9, 'Joe');
INSERT INTO public.tennis_players (id, player_name) VALUES (10, 'Martha');