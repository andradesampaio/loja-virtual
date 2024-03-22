CREATE TABLE IF NOT EXISTS produto (
  id SERIAL PRIMARY KEY,
  nome VARCHAR(255),
  descricao TEXT,
  preco DECIMAL(10, 2),
  quantidade INT
);

INSERT INTO produto (nome, descricao, preco, quantidade) VALUES
('iPhone 13', 'Smartphone da Apple com 128GB', 6999.00, 50),
('Galaxy S21', 'Smartphone da Samsung com 128GB', 5499.00, 70),
('Macbook Pro', 'Notebook da Apple com 512GB SSD', 13999.00, 30),
('Dell XPS 13', 'Notebook da Dell com 512GB SSD', 8999.00, 40),
('iPad Pro', 'Tablet da Apple com 256GB', 7999.00, 60),
('Camiseta Polo', 'Camiseta Polo de algodão, tamanho M', 79.90, 100),
('Calça Jeans', 'Calça Jeans azul, tamanho 40', 119.90, 75),
('Vestido Floral', 'Vestido Floral tamanho P', 149.90, 50),
('Blusa de Seda', 'Blusa de Seda tamanho M', 89.90, 60),
('Saia Plissada', 'Saia Plissada tamanho G', 99.90, 80);