## Projeto de Simulação de Home Broker

O objetivo deste projeto é criar uma simulação de home broker, envolvendo a compra e venda de ações. O foco é utilizar as tecnologias mais atuais para demonstrar minhas habilidades profissionais. Como em qualquer design ou arquitetura, tudo pode ser modificado ao longo do tempo, conforme os trade-offs se tornam mais evidentes.

### Arquitetura Baseada em Eventos

#### Decisões Tomadas

1. **Criação de uma Biblioteca de Utilitários (Utils):**
   - **Motivação:** No Go, muitas vezes precisamos implementar funcionalidades manualmente. Para evitar a recriação de pacotes básicos em cada microserviço, decidi centralizar os códigos mais genéricos em uma biblioteca de utilitários.
   - **Prós:** Evita a duplicação de código de 'baixo nível'.
   - **Contras:** Pode levar a um alto acoplamento e menor coesão entre os componentes.

### Próximos objetivos

1. **Criar o Design Básico:**
   - Definir a estrutura inicial do sistema, incluindo as principais entidades e fluxos.

2. **Desenhar a arquitetura:**
   - Desenhar os eventos básicos do sistema.
     
3. **Implementar o Sistema Básico de Criação/Login de Usuário:**
   - Desenvolver as funcionalidades essenciais para que os usuários possam se cadastrar e fazer login no sistema.

