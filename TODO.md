# TODO

## Itens

### 1. Estrutura Inicial do Projeto
- [ ] Criar a estrutura de diretórios do projeto (handlers, models, services, etc.).

### 2. Configurar o Servidor Web
- [ ] Escolher e configurar um framework web para Golang (Gin, Echo, etc.).
- [ ] Implementar um endpoint de webhook que possa receber requests POST do GitHub.

### 3. Autenticação de Webhooks
- [ ] Implementar a verificação de assinaturas de webhook para garantir que os eventos são enviados pelo GitHub.

### 4. Processamento de Webhooks
- [ ] Desenvolver a lógica para parsear o payload do webhook e extrair informações necessárias (branch, commits, etc.).
- [ ] Definir uma estrutura de dados para mapear as configurações do arquivo YAML.

### 5. Integração com o Sistema de Arquivos
- [ ] Implementar funcionalidade para ler e parsear arquivos YAML.
- [ ] Desenvolver um sistema de mapeamento das configurações YAML para as ações de CI/CD.

### 6. Execução de Tarefas
- [ ] Criar funções para executar comandos shell baseados nas regras definidas pelo arquivo YAML.
- [ ] Implementar a lógica para lidar com diferentes branches e comandos de forma condicional.

### 7. CLI (Interface de Linha de Comando)
- [ ] Desenvolver uma CLI para interação básica com o sistema (start, stop, status, logs).

### 8. Interface Web (Opcional)
- [ ] Se necessário, desenvolver uma interface web básica para visualizar o status das tarefas e configurações.

### 9. Logging e Monitoramento
- [ ] Configurar um sistema de logging para registrar as atividades do sistema.
- [ ] Opcionalmente, integrar com ferramentas de monitoramento externas.

### 10. Testes
- [ ] Escrever testes unitários e de integração.
- [ ] Configurar um ambiente de teste para simular os webhooks do GitHub.

### 11. Documentação
- [ ] Documentar o projeto, incluindo configuração, uso e troubleshooting.

### 12. Deployment
- [ ] Configurar o ambiente de produção ou staging.
- [ ] Automatizar o processo de deployment, se possível, usando ferramentas como Docker, Kubernetes, ou simples scripts.

### 13. Configuração como Serviço
- [ ] Criar e configurar um arquivo de serviço systemd para permitir que a aplicação rode como um serviço.


## Estrutura

[Desenho tree](tree.nathanfriend.com)

```plaintext
.
├── LICENSE
├── README.md
└── TODO.md/
    └── api/
        ├── main.go
        ├── router/
        │   └── router.go
        └── handler/
            └── action.go
```

