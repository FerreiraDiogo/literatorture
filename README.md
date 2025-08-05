# literatorture
Literatorture is a simple dictionary app running on terminal built for learning purposes

## Requisitos Funcionais

### 1. Map palavra->definição com operações CRUD

**Detalhamento:**

- **Create:** Adicionar nova palavra com definição
- **Read:** Buscar definição de uma palavra específica
- **Update:** Modificar definição existente
- **Delete:** Remover palavra do dicionário

### 2. Busca por prefixo e sufixo

**Detalhamento:**

- Implementar função que retorna todas as palavras que começam com um prefixo específico
- Implementar função que retorna todas as palavras que terminam com um sufixo específico

### 3. Contagem de palavras por letra inicial

**Detalhamento:**

- Função que retorna um map[string]int com estatísticas por letra
- Exibir estatísticas de forma organizada (A-Z)

### 4. Exportar/importar dicionário (JSON/CSV)

**Detalhamento:**

- Função para salvar dicionário em formato JSON estruturado
- Função para salvar em CSV (palavra, definição)
- Funções correspondentes para importar de ambos os formatos

## Requisitos Não Funcionais

### 1. Usar maps nativos do Go

**Detalhamento:**

- Estrutura principal: `map[string]string`
- Índices auxiliares para buscas otimizadas

### 2. Busca case-insensitive

**Detalhamento:**

- Normalizar todas as chaves para lowercase
- Manter versão original para exibição

### 3. Suportar até 10,000 palavras

**Detalhamento:**

- Validar limite durante inserção
- Implementar paginação para exibição

### 4. Interface de linha de comando amigável

**Detalhamento:**

- Menu interativo com opções numeradas
- Mensagens de erro claras
- Help/ajuda para cada comando