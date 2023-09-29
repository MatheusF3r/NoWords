Bom dia/tarde/noite!

Produzi este script em uma hora vaga, pelo simples fato de nao ter achado meios de verificar um caractere em string em GO.
Não há nada de muito brilhante nisto, e desenvolvi, por pura curiosidade, e necessidade para um projeto futuro.

Segue demais informações.


OPA, LETRA NÃO IDENTIFICADA NA POSIÇÃO EXIGIDA

Precisava verificar se uma string tinha na posição exigida, a letra informada.
Até ai tranquilo, bastava criar um loop para passar letra por letra, e encontra-lá na possição solicitada.

Pra minha surpresa, consegui erro atrás de erro.
Assim, resolvi buscar no Google, funções ou metodos para resolução deste problema,

Me deparei com alguns metodos, como o uso da função LastIndex, do package strings.
Entretanto, ele retorna a última menção do valor solicitado na string.

E qual o problema disto? Se você tem uma string em que não há repetição de caractere, nenhuma!
Exemplo:
Suponhamos que você tenha um dicionário com três palavras:
ana, ama, mar, azul, claro.

Você quer uma palavra de três letras, começando com "a".
O resultado esperado:
>>> ana ; ama

O retorno obtido:
>>> Nulo

Isto se dá, pois ao fazer a verificação o programa identificava que havia letra "a", mas somente no final
palavra, e não no começo. Desta forma, a existência de uma palavra de três letras, começando com "a" era
desconhecida.

Pesquisando mais um pouco, me deparei com a seguinte explicação sobre as strings do Golang:
[https://medium.com/@onezino.moreira/entendendo-string-em-go-c4628055f60d]

Em um resumo... o Golang trata as strings de uma maneira diferente das demais  linguagens de programação.
Em GO uma string é uma slice (mais simples, uma lista) representada diretamente em bytes.
E isto é muito massa! No caso, as demais linguagens tratam as strings como uma lista. Exemplo:

"batata"

A palavra "Batata" seria uma lista assim: [B][a][t][a][t][a].
Neste caso, poderiamos pegar diretamente uma letra na palavra "batata" sem o menor problema.

Mas no Go não ocorre assim. Em GO a mesma lista é assim: [98][97][116][97][116][97].
Cada número representa um caractere em UTF-8. Se quisessemos pegar qualquer letra,
pegariamos na realidade um número.

Então, juntando A com B pra resolver este problema, por que não ao invés de buscar diretamente por letra
não busco por caractere? E esta foi a solução.

Ó código filtra por palavras com a quantidade exigida de letras.
Depois passa letra por letra buscando na posição, o código que bate com o código da letra.



IMPLEMENTAR

Logica baseado em padroes linguisticos que calcule as chances de ter tal palavra.
1 LOGICA - VOGAIS
"TODA CONSOANTE QUE NAO TEM SUPORTE, POSSUI UMA VOGAL EM SEGUIDA"
EX: RAIVA - VOGAL: A; SANTO - VOGAL: A; SONO - VOGAL: O
"TODA CONSOANTE QUE SE ESTA UTILIZANDO DE UM SUPORTE, POSSUI UMA VOGAL EM SEGUIDA"
EX: PSEUDO - SUPORTE: S VOGAL: E; DRAMA - SUPORTE: R VOGAL: A

2 LOGICA - SUPORTE E NAO SUPORTE
SUPORTE É A LETRA QUE ACOMPANHA UMA CONSOANTE, E QUE SE SEPARADA A PALAVRA EM SILABAS PERMANECEM JUNTAS.
*LEMBRANDO QUE NEM TODA LETRA PRECISA DE UM SUPORTE.
H J K M Q R S V W X Y Z - LETRAS QUE NÃO TEM SUPORTE
H R S N - LETRAS QUE PRESTAM SUPORTE, MAS NAO TEM SUPORTE
EX: CHAMA - SUPORTE: H; SOBRA - SUPORTE: R; PNEU - SUPORTE: N
OBS: SEDENTO - NT, MAS SAO SILABAS DISTINTAS. ENTAO NAO HA UTILIZACAO DE SUPORTE

3 LOGICA - MEN (EM OBSERVACAO)
"PALAVRAS COM MEN NA PENULTIMA SILABA TENDEM A TERMINAR COM TO NO FINAL"
CASAMENTO - PAGAMENTO
