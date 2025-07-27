# 🖥️ Servidor Estático com Currículo em Go

Este é um projeto simples feito em Go que serve arquivos estáticos em um servidor web local. Ele exibe um site com HTML e CSS, e permite baixar um currículo em PDF.

---

## 🧠 Proposta

Criei este projeto para treinar como um servidor HTTP funciona em Go e, ao mesmo tempo, mostrar meu currículo de forma acessível em um navegador, sem depender de serviços externos.

---

## 🚀 Como usar

### 🔧 Requisitos:
- Go instalado (https://go.dev)
- Navegador Web
- VS Code (opcional, para facilitar edição e execução)

### ▶️ Para rodar:
- Abra a pasta do projeto no VS Code.
- Abra o terminal integrado do VS Code (`Ctrl + ``).
- No terminal, execute o comando:

go run main.go
Abra o navegador e acesse:

Abra o navegador e acesse o endereço exibido no terminal; ele aparecerá assim, por exemplo:
http://localhost:8080

## ⚠️ Aviso do Windows Defender Firewall
Ao rodar o comando:
go run main.go
pela primeira vez, o Windows Defender Firewall pode exibir um alerta informando que o aplicativo main está tentando se comunicar pela rede.

Esse aviso aparece porque o servidor Go está abrindo uma porta local para aceitar conexões HTTP, mesmo que apenas na sua própria máquina (localhost). O Windows interpreta isso como uma tentativa de comunicação de rede e, por segurança, pede sua permissão.

# ✅ O que fazer:
Marque a opção:
Redes privadas, como minha rede doméstica ou corporativa

Clique em: Permitir acesso

Importante: Não é necessário marcar a opção de redes públicas (como redes de aeroportos ou cafés), a menos que você esteja em um ambiente que realmente exija isso — o que não é comum para desenvolvimento local.

Depois disso, o servidor continuará funcionando normalmente e você poderá acessar o projeto pelo navegador no endereço informado no terminal.

## 🗂️ Estrutura de arquivos
Servidor Estatico/
├── main.go              → Código do servidor
├── public/
│   ├── index.html       → Página HTML
│   ├── style.css        → Estilo da página
│   └── go.png           → Imagem ilustrativa
├── main.go              → Arquivo principal do programa em Go, que inicia o servidor

## 💡 Exemplo de uso
Após iniciar o servidor, você pode abrir seu navegador e acessar o localhost.

Esse localhost carrega os arquivos da pasta public, e o conteúdo que será exibido no seu navegador será o que estiver dentro do arquivo index.html, que já está nessa pasta.

Por exemplo, você pode deixar um título, uma imagem e um botão visível nessa página, de acordo com o conteúdo que estiver no index.html.

Caso deseje personalizar o visual ou o texto exibido, basta editar o arquivo index.html e os arquivos relacionados (como CSS ou imagens) que também ficam na pasta public.

## 📚 Tecnologias usadas
Go para o servidor HTTP
HTML + CSS para a interface
PDF como recurso de download

# 🙋‍♂️ Autor
Paulo Victor
Estudante de programação, desenvolvedor em formação.
📧 pvrioslima@gmail.com
🔗 [LinkedIn](https://www.linkedin.com/in/paulovictor-dev)
🐙 [GitHub](https://github.com/paulovrl0624)