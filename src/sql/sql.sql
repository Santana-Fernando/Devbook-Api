-- Criação do banco (rodar separado, se necessário)
CREATE DATABASE devbook;

-- Conectar no banco (via terminal)
-- \c devbook

-- Drop das tabelas (ordem importa por causa das FK)
DROP TABLE IF EXISTS publicacoes;
DROP TABLE IF EXISTS seguidores;
DROP TABLE IF EXISTS usuarios;

-- Tabela usuarios
CREATE TABLE usuarios (
    id SERIAL PRIMARY KEY,
    nome VARCHAR(50) NOT NULL,
    nick VARCHAR(50) NOT NULL UNIQUE,
    email VARCHAR(50) NOT NULL UNIQUE,
    senha VARCHAR(100) NOT NULL,
    criado_em TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Tabela seguidores (relacionamento N:N)
CREATE TABLE seguidores (
    usuario_id INT NOT NULL,
    seguidor_id INT NOT NULL,

    PRIMARY KEY (usuario_id, seguidor_id),

    CONSTRAINT fk_usuario
        FOREIGN KEY (usuario_id)
        REFERENCES usuarios(id)
        ON DELETE CASCADE,

    CONSTRAINT fk_seguidor
        FOREIGN KEY (seguidor_id)
        REFERENCES usuarios(id)
        ON DELETE CASCADE
);

-- Tabela publicacoes
CREATE TABLE publicacoes (
    id SERIAL PRIMARY KEY,
    titulo VARCHAR(50) NOT NULL,
    conteudo VARCHAR(300) NOT NULL,
    autor_id INT NOT NULL,
    curtidas INT DEFAULT 0,
    criada_em TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

    CONSTRAINT fk_autor
        FOREIGN KEY (autor_id)
        REFERENCES usuarios(id)
        ON DELETE CASCADE
);