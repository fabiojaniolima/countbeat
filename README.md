# Countbeat

Este Beat foi desenvolvido considerando a necessidade de quantificar o total de arquivos dado um diretório no sistema operacional. No estado atual o Beat irá quantificar o total de arquivos considerando para tanto a máscara (match) passado por meio da definição do campo `path` no arquivo `countbeat.yml`.

> As palavras-chave "DEVE", "NÃO DEVE", "REQUER", "DEVERIA", "NÃO DEVERIA", "PODERIA", "NÃO PODERIA", "RECOMENDÁVEL", "PODE", e "OPCIONAL" neste documento devem ser interpretadas como descritas no [RFC 2119](http://tools.ietf.org/html/rfc2119). Tradução livre [RFC 2119 pt-br](http://rfc.pt.webiwg.org/rfc2119).

Versão: baseado no [elastic/beats v7.7](https://github.com/elastic/beats/tree/7.7)

## Considerações

- Você PODE realizar o download dos binários para seu respectivo sistema operacional ou realizar sua própria compilação.

### Pacote pré-compilado

Realize o download do arquivo `*bin.zip` referente ao seu sistema operacional em: https://github.com/fabiojaniolima/countbeat/releases

### Realizar compilação

Pré-requisitos:
- `go 1.13.10`
- `git`
- `Python 3.7 ou superior`
- `python virtualenv`

> É importante que as variáveis `GOROOT` e `GOPATH` estejam corretamente definidas, bem como o mapeamento do subdiretório `bin`. Para facilitar a vida considere utilizar o [GVM]( https://github.com/moovweb/gvm) se possível.

Caso não tenha o "Mage" instalado, você DEVE realizar a instalação do mesmo:
```
go get github.com/magefile/mage
```
> Caso esteja utilizando o GVM o binário será automaticamente incluído na sua variável PATH. Do contrário registre o binário manualmente.

```
mkdir -p $GOPATH/src/github.com/fabiojaniolima
git clone https://github.com/fabiojaniolima/countbeat.git $GOPATH/src/github.com/fabiojaniolima/countbeat
```

Baixe as dependências:
```
mage vendorUpdate
```

Uma das formas de realizar a compilação é executar:
```
go build -o countbeat main.go
```

> É RECOMENDÁVEL executar `mage update` e `mage fmt` sempre que alterar arquivos nos diretórios: `_meta` e `config`.

Para realizar uma compilação cruzada, ou seja, para outro sistema ou arquitetura considere utilizar:
```
env GOOS=windows GOARCH=amd64 go build -o countbeat.exe main.go
```

## Utilizando

O Countbeat possui um funcionamento similar a qualquer outro Beat que faz parte da Stack ELK. Basicamente você configura o arquivo `countbeat.yml` e chama o binário de nome `countbeat` que foi baixado ou compilado.

A contagem dos arquivos será armazenada no campo `counter`, este valor será enviado para o output definido no arquivo de configuração, ou seja, segue o comportamento padrão de qualquer Beat da família.

## Referências

Alguns links relevantes:
- https://www.elastic.co/guide/en/beats/devguide/current/new-beat.html
- https://www.elastic.co/guide/en/beats/devguide/current/beater-interface.html
