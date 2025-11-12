# 🔍 Monitoramento de Sites

Sistema de monitoramento HTTP desenvolvido em Go para verificar a disponibilidade e performance de websites em tempo real.

## 🎯 Sobre o Projeto

Aplicação CLI (Command Line Interface) que realiza requisições HTTP periódicas para monitorar o status de sites, registrando tempo de resposta e disponibilidade. Ideal para DevOps, SRE e desenvolvedores que precisam garantir uptime de aplicações.

## ⚡ Tecnologias Utilizadas

- **Go (Golang)** - Linguagem de programação
- **HTTP Client nativo** - Requisições HTTP
- **Goroutines** - Concorrência e paralelismo
- **JSON** - Persistência de configurações

## ✨ Funcionalidades

- ✅ Monitoramento contínuo de múltiplos sites simultaneamente
- ✅ Verificação de status HTTP (200, 404, 500, etc)
- ✅ Medição de tempo de resposta
- ✅ Registro de histórico de disponibilidade
- ✅ Alertas quando sites ficam fora do ar
- ✅ Execução concorrente usando Goroutines

## 🚀 Como Usar

### Pré-requisitos

- [Go 1.19+](https://golang.org/dl/) instalado

### Instalação

```bash
# Clone o repositório
git clone https://github.com/YuriLuiz1/monitoramento-sites.git
cd monitoramento-sites

# Baixe as dependências
go mod download

# Compile o projeto
go build -o monitor
```

### Configuração

Crie um arquivo `sites.txt` com os sites que deseja monitorar:

```txt **Nesse mesmo formato abaixo**
https://www.google.com.br
https://www.youtube.com.br
```

### Execução

```bash
# Execute o monitoramento
./monitor

# Ou diretamente com Go
go run main.go
```

## 📊 Exemplo de Saída

```
[2024-10-28 14:30:15] ✅ https://google.com - Status: 200 - Tempo: 145ms
[2024-10-28 14:30:15] ✅ https://github.com - Status: 200 - Tempo: 289ms
[2024-10-28 14:30:16] ❌ https://siteindisponivel.com - Status: 0 - ERRO: timeout

```
monitoramento-sites/
├── main.go              # Ponto de entrada
├── monitor/
│   ├── checker.go       # Lógica de verificação HTTP
│   ├── logger.go        # Sistema de logs
│   └── config.go        # Gerenciamento de configurações
├── sites.json           # Lista de sites para monitorar
├── logs/                # Diretório de logs históricos
└── go.mod
```

## 🔧 Conceitos de Go Aplicados

Este projeto demonstra o uso de:

- **HTTP Client** - Requisições HTTP com timeout personalizado
- **Error Handling** - Tratamento robusto de erros
- **Time e Ticker** - Execução periódica
- **Logging** - Registro estruturado de eventos

## 📈 Melhorias Futuras

- [ ] Interface web para visualização em tempo real
- [ ] Notificações por email/Slack/Discord
- [ ] Métricas de uptime (99.9%, 99.99%)
- [ ] Gráficos de tempo de resposta
- [ ] Suporte a autenticação HTTP
- [ ] Docker container para facilitar deploy
- [ ] Exportação de métricas para Prometheus

## 💡 Por que Go?

Go foi escolhido para este projeto por:

- **Performance**: Requisições HTTP concorrentes extremamente rápidas
- **Concorrência nativa**: Goroutines facilitam monitoramento paralelo
- **Binário único**: Deploy simplificado sem dependências
- **Baixo consumo de recursos**: Ideal para rodar 24/7

## 🎓 Aprendizados

Desenvolvi este projeto para:
- Trabalhar com HTTP clients e timeouts
- Implementar sistemas de logging
- Criar aplicações CLI úteis e práticas

## 👤 Autor

**Yuri Luiz**

- GitHub: [@YuriLuiz1](https://github.com/YuriLuiz1)
- LinkedIn: [https://www.linkedin.com/in/yuri-luiz/]

## 📄 Licença

Este projeto está sob a licença MIT.

---

💚 Desenvolvido com Go | ⭐ Deixe uma estrela se foi útil!
