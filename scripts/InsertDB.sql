INSERT INTO desafiotecnicoprincipal.accounts(name, cpf, secret, balance, created_at)
	VALUES ('Docker User One', '000.000.000-00', '$2y$14$NW7a5zo2uwF.Jpw9KqiQKe5hD7Zb7LJuDVnRcHup5fZMipGzOGP0W', 1000, 1655841024),
	('Docker User Two', '111.111.111-11', '$2y$14$3IWvF8xcnRrLfPe5r/uJLeSdw1SdyKLvAzBLm68CXG257HJmpW0dO', 2000, 1655841032),
	('Docker User Three', '123.123.123-00', '$2y$14$gRlCa2RtXKDe8S249SrIj.n49CZOozFnWg0u8dLr1GEjUL5a96/p.', 123000, 1655841042),
	('Docker User Four', '423.157.174-76', '$2y$14$l/9J76rGUhAZwZtGquMnxeuxOIdHXmoXv8VJwYvYQZP4qGrvM5XRK', 123456, 1655841076);

INSERT INTO desafiotecnicoprincipal.transfers(account_origin_id, account_destination_id, amount, created_at)
	VALUES (1, 2, 10, 1655841124),
	(1, 3, 1, 1655841023),
	(2, 4, 25, 1655841124),
	(2, 1, 123, 1655841023);