# Uptime Monitor (Go)

Um mini serviço de **Uptime Monitor** desenvolvido em **Go**, com foco em aprendizado prático da linguagem, organização de projeto e uso de concorrência básica.

Este projeto permite cadastrar URLs e acompanhar se elas estão **UP** ou **DOWN**, servindo como base para entender como serviços backend reais funcionam em Go.

---

## 🧠 O que o serviço faz

- Sobe um servidor HTTP
- Permite cadastrar URLs para monitoramento
- Armazena os checks em memória
- Expõe endpoints para consulta
- Possui endpoint de health check

---

## ➕ Como funciona a criação de um check

A criação de um check acontece através do endpoint `POST /checks` e segue um fluxo simples e bem definido.

### Fluxo de criação

1. **Recebimento da requisição**
   - A API recebe um JSON contendo a URL a ser monitorada.
   - Exemplo:
     ```json
     {
       "url": "https://google.com"
     }
     ```

2. **Geração de um ID único**
   - Um identificador é gerado para representar unicamente o check.

3. **Criação do objeto de domínio**
   - O check é criado com os seguintes campos:
     - `ID`: identificador único
     - `URL`: URL informada
     - `Status`: inicializado como `DOWN`
     - `LastChecked`: valor zero (`0001-01-01T00:00:00Z`)

4. **Armazenamento em memória**
   - O check é salvo em memória.
   - O acesso é protegido com `sync.Mutex` para evitar problemas de concorrência.

5. **Retorno da resposta**
   - A API retorna o check criado em formato JSON.

### Exemplo de resposta

```json
{
  "ID": "a1b2c3d4e5f6",
  "URL": "https://google.com",
  "Status": "DOWN",
  "LastChecked": "0001-01-01T00:00:00Z"
}