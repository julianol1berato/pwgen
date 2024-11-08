# 🔐 Password Generator | 9Level Network

Um gerador de senhas seguro e elegante com interface inspirada no ParrotOS Security. Gere senhas fortes com facilidade, escolhendo entre diferentes tipos de caracteres e comprimentos personalizados.

![License](https://img.shields.io/badge/license-MIT-green)
![Docker](https://img.shields.io/badge/Docker-ready-blue)
![Go Version](https://img.shields.io/badge/Go-1.21-00ADD8)

## 🚀 Características

- Interface moderna inspirada no ParrotOS Security
- Geração de senhas de 7 a 100 caracteres
- Opções personalizáveis:
  - Letras maiúsculas
  - Letras minúsculas
  - Números
  - Símbolos especiais
- Cópia para área de transferência com um clique
- Backend robusto em Go
- Containerizado com Docker
- Proxy reverso com Traefik
- Design responsivo

## 🛠️ Tecnologias Utilizadas

- Backend: Go 1.21
- Frontend: HTML5, CSS3, JavaScript
- Container: Docker & Docker Compose
- Proxy Reverso: Traefik v3.2.0
- UI/UX: Inspirado no ParrotOS Security

## 📦 Instalação

1. Clone o repositório
```bash
git clone https://github.com/julianol1berato/pwgen.git
cd pwgen
```

2. Configure o DNS
```
Aponte pw.9level.network para o IP do seu servidor
```

3. Execute com Docker Compose
```bash
docker-compose up -d
```

## 🔧 Configuração

O serviço estará disponível em `http://pw.9level.network`

Requisitos do sistema:
- Docker
- Docker Compose
- Porta 80 liberada
- DNS configurado

## 🔒 Segurança

- Geração de senhas totalmente no lado do servidor
- Sem armazenamento de senhas geradas
- Comunicação via HTTP - Uso em conjunto com CloudFlare Tunnels
- Validações tanto no frontend quanto no backend

## 🤝 Contribuindo

1. Faça um Fork do projeto
2. Crie sua Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Commit suas mudanças (`git commit -m 'Add some AmazingFeature'`)
4. Push para a Branch (`git push origin feature/AmazingFeature`)
5. Abra um Pull Request

## 📝 Licença

Este projeto está sob a licença MIT. Veja o arquivo [LICENSE](LICENSE) para mais detalhes.

## ✨ Agradecimentos

- Inspirado no design do ParrotOS Security
- Comunidade Go
- Contribuidores do projeto

## 📬 Contato

Juliano Liberato - [@juliano.liberato](https://bio.9level.network)

Link do Projeto: [https://github.com/julianol1berato/pwgen](https://github.com/julianol1berato/pwgen)
