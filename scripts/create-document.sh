#!/bin/bash
NAMES=(
  João
  Ricardo
  Joana
  Larissa
  Isabela
  Otávio
  Jussara
  Agnes
  Tatiana
  Thiago
  Mateus
  Olavo
  Teodoro
  Luis
  Gustavo
)

PROFESSIONS=(
  Estudante
  Programador
  Gerente
  Atendente
)

NAME_INDEX=$((RANDOM % ${#NAMES[@]}))
PROF_INDEX=$((RANDOM % ${#PROFESSIONS[@]}))

DOCUMENT=$(cat <<EOF
{
  "name": "${NAMES[$NAME_INDEX]}",
  "profession": "${PROFESSIONS[$PROF_INDEX]}"
}
EOF
)

curl -XPOST -H 'Content-Type: application/json' -d "$DOCUMENT" localhost:8080/api/v1/documents
