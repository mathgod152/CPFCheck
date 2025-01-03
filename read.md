# README

## Iniciando o Serviço
Para iniciar o serviço, você deve entrar na pasta /server

```bash
cd /server - linux
cd .\server\ windows 
```

### Agora precisamos buildar a imagem com Docker

Para fazer o build você deve seguir estes passos:


```bash
docker docker build --pull --rm -f "Dockerfile" -t cpfcheck:latest .
```
Lembrando que caso queira mudar o nome da imagem, você deve mudá-la no docker-compose também

Depois de fazer o build da imagem, vamos subir o container:

```
estando na pasta server:

cd /.deloy

docker compose  up 
```

Agora é só utilizar a aplicação na porta :5000

### Service Routes
Todas as rotas da aplicação do lado servidor estão no arquivo “server/manual-test/webtest.http”

# Decisões Arquiteturais:

Servidor: A escolha da arquitetura do servidor foi fundamentada em princípios de Clean Architecture e SOLID, visando garantir baixo acoplamento entre os componentes do sistema. Embora o porte da aplicação não exigisse uma arquitetura tão complexa, optei por adotá-la para demonstrar domínio e aderência a boas práticas de desenvolvimento de software.

Interface de Usuário: Para a camada de interface, foi selecionado o framework Svelte, uma tecnologia com a qual tenho ampla experiência e confiança, o que garante maior eficiência no desenvolvimento e manutenção da aplicação.

# Possíveis Implementações e Melhorias:

A API pode ser aprimorada com a inclusão de testes de validação mais robustos, além dos testes de implementação já existentes. Há uma necessidade de maior foco em validações de dados e no fortalecimento das camadas de segurança da aplicação. Além disso, as variáveis atualmente fixadas devem ser externalizadas para um arquivo de configuração, promovendo uma melhor organização e flexibilidade.

A interface de usuário pode ser otimizada, com a aplicação de boas práticas para garantir um código mais limpo e eficiente. Além disso, há oportunidades de aprimorar o design, buscando uma experiência mais intuitiva e visualmente agradável para os usuários.