curl -s https://learn.zone01dakar.sn/assets/superhero/all.json | jq  --argjson id "$HERO_ID" ' .[] | select( .id == $id ) | .connections.relatives' | tr -d '"'