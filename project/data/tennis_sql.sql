-- Enable the uuid-ossp extension
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

INSERT INTO player (name, age, country) VALUES 
('Roger Federer', 39, 'Switzerland'),
('Rafael Nadal', 35, 'Spain'),
('Novak Djokovic', 34, 'Serbia'),
('Andy Murray', 34, 'United Kingdom');

INSERT INTO court (id, location, surface) VALUES 
(uuid_generate_v4(), 'All England Lawn Tennis and Croquet Club', 'Grass'),
(uuid_generate_v4(), 'Stade Roland Garros', 'Clay');

INSERT INTO referee (name, country) VALUES 
('Carlos Ramos', 'Portugal'),
('Eva Asderaki', 'Greece');

INSERT INTO match (player1_id, player2_id, court_id, referee_id, date) VALUES 
((SELECT id FROM player WHERE name = 'Roger Federer'), (SELECT id FROM player WHERE name = 'Rafael Nadal'), (SELECT id FROM court WHERE location = 'All England Lawn Tennis and Croquet Club'), (SELECT id FROM referee WHERE name = 'Carlos Ramos'), '2022-07-03 14:00:00'),
((SELECT id FROM player WHERE name = 'Novak Djokovic'), (SELECT id FROM player WHERE name = 'Andy Murray'), (SELECT id FROM court WHERE location = 'Stade Roland Garros'), (SELECT id FROM referee WHERE name = 'Eva Asderaki'), '2022-06-05 14:00:00');

INSERT INTO score (match_id, player_id, set1, set2, set3) VALUES 
((SELECT id FROM match WHERE player1_id = (SELECT id FROM player WHERE name = 'Roger Federer') AND player2_id = (SELECT id FROM player WHERE name = 'Rafael Nadal')), (SELECT id FROM player WHERE name = 'Roger Federer'), 6, 4, 7),
((SELECT id FROM match WHERE player1_id = (SELECT id FROM player WHERE name = 'Roger Federer') AND player2_id = (SELECT id FROM player WHERE name = 'Rafael Nadal')), (SELECT id FROM player WHERE name = 'Rafael Nadal'), 4, 6, 5),
((SELECT id FROM match WHERE player1_id = (SELECT id FROM player WHERE name = 'Novak Djokovic') AND player2_id = (SELECT id FROM player WHERE name = 'Andy Murray')), (SELECT id FROM player WHERE name = 'Novak Djokovic'), 6, 6, 6),
((SELECT id FROM match WHERE player1_id = (SELECT id FROM player WHERE name = 'Novak Djokovic') AND player2_id = (SELECT id FROM player WHERE name = 'Andy Murray')), (SELECT id FROM player WHERE name = 'Andy Murray'), 4, 4, 4);
