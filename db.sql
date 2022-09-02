CREATE TABLE `livros` (
  `id` int(6) unsigned NOT NULL AUTO_INCREMENT,
  `título`  NOT NULL,
  `categoria`  NOT NULL,
  `autor` NOT NULL,
  `sinopse` NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=0 DEFAULT CHARSET=latin1;
